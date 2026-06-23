# Performance & Scalability Audit

Goal: keep the system fast (< 500 ms for most operations) at 100k → 1M+ orders.

## Root cause of the "dies at ~10k orders" symptom

A self-reinforcing loop:

1. Every order mutation calls `BroadcastOrders()` over WebSocket.
2. Every connected panel reacts by re-fetching its order list.
3. Several list endpoints ran `Find(&orders)` with **no LIMIT**, loading the entire,
   ever-growing `orders` table (plus `Items.Product`, `User`) into Go memory on every fetch.

So the per-click cost grew linearly with total order count. At ~10k rows × several
concurrent panels × every mutation, the backend spent all its time serializing the whole
history. Postgres also had to sequentially scan because the hot predicates weren't all indexed.

## What was changed (implemented in this branch)

| # | Problem | Critical | Fix | Expected effect |
|---|---------|----------|-----|-----------------|
| 1 | `GetPickupOrders` loads ALL of a worker's orders every refresh | 🔴 High | `LIMIT/OFFSET` via shared `paginate()` (default 500, cap 1000) + partial index | O(table) → O(page). Panel load constant regardless of history size |
| 2 | `GetManagerOrders`, `GetNurseOrders`, `GetDoctorOrders`, `GetUserOrders` unbounded | 🔴 High | Same `paginate()` bounding | Constant memory/latency per request |
| 3 | `GetMarketologDebt` loads every order+item into Go to sum an all-time total | 🔴 High | Rewritten as a single SQL `GROUP BY` aggregation (constant memory) | No OOM at 1M rows; debt computed in DB |
| 4 | Hot predicates not indexed (doctor_id, marketolog+status, active-orders list) | 🟠 Med | Added composite + **partial** index `WHERE archived=false AND is_deleted=false` | Index-only/range scans instead of seq scans |
| 5 | No rate limiting → a single client can flood Postgres | 🟠 Med | In-process per-IP token bucket (`middleware.RateLimit`, 30 rps / 60 burst) | Protects DB from accidental/abusive floods |
| 6 | No graceful shutdown / server timeouts → dropped requests, leaked conns on deploy | 🟠 Med | `http.Server` with timeouts + SIGTERM `Shutdown` + DB pool close | Clean deploys, slow-loris protection |
| 7 | Postgres ran with defaults (4MB work_mem, 128MB shared_buffers) | 🟠 Med | Tuned `shared_buffers/effective_cache_size/work_mem/maintenance_work_mem` in compose | Hot indexes + sorts stay in RAM |
| 8 | No visibility into slow queries | 🟢 Low | `log_min_duration_statement=500` | Every >500ms query is logged for follow-up |

Connection pooling was already bounded (`SetMaxOpenConns(25)`), and core indexes
(`order_items.order_id`, `orders.created_at/status/...`) already existed.

## Example: before / after (item #3, marketolog debt)

Before — loads the marketolog's entire order history into memory, then loops:

```go
var orders []models.Order
database.DB.Where("marketolog_id = ? AND status != ? AND is_deleted = ?", id, "cancelled", false).
    Preload("Items.Product").Find(&orders) // unbounded
for _, o := range orders { for _, it := range o.Items { /* sum in Go */ } }
```

After — a single constant-memory aggregation in Postgres:

```sql
SELECT p.name,
       COALESCE(SUM(oi.price),0) AS revenue,
       COALESCE(SUM(CASE WHEN oi.unit_type='piece' THEN oi.quantity ELSE 0 END),0) AS pieces,
       COALESCE(SUM(CASE WHEN oi.unit_type<>'piece' THEN oi.quantity ELSE 0 END),0) AS capsules
FROM order_items oi
JOIN orders o ON o.id = oi.order_id
LEFT JOIN products p ON p.id = oi.product_id
WHERE o.marketolog_id = ? AND o.status <> 'cancelled' AND o.is_deleted = false AND oi.quantity > 0
GROUP BY p.name;
```

## Lossless pagination (implemented, no behaviour lost)

To make the bounding above safe — the pickup panel computes history and payment-breakdown
stats client-side over whatever it loaded — the frontend now loads the most recent 500
first (instant) and shows a **"Показать ещё"** button that fetches the next page and
**appends** it. Every existing client-side computation runs unchanged over the loaded set;
full history is always reachable, just lazily. Per-worker lists (manager/nurse/doctor) keep
a generous 2000-row bound because ManagerPanel sums an all-time stat over its list.

## Refetch-storm fixes (implemented, no numbers changed)

At scale the felt lag came from the real-time refresh loop, not single queries (analytics
is already date-range bounded on the `created_at` index, so daily/weekly stay fast even at
500k rows). Two safe changes remove it:

- **Debounced realtime signal** (`frontend/src/stores/stock.js`). A busy till fires many
  `orders` WebSocket events per second; each one made every panel reload its list +
  analytics. The version bump is now throttled: the first event fires immediately (instant
  feedback for a single sale), further events in the same ~1s window collapse into one
  trailing bump. A burst of 30 sales costs one refresh instead of 30. Live stock numbers
  are untouched and still update instantly.
- **Identity-scoped response cache** (`backend/handlers/cache.go`, `CacheGET`). Successful
  analytics/stats GETs are cached for 15s, keyed per caller, and invalidated the instant
  orders/products change (a `dataVersion` counter bumped by `BroadcastOrders`/`Products`).
  Multiple admins, tab switches and re-renders are served from cache instead of re-scanning
  the orders table. Because any real change bumps the version and forces recompute, **no
  cached value can ever differ from a fresh one** — the computation code is unchanged.

## Recommended next (NOT yet implemented — need product/infra decision)

These need your sign-off because they add infrastructure or are a larger change:

- **Server-side analytics aggregation.** `GetAnalytics`/`GetWorkerAnalytics` still pull
  in-range orders into Go. For multi-year ranges at 1M+ rows, move the SUM/GROUP BY into
  SQL (same pattern as #3) and cache the result for ~30–60s.
- **Redis** (or in-process TTL cache) for read-heavy reference data (public products,
  settings, BTS branches). Single-instance: an in-process cache is enough; multi-instance: Redis.
- **Table partitioning** of `orders` by `created_at` (monthly `RANGE`) once you cross a
  few million rows, so old partitions can be detached/archived and scans hit one partition.
- **Monitoring stack**: Prometheus + Grafana + postgres_exporter + nginx_exporter. The
  `log_min_duration_statement=500` setting already surfaces slow queries in the DB log.
