<template>
  <div class="bg-white">
    <Navbar />
    <CartDrawer @checkout="placeOrder" />

    <!-- Hero -->
    <section class="relative min-h-[100vh] flex items-center bg-brand-950 grain overflow-hidden">
      <!-- Decorative elements -->
      <div class="absolute inset-0">
        <div class="absolute top-1/4 -left-32 w-[500px] h-[500px] bg-brand-800/30 rounded-full blur-[120px] animate-float"></div>
        <div class="absolute bottom-1/4 -right-32 w-[400px] h-[400px] bg-gold-500/10 rounded-full blur-[100px] animate-float-delay"></div>
        <div class="absolute top-20 right-1/4 w-2 h-2 bg-gold-400 rounded-full opacity-60"></div>
        <div class="absolute bottom-32 left-1/3 w-1.5 h-1.5 bg-brand-400 rounded-full opacity-40"></div>
        <div class="absolute top-1/2 right-20 w-1 h-1 bg-gold-300 rounded-full opacity-50"></div>
      </div>
      <!-- Grid pattern -->
      <div class="absolute inset-0 opacity-[0.03]" style="background-image: linear-gradient(rgba(255,255,255,.1) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,.1) 1px, transparent 1px); background-size: 60px 60px;"></div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-32 pb-24 md:pt-40 md:pb-32 w-full">
        <div class="max-w-3xl">
          <div class="section-label text-gold-400 mb-6 animate-fade-up" style="animation-delay: 0.1s">
            <span class="w-8 h-px bg-gold-400"></span>
            Профессиональная трихология
          </div>

          <h1 class="heading-xl text-white mb-6 animate-fade-up text-balance" style="animation-delay: 0.2s">
            Здоровые волосы — это
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-gold-300 to-gold-500">реальность</span>
          </h1>

          <p class="text-lg md:text-xl text-white/60 leading-relaxed mb-10 max-w-xl animate-fade-up" style="animation-delay: 0.3s">
            Профессиональные препараты для лечения и восстановления волос.
            Проверенные средства от ведущих специалистов.
          </p>

          <div class="flex flex-col sm:flex-row gap-4 animate-fade-up" style="animation-delay: 0.4s">
            <a href="#products" class="btn-gold text-center">
              Каталог препаратов
            </a>
            <a href="#doctor" class="btn-outline text-center">
              О специалисте
            </a>
          </div>
        </div>

        <!-- Stats row in hero -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6 mt-20 md:mt-28 pt-10 border-t border-white/10 animate-fade-up" style="animation-delay: 0.6s">
          <div>
            <div class="text-3xl md:text-4xl font-serif text-white">10<span class="text-gold-400">+</span></div>
            <div class="text-sm text-white/40 mt-1 font-medium">Лет опыта</div>
          </div>
          <div>
            <div class="text-3xl md:text-4xl font-serif text-white">5000<span class="text-gold-400">+</span></div>
            <div class="text-sm text-white/40 mt-1 font-medium">Довольных пациентов</div>
          </div>
          <div>
            <div class="text-3xl md:text-4xl font-serif text-white">30<span class="text-gold-400">+</span></div>
            <div class="text-sm text-white/40 mt-1 font-medium">Видов препаратов</div>
          </div>
          <div>
            <div class="text-3xl md:text-4xl font-serif text-white">98<span class="text-gold-400">%</span></div>
            <div class="text-sm text-white/40 mt-1 font-medium">Эффективность</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Trust bar -->
    <section class="bg-brand-900 py-5 border-y border-brand-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-wrap items-center justify-center gap-x-10 gap-y-3 text-sm text-brand-300">
          <div class="flex items-center gap-2">
            <svg class="w-4 h-4 text-gold-400" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M16.403 12.652a3 3 0 000-5.304 3 3 0 00-3.75-3.751 3 3 0 00-5.305 0 3 3 0 00-3.751 3.75 3 3 0 000 5.305 3 3 0 003.75 3.751 3 3 0 005.305 0 3 3 0 003.751-3.75zm-2.546-4.46a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd" /></svg>
            Сертифицированные препараты
          </div>
          <div class="flex items-center gap-2">
            <svg class="w-4 h-4 text-gold-400" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" /></svg>
            Безопасная оплата
          </div>
          <div class="flex items-center gap-2">
            <svg class="w-4 h-4 text-gold-400" fill="currentColor" viewBox="0 0 20 20"><path d="M8 16.25a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75z" /><path fill-rule="evenodd" d="M4 4a3 3 0 013-3h6a3 3 0 013 3v12a3 3 0 01-3 3H7a3 3 0 01-3-3V4zm3-1.5A1.5 1.5 0 005.5 4v12A1.5 1.5 0 007 17.5h6a1.5 1.5 0 001.5-1.5V4A1.5 1.5 0 0013 2.5H7z" clip-rule="evenodd" /></svg>
            Консультация онлайн
          </div>
          <div class="flex items-center gap-2">
            <svg class="w-4 h-4 text-gold-400" fill="currentColor" viewBox="0 0 20 20"><path d="M10 2a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 2zm0 13a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 15zm0-8a3 3 0 100 6 3 3 0 000-6zm-4.243-2.757a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.061zM2 10a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 012 10zm13 0a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 0115 10zM4.343 4.343a.75.75 0 011.061 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.061z" /></svg>
            Доставка по Узбекистану
          </div>
        </div>
      </div>
    </section>

    <DoctorSection />

    <!-- Products -->
    <section id="products" class="py-24 bg-gradient-to-b from-gray-50 to-white relative">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <div class="section-label text-brand-600 justify-center mb-4">
            <span class="w-8 h-px bg-brand-600"></span>
            Каталог
            <span class="w-8 h-px bg-brand-600"></span>
          </div>
          <h2 class="heading-lg text-gray-900 mb-4">Наши препараты</h2>
          <p class="text-gray-500 text-lg max-w-2xl mx-auto leading-relaxed">
            Эффективные средства для лечения и восстановления волос, подобранные нашими специалистами
          </p>
        </div>

        <div v-if="loading" class="text-center py-20">
          <div class="inline-block w-10 h-10 border-[3px] border-brand-200 border-t-brand-700 rounded-full animate-spin"></div>
          <p class="mt-4 text-gray-400 font-medium">Загрузка...</p>
        </div>

        <div v-else-if="products.length === 0" class="text-center py-20">
          <div class="w-20 h-20 bg-gray-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-10 h-10 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <p class="text-gray-400 text-lg font-medium">Препараты скоро появятся</p>
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
    <section class="relative overflow-hidden grain">
      <div class="absolute inset-0 bg-gradient-to-br from-brand-900 via-brand-950 to-brand-950"></div>
      <div class="absolute top-0 right-0 w-96 h-96 bg-gold-500/5 rounded-full blur-[100px]"></div>
      <div class="absolute bottom-0 left-0 w-72 h-72 bg-brand-600/10 rounded-full blur-[80px]"></div>
      <div class="relative max-w-4xl mx-auto px-4 py-24 text-center">
        <div class="section-label text-gold-400 justify-center mb-6">
          <span class="w-8 h-px bg-gold-400/50"></span>
          Начните сегодня
          <span class="w-8 h-px bg-gold-400/50"></span>
        </div>
        <h2 class="heading-lg text-white mb-5 text-balance">Начните лечение уже сегодня</h2>
        <p class="text-white/50 text-lg mb-10 max-w-lg mx-auto leading-relaxed">
          Зарегистрируйтесь и получите доступ к полному каталогу препаратов от специалистов
        </p>
        <router-link
          v-if="!authStore.isLoggedIn"
          to="/register"
          class="btn-gold text-center inline-block"
        >
          Зарегистрироваться
        </router-link>
      </div>
    </section>

    <Footer />

    <!-- Order success modal -->
    <Teleport to="body">
      <div v-if="showSuccess" class="fixed inset-0 z-[60] flex items-center justify-center">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="showSuccess = false"></div>
        <div class="relative bg-white rounded-3xl p-10 max-w-sm mx-4 text-center shadow-2xl animate-scale-in">
          <div class="w-20 h-20 bg-brand-50 rounded-2xl flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h3 class="text-2xl font-serif text-gray-900 mb-2">Заказ оформлен!</h3>
          <p class="text-gray-500 mb-8 leading-relaxed">Мы свяжемся с вами в ближайшее время для подтверждения</p>
          <button @click="showSuccess = false" class="btn-primary w-full">
            Отлично
          </button>
        </div>
      </div>
    </Teleport>
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
import { useAuthStore, api } from '../stores/auth'

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
      quantity: item.quantity,
      unit_type: item.unit_type || 'pack'
    }))

    await api.post('/orders', {
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
