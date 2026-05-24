<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 flex items-center justify-between h-16">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-purple-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div>
            <h1 class="text-base font-bold text-gray-800">Панель менеджера</h1>
            <p v-if="authStore.worker" class="text-xs text-gray-400">{{ authStore.worker.name }}</p>
          </div>
        </div>
        <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">Выйти</button>
      </div>
    </header>

    <div class="max-w-4xl mx-auto px-4 sm:px-6 mt-6 pb-12 space-y-6">

      <!-- ===== Create marketplace sale ===== -->
      <div class="bg-white rounded-xl shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b flex items-center gap-3">
          <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center">
            <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4"/></svg>
          </div>
          <h2 class="font-bold text-gray-800">Продажа на маркетплейсе</h2>
        </div>
        <div class="px-6 py-5 space-y-4">
          <!-- Add item -->
          <div class="flex gap-2 flex-wrap items-end">
            <div class="flex-1 min-w-[180px]">
              <label class="block text-xs font-medium text-gray-500 mb-1">Препарат</label>
              <select v-model="productId" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500">
                <option value="">Выберите препарат</option>
                <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
              </select>
            </div>
            <div class="w-24">
              <label class="block text-xs font-medium text-gray-500 mb-1">Кол-во</label>
              <input v-model.number="qty" type="number" min="1" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500" />
            </div>
            <div class="flex items-end pb-1.5"><span class="text-sm text-gray-600 font-medium px-1">капс.</span></div>
            <button @click="addItem" :disabled="!productId || qty < 1"
              class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition text-sm font-medium disabled:opacity-40">+ Добавить</button>
          </div>

          <!-- Items -->
          <div v-if="items.length" class="border rounded-xl overflow-hidden">
            <div v-for="(item, idx) in items" :key="idx" class="flex items-center justify-between px-4 py-2.5 border-b last:border-0 bg-gray-50">
              <div>
                <span class="font-medium text-gray-800 text-sm">{{ item.name }}</span>
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} капс.</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-semibold text-gray-700 text-sm">{{ formatPrice(item.price) }} сўм</span>
                <button @click="items.splice(idx, 1)" class="text-red-400 hover:text-red-600 transition">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between px-4 py-2 bg-white">
              <span class="text-sm font-semibold text-gray-700">Итого:</span>
              <span class="font-bold text-purple-600">{{ formatPrice(total) }} сўм</span>
            </div>
          </div>

          <!-- Marketplace -->
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
            {{ submitting ? 'Запись...' : 'Записать продажу' }}
          </button>

          <div v-if="success" class="bg-purple-50 border border-purple-200 rounded-lg px-4 py-3 text-sm text-purple-700">
            Продажа записана. Код: <strong>{{ success }}</strong>
          </div>
        </div>
      </div>

      <!-- ===== My analytics ===== -->
      <div class="bg-white rounded-xl shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b flex items-center justify-between">
          <h2 class="font-bold text-gray-800">Моя аналитика</h2>
          <div class="flex gap-1.5 flex-wrap items-center">
            <button v-for="p in periods" :key="p.v" @click="selectPeriod(p.v)"
              :class="period === p.v ? 'bg-indigo-600 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'"
              class="px-3 py-1.5 rounded-lg text-xs font-medium transition">{{ p.l }}</button>
            <input v-if="period === 'custom'" v-model="customDate" type="date" @change="loadAnalytics"
              class="border border-gray-300 rounded-lg px-2 py-1 text-xs" />
          </div>
        </div>
        <div class="px-6 py-5">
          <div v-if="analyticsLoading" class="text-gray-400 text-sm py-6 text-center">Загрузка...</div>
          <template v-else-if="analytics">
            <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Продаж</p><p class="text-2xl font-bold text-gray-800">{{ analytics.total_orders }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Выручка</p><p class="text-lg font-bold text-emerald-600">{{ formatPrice(analytics.total_revenue) }} сўм</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Создано</p><p class="text-2xl font-bold text-purple-600">{{ analytics.created_count }}</p></div>
            </div>
            <div class="flex items-end gap-0.5 h-28 pt-3">
              <div v-for="(pt,i) in analytics.points" :key="i"
                class="flex-1 bg-indigo-200 hover:bg-indigo-400 rounded-t transition-all min-h-[2px]"
                :style="{ height: (pt.revenue / maxRevenue * 100) + '%' }"
                :title="pt.label + ': ' + formatPrice(pt.revenue) + ' сўм'"></div>
            </div>
          </template>
        </div>
      </div>

      <!-- ===== My sales ===== -->
      <div>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-bold text-gray-800">Мои продажи</h2>
          <button @click="loadOrders" class="text-sm text-purple-600 hover:text-purple-800">Обновить</button>
        </div>
        <div class="space-y-3">
          <div v-for="order in orders" :key="order.id" class="bg-white rounded-xl shadow-sm p-5">
            <div class="flex justify-between items-start gap-3">
              <div>
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-lg font-bold text-purple-700 tracking-widest">{{ order.order_code }}</span>
                  <span v-if="order.sales_channel" class="text-xs font-semibold px-2 py-0.5 rounded bg-purple-100 text-purple-700">{{ channelLabel(order.sales_channel) }}</span>
                </div>
                <p v-if="order.offline_note" class="text-sm text-gray-600">{{ order.offline_note }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
              </div>
              <p class="font-bold text-gray-800">{{ formatPrice(orderTotal(order)) }} сўм</p>
            </div>
            <div class="mt-3 pt-3 border-t border-gray-100 space-y-1">
              <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm text-gray-600">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} капс.</span></span>
                <span class="font-medium">{{ formatPrice(item.price) }} сўм</span>
              </div>
            </div>
          </div>
          <div v-if="orders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">Продаж пока нет</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

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
const productId = ref('')
const qty = ref(1)
const items = ref([])
const channel = ref('ozon')
const note = ref('')
const submitting = ref(false)
const success = ref('')
const orders = ref([])

const total = computed(() => items.value.reduce((s, i) => s + i.price, 0))
const canSubmit = computed(() => items.value.length > 0 && !!channel.value)

function formatPrice(p) { return new Intl.NumberFormat('ru-RU').format(Math.round(p || 0)) }
function orderTotal(order) { return order.items?.reduce((s, i) => s + i.price, 0) || 0 }

async function loadProducts() {
  try {
    const res = await api.get('/manager/products')
    products.value = res.data || []
    for (const p of products.value) p.price_per_pack = p.price_per_pill * p.quantity_per_pack
  } catch (e) { console.error(e) }
}

async function loadOrders() {
  try {
    const res = await api.get('/manager/orders')
    orders.value = res.data || []
  } catch (e) { console.error(e) }
}

function addItem() {
  if (!productId.value || qty.value < 1) return
  const product = products.value.find(p => p.id === productId.value)
  if (!product) return
  items.value.push({ product_id: product.id, name: product.name, quantity: qty.value, price: product.price_per_pack * qty.value })
  productId.value = ''
  qty.value = 1
}

async function submitSale() {
  if (!canSubmit.value) return
  submitting.value = true
  success.value = ''
  try {
    const res = await api.post('/manager/sale', {
      items: items.value.map(i => ({ product_id: i.product_id, quantity: i.quantity, unit_type: 'pack' })),
      offline_note: note.value,
      sales_channel: channel.value,
    })
    success.value = res.data.order_code
    items.value = []
    note.value = ''
    channel.value = 'ozon'
    loadOrders()
    loadAnalytics()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при записи')
  } finally { submitting.value = false }
}

// Analytics
const period = ref('daily')
const customDate = ref('')
const analytics = ref(null)
const analyticsLoading = ref(false)
const maxRevenue = computed(() => Math.max(1, ...(analytics.value?.points || []).map(p => p.revenue)))

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const params = { period: period.value }
    if (period.value === 'custom') params.date = customDate.value
    const res = await api.get('/manager/analytics', { params })
    analytics.value = res.data
  } catch (e) { console.error(e) } finally { analyticsLoading.value = false }
}

function selectPeriod(p) {
  period.value = p
  if (p !== 'custom') loadAnalytics()
}

function logout() {
  authStore.workerLogout()
  router.push('/admin/login')
}

onMounted(() => {
  loadProducts()
  loadOrders()
  loadAnalytics()
})
</script>
