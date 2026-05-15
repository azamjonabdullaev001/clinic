<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center justify-between h-16">
        <div class="flex items-center gap-3">
          <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <div>
            <h1 class="text-lg font-bold text-gray-800">Пункт выдачи</h1>
            <p v-if="authStore.worker" class="text-xs text-gray-400">{{ authStore.worker.name }}</p>
          </div>
        </div>
        <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">Выйти</button>
      </div>
    </header>

    <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 mt-6 pb-12 space-y-6">

      <!-- Search by code -->
      <div class="bg-white rounded-xl shadow-sm p-6">
        <h2 class="text-lg font-bold text-gray-800 mb-4">Поиск заказа по коду</h2>
        <div class="flex gap-3">
          <input
            v-model="searchCode"
            type="text"
            maxlength="6"
            placeholder="Введите 6-значный код"
            class="flex-1 border border-gray-300 rounded-lg px-4 py-3 text-2xl font-bold tracking-widest text-center focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
            @keyup.enter="searchByCode"
          />
          <button
            @click="searchByCode"
            :disabled="searchCode.length < 6 || searching"
            class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition font-medium disabled:opacity-40"
          >
            {{ searching ? 'Поиск...' : 'Найти' }}
          </button>
        </div>
        <p v-if="searchError" class="mt-3 text-red-500 text-sm">{{ searchError }}</p>

        <!-- Found order -->
        <div v-if="foundOrder" class="mt-6 border-2 border-blue-200 rounded-xl p-5 bg-blue-50">
          <div class="flex items-center justify-between mb-4">
            <div>
              <div class="flex items-center gap-2 mb-1">
                <span class="text-2xl font-bold text-blue-700 tracking-widest">{{ foundOrder.order_code }}</span>
                <span :class="statusClass(foundOrder.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(foundOrder.status) }}</span>
              </div>
              <p class="font-semibold text-gray-800">{{ foundOrder.user?.first_name }} {{ foundOrder.user?.last_name }}</p>
              <p v-if="foundOrder.user?.middle_name" class="text-sm text-gray-600">{{ foundOrder.user.middle_name }}</p>
              <p class="text-sm text-gray-500">+{{ foundOrder.phone }}</p>
              <p class="text-xs text-gray-400 mt-0.5">{{ new Date(foundOrder.created_at).toLocaleString('ru-RU') }}</p>
              <!-- Delivery address -->
              <div v-if="foundOrder.delivery_address" class="mt-2 flex items-start gap-1.5">
                <svg class="w-4 h-4 text-blue-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                </svg>
                <span class="text-sm text-blue-700 font-medium">{{ foundOrder.delivery_address }}</span>
              </div>
            </div>
            <div class="text-right">
              <p class="text-xs text-gray-400">Итого</p>
              <p class="text-2xl font-bold text-blue-700">{{ formatPrice(orderTotal(foundOrder)) }} <span class="text-sm font-normal">сўм</span></p>
            </div>
          </div>

          <div class="border-t border-blue-200 pt-4 space-y-2 mb-4">
            <div v-for="item in foundOrder.items" :key="item.id" class="flex justify-between items-center py-1.5 border-b border-blue-100 last:border-0">
              <div>
                <span class="font-medium text-gray-800">{{ item.product?.name }}</span>
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} {{ item.unit_type === 'piece' ? 'шт' : 'упак.' }}</span>
              </div>
              <span class="font-semibold text-gray-700">{{ formatPrice(item.price) }} сўм</span>
            </div>
          </div>

          <div class="flex gap-2 flex-wrap">
            <button
              v-if="foundOrder.status === 'pending'"
              @click="updateStatus(foundOrder, 'confirmed')"
              class="flex-1 bg-blue-600 text-white py-2.5 rounded-lg hover:bg-blue-700 transition font-medium text-sm"
            >
              ✓ Подтвердить
            </button>
            <button
              v-if="foundOrder.status === 'confirmed' || foundOrder.status === 'shipped'"
              @click="updateStatus(foundOrder, 'in_transit')"
              class="flex-1 bg-orange-500 text-white py-2.5 rounded-lg hover:bg-orange-600 transition font-medium text-sm"
            >
              🚚 Передан в доставку
            </button>
            <button
              v-if="foundOrder.status !== 'delivered' && foundOrder.status !== 'cancelled'"
              @click="updateStatus(foundOrder, 'delivered')"
              class="flex-1 bg-green-600 text-white py-2.5 rounded-lg hover:bg-green-700 transition font-medium text-sm"
            >
              ✓ Выдать заказ
            </button>
            <button
              v-if="foundOrder.status !== 'cancelled' && foundOrder.status !== 'delivered'"
              @click="updateStatus(foundOrder, 'cancelled')"
              class="flex-1 bg-red-50 text-red-600 border border-red-200 py-2.5 rounded-lg hover:bg-red-100 transition font-medium text-sm"
            >
              Отменить
            </button>
            <button
              @click="openChat(foundOrder)"
              class="bg-indigo-50 text-indigo-600 border border-indigo-200 py-2.5 px-4 rounded-lg hover:bg-indigo-100 transition font-medium text-sm flex items-center gap-1.5"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
              </svg>
              Написать
            </button>
          </div>
        </div>
      </div>

      <!-- All orders list -->
      <div>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-bold text-gray-800">Все заказы</h2>
          <button @click="loadOrders" class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            Обновить
          </button>
        </div>

        <!-- Filter tabs -->
        <div class="flex gap-2 mb-4 flex-wrap">
          <button
            v-for="f in filters"
            :key="f.value"
            @click="activeFilter = f.value"
            :class="activeFilter === f.value ? 'bg-blue-600 text-white' : 'bg-white text-gray-600 hover:bg-gray-50'"
            class="px-4 py-2 rounded-lg text-sm font-medium border border-gray-200 transition"
          >
            {{ f.label }}
            <span v-if="f.value !== 'all'" class="ml-1.5 text-xs opacity-75">({{ orderCountByStatus(f.value) }})</span>
          </button>
        </div>

        <div class="space-y-3">
          <div
            v-for="order in filteredOrders"
            :key="order.id"
            class="bg-white rounded-xl shadow-sm p-5 hover:shadow-md transition-shadow"
          >
            <div class="flex flex-col sm:flex-row justify-between items-start gap-3">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-xl font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                </div>
                <p class="font-semibold text-gray-800">{{ order.user?.first_name }} {{ order.user?.last_name }}
                  <span v-if="order.user?.middle_name" class="font-normal text-gray-600"> {{ order.user.middle_name }}</span>
                </p>
                <p class="text-sm text-gray-500">+{{ order.phone }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                <!-- Delivery address -->
                <div v-if="order.delivery_address" class="mt-1.5 flex items-start gap-1.5">
                  <svg class="w-3.5 h-3.5 text-orange-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                  </svg>
                  <span class="text-xs text-orange-700">{{ order.delivery_address }}</span>
                </div>
              </div>
              <div class="text-right flex-shrink-0">
                <p class="font-bold text-gray-800">{{ formatPrice(orderTotal(order)) }} сўм</p>
                <p class="text-xs text-gray-400">{{ order.items?.length }} позиций</p>
              </div>
            </div>

            <!-- Items summary -->
            <div class="mt-3 pt-3 border-t border-gray-100 space-y-1">
              <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm text-gray-600">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? 'шт' : 'упак.' }}</span></span>
                <span class="font-medium">{{ formatPrice(item.price) }} сўм</span>
              </div>
            </div>

            <!-- Actions -->
            <div class="mt-3 flex gap-2 flex-wrap">
              <template v-if="order.status !== 'delivered' && order.status !== 'cancelled'">
                <button
                  v-if="order.status === 'pending'"
                  @click="updateStatus(order, 'confirmed')"
                  class="bg-blue-600 text-white px-4 py-1.5 rounded-lg hover:bg-blue-700 transition text-sm font-medium"
                >
                  Подтвердить
                </button>
                <button
                  v-if="order.status === 'confirmed' || order.status === 'shipped'"
                  @click="updateStatus(order, 'in_transit')"
                  class="bg-orange-500 text-white px-4 py-1.5 rounded-lg hover:bg-orange-600 transition text-sm font-medium"
                >
                  🚚 В пути
                </button>
                <button
                  @click="updateStatus(order, 'delivered')"
                  class="bg-green-600 text-white px-4 py-1.5 rounded-lg hover:bg-green-700 transition text-sm font-medium"
                >
                  ✓ Выдать
                </button>
                <button
                  @click="updateStatus(order, 'cancelled')"
                  class="bg-red-50 text-red-600 border border-red-200 px-4 py-1.5 rounded-lg hover:bg-red-100 transition text-sm font-medium"
                >
                  Отменить
                </button>
              </template>
              <button
                @click="openChat(order)"
                class="bg-indigo-50 text-indigo-600 border border-indigo-200 px-3 py-1.5 rounded-lg hover:bg-indigo-100 transition text-sm font-medium flex items-center gap-1"
              >
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
                </svg>
                Написать
              </button>
            </div>
          </div>

          <div v-if="filteredOrders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            Нет заказов
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Chat panel -->
  <div v-if="chatOpen" class="fixed inset-0 z-50 flex items-end sm:items-center justify-center sm:justify-end" @click.self="chatOpen = false">
    <div class="w-full sm:w-96 h-[55vh] sm:h-[70vh] bg-white sm:rounded-tl-2xl sm:rounded-bl-2xl shadow-2xl flex flex-col sm:mr-0 sm:mt-0 rounded-t-2xl">
      <!-- Chat header -->
      <div class="flex items-center justify-between px-4 py-3 border-b bg-indigo-600 rounded-t-2xl sm:rounded-tl-2xl">
        <div>
          <p class="text-white font-semibold text-sm">{{ chatUserName }}</p>
          <p class="text-indigo-200 text-xs">Чат с клиентом</p>
        </div>
        <button @click="chatOpen = false" class="text-white/70 hover:text-white transition">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <!-- Messages -->
      <div ref="chatMessagesEl" class="flex-1 overflow-y-auto p-4 space-y-2 bg-gray-50">
        <div v-if="chatLoading" class="flex justify-center pt-8 text-gray-400 text-sm">Загрузка...</div>
        <template v-else>
          <div v-if="chatMessages.length === 0" class="text-center text-gray-400 text-sm pt-8">Нет сообщений</div>
          <div
            v-for="msg in chatMessages"
            :key="msg.id"
            class="flex"
            :class="msg.sender_role === 'user' ? 'justify-start' : 'justify-end'"
          >
            <div
              class="max-w-[75%] px-3 py-2 rounded-xl text-sm"
              :class="msg.sender_role === 'user'
                ? 'bg-white text-gray-800 shadow-sm'
                : 'bg-indigo-600 text-white'"
            >
              <p class="text-[10px] mb-0.5 opacity-60">{{ msg.sender_role === 'user' ? 'Клиент' : msg.sender_role === 'worker' ? 'Работник' : 'Администратор' }}</p>
              {{ msg.message }}
              <p class="text-[10px] mt-0.5 opacity-50 text-right">{{ new Date(msg.created_at).toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }) }}</p>
            </div>
          </div>
        </template>
      </div>
      <!-- Input -->
      <div class="px-3 py-3 border-t flex gap-2">
        <input
          v-model="chatMsg"
          @keyup.enter="sendWorkerMessage"
          placeholder="Введите сообщение..."
          class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <button
          @click="sendWorkerMessage"
          :disabled="!chatMsg.trim() || chatSending"
          class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition disabled:opacity-40"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const orders = ref([])
const searchCode = ref('')
const foundOrder = ref(null)
const searchError = ref('')
const searching = ref(false)
const activeFilter = ref('all')

const filters = [
  { label: 'Все', value: 'all' },
  { label: 'Ожидает', value: 'pending' },
  { label: 'Подтверждён', value: 'confirmed' },
  { label: 'Отправлен', value: 'shipped' },
  { label: 'В пути', value: 'in_transit' },
  { label: 'Выдан', value: 'delivered' },
  { label: 'Отменён', value: 'cancelled' },
]

const filteredOrders = computed(() => {
  if (activeFilter.value === 'all') return orders.value
  return orders.value.filter(o => o.status === activeFilter.value)
})

function orderCountByStatus(status) {
  return orders.value.filter(o => o.status === status).length
}

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

function statusLabel(status) {
  const labels = {
    pending: 'Ожидает',
    confirmed: 'Подтверждён',
    shipped: 'Отправлен',
    in_transit: 'В пути',
    delivered: 'Выдан',
    cancelled: 'Отменён',
  }
  return labels[status] || status
}

function statusClass(status) {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-700',
    confirmed: 'bg-blue-100 text-blue-700',
    shipped: 'bg-purple-100 text-purple-700',
    in_transit: 'bg-orange-100 text-orange-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-700',
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

async function loadOrders() {
  try {
    const res = await api.get('/pickup/orders')
    orders.value = res.data || []
  } catch (e) {
    console.error(e)
  }
}

async function searchByCode() {
  if (searchCode.value.length < 6) return
  searchError.value = ''
  foundOrder.value = null
  searching.value = true
  try {
    const res = await api.get(`/pickup/orders/code/${searchCode.value}`)
    foundOrder.value = res.data
  } catch (e) {
    searchError.value = e.response?.data?.error || 'Заказ не найден'
  } finally {
    searching.value = false
  }
}

async function updateStatus(order, status) {
  try {
    const res = await api.put(`/pickup/orders/${order.id}/status`, { status })
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx !== -1) orders.value[idx] = res.data
    if (foundOrder.value?.id === order.id) foundOrder.value = res.data
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при обновлении статуса')
  }
}

function logout() {
  authStore.workerLogout()
  router.push('/admin/login')
}

// Chat
const chatOpen = ref(false)
const chatUserName = ref('')
const chatThreadId = ref(null)
const chatMessages = ref([])
const chatMsg = ref('')
const chatSending = ref(false)
const chatLoading = ref(false)
const chatMessagesEl = ref(null)

async function openChat(order) {
  chatUserName.value = [order.user?.first_name, order.user?.last_name].filter(Boolean).join(' ') || order.phone
  chatOpen.value = true
  chatMessages.value = []
  chatThreadId.value = null
  chatLoading.value = true
  try {
    const res = await api.get('/pickup/support/threads')
    const threads = res.data || []
    const thread = threads.find(t => t.user_id === order.user_id || t.user?.id === order.user_id)
    if (thread) {
      chatThreadId.value = thread.id
      const detail = await api.get(`/pickup/support/threads/${thread.id}`)
      chatMessages.value = detail.data.messages || []
    }
  } catch (e) {
    console.error(e)
  } finally {
    chatLoading.value = false
    await nextTick()
    scrollChatToBottom()
  }
}

function scrollChatToBottom() {
  if (chatMessagesEl.value) {
    chatMessagesEl.value.scrollTop = chatMessagesEl.value.scrollHeight
  }
}

async function sendWorkerMessage() {
  if (!chatMsg.value.trim() || chatSending.value) return
  const text = chatMsg.value.trim()

  if (!chatThreadId.value) {
    alert('Пользователь ещё не начал переписку')
    return
  }

  chatSending.value = true
  try {
    const res = await api.post(`/pickup/support/threads/${chatThreadId.value}/reply`, { message: text })
    chatMessages.value.push(res.data)
    chatMsg.value = ''
    await nextTick()
    scrollChatToBottom()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при отправке')
  } finally {
    chatSending.value = false
  }
}

onMounted(() => {
  loadOrders()
})
</script>
