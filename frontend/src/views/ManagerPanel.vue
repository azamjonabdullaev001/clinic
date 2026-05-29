<template>
  <div class="min-h-screen flex" :class="{ 'night-mode': night }">

    <!-- ===== SIDEBAR ===== -->
    <aside class="worker-sidebar flex-shrink-0 flex flex-col"
      style="width:200px;min-height:100vh;position:sticky;top:0;height:100vh;overflow-y:auto;background:#111827;border-right:1px solid rgba(255,255,255,0.06);">

      <!-- Brand -->
      <div class="px-4 py-5 flex items-center gap-3"
        style="border-bottom:1px solid rgba(255,255,255,0.06);">
        <div class="w-9 h-9 rounded-xl bg-purple-600 flex items-center justify-center flex-shrink-0 shadow-lg">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-white text-sm font-bold leading-tight">{{ txt.title }}</p>
          <p class="text-gray-400 text-xs leading-tight truncate">{{ authStore.worker?.name }}</p>
        </div>
      </div>

      <!-- Nav items -->
      <nav class="flex-1 px-3 py-3 space-y-0.5">
        <button
          v-for="s in [
            {k:'sale',      l:txt.nav_sale,      d:'M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z'},
            {k:'stock',     l:txt.nav_stock,     d:'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4'},
            {k:'analytics', l:txt.nav_analytics, d:'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'}
          ]"
          :key="s.k"
          @click="tab = s.k"
          :class="tab === s.k
            ? 'bg-purple-600 text-white shadow-lg shadow-purple-900/30'
            : 'text-gray-400 hover:bg-white/5 hover:text-gray-200'"
          class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-all text-left">
          <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" :d="s.d"/>
          </svg>
          {{ s.l }}
        </button>
      </nav>

      <!-- Bottom: lang + night + logout -->
      <div class="px-4 pb-5 pt-3 space-y-3"
        style="border-top:1px solid rgba(255,255,255,0.06);">
        <div class="flex items-center gap-1.5">
          <button @click="lang='ru'"
            :class="lang==='ru' ? 'text-white font-bold' : 'text-gray-500 hover:text-gray-300'"
            class="text-sm transition">RU</button>
          <span class="text-gray-700 text-sm">|</span>
          <button @click="lang='uz'"
            :class="lang==='uz' ? 'text-white font-bold' : 'text-gray-500 hover:text-gray-300'"
            class="text-sm transition">UZ</button>
        </div>
        <button @click="toggleNight"
          class="flex items-center gap-2 text-sm transition w-full"
          :class="night ? 'text-amber-400' : 'text-gray-500 hover:text-gray-300'">
          <svg v-if="!night" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>
          </svg>
          <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
          {{ night ? txt.day_mode : txt.night_mode }}
        </button>
        <button @click="logout"
          class="flex items-center gap-2 text-red-400 hover:text-red-300 text-sm font-medium transition w-full">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
          </svg>
          {{ txt.logout }}
        </button>
      </div>
    </aside>

    <!-- ===== MAIN CONTENT ===== -->
    <div class="flex-1 overflow-auto bg-gray-100">
      <div class="p-6 space-y-6 max-w-5xl mx-auto">

        <!-- ===== SALE ===== -->
        <div v-show="tab === 'sale'" class="bg-white rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b flex items-center gap-3">
            <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17"/></svg>
            </div>
            <h2 class="font-bold text-gray-800">Продажа на маркетплейсе</h2>
          </div>
          <div class="px-6 py-5 space-y-4">
            <!-- Unit toggle -->
            <div class="flex gap-2">
              <button @click="setUnit('pack')" :class="saleUnit === 'pack' ? 'bg-purple-600 text-white border-purple-600' : 'bg-white text-gray-600 border-gray-300'" class="px-3 py-1.5 rounded-lg text-sm font-medium border transition">Капсула</button>
              <button @click="setUnit('piece')" :class="saleUnit === 'piece' ? 'bg-purple-600 text-white border-purple-600' : 'bg-white text-gray-600 border-gray-300'" class="px-3 py-1.5 rounded-lg text-sm font-medium border transition">Штука</button>
            </div>
            <!-- Product grid (tap to add) -->
            <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
              <button v-for="p in products" :key="p.id" @click="addToCart(p)" :disabled="availOf(p.id) <= cartQty(p.id)"
                class="text-left bg-gray-50 rounded-xl p-3 border-2 transition relative disabled:opacity-50"
                :class="cartQty(p.id) > 0 ? 'border-purple-400 bg-purple-50' : 'border-transparent hover:border-purple-200'">
                <div v-if="cartQty(p.id) > 0" class="absolute -top-2 -right-2 w-6 h-6 bg-purple-600 text-white text-xs font-bold rounded-full flex items-center justify-center">{{ cartQty(p.id) }}</div>
                <p class="text-xs font-semibold text-gray-800 leading-tight mb-1 line-clamp-2">{{ p.name }}</p>
                <p class="text-xs font-bold text-purple-600">{{ formatPrice(saleUnit === 'piece' ? p.price_per_pill : p.price_per_pack) }} сўм</p>
                <p class="text-[11px]" :class="availOf(p.id) > 0 ? 'text-gray-400' : 'text-red-500'">склад: {{ saleUnit === 'piece' ? (stockOf(p.id) + ' шт') : (capsulesOf(p.id) + ' капс') }}</p>
              </button>
            </div>

            <!-- Cart -->
            <div v-if="items.length" class="border rounded-xl overflow-hidden">
              <div v-for="item in items" :key="item.product_id" class="flex items-center justify-between px-4 py-2.5 border-b last:border-0 bg-gray-50">
                <span class="font-medium text-gray-800 text-sm flex-1 truncate">{{ item.name }} <span class="text-gray-400">({{ item.unit_type === 'piece' ? 'шт' : 'капс' }})</span></span>
                <div class="flex items-center gap-1.5 mx-2">
                  <button @click="decQty(item)" class="w-7 h-7 bg-gray-200 rounded-lg font-bold text-sm">−</button>
                  <span class="w-7 text-center text-sm font-bold">{{ item.quantity }}</span>
                  <button @click="addToCart(productById(item.product_id))" :disabled="availOf(item.product_id) <= item.quantity" class="w-7 h-7 bg-purple-100 text-purple-700 rounded-lg font-bold text-sm disabled:opacity-40">+</button>
                </div>
                <span class="font-semibold text-gray-700 text-sm w-24 text-right">{{ formatPrice(item.price) }} сўм</span>
              </div>
              <div class="flex items-center justify-between px-4 py-2 bg-white">
                <span class="text-sm font-semibold text-gray-700">Итого:</span>
                <span class="font-bold text-purple-600">{{ formatPrice(total) }} сўм</span>
              </div>
            </div>

            <div v-if="items.length">
              <label class="block text-xs font-medium text-gray-500 mb-1.5">Площадка</label>
              <div class="flex gap-2 flex-wrap">
                <button v-for="m in marketplaces" :key="m.value" @click="channel = m.value"
                  :class="channel === m.value ? 'bg-purple-600 text-white border-purple-600' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ m.label }}</button>
              </div>
            </div>

            <input v-if="items.length" v-model="note" placeholder="Покупатель / заметка (необязательно)"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500" />

            <button @click="submitSale" :disabled="!canSubmit || submitting"
              class="w-full bg-purple-600 text-white py-2.5 rounded-lg hover:bg-purple-700 transition font-medium text-sm disabled:opacity-40">
              {{ submitting ? 'Запись...' : 'Подтвердить продажу' }}
            </button>
            <div v-if="success" class="bg-purple-50 border border-purple-200 rounded-lg px-4 py-3 text-sm text-purple-700">Продажа записана ✓ Списано со склада.</div>
            <div v-if="saleError" class="bg-red-50 border border-red-200 rounded-lg px-4 py-3 text-sm text-red-600">{{ saleError }}</div>
          </div>
        </div>

        <!-- ===== STOCK + ORDERS ===== -->
        <div v-show="tab === 'stock'" class="space-y-6">
          <!-- Add incoming stock -->
          <div class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b"><h2 class="font-bold text-gray-800">Склад — приход товара</h2></div>
            <div class="px-6 py-5 space-y-3">
              <div class="flex gap-2 flex-wrap items-end">
                <div class="flex-1 min-w-[180px]">
                  <label class="block text-xs font-medium text-gray-500 mb-1">Препарат</label>
                  <select v-model="stockProductId" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500">
                    <option value="">Выберите препарат</option>
                    <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
                  </select>
                </div>
                <div class="w-28">
                  <label class="block text-xs font-medium text-gray-500 mb-1">Кол-во (капс.)</label>
                  <input v-model.number="stockQty" type="number" min="1" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500" />
                </div>
                <button @click="addStock" :disabled="!stockProductId || stockQty < 1 || addingStock"
                  class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition text-sm font-medium disabled:opacity-40">+ Пополнить</button>
              </div>
            </div>
          </div>

          <!-- Stock list -->
          <div class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-6 py-4 border-b"><h2 class="font-bold text-gray-800">Склад</h2></div>
            <div class="divide-y divide-gray-100">
              <div v-for="s in stock" :key="s.id" class="flex justify-between items-center px-6 py-3 text-sm">
                <span class="text-gray-700">{{ s.product?.name }}</span>
                <span class="px-2.5 py-1 rounded-lg font-bold" :class="stockOf(s.product_id) > 0 ? 'text-green-700 bg-green-50' : 'text-red-600 bg-red-50'">{{ capsulesOf(s.product_id) }} капс / {{ stockOf(s.product_id) }} шт</span>
              </div>
              <div v-if="stock.length === 0" class="px-6 py-8 text-center text-gray-400 text-sm">Склад пуст</div>
            </div>
          </div>

          <!-- My orders -->
          <div>
            <h2 class="text-lg font-bold text-gray-800 mb-3">Мои продажи</h2>
            <div class="space-y-3">
              <div v-for="order in orders" :key="order.id" class="bg-white rounded-xl shadow-sm p-5">
                <div class="flex justify-between items-start gap-3">
                  <div>
                    <span v-if="order.sales_channel" class="text-xs font-semibold px-2 py-0.5 rounded bg-purple-100 text-purple-700">{{ channelLabel(order.sales_channel) }}</span>
                    <p v-if="order.offline_note" class="text-sm text-gray-600 mt-1">{{ order.offline_note }}</p>
                    <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                  </div>
                  <p class="font-bold text-gray-800">{{ formatPrice(orderTotal(order)) }} сўм</p>
                </div>
                <div class="mt-3 pt-3 border-t border-gray-100 space-y-1">
                  <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm text-gray-600">
                    <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? 'шт' : 'капс' }}</span></span>
                    <span class="font-medium">{{ formatPrice(item.price) }} сўм</span>
                  </div>
                </div>
              </div>
              <div v-if="orders.length === 0" class="bg-white rounded-xl shadow-sm p-10 text-center text-gray-400">Продаж пока нет</div>
            </div>
          </div>
        </div>

        <!-- ===== ANALYTICS ===== -->
        <div v-show="tab === 'analytics'" class="bg-white rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b flex items-center justify-between flex-wrap gap-2">
            <h2 class="font-bold text-gray-800">Аналитика</h2>
            <div class="flex gap-1.5 flex-wrap items-center">
              <button v-for="p in periods" :key="p.v" @click="selectPeriod(p.v)"
                :class="period === p.v ? 'bg-indigo-600 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'"
                class="px-3 py-1.5 rounded-lg text-xs font-medium transition">{{ p.l }}</button>
              <input v-if="period === 'custom'" v-model="customDate" type="date" @change="loadAnalytics" class="border border-gray-300 rounded-lg px-2 py-1 text-xs" />
            </div>
          </div>
          <div class="px-6 py-5">
            <div v-if="analyticsLoading" class="text-gray-400 text-sm py-6 text-center">Загрузка...</div>
            <template v-else-if="analytics">
              <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 mb-4">
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Продаж</p><p class="text-2xl font-bold text-gray-800">{{ analytics.total_orders }}</p></div>
                <div class="bg-purple-50 rounded-xl p-4"><p class="text-xs text-purple-600 mb-1">Сумма (долг)</p><p class="text-lg font-bold text-purple-700">{{ formatPrice(analytics.total_revenue) }} сўм</p></div>
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Капсул</p><p class="text-2xl font-bold text-gray-800">{{ analytics.total_capsules }}</p></div>
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Штук</p><p class="text-2xl font-bold text-gray-800">{{ analytics.total_pieces }}</p></div>
              </div>
              <LineChart :points="analytics.points || []" color="#a855f7" />
              <div v-if="(analytics.products || []).length" class="mt-4 overflow-x-auto rounded-xl border border-purple-100">
                <table class="w-full text-sm">
                  <thead class="bg-purple-50">
                    <tr>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Препарат</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Капсул</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Штук</th>
                      <th class="text-right px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Сумма</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-purple-50">
                    <tr v-for="(p, i) in analytics.products" :key="i">
                      <td class="px-4 py-2 font-medium text-gray-800">{{ p.product_name }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ p.capsules }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ p.pieces }}</td>
                      <td class="px-4 py-2 text-right font-bold text-purple-700">{{ formatPrice(p.revenue) }} сўм</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'
import { useNight } from '../stores/night'
import { useStockSocket, displayStock, realtime } from '../stores/stock'
import LineChart from '../components/LineChart.vue'

const authStore = useAuthStore()
const router = useRouter()
const { night, toggle: toggleNight } = useNight()
useStockSocket()

const lang = ref(localStorage.getItem('managerLang') || 'ru')
watch(lang, v => localStorage.setItem('managerLang', v))

const mgTexts = {
  ru: {
    title: 'Панель маркетолога',
    nav_sale: 'Продажа',
    nav_stock: 'Склад и заказы',
    nav_analytics: 'Аналитика',
    logout: 'Выйти',
    night_mode: 'Ночной',
    day_mode: 'Дневной',
  },
  uz: {
    title: 'Marketolog paneli',
    nav_sale: 'Sotuv',
    nav_stock: 'Ombor va buyurtmalar',
    nav_analytics: 'Tahlil',
    logout: 'Chiqish',
    night_mode: 'Tungi',
    day_mode: 'Kunduzgi',
  }
}
const txt = computed(() => mgTexts[lang.value] || mgTexts.ru)

const tab = ref('sale')

const marketplaces = [
  { value: 'ozon', label: 'Ozon' },
  { value: 'yandex', label: 'Yandex Market' },
  { value: 'wildberries', label: 'Wildberries' },
  { value: 'uzum', label: 'Uzum Market' },
  { value: 'other', label: 'Другое' },
]
const channelLabels = { ozon: 'Ozon', yandex: 'Yandex', wildberries: 'Wildberries', uzum: 'Uzum', other: 'Маркетплейс' }
function channelLabel(v) { return channelLabels[v] || v }

const periods = [
  { v: 'daily', l: 'Сегодня' },
  { v: 'weekly', l: 'Неделя' },
  { v: 'monthly', l: 'Месяц' },
  { v: 'custom', l: 'Дата' },
]

const products = ref([])
const stock = ref([])
const items = ref([])
const saleUnit = ref('pack')
const channel = ref('ozon')
const note = ref('')
const submitting = ref(false)
const success = ref('')
const saleError = ref('')
const orders = ref([])

// Add-stock form
const stockProductId = ref('')
const stockQty = ref(1)
const addingStock = ref(false)

const total = computed(() => items.value.reduce((s, i) => s + i.price, 0))
const canSubmit = computed(() => items.value.length > 0 && !!channel.value)

const stockMap = computed(() => {
  const m = {}
  for (const s of stock.value) m[s.product_id] = s.quantity
  return m
})
function stockOf(productId) { return displayStock(productId, stockMap.value[productId] || 0) }
function qppOf(productId) { const p = products.value.find(x => x.id === productId); return p && p.quantity_per_pack > 0 ? p.quantity_per_pack : 1 }
function capsulesOf(productId) { return Math.floor(stockOf(productId) / qppOf(productId)) }
function availOf(productId) { return saleUnit.value === 'piece' ? stockOf(productId) : capsulesOf(productId) }
function productById(id) { return products.value.find(p => p.id === id) }
function cartQty(productId) { const i = items.value.find(x => x.product_id === productId); return i ? i.quantity : 0 }
function setUnit(u) { if (u !== saleUnit.value) { items.value = []; saleUnit.value = u } }

function formatPrice(p) { return new Intl.NumberFormat('ru-RU').format(Math.round(p || 0)) }
function orderTotal(order) { return order.items?.reduce((s, i) => s + i.price, 0) || 0 }

async function loadProducts() {
  try {
    const res = await api.get('/manager/products')
    products.value = res.data || []
    for (const p of products.value) p.price_per_pack = p.price_per_pill * p.quantity_per_pack
  } catch (e) { console.error(e) }
}
async function loadStock() {
  try { stock.value = (await api.get('/manager/stock')).data || [] } catch (e) { console.error(e) }
}
async function loadOrders() {
  try { orders.value = (await api.get('/manager/orders')).data || [] } catch (e) { console.error(e) }
}

function unitPrice(product) { return saleUnit.value === 'piece' ? (product.price_per_pill || 0) : (product.price_per_pack || 0) }
function addToCart(product) {
  if (!product) return
  if (availOf(product.id) <= cartQty(product.id)) return // not enough stock
  const existing = items.value.find(i => i.product_id === product.id)
  if (existing) { existing.quantity++; existing.price = existing.quantity * unitPrice(product) }
  else items.value.push({ product_id: product.id, name: product.name, quantity: 1, unit_type: saleUnit.value, price: unitPrice(product) })
}
function decQty(item) {
  item.quantity--
  if (item.quantity <= 0) { items.value = items.value.filter(i => i.product_id !== item.product_id) }
  else item.price = item.quantity * unitPrice(productById(item.product_id) || {})
}

async function submitSale() {
  if (!canSubmit.value) return
  submitting.value = true; success.value = ''; saleError.value = ''
  try {
    await api.post('/manager/sale', {
      items: items.value.map(i => ({ product_id: i.product_id, quantity: i.quantity, unit_type: i.unit_type || 'pack' })),
      offline_note: note.value,
      sales_channel: channel.value,
    })
    success.value = '1'
    items.value = []; note.value = ''; channel.value = 'ozon'; saleUnit.value = 'pack'
    loadStock(); loadOrders(); loadAnalytics()
  } catch (e) {
    saleError.value = e.response?.data?.error || 'Ошибка при записи'
  } finally { submitting.value = false }
}

async function addStock() {
  if (!stockProductId.value || stockQty.value < 1) return
  addingStock.value = true
  try {
    await api.post('/manager/stock', { product_id: stockProductId.value, quantity: stockQty.value, unit_type: 'pack' })
    stockProductId.value = ''; stockQty.value = 1
    loadStock()
  } catch (e) { alert(e.response?.data?.error || 'Ошибка') } finally { addingStock.value = false }
}

// Analytics
const period = ref('daily')
const customDate = ref('')
const analytics = ref(null)
const analyticsLoading = ref(false)

const soldCapsules = computed(() =>
  orders.value.reduce((s, o) => s + (o.items || []).reduce((a, i) => a + (i.quantity || 0), 0), 0)
)

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const params = { period: period.value }
    if (period.value === 'custom') params.date = customDate.value
    analytics.value = (await api.get('/manager/analytics', { params })).data
  } catch (e) { console.error(e) } finally { analyticsLoading.value = false }
}
function selectPeriod(p) { period.value = p; if (p !== 'custom') loadAnalytics() }

function logout() { authStore.workerLogout(); router.push('/admin/login') }

let stockPoll = null
onMounted(() => {
  loadProducts(); loadStock(); loadOrders(); loadAnalytics()
  stockPoll = setInterval(() => {
    if (tab.value === 'sale' || tab.value === 'stock') loadStock()
  }, 7000)
})
onUnmounted(() => { if (stockPoll) clearInterval(stockPoll) })

// Real-time: a marketolog's debt orders/analytics update without a reload.
watch(() => realtime.ordersVersion, () => { loadOrders(); loadAnalytics() })
</script>
