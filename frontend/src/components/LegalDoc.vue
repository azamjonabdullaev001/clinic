<template>
  <div class="min-h-screen bg-gradient-to-b from-sky-50 via-white to-stone-50 flex flex-col">
    <Navbar />

    <main class="flex-1 max-w-3xl w-full mx-auto px-4 sm:px-6 lg:px-8 pt-24 sm:pt-28 pb-16">
      <!-- Header -->
      <div class="mb-8">
        <router-link to="/" class="inline-flex items-center gap-1.5 text-sm text-stone-400 hover:text-brand-600 transition mb-4">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
          </svg>
          {{ backLabel }}
        </router-link>
        <h1 class="text-3xl sm:text-4xl font-black text-stone-900 tracking-tight">{{ title }}</h1>
        <p v-if="updated" class="mt-2 text-sm text-stone-400">{{ updatedLabel }}: {{ updated }}</p>
      </div>

      <!-- Body -->
      <article class="bg-white border border-stone-200 rounded-3xl shadow-sm p-6 sm:p-8 space-y-8">
        <section v-for="(sec, i) in sections" :key="i">
          <h2 v-if="sec.h" class="text-lg sm:text-xl font-bold text-stone-900 mb-3">
            <span v-if="numbered" class="text-brand-600">{{ i + 1 }}.</span> {{ sec.h }}
          </h2>
          <div class="space-y-3">
            <template v-for="(block, j) in sec.blocks" :key="j">
              <p v-if="block.t === 'p'" class="text-[15px] leading-relaxed text-stone-600">{{ block.text }}</p>
              <ul v-else-if="block.t === 'ul'" class="space-y-1.5 pl-1">
                <li v-for="(item, k) in block.items" :key="k" class="flex gap-2.5 text-[15px] leading-relaxed text-stone-600">
                  <span class="mt-2 w-1.5 h-1.5 rounded-full bg-brand-400 shrink-0"></span>
                  <span>{{ item }}</span>
                </li>
              </ul>
              <dl v-else-if="block.t === 'kv'" class="rounded-2xl bg-stone-50 border border-stone-200 divide-y divide-stone-200 overflow-hidden">
                <div v-for="(row, k) in block.rows" :key="k" class="flex flex-col sm:flex-row sm:items-center gap-0.5 sm:gap-4 px-4 py-2.5">
                  <dt class="text-xs sm:text-sm text-stone-400 sm:w-56 shrink-0">{{ row[0] }}</dt>
                  <dd class="text-sm font-medium text-stone-700 break-words">{{ row[1] }}</dd>
                </div>
              </dl>
            </template>
          </div>
        </section>

        <!-- Payment systems footer inside doc -->
        <div v-if="showLogos" class="pt-6 border-t border-stone-100">
          <p class="text-xs text-stone-400 mb-3">{{ logosLabel }}</p>
          <PaymentLogos />
        </div>
      </article>
    </main>

    <SimpleFooter />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import Navbar from './Navbar.vue'
import SimpleFooter from './SimpleFooter.vue'
import PaymentLogos from './PaymentLogos.vue'
import { useLangStore } from '../stores/lang'

defineProps({
  title: { type: String, required: true },
  updated: { type: String, default: '' },
  sections: { type: Array, default: () => [] },
  numbered: { type: Boolean, default: true },
  showLogos: { type: Boolean, default: false },
})

const langStore = useLangStore()
const uz = computed(() => langStore.current === 'uz')
const backLabel = computed(() => (uz.value ? 'Bosh sahifaga' : 'На главную'))
const updatedLabel = computed(() => (uz.value ? 'Yangilangan' : 'Обновлено'))
const logosLabel = computed(() => (uz.value ? 'Qabul qilinadigan to‘lov tizimlari' : 'Принимаемые платёжные системы'))
</script>
