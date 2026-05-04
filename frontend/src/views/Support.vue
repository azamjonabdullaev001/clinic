<template>
  <div class="min-h-screen bg-gradient-to-b from-sky-50 via-white to-stone-50">
    <Navbar />

    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10 lg:py-14">
      <div class="grid lg:grid-cols-5 gap-6 lg:gap-8">
        <div class="lg:col-span-2 bg-white border border-stone-200 rounded-3xl p-5 sm:p-6 shadow-sm">
          <h1 class="text-2xl sm:text-3xl font-black text-stone-900 mb-2">{{ t.support_title }}</h1>
          <p class="text-stone-500 text-sm sm:text-base mb-5">{{ t.support_subtitle }}</p>

          <h2 class="text-xs font-semibold uppercase tracking-widest text-stone-400 mb-3">{{ t.support_faq_title }}</h2>
          <div class="space-y-2 max-h-[62vh] overflow-y-auto pr-1">
            <div v-for="faq in faqs" :key="faq.id" class="rounded-2xl border border-stone-200 overflow-hidden">
              <button
                @click="openFaqId = openFaqId === faq.id ? null : faq.id"
                class="w-full px-4 py-3 text-left bg-stone-50 hover:bg-stone-100 transition flex items-center justify-between"
              >
                <span class="font-semibold text-stone-700 text-sm">{{ faq.question }}</span>
                <svg class="w-4 h-4 text-stone-500 transition-transform" :class="openFaqId === faq.id ? 'rotate-180' : ''" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                </svg>
              </button>
              <div v-if="openFaqId === faq.id" class="px-4 py-3 bg-white border-t border-stone-200 space-y-2">
                <p v-for="answer in faq.answers" :key="answer.id" class="text-sm text-stone-600 leading-relaxed">{{ answer.text }}</p>
              </div>
            </div>

            <div v-if="faqs.length === 0" class="text-sm text-stone-400 border border-dashed border-stone-200 rounded-2xl px-4 py-6">
              {{ t.support_faq_empty }}
            </div>
          </div>
        </div>

        <div class="lg:col-span-3 bg-white border border-stone-200 rounded-3xl shadow-sm flex flex-col h-[78vh]">
          <div class="px-5 sm:px-6 py-4 border-b border-stone-200 flex items-center justify-between">
            <div>
              <h2 class="text-lg sm:text-xl font-bold text-stone-900">{{ t.support_chat_title }}</h2>
              <p class="text-xs sm:text-sm text-stone-500">{{ t.support_chat_hint }}</p>
            </div>
          </div>

          <div ref="messagesBox" class="flex-1 overflow-y-auto px-4 sm:px-6 py-5 space-y-3 bg-stone-50/60">
            <div v-for="msg in messages" :key="msg.id" :class="msg.sender_role === 'user' ? 'justify-end' : 'justify-start'" class="flex">
              <div
                :class="msg.sender_role === 'user' ? 'bg-brand-700 text-white rounded-br-md' : 'bg-white text-stone-700 rounded-bl-md border border-stone-200'"
                class="max-w-[86%] sm:max-w-[70%] px-4 py-3 rounded-2xl shadow-sm"
              >
                <p class="text-sm leading-relaxed whitespace-pre-line">{{ msg.message }}</p>
                <p class="text-[10px] mt-1.5 opacity-70">{{ formatDate(msg.created_at) }}</p>
              </div>
            </div>

            <div v-if="messages.length === 0" class="text-center text-stone-400 text-sm py-14">
              {{ t.support_no_messages }}
            </div>
          </div>

          <form @submit.prevent="sendMessage" class="p-4 sm:p-5 border-t border-stone-200 bg-white">
            <div class="flex gap-3 items-end">
              <textarea
                v-model="draft"
                rows="2"
                :placeholder="t.support_input_placeholder"
                class="flex-1 border border-stone-300 rounded-2xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500 resize-none"
              ></textarea>
              <button
                type="submit"
                :disabled="sending || !draft.trim()"
                class="bg-brand-700 text-white px-5 py-3 rounded-2xl font-semibold hover:bg-brand-800 transition disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {{ sending ? t.support_sending : t.support_send }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed, nextTick, onMounted, onUnmounted, ref } from 'vue'
import Navbar from '../components/Navbar.vue'
import { api } from '../stores/auth'
import { useLangStore } from '../stores/lang'

const langStore = useLangStore()
const t = computed(() => langStore.t)

const faqs = ref([])
const openFaqId = ref(null)
const messages = ref([])
const draft = ref('')
const sending = ref(false)
const messagesBox = ref(null)
let refreshTimer = null

function formatDate(value) {
  return new Date(value).toLocaleString('ru-RU')
}

async function loadFaqs() {
  try {
    const res = await api.get('/faqs')
    faqs.value = res.data || []
  } catch (e) {
    faqs.value = []
  }
}

async function loadThread(scrollToBottom = false) {
  try {
    const res = await api.get('/support/thread')
    messages.value = res.data?.messages || []
    if (scrollToBottom) {
      await nextTick()
      if (messagesBox.value) {
        messagesBox.value.scrollTop = messagesBox.value.scrollHeight
      }
    }
  } catch (e) {
    messages.value = []
  }
}

async function sendMessage() {
  const text = draft.value.trim()
  if (!text) return

  sending.value = true
  try {
    await api.post('/support/messages', { message: text })
    draft.value = ''
    await loadThread(true)
  } finally {
    sending.value = false
  }
}

onMounted(async () => {
  await Promise.all([loadFaqs(), loadThread(true)])
  refreshTimer = setInterval(() => {
    loadThread(false)
  }, 5000)
})

onUnmounted(() => {
  if (refreshTimer) clearInterval(refreshTimer)
})
</script>
