<template>
  <div class="min-h-screen bg-brand-950 flex items-center justify-center px-4 relative overflow-hidden">
    <div class="absolute top-1/3 -right-32 w-96 h-96 bg-brand-800/20 rounded-full blur-[100px]"></div>
    <div class="absolute bottom-1/4 -left-32 w-72 h-72 bg-gold-500/5 rounded-full blur-[80px]"></div>

    <div class="relative w-full max-w-md">
      <div class="bg-white rounded-3xl shadow-2xl p-8 md:p-10">
        <div class="text-center mb-8">
          <div class="w-14 h-14 bg-brand-800 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg shadow-brand-800/20">
            <svg class="w-7 h-7 text-gold-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
            </svg>
          </div>
          <h1 class="text-xl font-serif text-gray-900">Админ-панель</h1>
          <p class="text-gray-400 text-xs mt-1">Введите данные для входа</p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Номер телефона</label>
            <div class="flex">
              <span class="inline-flex items-center px-4 rounded-l-xl border border-r-0 border-gray-200 bg-gray-50 text-gray-400 text-sm font-medium">+998</span>
              <input
                v-model="phone"
                type="tel"
                maxlength="9"
                placeholder="901234567"
                class="flex-1 border border-gray-200 rounded-r-xl px-4 py-3 text-gray-900 placeholder-gray-300 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all"
                required
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Пароль</label>
            <input
              v-model="password"
              type="password"
              placeholder="••••••••"
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-gray-900 placeholder-gray-300 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all"
              required
            />
          </div>

          <div v-if="error" class="bg-red-50 text-red-600 text-sm p-3.5 rounded-xl border border-red-100">{{ error }}</div>

          <button
            type="submit"
            :disabled="loading"
            class="w-full btn-primary py-3.5 rounded-xl disabled:opacity-50"
          >
            {{ loading ? 'Вход...' : 'Войти' }}
          </button>
        </form>

        <div class="mt-8 text-center">
          <router-link to="/" class="text-gray-400 text-sm hover:text-brand-700 inline-flex items-center gap-1.5 transition-colors">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
            Вернуться на сайт
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const phone = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    await authStore.adminLogin({
      phone: '998' + phone.value,
      password: password.value
    })
    router.push('/admin')
  } catch (e) {
    error.value = e.response?.data?.error || 'Неверные учетные данные'
  } finally {
    loading.value = false
  }
}
</script>
