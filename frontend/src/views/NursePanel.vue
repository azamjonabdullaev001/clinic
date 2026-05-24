<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-3xl mx-auto px-4 sm:px-6 flex items-center justify-between h-16">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-teal-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </div>
          <div>
            <h1 class="text-base font-bold text-gray-800">{{ txt.title }}</h1>
            <p v-if="authStore.worker" class="text-xs text-gray-400">{{ authStore.worker.name }}</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <!-- Language switcher -->
          <div class="flex items-center bg-gray-100 rounded-lg p-0.5 gap-0.5">
            <button @click="lang = 'ru'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'ru' ? 'bg-teal-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">RU</button>
            <button @click="lang = 'uz'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'uz' ? 'bg-teal-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">UZ</button>
          </div>
          <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">{{ txt.logout }}</button>
        </div>
      </div>
    </header>

    <div class="max-w-3xl mx-auto px-4 sm:px-6 mt-6 pb-12 space-y-6">

      <!-- Create Order -->
      <div class="bg-white rounded-xl shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b flex items-center gap-3">
          <div class="w-8 h-8 bg-teal-100 rounded-lg flex items-center justify-center">
            <svg class="w-4 h-4 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
            </svg>
          </div>
          <h2 class="font-bold text-gray-800">{{ txt.create_order }}</h2>
        </div>

        <!-- Success card -->
        <div v-if="createdCode" class="mx-6 my-4 bg-teal-50 border-2 border-teal-300 rounded-xl p-6 text-center">
          <svg class="w-12 h-12 text-teal-500 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-sm text-teal-700 mb-2">{{ txt.success_title }}</p>
          <p class="text-5xl font-bold tracking-[0.3em] text-teal-700 mb-3">{{ createdCode }}</p>
          <p class="text-xs text-teal-600">{{ txt.success_desc }}</p>
          <button @click="createdCode = ''" class="mt-4 text-sm text-teal-700 underline hover:no-underline">{{ txt.new_order }}</button>
        </div>

        <div v-else class="px-6 py-5 space-y-5">
          <!-- Patient name -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">{{ txt.patient_fname }} <span class="text-red-400">*</span></label>
              <input v-model="patientFName" type="text"
                class="w-full border border-gray-300 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
                :placeholder="txt.patient_fname" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">{{ txt.patient_lname }} <span class="text-red-400">*</span></label>
              <input v-model="patientLName" type="text"
                class="w-full border border-gray-300 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
                :placeholder="txt.patient_lname" />
            </div>
          </div>

          <!-- Add item -->
          <div>
            <label class="block text-sm font-medium text-gray-600 mb-2">{{ txt.add_product }}</label>
            <div class="flex gap-2 flex-wrap items-end">
              <div class="flex-1 min-w-[180px]">
                <select v-model="selProductId"
                  class="w-full border border-gray-300 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500 transition">
                  <option value="">{{ txt.select_product }}</option>
                  <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
                </select>
              </div>
              <div class="w-24">
                <input v-model.number="selQty" type="number" min="1"
                  class="w-full border border-gray-300 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
              </div>
              <div class="flex items-end pb-1">
                <span class="text-sm text-gray-600 font-medium px-2">{{ txt.pack }}</span>
              </div>
              <button @click="addItem" :disabled="!selProductId || selQty < 1"
                class="bg-teal-600 text-white px-5 py-3 rounded-xl hover:bg-teal-700 transition font-medium disabled:opacity-40">
                + {{ txt.add }}
              </button>
            </div>
          </div>

          <!-- Items list -->
          <div v-if="items.length" class="border rounded-xl overflow-hidden">
            <div v-for="(item, idx) in items" :key="idx"
              class="flex items-center justify-between px-4 py-3 border-b last:border-0 bg-gray-50">
              <div>
                <span class="font-medium text-gray-800">{{ item.name }}</span>
                <span class="text-gray-500 ml-2">× {{ item.quantity }} {{ item.unit_type === 'pack' ? txt.pack : txt.piece }}</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-semibold text-gray-700">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                <button @click="items.splice(idx, 1)" class="text-red-400 hover:text-red-600 transition">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                </button>
              </div>
            </div>
            <div class="flex justify-between px-4 py-2.5 bg-white font-semibold text-sm">
              <span class="text-gray-600">{{ txt.total }}:</span>
              <span class="text-teal-700">{{ formatPrice(items.reduce((s, i) => s + i.price, 0)) }} {{ txt.sum }}</span>
            </div>
          </div>

          <div v-if="formError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ formError }}</div>

          <button @click="submitOrder" :disabled="!canSubmit || submitting"
            class="w-full bg-teal-600 text-white py-3.5 rounded-xl hover:bg-teal-700 transition font-semibold text-base disabled:opacity-40">
            {{ submitting ? txt.submitting : txt.submit }}
          </button>
        </div>
      </div>

      <!-- My Orders -->
      <div>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-bold text-gray-800">{{ txt.my_orders }}</h2>
          <button @click="loadOrders" class="text-sm text-teal-600 hover:text-teal-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            {{ txt.refresh }}
          </button>
        </div>

        <div class="space-y-3">
          <div v-for="order in myOrders" :key="order.id"
            class="bg-white rounded-xl shadow-sm p-5">
            <div class="flex items-center justify-between mb-3">
              <div>
                <div class="flex items-center gap-2 mb-1">
                  <span class="text-2xl font-bold tracking-[0.2em] text-teal-700">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                </div>
                <p class="font-semibold text-gray-800">{{ order.patient_first_name }} {{ order.patient_last_name }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
              </div>
              <div class="text-right">
                <p class="text-xs text-gray-400">{{ txt.total }}</p>
                <p class="font-bold text-gray-800">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
              </div>
            </div>
            <div class="border-t pt-3 space-y-1">
              <div v-for="item in order.items.filter(i => i.quantity > 0)" :key="item.id" class="flex justify-between text-sm text-gray-600">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ txt.pack }}</span></span>
                <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
              </div>
            </div>
          </div>

          <div v-if="myOrders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">
            <svg class="w-10 h-10 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            {{ txt.no_orders }}
          </div>
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

const lang = ref(localStorage.getItem('nurseLang') || 'ru')
const watchLang = () => { localStorage.setItem('nurseLang', lang.value) }

const texts = {
  ru: {
    title: 'Медицинский персонал',
    create_order: 'Оформить заказ',
    patient_fname: 'Имя пациента',
    patient_lname: 'Фамилия пациента',
    add_product: 'Добавить препарат',
    select_product: 'Выберите препарат',
    quantity: 'Количество',
    pack: 'капс.',
    piece: 'шт',
    add: 'Добавить',
    total: 'Итого',
    sum: 'сўм',
    submit: 'Оформить заказ',
    submitting: 'Оформление...',
    success_title: 'Заказ оформлен!',
    success_desc: 'Передайте этот код пациенту для получения препаратов',
    new_order: 'Новый заказ',
    my_orders: 'Мои заказы',
    refresh: 'Обновить',
    no_orders: 'Нет заказов',
    logout: 'Выйти',
    status_pending: 'Ожидает',
    status_delivered: 'Выдан',
    status_cancelled: 'Отменён',
  },
  uz: {
    title: 'Tibbiy xodim',
    create_order: 'Buyurtma rasmiylashtirish',
    patient_fname: 'Bemorning ismi',
    patient_lname: 'Bemorning familiyasi',
    add_product: "Dori qo'shish",
    select_product: 'Dori tanlang',
    quantity: 'Miqdor',
    pack: 'kapsula',
    piece: 'dona',
    add: "Qo'shish",
    total: 'Jami',
    sum: "so'm",
    submit: 'Buyurtma berish',
    submitting: 'Yuborilmoqda...',
    success_title: 'Buyurtma rasmiylashtirildi!',
    success_desc: 'Ushbu kodni bemorga dori olish uchun bering',
    new_order: 'Yangi buyurtma',
    my_orders: 'Mening buyurtmalarim',
    refresh: 'Yangilash',
    no_orders: "Buyurtmalar yo'q",
    logout: 'Chiqish',
    status_pending: 'Kutilmoqda',
    status_delivered: 'Berildi',
    status_cancelled: 'Bekor qilindi',
  }
}

const txt = computed(() => {
  watchLang()
  return texts[lang.value] || texts.ru
})

const products = ref([])
const myOrders = ref([])
const patientFName = ref('')
const patientLName = ref('')
const selProductId = ref('')
const selQty = ref(1)
const selUnit = ref('pack')
const items = ref([])
const submitting = ref(false)
const formError = ref('')
const createdCode = ref('')

const canSubmit = computed(() =>
  patientFName.value.trim() && patientLName.value.trim() && items.value.length > 0
)

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

function statusLabel(status) {
  const t = txt.value
  const m = { pending: t.status_pending, delivered: t.status_delivered, cancelled: t.status_cancelled }
  return m[status] || status
}

function statusClass(status) {
  const m = {
    pending: 'bg-yellow-100 text-yellow-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-700',
  }
  return m[status] || 'bg-gray-100 text-gray-700'
}

function addItem() {
  if (!selProductId.value || selQty.value < 1) return
  const product = products.value.find(p => p.id === selProductId.value)
  if (!product) return
  const price = product.price_per_pack * selQty.value
  items.value.push({
    product_id: product.id,
    name: product.name,
    quantity: selQty.value,
    unit_type: 'pack',
    price,
  })
  selProductId.value = ''
  selQty.value = 1
}

async function submitOrder() {
  if (!canSubmit.value) return
  formError.value = ''
  submitting.value = true
  try {
    const res = await api.post('/nurse/orders', {
      patient_first_name: patientFName.value.trim(),
      patient_last_name: patientLName.value.trim(),
      items: items.value.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_type: i.unit_type,
      })),
    })
    createdCode.value = res.data.order_code
    patientFName.value = ''
    patientLName.value = ''
    items.value = []
    loadOrders()
  } catch (e) {
    formError.value = e.response?.data?.error || 'Ошибка при оформлении'
  } finally {
    submitting.value = false
  }
}

async function loadOrders() {
  try {
    const res = await api.get('/nurse/orders')
    myOrders.value = res.data || []
  } catch (e) { console.error(e) }
}

async function loadProducts() {
  try {
    const res = await api.get('/nurse/products')
    products.value = res.data || []
    for (const p of products.value) {
      p.price_per_pack = p.price_per_pill * p.quantity_per_pack
    }
  } catch (e) { console.error(e) }
}

function logout() {
  authStore.workerLogout()
  router.push('/admin/login')
}

onMounted(() => {
  loadProducts()
  loadOrders()
})
</script>
