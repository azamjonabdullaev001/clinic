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
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">В упаковке</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Цена/шт</th>
                  <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Цена/упак.</th>
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
        <h2 class="text-2xl font-bold text-gray-800 mb-6">Заказы</h2>
        <div class="space-y-4">
          <div v-for="order in orders" :key="order.id" class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex flex-col sm:flex-row justify-between items-start gap-4 mb-4">
              <div>
                <div class="flex items-center gap-2 mb-1">
                  <span class="text-xs font-medium text-gray-400 bg-gray-100 px-2 py-0.5 rounded">#{{ order.id }}</span>
                  <span v-if="order.order_code" class="text-sm font-bold text-blue-700 tracking-widest bg-blue-50 px-2 py-0.5 rounded">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">
                    {{ statusLabel(order.status) }}
                  </span>
                </div>
                <h3 class="font-semibold text-gray-800 text-lg">
                  {{ order.user?.first_name }} {{ order.user?.last_name }}
                </h3>
                <p class="text-sm text-gray-500">+{{ order.phone }}</p>
                <p class="text-xs text-gray-400 mt-1">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
              </div>
              <select
                :value="order.status"
                @change="updateOrderStatus(order.id, $event.target.value)"
                class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 bg-white"
              >
                <option value="pending">Ожидает</option>
                <option value="confirmed">Подтверждён</option>
                <option value="shipped">Отправлен</option>
                <option value="delivered">Доставлен</option>
                <option value="cancelled">Отменён</option>
              </select>
            </div>

            <div class="border-t pt-4 space-y-2">
              <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm py-1">
                <span class="text-gray-600">{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? 'шт' : 'упак.' }}</span></span>
                <span class="font-medium text-gray-700">{{ formatPrice(item.price) }} сўм</span>
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
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">Работники пункта выдачи</h2>
          <button @click="showWorkerModal = true" class="bg-teal-600 text-white px-5 py-2.5 rounded-lg hover:bg-teal-700 transition font-medium flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
            Добавить работника
          </button>
        </div>

        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <table class="w-full">
            <thead class="bg-gray-50 border-b">
              <tr>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Имя</th>
                <th class="text-left px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Телефон</th>
                <th class="text-right px-5 py-3.5 text-xs font-semibold text-gray-500 uppercase tracking-wider">Действия</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-100">
              <tr v-for="w in workers" :key="w.id" class="hover:bg-gray-50 transition">
                <td class="px-5 py-3 font-medium text-gray-800">{{ w.name }}</td>
                <td class="px-5 py-3 text-gray-500">+{{ w.phone }}</td>
                <td class="px-5 py-3 text-right">
                  <button @click="deleteWorker(w.id)" class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition" title="Удалить">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                  </button>
                </td>
              </tr>
              <tr v-if="workers.length === 0">
                <td colspan="3" class="px-5 py-12 text-center text-gray-400">
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
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Кол-во в упаковке <span class="text-red-400">*</span></label>
              <input v-model.number="productForm.quantity_per_pack" type="number" min="1" required
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1.5">Цена за 1 шт (сўм) <span class="text-red-400">*</span></label>
              <input v-model.number="productForm.price_per_pill" type="number" min="0" step="100" required
                class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
          </div>

          <div v-if="productForm.quantity_per_pack > 0 && productForm.price_per_pill > 0" class="bg-teal-50 rounded-lg p-3.5 border border-teal-100">
            <div class="text-sm text-teal-800">
              <span class="text-teal-600">Цена за упаковку ({{ productForm.quantity_per_pack }} шт):</span>
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

    <!-- Worker Modal -->
    <div v-if="showWorkerModal" class="fixed inset-0 z-50 flex items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showWorkerModal = false"></div>
      <div class="relative bg-white rounded-2xl p-6 w-full max-w-md mx-4 shadow-2xl">
        <h3 class="text-xl font-bold text-gray-800 mb-5">Добавить работника</h3>
        <form @submit.prevent="saveWorker" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Имя <span class="text-red-400">*</span></label>
            <input v-model="workerForm.name" type="text" required
              class="w-full border border-gray-300 rounded-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Номер телефона <span class="text-red-400">*</span></label>
            <div class="flex">
              <span class="inline-flex items-center px-3.5 rounded-l-lg border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm font-medium">+998</span>
              <input v-model="workerForm.phoneDigits" type="tel" maxlength="9" required
                class="flex-1 border border-gray-300 rounded-r-lg px-3 py-2.5 focus:outline-none focus:ring-2 focus:ring-teal-500 transition" />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1.5">Пароль <span class="text-red-400">*</span></label>
            <input v-model="workerForm.password" type="password" required minlength="6"
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
              {{ savingWorker ? 'Сохранение...' : 'Добавить' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Image upload hidden input (for table button) -->
    <input ref="imageInput" type="file" accept="image/*" class="hidden" @change="handleImageUpload" />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const activeTab = ref('products')
const tabs = [
  { id: 'products', label: 'Препараты' },
  { id: 'orders', label: 'Заказы' },
  { id: 'workers', label: 'Работники' },
  { id: 'faq', label: 'FAQ' },
  { id: 'support', label: 'Поддержка' },
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
  price_per_pill: 6500
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
const workerForm = reactive({ name: '', phoneDigits: '', password: '' })
const workerError = ref('')
const savingWorker = ref(false)

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

function statusLabel(status) {
  const labels = {
    pending: 'Ожидает',
    confirmed: 'Подтверждён',
    shipped: 'Отправлен',
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

async function saveWorker() {
  workerError.value = ''
  savingWorker.value = true
  try {
    await api.post('/admin/workers', {
      name: workerForm.name,
      phone: '998' + workerForm.phoneDigits,
      password: workerForm.password
    })
    showWorkerModal.value = false
    workerForm.name = ''
    workerForm.phoneDigits = ''
    workerForm.password = ''
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
  } else {
    productForm.name = ''
    productForm.description = ''
    productForm.quantity_per_pack = 60
    productForm.price_per_pill = 6500
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

onMounted(() => {
  loadProducts()
  loadOrders()
  loadProfile()
  loadWorkers()
  loadFaqs()
  loadSupportThreads()
})
</script>
