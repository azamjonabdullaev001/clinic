<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b sticky top-0 z-10">
      <div class="max-w-5xl mx-auto px-4 sm:px-6 flex items-center justify-between h-16">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
          </div>
          <div>
            <h1 class="text-base font-bold text-gray-800">{{ txt.title }}</h1>
            <p v-if="authStore.doctor" class="text-xs text-gray-400">{{ authStore.doctor.name }}</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex items-center bg-gray-100 rounded-lg p-0.5 gap-0.5">
            <button @click="lang = 'ru'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'ru' ? 'bg-blue-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">RU</button>
            <button @click="lang = 'uz'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'uz' ? 'bg-blue-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">UZ</button>
          </div>
          <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">{{ txt.logout }}</button>
        </div>
      </div>
      <!-- Tabs -->
      <div class="max-w-5xl mx-auto px-4 sm:px-6 flex gap-1 border-t">
        <button v-for="tab in tabs" :key="tab.key" @click="activeTab = tab.key"
          class="px-4 py-3 text-sm font-medium border-b-2 transition-colors"
          :class="activeTab === tab.key ? 'border-blue-600 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700'">
          {{ tab.label }}
          <span v-if="tab.key === 'order' && cartItems.length" class="ml-1.5 bg-blue-600 text-white text-xs rounded-full px-1.5 py-0.5">{{ cartItems.length }}</span>
        </button>
      </div>
    </header>

    <div class="max-w-5xl mx-auto px-4 sm:px-6 mt-4 pb-12">

      <!-- ===== ORDER TAB ===== -->
      <div v-if="activeTab === 'order'">

        <!-- Success card -->
        <div v-if="createdCode" class="bg-blue-50 border-2 border-blue-300 rounded-xl p-8 text-center mb-6">
          <svg class="w-14 h-14 text-blue-500 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-sm text-blue-700 mb-2">{{ txt.success_title }}</p>
          <p class="text-6xl font-bold tracking-[0.3em] text-blue-700 mb-3">{{ createdCode }}</p>
          <p class="text-xs text-blue-600">{{ txt.success_desc }}</p>
          <button @click="createdCode = ''" class="mt-4 text-sm text-blue-700 underline hover:no-underline">{{ txt.new_order }}</button>
        </div>

        <div v-else class="space-y-5">
          <!-- Product catalog -->
          <div class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-5 py-4 border-b flex items-center gap-2">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
              </svg>
              <h2 class="font-bold text-gray-800">{{ txt.catalog }}</h2>
              <span class="ml-auto text-xs text-gray-400">{{ txt.click_to_add }}</span>
            </div>
            <div class="p-4">
              <div v-if="productsLoading" class="flex justify-center py-8 text-gray-400">{{ txt.loading }}</div>
              <div v-else class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
                <div v-for="product in products" :key="product.id"
                  @click="addToCart(product)"
                  class="relative bg-gray-50 rounded-xl p-3 cursor-pointer hover:bg-blue-50 hover:shadow-md transition-all border-2 border-transparent hover:border-blue-200 group select-none"
                  :class="cartQty(product.id) > 0 ? 'border-blue-400 bg-blue-50' : ''">
                  <!-- Cart badge -->
                  <div v-if="cartQty(product.id) > 0"
                    class="absolute -top-2 -right-2 w-6 h-6 bg-blue-600 text-white text-xs font-bold rounded-full flex items-center justify-center shadow-md">
                    {{ cartQty(product.id) }}
                  </div>
                  <!-- Product image -->
                  <div class="w-full aspect-square rounded-lg overflow-hidden bg-white mb-2 flex items-center justify-center">
                    <img v-if="product.image_path" :src="`/uploads/${product.image_path}`"
                      class="w-full h-full object-cover" :alt="product.name" />
                    <div v-else class="w-full h-full bg-blue-100 flex items-center justify-center">
                      <svg class="w-8 h-8 text-blue-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                      </svg>
                    </div>
                  </div>
                  <p class="text-xs font-semibold text-gray-800 leading-tight mb-1 line-clamp-2">{{ product.name }}</p>
                  <p v-if="product.description" class="text-xs text-gray-400 leading-tight mb-1 line-clamp-2">{{ product.description }}</p>
                  <p class="text-xs font-bold text-blue-600">{{ formatPrice(product.price_per_pack) }} {{ txt.sum }}</p>
                  <p class="text-xs text-gray-400">{{ product.quantity_per_pack }} {{ txt.piece }}</p>
                  <!-- Add indicator -->
                  <div class="mt-1.5 text-center">
                    <span class="text-xs text-blue-500 font-medium group-hover:text-blue-700 transition">
                      {{ cartQty(product.id) > 0 ? '+ ' + txt.add_more : txt.tap_to_add }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Cart / Order form -->
          <div v-if="cartItems.length > 0" class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-5 py-4 border-b flex items-center gap-2">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              <h2 class="font-bold text-gray-800">{{ txt.order_form }}</h2>
            </div>
            <div class="px-5 py-4 space-y-4">
              <!-- Patient name -->
              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.patient_fname }} <span class="text-red-400">*</span></label>
                  <input v-model="patientFName" type="text"
                    class="w-full border border-gray-300 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
                    :placeholder="txt.patient_fname" />
                </div>
                <div>
                  <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.patient_lname }} <span class="text-red-400">*</span></label>
                  <input v-model="patientLName" type="text"
                    class="w-full border border-gray-300 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
                    :placeholder="txt.patient_lname" />
                </div>
              </div>

              <!-- Cart items -->
              <div class="border rounded-xl overflow-hidden">
                <div v-for="item in cartItems" :key="item.product_id"
                  class="flex items-center gap-3 px-4 py-3 border-b last:border-0 bg-gray-50">
                  <span class="flex-1 text-sm font-medium text-gray-800 truncate">{{ item.name }}</span>
                  <div class="flex items-center gap-1.5">
                    <button @click.stop="decreaseCart(item.product_id)"
                      class="w-7 h-7 bg-gray-200 hover:bg-gray-300 rounded-lg flex items-center justify-center text-gray-700 transition text-sm font-bold">−</button>
                    <span class="w-8 text-center text-sm font-bold text-gray-800">{{ item.quantity }}</span>
                    <button @click.stop="increaseCart(item.product_id)"
                      class="w-7 h-7 bg-blue-100 hover:bg-blue-200 rounded-lg flex items-center justify-center text-blue-700 transition text-sm font-bold">+</button>
                  </div>
                  <span class="text-xs text-gray-400 w-16 text-right">{{ txt.pack }}</span>
                  <span class="font-semibold text-gray-700 text-sm w-24 text-right">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                  <button @click.stop="removeFromCart(item.product_id)" class="text-red-400 hover:text-red-600 transition flex-shrink-0">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                  </button>
                </div>
                <div class="flex justify-between px-4 py-3 bg-white font-bold text-sm">
                  <span class="text-gray-600">{{ txt.total }}:</span>
                  <span class="text-blue-700">{{ formatPrice(cartTotal) }} {{ txt.sum }}</span>
                </div>
              </div>

              <div v-if="formError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ formError }}</div>

              <button @click="submitOrder" :disabled="!canSubmit || submitting"
                class="w-full bg-blue-600 text-white py-3.5 rounded-xl hover:bg-blue-700 transition font-bold text-base disabled:opacity-40">
                {{ submitting ? txt.submitting : txt.submit }}
              </button>
            </div>
          </div>

          <div v-else class="bg-white rounded-xl shadow-sm p-8 text-center text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <p class="text-sm">{{ txt.catalog_hint }}</p>
          </div>
        </div>
      </div>

      <!-- ===== MY ORDERS TAB ===== -->
      <div v-if="activeTab === 'history'">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-bold text-gray-800">{{ txt.my_orders }}</h2>
          <button @click="loadOrders" class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            {{ txt.refresh }}
          </button>
        </div>

        <div class="space-y-3">
          <div v-for="order in myOrders" :key="order.id"
            class="bg-white rounded-xl shadow-sm p-5">
            <div class="flex items-start justify-between mb-3">
              <div>
                <div class="flex items-center gap-2 mb-1">
                  <span class="text-2xl font-bold tracking-[0.2em] text-blue-700">{{ order.order_code }}</span>
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
              <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm text-gray-600">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'pack' ? txt.pack : txt.piece }}</span></span>
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

      <!-- ===== ANALYTICS TAB ===== -->
      <div v-if="activeTab === 'analytics'">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-bold text-gray-800">{{ txt.analytics }}</h2>
          <button @click="loadAnalytics" class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            {{ txt.refresh }}
          </button>
        </div>

        <div v-if="analyticsLoading" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">{{ txt.loading }}</div>

        <div v-else class="space-y-4">
          <!-- Stats cards -->
          <div class="grid grid-cols-2 gap-4">
            <div class="bg-white rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-400 mb-1">{{ txt.total_orders }}</p>
              <p class="text-3xl font-bold text-blue-700">{{ analytics.total_orders || 0 }}</p>
            </div>
            <div class="bg-white rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-400 mb-1">{{ txt.total_revenue }}</p>
              <p class="text-2xl font-bold text-green-700">{{ formatPrice(analytics.total_revenue || 0) }}</p>
              <p class="text-xs text-gray-400">{{ txt.sum }}</p>
            </div>
          </div>

          <!-- Top products -->
          <div class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-5 py-4 border-b">
              <h3 class="font-bold text-gray-800">{{ txt.top_products }}</h3>
            </div>
            <div v-if="analytics.top_products?.length">
              <div v-for="(p, i) in analytics.top_products" :key="p.product_id"
                class="flex items-center gap-3 px-5 py-3 border-b last:border-0">
                <span class="w-6 h-6 bg-blue-100 text-blue-700 text-xs font-bold rounded-full flex items-center justify-center flex-shrink-0">{{ i + 1 }}</span>
                <span class="flex-1 text-sm font-medium text-gray-800">{{ p.product_name }}</span>
                <div class="text-right">
                  <p class="text-sm font-bold text-gray-700">{{ formatPrice(p.revenue) }} {{ txt.sum }}</p>
                  <p class="text-xs text-gray-400">{{ txt.orders }}: {{ p.order_count }}</p>
                </div>
              </div>
            </div>
            <div v-else class="px-5 py-8 text-center text-gray-400 text-sm">{{ txt.no_data }}</div>
          </div>

          <!-- Recent orders in analytics -->
          <div class="bg-white rounded-xl shadow-sm overflow-hidden">
            <div class="px-5 py-4 border-b">
              <h3 class="font-bold text-gray-800">{{ txt.recent_orders }}</h3>
            </div>
            <div v-if="analytics.recent_orders?.length">
              <div v-for="order in analytics.recent_orders" :key="order.id"
                class="flex items-center gap-3 px-5 py-3 border-b last:border-0">
                <div class="flex-1">
                  <div class="flex items-center gap-2">
                    <span class="font-bold text-blue-700 tracking-widest text-sm">{{ order.order_code }}</span>
                    <span :class="statusClass(order.status)" class="text-xs px-1.5 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                  </div>
                  <p class="text-xs text-gray-500">{{ order.patient_first_name }} {{ order.patient_last_name }}</p>
                  <p class="text-xs text-gray-400">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                </div>
                <p class="font-bold text-gray-700 text-sm">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
              </div>
            </div>
            <div v-else class="px-5 py-8 text-center text-gray-400 text-sm">{{ txt.no_orders }}</div>
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

const lang = ref(localStorage.getItem('doctorLang') || 'ru')
const watchLang = () => { localStorage.setItem('doctorLang', lang.value) }

const texts = {
  ru: {
    title: 'Панель врача',
    catalog: 'Каталог препаратов',
    click_to_add: 'Нажмите на препарат для добавления',
    tap_to_add: '+ Нажать для добавления',
    add_more: 'Ещё одну',
    order_form: 'Оформление заказа',
    patient_fname: 'Имя пациента',
    patient_lname: 'Фамилия пациента',
    pack: 'упак.',
    piece: 'шт.',
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
    loading: 'Загрузка...',
    catalog_hint: 'Нажмите на препарат в каталоге чтобы добавить его в заказ',
    analytics: 'Аналитика',
    total_orders: 'Всего заказов',
    total_revenue: 'Общая выручка',
    top_products: 'Топ препаратов',
    recent_orders: 'Последние заказы',
    orders: 'заказов',
    no_data: 'Нет данных',
    status_pending: 'Ожидает',
    status_confirmed: 'Подтверждён',
    status_shipped: 'Отправлен',
    status_in_transit: 'В пути',
    status_delivered: 'Выдан',
    status_cancelled: 'Отменён',
    add: 'Добавить',
  },
  uz: {
    title: 'Shifokor paneli',
    catalog: 'Dorilar katalogi',
    click_to_add: 'Dori qo\'shish uchun bosing',
    tap_to_add: '+ Qo\'shish uchun bosing',
    add_more: 'Yana bittasi',
    order_form: 'Buyurtmani rasmiylashtirish',
    patient_fname: 'Bemorning ismi',
    patient_lname: 'Bemorning familiyasi',
    pack: 'quti',
    piece: 'dona',
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
    loading: 'Yuklanmoqda...',
    catalog_hint: "Buyurtmaga qo'shish uchun katalogdagi dorini bosing",
    analytics: 'Tahlil',
    total_orders: 'Jami buyurtmalar',
    total_revenue: 'Umumiy daromad',
    top_products: 'Top dorilar',
    recent_orders: 'Oxirgi buyurtmalar',
    orders: 'buyurtma',
    no_data: "Ma'lumot yo'q",
    status_pending: 'Kutilmoqda',
    status_confirmed: 'Tasdiqlangan',
    status_shipped: 'Yuborilgan',
    status_in_transit: "Yo'lda",
    status_delivered: 'Berildi',
    status_cancelled: 'Bekor qilindi',
    add: "Qo'shish",
  }
}

const txt = computed(() => {
  watchLang()
  return texts[lang.value] || texts.ru
})

const activeTab = ref('order')
const tabs = computed(() => [
  { key: 'order', label: lang.value === 'ru' ? 'Новый заказ' : 'Yangi buyurtma' },
  { key: 'history', label: lang.value === 'ru' ? 'Мои заказы' : 'Buyurtmalarim' },
  { key: 'analytics', label: lang.value === 'ru' ? 'Аналитика' : 'Tahlil' },
])

// Products & cart
const products = ref([])
const productsLoading = ref(false)
const cartItems = ref([]) // { product_id, name, quantity, price_per_pack, price }

function cartQty(productId) {
  const item = cartItems.value.find(i => i.product_id === productId)
  return item ? item.quantity : 0
}

function addToCart(product) {
  const existing = cartItems.value.find(i => i.product_id === product.id)
  if (existing) {
    existing.quantity++
    existing.price = existing.price_per_pack * existing.quantity
  } else {
    cartItems.value.push({
      product_id: product.id,
      name: product.name,
      quantity: 1,
      price_per_pack: product.price_per_pack,
      price: product.price_per_pack,
    })
  }
}

function increaseCart(productId) {
  const item = cartItems.value.find(i => i.product_id === productId)
  if (item) {
    item.quantity++
    item.price = item.price_per_pack * item.quantity
  }
}

function decreaseCart(productId) {
  const idx = cartItems.value.findIndex(i => i.product_id === productId)
  if (idx === -1) return
  if (cartItems.value[idx].quantity <= 1) {
    cartItems.value.splice(idx, 1)
  } else {
    cartItems.value[idx].quantity--
    cartItems.value[idx].price = cartItems.value[idx].price_per_pack * cartItems.value[idx].quantity
  }
}

function removeFromCart(productId) {
  const idx = cartItems.value.findIndex(i => i.product_id === productId)
  if (idx !== -1) cartItems.value.splice(idx, 1)
}

const cartTotal = computed(() => cartItems.value.reduce((s, i) => s + i.price, 0))

// Order form
const patientFName = ref('')
const patientLName = ref('')
const submitting = ref(false)
const formError = ref('')
const createdCode = ref('')

const canSubmit = computed(() =>
  patientFName.value.trim() && patientLName.value.trim() && cartItems.value.length > 0
)

// Orders history
const myOrders = ref([])

// Analytics
const analytics = ref({})
const analyticsLoading = ref(false)

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

function statusLabel(status) {
  const t = txt.value
  const m = {
    pending: t.status_pending,
    confirmed: t.status_confirmed,
    shipped: t.status_shipped,
    in_transit: t.status_in_transit,
    delivered: t.status_delivered,
    cancelled: t.status_cancelled,
  }
  return m[status] || status
}

function statusClass(status) {
  const m = {
    pending: 'bg-yellow-100 text-yellow-700',
    confirmed: 'bg-blue-100 text-blue-700',
    shipped: 'bg-purple-100 text-purple-700',
    in_transit: 'bg-orange-100 text-orange-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-700',
  }
  return m[status] || 'bg-gray-100 text-gray-700'
}

async function submitOrder() {
  if (!canSubmit.value) return
  formError.value = ''
  submitting.value = true
  try {
    const res = await api.post('/doctor/orders', {
      patient_first_name: patientFName.value.trim(),
      patient_last_name: patientLName.value.trim(),
      items: cartItems.value.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_type: 'pack',
      })),
    })
    createdCode.value = res.data.order_code
    patientFName.value = ''
    patientLName.value = ''
    cartItems.value = []
    loadOrders()
  } catch (e) {
    formError.value = e.response?.data?.error || 'Ошибка при оформлении'
  } finally {
    submitting.value = false
  }
}

async function loadOrders() {
  try {
    const res = await api.get('/doctor/orders')
    myOrders.value = res.data || []
  } catch (e) { console.error(e) }
}

async function loadProducts() {
  productsLoading.value = true
  try {
    const res = await api.get('/doctor/products')
    products.value = res.data || []
    for (const p of products.value) {
      p.price_per_pack = p.price_per_pill * p.quantity_per_pack
    }
  } catch (e) { console.error(e) }
  finally { productsLoading.value = false }
}

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const res = await api.get('/doctor/analytics')
    analytics.value = res.data || {}
  } catch (e) { console.error(e) }
  finally { analyticsLoading.value = false }
}

function logout() {
  authStore.doctorLogout()
  router.push('/admin/login')
}

onMounted(() => {
  loadProducts()
  loadOrders()
  loadAnalytics()
})
</script>
