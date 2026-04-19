<template>
  <Teleport to="body">
    <div v-if="cartStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="cartStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white/95 backdrop-blur-2xl shadow-2xl flex flex-col animate-slide-right">
        <!-- Header -->
        <div class="flex items-center justify-between px-6 py-5 border-b border-stone-100">
          <div>
            <h2 class="text-xl font-serif text-stone-900">Корзина</h2>
            <p v-if="cartStore.totalItems > 0" class="text-xs text-stone-400 mt-0.5">{{ cartStore.totalItems }} товаров</p>
          </div>
          <button @click="cartStore.toggle()" class="p-2 hover:bg-stone-100 rounded-xl hover:rotate-90 transition-all duration-300">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Empty state -->
        <div v-if="cartStore.items.length === 0" class="flex-1 flex flex-col items-center justify-center px-8">
          <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
            <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
            </svg>
          </div>
          <p class="text-stone-800 text-lg font-medium mb-1">Корзина пуста</p>
          <p class="text-stone-400 text-sm">Добавьте препараты из каталога</p>
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
                  <button
                    @click="cartStore.removeItem(item.product_id)"
                    class="text-stone-300 hover:text-red-400 transition-colors p-0.5 flex-shrink-0"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
                <p class="text-brand-600 font-semibold text-sm mt-0.5">
                  {{ formatPrice(item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack) }} сўм
                  <span class="text-stone-400 font-normal text-xs">/ {{ item.unit_type === 'piece' ? 'шт' : 'упак' }}</span>
                </p>
              </div>
            </div>

            <!-- Unit type + quantity -->
            <div class="flex items-center justify-between mt-3 pt-3 border-t border-stone-200/60">
              <div class="flex items-center gap-1">
                <button
                  @click="cartStore.updateUnitType(item.product_id, 'piece')"
                  :class="item.unit_type === 'piece' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                  class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                >шт</button>
                <button
                  @click="cartStore.updateUnitType(item.product_id, 'pack')"
                  :class="(item.unit_type || 'pack') === 'pack' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                  class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                >коробка</button>
              </div>

              <div class="flex items-center gap-2">
                <button
                  @click="cartStore.updateQuantity(item.product_id, item.quantity - 1)"
                  class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600"
                >−</button>
                <span class="text-sm font-bold w-6 text-center text-stone-900">{{ item.quantity }}</span>
                <button
                  @click="cartStore.updateQuantity(item.product_id, item.quantity + 1)"
                  class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600"
                >+</button>
              </div>
            </div>

            <div class="text-right mt-2">
              <span class="text-xs font-semibold text-stone-500">Итого: <span class="text-brand-700">{{ formatPrice(itemTotal(item)) }} сўм</span></span>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div v-if="cartStore.items.length > 0" class="border-t border-stone-100 px-6 py-5 bg-stone-50/50">
          <div class="flex justify-between items-baseline mb-4">
            <span class="text-sm font-medium text-stone-500">Итого к оплате</span>
            <span class="text-2xl font-serif text-brand-700">{{ formatPrice(cartStore.totalPrice) }} <span class="text-sm font-sans">сўм</span></span>
          </div>
          <button
            @click="checkout"
            class="w-full btn-primary py-4 text-base rounded-xl"
          >
            Оформить заказ
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()

const emit = defineEmits(['checkout'])

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
</script>
