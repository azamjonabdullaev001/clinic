<template>
  <nav class="bg-white shadow-sm sticky top-0 z-40 border-b">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <router-link to="/" class="flex items-center space-x-2.5">
          <div class="w-9 h-9 bg-teal-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </div>
          <span class="text-xl font-bold text-gray-800">Hair <span class="text-teal-600">Clinic</span></span>
        </router-link>

        <div class="hidden md:flex items-center space-x-6">
          <a href="#doctor" class="text-sm text-gray-600 hover:text-teal-600 transition font-medium">О нас</a>
          <a href="#products" class="text-sm text-gray-600 hover:text-teal-600 transition font-medium">Препараты</a>
        </div>

        <div class="flex items-center space-x-3">
          <button @click="cartStore.toggle()" class="relative p-2 text-gray-500 hover:text-teal-600 transition rounded-lg hover:bg-gray-100">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 100 4 2 2 0 000-4z" />
            </svg>
            <span
              v-if="cartStore.totalItems > 0"
              class="absolute -top-0.5 -right-0.5 bg-red-500 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1"
            >
              {{ cartStore.totalItems }}
            </span>
          </button>

          <template v-if="authStore.isLoggedIn">
            <div class="hidden sm:flex items-center gap-2 text-sm">
              <div class="w-8 h-8 bg-teal-100 text-teal-700 rounded-full flex items-center justify-center font-semibold text-xs">
                {{ authStore.user?.first_name?.charAt(0) }}
              </div>
              <span class="text-gray-700 font-medium">{{ authStore.user?.first_name }}</span>
            </div>
            <button @click="authStore.logout()" class="text-sm text-gray-400 hover:text-red-500 transition">Выйти</button>
          </template>
          <template v-else>
            <router-link to="/login" class="text-sm text-gray-600 hover:text-teal-600 font-medium transition hidden sm:block">Войти</router-link>
            <router-link to="/register" class="bg-teal-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-teal-700 transition shadow-sm">
              Регистрация
            </router-link>
          </template>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'

const authStore = useAuthStore()
const cartStore = useCartStore()
</script>
