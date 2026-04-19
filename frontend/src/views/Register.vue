<template>
  <div class="min-h-screen bg-brand-950 flex items-center justify-center px-4 py-10 relative overflow-hidden">
    <!-- Background decoration -->
    <div class="absolute top-1/4 -right-32 w-96 h-96 bg-brand-800/20 rounded-full blur-[100px]"></div>
    <div class="absolute bottom-1/3 -left-32 w-72 h-72 bg-gold-500/5 rounded-full blur-[80px]"></div>

    <div class="relative w-full max-w-md">
      <!-- Back link -->
      <router-link to="/" class="inline-flex items-center gap-1.5 text-sm text-white/40 hover:text-white/70 transition-colors mb-8">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
        На главную
      </router-link>

      <!-- Card -->
      <div class="bg-white rounded-3xl shadow-2xl p-8 md:p-10">
        <div class="flex items-center gap-3 mb-8">
          <div class="w-10 h-10 rounded-xl bg-brand-800 flex items-center justify-center">
            <svg class="w-5 h-5 text-gold-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 007.92 12.446A9 9 0 1112 2.992z" />
            </svg>
          </div>
          <div>
            <h1 class="text-xl font-serif text-gray-900">Регистрация</h1>
            <p class="text-xs text-gray-400">Создайте аккаунт для заказа препаратов</p>
          </div>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-4">
          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Имя <span class="text-red-400">*</span></label>
              <input v-model="form.first_name" type="text" required
                class="w-full border border-gray-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Фамилия <span class="text-red-400">*</span></label>
              <input v-model="form.last_name" type="text" required
                class="w-full border border-gray-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900" />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Отчество</label>
            <input v-model="form.middle_name" type="text"
              class="w-full border border-gray-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Номер телефона <span class="text-red-400">*</span></label>
            <div class="flex">
              <span class="inline-flex items-center px-4 rounded-l-xl border border-r-0 border-gray-200 bg-gray-50 text-gray-400 text-sm font-medium">+998</span>
              <input
                v-model="phone"
                type="tel"
                maxlength="9"
                placeholder="901234567"
                class="flex-1 border border-gray-200 rounded-r-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900"
                required
              />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Пароль <span class="text-red-400">*</span></label>
            <input v-model="form.password" type="password" minlength="6" required
              class="w-full border border-gray-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900"
              placeholder="Минимум 6 символов" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Подтверждение пароля <span class="text-red-400">*</span></label>
            <input v-model="form.confirm_password" type="password" required
              class="w-full border border-gray-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-600/20 focus:border-brand-600 transition-all text-gray-900"
              placeholder="Повторите пароль" />
          </div>

          <div v-if="error" class="bg-red-50 text-red-600 text-sm p-3.5 rounded-xl border border-red-100">{{ error }}</div>

          <button
            type="submit"
            :disabled="loading"
            class="w-full btn-primary py-3.5 rounded-xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ loading ? 'Регистрация...' : 'Зарегистрироваться' }}
          </button>
        </form>

        <p class="mt-8 text-center text-gray-400 text-sm">
          Уже есть аккаунт?
          <router-link to="/login" class="text-brand-700 hover:text-brand-800 font-semibold transition-colors"> Войдите</router-link>
        </p>
      </div>
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
