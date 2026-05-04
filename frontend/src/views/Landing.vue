<template>
  <div class="bg-surface">
    <Navbar />
    <CartDrawer ref="cartDrawerRef" @checkout="placeOrder" />
    <OrdersDrawer />

    <!-- Hero -->
    <section class="relative min-h-screen flex items-center bg-gradient-to-br from-brand-700 via-brand-800 to-brand-900 overflow-hidden pt-[96px] md:pt-[112px]">
      <!-- Animated decorative blobs -->
      <div class="absolute top-0 right-0 w-[700px] h-[700px] bg-brand-500/20 rounded-full -translate-y-1/4 translate-x-1/4 animate-pulse-soft"></div>
      <div class="absolute bottom-0 left-0 w-[500px] h-[500px] bg-brand-600/20 rounded-full translate-y-1/4 -translate-x-1/4 animate-pulse-soft" style="animation-delay: 1.5s"></div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 pb-24 md:pt-24 md:pb-32 w-full">
        <div class="grid md:grid-cols-2 gap-12 lg:gap-20 items-center">
          <!-- LEFT: Text -->
          <div>
            <div class="flex items-center gap-3 mb-6 animate-fade-up text-brand-200 text-xs font-semibold tracking-widest uppercase" style="animation-delay: 0.1s; opacity: 0">
              <span class="w-8 h-px bg-brand-400"></span>
              {{ t.hero_label }}
            </div>

            <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight mb-6 animate-fade-up" style="animation-delay: 0.25s; opacity: 0">
              {{ t.hero_title1 }}
              <span class="text-brand-300">{{ t.hero_title2 }}</span>
            </h1>

            <p class="text-lg md:text-xl text-blue-100/80 leading-relaxed mb-10 max-w-xl animate-fade-up" style="animation-delay: 0.4s; opacity: 0">
              {{ t.hero_desc }}
            </p>

            <div class="flex flex-col sm:flex-row gap-4 animate-fade-up" style="animation-delay: 0.55s; opacity: 0">
              <a href="#products" class="inline-flex items-center justify-center gap-2 bg-white text-brand-700 px-8 py-3.5 rounded-full font-semibold hover:bg-brand-50 hover:-translate-y-0.5 hover:shadow-xl hover:shadow-white/20 active:translate-y-0 transition-all duration-300 shadow-lg group">
                {{ t.hero_catalog }}
                <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
              </a>
              <a href="#doctor" class="inline-flex items-center justify-center gap-2 border border-white/30 text-white px-8 py-3.5 rounded-full font-medium hover:bg-white/10 hover:border-white/50 transition-all duration-300">
                {{ t.hero_about }}
              </a>
            </div>

            <!-- Stats row -->
            <div class="grid grid-cols-2 gap-6 mt-14 pt-10 border-t border-white/20 animate-fade-up" style="animation-delay: 0.7s; opacity: 0">
              <div>
                <div class="text-3xl md:text-4xl font-bold text-white">40<span class="text-brand-300">+</span></div>
                <div class="text-sm text-blue-200 mt-1">{{ t.hero_years }}</div>
              </div>
              <div>
                <div class="text-3xl md:text-4xl font-bold text-white">500000<span class="text-brand-300">+</span></div>
                <div class="text-sm text-blue-200 mt-1">{{ t.hero_patients }}</div>
              </div>
            </div>
          </div>

          <!-- RIGHT: Doctor image -->
          <div class="relative hidden md:flex justify-center animate-fade-up" style="animation-delay: 0.35s; opacity: 0">
            <div class="relative w-full max-w-md">
              <!-- Decorative ring -->
              <div class="absolute -inset-4 rounded-3xl bg-white/10 backdrop-blur-sm border border-white/20 z-10"></div>
              <img
                src="/images/patients/3%20landing.jpg"
                alt="Специалист клиники"
                class="relative z-20 rounded-2xl w-full object-cover object-top shadow-2xl shadow-brand-900/40 scale-110"
                style="max-height: 550px;"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Trust bar -->
    <section class="bg-brand-50/70 py-5 border-y border-brand-100/60 backdrop-blur-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-wrap items-center justify-center gap-x-10 gap-y-3 text-sm text-brand-700">
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M16.403 12.652a3 3 0 000-5.304 3 3 0 00-3.75-3.751 3 3 0 00-5.305 0 3 3 0 00-3.751 3.75 3 3 0 000 5.305 3 3 0 003.75 3.751 3 3 0 005.305 0 3 3 0 003.751-3.75zm-2.546-4.46a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd" /></svg>
            {{ t.trust_certified }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" /></svg>
            {{ t.trust_payment }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path d="M8 16.25a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75z" /><path fill-rule="evenodd" d="M4 4a3 3 0 013-3h6a3 3 0 013 3v12a3 3 0 01-3 3H7a3 3 0 01-3-3V4zm3-1.5A1.5 1.5 0 005.5 4v12A1.5 1.5 0 007 17.5h6a1.5 1.5 0 001.5-1.5V4A1.5 1.5 0 0013 2.5H7z" clip-rule="evenodd" /></svg>
            {{ t.trust_consult }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path d="M10 2a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 2zm0 13a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 15zm0-8a3 3 0 100 6 3 3 0 000-6zm-4.243-2.757a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.061zM2 10a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 012 10zm13 0a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 0115 10zM4.343 4.343a.75.75 0 011.061 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.061z" /></svg>
            {{ t.trust_delivery }}
          </div>
        </div>
      </div>
    </section>

    <DoctorSection />

    <!-- Patients Photos Carousel -->
    <section id="patients" class="py-24 bg-gray-50 overflow-hidden">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center mb-14">
        <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-semibold tracking-widest uppercase mb-4">
          <span class="w-8 h-px bg-brand-400"></span>
          {{ t.patients_label }}
          <span class="w-8 h-px bg-brand-400"></span>
        </div>
        <h2 class="text-3xl md:text-4xl font-bold text-stone-900 mb-4">{{ t.patients_title }}</h2>
        <p class="text-stone-500 text-lg max-w-xl mx-auto">{{ t.patients_subtitle }}</p>
      </div>

      <!-- Auto-loop gallery -->
      <div class="relative overflow-hidden">
        <!-- Left gradient fade -->
        <div class="absolute left-0 top-0 bottom-0 w-16 lg:w-28 bg-gradient-to-r from-gray-50 to-transparent z-10 pointer-events-none"></div>
        <!-- Right gradient fade -->
        <div class="absolute right-0 top-0 bottom-0 w-16 lg:w-28 bg-gradient-to-l from-gray-50 to-transparent z-10 pointer-events-none"></div>

        <div
          ref="patientsScroller"
          class="overflow-x-auto px-3 sm:px-6 pb-4 cursor-grab active:cursor-grabbing select-none [scrollbar-width:none] [&::-webkit-scrollbar]:hidden"
          @wheel.passive="pauseAutoScroll"
          @scroll.passive="handleInfiniteScroll"
          @mousedown="onDragStart"
          @mousemove="onDragMove"
          @mouseup="onDragEnd"
          @mouseleave="onDragEnd"
          @touchstart="onDragStart"
          @touchmove="onDragMove"
          @touchend="onDragEnd"
        >
          <div class="flex gap-5 w-max">
            <template v-for="rep in 2" :key="rep">
              <div
                v-for="(photo, idx) in patientPhotos"
                :key="`${rep}-${photo}`"
                class="shrink-0 w-64 h-80 rounded-2xl overflow-hidden shadow-lg group cursor-pointer"
              >
                <img
                  :src="`/images/patients/${photo}`"
                  :alt="`Фото с пациентом ${idx + 1}`"
                  class="w-full h-full object-cover object-top group-hover:scale-105 transition-transform duration-500"
                  draggable="false"
                />
              </div>
            </template>
          </div>
        </div>
      </div>
    </section>

    <!-- Products -->
    <section id="products" class="py-24 bg-gradient-to-b from-white via-surface-warm to-surface relative">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16 reveal" ref="productsHeader">
          <div class="section-label justify-center mb-4">
            <span class="w-8 h-px bg-brand-400"></span>
            {{ t.products_label }}
            <span class="w-8 h-px bg-brand-400"></span>
          </div>
          <h2 class="heading-lg mb-4">{{ t.products_title }}</h2>
          <p class="text-stone-500 text-lg max-w-2xl mx-auto leading-relaxed">{{ t.products_desc }}</p>
        </div>

        <div v-if="loading" class="text-center py-20">
          <div class="inline-block w-10 h-10 border-[3px] border-brand-200 border-t-brand-600 rounded-full animate-spin"></div>
          <p class="mt-4 text-stone-400 font-medium">{{ t.products_loading }}</p>
        </div>

        <div v-else-if="products.length === 0" class="text-center py-20">
          <div class="w-20 h-20 bg-stone-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-10 h-10 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <p class="text-stone-400 text-lg font-medium">{{ t.products_empty }}</p>
        </div>

        <div v-else class="grid grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3 sm:gap-6 stagger-children">
          <ProductCard
            v-for="(product, index) in products"
            :key="product.id"
            :product="product"
            :style="{ '--stagger': index }"
            class="reveal"
            :ref="el => { if (el) productRefs[index] = el.$el || el }"
            @add-to-cart="cartStore.addItem"
          />
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="relative overflow-hidden reveal" ref="ctaSection">
      <div class="absolute inset-0 bg-gradient-to-br from-brand-700 via-brand-800 to-brand-900"></div>
      <div class="absolute top-0 right-0 w-96 h-96 bg-brand-500/10 rounded-full blur-[100px] animate-pulse-soft"></div>
      <div class="absolute bottom-0 left-0 w-72 h-72 bg-brand-400/10 rounded-full blur-[80px] animate-pulse-soft" style="animation-delay: 2s"></div>
      <div class="relative max-w-4xl mx-auto px-4 py-24 text-center">
        <div class="section-label !text-brand-200 justify-center mb-6">
          <span class="w-8 h-px bg-brand-400/50"></span>
          {{ t.cta_label }}
          <span class="w-8 h-px bg-brand-400/50"></span>
        </div>
        <h2 class="heading-lg !text-white mb-5 text-balance">{{ t.cta_title }}</h2>
        <p class="text-brand-200/70 text-lg mb-10 max-w-lg mx-auto leading-relaxed">{{ t.cta_desc }}</p>
        <router-link
          v-if="!authStore.isLoggedIn"
          to="/register"
          class="inline-block bg-white text-brand-800 px-8 py-4 rounded-full font-semibold hover:bg-brand-50 hover:shadow-2xl hover:shadow-white/10 hover:-translate-y-1 active:scale-[0.97] active:translate-y-0 transition-all duration-300 shadow-lg group"
        >
          <span class="inline-flex items-center gap-2">
            {{ t.cta_register }}
            <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
          </span>
        </router-link>
      </div>
    </section>

    <Footer />

    <!-- Order success modal -->
    <Teleport to="body">
      <div v-if="showSuccess" class="fixed inset-0 z-[60] flex items-center justify-center">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-md animate-fade-in" @click="showSuccess = false"></div>
        <div class="relative bg-white rounded-3xl p-10 max-w-sm mx-4 text-center shadow-2xl animate-scale-in">
          <div class="w-20 h-20 bg-brand-50 rounded-2xl flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h3 class="text-2xl font-serif text-stone-900 mb-2">Заказ оформлен!</h3>
          <p class="text-stone-500 mb-4 leading-relaxed">Мы свяжемся с вами в ближайшее время для подтверждения</p>
          <div v-if="lastOrderCode" class="bg-blue-50 border border-blue-200 rounded-2xl p-4 mb-6">
            <p class="text-xs text-blue-500 font-medium uppercase tracking-wider mb-1">Ваш код заказа</p>
            <p class="text-4xl font-bold text-blue-700 tracking-widest">{{ lastOrderCode }}</p>
            <p class="text-xs text-blue-400 mt-1">Назовите этот код на пункте выдачи</p>
          </div>
          <button @click="showSuccess = false" class="btn-primary w-full">
            Отлично
          </button>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import DoctorSection from '../components/DoctorSection.vue'
import ProductCard from '../components/ProductCard.vue'
import CartDrawer from '../components/CartDrawer.vue'
import OrdersDrawer from '../components/OrdersDrawer.vue'
import Footer from '../components/Footer.vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore, api } from '../stores/auth'
import { useLangStore } from '../stores/lang'

const cartStore = useCartStore()
const authStore = useAuthStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)
const cartDrawerRef = ref(null)
const products = ref([])
const loading = ref(true)
const showSuccess = ref(false)
const lastOrderCode = ref('')

const patientPhotos = [
  'photo_2026-04-21_23-34-03.jpg',
  'photo_2026-04-21_23-34-31.jpg',
  'photo_2026-04-21_23-34-39.jpg',
  'photo_2026-04-21_23-34-44.jpg',
  'photo_2026-04-21_23-34-52.jpg',
  'photo_2026-04-21_23-34-57.jpg',
  'photo_2026-04-21_23-35-02.jpg',
  'photo_2026-04-21_23-35-10.jpg',
]

const AUTO_SCROLL_PAUSE_MS = 3000
const AUTO_SCROLL_SPEED = 0.55

// Scroll reveal refs
const productsHeader = ref(null)
const ctaSection = ref(null)
const productRefs = ref({})
const patientsScroller = ref(null)

const dragState = {
  active: false,
  startX: 0,
  scrollLeft: 0,
}

let observer = null
let autoScrollRaf = null
let autoScrollResumeTimer = null
const autoScrollPaused = ref(false)

function startAutoScroll() {
  const step = () => {
    const scroller = patientsScroller.value
    if (scroller && !autoScrollPaused.value) {
      const loopWidth = scroller.scrollWidth / 2
      if (loopWidth > 0) {
        scroller.scrollLeft += AUTO_SCROLL_SPEED
        if (scroller.scrollLeft >= loopWidth) {
          scroller.scrollLeft -= loopWidth
        }
      }
    }
    autoScrollRaf = requestAnimationFrame(step)
  }

  if (autoScrollRaf) cancelAnimationFrame(autoScrollRaf)
  autoScrollRaf = requestAnimationFrame(step)
}

function pauseAutoScroll() {
  autoScrollPaused.value = true
  if (autoScrollResumeTimer) clearTimeout(autoScrollResumeTimer)
  autoScrollResumeTimer = setTimeout(() => {
    autoScrollPaused.value = false
  }, AUTO_SCROLL_PAUSE_MS)
}

function handleInfiniteScroll() {
  const scroller = patientsScroller.value
  if (!scroller) return

  const loopWidth = scroller.scrollWidth / 2
  if (loopWidth <= 0) return

  if (scroller.scrollLeft >= loopWidth) {
    scroller.scrollLeft -= loopWidth
  }
}

function setupObserver() {
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible')
          observer.unobserve(entry.target)
        }
      })
    },
    { threshold: 0.1, rootMargin: '0px 0px -40px 0px' }
  )

  // Observe direct refs
  if (productsHeader.value) observer.observe(productsHeader.value)
  if (ctaSection.value) observer.observe(ctaSection.value)

  // Observe all .reveal elements
  document.querySelectorAll('.reveal, .reveal-left, .reveal-right, .reveal-scale').forEach((el) => {
    observer.observe(el)
  })
}

onMounted(async () => {
  try {
    const res = await axios.get('/api/products')
    products.value = res.data || []
  } catch (e) {
    console.error('Failed to load products:', e)
  } finally {
    loading.value = false
  }

  await nextTick()
  setupObserver()
  startAutoScroll()
})

onUnmounted(() => {
  if (observer) observer.disconnect()
  if (autoScrollRaf) cancelAnimationFrame(autoScrollRaf)
  if (autoScrollResumeTimer) clearTimeout(autoScrollResumeTimer)
})

function getClientX(event) {
  if ('touches' in event && event.touches.length > 0) return event.touches[0].clientX
  if ('changedTouches' in event && event.changedTouches.length > 0) return event.changedTouches[0].clientX
  return event.clientX
}

function onDragStart(event) {
  if (!patientsScroller.value) return
  pauseAutoScroll()
  dragState.active = true
  dragState.startX = getClientX(event)
  dragState.scrollLeft = patientsScroller.value.scrollLeft
}

function onDragMove(event) {
  if (!dragState.active || !patientsScroller.value) return
  const currentX = getClientX(event)
  const walk = currentX - dragState.startX
  patientsScroller.value.scrollLeft = dragState.scrollLeft - walk
}

function onDragEnd() {
  dragState.active = false
  pauseAutoScroll()
}

async function placeOrder() {
  try {
    const items = cartStore.items.map(item => ({
      product_id: item.product_id,
      quantity: item.quantity,
      unit_type: item.unit_type || 'pack'
    }))

    const res = await api.post('/orders', {
      items,
      phone: authStore.user.phone
    })

    lastOrderCode.value = res.data.order_code || ''
    showSuccess.value = true
    cartDrawerRef.value?.showOrders()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при оформлении заказа')
  }
}
</script>
