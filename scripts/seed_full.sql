-- ============================================================================
--  seed_full.sql  —  Complete demo-data rebuild for the clinic app  (v4)
-- ----------------------------------------------------------------------------
--  Run on the server:
--     docker cp seed_full.sql clinic-db-1:/seed_full.sql
--     docker exec -i clinic-db-1 psql -U clinic -d clinic -f /seed_full.sql
--
--  What it does (one transaction, fully repeatable):
--    1. Adds the is_edited / is_deleted columns if the backend hasn't migrated yet.
--    2. DELETES every order, order item and worker stock  (=> analytics reset).
--    3. Resets each product to 60 000 штук = 1000 «капс» (packs of 60).
--    4. Generates ~12 MONTHS of REALISTIC orders, ~300-400 mln sum/month revenue.
--    5. Recomputes warehouse stock as a CLOSED system so analytics reconciles.
--
--  STATUS MODEL (final, 4 lifecycle states + 2 admin-only overlays):
--    • pending     «Ожидает»       — ONLINE only, start of the online flow
--    • in_transit  «В пути»         — ONLINE (cashier dispatched it)
--    • delivered   «Выдано»         — ONLINE *and* OFFLINE terminal state
--    • cancelled   «Отменено»       — ONLINE and OFFLINE (= «Отклонено», merged)
--    • is_edited   «Редактировано»  — overlay, RED border, admin panel only
--    • is_deleted  «Удалено»        — overlay, BLACK border, admin panel only,
--                                     hidden from pickup / customer / analytics
--    (the old «confirmed»/«shipped» states are gone)
--
--  UNITS:
--    • 1 «капс» = 1 упаковка = 60 капсул (штук). 6500 сум/штука, 390 000/упаковка.
--    • products.stock_quantity is stored in ШТУК (individual capsules/pieces).
--    • ONLINE sales are always WHOLE PACKS. OFFLINE channels (offline walk-in,
--      doctor, vip/own-patient, marketolog) may sell loose PIECES («штучно»), so
--      stock is no longer forced to a multiple of 60 — but it stays EXACT in штук
--      and «На складе» + «Продано» still reconcile one-to-one.
--
--  ACCOUNTS (use whatever exists in the DB — never create or recreate any):
--    • pickup points (role='pickup')  — orders split 50/50 between them
--    • marketolog(s) (role='manager') — marketolog sales assigned round-robin
--    • doctor = first doctors row;  users = the users table (online buyers)
-- ============================================================================

\set ON_ERROR_STOP on

BEGIN;

-- 1. Make sure the new overlay columns exist even if the backend hasn't migrated.
ALTER TABLE orders ADD COLUMN IF NOT EXISTS is_edited       boolean DEFAULT false;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS is_deleted      boolean DEFAULT false;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS deleted_reason  text    DEFAULT '';
ALTER TABLE orders ADD COLUMN IF NOT EXISTS deleted_by_name text    DEFAULT '';
ALTER TABLE orders ADD COLUMN IF NOT EXISTS deleted_by_role text    DEFAULT '';

-- 2. Full wipe of orders / analytics / stock counters
DELETE FROM order_items;
DELETE FROM orders;
DELETE FROM worker_stocks;

-- 3. Reset every product to 60 000 штук (= 1000 капс). Mark conversion done so the
--    backend's one-time capsules→pieces migration never multiplies these again.
UPDATE products SET stock_quantity = 60000, stock_converted = true;

-- 4. Generate realistic orders
DO $SEED$
DECLARE
  v_counter  BIGINT := 0;
  v_order_id BIGINT;
  v_pid INT; v_qty INT; v_unit TEXT; v_price NUMERIC;
  v_status TEXT; v_offline BOOLEAN; v_pay TEXT;
  v_user INT; v_worker INT; v_doctor INT; v_mkt INT;
  v_vip BOOLEAN; v_nurse BOOLEAN; v_refby TEXT; v_channel TEXT;
  v_fname TEXT; v_lname TEXT; v_addr TEXT; v_phone TEXT; v_code TEXT;
  v_cr TEXT; v_cb TEXT; v_crl TEXT; v_ts TIMESTAMPTZ;
  v_otype TEXT; v_rand FLOAT; v_day_rev NUMERIC; v_target NUMERIC;
  v_secs INT; v_hour INT; d INT; guard INT;
  v_edited BOOLEAN; v_deleted BOOLEAN; v_del_reason TEXT; v_del_name TEXT;
  v_wname TEXT;

  mkt_ids   INT[];
  pick_ids  INT[];
  pick_name TEXT[];
  pids INT[];
  uids INT[];
  fn   TEXT[] := ARRAY['Ali','Zulfiya','Bobur','Malika','Jasur','Nilufar','Sherzod','Dildora',
                        'Ulugbek','Feruza','Hamid','Sarvinoz','Otabek','Kamola','Ravshan',
                        'Shohruh','Mohira','Behruz','Nargiza','Dilshod'];
  ln   TEXT[] := ARRAY['Karimov','Xasanova','Raximov','Yusupova','Toshmatov','Mirzayeva',
                        'Abdullayev','Normatova','Sultonov','Ergasheva','Hasanov','Qodirov',
                        'Ismoilov','Nazarov','Yunusov'];
  addrs TEXT[] := ARRAY[
    'Toshkent, Yunusobod t., Musajon ko''chasi 12',
    'Toshkent, Chilonzor t., 7-mavze, 45-uy',
    'Toshkent, Mirzo Ulugbek t., Qoratosh 8',
    'Toshkent, Sergeli t., Yangi hayot 33',
    'Samarqand sh., Registon ko''chasi 12',
    'Buxoro v., Qorovulbozor, 5-uy',
    'Namangan sh., Mustaqillik ko''chasi 8',
    'Fargona v., Margilon sh., 3-kvartal',
    'Andijon sh., 25-kvartal, 3-uy',
    'Qashqadaryo, Shahrisabz, 17-uy'];
  chans TEXT[] := ARRAY['telegram','instagram','facebook','whatsapp'];
  del_reasons TEXT[] := ARRAY[
    'Xato kiritilgan buyurtma',
    'Mijoz boshqa buyurtma berdi, dublikat',
    'Test buyurtma, o''chirildi',
    'Noto''g''ri mahsulot tanlangan'];

  PRICE_PILL NUMERIC := 6500;
  QPP INT := 60;
  DOC_ID   INT;      -- referring doctor
  SPAN INT := 364;   -- 0..364 => 365 days (12 months)
  BASE TIMESTAMPTZ := '2025-06-01 00:00:00+05';
BEGIN
  PERFORM setseed(0.4242);

  SELECT array_agg(id ORDER BY id) INTO pids FROM products;
  IF pids IS NULL THEN RAISE EXCEPTION 'No products to sell'; END IF;

  -- marketolog (manager) accounts — assign sales round-robin to whoever exists.
  SELECT array_agg(id ORDER BY id) INTO mkt_ids FROM workers WHERE role='manager';

  -- pickup points (role='pickup') — orders split evenly between them.
  SELECT array_agg(id ORDER BY id), array_agg(name ORDER BY id)
    INTO pick_ids, pick_name FROM workers WHERE role='pickup';
  IF pick_ids IS NULL THEN RAISE EXCEPTION 'No pickup workers'; END IF;

  SELECT id INTO DOC_ID FROM doctors ORDER BY id LIMIT 1;

  SELECT array_agg(id ORDER BY id) INTO uids FROM users;
  IF uids IS NULL THEN uids := ARRAY[NULL]::INT[]; END IF;

  FOR d IN 0..SPAN LOOP
    v_target  := 10000000 + random()*3000000;   -- 10-13 mln main revenue / day  (~300-400 mln/month)
    v_day_rev := 0;
    guard     := 0;

    WHILE v_day_rev < v_target AND guard < 300 LOOP
      guard := guard + 1;
      v_counter := v_counter + 1;

      v_doctor := NULL; v_mkt := NULL; v_vip := false; v_nurse := false;
      v_refby := ''; v_channel := ''; v_fname := ''; v_lname := '';
      v_addr := ''; v_worker := NULL; v_offline := false; v_user := NULL;
      v_cr := ''; v_cb := ''; v_crl := '';
      v_edited := false; v_deleted := false; v_del_reason := ''; v_del_name := '';

      v_pid := pids[(floor(random()*array_length(pids,1))+1)::INT];

      -- pick a pickup point for this transaction (50/50 across the two)
      v_worker := pick_ids[(floor(random()*array_length(pick_ids,1))+1)::INT];
      v_wname  := pick_name[(array_position(pick_ids, v_worker))];

      v_rand := random();
      IF    v_rand < 0.28 THEN v_otype := 'online';
      ELSIF v_rand < 0.55 THEN v_otype := 'offline';
      ELSIF v_rand < 0.68 THEN v_otype := 'doctor';
      ELSIF v_rand < 0.76 THEN v_otype := 'vip';
      ELSE                     v_otype := 'marketolog';
      END IF;

      -- UNIT & QUANTITY.
      --   ONLINE  → always whole packs (1-3).
      --   OFFLINE channels → 30% loose pieces (6-59 шт), otherwise whole packs.
      IF v_otype = 'online' THEN
        v_unit := 'pack';
        v_rand := random();
        IF    v_rand < 0.65 THEN v_qty := 1;
        ELSIF v_rand < 0.90 THEN v_qty := 2;
        ELSE                     v_qty := 3; END IF;
        v_price := v_qty * PRICE_PILL * QPP;
      ELSE
        IF random() < 0.30 THEN
          v_unit := 'piece';
          v_qty  := 6 + floor(random()*54)::INT;     -- 6..59 loose capsules
          v_price := v_qty * PRICE_PILL;
        ELSE
          v_unit := 'pack';
          v_rand := random();
          IF    v_rand < 0.65 THEN v_qty := 1;
          ELSIF v_rand < 0.90 THEN v_qty := 2;
          ELSE                     v_qty := 3; END IF;
          v_price := v_qty * PRICE_PILL * QPP;
        END IF;
      END IF;

      CASE v_otype
        WHEN 'online' THEN
          v_offline := false;
          v_user := uids[(floor(random()*array_length(uids,1))+1)::INT];
          v_addr := addrs[(floor(random()*array_length(addrs,1))+1)::INT];
          IF random() < 0.55 THEN v_pay := 'online'; ELSE v_pay := 'card'; END IF;
          v_rand := random();
          IF    v_rand < 0.12 THEN v_status := 'pending';   v_worker := NULL;  -- not yet handled
          ELSIF v_rand < 0.22 THEN v_status := 'cancelled';
                v_cr := 'Mijoz buyurtmani bekor qildi'; v_cb := 'Mijoz'; v_crl := 'user';
          ELSIF v_rand < 0.42 THEN v_status := 'in_transit';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'offline' THEN
          v_offline := true;
          IF random() < 0.25 THEN v_user := NULL;
          ELSE v_user := uids[(floor(random()*array_length(uids,1))+1)::INT]; END IF;
          v_rand := random();
          IF    v_rand < 0.55 THEN v_pay := 'cash';
          ELSIF v_rand < 0.85 THEN v_pay := 'terminal';
          ELSE                     v_pay := 'card'; END IF;
          v_rand := random();
          IF    v_rand < 0.07 THEN v_status := 'cancelled';
                v_cr := 'Mijoz olib ketmadi'; v_cb := v_wname; v_crl := 'pickup';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'doctor' THEN
          v_offline := true; v_doctor := DOC_ID;
          v_refby := 'Karimov'; v_nurse := (random() < 0.50);
          v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
          v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
          IF random() < 0.60 THEN v_pay := 'cash'; ELSE v_pay := 'terminal'; END IF;
          v_rand := random();
          IF    v_rand < 0.07 THEN v_status := 'cancelled';
                v_cr := 'Bemor kelmadi'; v_cb := v_wname; v_crl := 'pickup';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'vip' THEN
          v_offline := true; v_doctor := DOC_ID;
          v_vip := true; v_refby := 'Karimov'; v_nurse := (random() < 0.70);
          v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
          v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
          v_pay := ''; v_price := 0;                 -- own patient: free
          v_rand := random();
          IF    v_rand < 0.03 THEN v_status := 'cancelled';
                v_cr := 'Bekor qilindi'; v_cb := v_wname; v_crl := 'pickup';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'marketolog' THEN
          v_offline := true; v_worker := NULL;
          IF mkt_ids IS NOT NULL THEN
            v_mkt := mkt_ids[(floor(random()*array_length(mkt_ids,1))+1)::INT];
          END IF;
          v_pay := 'marketplace';
          v_channel := chans[(floor(random()*array_length(chans,1))+1)::INT];
          IF random() < 0.30 THEN v_user := NULL;
          ELSE v_user := uids[(floor(random()*array_length(uids,1))+1)::INT]; END IF;
          v_rand := random();
          IF    v_rand < 0.10 THEN v_status := 'cancelled';
                v_cr := 'Mijoz rad etdi'; v_cb := 'Marketolog'; v_crl := 'manager';
          ELSE                     v_status := 'delivered'; END IF;
      END CASE;

      IF v_otype IN ('online','offline') AND random() < 0.25 THEN
        v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
        v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
      END IF;

      -- OVERLAYS (admin-only). Only meaningful on non-cancelled orders.
      IF v_status <> 'cancelled' THEN
        -- ~3% of delivered orders were edited at the pickup point (red border).
        IF v_status = 'delivered' AND random() < 0.03 THEN v_edited := true; END IF;
        -- ~1.2% of orders were deleted by a cashier (black border, admin only).
        IF random() < 0.012 THEN
          v_deleted := true;
          v_del_reason := del_reasons[(floor(random()*array_length(del_reasons,1))+1)::INT];
          IF v_worker IS NOT NULL THEN v_del_name := v_wname;
          ELSE v_del_name := pick_name[1]; END IF;
        END IF;
      END IF;

      v_phone := '998' || lpad((floor(random()*1000000000))::BIGINT::TEXT, 9, '0');
      v_code  := upper(lpad(to_hex(v_counter), 6, '0'));
      v_hour  := 8 + (floor(random()*13))::INT;       -- business hours 8..20
      v_secs  := v_hour*3600 + (floor(random()*3600))::INT;
      v_ts    := BASE + (d || ' days')::INTERVAL + (v_secs || ' seconds')::INTERVAL;

      INSERT INTO orders (
        user_id, status, phone, created_at, order_code,
        delivery_address, is_offline, offline_note, latitude, longitude, referred_by,
        worker_id, is_nurse_order, patient_f_name, patient_l_name,
        doctor_id, is_v_ip, payment_method, card_type,
        cancellation_reason, cancelled_by_name, cancelled_by_role,
        sales_channel, is_returned, return_reason,
        marketolog_id, archived, archive_reason,
        is_edited, is_deleted, deleted_reason, deleted_by_name, deleted_by_role
      ) VALUES (
        v_user, v_status, v_phone, v_ts, v_code,
        v_addr, v_offline, '', 0, 0, v_refby,
        v_worker, v_nurse, v_fname, v_lname,
        v_doctor, v_vip, v_pay, '',
        v_cr, v_cb, v_crl,
        v_channel, false, '',
        v_mkt, false, '',
        v_edited, v_deleted, v_del_reason, v_del_name,
        CASE WHEN v_deleted THEN 'pickup' ELSE '' END
      ) RETURNING id INTO v_order_id;

      INSERT INTO order_items
        (order_id, product_id, quantity, unit_type, price, original_quantity)
      VALUES
        (v_order_id, v_pid, v_qty, v_unit, v_price, v_qty);

      -- advance the day's revenue target only for what the main revenue chart shows:
      -- non-cancelled, non-deleted, non-marketolog (vip price 0 adds nothing).
      IF v_status <> 'cancelled' AND NOT v_deleted AND v_mkt IS NULL THEN
        v_day_rev := v_day_rev + v_price;
      END IF;
    END LOOP;
  END LOOP;

  RAISE NOTICE 'Created % orders over % days', v_counter, SPAN + 1;
END $SEED$;

-- 5. CLOSED-SYSTEM stock reconciliation.
--    stock = 60 000 − штук продано, with the SAME filter the product analytics uses
--    (status <> 'cancelled' AND archived = false AND is_deleted = false). No
--    restocking ⇒ «На складе» + «Продано» always add back up to 60 000 штук.
UPDATE products SET stock_quantity = 60000;
UPDATE products p SET stock_quantity = 60000 - c.pieces
FROM (
  SELECT oi.product_id AS pid,
         SUM(CASE WHEN oi.unit_type = 'pack' THEN oi.quantity * 60 ELSE oi.quantity END) AS pieces
  FROM order_items oi
  JOIN orders o ON o.id = oi.order_id
  WHERE o.status <> 'cancelled' AND o.archived = false AND o.is_deleted = false
  GROUP BY oi.product_id
) c WHERE c.pid = p.id;

COMMIT;

-- ── verification ────────────────────────────────────────────────────────────
\echo ''
\echo '=== orders / items ==='
SELECT (SELECT COUNT(*) FROM orders) AS orders, (SELECT COUNT(*) FROM order_items) AS items;
\echo '=== status ==='
SELECT status, COUNT(*) FROM orders GROUP BY status ORDER BY 2 DESC;
\echo '=== overlays (edited / deleted) ==='
SELECT COUNT(*) FILTER (WHERE is_edited)  AS edited,
       COUNT(*) FILTER (WHERE is_deleted) AS deleted FROM orders;
\echo '=== unit split (pack vs piece) ==='
SELECT unit_type, COUNT(*) FROM order_items GROUP BY unit_type ORDER BY 2 DESC;
\echo '=== pickup point split (worker_id) ==='
SELECT worker_id, COUNT(*) FROM orders WHERE worker_id IS NOT NULL GROUP BY worker_id ORDER BY worker_id;
\echo '=== channel ==='
SELECT CASE WHEN marketolog_id IS NOT NULL THEN 'marketolog'
            WHEN is_v_ip THEN 'vip'
            WHEN doctor_id IS NOT NULL THEN 'doctor'
            WHEN is_offline THEN 'offline' ELSE 'online' END AS channel, COUNT(*)
FROM orders GROUP BY 1 ORDER BY 2 DESC;
\echo '=== monthly main revenue (non-cancelled, non-deleted, non-marketolog) ==='
SELECT to_char(date_trunc('month',o.created_at),'YYYY-MM') AS month, round(SUM(oi.price)) AS revenue, COUNT(DISTINCT o.id) AS orders
FROM orders o JOIN order_items oi ON oi.order_id=o.id
WHERE o.status<>'cancelled' AND o.is_deleted=false AND o.marketolog_id IS NULL
GROUP BY 1 ORDER BY 1;
\echo '=== stock reconcile: sklad_kaps + prodano_kaps (whole packs) ==='
WITH consumed AS (
  SELECT oi.product_id pid,
         SUM(CASE WHEN oi.unit_type='pack' THEN oi.quantity*60 ELSE oi.quantity END) sold_pieces
  FROM order_items oi JOIN orders o ON o.id=oi.order_id
  WHERE o.status<>'cancelled' AND o.archived=false AND o.is_deleted=false GROUP BY 1)
SELECT p.id, p.name, p.stock_quantity sklad_sht, COALESCE(c.sold_pieces,0) prodano_sht,
       p.stock_quantity + COALESCE(c.sold_pieces,0) AS summa_sht
FROM products p LEFT JOIN consumed c ON c.pid=p.id ORDER BY p.id LIMIT 5;
