<template>
  <Teleport to="body">
    <div v-if="cartStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="cartStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white/95 backdrop-blur-2xl shadow-2xl flex flex-col animate-slide-right">

        <!-- Header -->
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100">
          <h2 class="text-xl font-serif text-stone-900">
            {{ activeTab === 'cart' ? t.cart_title : t.orders_title }}
          </h2>
          <button @click="cartStore.toggle()" class="p-2 hover:bg-stone-100 rounded-xl hover:rotate-90 transition-all duration-300">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Tabs -->
        <div class="flex border-b border-stone-100 bg-stone-50/60">
          <button
            @click="switchTab('cart')"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-200"
            :class="activeTab === 'cart' ? 'text-brand-700 border-b-2 border-brand-600 bg-white' : 'text-stone-400 hover:text-stone-600'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z" />
            </svg>
            {{ t.cart_title }}
            <span v-if="cartStore.totalItems > 0" class="bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
              {{ cartStore.totalItems }}
            </span>
          </button>
          <button
            v-if="authStore.isLoggedIn"
            @click="switchTab('orders')"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-200"
            :class="activeTab === 'orders' ? 'text-brand-700 border-b-2 border-brand-600 bg-white' : 'text-stone-400 hover:text-stone-600'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            {{ t.orders_title }}
          </button>
        </div>

        <!-- ===== TAB: CART ===== -->
        <template v-if="activeTab === 'cart'">
          <!-- Empty state -->
          <div v-if="cartStore.items.length === 0" class="flex-1 flex flex-col items-center justify-center px-8">
            <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
              <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
              </svg>
            </div>
            <p class="text-stone-800 text-lg font-medium mb-1">{{ t.cart_empty }}</p>
            <p class="text-stone-400 text-sm">{{ t.cart_empty_sub }}</p>
          </div>

          <!-- Items -->
          <div v-else class="flex-1 overflow-y-auto px-6 py-5 space-y-4">
            <div
              v-for="item in cartStore.items"
              :key="item.product_id"
              class="bg-stone-50 rounded-2xl p-4 border border-stone-100 hover:border-brand-100 transition-colors duration-300"
            >
              <div class="flex gap-3">
                <div class="w-14 h-14 bg-stone-200 rounded-xl overflow-hidden flex-shrink-0 flex items-center justify-center">
                  <img v-if="item.image_path" :src="item.image_path" class="w-full h-full object-cover" />
                  <svg v-else class="w-6 h-6 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                  </svg>
                </div>
                <div class="flex-1 min-w-0">
                  <div class="flex items-start justify-between gap-2">
                    <h4 class="font-semibold text-stone-900 text-sm truncate">{{ item.name }}</h4>
                    <button @click="cartStore.removeItem(item.product_id)" class="text-stone-300 hover:text-red-400 transition-colors p-0.5 flex-shrink-0">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </button>
                  </div>
                  <p class="text-brand-600 font-semibold text-sm mt-0.5">
                    {{ formatPrice(item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack) }} {{ t.currency }}
                    <span class="text-stone-400 font-normal text-xs">/ {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</span>
                  </p>
                </div>
              </div>

              <div class="flex items-center justify-between mt-3 pt-3 border-t border-stone-200/60">
                <div class="flex items-center gap-1">
                  <button
                    @click="cartStore.updateUnitType(item.product_id, 'piece')"
                    :class="item.unit_type === 'piece' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                    class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                  >{{ t.unit_piece }}</button>
                  <button
                    @click="cartStore.updateUnitType(item.product_id, 'pack')"
                    :class="(item.unit_type || 'pack') === 'pack' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                    class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                  >{{ t.unit_pack }}</button>
                </div>
                <div class="flex items-center gap-2">
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity - 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600">−</button>
                  <span class="text-sm font-bold w-6 text-center text-stone-900">{{ item.quantity }}</span>
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity + 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600">+</button>
                </div>
              </div>

              <div class="text-right mt-2">
                <span class="text-xs font-semibold text-stone-500">{{ t.cart_item_total }} <span class="text-brand-700">{{ formatPrice(itemTotal(item)) }} {{ t.currency }}</span></span>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div v-if="cartStore.items.length > 0" class="border-t border-stone-100 px-6 py-5 bg-stone-50/50">
            <div class="flex justify-between items-baseline mb-4">
              <span class="text-sm font-medium text-stone-500">{{ t.cart_total }}</span>
              <span class="text-2xl font-serif text-brand-700">{{ formatPrice(cartStore.totalPrice) }} <span class="text-sm font-sans">{{ t.currency }}</span></span>
            </div>
            <button @click="checkout" class="w-full btn-primary py-4 text-base rounded-xl">
              {{ t.cart_checkout }}
            </button>
          </div>
        </template>

        <!-- ===== TAB: MY ORDERS ===== -->
        <template v-else>
          <!-- Loading -->
          <div v-if="ordersLoading" class="flex-1 flex items-center justify-center">
            <div class="w-8 h-8 border-2 border-brand-600 border-t-transparent rounded-full animate-spin"></div>
          </div>

          <!-- Empty -->
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
            <!-- Active orders -->
            <div
              v-for="order in activeOrders"
              :key="order.id"
              class="rounded-2xl border border-stone-100 bg-stone-50 overflow-hidden"
            >
              <!-- Code + status -->
              <div class="flex items-center justify-between px-4 py-3 bg-white border-b border-stone-100">
                <div>
                  <p class="text-[10px] text-stone-400 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                  <p class="text-2xl font-bold text-brand-700 tracking-widest">{{ order.order_code }}</p>
                </div>
                <span class="px-3 py-1.5 rounded-xl text-xs font-semibold" :class="statusClass(order.status)">
                  {{ statusLabel(order.status) }}
                </span>
              </div>

              <!-- Items inside order -->
              <div class="px-4 py-3 space-y-2.5">
                <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2">
                  <div class="flex items-center gap-2 min-w-0">
                    <div class="w-9 h-9 rounded-lg bg-stone-200 overflow-hidden flex-shrink-0 flex items-center justify-center">
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

              <!-- Order total -->
              <div class="px-4 py-2.5 border-t border-stone-100 flex justify-between items-center">
                <span class="text-xs text-stone-400">{{ t.cart_total }}</span>
                <span class="font-bold text-brand-700">{{ formatPrice(orderTotal(order)) }} {{ t.currency }}</span>
              </div>
            </div>

            <!-- Cancelled orders -->
            <template v-if="cancelledOrders.length > 0">
              <div class="flex items-center gap-2 pt-2">
                <div class="flex-1 h-px bg-red-200"></div>
                <span class="text-xs font-semibold text-red-400 uppercase tracking-wider">{{ t.status_cancelled }}</span>
                <div class="flex-1 h-px bg-red-200"></div>
              </div>
              <div
                v-for="order in cancelledOrders"
                :key="order.id"
                class="rounded-2xl border border-red-200 bg-red-50/60 overflow-hidden"
              >
                <div class="flex items-center justify-between px-4 py-3 bg-red-50 border-b border-red-100">
                  <div>
                    <p class="text-[10px] text-red-300 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                    <p class="text-2xl font-bold text-red-400 tracking-widest">{{ order.order_code }}</p>
                  </div>
                  <span class="px-3 py-1.5 rounded-xl text-xs font-semibold bg-red-100 text-red-600">{{ t.status_cancelled }}</span>
                </div>
                <div class="px-4 py-3 space-y-2">
                  <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2 opacity-60">
                    <div class="flex items-center gap-2 min-w-0">
                      <div class="w-9 h-9 rounded-lg bg-red-100 overflow-hidden flex-shrink-0 flex items-center justify-center">
                        <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover grayscale" />
                        <svg v-else class="w-4 h-4 text-red-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                        </svg>
                      </div>
                      <div class="min-w-0">
                        <p class="font-medium text-stone-600 text-sm truncate line-through">{{ item.product?.name }}</p>
                        <p class="text-stone-400 text-xs">{{ item.quantity }} {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</p>
                      </div>
                    </div>
                    <p class="font-semibold text-stone-400 text-sm flex-shrink-0 line-through">{{ formatPrice(item.price) }} {{ t.currency }}</p>
                  </div>
                </div>
              </div>
            </template>
          </div>
        </template>

      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useLangStore } from '../stores/lang'
import { useRouter } from 'vue-router'
import { api } from '../stores/auth'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()
const langStore = useLangStore()
const t = computed(() => langStore.t)

const emit = defineEmits(['checkout'])

// Tabs
const activeTab = ref('cart')

function switchTab(tab) {
  activeTab.value = tab
  if (tab === 'orders') loadOrders()
}

// Reset to cart tab when drawer closes
watch(() => cartStore.isOpen, (val) => {
  if (!val) activeTab.value = 'cart'
})

// Orders
const orders = ref([])
const ordersLoading = ref(false)

async function loadOrders() {
  ordersLoading.value = true
  try {
    const res = await api.get('/orders')
    orders.value = res.data || []
  } catch {
    orders.value = []
  } finally {
    ordersLoading.value = false
  }
}

const activeOrders = computed(() => orders.value.filter(o => o.status !== 'cancelled'))
const cancelledOrders = computed(() => orders.value.filter(o => o.status === 'cancelled'))

function orderTotal(order) {
  return (order.items || []).reduce((sum, i) => sum + i.price, 0)
}

function statusLabel(status) {
  const map = {
    pending: t.value.status_pending,
    confirmed: t.value.status_confirmed,
    shipped: t.value.status_confirmed,
    delivered: t.value.status_delivered,
    cancelled: t.value.status_cancelled,
  }
  return map[status] || status
}

function statusClass(status) {
  const map = {
    pending: 'bg-yellow-100 text-yellow-700',
    confirmed: 'bg-blue-100 text-blue-700',
    shipped: 'bg-blue-100 text-blue-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-600',
  }
  return map[status] || 'bg-stone-100 text-stone-600'
}

// Cart helpers
function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

function itemTotal(item) {
  const unitPrice = item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack
  return unitPrice * item.quantity
}

function checkout() {
  if (!authStore.isLoggedIn) {
    cartStore.toggle()
    router.push('/login')
    return
  }
  emit('checkout')
}

function showOrders() {
  activeTab.value = 'orders'
  loadOrders()
}

defineExpose({ showOrders })
</script>
