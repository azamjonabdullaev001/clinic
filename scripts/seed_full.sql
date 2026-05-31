-- ============================================================================
--  seed_full.sql  —  Complete demo-data rebuild for the clinic app
-- ----------------------------------------------------------------------------
--  Run on the server:
--     docker cp seed_full.sql clinic-db-1:/seed_full.sql
--     docker exec -i clinic-db-1 psql -U clinic -d clinic -f /seed_full.sql
--
--  What it does (one transaction, fully repeatable):
--    1. Ensures marketolog (role='manager') workers exist — needed so orders
--       are visible in the marketolog panel. Password for the 3 created here: Market2026
--    2. DELETES every order, order item and worker stock  (=> analytics reset).
--    3. Resets each product to 60 000 штук = 1000 «капс» (packs of 60).
--    4. Generates ~6 months of REALISTIC orders, ~8-10 mln sum/day main revenue.
--    5. Recomputes warehouse stock as a CLOSED system:
--         stock = 60 000 − штук продано   (same filter the product analytics uses)
--       so «На складе» + «Продано» reconcile one-to-one.
--
--  Units (as the UI shows them):
--    • 1 «капс» = 1 упаковка = 60 капсул (штук). price 6500 сум/штука, 390 000/упаковка.
--    • products.stock_quantity is stored in ШТУК (individual capsules/pieces).
--    • Analytics «Продано капсул» = number of packs sold; «штучно» = loose pieces.
--
--  Channels & rules:
--    • online      — registered user, delivery, pay online/card; statuses incl. PENDING
--    • offline     — pickup point walk-in (cashier = pickup worker)
--    • doctor      — referred by doctor (paid)
--    • vip         — own patient, FREE (price 0)
--    • marketolog  — social-media sale, assigned to a manager (visible in panel)
--    • PENDING only for online. cancelled & confirmed in BOTH online and offline.
--    • vip & marketolog are OFFLINE only (never online).
-- ============================================================================

\set ON_ERROR_STOP on

-- 1. Marketolog workers (idempotent). Password for all three: Market2026
INSERT INTO workers (name, phone, password, role) VALUES
  ('Dilshod Marketolog', '998900000001', '$2b$10$NGBgHPO0xfdoS7yOOLj3j.f68tNL7H4C1kkVxzCItqYmZW6G3zyfu', 'manager'),
  ('Nigora Marketolog',  '998900000002', '$2b$10$NGBgHPO0xfdoS7yOOLj3j.f68tNL7H4C1kkVxzCItqYmZW6G3zyfu', 'manager'),
  ('Shahzod Marketolog', '998900000003', '$2b$10$NGBgHPO0xfdoS7yOOLj3j.f68tNL7H4C1kkVxzCItqYmZW6G3zyfu', 'manager')
ON CONFLICT (phone) DO NOTHING;

BEGIN;

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

  mkt_ids INT[];
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

  PRICE_PILL NUMERIC := 6500;
  QPP INT := 60;
  W_PICKUP INT;      -- pickup cashier
  DOC_ID   INT;      -- referring doctor
  SPAN INT := 181;   -- 0..181 => 182 days
  BASE TIMESTAMPTZ := '2025-12-01 00:00:00+05';
BEGIN
  PERFORM setseed(0.4242);

  SELECT array_agg(id ORDER BY id) INTO pids FROM products;
  IF pids IS NULL THEN RAISE EXCEPTION 'No products to sell'; END IF;

  SELECT array_agg(id ORDER BY id) INTO mkt_ids FROM workers WHERE role='manager';
  IF mkt_ids IS NULL THEN RAISE EXCEPTION 'No marketolog (manager) workers'; END IF;

  -- pickup cashier: first worker with role 'pickup' (fallback: any worker)
  SELECT id INTO W_PICKUP FROM workers WHERE role='pickup' ORDER BY id LIMIT 1;
  IF W_PICKUP IS NULL THEN SELECT id INTO W_PICKUP FROM workers ORDER BY id LIMIT 1; END IF;

  -- referring doctor (optional)
  SELECT id INTO DOC_ID FROM doctors ORDER BY id LIMIT 1;

  -- registered users for online/offline orders
  SELECT array_agg(id ORDER BY id) INTO uids FROM users;
  IF uids IS NULL THEN uids := ARRAY[NULL]::INT[]; END IF;

  FOR d IN 0..SPAN LOOP
    v_target  := 8000000 + random()*2000000;   -- 8-10 mln main revenue / day
    v_day_rev := 0;
    guard     := 0;

    WHILE v_day_rev < v_target AND guard < 250 LOOP
      guard := guard + 1;
      v_counter := v_counter + 1;

      v_doctor := NULL; v_mkt := NULL; v_vip := false; v_nurse := false;
      v_refby := ''; v_channel := ''; v_fname := ''; v_lname := '';
      v_addr := ''; v_worker := NULL; v_offline := false; v_user := NULL;
      v_cr := ''; v_cb := ''; v_crl := '';

      v_pid := pids[(floor(random()*array_length(pids,1))+1)::INT];

      v_rand := random();
      IF    v_rand < 0.28 THEN v_otype := 'online';
      ELSIF v_rand < 0.55 THEN v_otype := 'offline';
      ELSIF v_rand < 0.68 THEN v_otype := 'doctor';
      ELSIF v_rand < 0.76 THEN v_otype := 'vip';
      ELSE                     v_otype := 'marketolog';
      END IF;

      -- 80% packs, 20% loose pieces
      IF random() < 0.80 THEN
        v_unit := 'pack';
        v_rand := random();
        IF    v_rand < 0.65 THEN v_qty := 1;
        ELSIF v_rand < 0.90 THEN v_qty := 2;
        ELSE                     v_qty := 3; END IF;
        v_price := v_qty * PRICE_PILL * QPP;
      ELSE
        v_unit := 'piece';
        v_rand := random();
        IF    v_rand < 0.40 THEN v_qty := 1;
        ELSIF v_rand < 0.70 THEN v_qty := 2;
        ELSIF v_rand < 0.85 THEN v_qty := 3;
        ELSIF v_rand < 0.93 THEN v_qty := 4;
        ELSE                     v_qty := 5; END IF;
        v_price := v_qty * PRICE_PILL;
      END IF;

      CASE v_otype
        WHEN 'online' THEN
          v_offline := false;
          v_user := uids[(floor(random()*array_length(uids,1))+1)::INT];
          v_addr := addrs[(floor(random()*array_length(addrs,1))+1)::INT];
          IF random() < 0.55 THEN v_pay := 'online'; ELSE v_pay := 'card'; END IF;
          v_rand := random();
          IF    v_rand < 0.12 THEN v_status := 'pending';
          ELSIF v_rand < 0.19 THEN v_status := 'cancelled';
                v_cr := 'Mijoz buyurtmani bekor qildi'; v_cb := 'Mijoz'; v_crl := 'user';
          ELSIF v_rand < 0.40 THEN v_status := 'confirmed';
          ELSIF v_rand < 0.60 THEN v_status := 'shipped';
          ELSIF v_rand < 0.78 THEN v_status := 'in_transit';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'offline' THEN
          v_offline := true; v_worker := W_PICKUP;
          IF random() < 0.25 THEN v_user := NULL;
          ELSE v_user := uids[(floor(random()*array_length(uids,1))+1)::INT]; END IF;
          v_rand := random();
          IF    v_rand < 0.55 THEN v_pay := 'cash';
          ELSIF v_rand < 0.85 THEN v_pay := 'terminal';
          ELSE                     v_pay := 'card'; END IF;
          v_rand := random();
          IF    v_rand < 0.07 THEN v_status := 'cancelled';
                v_cr := 'Mijoz olib ketmadi'; v_cb := 'Kassir'; v_crl := 'worker';
          ELSIF v_rand < 0.45 THEN v_status := 'confirmed';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'doctor' THEN
          v_offline := true; v_worker := W_PICKUP; v_doctor := DOC_ID;
          v_refby := 'Karimov'; v_nurse := (random() < 0.50);
          v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
          v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
          IF random() < 0.60 THEN v_pay := 'cash'; ELSE v_pay := 'terminal'; END IF;
          v_rand := random();
          IF    v_rand < 0.07 THEN v_status := 'cancelled';
                v_cr := 'Bemor kelmadi'; v_cb := 'Hamshira'; v_crl := 'nurse';
          ELSIF v_rand < 0.42 THEN v_status := 'confirmed';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'vip' THEN
          v_offline := true; v_worker := W_PICKUP; v_doctor := DOC_ID;
          v_vip := true; v_refby := 'Karimov'; v_nurse := (random() < 0.70);
          v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
          v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
          v_pay := 'cash'; v_price := 0;            -- own patient: free
          v_rand := random();
          IF    v_rand < 0.03 THEN v_status := 'cancelled';
                v_cr := 'Bekor qilindi'; v_cb := 'Hamshira'; v_crl := 'nurse';
          ELSIF v_rand < 0.25 THEN v_status := 'confirmed';
          ELSE                     v_status := 'delivered'; END IF;

        WHEN 'marketolog' THEN
          v_offline := true; v_worker := NULL;
          v_mkt := mkt_ids[(floor(random()*array_length(mkt_ids,1))+1)::INT];
          v_pay := 'cash';
          v_channel := chans[(floor(random()*array_length(chans,1))+1)::INT];
          IF random() < 0.30 THEN v_user := NULL;
          ELSE v_user := uids[(floor(random()*array_length(uids,1))+1)::INT]; END IF;
          v_rand := random();
          IF    v_rand < 0.10 THEN v_status := 'cancelled';
                v_cr := 'Mijoz rad etdi'; v_cb := 'Marketolog'; v_crl := 'manager';
          ELSIF v_rand < 0.32 THEN v_status := 'confirmed';
          ELSIF v_rand < 0.50 THEN v_status := 'shipped';
          ELSIF v_rand < 0.65 THEN v_status := 'in_transit';
          ELSE                     v_status := 'delivered'; END IF;
      END CASE;

      IF v_otype IN ('online','offline') AND random() < 0.25 THEN
        v_fname := fn[(floor(random()*array_length(fn,1))+1)::INT];
        v_lname := ln[(floor(random()*array_length(ln,1))+1)::INT];
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
        marketolog_id, archived, archive_reason
      ) VALUES (
        v_user, v_status, v_phone, v_ts, v_code,
        v_addr, v_offline, '', 0, 0, v_refby,
        v_worker, v_nurse, v_fname, v_lname,
        v_doctor, v_vip, v_pay, '',
        v_cr, v_cb, v_crl,
        v_channel, false, '',
        v_mkt, false, ''
      ) RETURNING id INTO v_order_id;

      INSERT INTO order_items
        (order_id, product_id, quantity, unit_type, price, original_quantity)
      VALUES
        (v_order_id, v_pid, v_qty, v_unit, v_price, v_qty);

      -- advance day target only for what shows on the main revenue chart:
      -- non-cancelled, non-marketolog (vip price 0 adds nothing)
      IF v_status <> 'cancelled' AND v_mkt IS NULL THEN
        v_day_rev := v_day_rev + v_price;
      END IF;
    END LOOP;
  END LOOP;

  RAISE NOTICE 'Created % orders over % days', v_counter, SPAN + 1;
END $SEED$;

-- 5. CLOSED-SYSTEM stock reconciliation.
--    stock = 60 000 − штук продано, with the SAME filter the product analytics
--    uses (status <> 'cancelled' AND archived = false). No restocking, so
--    «На складе» + «Продано» always add back up to 60 000 штук (1000 капс).
UPDATE products SET stock_quantity = 60000;
UPDATE products p SET stock_quantity = 60000 - c.pieces
FROM (
  SELECT oi.product_id AS pid,
         SUM(CASE WHEN oi.unit_type = 'pack' THEN oi.quantity * 60 ELSE oi.quantity END) AS pieces
  FROM order_items oi
  JOIN orders o ON o.id = oi.order_id
  WHERE o.status <> 'cancelled' AND o.archived = false
  GROUP BY oi.product_id
) c WHERE c.pid = p.id;

COMMIT;

-- ── verification ────────────────────────────────────────────────────────────
\echo ''
\echo '=== orders / items ==='
SELECT (SELECT COUNT(*) FROM orders) AS orders, (SELECT COUNT(*) FROM order_items) AS items;
\echo '=== status ==='
SELECT status, COUNT(*) FROM orders GROUP BY status ORDER BY 2 DESC;
\echo '=== channel ==='
SELECT CASE WHEN marketolog_id IS NOT NULL THEN 'marketolog'
            WHEN is_v_ip THEN 'vip'
            WHEN doctor_id IS NOT NULL THEN 'doctor'
            WHEN is_offline THEN 'offline' ELSE 'online' END AS channel, COUNT(*)
FROM orders GROUP BY 1 ORDER BY 2 DESC;
\echo '=== main revenue / day (non-cancelled, non-marketolog) ==='
SELECT round(avg(rev)) avg_day, min(rev) min_day, max(rev) max_day, count(*) days
FROM (SELECT date_trunc('day',o.created_at) d, SUM(oi.price) rev
      FROM orders o JOIN order_items oi ON oi.order_id=o.id
      WHERE o.status<>'cancelled' AND o.marketolog_id IS NULL GROUP BY 1) t;
\echo '=== stock reconcile: sklad_kaps + prodano_kaps must equal ~1000 ==='
WITH consumed AS (
  SELECT oi.product_id pid, SUM(CASE WHEN oi.unit_type='pack' THEN oi.quantity ELSE 0 END) sold_packs
  FROM order_items oi JOIN orders o ON o.id=oi.order_id
  WHERE o.status<>'cancelled' AND o.archived=false GROUP BY 1)
SELECT p.id, p.name, p.stock_quantity sht, floor(p.stock_quantity/60.0) sklad_kaps,
       COALESCE(c.sold_packs,0) prodano_kaps,
       floor(p.stock_quantity/60.0)+COALESCE(c.sold_packs,0) summa
FROM products p LEFT JOIN consumed c ON c.pid=p.id ORDER BY p.id;
