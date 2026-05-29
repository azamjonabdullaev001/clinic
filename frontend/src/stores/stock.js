import { reactive } from 'vue'

// Live stock quantities keyed by product id, kept in sync via WebSocket.
// Components read liveStock[productId] and fall back to their loaded value.
export const liveStock = reactive({})

// Real-time signals: `ordersVersion` bumps whenever any order changes anywhere,
// so panels can watch it and refresh their lists/analytics without a page reload.
export const realtime = reactive({ ordersVersion: 0 })

let ws = null
let started = false

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
      else if (m && m.type === 'orders') realtime.ordersVersion++
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
