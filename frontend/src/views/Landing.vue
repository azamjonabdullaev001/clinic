<template>
  <div>
    <Navbar />
    <CartDrawer @checkout="placeOrder" />

    <!-- Hero -->
    <section class="relative bg-gradient-to-br from-teal-700 via-teal-600 to-emerald-600 text-white overflow-hidden">
      <div class="absolute inset-0 opacity-10">
        <div class="absolute top-10 left-10 w-72 h-72 bg-white rounded-full blur-3xl"></div>
        <div class="absolute bottom-10 right-10 w-96 h-96 bg-white rounded-full blur-3xl"></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24 md:py-32 text-center">
        <h1 class="text-4xl md:text-6xl font-extrabold mb-6 leading-tight">
          Здоровые волосы —<br />это <span class="text-yellow-300">реальность</span>
        </h1>
        <p class="text-xl md:text-2xl text-teal-100 mb-10 max-w-3xl mx-auto leading-relaxed">
          Профессиональные препараты для лечения и восстановления волос.
          Проверенные средства от ведущих специалистов трихологии.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <a href="#products" class="bg-white text-teal-700 px-8 py-4 rounded-xl font-bold text-lg hover:bg-teal-50 transition shadow-lg">
            Каталог препаратов
          </a>
          <a href="#doctor" class="border-2 border-white/50 text-white px-8 py-4 rounded-xl font-bold text-lg hover:bg-white/10 transition">
            О специалисте
          </a>
        </div>
      </div>
    </section>

    <!-- Stats -->
    <section class="bg-white py-12 border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
          <div>
            <div class="text-3xl font-extrabold text-teal-600">10+</div>
            <div class="text-gray-500 mt-1">Лет опыта</div>
          </div>
          <div>
            <div class="text-3xl font-extrabold text-teal-600">5000+</div>
            <div class="text-gray-500 mt-1">Довольных пациентов</div>
          </div>
          <div>
            <div class="text-3xl font-extrabold text-teal-600">30+</div>
            <div class="text-gray-500 mt-1">Видов препаратов</div>
          </div>
          <div>
            <div class="text-3xl font-extrabold text-teal-600">98%</div>
            <div class="text-gray-500 mt-1">Эффективность</div>
          </div>
        </div>
      </div>
    </section>

    <DoctorSection />

    <!-- Products -->
    <section id="products" class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-14">
          <h2 class="text-3xl md:text-4xl font-extrabold text-gray-800 mb-4">Наши препараты</h2>
          <p class="text-gray-500 text-lg max-w-2xl mx-auto">
            Эффективные средства для лечения и восстановления волос, подобранные нашими специалистами
          </p>
        </div>

        <div v-if="loading" class="text-center py-16">
          <div class="inline-block w-10 h-10 border-4 border-teal-600 border-t-transparent rounded-full animate-spin"></div>
          <p class="mt-4 text-gray-400">Загрузка...</p>
        </div>

        <div v-else-if="products.length === 0" class="text-center py-16">
          <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <p class="text-gray-400 text-lg">Препараты скоро появятся</p>
        </div>

        <div v-else class="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          <ProductCard
            v-for="product in products"
            :key="product.id"
            :product="product"
            @add-to-cart="cartStore.addItem"
          />
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="bg-teal-700 text-white py-16">
      <div class="max-w-4xl mx-auto px-4 text-center">
        <h2 class="text-3xl font-bold mb-4">Начните лечение уже сегодня</h2>
        <p class="text-teal-100 text-lg mb-8">
          Зарегистрируйтесь и получите доступ к полному каталогу препаратов
        </p>
        <router-link
          v-if="!authStore.isLoggedIn"
          to="/register"
          class="inline-block bg-white text-teal-700 px-8 py-4 rounded-xl font-bold text-lg hover:bg-teal-50 transition shadow-lg"
        >
          Зарегистрироваться
        </router-link>
      </div>
    </section>

    <Footer />

    <!-- Order success modal -->
    <div v-if="showSuccess" class="fixed inset-0 z-[60] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showSuccess = false"></div>
      <div class="relative bg-white rounded-2xl p-8 max-w-sm mx-4 text-center shadow-2xl">
        <div class="w-20 h-20 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-5">
          <svg class="w-10 h-10 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-gray-800 mb-2">Заказ оформлен!</h3>
        <p class="text-gray-500 mb-6">Мы свяжемся с вами в ближайшее время для подтверждения</p>
        <button @click="showSuccess = false" class="bg-teal-600 text-white px-8 py-3 rounded-xl hover:bg-teal-700 transition font-semibold">
          Отлично
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import DoctorSection from '../components/DoctorSection.vue'
import ProductCard from '../components/ProductCard.vue'
import CartDrawer from '../components/CartDrawer.vue'
import Footer from '../components/Footer.vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'

const cartStore = useCartStore()
const authStore = useAuthStore()
const products = ref([])
const loading = ref(true)
const showSuccess = ref(false)

onMounted(async () => {
  try {
    const res = await axios.get('/api/products')
    products.value = res.data || []
  } catch (e) {
    console.error('Failed to load products:', e)
  } finally {
    loading.value = false
  }
})

async function placeOrder() {
  try {
    const items = cartStore.items.map(item => ({
      product_id: item.product_id,
      quantity: item.quantity
    }))

    await authStore.api.post('/orders', {
      items,
      phone: authStore.user.phone
    })

    cartStore.clear()
    cartStore.toggle()
    showSuccess.value = true
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при оформлении заказа')
  }
}
</script>
