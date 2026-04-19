<template>
  <Teleport to="body">
    <div v-if="cartStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="cartStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white shadow-2xl flex flex-col">
        <!-- Header -->
        <div class="flex items-center justify-between px-5 py-4 border-b">
          <h2 class="text-xl font-bold text-gray-800">
            Корзина
            <span v-if="cartStore.totalItems > 0" class="text-gray-400 font-normal text-base">({{ cartStore.totalItems }})</span>
          </h2>
          <button @click="cartStore.toggle()" class="p-2 hover:bg-gray-100 rounded-lg transition">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Empty state -->
        <div v-if="cartStore.items.length === 0" class="flex-1 flex flex-col items-center justify-center px-8">
          <svg class="w-20 h-20 text-gray-200 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 100 4 2 2 0 000-4z" />
          </svg>
          <p class="text-gray-400 text-lg font-medium">Корзина пуста</p>
          <p class="text-gray-300 text-sm mt-1">Добавьте препараты из каталога</p>
        </div>

        <!-- Items -->
        <div v-else class="flex-1 overflow-y-auto px-5 py-4 space-y-3">
          <div
            v-for="item in cartStore.items"
            :key="item.product_id"
            class="flex gap-3 bg-gray-50 rounded-xl p-3 border border-gray-100"
          >
            <div class="w-16 h-16 bg-gray-200 rounded-lg overflow-hidden flex-shrink-0 flex items-center justify-center">
              <img v-if="item.image_path" :src="item.image_path" class="w-full h-full object-cover" />
              <svg v-else class="w-8 h-8 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <h4 class="font-medium text-gray-800 text-sm truncate">{{ item.name }}</h4>
              <p class="text-teal-600 font-semibold text-sm">{{ formatPrice(item.price_per_pack) }} сўм</p>
              <div class="flex items-center gap-2 mt-1.5">
                <button
                  @click="cartStore.updateQuantity(item.product_id, item.quantity - 1)"
                  class="w-7 h-7 rounded-md bg-gray-200 hover:bg-gray-300 flex items-center justify-center text-sm font-medium transition"
                >−</button>
                <span class="text-sm font-semibold w-6 text-center">{{ item.quantity }}</span>
                <button
                  @click="cartStore.updateQuantity(item.product_id, item.quantity + 1)"
                  class="w-7 h-7 rounded-md bg-gray-200 hover:bg-gray-300 flex items-center justify-center text-sm font-medium transition"
                >+</button>
                <button
                  @click="cartStore.removeItem(item.product_id)"
                  class="ml-auto text-red-400 hover:text-red-600 transition p-1"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div v-if="cartStore.items.length > 0" class="border-t px-5 py-4 space-y-3 bg-gray-50">
          <div class="flex justify-between text-lg font-bold">
            <span class="text-gray-700">Итого:</span>
            <span class="text-teal-600">{{ formatPrice(cartStore.totalPrice) }} сўм</span>
          </div>
          <button
            @click="checkout"
            class="w-full bg-teal-600 text-white py-3.5 rounded-xl hover:bg-teal-700 transition font-semibold text-lg shadow-sm"
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

function checkout() {
  if (!authStore.isLoggedIn) {
    cartStore.toggle()
    router.push('/login')
    return
  }
  emit('checkout')
}
</script>
