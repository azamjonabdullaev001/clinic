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
            <h1 class="text-lg font-bold text-gray-800">{{ txt.title }}</h1>
            <p v-if="authStore.worker" class="text-xs text-gray-400">{{ authStore.worker.name }}</p>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <!-- Language switcher -->
          <div class="flex items-center bg-gray-100 rounded-lg p-0.5 gap-0.5">
            <button @click="lang = 'ru'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'ru' ? 'bg-blue-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">RU</button>
            <button @click="lang = 'uz'" class="text-xs font-semibold px-2.5 py-1 rounded-md transition-all"
              :class="lang === 'uz' ? 'bg-blue-600 text-white shadow-sm' : 'text-gray-500 hover:text-gray-700'">UZ</button>
          </div>
          <button @click="logout" class="text-sm text-red-500 hover:text-red-700 font-medium transition">{{ txt.logout }}</button>
        </div>
      </div>
    </header>

    <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 mt-6 pb-12 space-y-6">

      <!-- ===== OFFLINE (Nurse) Order Section ===== -->
      <div class="bg-white rounded-xl shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b flex items-center gap-3">
          <div class="w-8 h-8 bg-teal-100 rounded-lg flex items-center justify-center">
            <svg class="w-4 h-4 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h2 class="font-bold text-gray-800">{{ txt.nurse_section }}</h2>
        </div>
        <div class="px-6 py-5">
          <p class="text-sm text-gray-500 mb-4">{{ txt.nurse_desc }}</p>
          <div class="flex gap-3">
            <input
              v-model="nurseCode"
              type="text"
              maxlength="5"
              :placeholder="txt.nurse_placeholder"
              class="flex-1 border-2 border-gray-300 rounded-xl px-5 py-4 text-3xl font-bold tracking-[0.4em] text-center focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-teal-400 transition"
              @keyup.enter="searchNurseOrder"
            />
            <button
              @click="searchNurseOrder"
              :disabled="nurseCode.length < 5 || nurseSearching"
              class="bg-teal-600 text-white px-7 py-4 rounded-xl hover:bg-teal-700 transition font-semibold text-base disabled:opacity-40"
            >
              {{ nurseSearching ? txt.searching : txt.find }}
            </button>
          </div>
          <p v-if="nurseSearchError" class="mt-3 text-red-500 text-sm">{{ nurseSearchError }}</p>

          <!-- Found nurse order -->
          <div v-if="nurseOrder" class="mt-5 border-2 border-teal-200 rounded-xl p-5 bg-teal-50">
            <div v-if="nurseConfirmed" class="text-center py-4">
              <svg class="w-14 h-14 text-teal-500 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <p class="text-xl font-bold text-teal-700">{{ txt.payment_success }}</p>
              <button @click="resetNurse" class="mt-4 text-sm text-teal-600 underline hover:no-underline">{{ txt.new_search }}</button>
            </div>
            <template v-else>
              <div class="flex items-start justify-between mb-4">
                <div>
                  <p class="text-2xl font-bold tracking-[0.2em] text-teal-700 mb-1">{{ nurseOrder.order_code }}</p>
                  <p class="font-semibold text-gray-800 text-lg">{{ nurseOrder.patient_first_name }} {{ nurseOrder.patient_last_name }}</p>
                  <p class="text-xs text-gray-400 mt-0.5">{{ new Date(nurseOrder.created_at).toLocaleString('ru-RU') }}</p>
                  <div v-if="nurseOrder.latitude && nurseOrder.longitude" class="mt-1 flex items-center gap-2">
                    <a :href="`https://www.openstreetmap.org/?mlat=${nurseOrder.latitude}&mlon=${nurseOrder.longitude}#map=15/${nurseOrder.latitude}/${nurseOrder.longitude}`"
                      target="_blank" rel="noopener noreferrer"
                      class="inline-flex items-center gap-1 text-xs text-blue-600 hover:text-blue-800 font-medium transition-colors">
                      <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                      {{ txt.view_on_map }}
                    </a>
                  </div>
                </div>
                <div class="text-right">
                  <p class="text-xs text-gray-400">{{ txt.total }}</p>
                  <p class="text-2xl font-bold text-teal-700">{{ formatPrice(orderTotal(nurseOrder)) }} <span class="text-sm font-normal">{{ txt.sum }}</span></p>
                </div>
              </div>

              <!-- Items display / edit mode -->
              <div class="border-t border-teal-200 pt-3 mb-4">
                <div v-if="!editingNurseItems">
                  <!-- View mode -->
                  <div class="space-y-1.5 mb-3">
                    <div v-for="item in nurseOrder.items" :key="item.id" class="flex justify-between text-sm text-gray-700">
                      <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span></span>
                      <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                    </div>
                  </div>
                  <button @click="startEditItems"
                    class="text-xs text-teal-600 hover:text-teal-800 font-medium flex items-center gap-1 transition">
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                    {{ txt.edit_items }}
                  </button>
                </div>

                <!-- Edit mode -->
                <div v-else class="space-y-3">
                  <!-- Existing editable items -->
                  <div v-for="(item, idx) in editItems" :key="idx"
                    class="flex items-center gap-2 bg-white border border-gray-200 rounded-xl px-3 py-2">
                    <span class="flex-1 text-sm font-medium text-gray-800 truncate">{{ item.name }}</span>
                    <input v-model.number="item.quantity" type="number" min="1"
                      class="w-16 border border-gray-300 rounded-lg px-2 py-1 text-sm text-center focus:outline-none focus:ring-1 focus:ring-teal-500" />
                    <select v-model="item.unit_type"
                      class="w-24 border border-gray-300 rounded-lg px-2 py-1 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500">
                      <option value="pack">{{ txt.pack }}</option>
                      <option value="piece">{{ txt.piece }}</option>
                    </select>
                    <button @click="editItems.splice(idx, 1)" class="text-red-400 hover:text-red-600 flex-shrink-0">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                    </button>
                  </div>
                  <!-- Add product row -->
                  <div class="flex gap-2 items-end flex-wrap">
                    <select v-model="editAddProductId"
                      class="flex-1 min-w-[140px] border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500">
                      <option value="">{{ txt.select_product }}</option>
                      <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
                    </select>
                    <input v-model.number="editAddQty" type="number" min="1"
                      class="w-16 border border-gray-300 rounded-lg px-2 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500" />
                    <select v-model="editAddUnit"
                      class="w-24 border border-gray-300 rounded-lg px-2 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500">
                      <option value="pack">{{ txt.pack }}</option>
                      <option value="piece">{{ txt.piece }}</option>
                    </select>
                    <button @click="addEditItem" :disabled="!editAddProductId || editAddQty < 1"
                      class="bg-teal-600 text-white px-3 py-2 rounded-lg text-sm font-medium hover:bg-teal-700 transition disabled:opacity-40">
                      + {{ txt.add }}
                    </button>
                  </div>
                  <!-- Save / cancel edit -->
                  <div class="flex gap-2">
                    <button @click="cancelEditItems"
                      class="flex-1 border border-gray-300 py-2 rounded-lg text-sm font-medium hover:bg-gray-50 transition">
                      {{ txt.cancel_edit }}
                    </button>
                    <button @click="saveEditItems" :disabled="savingEditItems || editItems.length === 0"
                      class="flex-1 bg-teal-600 text-white py-2 rounded-lg text-sm font-bold hover:bg-teal-700 transition disabled:opacity-40">
                      {{ savingEditItems ? txt.saving_items : txt.save_items }}
                    </button>
                  </div>
                </div>
              </div>

              <!-- Name verification -->
              <div v-if="!editingNurseItems" class="bg-white border border-teal-200 rounded-xl p-4 mb-4">
                <p class="text-sm font-medium text-gray-700 mb-2">{{ txt.verify_name }}</p>
                <input v-model="verifyName" type="text"
                  class="w-full border-2 border-gray-200 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500 transition"
                  :placeholder="txt.enter_patient_name" />
                <p v-if="verifyError" class="mt-2 text-red-500 text-sm">{{ verifyError }}</p>
              </div>

              <button v-if="!editingNurseItems" @click="confirmNurseOrder"
                :disabled="!verifyName.trim() || nurseConfirming"
                class="w-full bg-teal-600 text-white py-4 rounded-xl hover:bg-teal-700 transition font-bold text-base disabled:opacity-40">
                {{ nurseConfirming ? txt.confirming : txt.confirm_issue }}
              </button>
            </template>
          </div>
        </div>
      </div>

      <!-- ===== Direct Offline Sale (existing) ===== -->
      <div class="bg-white rounded-xl shadow-sm overflow-hidden">
        <button
          @click="offlineOpen = !offlineOpen"
          class="w-full flex items-center justify-between px-6 py-4 hover:bg-gray-50 transition"
        >
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 bg-emerald-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
            </div>
            <span class="font-semibold text-gray-800">{{ txt.offline_sale }}</span>
          </div>
          <svg class="w-5 h-5 text-gray-400 transition-transform" :class="offlineOpen ? 'rotate-180' : ''" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
          </svg>
        </button>

        <div v-if="offlineOpen" class="border-t px-6 py-5 space-y-4">
          <div class="flex gap-2 flex-wrap items-end">
            <div class="flex-1 min-w-[180px]">
              <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.product }}</label>
              <select v-model="offlineProductId" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500">
                <option value="">{{ txt.select_product }}</option>
                <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
              </select>
            </div>
            <div class="w-24">
              <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.qty }}</label>
              <input v-model.number="offlineQty" type="number" min="1" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500" />
            </div>
            <div class="w-28">
              <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.unit }}</label>
              <select v-model="offlineUnit" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500">
                <option value="pack">{{ txt.pack }}</option>
                <option value="piece">{{ txt.piece }}</option>
              </select>
            </div>
            <button
              @click="addOfflineItem"
              :disabled="!offlineProductId || offlineQty < 1"
              class="bg-emerald-600 text-white px-4 py-2 rounded-lg hover:bg-emerald-700 transition text-sm font-medium disabled:opacity-40"
            >+ {{ txt.add }}</button>
          </div>

          <div v-if="offlineItems.length" class="border rounded-xl overflow-hidden">
            <div v-for="(item, idx) in offlineItems" :key="idx" class="flex items-center justify-between px-4 py-2.5 border-b last:border-0 bg-gray-50">
              <div>
                <span class="font-medium text-gray-800 text-sm">{{ item.name }}</span>
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} {{ item.unit_type === 'pack' ? txt.pack : txt.piece }}</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-semibold text-gray-700 text-sm">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                <button @click="offlineItems.splice(idx, 1)" class="text-red-400 hover:text-red-600 transition">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between px-4 py-2 bg-white">
              <span class="text-sm font-semibold text-gray-700">{{ txt.total }}:</span>
              <span class="font-bold text-emerald-600">{{ formatPrice(offlineItems.reduce((s,i)=>s+i.price,0)) }} {{ txt.sum }}</span>
            </div>
          </div>

          <div class="flex gap-2 flex-wrap">
            <input
              v-model="offlineNote"
              :placeholder="txt.buyer_name"
              class="flex-1 min-w-[200px] border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500"
            />
            <button
              @click="submitOfflineSale"
              :disabled="!offlineItems.length || offlineSubmitting"
              class="bg-emerald-600 text-white px-6 py-2 rounded-lg hover:bg-emerald-700 transition font-medium text-sm disabled:opacity-40"
            >{{ offlineSubmitting ? txt.saving : txt.record_sale }}</button>
          </div>

          <div v-if="offlineSuccess" class="bg-emerald-50 border border-emerald-200 rounded-lg px-4 py-3 text-sm text-emerald-700">
            {{ txt.sale_recorded }} <strong>{{ offlineSuccess }}</strong>
          </div>
        </div>
      </div>

      <!-- ===== Online order search by 6-digit code ===== -->
      <div class="bg-white rounded-xl shadow-sm p-6">
        <h2 class="text-lg font-bold text-gray-800 mb-4">{{ txt.search_online }}</h2>
        <div class="flex gap-3">
          <input
            v-model="searchCode"
            type="text"
            maxlength="6"
            :placeholder="txt.enter_6_code"
            class="flex-1 border border-gray-300 rounded-lg px-4 py-3 text-2xl font-bold tracking-widest text-center focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
            @keyup.enter="searchByCode"
          />
          <button
            @click="searchByCode"
            :disabled="searchCode.length < 6 || searching"
            class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition font-medium disabled:opacity-40"
          >
            {{ searching ? txt.searching : txt.find }}
          </button>
        </div>
        <p v-if="searchError" class="mt-3 text-red-500 text-sm">{{ searchError }}</p>

        <!-- Found online order -->
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
              <div v-if="foundOrder.delivery_address" class="mt-2 flex items-start gap-1.5">
                <svg class="w-4 h-4 text-blue-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                </svg>
                <span class="text-sm text-blue-700 font-medium">{{ foundOrder.delivery_address }}</span>
              </div>
              <div v-if="foundOrder.latitude && foundOrder.longitude" class="mt-1 flex items-center gap-2">
                <a :href="`https://www.openstreetmap.org/?mlat=${foundOrder.latitude}&mlon=${foundOrder.longitude}#map=15/${foundOrder.latitude}/${foundOrder.longitude}`"
                  target="_blank" rel="noopener noreferrer"
                  class="inline-flex items-center gap-1 text-xs text-blue-600 hover:text-blue-800 font-medium transition-colors">
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                  {{ txt.view_on_map }}
                </a>
              </div>
            </div>
            <div class="text-right">
              <p class="text-xs text-gray-400">{{ txt.total }}</p>
              <p class="text-2xl font-bold text-blue-700">{{ formatPrice(orderTotal(foundOrder)) }} <span class="text-sm font-normal">{{ txt.sum }}</span></p>
            </div>
          </div>

          <div class="border-t border-blue-200 pt-4 space-y-2 mb-4">
            <div v-for="item in foundOrder.items" :key="item.id" class="flex justify-between items-center py-1.5 border-b border-blue-100 last:border-0">
              <div>
                <span class="font-medium text-gray-800">{{ item.product?.name }}</span>
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span>
              </div>
              <span class="font-semibold text-gray-700">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
            </div>
          </div>

          <div class="flex gap-2 flex-wrap">
            <button v-if="foundOrder.status === 'pending'" @click="updateStatus(foundOrder, 'confirmed')"
              class="flex-1 bg-blue-600 text-white py-2.5 rounded-lg hover:bg-blue-700 transition font-medium text-sm">✓ {{ txt.confirm }}</button>
            <button v-if="foundOrder.status === 'confirmed' || foundOrder.status === 'shipped'" @click="updateStatus(foundOrder, 'in_transit')"
              class="flex-1 bg-orange-500 text-white py-2.5 rounded-lg hover:bg-orange-600 transition font-medium text-sm">🚚 {{ txt.in_transit }}</button>
            <button v-if="foundOrder.status === 'in_transit'" @click="updateStatus(foundOrder, 'delivered')"
              class="flex-1 bg-green-600 text-white py-2.5 rounded-lg hover:bg-green-700 transition font-medium text-sm">✓ {{ txt.deliver }}</button>
            <button v-if="foundOrder.status !== 'cancelled' && foundOrder.status !== 'delivered'" @click="updateStatus(foundOrder, 'cancelled')"
              class="flex-1 bg-red-50 text-red-600 border border-red-200 py-2.5 rounded-lg hover:bg-red-100 transition font-medium text-sm">{{ txt.cancel }}</button>
            <button @click="openChat(foundOrder)"
              class="bg-indigo-50 text-indigo-600 border border-indigo-200 py-2.5 px-4 rounded-lg hover:bg-indigo-100 transition font-medium text-sm flex items-center gap-1.5">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
              </svg>
              {{ txt.write }}
            </button>
          </div>
        </div>
      </div>

      <!-- ===== All orders list ===== -->
      <div>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-bold text-gray-800">{{ txt.all_orders }}</h2>
          <button @click="loadOrders" class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            {{ txt.refresh }}
          </button>
        </div>

        <div class="flex gap-2 mb-4 flex-wrap">
          <button v-for="f in filters" :key="f.value" @click="activeFilter = f.value"
            :class="activeFilter === f.value ? 'bg-blue-600 text-white' : 'bg-white text-gray-600 hover:bg-gray-50'"
            class="px-4 py-2 rounded-lg text-sm font-medium border border-gray-200 transition">
            {{ f.label }}
            <span v-if="f.value !== 'all'" class="ml-1.5 text-xs opacity-75">({{ orderCountByStatus(f.value) }})</span>
          </button>
        </div>

        <div class="space-y-3">
          <div v-for="order in filteredOrders" :key="order.id"
            class="bg-white rounded-xl shadow-sm p-5 hover:shadow-md transition-shadow">
            <div class="flex flex-col sm:flex-row justify-between items-start gap-3">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-xl font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                  <span v-if="order.is_offline" class="text-xs font-medium px-2 py-0.5 rounded bg-gray-100 text-gray-600">{{ txt.offline_badge }}</span>
                </div>
                <p class="font-semibold text-gray-800">
                  {{ order.is_offline ? order.offline_note : (order.user?.first_name + ' ' + order.user?.last_name) }}
                </p>
                <p class="text-sm text-gray-500" v-if="!order.is_offline">+{{ order.phone }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                <div v-if="order.delivery_address" class="mt-1.5 flex items-start gap-1.5">
                  <svg class="w-3.5 h-3.5 text-orange-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                  </svg>
                  <span class="text-xs text-orange-700">{{ order.delivery_address }}</span>
                </div>
                <div v-if="order.latitude && order.longitude" class="mt-1 flex items-center gap-2">
                  <a :href="`https://www.openstreetmap.org/?mlat=${order.latitude}&mlon=${order.longitude}#map=15/${order.latitude}/${order.longitude}`"
                    target="_blank" rel="noopener noreferrer"
                    class="inline-flex items-center gap-1 text-xs text-blue-600 hover:text-blue-800 font-medium transition-colors">
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                    {{ txt.view_on_map }}
                  </a>
                </div>
              </div>
              <div class="text-right flex-shrink-0">
                <p class="font-bold text-gray-800">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
                <p class="text-xs text-gray-400">{{ order.items?.length }} {{ txt.positions }}</p>
              </div>
            </div>

            <div class="mt-3 pt-3 border-t border-gray-100 space-y-1">
              <div v-for="item in order.items" :key="item.id" class="flex justify-between text-sm text-gray-600">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span></span>
                <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
              </div>
            </div>

            <div class="mt-3 flex gap-2 flex-wrap">
              <button v-if="order.status === 'pending'" @click="updateStatus(order, 'confirmed')"
                class="bg-blue-600 text-white px-4 py-1.5 rounded-lg hover:bg-blue-700 transition text-sm font-medium">{{ txt.confirm }}</button>
              <button v-if="order.status === 'confirmed' || order.status === 'shipped'" @click="updateStatus(order, 'in_transit')"
                class="bg-orange-500 text-white px-4 py-1.5 rounded-lg hover:bg-orange-600 transition text-sm font-medium">🚚 {{ txt.in_transit }}</button>
              <button v-if="order.status === 'in_transit'" @click="updateStatus(order, 'delivered')"
                class="bg-green-600 text-white px-4 py-1.5 rounded-lg hover:bg-green-700 transition text-sm font-medium">✓ {{ txt.deliver }}</button>
              <button v-if="order.status !== 'cancelled' && order.status !== 'delivered'" @click="updateStatus(order, 'cancelled')"
                class="bg-red-50 text-red-600 border border-red-200 px-4 py-1.5 rounded-lg hover:bg-red-100 transition text-sm font-medium">{{ txt.cancel }}</button>
              <button v-if="!order.is_offline" @click="openChat(order)"
                class="bg-indigo-50 text-indigo-600 border border-indigo-200 px-3 py-1.5 rounded-lg hover:bg-indigo-100 transition text-sm font-medium flex items-center gap-1">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
                </svg>
                {{ txt.write }}
              </button>
            </div>
          </div>

          <div v-if="filteredOrders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            {{ txt.no_orders }}
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Chat panel -->
  <div v-if="chatOpen" class="fixed inset-0 z-50 flex items-end sm:items-center justify-center sm:justify-end" @click.self="chatOpen = false">
    <div class="w-full sm:w-96 h-[55vh] sm:h-[70vh] bg-white sm:rounded-tl-2xl sm:rounded-bl-2xl shadow-2xl flex flex-col sm:mr-0 sm:mt-0 rounded-t-2xl">
      <div class="flex items-center justify-between px-4 py-3 border-b bg-indigo-600 rounded-t-2xl sm:rounded-tl-2xl">
        <div>
          <p class="text-white font-semibold text-sm">{{ chatUserName }}</p>
          <p class="text-indigo-200 text-xs">{{ txt.chat_with_client }}</p>
        </div>
        <button @click="chatOpen = false" class="text-white/70 hover:text-white transition">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div ref="chatMessagesEl" class="flex-1 overflow-y-auto p-4 space-y-2 bg-gray-50">
        <div v-if="chatLoading" class="flex justify-center pt-8 text-gray-400 text-sm">{{ txt.loading }}</div>
        <template v-else>
          <div v-if="chatMessages.length === 0" class="text-center text-gray-400 text-sm pt-8">{{ txt.no_messages }}</div>
          <div v-for="msg in chatMessages" :key="msg.id" class="flex"
            :class="msg.sender_role === 'user' ? 'justify-start' : 'justify-end'">
            <div class="max-w-[75%] px-3 py-2 rounded-xl text-sm"
              :class="msg.sender_role === 'user' ? 'bg-white text-gray-800 shadow-sm' : 'bg-indigo-600 text-white'">
              <p class="text-[10px] mb-0.5 opacity-60">{{ msg.sender_role === 'user' ? txt.client : msg.sender_role === 'worker' ? txt.worker : txt.admin }}</p>
              {{ msg.message }}
              <p class="text-[10px] mt-0.5 opacity-50 text-right">{{ new Date(msg.created_at).toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }) }}</p>
            </div>
          </div>
        </template>
      </div>
      <div class="px-3 py-3 border-t flex gap-2">
        <input v-model="chatMsg" @keyup.enter="sendWorkerMessage" :placeholder="txt.type_message"
          class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
        <button @click="sendWorkerMessage" :disabled="!chatMsg.trim() || chatSending"
          class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition disabled:opacity-40">
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

const lang = ref(localStorage.getItem('pickupLang') || 'ru')
const watchLang = () => { localStorage.setItem('pickupLang', lang.value) }

const texts = {
  ru: {
    title: 'Пункт выдачи',
    logout: 'Выйти',
    nurse_section: 'Выдача по коду медсестры (5 цифр)',
    nurse_desc: 'Введите 5-значный код, который пациент получил от медсестры',
    nurse_placeholder: 'XXXXX',
    find: 'Найти',
    searching: 'Поиск...',
    verify_name: 'Подтвердите личность пациента',
    enter_patient_name: 'Введите имя пациента',
    confirm_issue: '✓ Выдать и подтвердить оплату',
    confirming: 'Подтверждение...',
    payment_success: 'Оплата прошла успешно!',
    new_search: 'Новый поиск',
    offline_sale: 'Прямая офлайн продажа',
    search_online: 'Поиск онлайн заказа по коду',
    enter_6_code: 'Введите 6-значный код',
    all_orders: 'Все заказы',
    refresh: 'Обновить',
    product: 'Препарат',
    select_product: 'Выберите препарат',
    qty: 'Кол-во',
    unit: 'Ед.',
    pack: 'упак.',
    piece: 'шт',
    add: 'Добавить',
    total: 'Итого',
    sum: 'сўм',
    buyer_name: 'Имя покупателя (необязательно)',
    saving: 'Запись...',
    record_sale: 'Записать продажу',
    sale_recorded: 'Продажа записана. Код:',
    confirm: 'Подтвердить',
    in_transit: 'В пути',
    deliver: 'Выдать',
    cancel: 'Отменить',
    write: 'Написать',
    offline_badge: 'Офлайн',
    positions: 'позиций',
    no_orders: 'Нет заказов',
    chat_with_client: 'Чат с клиентом',
    loading: 'Загрузка...',
    no_messages: 'Нет сообщений',
    client: 'Клиент',
    worker: 'Работник',
    admin: 'Администратор',
    type_message: 'Введите сообщение...',
    status_pending: 'Ожидает',
    status_confirmed: 'Подтверждён',
    status_shipped: 'Отправлен',
    status_in_transit: 'В пути',
    status_delivered: 'Выдан',
    status_cancelled: 'Отменён',
    verify_error: 'Имя не совпадает. Проверьте данные пациента.',
    edit_items: 'Редактировать товары',
    save_items: 'Сохранить изменения',
    saving_items: 'Сохранение...',
    cancel_edit: 'Отмена',
    view_on_map: 'На карте',
  },
  uz: {
    title: 'Berish punkti',
    logout: 'Chiqish',
    nurse_section: 'Hamshira kodi bo\'yicha berish (5 raqam)',
    nurse_desc: 'Bemor hamshiradan olgan 5 raqamli kodni kiriting',
    nurse_placeholder: 'XXXXX',
    find: 'Topish',
    searching: 'Qidirilmoqda...',
    verify_name: 'Bemorning shaxsini tasdiqlang',
    enter_patient_name: 'Bemorning ismini kiriting',
    confirm_issue: '✓ Berish va to\'lovni tasdiqlash',
    confirming: 'Tasdiqlanmoqda...',
    payment_success: 'To\'lov muvaffaqiyatli o\'tkazildi!',
    new_search: 'Yangi qidiruv',
    offline_sale: 'To\'g\'ridan-to\'g\'ri oflayn sotuv',
    search_online: 'Onlayn buyurtmani kod bo\'yicha qidirish',
    enter_6_code: '6 raqamli kodni kiriting',
    all_orders: 'Barcha buyurtmalar',
    refresh: 'Yangilash',
    product: 'Dori',
    select_product: 'Dori tanlang',
    qty: 'Miqdor',
    unit: 'Birlik',
    pack: 'quti',
    piece: 'dona',
    add: "Qo'shish",
    total: 'Jami',
    sum: "so'm",
    buyer_name: 'Xaridor ismi (ixtiyoriy)',
    saving: 'Saqlanmoqda...',
    record_sale: 'Sotuvni yozish',
    sale_recorded: 'Sotuv yozildi. Kod:',
    confirm: 'Tasdiqlash',
    in_transit: 'Yo\'lda',
    deliver: 'Berish',
    cancel: 'Bekor qilish',
    write: 'Yozish',
    offline_badge: 'Oflayn',
    positions: 'ta mahsulot',
    no_orders: "Buyurtmalar yo'q",
    chat_with_client: 'Mijoz bilan suhbat',
    loading: 'Yuklanmoqda...',
    no_messages: "Xabarlar yo'q",
    client: 'Mijoz',
    worker: 'Xodim',
    admin: 'Administrator',
    type_message: 'Xabar kiriting...',
    status_pending: 'Kutilmoqda',
    status_confirmed: 'Tasdiqlangan',
    status_shipped: 'Yuborilgan',
    status_in_transit: "Yo'lda",
    status_delivered: 'Berildi',
    status_cancelled: 'Bekor qilindi',
    verify_error: "Ism mos kelmadi. Bemorning ma'lumotlarini tekshiring.",
    edit_items: "Mahsulotlarni tahrirlash",
    save_items: "O'zgarishlarni saqlash",
    saving_items: "Saqlanmoqda...",
    cancel_edit: "Bekor qilish",
    view_on_map: "Xaritada ko'rish",
  }
}

const txt = computed(() => {
  watchLang()
  return texts[lang.value] || texts.ru
})

// ===== Orders =====
const orders = ref([])
const activeFilter = ref('all')

const filters = computed(() => [
  { label: txt.value.status_pending.split(' ')[0] === 'Все' ? 'Все' : (lang.value === 'ru' ? 'Все' : 'Barchasi'), value: 'all' },
  { label: txt.value.status_pending, value: 'pending' },
  { label: txt.value.status_confirmed, value: 'confirmed' },
  { label: txt.value.status_shipped, value: 'shipped' },
  { label: txt.value.status_in_transit, value: 'in_transit' },
  { label: txt.value.status_delivered, value: 'delivered' },
  { label: txt.value.status_cancelled, value: 'cancelled' },
])

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
  const t = txt.value
  const m = {
    pending: t.status_pending,
    confirmed: t.status_confirmed,
    shipped: t.status_shipped,
    in_transit: t.status_in_transit,
    delivered: t.status_delivered,
    cancelled: t.status_cancelled,
  }
  return m[status] || status
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
  } catch (e) { console.error(e) }
}

// ===== Nurse order (5-digit) =====
const nurseCode = ref('')
const nurseOrder = ref(null)
const nurseSearchError = ref('')
const nurseSearching = ref(false)
const nurseConfirmed = ref(false)
const nurseConfirming = ref(false)
const verifyName = ref('')
const verifyError = ref('')

// Edit items for nurse/doctor pre-order
const editingNurseItems = ref(false)
const editItems = ref([])
const editAddProductId = ref('')
const editAddQty = ref(1)
const editAddUnit = ref('pack')
const savingEditItems = ref(false)

function startEditItems() {
  editItems.value = (nurseOrder.value.items || []).map(item => ({
    product_id: item.product_id,
    name: item.product?.name || '',
    quantity: item.quantity,
    unit_type: item.unit_type,
  }))
  editAddProductId.value = ''
  editAddQty.value = 1
  editAddUnit.value = 'pack'
  editingNurseItems.value = true
}

function cancelEditItems() {
  editingNurseItems.value = false
}

function addEditItem() {
  if (!editAddProductId.value || editAddQty.value < 1) return
  const product = allProducts.value.find(p => p.id === editAddProductId.value)
  if (!product) return
  editItems.value.push({
    product_id: product.id,
    name: product.name,
    quantity: editAddQty.value,
    unit_type: editAddUnit.value,
  })
  editAddProductId.value = ''
  editAddQty.value = 1
  editAddUnit.value = 'pack'
}

async function saveEditItems() {
  if (!nurseOrder.value || editItems.value.length === 0) return
  savingEditItems.value = true
  try {
    const res = await api.put(`/pickup/orders/${nurseOrder.value.id}/items`, {
      items: editItems.value.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_type: i.unit_type,
      })),
    })
    nurseOrder.value = res.data
    editingNurseItems.value = false
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при сохранении')
  } finally {
    savingEditItems.value = false
  }
}

async function searchNurseOrder() {
  if (nurseCode.value.length < 5) return
  nurseSearchError.value = ''
  nurseOrder.value = null
  nurseConfirmed.value = false
  verifyName.value = ''
  verifyError.value = ''
  nurseSearching.value = true
  try {
    const res = await api.get(`/pickup/nurse-order/${nurseCode.value}`)
    nurseOrder.value = res.data
  } catch (e) {
    nurseSearchError.value = e.response?.data?.error || (lang.value === 'uz' ? "Buyurtma topilmadi" : 'Заказ не найден')
  } finally {
    nurseSearching.value = false
  }
}

async function confirmNurseOrder() {
  if (!nurseOrder.value || !verifyName.value.trim()) return
  verifyError.value = ''

  const expected = (nurseOrder.value.patient_first_name + ' ' + nurseOrder.value.patient_last_name).toLowerCase()
  const entered = verifyName.value.trim().toLowerCase()
  if (!expected.includes(entered) && !entered.includes(expected.split(' ')[0])) {
    verifyError.value = txt.value.verify_error
    return
  }

  nurseConfirming.value = true
  try {
    await api.put(`/pickup/orders/${nurseOrder.value.id}/status`, { status: 'delivered' })
    nurseConfirmed.value = true
    loadOrders()
  } catch (e) {
    verifyError.value = e.response?.data?.error || 'Ошибка'
  } finally {
    nurseConfirming.value = false
  }
}

function resetNurse() {
  nurseCode.value = ''
  nurseOrder.value = null
  nurseConfirmed.value = false
  verifyName.value = ''
  verifyError.value = ''
  nurseSearchError.value = ''
  editingNurseItems.value = false
  editItems.value = []
}

// ===== Online order search (6-digit) =====
const searchCode = ref('')
const foundOrder = ref(null)
const searchError = ref('')
const searching = ref(false)

async function searchByCode() {
  if (searchCode.value.length < 6) return
  searchError.value = ''
  foundOrder.value = null
  searching.value = true
  try {
    const res = await api.get(`/pickup/orders/code/${searchCode.value}`)
    foundOrder.value = res.data
  } catch (e) {
    searchError.value = e.response?.data?.error || (lang.value === 'uz' ? "Buyurtma topilmadi" : 'Заказ не найден')
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

// ===== Chat =====
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
  } catch (e) { console.error(e) }
  finally {
    chatLoading.value = false
    await nextTick()
    scrollChatToBottom()
  }
}

function scrollChatToBottom() {
  if (chatMessagesEl.value) chatMessagesEl.value.scrollTop = chatMessagesEl.value.scrollHeight
}

async function sendWorkerMessage() {
  if (!chatMsg.value.trim() || chatSending.value) return
  if (!chatThreadId.value) { alert(lang.value === 'uz' ? "Foydalanuvchi hali muloqot boshlamagan" : 'Пользователь ещё не начал переписку'); return }
  const text = chatMsg.value.trim()
  chatSending.value = true
  try {
    const res = await api.post(`/pickup/support/threads/${chatThreadId.value}/reply`, { message: text })
    chatMessages.value.push(res.data)
    chatMsg.value = ''
    await nextTick()
    scrollChatToBottom()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при отправке')
  } finally { chatSending.value = false }
}

// ===== Direct Offline Sale =====
const offlineOpen = ref(false)
const allProducts = ref([])
const offlineProductId = ref('')
const offlineQty = ref(1)
const offlineUnit = ref('pack')
const offlineItems = ref([])
const offlineNote = ref('')
const offlineSubmitting = ref(false)
const offlineSuccess = ref('')

async function loadProducts() {
  try {
    const res = await api.get('/products')
    allProducts.value = res.data || []
    for (const p of allProducts.value) p.price_per_pack = p.price_per_pill * p.quantity_per_pack
  } catch (e) { console.error(e) }
}

function addOfflineItem() {
  if (!offlineProductId.value || offlineQty.value < 1) return
  const product = allProducts.value.find(p => p.id === offlineProductId.value)
  if (!product) return
  const price = offlineUnit.value === 'piece'
    ? product.price_per_pill * offlineQty.value
    : product.price_per_pack * offlineQty.value
  offlineItems.value.push({ product_id: product.id, name: product.name, quantity: offlineQty.value, unit_type: offlineUnit.value, price })
  offlineProductId.value = ''
  offlineQty.value = 1
  offlineUnit.value = 'pack'
}

async function submitOfflineSale() {
  if (!offlineItems.value.length) return
  offlineSubmitting.value = true
  offlineSuccess.value = ''
  try {
    const res = await api.post('/pickup/offline-sale', {
      items: offlineItems.value.map(i => ({ product_id: i.product_id, quantity: i.quantity, unit_type: i.unit_type })),
      offline_note: offlineNote.value,
    })
    offlineSuccess.value = res.data.order_code
    offlineItems.value = []
    offlineNote.value = ''
    loadOrders()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при записи')
  } finally { offlineSubmitting.value = false }
}

onMounted(() => {
  loadOrders()
  loadProducts()
})
</script>
