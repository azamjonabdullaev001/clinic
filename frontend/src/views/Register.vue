<template>
  <div class="min-h-screen bg-surface flex items-center justify-center px-4 py-10 relative overflow-hidden">
    <!-- Background decoration -->
    <div class="absolute top-1/4 -right-32 w-96 h-96 bg-brand-200/15 rounded-full blur-[100px] animate-pulse-soft"></div>
    <div class="absolute bottom-1/3 -left-32 w-72 h-72 bg-brand-100/15 rounded-full blur-[80px] animate-pulse-soft" style="animation-delay: 2s"></div>

    <div class="relative w-full max-w-md">
      <!-- Back link -->
      <router-link to="/" class="inline-flex items-center gap-1.5 text-sm text-stone-400 hover:text-stone-600 transition-colors mb-8">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
        На главную
      </router-link>

      <!-- Card -->
      <div class="bg-white/90 backdrop-blur-xl rounded-3xl shadow-xl shadow-stone-200/40 p-8 md:p-10 animate-fade-up" style="animation-delay: 0.1s; opacity: 0">
        <div class="flex items-center gap-3 mb-8">
          <div class="w-10 h-10 rounded-xl bg-brand-600 flex items-center justify-center hover:scale-110 transition-transform duration-300">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 007.92 12.446A9 9 0 1112 2.992z" />
            </svg>
          </div>
          <div>
            <h1 class="text-xl font-serif text-stone-900">Регистрация</h1>
            <p class="text-xs text-stone-400">Создайте аккаунт для заказа препаратов</p>
          </div>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-4">
          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-sm font-medium text-stone-700 mb-2">Имя <span class="text-red-400">*</span></label>
              <input v-model="form.first_name" type="text" required
                class="w-full border border-stone-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900" />
            </div>
            <div>
              <label class="block text-sm font-medium text-stone-700 mb-2">Фамилия <span class="text-red-400">*</span></label>
              <input v-model="form.last_name" type="text" required
                class="w-full border border-stone-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900" />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-stone-700 mb-2">Отчество</label>
            <input v-model="form.middle_name" type="text"
              class="w-full border border-stone-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900" />
          </div>

          <div>
            <label class="block text-sm font-medium text-stone-700 mb-2">Номер телефона <span class="text-red-400">*</span></label>
            <div class="flex">
              <span class="inline-flex items-center px-4 rounded-l-xl border border-r-0 border-stone-200 bg-stone-50 text-stone-400 text-sm font-medium">+998</span>
              <input
                v-model="phone"
                type="tel"
                maxlength="9"
                placeholder="901234567"
                class="flex-1 border border-stone-200 rounded-r-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900"
                required
              />
            </div>
          </div>

          <!-- Delivery Address -->
          <div>
            <label class="block text-sm font-medium text-stone-700 mb-2">
              Адрес доставки <span class="text-red-400">*</span>
            </label>
            <div class="relative">
              <input
                v-model="form.delivery_address"
                type="text"
                placeholder="Например: Андижанская область, Кургантепинский район, г. Карасу"
                class="w-full border rounded-xl px-4 py-3 pr-12 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900"
                :class="addressError ? 'border-red-400 bg-red-50/30' : 'border-stone-200'"
              />
              <button
                type="button"
                @click="detectLocation"
                :disabled="locating"
                ref="geoBtn"
                class="absolute right-2 top-1/2 -translate-y-1/2 p-2 rounded-lg text-stone-400 hover:text-brand-600 hover:bg-brand-50 transition-all disabled:opacity-40"
                title="Определить моё местоположение"
              >
                <svg v-if="!locating" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                </svg>
                <svg v-else class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
              </button>
            </div>
            <p v-if="locationError" class="text-xs text-red-500 mt-1">{{ locationError }}</p>
            <p v-else-if="addressError" class="text-xs text-red-500 mt-1">{{ addressError }}</p>
            <p v-else class="text-xs text-stone-400 mt-1.5 flex items-center gap-1.5">
              <span class="inline-flex items-center gap-1">
                Не знаете свой адрес?
                <span class="relative inline-flex items-center gap-1 text-brand-600 font-medium">
                  Нажмите на значок
                  <svg class="w-3.5 h-3.5 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"/></svg>
                  <span class="animate-hint-arrow absolute -right-6 top-0 text-brand-500 font-bold">→</span>
                </span>
              </span>
            </p>
          </div>

          <div>
            <label class="block text-sm font-medium text-stone-700 mb-2">Пароль <span class="text-red-400">*</span></label>
            <input v-model="form.password" type="password" minlength="6" required
              class="w-full border border-stone-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900"
              placeholder="Минимум 6 символов" />
          </div>

          <div>
            <label class="block text-sm font-medium text-stone-700 mb-2">Подтверждение пароля <span class="text-red-400">*</span></label>
            <input v-model="form.confirm_password" type="password" required
              class="w-full border border-stone-200 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900"
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

        <p class="mt-8 text-center text-stone-400 text-sm">
          Уже есть аккаунт?
          <router-link to="/login" class="text-brand-600 hover:text-brand-700 font-semibold transition-colors"> Войдите</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const phone = ref('')
const form = reactive({
  first_name: '',
  last_name: '',
  middle_name: '',
  delivery_address: '',
  password: '',
  confirm_password: ''
})
const error = ref('')
const loading = ref(false)
const locating = ref(false)
const locationError = ref('')
const addressError = ref('')
const geoBtn = ref(null)

async function detectLocation() {
  if (!navigator.geolocation) {
    locationError.value = 'Геолокация не поддерживается вашим браузером'
    return
  }
  locating.value = true
  locationError.value = ''
  navigator.geolocation.getCurrentPosition(
    async (pos) => {
      try {
        const { latitude, longitude } = pos.coords
        const res = await fetch(
          `https://nominatim.openstreetmap.org/reverse?lat=${latitude}&lon=${longitude}&format=json&accept-language=ru`,
          { headers: { 'Accept-Language': 'ru' } }
        )
        const data = await res.json()
        if (data && data.address) {
          const a = data.address
          const parts = []
          if (a.state) parts.push(a.state)
          if (a.county || a.district) parts.push(a.county || a.district)
          if (a.city || a.town || a.village) parts.push(a.city || a.town || a.village)
          if (a.road) parts.push(a.road)
          form.delivery_address = parts.join(', ') || data.display_name
        }
      } catch {
        locationError.value = 'Не удалось определить адрес по координатам'
      } finally {
        locating.value = false
      }
    },
    (err) => {
      locating.value = false
      if (err.code === 1) {
        locationError.value = 'Доступ к геолокации запрещён'
      } else {
        locationError.value = 'Не удалось получить местоположение'
      }
    },
    { timeout: 10000 }
  )
}

onMounted(() => {
  // Auto-request location on page load
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
      async (pos) => {
        try {
          const { latitude, longitude } = pos.coords
          const res = await fetch(
            `https://nominatim.openstreetmap.org/reverse?lat=${latitude}&lon=${longitude}&format=json`,
            { headers: { 'Accept-Language': 'ru' } }
          )
          const data = await res.json()
          if (data && data.address && !form.delivery_address) {
            const a = data.address
            const parts = []
            if (a.state) parts.push(a.state)
            if (a.county || a.district) parts.push(a.county || a.district)
            if (a.city || a.town || a.village) parts.push(a.city || a.town || a.village)
            if (a.road) parts.push(a.road)
            form.delivery_address = parts.join(', ') || data.display_name
          }
        } catch { /* silently ignore */ }
      },
      () => { /* silently ignore denial */ },
      { timeout: 8000 }
    )
  }
})

async function handleRegister() {
  error.value = ''
  addressError.value = ''

  if (form.password !== form.confirm_password) {
    error.value = 'Пароли не совпадают'
    return
  }

  if (phone.value.length !== 9) {
    error.value = 'Введите 9 цифр номера телефона'
    return
  }

  if (!form.delivery_address.trim()) {
    addressError.value = 'Укажите адрес доставки. Нажмите на значок геолокации для автоматического определения.'
    geoBtn.value?.scrollIntoView({ behavior: 'smooth', block: 'center' })
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

<style scoped>
@keyframes hint-arrow {
  0%, 100% { transform: translateX(0); opacity: 1; }
  50% { transform: translateX(5px); opacity: 0.5; }
}
.animate-hint-arrow {
  animation: hint-arrow 1s ease-in-out infinite;
}
</style>
