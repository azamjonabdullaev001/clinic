<template>
  <div class="group bg-white rounded-2xl overflow-hidden card-hover border border-gray-100/80">
    <!-- Image -->
    <div class="aspect-[4/3] bg-gradient-to-br from-gray-50 to-gray-100 relative overflow-hidden">
      <img
        v-if="product.image_path"
        :src="product.image_path"
        :alt="product.name"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
      />
      <div v-else class="w-full h-full flex flex-col items-center justify-center">
        <div class="w-14 h-14 rounded-2xl bg-gray-200/60 flex items-center justify-center mb-2">
          <svg class="w-7 h-7 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
          </svg>
        </div>
        <span class="text-xs text-gray-300 font-medium">Нет фото</span>
      </div>
      <!-- Overlay gradient -->
      <div class="absolute inset-0 bg-gradient-to-t from-black/5 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
    </div>

    <!-- Content -->
    <div class="p-5">
      <h3 class="font-semibold text-gray-900 text-base mb-1 truncate">{{ product.name }}</h3>
      <p v-if="product.description" class="text-gray-400 text-sm mb-4 line-clamp-2 leading-relaxed">{{ product.description }}</p>

      <!-- Prices -->
      <div class="space-y-2 mb-5">
        <div class="flex items-center justify-between">
          <span class="text-xs font-medium text-gray-400 uppercase tracking-wide">1 таблетка</span>
          <span class="text-sm font-semibold text-gray-600">{{ formatPrice(product.price_per_pill) }} сўм</span>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-xs font-medium text-gray-400 uppercase tracking-wide">Упаковка <span class="normal-case">({{ product.quantity_per_pack }} шт)</span></span>
          <span class="text-lg font-bold text-brand-800">{{ formatPrice(product.price_per_pack) }} сўм</span>
        </div>
      </div>

      <!-- Add to cart -->
      <button
        @click="$emit('add-to-cart', product)"
        class="w-full bg-brand-800 text-white py-3 rounded-xl hover:bg-brand-900 active:scale-[0.98]
               transition-all duration-200 font-medium flex items-center justify-center gap-2
               shadow-lg shadow-brand-800/10 hover:shadow-brand-800/20"
      >
        <svg class="w-4.5 h-4.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
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
