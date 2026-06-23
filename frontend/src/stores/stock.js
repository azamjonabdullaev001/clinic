import { reactive } from 'vue'

// Live stock quantities keyed by product id, kept in sync via WebSocket.
// Components read liveStock[productId] and fall back to their loaded value.
export const liveStock = reactive({})

// Real-time signals: `ordersVersion` bumps whenever any order changes anywhere and
// `productsVersion` whenever the product catalog changes — panels watch these and refresh
// their lists/analytics without a page reload.
export const realtime = reactive({ ordersVersion: 0, productsVersion: 0 })

let ws = null
let started = false

// Coalesce bursts of order/product events. A busy till can fire many "orders" messages
// per second; without throttling every panel would reload its list + analytics on each
// one. This fires the first bump immediately (instant feedback for a single sale) then
// collapses any further bumps in the same ~1s window into one trailing bump — so the data
// is always fresh within a second, but a burst costs one refresh instead of dozens.
function makeBumper(key, interval = 1000) {
  let last = 0
  let pending = false
  return () => {
    const now = Date.now()
    const since = now - last
    if (since >= interval) {
      last = now
      realtime[key]++
    } else if (!pending) {
      pending = true
      setTimeout(() => {
        pending = false
        last = Date.now()
        realtime[key]++
      }, interval - since)
    }
  }
}
const bumpOrders = makeBumper('ordersVersion')
const bumpProducts = makeBumper('productsVersion')

function connect() {
  const proto = location.protocol === 'https:' ? 'wss' : 'ws'
  try {
    ws = new WebSocket(`${proto}://${location.host}/api/ws/stock`)
  } catch {
    setTimeout(connect, 4000)
    return
  }
  ws.onmessage = (e) => {
    try {
      const m = JSON.parse(e.data)
      if (m && m.type === 'stock') liveStock[m.product_id] = m.quantity
      else if (m && m.type === 'orders') bumpOrders()
      else if (m && m.type === 'products') bumpProducts()
    } catch { /* ignore */ }
  }
  ws.onclose = () => { ws = null; setTimeout(connect, 4000) }
  ws.onerror = () => { try { ws && ws.close() } catch { /* ignore */ } }
}

// Connect once for the app session.
export function useStockSocket() {
  if (!started) {
    started = true
    connect()
  }
  return { liveStock }
}

// Quantity to display: live value if known, otherwise the loaded fallback.
export function displayStock(productId, fallback) {
  return liveStock[productId] !== undefined ? liveStock[productId] : (fallback || 0)
}
