# Demo-data rebuild

`seed_full.sql` recreates the whole demo dataset from scratch, idempotently.

## What it does
1. Ensures 3 marketolog (`role=manager`) workers exist — password **`Market2026`**
   (phones `998900000001/2/3`). Needed so orders show up in the marketolog panel.
2. Deletes **all** orders, order items and worker stocks (analytics is derived from
   orders, so it resets too).
3. Resets every product to **60 000 штук = 1000 «капс»** (1 «капс»/упаковка = 60 капсул).
4. Generates ~6 months (2025-12-01 … 2026-05-31) of realistic orders:
   - **~8–10 mln сум/day** main revenue
   - channels: online / offline / doctor / vip (own patient, free) / marketolog
   - `pending` only on online; `cancelled` & `confirmed` in both online and offline
   - vip & marketolog are offline-only
5. Recomputes warehouse stock as a **closed system**: `stock = 60 000 − штук продано`
   (same filter the product analytics uses), so **«На складе» + «Продано» = 1000 капс**
   for every product — analytics and stock match one-to-one.

## Run it (on the server)
```bash
bash scripts/run_seed_full.sh
```
or manually:
```bash
docker cp scripts/seed_full.sql clinic-db-1:/seed_full.sql
docker exec -i clinic-db-1 psql -U clinic -d clinic -v ON_ERROR_STOP=1 -f /seed_full.sql
```

The script prints verification tables at the end (status mix, channel mix,
revenue/day, and the stock reconciliation check).

> The `*.py` helpers and `*backup*.sql` in this folder are local-only (they contain
> server credentials / large dumps) and are git-ignored — never commit them.
