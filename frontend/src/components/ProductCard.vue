<template>
  <div class="group bg-white rounded-2xl overflow-hidden card-hover border border-stone-100/80 hover:border-brand-200/50 flex flex-col">
    <!-- Image -->
    <div class="aspect-[4/3] bg-gradient-to-br from-surface to-stone-100 relative overflow-hidden shrink-0">
      <img
        v-if="product.image_path"
        :src="product.image_path"
        :alt="product.name"
        class="w-full h-full object-cover transition-all duration-700 ease-out group-hover:scale-110"
      />
      <div v-else class="w-full h-full flex flex-col items-center justify-center">
        <div class="w-14 h-14 rounded-2xl bg-stone-200/60 flex items-center justify-center mb-2 group-hover:bg-brand-100/50 transition-colors duration-500">
          <svg class="w-7 h-7 text-stone-300 group-hover:text-brand-400 transition-colors duration-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
          </svg>
        </div>
        <span class="text-xs text-stone-300 font-medium">{{ t.no_photo }}</span>
      </div>
      <!-- Overlay gradient -->
      <div class="absolute inset-0 bg-gradient-to-t from-black/8 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
    </div>

    <!-- Content -->
    <div
      class="p-3 sm:p-5 flex flex-col flex-1 cursor-pointer"
      @click="openDescription"
    >
      <h3 class="font-semibold text-stone-900 text-sm sm:text-base mb-1 truncate group-hover:text-brand-700 transition-colors duration-300">{{ product.name }}</h3>
      <p v-if="product.description" class="text-stone-400 text-xs sm:text-sm mb-3 line-clamp-2 sm:line-clamp-3 leading-relaxed">{{ shortDescription }}</p>

      <!-- Prices -->
      <div class="space-y-1 sm:space-y-2 mb-3 sm:mb-5 mt-auto">
        <div class="flex items-center justify-between">
          <span class="text-[10px] sm:text-xs font-medium text-stone-400 uppercase tracking-wide">{{ t.product_one_pill }}</span>
          <span class="text-xs sm:text-sm font-semibold text-stone-600">{{ formatPrice(product.price_per_pill) }} {{ t.currency }}</span>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-[10px] sm:text-xs font-medium text-stone-400 uppercase tracking-wide leading-tight">{{ t.product_pack }} <span class="normal-case">({{ product.quantity_per_pack }} {{ t.unit_piece }})</span></span>
          <span class="text-sm sm:text-lg font-bold text-brand-700">{{ formatPrice(product.price_per_pack) }} {{ t.currency }}</span>
        </div>
      </div>

      <!-- Add to cart -->
      <button
        @click.stop="handleAddToCart"
        class="relative w-full bg-brand-700 text-white py-2 sm:py-3 rounded-xl hover:bg-brand-800 hover:shadow-xl hover:shadow-brand-700/20 hover:-translate-y-0.5
               active:scale-[0.97] active:translate-y-0
               transition-all duration-300 ease-out font-medium flex items-center justify-center gap-1 sm:gap-2
               shadow-lg shadow-brand-700/15 group/btn text-xs sm:text-sm mt-auto"
        :class="{'animate-add-to-cart': isAdding}"
      >
        <span
          v-if="isAdding"
          class="absolute -top-2 right-3 bg-emerald-500 text-white text-[10px] font-bold rounded-full w-6 h-6 flex items-center justify-center shadow-lg shadow-emerald-500/30 animate-pop"
        >
          +1
        </span>
        <svg class="w-4 h-4 sm:w-5 sm:h-5 group-hover/btn:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
        </svg>
        {{ t.add_to_cart }}
      </button>
    </div>

    <!-- Mini panel with full description -->
    <Teleport to="body">
      <div
        v-if="isDescriptionOpen"
        class="fixed inset-0 z-[100] bg-stone-900/60 backdrop-blur-sm flex items-center justify-center p-2 sm:p-4"
        @click.self="closeDescription"
      >
        <div class="w-[96vw] max-h-[92vh] bg-white rounded-3xl shadow-2xl border border-stone-100 p-4 sm:p-6 lg:p-8 overflow-y-auto lg:overflow-hidden lg:flex lg:flex-col lg:h-[88vh]">
          <div class="flex items-start justify-between gap-4 mb-5 flex-shrink-0">
            <h4 class="text-base sm:text-lg font-bold text-stone-900">{{ t.product_description }}</h4>
            <button
              @click="closeDescription"
              class="text-xs font-semibold text-stone-500 hover:text-stone-700 flex-shrink-0"
            >
              {{ t.product_close }}
            </button>
          </div>

          <div class="grid lg:grid-cols-2 gap-5 lg:gap-8 lg:flex-1 lg:min-h-0">
            <div
              class="rounded-2xl overflow-hidden bg-stone-100 transition-all duration-700 ease-out h-56 sm:h-64 lg:h-full"
              :class="animatePanel ? 'opacity-100 translate-x-0' : 'opacity-0 -translate-x-20'"
            >
              <img
                v-if="product.image_path"
                :src="product.image_path"
                :alt="product.name"
                class="w-full h-full object-contain"
              />
              <div v-else class="w-full h-full flex items-center justify-center text-stone-400 text-sm">
                {{ t.no_photo }}
              </div>
            </div>

            <div
              class="flex flex-col transition-all duration-700 ease-out lg:overflow-y-auto lg:h-full"
              :class="animatePanel ? 'opacity-100 translate-x-0' : 'opacity-0 translate-x-20'"
            >
              <h5 class="text-3xl lg:text-4xl font-black text-stone-900 mb-2 uppercase tracking-wide">{{ product.name }}</h5>
              <p class="text-brand-700 text-2xl lg:text-3xl font-extrabold mb-5 uppercase tracking-wide">{{ formatPrice(product.price_per_pack) }} {{ t.currency }}</p>

              <div class="space-y-3 pr-1 mb-6">
                <p
                  v-for="(line, index) in animatedLines"
                  :key="`${index}-${line}`"
                  class="text-base lg:text-lg font-semibold text-stone-700 leading-relaxed opacity-0 translate-y-2 animate-line-in uppercase"
                  :style="{ animationDelay: `${index * 140}ms` }"
                >
                  {{ line }}
                </p>
              </div>

              <button
                @click="$emit('add-to-cart', product); closeDescription()"
                class="mt-auto w-full sm:w-auto sm:min-w-64 bg-brand-700 text-white py-3.5 px-7 rounded-2xl hover:bg-brand-800 hover:shadow-xl hover:shadow-brand-700/20 hover:-translate-y-0.5 active:scale-[0.98] transition-all duration-300 font-bold text-sm uppercase tracking-wide"
              >
                {{ t.add_to_cart }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useLangStore } from '../stores/lang'

const langStore = useLangStore()
const t = computed(() => langStore.t)

const props = defineProps({
  product: { type: Object, required: true }
})

const emit = defineEmits(['add-to-cart'])

const isDescriptionOpen = ref(false)
const isAdding = ref(false)
let addTimer = null
const animationNonce = ref(0)
const animatePanel = ref(false)

function handleAddToCart() {
  if (addTimer) {
    clearTimeout(addTimer)
  }
  isAdding.value = true
  addTimer = setTimeout(() => {
    isAdding.value = false
    addTimer = null
  }, 500)
  emit('add-to-cart', props.product)
}

const descriptionWords = computed(() => {
  return (props.product.description || '').trim().split(/\s+/).filter(Boolean)
})

const shortDescription = computed(() => {
  const DESCRIPTION_PREVIEW_WORDS = 20
  if (!props.product.description) return ''
  if (descriptionWords.value.length <= DESCRIPTION_PREVIEW_WORDS) return props.product.description
  return `${descriptionWords.value.slice(0, DESCRIPTION_PREVIEW_WORDS).join(' ')}...`
})

const animatedLines = computed(() => {
  animationNonce.value
  const text = (props.product.description || '').trim()
  if (!text) return ['Описание пока не добавлено.']

  const byLine = text.split('\n').map(line => line.trim()).filter(Boolean)
  if (byLine.length > 1) return byLine

  return text
    .split(/(?<=[.!?])\s+/)
    .map(line => line.trim())
    .filter(Boolean)
})

function openDescription() {
  animationNonce.value++
  isDescriptionOpen.value = true
  requestAnimationFrame(() => {
    animatePanel.value = true
  })
}

function closeDescription() {
  animatePanel.value = false
  isDescriptionOpen.value = false
}

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}
</script>

<style scoped>
.animate-line-in {
  animation: lineIn 0.45s ease forwards;
}

@keyframes lineIn {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
  .animate-add-to-cart {
    animation: addButton 0.35s ease;
  }

  .animate-pop {
    animation: popUp 0.45s ease forwards;
  }

  @keyframes addButton {
    0% {
      transform: scale(1);
      box-shadow: 0 20px 35px rgba(34, 197, 94, 0.18);
    }
    50% {
      transform: scale(1.04);
      box-shadow: 0 24px 42px rgba(34, 197, 94, 0.24);
    }
    100% {
      transform: scale(1);
      box-shadow: 0 20px 35px rgba(34, 197, 94, 0.18);
    }
  }

  @keyframes popUp {
    0% {
      opacity: 0;
      transform: translateY(4px) scale(0.9);
    }
    50% {
      opacity: 1;
      transform: translateY(-6px) scale(1.05);
    }
    100% {
      opacity: 0;
      transform: translateY(-14px) scale(1);
    }
  }
</style>
