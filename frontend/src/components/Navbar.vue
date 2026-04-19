<template>
  <nav class="fixed top-0 inset-x-0 z-40 transition-all duration-500"
       :class="scrolled ? 'bg-white/95 backdrop-blur-2xl shadow-lg shadow-stone-900/[0.04] border-b border-stone-200/50' : 'bg-surface/80 backdrop-blur-xl'">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-[72px]">
        <!-- Logo -->
        <router-link to="/" class="flex items-center gap-3 group">
          <div class="w-10 h-10 rounded-xl bg-brand-600 flex items-center justify-center shadow-md shadow-brand-600/20 group-hover:shadow-lg group-hover:shadow-brand-600/30 group-hover:scale-105 transition-all duration-300">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 007.92 12.446A9 9 0 1112 2.992z" />
            </svg>
          </div>
          <div class="leading-tight">
            <span class="text-lg font-bold tracking-tight text-stone-900 group-hover:text-brand-700 transition-colors duration-300">Hair Clinic</span>
            <span class="block text-[10px] font-medium tracking-widest uppercase text-brand-600">Трихология</span>
          </div>
        </router-link>

        <!-- Center nav -->
        <div class="hidden md:flex items-center gap-1">
          <a href="#doctor" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            О специалисте
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <a href="#products" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            Препараты
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
          <a href="#contacts" class="relative text-sm font-medium text-stone-500 hover:text-brand-700 px-4 py-2 rounded-xl hover:bg-brand-50/60 transition-all duration-300 group">
            Контакты
            <span class="absolute bottom-0.5 left-1/2 -translate-x-1/2 w-0 h-0.5 bg-brand-500 rounded-full group-hover:w-5 transition-all duration-300"></span>
          </a>
        </div>

        <!-- Right side -->
        <div class="flex items-center gap-2">
          <!-- Cart -->
          <button @click="cartStore.toggle()"
                  class="relative p-2.5 rounded-xl text-stone-500 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300 group">
            <svg class="w-5 h-5 group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
            </svg>
            <span v-if="cartStore.totalItems > 0"
                  class="absolute -top-0.5 -right-0.5 bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1 animate-scale-in">
              {{ cartStore.totalItems }}
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
              Выйти
            </button>
          </template>
          <template v-else>
            <router-link to="/login"
                         class="hidden sm:block text-sm font-medium px-4 py-2 rounded-xl text-stone-600 hover:text-brand-700 hover:bg-brand-50/60 transition-all duration-300">
              Войти
            </router-link>
            <router-link to="/register"
                         class="text-sm font-semibold px-5 py-2.5 rounded-xl bg-brand-700 text-white hover:bg-brand-800 hover:shadow-lg hover:shadow-brand-700/20 hover:-translate-y-0.5 active:translate-y-0 transition-all duration-300 shadow-sm shadow-brand-700/20">
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
