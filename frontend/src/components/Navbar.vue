<template>
  <nav class="fixed top-0 inset-x-0 z-40 transition-all duration-300"
       :class="scrolled ? 'bg-white/80 backdrop-blur-xl shadow-sm border-b border-gray-100' : 'bg-transparent'">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-[72px]">
        <!-- Logo -->
        <router-link to="/" class="flex items-center gap-3 group">
          <div class="w-10 h-10 rounded-xl bg-brand-800 flex items-center justify-center shadow-lg shadow-brand-800/20 group-hover:shadow-brand-800/30 transition-shadow">
            <svg class="w-5 h-5 text-gold-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 007.92 12.446A9 9 0 1112 2.992z" />
            </svg>
          </div>
          <div class="leading-tight">
            <span class="text-lg font-bold tracking-tight" :class="scrolled ? 'text-brand-900' : 'text-white'">Hair Clinic</span>
            <span class="block text-[10px] font-medium tracking-widest uppercase" :class="scrolled ? 'text-gold-600' : 'text-gold-300'">Трихология</span>
          </div>
        </router-link>

        <!-- Center nav -->
        <div class="hidden md:flex items-center gap-8">
          <a href="#doctor" class="text-sm font-medium transition-colors duration-200"
             :class="scrolled ? 'text-gray-500 hover:text-brand-800' : 'text-white/70 hover:text-white'">
            О специалисте
          </a>
          <a href="#products" class="text-sm font-medium transition-colors duration-200"
             :class="scrolled ? 'text-gray-500 hover:text-brand-800' : 'text-white/70 hover:text-white'">
            Препараты
          </a>
          <a href="#contacts" class="text-sm font-medium transition-colors duration-200"
             :class="scrolled ? 'text-gray-500 hover:text-brand-800' : 'text-white/70 hover:text-white'">
            Контакты
          </a>
        </div>

        <!-- Right side -->
        <div class="flex items-center gap-2">
          <!-- Cart -->
          <button @click="cartStore.toggle()"
                  class="relative p-2.5 rounded-xl transition-all duration-200"
                  :class="scrolled ? 'text-gray-500 hover:text-brand-800 hover:bg-brand-50' : 'text-white/80 hover:text-white hover:bg-white/10'">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
            </svg>
            <span v-if="cartStore.totalItems > 0"
                  class="absolute -top-0.5 -right-0.5 bg-gold-400 text-brand-950 text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1 shadow-sm">
              {{ cartStore.totalItems }}
            </span>
          </button>

          <!-- Auth -->
          <template v-if="authStore.isLoggedIn">
            <div class="hidden sm:flex items-center gap-2 ml-1">
              <div class="w-9 h-9 rounded-xl flex items-center justify-center font-semibold text-xs"
                   :class="scrolled ? 'bg-brand-100 text-brand-800' : 'bg-white/15 text-white'">
                {{ authStore.user?.first_name?.charAt(0) }}
              </div>
              <span class="text-sm font-medium" :class="scrolled ? 'text-gray-700' : 'text-white/90'">
                {{ authStore.user?.first_name }}
              </span>
            </div>
            <button @click="authStore.logout()"
                    class="text-xs font-medium transition-colors"
                    :class="scrolled ? 'text-gray-400 hover:text-red-500' : 'text-white/50 hover:text-white'">
              Выйти
            </button>
          </template>
          <template v-else>
            <router-link to="/login"
                         class="hidden sm:block text-sm font-medium px-4 py-2 rounded-xl transition-all duration-200"
                         :class="scrolled ? 'text-gray-600 hover:text-brand-800 hover:bg-brand-50' : 'text-white/80 hover:text-white hover:bg-white/10'">
              Войти
            </router-link>
            <router-link to="/register"
                         class="text-sm font-semibold px-5 py-2.5 rounded-xl transition-all duration-200 shadow-sm"
                         :class="scrolled ? 'bg-brand-800 text-white hover:bg-brand-900 shadow-brand-800/20' : 'bg-white text-brand-900 hover:bg-white/90'">
              Регистрация
            </router-link>
          </template>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'

const authStore = useAuthStore()
const cartStore = useCartStore()
const scrolled = ref(false)

function handleScroll() {
  scrolled.value = window.scrollY > 40
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll, { passive: true })
  handleScroll()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>
