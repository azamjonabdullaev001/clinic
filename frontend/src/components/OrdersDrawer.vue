<template>
  <Teleport to="body">
    <div v-if="ordersStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="ordersStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white/95 backdrop-blur-2xl shadow-2xl flex flex-col animate-slide-right">

        <!-- Header -->
        <div class="flex items-center justify-between px-6 py-5 border-b border-stone-100">
          <div>
            <h2 class="text-xl font-serif text-stone-900">{{ t.orders_title }}</h2>
            <p class="text-xs text-stone-400 mt-0.5">{{ activeOrders.length }} {{ t.orders_count }}</p>
          </div>
          <button @click="ordersStore.toggle()" class="p-2 hover:bg-stone-100 rounded-xl hover:rotate-90 transition-all duration-300">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="flex-1 flex items-center justify-center">
          <div class="w-8 h-8 border-2 border-brand-600 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <!-- Empty state -->
        <div v-else-if="orders.length === 0" class="flex-1 flex flex-col items-center justify-center px-8 text-center">
          <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
            <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <p class="text-stone-800 text-lg font-medium">{{ t.orders_empty }}</p>
        </div>

        <!-- Orders list -->
        <div v-else class="flex-1 overflow-y-auto px-6 py-5 space-y-4">

          <!-- Active orders (not cancelled) -->
          <div
            v-for="order in activeOrders"
            :key="order.id"
            class="rounded-2xl border border-stone-100 bg-stone-50 overflow-hidden"
          >
            <!-- Order header with code + status -->
            <div class="flex items-center justify-between px-4 py-3 bg-white border-b border-stone-100">
              <div>
                <p class="text-[10px] text-stone-400 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                <p class="text-2xl font-bold text-brand-700 tracking-widest">{{ order.order_code }}</p>
              </div>
              <span class="px-3 py-1.5 rounded-xl text-xs font-semibold" :class="statusClass(order.status)">
                {{ statusLabel(order.status) }}
              </span>
            </div>

            <!-- Items -->
            <div class="px-4 py-3 space-y-2.5">
              <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2">
                <div class="flex items-center gap-2 min-w-0">
                  <div class="w-8 h-8 rounded-lg bg-stone-200 overflow-hidden flex-shrink-0 flex items-center justify-center">
                    <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover" />
                    <svg v-else class="w-4 h-4 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                  </div>
                  <div class="min-w-0">
                    <p class="font-medium text-stone-900 text-sm truncate">{{ item.product?.name }}</p>
                    <p class="text-stone-400 text-xs">{{ item.quantity }} {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</p>
                  </div>
                </div>
                <p class="font-semibold text-stone-700 text-sm flex-shrink-0">{{ formatPrice(item.price) }} {{ t.currency }}</p>
              </div>
            </div>

            <!-- Payment / approval banner -->
            <div v-if="order.status === 'awaiting_payment'" class="mx-4 mb-3 bg-amber-50 border border-amber-200 rounded-xl px-3 py-2.5">
              <p class="text-sm font-semibold text-amber-800 mb-2">{{ t.status_awaiting_payment }}</p>
              <p class="text-xs text-amber-600 mb-3">{{ t.payment_scan_and_upload }}</p>
              <!-- QR Code (static file) -->
              <div class="flex justify-center mb-3">
                <img
                  v-if="!qrImgFailed"
                  src="/images/QRpayment.jpg"
                  class="w-48 h-48 object-contain rounded-xl border border-amber-200 bg-white p-1"
                  @error="qrImgFailed = true"
                />
                <div v-else class="w-48 h-48 flex flex-col items-center justify-center bg-white rounded-xl border border-amber-200 text-center px-2">
                  <svg class="w-8 h-8 text-amber-300 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z"/></svg>
                  <p class="text-xs text-amber-500">{{ t.payment_qr_not_configured }}</p>
                  <button @click="qrImgFailed = false" class="text-xs text-amber-600 font-medium mt-1 underline">{{ t.qr_retry }}</button>
                </div>
              </div>
              <input :ref="el => receiptInputs[order.id] = el" type="file" accept="image/*" class="hidden" @change="e => uploadReceipt(order.id, e)"/>
              <button @click="receiptInputs[order.id]?.click()" :disabled="receiptUploading === order.id" class="w-full bg-amber-600 text-white text-sm px-4 py-2.5 rounded-lg hover:bg-amber-700 transition disabled:opacity-50 font-medium">
                {{ receiptUploading === order.id ? t.uploading_text : t.upload_receipt_btn }}
              </button>
            </div>
            <div v-else-if="order.status === 'in_transit' || order.status === 'delivered'" class="mx-4 mb-3 bg-green-50 border border-green-200 rounded-xl px-3 py-2.5 flex items-center gap-2">
              <svg class="w-5 h-5 text-green-600 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/></svg>
              <p class="text-sm font-medium text-green-700">Ваш заказ одобрен! Спасибо, что выбрали нас 💚</p>
            </div>

            <!-- BTS cargo info -->
            <div v-if="order.bts_pickup_point || order.bts_tracking_number" class="mx-4 mb-3 bg-blue-50 border border-blue-200 rounded-xl px-3 py-2.5">
              <p class="text-xs font-semibold text-blue-700 mb-1.5 flex items-center gap-1">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                БТС Карго — ваш заказ в пути
              </p>
              <p v-if="order.bts_pickup_point" class="text-xs text-blue-800 font-medium">📍 {{ order.bts_pickup_point }}</p>
              <p v-if="order.bts_tracking_number" class="text-xs text-blue-600 mt-0.5">Трек: {{ order.bts_tracking_number }}</p>
              <button
                v-if="btsBranchFor(order)"
                @click="openBtsMap(order)"
                class="mt-2 w-full flex items-center justify-center gap-1.5 bg-blue-600 text-white text-xs font-semibold px-3 py-2 rounded-lg hover:bg-blue-700 transition"
              >
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                Показать маршрут доставки
              </button>
            </div>

            <!-- Total -->
            <div class="px-4 py-2.5 border-t border-stone-100 flex justify-between items-center">
              <span class="text-xs text-stone-400">{{ t.cart_total }}</span>
              <span class="font-bold text-brand-700">{{ formatPrice(orderTotal(order)) }} {{ t.currency }}</span>
            </div>

            <!-- Hide delivered/cancelled orders from history -->
            <div v-if="order.status === 'delivered'" class="px-4 pb-3 flex justify-end">
              <button @click="confirmHide(order)" class="text-xs text-stone-400 hover:text-red-500 transition flex items-center gap-1">
                <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                Удалить из истории
              </button>
            </div>
          </div>

          <!-- Cancelled orders section -->
          <template v-if="cancelledOrders.length > 0">
            <div class="pt-2">
              <div class="flex items-center gap-2 mb-3">
                <div class="flex-1 h-px bg-red-200"></div>
                <span class="text-xs font-semibold text-red-400 uppercase tracking-wider">{{ t.status_cancelled }}</span>
                <div class="flex-1 h-px bg-red-200"></div>
              </div>

              <div
                v-for="order in cancelledOrders"
                :key="order.id"
                class="rounded-2xl border border-red-200 bg-red-50/60 overflow-hidden mb-4"
              >
                <div class="flex items-center justify-between px-4 py-3 bg-red-50 border-b border-red-100">
                  <div>
                    <p class="text-[10px] text-red-300 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                    <p class="text-2xl font-bold text-red-400 tracking-widest">{{ order.order_code }}</p>
                  </div>
                  <div class="flex items-center gap-2">
                    <span class="px-3 py-1.5 rounded-xl text-xs font-semibold bg-red-100 text-red-600">
                      {{ t.status_cancelled }}
                    </span>
                    <button @click="confirmHide(order)" class="p-1.5 hover:bg-red-100 rounded-lg transition text-red-300 hover:text-red-500" title="Удалить из истории">
                      <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                    </button>
                  </div>
                </div>

                <div class="px-4 py-3 space-y-2">
                  <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2 opacity-60">
                    <div class="flex items-center gap-2 min-w-0">
                      <div class="w-8 h-8 rounded-lg bg-red-100 overflow-hidden flex-shrink-0 flex items-center justify-center">
                        <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover grayscale" />
                        <svg v-else class="w-4 h-4 text-red-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                        </svg>
                      </div>
                      <div class="min-w-0">
                        <p class="font-medium text-stone-700 text-sm truncate line-through">{{ item.product?.name }}</p>
                        <p class="text-stone-400 text-xs">{{ item.quantity }} {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</p>
                      </div>
                    </div>
                    <p class="font-semibold text-stone-500 text-sm flex-shrink-0 line-through">{{ formatPrice(item.price) }} {{ t.currency }}</p>
                  </div>
                </div>
              </div>
            </div>
          </template>

        </div>
      </div>
    </div>

    <!-- BTS Route Map Modal -->
    <div v-if="btsMapOpen" class="fixed inset-0 z-[300] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="closeBtsMap"></div>
      <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg mx-4 overflow-hidden flex flex-col" style="max-height:90vh">
        <div class="flex items-center justify-between px-5 py-4 border-b border-gray-100 bg-white flex-shrink-0">
          <div>
            <h3 class="text-base font-bold text-gray-900">Маршрут вашего заказа</h3>
            <p v-if="btsMapOrder" class="text-xs text-gray-500 mt-0.5">{{ btsMapOrder.order_code }} → {{ btsBranchFor(btsMapOrder)?.city }}</p>
          </div>
          <button @click="closeBtsMap" class="p-2 hover:bg-gray-100 rounded-xl transition">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
          </button>
        </div>
        <div v-if="btsMapLoading" class="flex items-center justify-center gap-3 text-gray-500 text-sm" style="height:380px">
          <svg class="w-5 h-5 animate-spin text-blue-500" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"/></svg>
          Строим маршрут…
        </div>
        <div v-else id="bts-route-map" style="height:380px;width:100%;flex-shrink:0"></div>
        <div v-if="btsMapOrder && btsBranchFor(btsMapOrder)" class="px-5 py-3 bg-gray-50 border-t border-gray-100 flex gap-4 text-xs flex-shrink-0">
          <div class="flex items-center gap-1.5">
            <span class="w-3 h-3 rounded-full bg-teal-600 flex-shrink-0"></span>
            <span class="text-gray-600 font-medium">Наша аптека (Андижан)</span>
          </div>
          <div class="flex items-center gap-1.5">
            <span class="w-3 h-3 rounded-full bg-blue-600 flex-shrink-0"></span>
            <span class="text-gray-600 font-medium">{{ btsBranchFor(btsMapOrder)?.name }}</span>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch, nextTick } from 'vue'
import { useOrdersStore } from '../stores/orders'
import { useLangStore } from '../stores/lang'
import { api } from '../stores/auth'
import { detectBtsBranch } from '../config/btsBranches'

// Clinic/pickup point coordinates (Andijan)
const CLINIC_LAT = 40.7760385
const CLINIC_LNG = 72.3784014

// BTS branches loaded from API
const btsBranches = ref([])
async function loadBtsBranches() {
  try {
    const res = await api.get('/bts-branches')
    btsBranches.value = res.data || []
  } catch { /* non-critical */ }
}

const ordersStore = useOrdersStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)

const orders = ref([])
const loading = ref(false)
const qrImgFailed = ref(false)
const receiptInputs = ref({})
const receiptUploading = ref(null)

const activeOrders = computed(() => orders.value.filter(o => o.status !== 'cancelled'))
const cancelledOrders = computed(() => orders.value.filter(o => o.status === 'cancelled'))

async function loadOrders() {
  loading.value = true
  try {
    const res = await api.get('/orders')
    orders.value = res.data || []
  } catch (e) {
    orders.value = []
  } finally {
    loading.value = false
  }
}

watch(() => ordersStore.isOpen, (val) => {
  if (val) {
    loadOrders()
    loadBtsBranches()
    qrImgFailed.value = false
  }
})

// ===== BTS Route Map =====
const btsMapOpen = ref(false)
const btsMapOrder = ref(null)
const btsMapLoading = ref(false)
let btsMapInstance = null

function btsBranchFor(order) {
  // Prefer exact branch saved on order (has verified coordinates)
  if (order.bts_branch_lat && order.bts_branch_lng) {
    return {
      name: order.bts_pickup_point || 'БТС пункт выдачи',
      address: '',
      lat: order.bts_branch_lat,
      lng: order.bts_branch_lng,
      city: '',
    }
  }
  // Fall back to API-configured branch matched by delivery address
  return detectBtsBranch(order.delivery_address, btsBranches.value)
}

async function openBtsMap(order) {
  btsMapOrder.value = order
  btsMapOpen.value = true
  btsMapLoading.value = true
  await nextTick()
  await initBtsMap(order)
}

function closeBtsMap() {
  btsMapOpen.value = false
  btsMapOrder.value = null
  btsMapLoading.value = false
  if (btsMapInstance) { btsMapInstance.remove(); btsMapInstance = null }
}

async function initBtsMap(order) {
  const branch = detectBtsBranch(order.delivery_address)
  if (!branch) { btsMapLoading.value = false; return }

  const L = await import('leaflet')
  await import('leaflet/dist/leaflet.css')

  delete L.Icon.Default.prototype._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })

  btsMapLoading.value = false
  await nextTick()

  const mapEl = document.getElementById('bts-route-map')
  if (!mapEl) return

  const midLat = (CLINIC_LAT + branch.lat) / 2
  const midLng = (CLINIC_LNG + branch.lng) / 2

  btsMapInstance = L.map('bts-route-map').setView([midLat, midLng], 7)
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
  }).addTo(btsMapInstance)

  // Clinic marker (teal)
  const clinicIcon = L.divIcon({
    html: '<div style="background:#0d9488;width:14px;height:14px;border-radius:50%;border:3px solid white;box-shadow:0 0 0 2px #0d9488"></div>',
    className: '', iconSize: [14, 14], iconAnchor: [7, 7],
  })
  L.marker([CLINIC_LAT, CLINIC_LNG], { icon: clinicIcon })
    .addTo(btsMapInstance)
    .bindPopup('<b>Наша аптека</b><br>Андижан')

  // BTS branch marker (blue)
  const btsIcon = L.divIcon({
    html: '<div style="background:#2563eb;width:16px;height:16px;border-radius:50%;border:3px solid white;box-shadow:0 0 0 2px #2563eb"></div>',
    className: '', iconSize: [16, 16], iconAnchor: [8, 8],
  })
  L.marker([branch.lat, branch.lng], { icon: btsIcon })
    .addTo(btsMapInstance)
    .bindPopup(`<b>${branch.name}</b><br>${branch.address}`)
    .openPopup()

  // OSRM road route
  try {
    const osrmUrl = `https://router.project-osrm.org/route/v1/driving/${CLINIC_LNG},${CLINIC_LAT};${branch.lng},${branch.lat}?overview=full&geometries=geojson`
    const res = await fetch(osrmUrl, { signal: AbortSignal.timeout(8000) })
    const data = await res.json()
    if (data.code === 'Ok' && data.routes?.[0]) {
      const coords = data.routes[0].geometry.coordinates.map(([lng, lat]) => [lat, lng])
      L.polyline(coords, { color: '#2563eb', weight: 4, opacity: 0.85 }).addTo(btsMapInstance)
    } else throw new Error('no route')
  } catch {
    L.polyline([[CLINIC_LAT, CLINIC_LNG], [branch.lat, branch.lng]], { color: '#2563eb', weight: 3, dashArray: '8 5', opacity: 0.8 }).addTo(btsMapInstance)
  }

  btsMapInstance.fitBounds([[CLINIC_LAT, CLINIC_LNG], [branch.lat, branch.lng]], { padding: [40, 40] })
}

async function confirmHide(order) {
  if (!confirm(`Удалить заказ #${order.order_code} из истории?`)) return
  try {
    await api.delete(`/orders/${order.id}/hide`)
    orders.value = orders.value.filter(o => o.id !== order.id)
  } catch (err) {
    alert(err.response?.data?.error || 'Ошибка при удалении')
  }
}

async function uploadReceipt(orderId, e) {
  const f = e.target.files[0]
  if (!f) return
  receiptUploading.value = orderId
  try {
    const fd = new FormData()
    fd.append('receipt', f)
    await api.post(`/orders/${orderId}/receipt`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    await loadOrders()
  } catch (err) {
    alert(err.response?.data?.error || 'Ошибка при отправке чека')
  } finally {
    receiptUploading.value = null
    e.target.value = ''
  }
}

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

function orderTotal(order) {
  return (order.items || []).reduce((sum, i) => sum + i.price, 0)
}

function statusLabel(status) {
  const map = {
    awaiting_payment: 'Ожидает оплаты',
    pending: t.value.status_pending,
    in_transit: t.value.status_in_transit,
    delivered: t.value.status_delivered,
    cancelled: t.value.status_cancelled,
  }
  return map[status] || status
}

function statusClass(status) {
  const map = {
    awaiting_payment: 'bg-amber-100 text-amber-700',
    pending: 'bg-yellow-100 text-yellow-700',
    in_transit: 'bg-orange-100 text-orange-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-600',
  }
  return map[status] || 'bg-stone-100 text-stone-600'
}
</script>
