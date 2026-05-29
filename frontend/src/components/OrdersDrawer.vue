<template>
  <Teleport to="body">
    <div v-if="ordersStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="ordersStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white/95 backdrop-blur-2xl shadow-2xl flex flex-col animate-slide-right">

        <!-- Header -->
        <div class="flex items-center justify-between px-6 py-5 border-b border-stone-100">
          <div>
            <h2 class="text-xl font-serif text-stone-900">{{ t.orders_title }}</h2>
            <p class="text-xs text-stone-400 mt-0.5">{{ activeOrders.length }} {{ t.orders_count }}</p>
          </div>
          <button @click="ordersStore.toggle()" class="p-2 hover:bg-stone-100 rounded-xl hover:rotate-90 transition-all duration-300">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="flex-1 flex items-center justify-center">
          <div class="w-8 h-8 border-2 border-brand-600 border-t-transparent rounded-full animate-spin"></div>
        </div>

        <!-- Empty state -->
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

          <!-- Active orders (not cancelled) -->
          <div
            v-for="order in activeOrders"
            :key="order.id"
            class="rounded-2xl border border-stone-100 bg-stone-50 overflow-hidden"
          >
            <!-- Order header with code + status -->
            <div class="flex items-center justify-between px-4 py-3 bg-white border-b border-stone-100">
              <div>
                <p class="text-[10px] text-stone-400 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                <p class="text-2xl font-bold text-brand-700 tracking-widest">{{ order.order_code }}</p>
              </div>
              <span class="px-3 py-1.5 rounded-xl text-xs font-semibold" :class="statusClass(order.status)">
                {{ statusLabel(order.status) }}
              </span>
            </div>

            <!-- Items -->
            <div class="px-4 py-3 space-y-2.5">
              <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2">
                <div class="flex items-center gap-2 min-w-0">
                  <div class="w-8 h-8 rounded-lg bg-stone-200 overflow-hidden flex-shrink-0 flex items-center justify-center">
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

            <!-- Total -->
            <div class="px-4 py-2.5 border-t border-stone-100 flex justify-between items-center">
              <span class="text-xs text-stone-400">{{ t.cart_total }}</span>
              <span class="font-bold text-brand-700">{{ formatPrice(orderTotal(order)) }} {{ t.currency }}</span>
            </div>
          </div>

          <!-- Cancelled orders section -->
          <template v-if="cancelledOrders.length > 0">
            <div class="pt-2">
              <div class="flex items-center gap-2 mb-3">
                <div class="flex-1 h-px bg-red-200"></div>
                <span class="text-xs font-semibold text-red-400 uppercase tracking-wider">{{ t.status_cancelled }}</span>
                <div class="flex-1 h-px bg-red-200"></div>
              </div>

              <div
                v-for="order in cancelledOrders"
                :key="order.id"
                class="rounded-2xl border border-red-200 bg-red-50/60 overflow-hidden mb-4"
              >
                <div class="flex items-center justify-between px-4 py-3 bg-red-50 border-b border-red-100">
                  <div>
                    <p class="text-[10px] text-red-300 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                    <p class="text-2xl font-bold text-red-400 tracking-widest">{{ order.order_code }}</p>
                  </div>
                  <span class="px-3 py-1.5 rounded-xl text-xs font-semibold bg-red-100 text-red-600">
                    {{ t.status_cancelled }}
                  </span>
                </div>

                <div class="px-4 py-3 space-y-2">
                  <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2 opacity-60">
                    <div class="flex items-center gap-2 min-w-0">
                      <div class="w-8 h-8 rounded-lg bg-red-100 overflow-hidden flex-shrink-0 flex items-center justify-center">
                        <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover grayscale" />
                        <svg v-else class="w-4 h-4 text-red-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                        </svg>
                      </div>
                      <div class="min-w-0">
                        <p class="font-medium text-stone-700 text-sm truncate line-through">{{ item.product?.name }}</p>
                        <p class="text-stone-400 text-xs">{{ item.quantity }} {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</p>
                      </div>
                    </div>
                    <p class="font-semibold text-stone-500 text-sm flex-shrink-0 line-through">{{ formatPrice(item.price) }} {{ t.currency }}</p>
                  </div>
                </div>
              </div>
            </div>
          </template>

        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useOrdersStore } from '../stores/orders'
import { useLangStore } from '../stores/lang'
import { api } from '../stores/auth'

const ordersStore = useOrdersStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)

const orders = ref([])
const loading = ref(false)

const activeOrders = computed(() => orders.value.filter(o => o.status !== 'cancelled'))
const cancelledOrders = computed(() => orders.value.filter(o => o.status === 'cancelled'))

watch(() => ordersStore.isOpen, async (val) => {
  if (val) {
    loading.value = true
    try {
      const res = await api.get('/orders/my')
      orders.value = res.data || []
    } catch (e) {
      orders.value = []
    } finally {
      loading.value = false
    }
  }
})

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

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
</script>
