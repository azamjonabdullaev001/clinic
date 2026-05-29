<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center justify-between h-16">
        <div class="flex items-center gap-3">
          <div class="w-8 h-8 bg-teal-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.066 2.573c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.573 1.066c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.066-2.573c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
          <h1 class="text-xl font-bold text-gray-800">Админ-панель</h1>
        </div>
        <div class="flex items-center gap-4">
          <router-link to="/" class="text-sm text-gray-500 hover:text-teal-600 transition">← На сайт</router-link>
          <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">Выйти</button>
        </div>
      </div>
    </header>

    <!-- Tabs -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-6">
      <div class="flex space-x-1 bg-white rounded-xl p-1.5 shadow-sm">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="activeTab === tab.id
            ? 'bg-teal-600 text-white shadow-sm'
            : 'text-gray-600 hover:bg-gray-100'"
          class="flex-1 py-2.5 px-4 rounded-lg font-medium text-sm transition"
        >
          {{ tab.label }}
        </button>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-6 pb-12">

      <!-- ===== Products Tab ===== -->
      <div v-if="activeTab === 'products'">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Препараты</h2>
          <button @click="openProductModal()" class="bg-teal-600 text-white px-5 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
            Добавить
          </button>
        </div>

        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50 border-b">
                <tr>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Фото</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Название</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">В капсуле</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">На складе</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Цена/шт</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Цена/капс.</th>
                  <th class="text-right px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Действия</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="product in products" :key="product.id" class="hover:bg-gray-50 transition">
                  <td class="px-5 py-3">
                    <div class="w-12 h-12 bg-gray-100 rounded-lg overflow-hidden flex items-center justify-center">
                      <img v-if="product.image_path" :src="product.image_path" class="w-full h-full object-cover" />
                      <svg v-else class="w-6 h-6 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </td>
                  <td class="px-5 py-3">
                    <div class="font-medium text-gray-800">{{ product.name }}</div>
                    <div v-if="product.description" class="text-xs text-gray-400 truncate max-w-xs mt-0.5">{{ product.description }}</div>
                  </td>
                  <td class="px-5 py-3 text-gray-600">{{ product.quantity_per_pack }} шт</td>
                  <td class="px-5 py-3">
                    <span :class="liveQty(product) > 0 ? 'text-green-700 bg-green-50' : 'text-red-600 bg-red-50'" class="px-2 py-0.5 rounded text-sm font-semibold">
                      {{ liveCaps(product) }} капс / {{ liveQty(product) }} шт
                    </span>
                  </td>
                  <td class="px-5 py-3 text-gray-600">{{ formatPrice(product.price_per_pill) }} сўм</td>
                  <td class="px-5 py-3 font-semibold text-teal-600">{{ formatPrice(product.price_per_pack) }} сўм</td>
                  <td class="px-5 py-3 text-right">
                    <div class="flex items-center justify-end gap-1">
                      <button @click="uploadImageFor(product)" class="p-2 text-blue-500 hover:bg-blue-50 rounded-lg transition" title="Загрузить фото">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" /></svg>
                      </button>
                      <button @click="openProductModal(product)" class="p-2 text-teal-500 hover:bg-teal-50 rounded-lg transition" title="Редактировать">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                      </button>
                      <button @click="deleteProduct(product.id)" class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition" title="Удалить">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="products.length === 0">
                  <td colspan="6" class="px-5 py-12 text-center text-gray-400">
                    <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                    Нет добавленных препаратов
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- ===== Orders Tab ===== -->
      <div v-if="activeTab === 'orders'">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Заказы</h2>
          <button v-if="orders.length" @click="deleteAllOrders"
            class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg font-medium transition-colors">
            Удалить все заказы
          </button>
        </div>
        <div class="space-y-4">
          <div v-for="order in orders" :key="order.id" class="bg-white rounded-xl shadow-sm p-6"
            :class="order.is_returned ? 'border-2 border-red-400' : ''">
            <div class="flex flex-col sm:flex-row justify-between items-start gap-4 mb-4">
              <div>
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-xs font-medium text-gray-400 bg-gray-100 px-2 py-0.5 rounded">#{{ order.id }}</span>
                  <span v-if="order.order_code" class="text-sm font-bold text-blue-700 tracking-widest bg-blue-50 px-2 py-0.5 rounded">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">
                    {{ statusLabel(order.status) }}
                  </span>
                  <span v-if="order.is_offline" class="text-xs font-semibold px-2 py-0.5 rounded bg-emerald-100 text-emerald-700">Офлайн</span>
                  <span v-if="paymentLabel(order)" :class="paymentBadgeClass(order)" class="text-xs font-semibold px-2 py-0.5 rounded">{{ paymentLabel(order) }}</span>
                  <span v-if="order.is_returned" class="text-xs font-semibold px-2 py-0.5 rounded bg-red-100 text-red-600">Возврат</span>
                </div>
                <h3 class="font-semibold text-gray-800 text-lg">
                  <template v-if="order.is_offline">
                    {{ order.offline_note || 'Покупатель (офлайн)' }}
                  </template>
                  <template v-else>
                    {{ order.user?.first_name }} {{ order.user?.last_name }}
                    <span v-if="order.user?.middle_name" class="font-normal text-gray-600 text-base"> {{ order.user.middle_name }}</span>
                  </template>
                </h3>
                <p v-if="!order.is_offline" class="text-sm text-gray-500">+{{ order.phone }}</p>
                <p class="text-xs text-gray-400 mt-1">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                <div v-if="order.delivery_address" class="mt-1.5 flex items-start gap-1.5">
                  <svg class="w-3.5 h-3.5 text-teal-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                  </svg>
                  <span class="text-xs text-teal-700">{{ order.delivery_address }}</span>
                </div>
                <div v-if="order.latitude && order.longitude" class="mt-1 flex items-center gap-2 flex-wrap">
                  <a
                    :href="`https://www.openstreetmap.org/?mlat=${order.latitude}&mlon=${order.longitude}#map=15/${order.latitude}/${order.longitude}`"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="inline-flex items-center gap-1 text-xs text-blue-600 hover:text-blue-800 font-medium transition-colors"
                  >
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                    На карте
                  </a>
                  <button
                    @click="openRouteModal(order)"
                    class="inline-flex items-center gap-1 text-xs text-emerald-600 hover:text-emerald-800 font-medium transition-colors"
                  >
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
                    Маршрут
                  </button>
                </div>
                <div v-if="order.referred_by" class="mt-1 flex items-center gap-1">
                  <svg class="w-3.5 h-3.5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                  <span class="text-xs text-purple-700">Рекомендовал: {{ order.referred_by }}</span>
                </div>
              </div>
              <div class="flex items-center gap-2">
                <span :class="statusClass(order.status)" class="text-xs font-semibold px-3 py-1.5 rounded-full">
                  {{ statusLabel(order.status) }}
                </span>
                <button
                  @click="deleteOrder(order.id)"
                  class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition"
                  title="Удалить заказ"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6M1 7h22M9 7V4a1 1 0 011-1h4a1 1 0 011 1v3" />
                  </svg>
                </button>
              </div>
            </div>

            <div v-if="order.is_returned && order.return_reason" class="mb-4 bg-red-50 border border-red-200 rounded-lg px-3 py-2 flex items-start gap-2">
              <svg class="w-4 h-4 text-red-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"/></svg>
              <div class="text-sm"><span class="font-semibold text-red-700">Причина возврата:</span><span class="text-red-600 ml-1">{{ order.return_reason }}</span></div>
            </div>

            <div v-if="order.status === 'cancelled' && (order.cancellation_reason || order.cancelled_by_name)" class="mb-4 bg-red-50 border border-red-200 rounded-lg px-3 py-2 space-y-1">
              <div v-if="order.cancellation_reason" class="flex items-start gap-2">
                <svg class="w-4 h-4 text-red-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div class="text-sm">
                  <span class="font-semibold text-red-700">Причина отмены:</span>
                  <span class="text-red-600 ml-1">{{ order.cancellation_reason }}</span>
                </div>
              </div>
              <div v-if="order.cancelled_by_name" class="flex items-start gap-2">
                <svg class="w-4 h-4 text-red-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                <div class="text-sm">
                  <span class="font-semibold text-red-700">Отменил:</span>
                  <span class="text-red-600 ml-1">{{ order.cancelled_by_name }}<span v-if="order.cancelled_by_role" class="text-red-400"> ({{ order.cancelled_by_role === 'nurse' ? 'медсестра' : 'пункт выдачи' }})</span></span>
                </div>
              </div>
            </div>

            <div class="border-t pt-4 space-y-2">
              <div v-for="item in order.items" :key="item.id" class="flex justify-between items-center text-sm py-1">
                <span class="text-gray-600">
                  {{ item.product?.name }}
                  <span class="text-gray-400" :class="{ 'line-through': item.quantity === 0 }">× {{ item.quantity === 0 ? item.original_quantity : item.quantity }} капс.</span>
                  <span v-if="itemDiffLabel(item)" :class="itemDiffClass(item)" class="ml-1 text-[11px] font-semibold px-1.5 py-0.5 rounded">{{ itemDiffLabel(item) }}</span>
                </span>
                <span class="font-medium text-gray-700">{{ item.quantity === 0 ? '—' : formatPrice(item.price) + ' сўм' }}</span>
              </div>
              <div class="flex justify-between font-bold text-base pt-3 border-t mt-2">
                <span>Итого:</span>
                <span class="text-teal-600">{{ formatPrice(orderTotal(order)) }} сўм</span>
              </div>
            </div>
          </div>

          <div v-if="orders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            Нет заказов
          </div>
        </div>
      </div>

      <!-- ===== Workers Tab ===== -->
      <div v-if="activeTab === 'workers'">
        <div class="flex flex-wrap justify-between items-center gap-3 mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Работники</h2>
          <div class="flex items-center gap-3">
            <router-link to="/admin/workers"
              class="flex items-center gap-2 border border-gray-300 text-gray-600 px-4 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium text-sm">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/>
              </svg>
              Открыть панель работников
            </router-link>
            <button @click="openWorkerModal()" class="bg-teal-600 text-white px-5 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium flex items-center gap-2">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
              Добавить работника
            </button>
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <table class="w-full">
            <thead class="bg-gray-50 border-b">
              <tr>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Имя</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Телефон</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Роль</th>
                <th class="text-right px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Действия</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-100">
              <template v-for="w in workers" :key="w.id">
                <tr class="hover:bg-gray-50 transition">
                  <td class="px-5 py-3 font-medium text-gray-800">{{ w.name }}</td>
                  <td class="px-5 py-3 text-gray-500 font-mono">+{{ w.phone }}</td>
                  <td class="px-5 py-3">
                    <span :class="w.role === 'nurse' ? 'bg-teal-50 text-teal-700' : w.role === 'manager' ? 'bg-purple-50 text-purple-700' : 'bg-blue-50 text-blue-700'"
                      class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium">
                      {{ w.role === 'nurse' ? 'Медсестра' : w.role === 'manager' ? 'Менеджер' : 'Пункт выдачи' }}
                    </span>
                  </td>
                  <td class="px-5 py-3 text-right">
                    <div class="flex items-center justify-end gap-1">
                      <button @click="toggleWorkerStats(w)" class="p-2 rounded-lg transition"
                        :class="expandedWorkerId === w.id ? 'text-emerald-600 bg-emerald-50' : 'text-gray-400 hover:bg-gray-100'"
                        title="Статистика">
                        <svg class="w-4 h-4 transition-transform duration-200" :class="expandedWorkerId === w.id ? 'rotate-180' : ''" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/></svg>
                      </button>
                      <button @click="openEditWorkerModal(w)" class="p-2 text-blue-500 hover:bg-blue-50 rounded-lg transition" title="Редактировать">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                      </button>
                      <button @click="deleteWorker(w.id)" class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition" title="Удалить">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="expandedWorkerId === w.id" class="bg-emerald-50/40">
                  <td colspan="4" class="px-5 py-4">
                    <div v-if="workerStatsLoading" class="flex items-center gap-2 text-sm text-gray-400 py-2">
                      <div class="w-4 h-4 border-2 border-emerald-400 border-t-transparent rounded-full animate-spin"></div>
                      Загрузка статистики...
                    </div>
                    <div v-else-if="workerStats" class="grid grid-cols-1 md:grid-cols-3 gap-3">
                      <div v-for="(period, key) in { today: 'Сегодня', week: 'За 7 дней', month: 'За 30 дней' }" :key="key"
                        class="bg-white rounded-xl border border-emerald-100 p-4">
                        <div class="text-xs font-semibold text-emerald-700 uppercase tracking-wide mb-2">{{ period }}</div>
                        <div class="space-y-1.5">
                          <div class="flex justify-between text-sm">
                            <span class="text-gray-500">Заказов:</span>
                            <span class="font-bold text-gray-800">{{ workerStats[key]?.orders || 0 }}</span>
                          </div>
                          <div class="flex justify-between text-sm">
                            <span class="text-gray-500">Товаров:</span>
                            <span class="font-bold text-gray-800">{{ workerStats[key]?.items || 0 }} шт</span>
                          </div>
                          <div class="flex justify-between text-sm pt-1.5 border-t border-emerald-50">
                            <span class="text-gray-500">Выручка:</span>
                            <span class="font-bold text-emerald-700">{{ formatPrice(workerStats[key]?.revenue || 0) }} сўм</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </td>
                </tr>
              </template>
              <tr v-if="workers.length === 0">
                <td colspan="4" class="px-5 py-12 text-center text-gray-400">
                  <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  Нет работников
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- ===== Doctors Tab ===== -->
      <div v-if="activeTab === 'doctors'">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Доктора</h2>
          <button @click="showDoctorModal = true" class="bg-purple-600 text-white px-5 py-2.5 rounded-lg hover:bg-purple-700 transition font-medium flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
            Добавить доктора
          </button>
        </div>
        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <table class="w-full">
            <thead class="bg-gray-50 border-b">
              <tr>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Имя доктора</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Специализация</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Телефон</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Добавлен</th>
                <th class="text-right px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Действия</th>
              </tr>
            </thead>
            <tbody>
              <template v-for="doc in doctors" :key="doc.id">
                <tr class="hover:bg-gray-50 transition border-b border-gray-100">
                  <td class="px-5 py-3 font-medium text-gray-800">
                    <div class="flex items-center gap-2">
                      <svg class="w-4 h-4 text-purple-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                      {{ doc.name }}
                    </div>
                  </td>
                  <td class="px-5 py-3">
                    <span v-if="doc.specialty" class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-medium bg-purple-50 text-purple-700">{{ doc.specialty }}</span>
                    <span v-else class="text-gray-400 text-sm">—</span>
                  </td>
                  <td class="px-5 py-3 text-gray-500 text-sm">
                    <span v-if="doc.phone">+{{ doc.phone }}</span>
                    <span v-else class="text-gray-300">—</span>
                  </td>
                  <td class="px-5 py-3 text-gray-500 text-sm">{{ new Date(doc.created_at).toLocaleDateString('ru-RU') }}</td>
                  <td class="px-5 py-3 text-right">
                    <div class="flex items-center justify-end gap-1">
                      <button @click="toggleDoctorStats(doc)" class="p-2 rounded-lg transition"
                        :class="expandedDoctorId === doc.id ? 'text-purple-600 bg-purple-50' : 'text-gray-400 hover:bg-gray-100'"
                        title="Статистика продаж">
                        <svg class="w-4 h-4 transition-transform duration-200" :class="expandedDoctorId === doc.id ? 'rotate-180' : ''" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/></svg>
                      </button>
                      <button @click="openEditDoctorModal(doc)" class="p-2 text-teal-500 hover:bg-teal-50 rounded-lg transition" title="Редактировать">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                      </button>
                      <button @click="deleteDoctor(doc.id)" class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition" title="Удалить">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                      </button>
                    </div>
                  </td>
                </tr>
                <!-- Stats expand row -->
                <tr v-if="expandedDoctorId === doc.id" class="bg-purple-50/40">
                  <td colspan="5" class="px-5 py-4">
                    <div v-if="doctorStatsLoading" class="flex items-center gap-2 text-sm text-gray-400 py-2">
                      <div class="w-4 h-4 border-2 border-purple-400 border-t-transparent rounded-full animate-spin"></div>
                      Загрузка статистики...
                    </div>
                    <div v-else-if="doctorStats">
                      <div class="flex items-center gap-3 mb-3">
                        <span class="text-sm font-semibold text-purple-800">{{ doctorStats.doctor.name }}</span>
                        <span class="text-xs bg-purple-100 text-purple-700 px-2 py-0.5 rounded-full font-medium">{{ doctorStats.total_orders }} заказ(ов)</span>
                      </div>
                      <div v-if="doctorStats.products && doctorStats.products.length > 0" class="overflow-x-auto rounded-xl border border-purple-100">
                        <table class="w-full text-sm">
                          <thead class="bg-purple-100/60">
                            <tr>
                              <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Препарат</th>
                              <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Капсул</th>
                              <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Штук</th>
                              <th class="text-right px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Выручка</th>
                            </tr>
                          </thead>
                          <tbody class="divide-y divide-purple-100">
                            <tr v-for="p in doctorStats.products" :key="p.product_id" class="hover:bg-purple-50 transition">
                              <td class="px-4 py-2 font-medium text-gray-800">{{ p.product_name }}</td>
                              <td class="px-4 py-2 text-gray-600">{{ p.total_packs }} капс.</td>
                              <td class="px-4 py-2 text-gray-600">{{ p.total_pieces }} шт</td>
                              <td class="px-4 py-2 text-right font-bold text-purple-700">{{ formatPrice(p.revenue) }} сўм</td>
                            </tr>
                          </tbody>
                        </table>
                      </div>
                      <p v-else class="text-sm text-gray-400 py-2">Продаж по этому доктору ещё нет</p>
                    </div>
                  </td>
                </tr>
              </template>
              <tr v-if="doctors.length === 0">
                <td colspan="5" class="px-5 py-12 text-center text-gray-400">
                  <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  Нет добавленных докторов
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- ===== Analytics Tab ===== -->
      <div v-if="activeTab === 'analytics'">
        <h2 class="text-2xl font-bold text-gray-800 mb-6">Аналитика продаж</h2>

        <!-- Period selector -->
        <div class="bg-white rounded-xl shadow-sm p-4 mb-5 flex flex-wrap gap-3 items-center">
          <div class="flex gap-2 flex-wrap">
            <button
              v-for="p in analyticsPeriods"
              :key="p.value"
              @click="selectPeriod(p.value)"
              :class="analyticsPeriod === p.value ? 'bg-teal-600 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'"
              class="px-4 py-2 rounded-lg text-sm font-medium transition"
            >{{ p.label }}</button>
          </div>
          <div v-if="analyticsPeriod === 'custom'" class="flex items-center gap-2">
            <input
              v-model="analyticsCustomDate"
              type="date"
              class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
              @change="loadAnalytics"
            />
          </div>
          <button @click="loadAnalytics" class="ml-auto text-sm text-teal-600 hover:text-teal-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            Обновить
          </button>
        </div>

        <!-- Summary cards -->
        <div class="grid grid-cols-2 gap-4 mb-5" v-if="analyticsData">
          <div class="bg-white rounded-xl shadow-sm p-5">
            <p class="text-sm text-gray-500 mb-1">Выручка за период</p>
            <p class="text-3xl font-bold text-teal-600">{{ formatPrice(analyticsData.total_revenue) }} <span class="text-base font-normal text-gray-500">сўм</span></p>
          </div>
          <div class="bg-white rounded-xl shadow-sm p-5">
            <p class="text-sm text-gray-500 mb-1">Заказов за период</p>
            <p class="text-3xl font-bold text-blue-600">{{ analyticsData.total_orders }}</p>
          </div>
        </div>

        <!-- Chart -->
        <div class="bg-white rounded-xl shadow-sm p-6 mb-5">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">
            {{ analyticsPeriods.find(p => p.value === analyticsPeriod)?.label }} — динамика продаж
          </h3>
          <div v-show="analyticsLoading" class="flex items-center justify-center text-gray-400" style="height:280px">Загрузка...</div>
          <div v-show="!analyticsLoading" style="position:relative;height:280px;width:100%">
            <canvas ref="analyticsCanvas"></canvas>
          </div>
        </div>

        <!-- Top 10 products -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden" v-if="analyticsData && analyticsData.top_products?.length">
          <div class="px-6 py-4 border-b border-gray-100">
            <h3 class="text-lg font-semibold text-gray-800">Топ-10 товаров по выручке</h3>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50 border-b">
                <tr>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">#</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Препарат</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Капсул</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Штук</th>
                  <th class="text-right px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Выручка</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="(p, i) in analyticsData.top_products" :key="p.product_id" class="hover:bg-gray-50 transition">
                  <td class="px-5 py-3 text-gray-400 font-medium">{{ i + 1 }}</td>
                  <td class="px-5 py-3 font-medium text-gray-800">{{ p.product_name }}</td>
                  <td class="px-5 py-3 text-gray-600">{{ p.total_packs }} капс.</td>
                  <td class="px-5 py-3 text-gray-600">{{ p.total_qty }} шт</td>
                  <td class="px-5 py-3 text-right font-bold text-teal-600">{{ formatPrice(p.revenue) }} сўм</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div v-else-if="analyticsData && !analyticsData.top_products?.length" class="bg-white rounded-xl shadow-sm p-10 text-center text-gray-400">
          Нет данных о продажах за выбранный период
        </div>

        <!-- Category breakdown (separate, non-overlapping) -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5" v-if="analyticsData && analyticsData.breakdown">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center gap-2">
            <svg class="w-5 h-5 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>
            <h3 class="text-lg font-semibold text-gray-800">Разбивка по категориям</h3>
            <span class="text-xs text-gray-400">каждый заказ учтён один раз</span>
          </div>
          <div class="px-6 py-4 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3">
            <div v-for="cat in breakdownCats" :key="cat.key" class="rounded-xl p-4 border" :class="cat.box">
              <p class="text-sm font-semibold mb-2" :class="cat.text">{{ cat.label }}</p>
              <div class="flex justify-between text-xs text-gray-500"><span>Заказов</span><span class="font-bold text-gray-800">{{ analyticsData.breakdown[cat.key]?.orders || 0 }}</span></div>
              <div class="flex justify-between text-xs text-gray-500"><span>Капсул</span><span class="font-bold text-gray-800">{{ analyticsData.breakdown[cat.key]?.capsules || 0 }}</span></div>
              <div class="flex justify-between text-xs text-gray-500"><span>Штук</span><span class="font-bold text-gray-800">{{ analyticsData.breakdown[cat.key]?.pieces || 0 }}</span></div>
              <div class="flex justify-between text-xs text-gray-500 mt-1 pt-1 border-t border-gray-100"><span>Сумма</span><span class="font-bold" :class="cat.text">{{ formatPrice(analyticsData.breakdown[cat.key]?.revenue || 0) }}</span></div>
            </div>
          </div>
        </div>

        <!-- Own patients (свои пациенты) -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5" v-if="analyticsData && analyticsData.vip && analyticsData.vip.total_orders > 0">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center gap-2">
            <svg class="w-5 h-5 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/></svg>
            <h3 class="text-lg font-semibold text-gray-800">Свои пациенты</h3>
            <span class="text-xs text-gray-400">бесплатно, отдельно от выручки</span>
          </div>
          <div class="px-6 py-4">
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 mb-4">
              <div class="bg-emerald-50 rounded-xl p-4"><p class="text-xs text-emerald-600 mb-1">Стоимость</p><p class="text-xl font-bold text-emerald-700">{{ formatPrice(analyticsData.vip.total_revenue) }} сўм</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Заказов</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.vip.total_orders }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Капсул</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.vip.total_capsules }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Штук</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.vip.total_pieces }}</p></div>
            </div>
            <div v-if="analyticsData.vip.products?.length" class="overflow-x-auto rounded-xl border border-emerald-100">
              <table class="w-full text-sm">
                <thead class="bg-emerald-50">
                  <tr>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-emerald-700 uppercase">Препарат</th>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-emerald-700 uppercase">Капсул</th>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-emerald-700 uppercase">Штук</th>
                    <th class="text-right px-4 py-2.5 text-xs font-semibold text-emerald-700 uppercase">Стоимость</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-emerald-50">
                  <tr v-for="(p, i) in analyticsData.vip.products" :key="i" class="hover:bg-emerald-50 transition">
                    <td class="px-4 py-2.5 font-medium text-gray-800">{{ p.product_name }}</td>
                    <td class="px-4 py-2.5 text-gray-600">{{ p.capsules }}</td>
                    <td class="px-4 py-2.5 text-gray-600">{{ p.pieces }}</td>
                    <td class="px-4 py-2.5 text-right font-bold text-emerald-700">{{ formatPrice(p.revenue) }} сўм</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Marketolog (debt) -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5" v-if="analyticsData && analyticsData.marketolog && analyticsData.marketolog.total_orders > 0">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center gap-2">
            <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/></svg>
            <h3 class="text-lg font-semibold text-gray-800">Маркетолог (долг)</h3>
            <span class="text-xs text-gray-400">отдельно от выручки</span>
          </div>
          <div class="px-6 py-4">
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 mb-4">
              <div class="bg-purple-50 rounded-xl p-4"><p class="text-xs text-purple-600 mb-1">Сумма (долг)</p><p class="text-xl font-bold text-purple-700">{{ formatPrice(analyticsData.marketolog.total_revenue) }} сўм</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Заказов</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.marketolog.total_orders }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Капсул</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.marketolog.total_capsules }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Штук</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.marketolog.total_pieces }}</p></div>
            </div>
            <div v-if="analyticsData.marketolog.products?.length" class="overflow-x-auto rounded-xl border border-purple-100">
              <table class="w-full text-sm">
                <thead class="bg-purple-50">
                  <tr>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-purple-700 uppercase">Препарат</th>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-purple-700 uppercase">Капсул</th>
                    <th class="text-left px-4 py-2.5 text-xs font-semibold text-purple-700 uppercase">Штук</th>
                    <th class="text-right px-4 py-2.5 text-xs font-semibold text-purple-700 uppercase">Сумма</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-purple-50">
                  <tr v-for="(p, i) in analyticsData.marketolog.products" :key="i" class="hover:bg-purple-50 transition">
                    <td class="px-4 py-2.5 font-medium text-gray-800">{{ p.product_name }}</td>
                    <td class="px-4 py-2.5 text-gray-600">{{ p.capsules }}</td>
                    <td class="px-4 py-2.5 text-gray-600">{{ p.pieces }}</td>
                    <td class="px-4 py-2.5 text-right font-bold text-purple-700">{{ formatPrice(p.revenue) }} сўм</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Doctor Referral Stats -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5" v-if="analyticsData && analyticsData.doctor_referrals?.length">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center gap-2">
            <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
            <h3 class="text-lg font-semibold text-gray-800">Реферальная статистика</h3>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50 border-b">
                <tr>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">#</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Кто рекомендовал</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Заказов</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Капсул</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Штук</th>
                  <th class="text-right px-5 py-3 text-xs font-semibold text-gray-500 uppercase">Выручка</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="(d, i) in analyticsData.doctor_referrals" :key="d.doctor_name" class="hover:bg-gray-50 transition">
                  <td class="px-5 py-3 text-gray-400 font-medium">{{ i + 1 }}</td>
                  <td class="px-5 py-3 font-medium text-gray-800 flex items-center gap-2">
                    <svg class="w-4 h-4 text-purple-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                    {{ d.doctor_name }}
                  </td>
                  <td class="px-5 py-3 text-gray-600">{{ d.order_count }}</td>
                  <td class="px-5 py-3 text-gray-600">{{ d.capsules }}</td>
                  <td class="px-5 py-3 text-gray-600">{{ d.pieces }}</td>
                  <td class="px-5 py-3 text-right font-bold text-purple-600">{{ formatPrice(d.total_revenue) }} сўм</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Per-Marketolog Detail Stats -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between gap-2 flex-wrap">
            <div class="flex items-center gap-2">
              <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5"/></svg>
              <h3 class="text-lg font-semibold text-gray-800">Аналитика маркетолога</h3>
            </div>
            <select v-model="perMktId" @change="loadPerMktStats"
              class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500 min-w-[180px]">
              <option value="">Выберите маркетолога</option>
              <option v-for="m in marketologList" :key="m.id" :value="m.id">{{ m.name }}</option>
            </select>
          </div>
          <div class="px-6 py-4">
            <div v-if="perMktStats">
              <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 mb-4">
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Заказов</p><p class="text-2xl font-bold text-gray-800">{{ perMktStats.total_orders }}</p></div>
                <div class="bg-purple-50 rounded-xl p-4"><p class="text-xs text-purple-600 mb-1">Сумма (долг)</p><p class="text-lg font-bold text-purple-700">{{ formatPrice(perMktStats.total_revenue) }} сўм</p></div>
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Капсул</p><p class="text-2xl font-bold text-gray-800">{{ perMktStats.total_capsules }}</p></div>
                <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">Штук</p><p class="text-2xl font-bold text-gray-800">{{ perMktStats.total_pieces }}</p></div>
              </div>
              <div v-if="(perMktStats.products || []).length" class="overflow-x-auto rounded-xl border border-purple-100">
                <table class="w-full text-sm">
                  <thead class="bg-purple-50"><tr>
                    <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Препарат</th>
                    <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Капсул</th>
                    <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Штук</th>
                    <th class="text-right px-4 py-2 text-xs font-semibold text-purple-700 uppercase">Сумма</th>
                  </tr></thead>
                  <tbody class="divide-y divide-purple-50">
                    <tr v-for="(p, i) in perMktStats.products" :key="i">
                      <td class="px-4 py-2 font-medium text-gray-800">{{ p.product_name }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ p.capsules }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ p.pieces }}</td>
                      <td class="px-4 py-2 text-right font-bold text-purple-700">{{ formatPrice(p.revenue) }} сўм</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
            <div v-else class="py-6 text-center text-gray-400 text-sm">Выберите маркетолога</div>
          </div>
        </div>

        <!-- Per-Doctor Detail Stats -->
        <div class="bg-white rounded-xl shadow-sm overflow-hidden mt-5">
          <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between gap-2 flex-wrap">
            <div class="flex items-center gap-2">
              <svg class="w-5 h-5 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg>
              <h3 class="text-lg font-semibold text-gray-800">Детальная статистика по врачу</h3>
            </div>
            <div class="flex items-center gap-2 flex-wrap">
              <select v-model="perDoctorSelectedId"
                class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 min-w-[180px]">
                <option value="">Выберите врача</option>
                <option v-for="doc in doctors" :key="doc.id" :value="doc.id">{{ doc.name }}</option>
              </select>
              <button @click="loadPerDoctorStats"
                :disabled="!perDoctorSelectedId || perDoctorStatsLoading"
                class="bg-indigo-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-indigo-700 transition disabled:opacity-40 flex items-center gap-1.5">
                <div v-if="perDoctorStatsLoading" class="w-3.5 h-3.5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                <svg v-else class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg>
                Показать статистику
              </button>
            </div>
          </div>
          <div class="px-6 py-4">
            <div v-if="perDoctorStatsLoading" class="flex items-center gap-2 text-sm text-gray-400 py-4">
              <div class="w-4 h-4 border-2 border-indigo-400 border-t-transparent rounded-full animate-spin"></div>
              Загрузка статистики...
            </div>
            <div v-else-if="perDoctorStats">
              <div class="flex items-center gap-3 mb-4">
                <div class="w-8 h-8 bg-indigo-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                </div>
                <div>
                  <p class="font-semibold text-gray-800">{{ perDoctorStats.doctor.name }}</p>
                  <p v-if="perDoctorStats.doctor.specialty" class="text-xs text-gray-400">{{ perDoctorStats.doctor.specialty }}</p>
                </div>
                <span class="ml-auto text-sm bg-indigo-100 text-indigo-700 px-3 py-1 rounded-full font-medium">{{ perDoctorStats.total_orders }} заказ(ов)</span>
              </div>
              <div v-if="perDoctorStats.products && perDoctorStats.products.length > 0" class="overflow-x-auto rounded-xl border border-indigo-100">
                <table class="w-full text-sm">
                  <thead class="bg-indigo-50">
                    <tr>
                      <th class="text-left px-4 py-2.5 text-xs font-semibold text-indigo-700 uppercase">Препарат</th>
                      <th class="text-left px-4 py-2.5 text-xs font-semibold text-indigo-700 uppercase">Капсул</th>
                      <th class="text-left px-4 py-2.5 text-xs font-semibold text-indigo-700 uppercase">Штук</th>
                      <th class="text-left px-4 py-2.5 text-xs font-semibold text-indigo-700 uppercase">Заказов</th>
                      <th class="text-right px-4 py-2.5 text-xs font-semibold text-indigo-700 uppercase">Выручка</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-indigo-50">
                    <tr v-for="p in perDoctorStats.products" :key="p.product_id" class="hover:bg-indigo-50 transition">
                      <td class="px-4 py-2.5 font-medium text-gray-800">{{ p.product_name }}</td>
                      <td class="px-4 py-2.5 text-gray-600">{{ p.total_packs }} капс.</td>
                      <td class="px-4 py-2.5 text-gray-600">{{ p.total_pieces }} шт</td>
                      <td class="px-4 py-2.5 text-gray-600">{{ p.order_count }}</td>
                      <td class="px-4 py-2.5 text-right font-bold text-indigo-700">{{ formatPrice(p.revenue) }} сўм</td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <p v-else class="text-sm text-gray-400 py-4 text-center">Продаж по этому врачу ещё нет</p>
            </div>
            <div v-else class="py-6 text-center text-gray-400 text-sm">
              Выберите врача и нажмите «Показать статистику»
            </div>
          </div>
        </div>
      </div>

      <!-- ===== Settings Tab ===== -->
      <div v-if="activeTab === 'faq'">
        <div class="grid lg:grid-cols-2 gap-6">
          <div class="bg-white rounded-xl shadow-sm p-6">
            <h2 class="text-2xl font-bold text-gray-800 mb-1">Частые вопросы</h2>
            <p class="text-sm text-gray-500 mb-6">Добавляйте вопросы и несколько ответов для каждого вопроса.</p>

            <form @submit.prevent="saveFaq" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1.5">Вопрос</label>
                <input
                  v-model="faqForm.question"
                  type="text"
                  placeholder="Например: Как принимать препарат?"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500"
                  required
                />
              </div>

              <div>
                <div class="flex items-center justify-between mb-2">
                  <label class="block text-sm font-medium text-gray-700">Ответы</label>
                  <button
                    type="button"
                    @click="addFaqAnswerField"
                    class="text-xs font-semibold text-teal-600 hover:text-teal-700"
                  >
                    + Добавить ответ
                  </button>
                </div>

                <div class="space-y-2">
                  <div v-for="(answer, index) in faqForm.answers" :key="index" class="flex gap-2 items-start">
                    <textarea
                      v-model="faqForm.answers[index]"
                      rows="2"
                      class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 resize-none"
                      :placeholder="`Ответ ${index + 1}`"
                    ></textarea>
                    <button
                      type="button"
                      @click="removeFaqAnswerField(index)"
                      class="mt-1 text-red-500 hover:text-red-700 p-1"
                      title="Удалить ответ"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>

              <div v-if="faqError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ faqError }}</div>

              <div class="flex gap-3">
                <button
                  type="submit"
                  :disabled="savingFaq"
                  class="bg-teal-600 text-white px-5 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium disabled:opacity-50"
                >
                  {{ savingFaq ? 'Сохранение...' : (editingFaqId ? 'Обновить вопрос' : 'Добавить вопрос') }}
                </button>
                <button
                  v-if="editingFaqId"
                  type="button"
                  @click="resetFaqForm"
                  class="border border-gray-300 px-5 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium"
                >
                  Отмена
                </button>
              </div>
            </form>
          </div>

          <div class="bg-white rounded-xl shadow-sm p-6">
            <h3 class="text-lg font-bold text-gray-800 mb-4">Список вопросов</h3>
            <div class="space-y-3 max-h-[640px] overflow-y-auto pr-1">
              <div v-for="faq in faqs" :key="faq.id" class="border border-gray-200 rounded-xl overflow-hidden">
                <button
                  @click="expandedFaqId = expandedFaqId === faq.id ? null : faq.id"
                  class="w-full text-left px-4 py-3 bg-gray-50 hover:bg-gray-100 transition flex justify-between items-center"
                >
                  <span class="font-medium text-gray-800">{{ faq.question }}</span>
                  <svg class="w-4 h-4 text-gray-500 transition-transform" :class="expandedFaqId === faq.id ? 'rotate-180' : ''" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>

                <div v-if="expandedFaqId === faq.id" class="px-4 py-3 bg-white border-t border-gray-200">
                  <ol class="space-y-2 mb-4">
                    <li v-for="answer in faq.answers" :key="answer.id" class="text-sm text-gray-600 leading-relaxed">
                      • {{ answer.text }}
                    </li>
                  </ol>
                  <div class="flex gap-2">
                    <button @click="editFaq(faq)" class="text-xs font-semibold text-teal-600 hover:text-teal-700">Редактировать</button>
                    <button @click="deleteFaq(faq.id)" class="text-xs font-semibold text-red-500 hover:text-red-600">Удалить</button>
                  </div>
                </div>
              </div>

              <div v-if="faqs.length === 0" class="text-sm text-gray-400 py-10 text-center border border-dashed border-gray-200 rounded-xl">
                Вопросов пока нет
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="activeTab === 'support'">
        <div class="grid lg:grid-cols-5 gap-6">
          <div class="lg:col-span-2 bg-white rounded-xl shadow-sm p-4 sm:p-5">
            <h2 class="text-xl font-bold text-gray-800 mb-4">Обращения пользователей</h2>

            <div v-if="supportLoading" class="text-sm text-gray-400">Загрузка...</div>

            <div v-else class="space-y-2 max-h-[72vh] overflow-y-auto pr-1">
              <button
                v-for="thread in supportThreads"
                :key="thread.id"
                @click="openSupportThread(thread.id)"
                class="w-full text-left p-3 rounded-xl border transition"
                :class="selectedThread?.id === thread.id ? 'border-teal-500 bg-teal-50' : 'border-gray-200 hover:bg-gray-50'"
              >
                <div class="font-semibold text-gray-800 truncate">{{ thread.user?.first_name }} {{ thread.user?.last_name }}</div>
                <div class="text-xs text-gray-500 truncate mt-1">+{{ thread.user?.phone }}</div>
                <div class="text-xs text-gray-400 truncate mt-1">{{ supportThreadPreview(thread) }}</div>
                <div class="text-[10px] text-gray-400 mt-1">{{ supportThreadTime(thread) }}</div>
              </button>

              <div v-if="supportThreads.length === 0" class="text-sm text-gray-400 border border-dashed border-gray-200 rounded-xl p-6 text-center">
                Пока нет обращений
              </div>
            </div>
          </div>

          <div class="lg:col-span-3 bg-white rounded-xl shadow-sm flex flex-col h-[78vh]">
            <div class="px-5 py-4 border-b border-gray-200" v-if="selectedThread">
              <h3 class="font-bold text-gray-800">{{ selectedThread.user?.first_name }} {{ selectedThread.user?.last_name }}</h3>
              <p class="text-sm text-gray-500">+{{ selectedThread.user?.phone }}</p>
            </div>

            <div class="flex-1 overflow-y-auto px-5 py-4 bg-gray-50/70 space-y-3">
              <template v-if="selectedThread">
                <div
                  v-for="msg in selectedThread.messages"
                  :key="msg.id"
                  class="flex"
                  :class="msg.sender_role === 'admin' ? 'justify-end' : 'justify-start'"
                >
                  <div
                    class="max-w-[78%] px-4 py-2.5 rounded-2xl text-sm leading-relaxed shadow-sm"
                    :class="msg.sender_role === 'admin' ? 'bg-teal-600 text-white rounded-br-md' : 'bg-white text-gray-700 border border-gray-200 rounded-bl-md'"
                  >
                    <p class="whitespace-pre-line">{{ msg.message }}</p>
                    <p class="text-[10px] mt-1 opacity-70">{{ new Date(msg.created_at).toLocaleString('ru-RU') }}</p>
                  </div>
                </div>
              </template>

              <div v-else class="h-full flex items-center justify-center text-sm text-gray-400">
                Выберите диалог слева
              </div>
            </div>

            <form @submit.prevent="sendSupportReply" class="p-4 border-t border-gray-200 bg-white" v-if="selectedThread">
              <div class="flex gap-3 items-end">
                <textarea
                  v-model="supportReply"
                  rows="2"
                  placeholder="Напишите ответ пользователю..."
                  class="flex-1 border border-gray-300 rounded-xl px-3 py-2.5 resize-none focus:outline-none focus:ring-2 focus:ring-teal-500"
                ></textarea>
                <button
                  type="submit"
                  :disabled="supportSending || !supportReply.trim()"
                  class="bg-teal-600 text-white px-5 py-2.5 rounded-xl hover:bg-teal-700 transition disabled:opacity-50"
                >
                  {{ supportSending ? 'Отправка...' : 'Ответить' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <!-- ===== Settings Tab ===== -->
      <!-- ===== News Tab ===== -->
      <div v-if="activeTab === 'news'">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Новости</h2>
          <button @click="openNewsModal()" class="bg-teal-600 text-white px-5 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
            Добавить новость
          </button>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-5">
          <div v-for="post in newsPosts" :key="post.id" class="bg-white rounded-xl shadow-sm overflow-hidden border border-gray-100">
            <div v-if="post.image_path" class="h-44 overflow-hidden">
              <img :src="post.image_path" :alt="post.title" class="w-full h-full object-cover" />
            </div>
            <div v-else-if="post.video_url" class="h-44 bg-gray-100 flex items-center justify-center text-gray-400">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.91 11.672a.375.375 0 010 .656l-5.603 3.113a.375.375 0 01-.557-.328V8.887c0-.286.307-.466.557-.327l5.603 3.112z"/></svg>
            </div>
            <div class="p-4">
              <p class="text-xs text-gray-400 mb-1">{{ new Date(post.created_at).toLocaleDateString('ru-RU') }}</p>
              <h3 class="font-semibold text-gray-800 mb-1.5">{{ post.title }}</h3>
              <p v-if="post.description" class="text-sm text-gray-500 line-clamp-3">{{ post.description }}</p>
              <div class="flex gap-2 mt-3">
                <button @click="openNewsModal(post)" class="text-teal-600 text-sm font-medium hover:text-teal-800 transition">Редактировать</button>
                <button @click="deleteNews(post.id)" class="text-red-500 text-sm font-medium hover:text-red-700 transition">Удалить</button>
              </div>
            </div>
          </div>

          <div v-if="newsPosts.length === 0" class="col-span-3 text-center py-20 text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" />
            </svg>
            Новостей пока нет
          </div>
        </div>
      </div>

      <!-- News Modal -->
      <div v-if="showNewsModal" class="fixed inset-0 z-50 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showNewsModal = false"></div>
        <div class="relative bg-white rounded-2xl p-6 w-full max-w-lg mx-4 shadow-2xl max-h-[90vh] overflow-y-auto">
          <h3 class="text-xl font-bold text-gray-800 mb-5">{{ editingNews ? 'Редактировать новость' : 'Добавить новость' }}</h3>
          <form @submit.prevent="saveNews" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Заголовок <span class="text-red-400">*</span></label>
              <input v-model="newsForm.title" required class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Описание</label>
              <textarea v-model="newsForm.description" rows="4" class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition resize-none"></textarea>
            </div>

            <!-- Existing images (when editing) -->
            <div v-if="editingNews && editingNews.images && editingNews.images.length > 0">
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Текущие изображения</label>
              <div class="flex flex-wrap gap-2">
                <div v-for="img in editingNews.images" :key="img.id" class="relative group">
                  <img :src="img.image_path" class="h-20 w-20 object-cover rounded-lg border border-gray-200" />
                  <button
                    type="button"
                    @click="deleteNewsImg(img.id)"
                    class="absolute -top-2 -right-2 w-5 h-5 bg-red-500 text-white rounded-full flex items-center justify-center text-xs opacity-0 group-hover:opacity-100 transition-opacity"
                  >×</button>
                </div>
              </div>
            </div>
            <!-- Single image (legacy) -->
            <div v-else-if="editingNews && editingNews.image_path && !newsImagePreview">
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Текущее изображение</label>
              <img :src="editingNews.image_path" class="h-24 rounded-lg object-cover border border-gray-200" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Добавить фотографии (несколько)</label>
              <input type="file" accept="image/*" multiple ref="newsImagesInput" @change="onNewsImagesSelect" class="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-teal-50 file:text-teal-700 file:font-medium" />
              <div v-if="newsImagePreviews.length > 0" class="flex flex-wrap gap-2 mt-2">
                <img v-for="(src, i) in newsImagePreviews" :key="i" :src="src" class="h-20 w-20 object-cover rounded-lg border border-gray-200" />
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Ссылка на видео (YouTube, Instagram, Facebook, Twitter)</label>
              <input v-model="newsForm.video_url" type="url" placeholder="https://www.youtube.com/watch?v=..." class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div v-if="newsError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ newsError }}</div>
            <div class="flex gap-3 pt-2">
              <button type="submit" :disabled="savingNews" class="flex-1 bg-teal-600 text-white py-2.5 rounded-lg hover:bg-teal-700 transition font-medium disabled:opacity-50">
                {{ savingNews ? 'Сохранение...' : 'Сохранить' }}
              </button>
              <button type="button" @click="showNewsModal = false" class="flex-1 bg-gray-100 text-gray-700 py-2.5 rounded-lg hover:bg-gray-200 transition font-medium">
                Отмена
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Edit Doctor Modal -->
      <div v-if="showEditDoctorModal" class="fixed inset-0 z-50 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showEditDoctorModal = false"></div>
        <div class="relative bg-white rounded-2xl p-6 w-full max-w-md mx-4 shadow-2xl">
          <h3 class="text-xl font-bold text-gray-800 mb-5">Редактировать доктора</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Имя доктора <span class="text-red-400">*</span></label>
              <input v-model="editDoctorForm.name" type="text" required placeholder="Например: Абдурахман Каримов"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Специализация</label>
              <input v-model="editDoctorForm.specialty" type="text" placeholder="Например: Невролог, Кардиолог..."
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона</label>
              <div class="flex">
                <span class="inline-flex items-center px-3 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">+998</span>
                <input v-model="editDoctorForm.phone" type="tel" maxlength="9" placeholder="901234567"
                  class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
              </div>
              <p class="text-xs text-gray-400 mt-1">Для доступа к панели врача (необязательно)</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Новый пароль</label>
              <input v-model="editDoctorForm.password" type="password" placeholder="Оставьте пустым, чтобы не менять"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div v-if="editDoctorError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ editDoctorError }}</div>
            <div class="flex gap-3">
              <button type="button" @click="showEditDoctorModal = false"
                class="flex-1 border border-gray-300 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium">Отмена</button>
              <button @click="saveEditDoctor" :disabled="savingEditDoctor"
                class="flex-1 bg-teal-600 text-white py-2.5 rounded-lg hover:bg-teal-700 transition font-medium disabled:opacity-50">
                {{ savingEditDoctor ? 'Сохранение...' : 'Сохранить' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Doctor Modal -->
      <div v-if="showDoctorModal" class="fixed inset-0 z-50 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showDoctorModal = false"></div>
        <div class="relative bg-white rounded-2xl p-6 w-full max-w-md mx-4 shadow-2xl">
          <h3 class="text-xl font-bold text-gray-800 mb-5">Добавить доктора</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Имя доктора <span class="text-red-400">*</span></label>
              <input v-model="doctorName" type="text" required placeholder="Например: Абдурахман Каримов"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-purple-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Специализация</label>
              <input v-model="doctorSpecialty" type="text" placeholder="Например: Невролог, Кардиолог, Гинеколог..."
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-purple-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона</label>
              <div class="flex">
                <span class="inline-flex items-center px-3 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">+998</span>
                <input v-model="doctorPhone" type="tel" maxlength="9" placeholder="901234567"
                  class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-purple-500 transition" />
              </div>
              <p class="text-xs text-gray-400 mt-1">Для доступа к панели врача (необязательно)</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Пароль</label>
              <input v-model="doctorPassword" type="password" placeholder="Минимум 6 символов (необязательно)"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-purple-500 transition" />
            </div>
            <div v-if="doctorError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ doctorError }}</div>
            <div class="flex gap-3">
              <button type="button" @click="showDoctorModal = false"
                class="flex-1 border border-gray-300 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium">Отмена</button>
              <button @click="saveDoctor" :disabled="savingDoctor"
                class="flex-1 bg-purple-600 text-white py-2.5 rounded-lg hover:bg-purple-700 transition font-medium disabled:opacity-50">
                {{ savingDoctor ? 'Сохранение...' : 'Добавить' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="activeTab === 'settings'">
        <h2 class="text-2xl font-bold text-gray-800 mb-6">Настройки</h2>
        <div class="bg-white rounded-xl shadow-sm p-6 max-w-lg">
          <form @submit.prevent="saveSettings" class="space-y-5">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона</label>
              <div class="flex">
                <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm font-medium">+998</span>
                <input
                  v-model="settings.phone"
                  type="tel"
                  maxlength="9"
                  class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
                />
              </div>
            </div>

            <hr class="border-gray-200" />

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Текущий пароль</label>
              <input v-model="settings.old_password" type="password" placeholder="Введите текущий пароль"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Новый пароль</label>
              <input v-model="settings.new_password" type="password" placeholder="Введите новый пароль"
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>

            <div v-if="settingsError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ settingsError }}</div>
            <div v-if="settingsSuccess" class="bg-green-50 text-green-600 text-sm p-3 rounded-lg">{{ settingsSuccess }}</div>

            <button type="submit" class="bg-teal-600 text-white px-6 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium">
              Сохранить изменения
            </button>
          </form>
        </div>
      </div>
    </div>

    <!-- Product Modal -->
    <div v-if="showProductModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showProductModal = false"></div>
      <div class="relative bg-white rounded-2xl p-6 w-full max-w-md mx-4 shadow-2xl max-h-[90vh] overflow-y-auto">
        <h3 class="text-xl font-bold text-gray-800 mb-5">
          {{ editingProduct ? 'Изменить препарат' : 'Добавить препарат' }}
        </h3>
        <form @submit.prevent="saveProduct" class="space-y-4">
          <!-- Image Upload -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Фото препарата</label>
            <div class="flex items-center gap-4">
              <div class="w-24 h-24 bg-gray-100 rounded-xl overflow-hidden flex items-center justify-center border-2 border-dashed border-gray-300 flex-shrink-0">
                <img v-if="imagePreview || (editingProduct && editingProduct.image_path)" 
                  :src="imagePreview || editingProduct.image_path" 
                  class="w-full h-full object-cover" />
                <svg v-else class="w-8 h-8 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              <div class="flex-1">
                <label class="cursor-pointer inline-flex items-center gap-2 bg-blue-50 text-blue-600 px-4 py-2.5 rounded-lg hover:bg-blue-100 transition font-medium text-sm border border-blue-200">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" /></svg>
                  Выбрать фото
                  <input type="file" accept="image/*" class="hidden" @change="onModalImageSelect" />
                </label>
                <p class="text-xs text-gray-400 mt-1.5">JPG, PNG, WEBP до 10 МБ</p>
              </div>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Название <span class="text-red-400">*</span></label>
            <input v-model="productForm.name" type="text" required
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Описание</label>
            <textarea v-model="productForm.description" rows="3"
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition resize-none"></textarea>
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Кол-во в капсуле <span class="text-red-400">*</span></label>
              <input v-model.number="productForm.quantity_per_pack" type="number" min="1" required
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Цена за 1 шт (сўм) <span class="text-red-400">*</span></label>
              <input v-model.number="productForm.price_per_pill" type="number" min="0" step="100" required
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Количество на складе (штук)</label>
            <input v-model.number="productForm.stock_quantity" type="number" min="0"
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
              placeholder="0" />
            <p class="text-xs text-gray-400 mt-1">
              <span v-if="productForm.stock_quantity > 0 && productForm.quantity_per_pack > 0">
                = {{ Math.floor(productForm.stock_quantity / productForm.quantity_per_pack) }} капсул на складе
              </span>
              <span v-else>Указывается в штуках</span>
            </p>
          </div>

          <div v-if="productForm.quantity_per_pack > 0 && productForm.price_per_pill > 0" class="bg-teal-50 rounded-lg p-3.5 border border-teal-100">
            <div class="text-sm text-teal-800">
              <span class="text-teal-600">Цена за капсулу ({{ productForm.quantity_per_pack }} шт):</span>
              <span class="font-bold text-teal-700 ml-1">{{ formatPrice(productForm.quantity_per_pack * productForm.price_per_pill) }} сўм</span>
            </div>
          </div>

          <div v-if="productError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ productError }}</div>

          <div class="flex gap-3 pt-2">
            <button type="button" @click="showProductModal = false"
              class="flex-1 border border-gray-300 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium">
              Отмена
            </button>
            <button type="submit" :disabled="savingProduct"
              class="flex-1 bg-teal-600 text-white py-2.5 rounded-lg hover:bg-teal-700 transition font-medium disabled:opacity-50">
              {{ savingProduct ? 'Сохранение...' : (editingProduct ? 'Сохранить' : 'Добавить') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Worker Modal (Add / Edit) -->
    <div v-if="showWorkerModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showWorkerModal = false"></div>
      <div class="relative bg-white rounded-2xl p-6 w-full max-w-md mx-4 shadow-2xl">
        <h3 class="text-xl font-bold text-gray-800 mb-5">{{ editingWorker ? 'Редактировать работника' : 'Добавить работника' }}</h3>
        <form @submit.prevent="saveWorker" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Имя <span class="text-red-400">*</span></label>
            <input v-model="workerForm.name" type="text" :required="!editingWorker"
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона <span class="text-red-400">*</span></label>
            <div class="flex">
              <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm font-medium">+998</span>
              <input v-model="workerForm.phoneDigits" type="tel" maxlength="9" :required="!editingWorker"
                class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Роль <span class="text-red-400">*</span></label>
            <select v-model="workerForm.role"
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition">
              <option value="pickup">Пункт выдачи</option>
              <option value="nurse">Медсестра / Медбрат</option>
              <option value="manager">Менеджер (маркетплейсы)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">
              Пароль <span v-if="!editingWorker" class="text-red-400">*</span>
              <span v-else class="text-gray-400 text-xs">(оставьте пустым, чтобы не менять)</span>
            </label>
            <input v-model="workerForm.password" type="password" :required="!editingWorker" minlength="6"
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>
          <div v-if="workerError" class="bg-red-50 text-red-600 text-sm p-3 rounded-lg">{{ workerError }}</div>
          <div class="flex gap-3 pt-2">
            <button type="button" @click="showWorkerModal = false"
              class="flex-1 border border-gray-300 py-2.5 rounded-lg hover:bg-gray-50 transition font-medium">
              Отмена
            </button>
            <button type="submit" :disabled="savingWorker"
              class="flex-1 bg-teal-600 text-white py-2.5 rounded-lg hover:bg-teal-700 transition font-medium disabled:opacity-50">
              {{ savingWorker ? 'Сохранение...' : (editingWorker ? 'Сохранить' : 'Добавить') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Image upload hidden input (for table button) -->
    <input ref="imageInput" type="file" accept="image/*" class="hidden" @change="handleImageUpload" />

    <!-- Route Modal -->
    <div v-if="showRouteModal" class="fixed inset-0 z-[60] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="closeRouteModal"></div>
      <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-2xl mx-4 overflow-hidden">
        <div class="flex items-center justify-between px-6 py-4 border-b border-gray-100">
          <div>
            <h3 class="text-lg font-bold text-gray-800">Маршрут доставки</h3>
            <p v-if="routeOrder" class="text-sm text-gray-500">Заказ {{ routeOrder.order_code }} · {{ routeOrder.delivery_address }}</p>
          </div>
          <button @click="closeRouteModal" class="p-2 hover:bg-gray-100 rounded-xl transition-colors">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
          </button>
        </div>
        <div v-if="routeLoading" class="flex items-center justify-center" style="height: 420px">
          <div class="text-center text-gray-400">
            <div class="w-8 h-8 border-2 border-teal-600 border-t-transparent rounded-full animate-spin mx-auto mb-3"></div>
            <p class="text-sm">Построение маршрута...</p>
          </div>
        </div>
        <div v-show="!routeLoading" style="height: 420px">
          <div id="admin-route-map" style="height: 100%; width: 100%;"></div>
        </div>
        <div class="px-6 py-3 bg-gray-50 border-t border-gray-100 text-xs text-gray-500 flex items-center gap-1">
          <svg class="w-3.5 h-3.5 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/></svg>
          Дорожный маршрут от аптеки до адреса доставки
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useStockSocket, displayStock, realtime } from '../stores/stock'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'
import { Chart, registerables } from 'chart.js'
Chart.register(...registerables)

const authStore = useAuthStore()
const router = useRouter()
useStockSocket()
function liveQty(p) { return displayStock(p.id, p.stock_quantity) }
function liveCaps(p) { return Math.floor(liveQty(p) / (p.quantity_per_pack > 0 ? p.quantity_per_pack : 1)) }

const activeTab = ref('products')
const tabs = [
  { id: 'products', label: 'Препараты' },
  { id: 'orders', label: 'Заказы' },
  { id: 'workers', label: 'Работники' },
  { id: 'doctors', label: 'Доктора' },
  { id: 'analytics', label: 'Аналитика' },
  { id: 'faq', label: 'FAQ' },
  { id: 'support', label: 'Поддержка' },
  { id: 'news', label: 'Новости' },
  { id: 'settings', label: 'Настройки' },
]

// Products
const products = ref([])
const showProductModal = ref(false)
const editingProduct = ref(null)
const productForm = reactive({
  name: '',
  description: '',
  quantity_per_pack: 60,
  price_per_pill: 6500,
  stock_quantity: 0
})
const productError = ref('')
const savingProduct = ref(false)
const imageInput = ref(null)
const imageUploadProductId = ref(null)
const modalImageFile = ref(null)
const imagePreview = ref(null)

// Orders
const orders = ref([])

// Workers
const workers = ref([])
const showWorkerModal = ref(false)
const editingWorker = ref(null)
const workerForm = reactive({ name: '', phoneDigits: '', password: '', role: 'pickup' })
const workerError = ref('')
const savingWorker = ref(false)

// Analytics
const analyticsPeriod = ref('daily')
const analyticsCustomDate = ref(new Date().toISOString().split('T')[0])
const analyticsData = ref(null)
const analyticsLoading = ref(false)
const analyticsCanvas = ref(null)
let analyticsChart = null

const analyticsPeriods = [
  { value: 'daily', label: 'День' },
  { value: 'weekly', label: 'Неделя' },
  { value: 'monthly', label: 'Месяц' },
  { value: 'custom', label: 'Выбрать дату' },
]

// FAQ
const faqs = ref([])
const faqForm = reactive({ question: '', answers: [''] })
const editingFaqId = ref(null)
const savingFaq = ref(false)
const faqError = ref('')
const expandedFaqId = ref(null)

// Support
const supportThreads = ref([])
const selectedThread = ref(null)
const supportReply = ref('')
const supportSending = ref(false)
const supportLoading = ref(false)

// Settings
const settings = reactive({ phone: '', old_password: '', new_password: '' })
const settingsError = ref('')
const settingsSuccess = ref('')

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

function paymentLabel(order) {
  if (order.is_vip) return 'Свой пациент · Бесплатно'
  const m = {
    cash: 'Наличные',
    terminal: 'Терминал',
    card: 'Карта',
    online: 'Онлайн (карта)',
  }
  return m[order.payment_method] || ''
}

// Difference between what the doctor prescribed and what was actually bought at the till.
function itemDiffLabel(item) {
  const orig = item.original_quantity || 0
  const qty = item.quantity || 0
  if (orig === 0 && qty > 0) return 'добавлено'
  if (qty === 0 && orig > 0) return 'убрано'
  if (qty < orig) return 'убрано ' + (orig - qty)
  if (qty > orig) return 'добавлено ' + (qty - orig)
  return ''
}

function itemDiffClass(item) {
  const orig = item.original_quantity || 0
  const qty = item.quantity || 0
  if (qty === 0) return 'bg-red-100 text-red-600'
  if (qty < orig) return 'bg-orange-100 text-orange-700'
  return 'bg-emerald-100 text-emerald-700'
}

function paymentBadgeClass(order) {
  if (order.is_vip) return 'bg-amber-100 text-amber-700'
  if (order.payment_method === 'cash') return 'bg-green-100 text-green-700'
  if (order.payment_method === 'online') return 'bg-indigo-100 text-indigo-700'
  return 'bg-blue-100 text-blue-700'
}

function statusLabel(status) {
  const labels = {
    pending: 'Ожидает',
    confirmed: 'Подтверждён',
    shipped: 'Отправлен',
    in_transit: 'В пути',
    delivered: 'Доставлен',
    cancelled: 'Отменён'
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
    cancelled: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

async function loadProducts() {
  try {
    const res = await api.get('/admin/products')
    products.value = res.data || []
  } catch (e) {
    console.error(e)
  }
}

async function loadOrders() {
  try {
    const res = await api.get('/admin/orders')
    orders.value = res.data || []
  } catch (e) {
    console.error(e)
  }
}

async function loadProfile() {
  try {
    const res = await api.get('/admin/profile')
    settings.phone = res.data.phone.replace('998', '')
  } catch (e) {
    console.error(e)
  }
}

async function loadWorkers() {
  try {
    const res = await api.get('/admin/workers')
    workers.value = res.data || []
  } catch (e) {
    console.error(e)
  }
}

function selectPeriod(period) {
  analyticsPeriod.value = period
  loadAnalytics()
}

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const params = new URLSearchParams({ period: analyticsPeriod.value })
    if (analyticsPeriod.value === 'custom' && analyticsCustomDate.value) {
      params.set('date', analyticsCustomDate.value)
    }
    const res = await api.get(`/admin/analytics?${params}`)
    analyticsData.value = res.data
    analyticsLoading.value = false
    if (perMktId.value) loadPerMktStats()
    await nextTick()
    renderAnalyticsChart(res.data)
  } catch (e) {
    console.error(e)
    analyticsLoading.value = false
  }
}

function renderAnalyticsChart(data) {
  if (!analyticsCanvas.value) return
  if (analyticsChart) {
    analyticsChart.destroy()
    analyticsChart = null
  }
  const points = data.points || []
  const labels = points.map(p => p.label)
  const revenues = points.map(p => p.revenue)
  const ordersData = points.map(p => p.orders)

  analyticsChart = new Chart(analyticsCanvas.value, {
    type: 'line',
    data: {
      labels,
      datasets: [
        {
          label: 'Выручка (сўм)',
          data: revenues,
          borderColor: '#0d9488',
          backgroundColor: 'rgba(13,148,136,0.1)',
          fill: true,
          tension: 0.4,
          pointRadius: 3,
          pointHoverRadius: 6,
          yAxisID: 'yRevenue',
        },
        {
          label: 'Заказы',
          data: ordersData,
          borderColor: '#3b82f6',
          backgroundColor: 'rgba(59,130,246,0.08)',
          fill: false,
          tension: 0.4,
          pointRadius: 3,
          pointHoverRadius: 6,
          yAxisID: 'yOrders',
        },
      ],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      interaction: { mode: 'index', intersect: false },
      plugins: {
        legend: { position: 'top' },
        tooltip: {
          callbacks: {
            label(ctx) {
              if (ctx.datasetIndex === 0) return ` Выручка: ${new Intl.NumberFormat('ru-RU').format(Math.round(ctx.raw))} сўм`
              return ` Заказов: ${ctx.raw}`
            }
          }
        }
      },
      scales: {
        yRevenue: {
          type: 'linear',
          position: 'left',
          beginAtZero: true,
          ticks: {
            callback: v => new Intl.NumberFormat('ru-RU', { notation: 'compact' }).format(v)
          },
          grid: { color: 'rgba(0,0,0,0.05)' },
        },
        yOrders: {
          type: 'linear',
          position: 'right',
          beginAtZero: true,
          grid: { drawOnChartArea: false },
          ticks: { stepSize: 1 },
        },
        x: {
          grid: { color: 'rgba(0,0,0,0.04)' },
          ticks: { maxRotation: 45, font: { size: 11 } }
        }
      },
    },
  })
}

async function loadFaqs() {
  try {
    const res = await api.get('/admin/faqs')
    faqs.value = res.data || []
  } catch (e) {
    console.error(e)
  }
}

function resetFaqForm() {
  editingFaqId.value = null
  faqForm.question = ''
  faqForm.answers = ['']
  faqError.value = ''
}

function addFaqAnswerField() {
  faqForm.answers.push('')
}

function removeFaqAnswerField(index) {
  if (faqForm.answers.length === 1) {
    faqForm.answers[0] = ''
    return
  }
  faqForm.answers.splice(index, 1)
}

function editFaq(faq) {
  editingFaqId.value = faq.id
  faqForm.question = faq.question
  faqForm.answers = (faq.answers || []).map(item => item.text)
  if (faqForm.answers.length === 0) faqForm.answers = ['']
  faqError.value = ''
}

async function saveFaq() {
  faqError.value = ''
  savingFaq.value = true

  const payload = {
    question: faqForm.question,
    answers: faqForm.answers
  }

  try {
    if (editingFaqId.value) {
      await api.put(`/admin/faqs/${editingFaqId.value}`, payload)
    } else {
      await api.post('/admin/faqs', payload)
    }
    resetFaqForm()
    await loadFaqs()
  } catch (e) {
    faqError.value = e.response?.data?.error || 'Ошибка при сохранении FAQ'
  } finally {
    savingFaq.value = false
  }
}

async function deleteFaq(id) {
  if (!confirm('Удалить этот вопрос?')) return
  try {
    await api.delete(`/admin/faqs/${id}`)
    if (expandedFaqId.value === id) expandedFaqId.value = null
    if (editingFaqId.value === id) resetFaqForm()
    await loadFaqs()
  } catch (e) {
    alert('Ошибка при удалении FAQ')
  }
}

function supportThreadPreview(thread) {
  const last = thread.messages?.[0]
  return last?.message || 'Нет сообщений'
}

function supportThreadTime(thread) {
  const last = thread.messages?.[0]
  const date = last?.created_at || thread.updated_at
  if (!date) return ''
  return new Date(date).toLocaleString('ru-RU')
}

async function loadSupportThreads() {
  supportLoading.value = true
  try {
    const res = await api.get('/admin/support/threads')
    supportThreads.value = res.data || []
    if (!selectedThread.value && supportThreads.value.length > 0) {
      await openSupportThread(supportThreads.value[0].id)
    }
  } catch (e) {
    console.error(e)
  } finally {
    supportLoading.value = false
  }
}

async function openSupportThread(id) {
  try {
    const res = await api.get(`/admin/support/threads/${id}`)
    selectedThread.value = res.data
  } catch (e) {
    console.error(e)
  }
}

async function sendSupportReply() {
  const message = supportReply.value.trim()
  if (!message || !selectedThread.value) return

  supportSending.value = true
  try {
    await api.post(`/admin/support/threads/${selectedThread.value.id}/reply`, { message })
    supportReply.value = ''
    await openSupportThread(selectedThread.value.id)
    await loadSupportThreads()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при отправке ответа')
  } finally {
    supportSending.value = false
  }
}

function openWorkerModal() {
  editingWorker.value = null
  workerForm.name = ''
  workerForm.phoneDigits = ''
  workerForm.password = ''
  workerForm.role = 'pickup'
  workerError.value = ''
  showWorkerModal.value = true
}

function openEditWorkerModal(w) {
  editingWorker.value = w
  workerForm.name = w.name
  workerForm.phoneDigits = w.phone.replace(/^998/, '')
  workerForm.password = ''
  workerForm.role = w.role || 'pickup'
  workerError.value = ''
  showWorkerModal.value = true
}

async function saveWorker() {
  workerError.value = ''
  savingWorker.value = true
  try {
    if (editingWorker.value) {
      const payload = { role: workerForm.role }
      if (workerForm.name) payload.name = workerForm.name
      if (workerForm.phoneDigits) payload.phone = '998' + workerForm.phoneDigits
      if (workerForm.password) payload.password = workerForm.password
      await api.put(`/admin/workers/${editingWorker.value.id}`, payload)
    } else {
      await api.post('/admin/workers', {
        name: workerForm.name,
        phone: '998' + workerForm.phoneDigits,
        password: workerForm.password,
        role: workerForm.role,
      })
    }
    showWorkerModal.value = false
    editingWorker.value = null
    workerForm.name = ''
    workerForm.phoneDigits = ''
    workerForm.password = ''
    workerForm.role = 'pickup'
    await loadWorkers()
  } catch (e) {
    workerError.value = e.response?.data?.error || 'Ошибка'
  } finally {
    savingWorker.value = false
  }
}

async function deleteWorker(id) {
  if (!confirm('Удалить этого работника?')) return
  try {
    await api.delete(`/admin/workers/${id}`)
    await loadWorkers()
  } catch (e) {
    alert('Ошибка при удалении')
  }
}

function openProductModal(product = null) {
  editingProduct.value = product
  modalImageFile.value = null
  imagePreview.value = null
  if (product) {
    productForm.name = product.name
    productForm.description = product.description || ''
    productForm.quantity_per_pack = product.quantity_per_pack
    productForm.price_per_pill = product.price_per_pill
    productForm.stock_quantity = product.stock_quantity || 0
  } else {
    productForm.name = ''
    productForm.description = ''
    productForm.quantity_per_pack = 60
    productForm.price_per_pill = 6500
    productForm.stock_quantity = 0
  }
  productError.value = ''
  showProductModal.value = true
}

function onModalImageSelect(e) {
  const file = e.target.files[0]
  if (!file) return
  modalImageFile.value = file
  imagePreview.value = URL.createObjectURL(file)
  e.target.value = ''
}

async function uploadImageForProduct(productId) {
  if (!modalImageFile.value) return
  const formData = new FormData()
  formData.append('image', modalImageFile.value)
  await api.post(`/admin/products/${productId}/image`, formData, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}

async function saveProduct() {
  productError.value = ''
  savingProduct.value = true
  try {
    let savedProduct
    if (editingProduct.value) {
      const res = await api.put(`/admin/products/${editingProduct.value.id}`, productForm)
      savedProduct = res.data
    } else {
      const res = await api.post('/admin/products', productForm)
      savedProduct = res.data
    }
    // Upload image if selected
    if (modalImageFile.value && savedProduct?.id) {
      await uploadImageForProduct(savedProduct.id)
    }
    showProductModal.value = false
    modalImageFile.value = null
    imagePreview.value = null
    await loadProducts()
  } catch (e) {
    productError.value = e.response?.data?.error || 'Ошибка'
  } finally {
    savingProduct.value = false
  }
}

async function deleteProduct(id) {
  if (!confirm('Удалить этот препарат?')) return
  try {
    await api.delete(`/admin/products/${id}`)
    await loadProducts()
  } catch (e) {
    alert('Ошибка при удалении')
  }
}

function uploadImageFor(product) {
  imageUploadProductId.value = product.id
  imageInput.value.click()
}

async function handleImageUpload(e) {
  const file = e.target.files[0]
  if (!file) return
  const formData = new FormData()
  formData.append('image', file)

  try {
    await api.post(`/admin/products/${imageUploadProductId.value}/image`, formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    await loadProducts()
  } catch (err) {
    alert('Ошибка при загрузке изображения')
  }
  e.target.value = ''
}

async function updateOrderStatus(id, status) {
  try {
    await api.put(`/admin/orders/${id}/status`, { status })
    await loadOrders()
  } catch (e) {
    alert('Ошибка при обновлении статуса')
  }
}

async function deleteOrder(id) {
  if (!confirm('Удалить этот заказ? Действие нельзя отменить.')) return
  try {
    await api.delete(`/admin/orders/${id}`)
    await loadOrders()
  } catch (e) {
    alert('Ошибка при удалении заказа')
  }
}

async function deleteAllOrders() {
  // Step 1: initial warning.
  if (!confirm('Удалить ВСЕ заказы? Они исчезнут из списков, но аналитика сохранится. Это действие нельзя отменить.')) return
  // Step 2: require the user to type the confirmation phrase or a reason.
  const input = prompt('Для подтверждения введите «удалить все» или укажите причину удаления:')
  if (input === null) return
  const text = input.trim()
  if (text === '') {
    alert('Подтверждение не получено. Удаление отменено.')
    return
  }
  const reason = text.toLowerCase() === 'удалить все' ? '' : text
  try {
    await api.delete('/admin/orders', { data: { reason } })
    await loadOrders()
  } catch (e) {
    alert('Ошибка при удалении заказов')
  }
}

async function saveSettings() {
  settingsError.value = ''
  settingsSuccess.value = ''

  const data = {}
  if (settings.phone) data.phone = '998' + settings.phone
  if (settings.new_password) {
    if (!settings.old_password) {
      settingsError.value = 'Введите текущий пароль'
      return
    }
    data.old_password = settings.old_password
    data.new_password = settings.new_password
  }

  try {
    await api.put('/admin/settings', data)
    settingsSuccess.value = 'Настройки успешно сохранены'
    settings.old_password = ''
    settings.new_password = ''
  } catch (e) {
    settingsError.value = e.response?.data?.error || 'Ошибка'
  }
}

function logout() {
  authStore.adminLogout()
  router.push('/admin/login')
}


// Doctors
const doctors = ref([])
const showDoctorModal = ref(false)
const doctorName = ref('')
const doctorSpecialty = ref('')
const doctorPhone = ref('')
const doctorPassword = ref('')
const doctorError = ref('')
const savingDoctor = ref(false)

// Edit doctor
const showEditDoctorModal = ref(false)
const editingDoctorId = ref(null)
const editDoctorForm = reactive({ name: '', specialty: '', phone: '', password: '' })
const editDoctorError = ref('')
const savingEditDoctor = ref(false)

// Doctor stats expand
const expandedDoctorId = ref(null)
const doctorStats = ref(null)
const doctorStatsLoading = ref(false)

// Per-doctor analytics detail
const perDoctorSelectedId = ref('')
const perDoctorStats = ref(null)
const perDoctorStatsLoading = ref(false)

// Marketolog analytics detail
const marketologList = ref([])
const perMktId = ref('')
const perMktStats = ref(null)

const breakdownCats = [
  { key: 'vip', label: 'Свои пациенты', box: 'border-emerald-100 bg-emerald-50/40', text: 'text-emerald-700' },
  { key: 'doctor', label: 'Через докторов', box: 'border-indigo-100 bg-indigo-50/40', text: 'text-indigo-700' },
  { key: 'regular', label: 'Обычные', box: 'border-gray-200 bg-gray-50', text: 'text-gray-700' },
  { key: 'marketolog', label: 'Маркетолог (долг)', box: 'border-purple-100 bg-purple-50/40', text: 'text-purple-700' },
]

async function loadMarketologList() {
  try { marketologList.value = (await api.get('/admin/marketologs')).data || [] } catch { marketologList.value = [] }
}
async function loadPerMktStats() {
  if (!perMktId.value) { perMktStats.value = null; return }
  try {
    const params = { period: analyticsPeriod.value }
    if (analyticsPeriod.value === 'custom') params.date = analyticsCustomDate.value
    perMktStats.value = (await api.get(`/admin/marketologs/${perMktId.value}/stats`, { params })).data
  } catch { perMktStats.value = null }
}

async function loadDoctors() {
  try {
    const res = await api.get('/admin/doctors')
    doctors.value = res.data || []
  } catch { doctors.value = [] }
}

async function saveDoctor() {
  if (!doctorName.value.trim()) return
  doctorError.value = ''
  savingDoctor.value = true
  try {
    const payload = {
      name: doctorName.value.trim(),
      specialty: doctorSpecialty.value.trim(),
    }
    if (doctorPhone.value.trim()) payload.phone = '998' + doctorPhone.value.trim()
    if (doctorPassword.value) payload.password = doctorPassword.value
    await api.post('/admin/doctors', payload)
    showDoctorModal.value = false
    doctorName.value = ''
    doctorSpecialty.value = ''
    doctorPhone.value = ''
    doctorPassword.value = ''
    await loadDoctors()
  } catch (e) {
    doctorError.value = e.response?.data?.error || 'Ошибка'
  } finally {
    savingDoctor.value = false
  }
}

async function deleteDoctor(id) {
  if (!confirm('Удалить этого доктора?')) return
  try {
    await api.delete(`/admin/doctors/${id}`)
    if (expandedDoctorId.value === id) expandedDoctorId.value = null
    await loadDoctors()
  } catch { alert('Ошибка при удалении') }
}

function openEditDoctorModal(doc) {
  editingDoctorId.value = doc.id
  editDoctorForm.name = doc.name
  editDoctorForm.specialty = doc.specialty || ''
  editDoctorForm.phone = doc.phone ? doc.phone.replace(/^998/, '') : ''
  editDoctorForm.password = ''
  editDoctorError.value = ''
  showEditDoctorModal.value = true
}

async function saveEditDoctor() {
  if (!editDoctorForm.name.trim()) return
  editDoctorError.value = ''
  savingEditDoctor.value = true
  try {
    const payload = {
      name: editDoctorForm.name.trim(),
      specialty: editDoctorForm.specialty.trim(),
    }
    if (editDoctorForm.phone.trim()) payload.phone = '998' + editDoctorForm.phone.trim()
    if (editDoctorForm.password) payload.password = editDoctorForm.password
    await api.put(`/admin/doctors/${editingDoctorId.value}`, payload)
    showEditDoctorModal.value = false
    await loadDoctors()
  } catch (e) {
    editDoctorError.value = e.response?.data?.error || 'Ошибка'
  } finally {
    savingEditDoctor.value = false
  }
}

async function toggleDoctorStats(doc) {
  if (expandedDoctorId.value === doc.id) {
    expandedDoctorId.value = null
    doctorStats.value = null
    return
  }
  expandedDoctorId.value = doc.id
  doctorStats.value = null
  doctorStatsLoading.value = true
  try {
    const res = await api.get(`/admin/doctors/${doc.id}/stats`)
    doctorStats.value = res.data
  } catch { doctorStats.value = { doctor: doc, total_orders: 0, products: [] } }
  finally { doctorStatsLoading.value = false }
}

// ===== Worker Stats =====
const expandedWorkerId = ref(null)
const workerStats = ref(null)
const workerStatsLoading = ref(false)

async function toggleWorkerStats(w) {
  if (expandedWorkerId.value === w.id) {
    expandedWorkerId.value = null
    workerStats.value = null
    return
  }
  expandedWorkerId.value = w.id
  workerStats.value = null
  workerStatsLoading.value = true
  try {
    const res = await api.get(`/admin/workers/${w.id}/stats`)
    workerStats.value = res.data
  } catch { workerStats.value = { today: { orders: 0, items: 0, revenue: 0 }, week: { orders: 0, items: 0, revenue: 0 }, month: { orders: 0, items: 0, revenue: 0 } } }
  finally { workerStatsLoading.value = false }
}

async function loadPerDoctorStats() {
  if (!perDoctorSelectedId.value) return
  perDoctorStats.value = null
  perDoctorStatsLoading.value = true
  try {
    const res = await api.get(`/admin/doctors/${perDoctorSelectedId.value}/stats`)
    perDoctorStats.value = res.data
  } catch {
    const doc = doctors.value.find(d => d.id === perDoctorSelectedId.value)
    perDoctorStats.value = { doctor: doc || { name: 'Врач' }, total_orders: 0, products: [] }
  } finally {
    perDoctorStatsLoading.value = false
  }
}

// News
const newsPosts = ref([])
const showNewsModal = ref(false)
const editingNews = ref(null)
const newsForm = reactive({ title: '', description: '', video_url: '' })
const newsImagesInput = ref(null)
const newsImageFiles = ref([])
const newsImagePreviews = ref([])
const newsError = ref('')
const savingNews = ref(false)

async function loadNews() {
  try {
    const res = await api.get('/admin/news')
    newsPosts.value = res.data || []
  } catch { newsPosts.value = [] }
}

function openNewsModal(post = null) {
  editingNews.value = post
  newsForm.title = post?.title || ''
  newsForm.description = post?.description || ''
  newsForm.video_url = post?.video_url || ''
  newsImageFiles.value = []
  newsImagePreviews.value = []
  newsError.value = ''
  showNewsModal.value = true
}

function onNewsImagesSelect(e) {
  const files = Array.from(e.target.files)
  newsImageFiles.value = files
  newsImagePreviews.value = files.map(f => URL.createObjectURL(f))
  e.target.value = ''
}

async function deleteNewsImg(imgId) {
  if (!confirm('Удалить это изображение?')) return
  try {
    await api.delete(`/admin/news/images/${imgId}`)
    if (editingNews.value) {
      editingNews.value = { ...editingNews.value, images: editingNews.value.images.filter(i => i.id !== imgId) }
    }
    await loadNews()
  } catch { alert('Ошибка при удалении изображения') }
}

async function saveNews() {
  newsError.value = ''
  savingNews.value = true
  try {
    const formData = new FormData()
    formData.append('title', newsForm.title)
    formData.append('description', newsForm.description)
    formData.append('video_url', newsForm.video_url)
    // Append multiple images
    for (const file of newsImageFiles.value) {
      formData.append('images', file)
    }

    if (editingNews.value) {
      await api.put(`/admin/news/${editingNews.value.id}`, formData, { headers: { 'Content-Type': 'multipart/form-data' } })
    } else {
      await api.post('/admin/news', formData, { headers: { 'Content-Type': 'multipart/form-data' } })
    }
    showNewsModal.value = false
    await loadNews()
  } catch (e) {
    newsError.value = e.response?.data?.error || 'Ошибка при сохранении'
  } finally {
    savingNews.value = false
  }
}

async function deleteNews(id) {
  if (!confirm('Удалить эту новость?')) return
  try {
    await api.delete(`/admin/news/${id}`)
    await loadNews()
  } catch { alert('Ошибка при удалении') }
}

// Route modal
const showRouteModal = ref(false)
const routeOrder = ref(null)
const routeLoading = ref(false)
let routeLeafletMap = null

async function openRouteModal(order) {
  routeOrder.value = order
  showRouteModal.value = true
  routeLoading.value = true
  await nextTick()
  await buildRouteMap(order)
}

function closeRouteModal() {
  showRouteModal.value = false
  routeOrder.value = null
  if (routeLeafletMap) {
    routeLeafletMap.remove()
    routeLeafletMap = null
  }
}

// Pharmacy coordinates cache (Andijan, Hodja 27)
let pharmacyCoords = null

async function getPharmacyCoords() {
  if (pharmacyCoords) return pharmacyCoords
  try {
    const res = await fetch(
      'https://nominatim.openstreetmap.org/search?q=%D1%83%D0%BB%D0%B8%D1%86%D0%B0+%D0%A5%D0%BE%D0%B4%D0%B6%D0%B0+27%2C+%D0%90%D0%BD%D0%B4%D0%B8%D0%B6%D0%B0%D0%BD%2C+%D0%A3%D0%B7%D0%B1%D0%B5%D0%BA%D0%B8%D1%81%D1%82%D0%B0%D0%BD&format=json&limit=1',
      { headers: { 'Accept-Language': 'ru' } }
    )
    const data = await res.json()
    if (data && data.length > 0) {
      pharmacyCoords = { lat: parseFloat(data[0].lat), lng: parseFloat(data[0].lon) }
      return pharmacyCoords
    }
  } catch { /* ignore */ }
  // Fallback: Andijan city center
  pharmacyCoords = { lat: 40.7821, lng: 72.3442 }
  return pharmacyCoords
}

async function buildRouteMap(order) {
  const L = await import('leaflet')
  await import('leaflet/dist/leaflet.css')

  delete L.Icon.Default.prototype._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })

  const mapEl = document.getElementById('admin-route-map')
  if (!mapEl) { routeLoading.value = false; return }

  // Pharmacy starting point: улица Ходжа 27, Андижан
  const pharmacy = await getPharmacyCoords()
  const startLat = pharmacy.lat
  const startLng = pharmacy.lng
  const destLat = order.latitude
  const destLng = order.longitude

  if (routeLeafletMap) {
    routeLeafletMap.remove()
    routeLeafletMap = null
  }

  routeLeafletMap = L.map('admin-route-map').setView(
    [(startLat + destLat) / 2, (startLng + destLng) / 2], 13
  )
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
  }).addTo(routeLeafletMap)

  // Pharmacy marker (green)
  const pharmacyIcon = L.divIcon({
    html: `<div style="background:#059669;width:14px;height:14px;border-radius:50%;border:2px solid white;box-shadow:0 1px 4px rgba(0,0,0,0.4)"></div>`,
    className: '',
    iconAnchor: [7, 7],
  })
  L.marker([startLat, startLng], { icon: pharmacyIcon })
    .addTo(routeLeafletMap)
    .bindPopup('<b>Аптека</b><br>ул. Ходжа 27, Андижан')

  // Customer marker (red)
  L.marker([destLat, destLng])
    .addTo(routeLeafletMap)
    .bindPopup(`<b>Адрес доставки</b><br>${order.delivery_address || ''}`)
    .openPopup()

  // Fetch road route from OSRM
  try {
    const url = `https://router.project-osrm.org/route/v1/driving/${startLng},${startLat};${destLng},${destLat}?overview=full&geometries=geojson`
    const res = await fetch(url)
    const data = await res.json()
    if (data.routes && data.routes.length > 0) {
      const coords = data.routes[0].geometry.coordinates.map(c => [c[1], c[0]])
      const polyline = L.polyline(coords, { color: '#0d9488', weight: 4, opacity: 0.85 }).addTo(routeLeafletMap)
      routeLeafletMap.fitBounds(polyline.getBounds(), { padding: [30, 30] })
    } else {
      // Fallback: straight line
      const line = L.polyline([[startLat, startLng], [destLat, destLng]], { color: '#6366f1', weight: 3, dashArray: '6 4' }).addTo(routeLeafletMap)
      routeLeafletMap.fitBounds(line.getBounds(), { padding: [30, 30] })
    }
  } catch {
    // Fallback: straight line
    const line = L.polyline([[startLat, startLng], [destLat, destLng]], { color: '#6366f1', weight: 3, dashArray: '6 4' }).addTo(routeLeafletMap)
    routeLeafletMap.fitBounds(line.getBounds(), { padding: [30, 30] })
  }

  routeLoading.value = false
}

onUnmounted(() => {
  if (routeLeafletMap) {
    routeLeafletMap.remove()
    routeLeafletMap = null
  }
})

watch(activeTab, async (tab) => {
  if (tab === 'analytics') {
    loadMarketologList()
    if (!analyticsData.value) {
      loadAnalytics()
    } else {
      await nextTick()
      renderAnalyticsChart(analyticsData.value)
    }
  }
  if (tab === 'news') {
    loadNews()
  }
  if (tab === 'doctors') {
    loadDoctors()
  }
})

// Real-time: refresh orders and analytics as soon as any order changes.
watch(() => realtime.ordersVersion, async () => {
  if (activeTab.value === 'orders') loadOrders()
  if (activeTab.value === 'analytics') loadAnalytics()
})

onMounted(() => {
  loadProducts()
  loadOrders()
  loadProfile()
  loadWorkers()
  loadFaqs()
  loadSupportThreads()
})
</script>
