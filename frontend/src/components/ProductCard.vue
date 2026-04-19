<template>
  <div class="bg-white rounded-xl shadow-sm overflow-hidden hover:shadow-md transition-shadow border border-gray-100 group">
    <div class="aspect-square bg-gray-50 flex items-center justify-center overflow-hidden">
      <img
        v-if="product.image_path"
        :src="product.image_path"
        :alt="product.name"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
      />
      <div v-else class="text-center">
        <svg class="w-16 h-16 text-gray-200 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
        </svg>
        <p class="text-xs text-gray-300 mt-2">Нет фото</p>
      </div>
    </div>

    <div class="p-4">
      <h3 class="font-semibold text-gray-800 text-lg mb-1 truncate">{{ product.name }}</h3>
      <p v-if="product.description" class="text-gray-400 text-sm mb-3 line-clamp-2 leading-relaxed">{{ product.description }}</p>

      <div class="space-y-1.5 mb-4">
        <div class="flex justify-between text-sm">
          <span class="text-gray-400">1 таблетка:</span>
          <span class="font-medium text-gray-600">{{ formatPrice(product.price_per_pill) }} сўм</span>
        </div>
        <div class="flex justify-between items-baseline">
          <span class="text-gray-400 text-sm">Упаковка ({{ product.quantity_per_pack }} шт):</span>
          <span class="font-bold text-teal-600 text-lg">{{ formatPrice(product.price_per_pack) }} сўм</span>
        </div>
      </div>

      <button
        @click="$emit('add-to-cart', product)"
        class="w-full bg-teal-600 text-white py-2.5 rounded-lg hover:bg-teal-700 active:bg-teal-800 transition font-medium flex items-center justify-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 100 4 2 2 0 000-4z" />
        </svg>
        В корзину
      </button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  product: { type: Object, required: true }
})

defineEmits(['add-to-cart'])

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}
</script>
