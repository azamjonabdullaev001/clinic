<template>
  <div class="min-h-screen bg-gray-900 flex items-center justify-center px-4">
    <div class="bg-gray-800 rounded-2xl shadow-2xl p-8 w-full max-w-md border border-gray-700">
      <div class="text-center mb-8">
        <div class="w-16 h-16 bg-teal-600 rounded-2xl flex items-center justify-center mx-auto mb-4 shadow-lg">
          <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-white">Админ-панель</h1>
        <p class="text-gray-400 text-sm mt-2">Введите данные для входа</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label class="block text-sm font-medium text-gray-300 mb-1.5">Номер телефона</label>
          <div class="flex">
            <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-600 bg-gray-700 text-gray-400 text-sm font-medium">+998</span>
            <input
              v-model="phone"
              type="tel"
              maxlength="9"
              placeholder="901234567"
              class="flex-1 bg-gray-700 border border-gray-600 rounded-r-lg px-3 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
              required
            />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-300 mb-1.5">Пароль</label>
          <input
            v-model="password"
            type="password"
            placeholder="••••••••"
            class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2.5 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
            required
          />
        </div>

        <div v-if="error" class="bg-red-900/50 text-red-400 text-sm p-3 rounded-lg border border-red-800">{{ error }}</div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-teal-600 text-white py-3 rounded-lg hover:bg-teal-700 transition font-semibold disabled:opacity-50"
        >
          {{ loading ? 'Вход...' : 'Войти' }}
        </button>
      </form>

      <div class="mt-6 text-center">
        <router-link to="/" class="text-gray-400 text-sm hover:text-teal-400 inline-flex items-center gap-1">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          Вернуться на сайт
        </router-link>
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
