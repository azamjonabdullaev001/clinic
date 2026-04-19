<template>
  <div class="min-h-screen bg-gradient-to-br from-teal-50 to-white flex items-center justify-center px-4">
    <div class="bg-white rounded-2xl shadow-xl p-8 w-full max-w-md">
      <router-link to="/" class="text-teal-600 text-sm hover:underline mb-6 inline-flex items-center gap-1">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
        На главную
      </router-link>
      <h1 class="text-2xl font-bold text-gray-800 mb-2">Вход</h1>
      <p class="text-gray-400 text-sm mb-6">Войдите в свой аккаунт</p>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона</label>
          <div class="flex">
            <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm font-medium">+998</span>
            <input
              v-model="phone"
              type="tel"
              maxlength="9"
              placeholder="901234567"
              class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent transition"
              required
            />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Пароль</label>
          <input
            v-model="password"
            type="password"
            placeholder="••••••••"
            class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent transition"
            required
          />
        </div>

        <div v-if="error" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ error }}</div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-teal-600 text-white py-3 rounded-lg hover:bg-teal-700 transition font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ loading ? 'Вход...' : 'Войти' }}
        </button>
      </form>

      <p class="mt-6 text-center text-gray-500 text-sm">
        Нет аккаунта?
        <router-link to="/register" class="text-teal-600 hover:underline font-medium"> Зарегистрируйтесь</router-link>
      </p>
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
  if (phone.value.length !== 9) {
    error.value = 'Введите 9 цифр номера телефона'
    return
  }
  loading.value = true
  try {
    await authStore.login({
      phone: '998' + phone.value,
      password: password.value
    })
    router.push('/')
  } catch (e) {
    error.value = e.response?.data?.error || 'Ошибка при входе'
  } finally {
    loading.value = false
  }
}
</script>
