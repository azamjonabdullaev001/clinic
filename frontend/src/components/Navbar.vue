<template>
  <nav class="fixed inset-x-0 top-0 z-50 bg-white border-b border-gray-100 transition-all duration-300"
       :class="scrolled ? 'shadow-sm' : ''">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-[72px]">

        <!-- Logo -->
        <router-link to="/" class="flex items-center gap-3 group shrink-0">
          <div class="w-10 h-10 rounded-xl overflow-hidden bg-white flex items-center justify-center shadow-md group-hover:shadow-lg group-hover:scale-105 transition-all duration-300 shrink-0">
            <img src="/images/patients/Jalilov.jpg" alt="Doctor Jalilov" class="w-full h-full object-contain" />
          </div>
          <div class="leading-tight">
            <span class="text-base font-bold tracking-tight text-slate-900 group-hover:text-brand-700 transition-colors duration-300">Doctor Jalilov</span>
            <span class="block text-[10px] font-medium text-slate-400">Sog'ligingiz – bizning ustuvorligimiz</span>
          </div>
        </router-link>

        <!-- Center nav (desktop) -->
        <div class="hidden md:flex items-center gap-1">
          <a href="/#products"
             class="relative text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-2 rounded-lg hover:bg-brand-50 transition-all duration-200 group">
            {{ t.nav_products_link }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <a href="/#news"
             class="relative text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-2 rounded-lg hover:bg-brand-50 transition-all duration-200 group">
            {{ t.nav_news }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <router-link to="/support"
             class="relative text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-2 rounded-lg hover:bg-brand-50 transition-all duration-200 group">
            {{ t.nav_support }}
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </router-link>
        </div>

        <!-- Right side -->
        <div class="flex items-center gap-2">

          <!-- Phone (desktop) -->
          <a href="tel:+998993251740"
             class="hidden lg:flex items-center gap-2 text-sm font-semibold text-slate-700 hover:text-brand-700 transition-colors duration-200 mr-1">
            <div class="w-8 h-8 rounded-lg bg-brand-50 flex items-center justify-center">
              <svg class="w-4 h-4 text-brand-700" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
              </svg>
            </div>
            +998 93 257 02 06
          </a>

          <!-- Cart -->
          <button @click="cartStore.toggle()"
                  class="relative p-2.5 rounded-xl text-slate-500 hover:text-brand-700 hover:bg-brand-50 transition-all duration-200">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z" />
            </svg>
            <span class="absolute -top-0.5 -right-0.5 bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1"
                  :class="countPulse ? 'animate-bump' : ''">
              {{ cartStore.totalItems }}
            </span>
          </button>

          <!-- Chat (logged-in) -->
          <button v-if="authStore.isLoggedIn"
                  @click="toggleChat"
                  class="relative p-2.5 rounded-xl text-slate-500 hover:text-brand-700 hover:bg-brand-50 transition-all duration-200">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
            <span v-if="unreadCount > 0"
                  class="absolute -top-0.5 -right-0.5 bg-red-500 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
              {{ unreadCount > 9 ? '9+' : unreadCount }}
            </span>
          </button>

          <!-- Profile button (logged-in) / Login link (not logged-in) -->
          <div v-if="authStore.isLoggedIn" class="relative profile-btn-container">
            <button @click.stop="profileOpen = !profileOpen"
                    class="flex items-center gap-1.5 p-1.5 rounded-xl hover:bg-brand-50 transition-all duration-200">
              <div class="w-8 h-8 rounded-full bg-brand-600 flex items-center justify-center text-white text-xs font-bold flex-shrink-0 select-none">
                {{ userInitials }}
              </div>
              <span class="hidden lg:block text-sm font-medium text-slate-700 max-w-[90px] truncate">{{ authStore.user?.first_name }}</span>
            </button>
            <!-- Profile dropdown -->
            <div v-if="profileOpen"
                 class="absolute right-0 top-full mt-2 w-56 bg-white rounded-2xl shadow-xl border border-slate-100 overflow-hidden z-[200]">
              <div class="px-4 py-3 border-b border-slate-100 bg-brand-50">
                <p class="font-semibold text-slate-900 text-sm">{{ authStore.user?.first_name }} {{ authStore.user?.last_name }}</p>
                <p v-if="authStore.user?.phone" class="text-xs text-brand-600 mt-0.5">{{ authStore.user.phone }}</p>
                <p v-else class="text-xs text-slate-400 mt-0.5 italic">{{ t.nav_profile_phone }}</p>
              </div>
              <button @click="authStore.logout(); profileOpen = false"
                      class="w-full flex items-center gap-2 px-4 py-3 text-sm text-red-500 hover:bg-red-50 transition-colors text-left">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"/>
                </svg>
                {{ t.nav_logout_icon_title }}
              </button>
            </div>
          </div>
          <router-link v-else to="/login"
                       class="hidden sm:flex items-center gap-1.5 px-3 py-2 rounded-xl text-slate-500 hover:text-brand-700 hover:bg-brand-50 transition-all duration-200">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z"/>
            </svg>
          </router-link>

          <!-- Language switcher (hidden on xs — accessible via hamburger menu) -->
          <div class="hidden sm:flex items-center bg-gray-100 rounded-lg p-0.5 gap-0.5">
            <button @click="langStore.setLang('ru')"
                    class="text-[11px] font-semibold px-2.5 py-1.5 rounded-md transition-all duration-200"
                    :class="langStore.current === 'ru' ? 'bg-brand-600 text-white shadow-sm' : 'text-slate-500 hover:text-slate-700'">RU</button>
            <button @click="langStore.setLang('uz')"
                    class="text-[11px] font-semibold px-2.5 py-1.5 rounded-md transition-all duration-200"
                    :class="langStore.current === 'uz' ? 'bg-brand-600 text-white shadow-sm' : 'text-slate-500 hover:text-slate-700'">UZ</button>
          </div>

          <!-- Hamburger (mobile) -->
          <button @click="mobileMenuOpen = !mobileMenuOpen"
                  class="md:hidden p-2.5 rounded-xl text-slate-500 hover:text-brand-700 hover:bg-brand-50 transition-all duration-200">
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

    <!-- Chat widget -->
    <teleport to="body">
      <div v-if="chatOpen"
           class="fixed z-[200] flex flex-col bg-white rounded-2xl shadow-2xl overflow-hidden"
           style="bottom:80px;right:16px;width:min(360px,calc(100vw - 32px));height:480px">
        <div class="flex items-center justify-between px-4 py-3 bg-brand-700 flex-shrink-0">
          <span class="text-white font-semibold text-sm">{{ t.nav_support }}</span>
          <button @click="chatOpen = false" class="text-white/70 hover:text-white p-1 rounded-lg hover:bg-white/10 transition">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <div ref="chatMsgContainer" class="flex-1 overflow-y-auto p-3 space-y-2 bg-gray-50" style="min-height:0">
          <div v-if="chatLoading" class="flex justify-center pt-8 text-gray-400 text-sm">{{ t.products_loading }}</div>
          <template v-else>
            <div v-if="chatMessages.length === 0" class="flex flex-col items-center justify-center h-full text-center text-gray-400 text-sm py-8">
              <svg class="w-10 h-10 mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
              </svg>
              {{ t.support_no_messages }}
            </div>
            <div v-for="msg in chatMessages" :key="msg.id" class="flex"
                 :class="msg.sender_role === 'user' ? 'justify-end' : 'justify-start'">
              <div class="max-w-[78%] px-3 py-2 rounded-2xl text-sm"
                   :class="msg.sender_role === 'user' ? 'bg-brand-600 text-white rounded-br-sm' : 'bg-white text-gray-800 shadow-sm rounded-bl-sm'">
                {{ msg.message }}
                <p class="text-[10px] mt-0.5 opacity-50 text-right">
                  {{ new Date(msg.created_at).toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }) }}
                </p>
              </div>
            </div>
          </template>
        </div>
        <div class="px-3 py-3 border-t bg-white flex gap-2 flex-shrink-0">
          <input v-model="chatNewMsg" @keyup.enter="sendUserMessage"
                 :placeholder="t.support_input_placeholder"
                 class="flex-1 border border-gray-200 rounded-xl px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-brand-500 bg-gray-50" />
          <button @click="sendUserMessage" :disabled="!chatNewMsg.trim() || chatSending"
                  class="bg-brand-600 text-white w-10 h-10 rounded-xl flex items-center justify-center hover:bg-brand-700 transition disabled:opacity-40 flex-shrink-0">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
            </svg>
          </button>
        </div>
      </div>
    </teleport>

    <!-- Mobile menu -->
    <div v-if="mobileMenuOpen" class="md:hidden border-t border-gray-100 bg-white px-4 py-4 flex flex-col gap-1">
      <a href="/#products" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50 transition-all">
        {{ t.nav_products_link }}
      </a>
      <a href="/#news" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50 transition-all">
        {{ t.nav_news }}
      </a>
      <router-link to="/support" @click="mobileMenuOpen = false"
         class="text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50 transition-all">
        {{ t.nav_support }}
      </router-link>
      <a href="tel:+998993251740"
         class="flex items-center gap-2 text-sm font-semibold text-slate-700 px-4 py-3 rounded-xl hover:bg-brand-50 transition-all">
        <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
        </svg>
        +998 93 257 02 06
      </a>
      <div class="border-t border-slate-100 mt-2 pt-3 flex flex-col gap-2">
        <div class="flex items-center gap-2 px-4 py-2">
          <button @click="langStore.setLang('ru')" class="text-xs font-semibold px-3 py-1.5 rounded-lg transition"
                  :class="langStore.current === 'ru' ? 'bg-brand-600 text-white' : 'bg-gray-100 text-slate-500'">RU</button>
          <button @click="langStore.setLang('uz')" class="text-xs font-semibold px-3 py-1.5 rounded-lg transition"
                  :class="langStore.current === 'uz' ? 'bg-brand-600 text-white' : 'bg-gray-100 text-slate-500'">UZ</button>
        </div>
        <template v-if="authStore.isLoggedIn">
          <!-- User info card -->
          <div class="px-4 py-3 bg-brand-50 rounded-xl mb-1">
            <p class="text-sm font-semibold text-slate-900">{{ authStore.user?.first_name }} {{ authStore.user?.last_name }}</p>
            <p v-if="authStore.user?.phone" class="text-xs text-brand-600 mt-0.5">{{ authStore.user.phone }}</p>
            <p v-else class="text-xs text-slate-400 mt-0.5 italic">{{ t.nav_profile_phone }}</p>
          </div>
          <button @click="authStore.logout(); mobileMenuOpen = false"
                  class="text-sm font-medium text-red-500 hover:text-red-600 px-4 py-3 rounded-xl hover:bg-red-50 transition text-left flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"/>
            </svg>
            {{ t.nav_logout_icon_title }}
          </button>
        </template>
        <template v-else>
          <router-link to="/login" @click="mobileMenuOpen = false"
                       class="text-sm font-medium text-slate-600 hover:text-brand-700 px-4 py-3 rounded-xl hover:bg-brand-50 transition">
            {{ t.nav_login }}
          </router-link>
          <router-link to="/register" @click="mobileMenuOpen = false"
                       class="text-sm font-semibold px-4 py-3 rounded-xl bg-brand-700 text-white hover:bg-brand-800 transition text-center">
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
const chatOpen = ref(false)
const profileOpen = ref(false)

const userInitials = computed(() => {
  const u = authStore.user
  if (!u) return '?'
  const first = (u.first_name || '').charAt(0).toUpperCase()
  const last = (u.last_name || '').charAt(0).toUpperCase()
  return (first + last) || '?'
})
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
  } catch { /* ignore */ }
}

async function loadChatThread() {
  chatLoading.value = true
  try {
    const res = await api.get('/support/thread')
    chatMessages.value = res.data.messages || []
    unreadCount.value = 0
  } catch { chatMessages.value = [] }
  finally {
    chatLoading.value = false
    await nextTick()
    if (chatMsgContainer.value) chatMsgContainer.value.scrollTop = chatMsgContainer.value.scrollHeight
  }
}

async function toggleChat() {
  chatOpen.value = !chatOpen.value
  if (chatOpen.value) await loadChatThread()
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
    if (chatMsgContainer.value) chatMsgContainer.value.scrollTop = chatMsgContainer.value.scrollHeight
  } catch (e) {
    alert(e.response?.data?.error || 'Xatolik')
  } finally { chatSending.value = false }
}

watch(() => authStore.isLoggedIn, (v) => {
  if (v) { loadUnreadCount(); unreadTimer = setInterval(loadUnreadCount, 30000) }
  else { clearInterval(unreadTimer); unreadCount.value = 0; chatOpen.value = false }
}, { immediate: true })

watch(() => cartStore.totalItems, (n, o) => {
  if (n > o) { countPulse.value = true; setTimeout(() => { countPulse.value = false }, 400) }
})

function handleScroll() {
  scrolled.value = window.scrollY > 40
  if (mobileMenuOpen.value) mobileMenuOpen.value = false
  if (profileOpen.value) profileOpen.value = false
}

function handleDocumentClick(e) {
  if (!e.target.closest('.profile-btn-container')) {
    profileOpen.value = false
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll, { passive: true })
  document.addEventListener('click', handleDocumentClick)
  handleScroll()
})
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  document.removeEventListener('click', handleDocumentClick)
  clearInterval(unreadTimer)
})
</script>

<style scoped>
@keyframes bump { 0% { transform: scale(1); } 30% { transform: scale(1.25); } 100% { transform: scale(1); } }
.animate-bump { animation: bump 0.35s ease; }
</style>
