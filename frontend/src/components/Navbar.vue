<template>
  <nav class="fixed inset-x-0 top-0 z-50 transition-all duration-500"
       :class="scrolled ? 'bg-white shadow-lg shadow-stone-900/[0.04] border-b border-stone-200/50' : 'bg-white/95'">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-[72px]">
        <!-- Logo -->
        <router-link to="/" class="flex items-center gap-3 group">
          <div class="w-10 h-10 rounded-xl bg-white border border-stone-200 flex items-center justify-center shadow-md group-hover:shadow-lg group-hover:scale-105 transition-all duration-300 overflow-hidden">
            <img src="/images/patients/Jalilov.jpg" alt="Doctor Jalilov logotipi" class="w-full h-full object-contain p-0.5" />
          </div>
          <div class="leading-tight">
            <span class="text-lg font-bold tracking-tight text-stone-900 group-hover:text-brand-700 transition-colors duration-300">Doctor Jalilov</span>
            <span class="block text-[10px] font-medium tracking-widest uppercase text-brand-600">{{ t.footer_trich }}</span>
          </div>
        </router-link>

        <!-- Center nav (desktop) -->
        <div class="hidden md:flex items-center gap-1">
          <a href="/#doctor" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            {{ t.nav_specialist }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <a href="/#products" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            {{ t.nav_products }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <a href="/#news" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            {{ t.nav_news }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <router-link to="/support" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            {{ t.nav_support }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </router-link>
        </div>

        <!-- Right side -->
        <div class="flex items-center gap-2">
          <!-- Language switcher -->
          <div class="flex items-center bg-stone-100 rounded-xl p-1 gap-0.5">
            <button
              @click="langStore.setLang('ru')"
              class="text-xs font-semibold px-3 py-1.5 rounded-lg transition-all duration-200"
              :class="langStore.current === 'ru' ? 'bg-brand-600 text-white shadow-sm' : 'text-stone-500 hover:text-stone-700'"
            >RU</button>
            <button
              @click="langStore.setLang('uz')"
              class="text-xs font-semibold px-3 py-1.5 rounded-lg transition-all duration-200"
              :class="langStore.current === 'uz' ? 'bg-brand-600 text-white shadow-sm' : 'text-stone-500 hover:text-stone-700'"
            >UZ</button>
          </div>
          <!-- Cart -->
          <button @click="cartStore.toggle()"
                  class="relative p-2.5 rounded-xl text-stone-500 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300 group">
            <svg class="w-5 h-5 group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
            </svg>
            <span v-if="cartStore.totalItems > 0"
                  :class="['absolute -top-0.5 -right-0.5 bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1 animate-scale-in', countPulse ? 'animate-bump' : '']">
              {{ cartStore.totalItems }}
            </span>
          </button>

          <!-- Chat notification (logged-in users only) -->
          <button v-if="authStore.isLoggedIn"
                  @click="toggleChat"
                  class="relative p-2.5 rounded-xl text-stone-500 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300 group">
            <svg class="w-5 h-5 group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
            <span v-if="unreadCount > 0"
                  class="absolute -top-0.5 -right-0.5 bg-red-500 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
              {{ unreadCount > 9 ? '9+' : unreadCount }}
            </span>
          </button>

          <!-- Auth -->
          <template v-if="authStore.isLoggedIn">
            <div class="hidden sm:flex items-center gap-2 ml-1 group cursor-default">
              <div class="w-9 h-9 rounded-xl bg-brand-100 text-brand-700 flex items-center justify-center font-semibold text-xs group-hover:bg-brand-200 transition-colors duration-300">
                {{ authStore.user?.first_name?.charAt(0) }}
              </div>
              <span class="text-sm font-medium text-stone-700">
                {{ authStore.user?.first_name }}
              </span>
            </div>
            <button @click="authStore.logout()"
                    class="text-xs font-medium text-stone-400 hover:text-red-500 transition-colors duration-300">
              {{ t.nav_logout }}
            </button>
          </template>
          <template v-else>
            <router-link to="/login"
                         class="hidden sm:block text-sm font-medium px-4 py-2 rounded-xl text-stone-600 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300">
              {{ t.nav_login }}
            </router-link>
            <router-link to="/register"
                         class="hidden sm:block text-sm font-semibold px-5 py-2.5 rounded-xl bg-brand-700 text-white hover:bg-brand-800 hover:shadow-lg hover:shadow-brand-700/20 hover:-translate-y-0.5 active:translate-y-0 transition-all duration-300 shadow-sm shadow-brand-700/20">
              {{ t.nav_register }}
            </router-link>
          </template>

          <!-- Hamburger (mobile only) -->
          <button
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="md:hidden p-2.5 rounded-xl text-stone-500 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300"
            :aria-expanded="mobileMenuOpen"
          >
            <svg v-if="!mobileMenuOpen" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mini-chat widget fixed bottom-right -->
    <teleport to="body">
      <div
        v-if="chatOpen"
        class="fixed z-[200] flex flex-col bg-white rounded-2xl shadow-2xl overflow-hidden"
        style="bottom:80px;right:16px;width:360px;height:480px"
      >
        <!-- Header -->
        <div class="flex items-center justify-between px-4 py-3 bg-brand-700 flex-shrink-0">
          <div class="flex items-center gap-2">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
            <span class="text-white font-semibold text-sm">{{ t.nav_support }}</span>
          </div>
          <button @click="chatOpen = false" class="text-white/70 hover:text-white transition p-1 rounded-lg hover:bg-white/10">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <!-- Messages -->
        <div ref="chatMsgContainer" class="flex-1 overflow-y-auto p-3 space-y-2 bg-gray-50" style="min-height:0">
          <div v-if="chatLoading" class="flex justify-center pt-8 text-gray-400 text-sm">{{ t.products_loading }}</div>
          <template v-else>
            <div v-if="chatMessages.length === 0" class="flex flex-col items-center justify-center h-full text-center text-gray-400 text-sm py-8">
              <svg class="w-10 h-10 mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
              </svg>
              {{ t.support_no_messages }}
            </div>
            <div
              v-for="msg in chatMessages"
              :key="msg.id"
              class="flex"
              :class="msg.sender_role === 'user' ? 'justify-end' : 'justify-start'"
            >
              <div
                class="max-w-[78%] px-3 py-2 rounded-2xl text-sm"
                :class="msg.sender_role === 'user'
                  ? 'bg-brand-600 text-white rounded-br-sm'
                  : 'bg-white text-gray-800 shadow-sm rounded-bl-sm'"
              >
                {{ msg.message }}
                <p class="text-[10px] mt-0.5 opacity-50 text-right">
                  {{ new Date(msg.created_at).toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }) }}
                </p>
              </div>
            </div>
          </template>
        </div>
        <!-- Input -->
        <div class="px-3 py-3 border-t bg-white flex gap-2 flex-shrink-0">
          <input
            v-model="chatNewMsg"
            @keyup.enter="sendUserMessage"
            :placeholder="t.support_input_placeholder"
            class="flex-1 border border-gray-200 rounded-xl px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-brand-500 bg-gray-50"
          />
          <button
            @click="sendUserMessage"
            :disabled="!chatNewMsg.trim() || chatSending"
            class="bg-brand-600 text-white w-10 h-10 rounded-xl flex items-center justify-center hover:bg-brand-700 transition disabled:opacity-40 flex-shrink-0"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
            </svg>
          </button>
        </div>
      </div>
    </teleport>

    <!-- Mobile menu dropdown -->
    <div
      v-if="mobileMenuOpen"
      class="md:hidden border-t border-stone-200/50 bg-white px-4 py-4 flex flex-col gap-1"
    >
      <a href="/#doctor" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-stone-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50/60 transition-all duration-200">
        {{ t.nav_specialist }}
      </a>
      <a href="/#products" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-stone-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50/60 transition-all duration-200">
        {{ t.nav_products }}
      </a>
      <a href="/#news" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-stone-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50/60 transition-all duration-200">
        {{ t.nav_news }}
      </a>
      <router-link to="/support" @click="mobileMenuOpen = false"
                   class="text-sm font-medium text-stone-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50/60 transition-all duration-200">
        {{ t.nav_support }}
      </router-link>
      <div class="border-t border-stone-100 mt-2 pt-3 flex flex-col gap-2">
        <template v-if="authStore.isLoggedIn">
          <div class="flex items-center gap-3 px-4 py-2">
            <div class="w-8 h-8 rounded-xl bg-brand-100 text-brand-700 flex items-center justify-center font-semibold text-xs">
              {{ authStore.user?.first_name?.charAt(0) }}
            </div>
            <span class="text-sm font-medium text-stone-700">{{ authStore.user?.first_name }}</span>
          </div>
          <button @click="authStore.logout(); mobileMenuOpen = false"
                  class="text-sm font-medium text-red-500 hover:text-red-600 px-4 py-3 rounded-xl hover:bg-red-50 transition-all duration-200 text-left">
            {{ t.nav_logout }}
          </button>
        </template>
        <template v-else>
          <router-link to="/login" @click="mobileMenuOpen = false"
                       class="text-sm font-medium text-stone-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50/60 transition-all duration-200">
            {{ t.nav_login }}
          </router-link>
          <router-link to="/register" @click="mobileMenuOpen = false"
                       class="text-sm font-semibold px-4 py-3 rounded-xl bg-brand-700 text-white hover:bg-brand-800 transition-all duration-200 text-center">
            {{ t.nav_register }}
          </router-link>
        </template>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, watch, onMounted, onUnmounted, computed, nextTick } from 'vue'
import { useAuthStore, api } from '../stores/auth'
import { useCartStore } from '../stores/cart'
import { useLangStore } from '../stores/lang'

const authStore = useAuthStore()
const cartStore = useCartStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)
const scrolled = ref(false)
const mobileMenuOpen = ref(false)
const countPulse = ref(false)

// Chat
const chatOpen = ref(false)
const chatMessages = ref([])
const chatNewMsg = ref('')
const chatSending = ref(false)
const chatLoading = ref(false)
const unreadCount = ref(0)
const chatMsgContainer = ref(null)
let unreadTimer = null

async function loadUnreadCount() {
  if (!authStore.isLoggedIn) return
  try {
    const res = await api.get('/support/unread-count')
    unreadCount.value = res.data.count || 0
  } catch (e) { /* ignore */ }
}

async function loadChatThread() {
  chatLoading.value = true
  try {
    const res = await api.get('/support/thread')
    chatMessages.value = res.data.messages || []
    unreadCount.value = 0
  } catch (e) {
    chatMessages.value = []
  } finally {
    chatLoading.value = false
    await nextTick()
    scrollChatBottom()
  }
}

function scrollChatBottom() {
  if (chatMsgContainer.value) {
    chatMsgContainer.value.scrollTop = chatMsgContainer.value.scrollHeight
  }
}

async function toggleChat() {
  chatOpen.value = !chatOpen.value
  if (chatOpen.value) {
    await loadChatThread()
  }
}

async function sendUserMessage() {
  if (!chatNewMsg.value.trim() || chatSending.value) return
  const text = chatNewMsg.value.trim()
  chatSending.value = true
  try {
    const res = await api.post('/support/messages', { message: text })
    chatMessages.value.push(res.data)
    chatNewMsg.value = ''
    await nextTick()
    scrollChatBottom()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при отправке')
  } finally {
    chatSending.value = false
  }
}

watch(() => authStore.isLoggedIn, (loggedIn) => {
  if (loggedIn) {
    loadUnreadCount()
    unreadTimer = setInterval(loadUnreadCount, 30000)
  } else {
    clearInterval(unreadTimer)
    unreadCount.value = 0
    chatOpen.value = false
  }
}, { immediate: true })

watch(() => cartStore.totalItems, (newValue, oldValue) => {
  if (newValue > oldValue) {
    countPulse.value = true
    setTimeout(() => {
      countPulse.value = false
    }, 400)
  }
})

function handleScroll() {
  scrolled.value = window.scrollY > 40
  if (mobileMenuOpen.value) mobileMenuOpen.value = false
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll, { passive: true })
  handleScroll()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  clearInterval(unreadTimer)
})
</script>

<style scoped>
@keyframes bump {
  0% {
    transform: scale(1);
  }
  30% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}
.animate-bump {
  animation: bump 0.35s ease;
}
</style>

