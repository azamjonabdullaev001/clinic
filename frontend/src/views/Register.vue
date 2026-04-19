<template>
  <div class="min-h-screen bg-gradient-to-br from-teal-50 to-white flex items-center justify-center px-4 py-8">
    <div class="bg-white rounded-2xl shadow-xl p-8 w-full max-w-md">
      <router-link to="/" class="text-teal-600 text-sm hover:underline mb-6 inline-flex items-center gap-1">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
        На главную
      </router-link>
      <h1 class="text-2xl font-bold text-gray-800 mb-2">Регистрация</h1>
      <p class="text-gray-400 text-sm mb-6">Создайте аккаунт для заказа препаратов</p>

      <form @submit.prevent="handleRegister" class="space-y-4">
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Имя <span class="text-red-400">*</span></label>
            <input v-model="form.first_name" type="text" required
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Фамилия <span class="text-red-400">*</span></label>
            <input v-model="form.last_name" type="text" required
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Отчество</label>
          <input v-model="form.middle_name" type="text"
            class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона <span class="text-red-400">*</span></label>
          <div class="flex">
            <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm font-medium">+998</span>
            <input
              v-model="phone"
              type="tel"
              maxlength="9"
              placeholder="901234567"
              class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
              required
            />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Пароль <span class="text-red-400">*</span></label>
          <input v-model="form.password" type="password" minlength="6" required
            class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
            placeholder="Минимум 6 символов" />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">Подтверждение пароля <span class="text-red-400">*</span></label>
          <input v-model="form.confirm_password" type="password" required
            class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
            placeholder="Повторите пароль" />
        </div>

        <div v-if="error" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ error }}</div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-teal-600 text-white py-3 rounded-lg hover:bg-teal-700 transition font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {{ loading ? 'Регистрация...' : 'Зарегистрироваться' }}
        </button>
      </form>

      <p class="mt-6 text-center text-gray-500 text-sm">
        Уже есть аккаунт?
        <router-link to="/login" class="text-teal-600 hover:underline font-medium"> Войдите</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const phone = ref('')
const form = reactive({
  first_name: '',
  last_name: '',
  middle_name: '',
  password: '',
  confirm_password: ''
})
const error = ref('')
const loading = ref(false)

async function handleRegister() {
  error.value = ''

  if (form.password !== form.confirm_password) {
    error.value = 'Пароли не совпадают'
    return
  }

  if (phone.value.length !== 9) {
    error.value = 'Введите 9 цифр номера телефона'
    return
  }

  loading.value = true
  try {
    await authStore.register({
      ...form,
      phone: '998' + phone.value
    })
    router.push('/')
  } catch (e) {
    error.value = e.response?.data?.error || 'Ошибка при регистрации'
  } finally {
    loading.value = false
  }
}
</script>
