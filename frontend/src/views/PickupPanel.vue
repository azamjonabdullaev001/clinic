<template>
  <div class="min-h-screen flex" :class="{ 'night-mode': night }">

    <!-- ===== SIDEBAR ===== -->
    <aside class="worker-sidebar flex-shrink-0 flex flex-col"
      style="width:200px;min-height:100vh;position:sticky;top:0;height:100vh;overflow-y:auto;background:#111827;border-right:1px solid rgba(255,255,255,0.06);">

      <!-- Brand -->
      <div class="px-4 py-5 flex items-center gap-3"
        style="border-bottom:1px solid rgba(255,255,255,0.06);">
        <div class="w-9 h-9 rounded-xl bg-blue-600 flex items-center justify-center flex-shrink-0 shadow-lg">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/>
          </svg>
        </div>
        <div class="min-w-0">
          <p class="text-white text-sm font-bold leading-tight">{{ txt.title }}</p>
          <p class="text-gray-400 text-xs leading-tight truncate">{{ authStore.worker?.name }}</p>
        </div>
      </div>

      <!-- Nav items -->
      <nav class="flex-1 px-3 py-3 space-y-0.5">
        <button
          v-for="s in [
            {k:'online',    l:txt.nav_online,    d:'M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064'},
            {k:'offline',   l:txt.nav_offline,   d:'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'},
            {k:'stock',     l:txt.nav_stock,     d:'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4'},
            {k:'analytics', l:txt.nav_analytics, d:'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'},
            {k:'history',   l:txt.nav_history,   d:'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'}
          ]"
          :key="s.k"
          @click="tab = s.k"
          :class="tab === s.k
            ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/30'
            : 'text-gray-400 hover:bg-white/5 hover:text-gray-200'"
          class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-all text-left">
          <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" :d="s.d"/>
          </svg>
          {{ s.l }}
        </button>
      </nav>

      <!-- Bottom: lang + night + logout -->
      <div class="px-4 pb-5 pt-3 space-y-3"
        style="border-top:1px solid rgba(255,255,255,0.06);">
        <div class="flex items-center gap-1.5">
          <button @click="lang='ru'"
            :class="lang==='ru' ? 'text-white font-bold' : 'text-gray-500 hover:text-gray-300'"
            class="text-sm transition">RU</button>
          <span class="text-gray-700 text-sm">|</span>
          <button @click="lang='uz'"
            :class="lang==='uz' ? 'text-white font-bold' : 'text-gray-500 hover:text-gray-300'"
            class="text-sm transition">UZ</button>
        </div>
        <button @click="toggleNight"
          class="flex items-center gap-2 text-sm transition w-full"
          :class="night ? 'text-amber-400' : 'text-gray-500 hover:text-gray-300'">
          <svg v-if="!night" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>
          </svg>
          <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
          {{ night ? txt.day_mode : txt.night_mode }}
        </button>
        <button @click="logout"
          class="flex items-center gap-2 text-red-400 hover:text-red-300 text-sm font-medium transition w-full">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
          </svg>
          {{ txt.logout }}
        </button>
      </div>
    </aside>

    <!-- ===== MAIN CONTENT ===== -->
    <div class="flex-1 overflow-auto bg-gray-100">
      <div class="p-6 space-y-6 max-w-5xl mx-auto">

      <!-- ===== OFFLINE (Nurse) Order Section ===== -->
      <div v-show="tab === 'offline'" class="bg-white rounded-xl shadow-sm overflow-hidden">
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
                    <div v-for="item in boughtItems(nurseOrder)" :key="item.id" class="flex justify-between text-sm text-gray-700">
                      <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ txt.pack }}</span></span>
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
                    <span class="text-xs text-gray-500 w-14 text-center">{{ txt.pack }}</span>
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
                    <span class="text-xs text-gray-500 w-14 text-center">{{ txt.pack }}</span>
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
      <div v-show="tab === 'offline'" class="bg-white rounded-xl shadow-sm overflow-hidden">
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
                <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }} — {{ capsulesOf(p.id) }} {{ txt.pack }} / {{ stockOf(p.id) }} {{ txt.piece }}</option>
              </select>
            </div>
            <div class="w-20">
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
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-semibold text-gray-700 text-sm">{{ formatPrice(saleType === 'vip' ? 0 : item.price) }} {{ txt.sum }}</span>
                <button @click="offlineItems.splice(idx, 1)" class="text-red-400 hover:text-red-600 transition">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between px-4 py-2 bg-white">
              <span class="text-sm font-semibold text-gray-700">{{ txt.total }}:</span>
              <span class="font-bold text-emerald-600">{{ formatPrice(saleType === 'vip' ? 0 : offlineTotal) }} {{ txt.sum }}</span>
            </div>
          </div>

          <!-- Sale type + payment method -->
          <div v-if="offlineItems.length" class="border rounded-xl px-4 py-3 space-y-3 bg-white">
            <div>
              <label class="block text-xs font-medium text-gray-500 mb-1.5">{{ txt.sale_type }}</label>
              <div class="flex gap-2 flex-wrap">
                <button @click="saleType = 'regular'"
                  :class="saleType === 'regular' ? 'bg-emerald-600 text-white border-emerald-600' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.sale_regular }}</button>
                <button @click="saleType = 'vip'"
                  :class="saleType === 'vip' ? 'bg-amber-500 text-white border-amber-500' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.vip_free }}</button>
                <button @click="saleType = 'marketolog'"
                  :class="saleType === 'marketolog' ? 'bg-purple-600 text-white border-purple-600' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.sale_marketolog }}</button>
              </div>
            </div>

            <!-- Marketolog select -->
            <div v-if="saleType === 'marketolog'">
              <label class="block text-xs font-medium text-gray-500 mb-1.5">{{ txt.choose_marketolog }}</label>
              <select v-model="offlineMarketolog" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500">
                <option :value="null">{{ txt.choose_marketolog }}</option>
                <option v-for="m in marketologs" :key="m.id" :value="m.id">{{ m.name }}</option>
              </select>
              <p v-if="marketologs.length === 0" class="text-xs text-gray-400 mt-1">{{ txt.no_marketologs }}</p>
            </div>

            <!-- Payment method (regular sales only) -->
            <div v-if="saleType === 'regular'">
              <label class="block text-xs font-medium text-gray-500 mb-1.5">{{ txt.payment_method }}</label>
              <div class="flex gap-2">
                <button v-for="pm in paymentMethods" :key="pm.value" @click="offlinePaymentMethod = pm.value"
                  :class="offlinePaymentMethod === pm.value ? 'bg-emerald-600 text-white border-emerald-600' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="flex-1 px-3 py-2 rounded-lg text-sm font-medium border transition">{{ pm.label }}</button>
              </div>
            </div>
          </div>

          <div v-if="offlineItems.length && saleType !== 'marketolog'">
            <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.referral_label }}</label>
            <input
              v-model="offlineReferral"
              list="offline-doctors"
              :placeholder="txt.referral_ph"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500"
            />
            <datalist id="offline-doctors">
              <option value="Самостоятельно"></option>
              <option v-for="d in allDoctors" :key="d.id" :value="d.name + (d.specialty ? ' (' + d.specialty + ')' : '')"></option>
            </datalist>
          </div>

          <div class="flex gap-2 flex-wrap">
            <input
              v-if="saleType !== 'marketolog'"
              v-model="offlineNote"
              :placeholder="txt.buyer_name"
              class="flex-1 min-w-[200px] border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500"
            />
            <button
              @click="submitOfflineSale"
              :disabled="!offlineCanSubmit || offlineSubmitting"
              class="bg-emerald-600 text-white px-6 py-2 rounded-lg hover:bg-emerald-700 transition font-medium text-sm disabled:opacity-40"
            >{{ offlineSubmitting ? txt.saving : txt.record_sale }}</button>
          </div>

          <div v-if="offlineSuccess" class="bg-emerald-50 border border-emerald-200 rounded-lg px-4 py-3 text-sm text-emerald-700 flex items-center gap-2">
            <svg class="w-5 h-5 text-emerald-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
            {{ txt.sold_ok }}
          </div>
        </div>
      </div>

      <!-- ===== My analytics ===== -->
      <div v-show="tab === 'analytics'" class="bg-white rounded-xl shadow-sm overflow-hidden">
        <div class="px-6 py-4 border-b flex items-center gap-3">
          <div class="w-8 h-8 bg-indigo-100 rounded-lg flex items-center justify-center">
            <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg>
          </div>
          <span class="font-semibold text-gray-800">{{ txt.analytics }}</span>
        </div>
        <div class="px-6 py-5 space-y-4">
          <div class="flex gap-2 flex-wrap items-center">
            <button v-for="p in [{v:'daily',l:txt.a_today},{v:'weekly',l:txt.a_week},{v:'monthly',l:txt.a_month},{v:'custom',l:txt.a_date}]" :key="p.v"
              @click="selectAnalyticsPeriod(p.v)"
              :class="analyticsPeriod===p.v ? 'bg-indigo-600 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'"
              class="px-4 py-2 rounded-lg text-sm font-medium transition">{{ p.l }}</button>
            <input v-if="analyticsPeriod==='custom'" v-model="analyticsDate" type="date" @change="loadAnalytics"
              class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
          <div v-if="analyticsLoading" class="text-gray-400 text-sm py-6 text-center">{{ txt.loading }}</div>
          <template v-else-if="analyticsData">
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">{{ txt.a_orders }}</p><p class="text-2xl font-bold text-gray-800">{{ analyticsData.total_orders }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">{{ txt.a_revenue }}</p><p class="text-lg font-bold text-emerald-600">{{ formatPrice(analyticsData.total_revenue) }} {{ txt.sum }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">{{ txt.a_created }}</p><p class="text-2xl font-bold text-blue-600">{{ analyticsData.created_count }}</p></div>
              <div class="bg-gray-50 rounded-xl p-4"><p class="text-xs text-gray-500 mb-1">{{ txt.a_confirmed }}</p><p class="text-2xl font-bold text-teal-600">{{ analyticsData.confirmed_count }}</p></div>
            </div>
            <div class="pt-2">
              <LineChart :points="analyticsData.points || []" color="#6366f1" />
            </div>

            <!-- Breakdown by customer category -->
            <div class="border-t pt-4 mt-2">
              <div class="flex items-center justify-between flex-wrap gap-2 mb-3">
                <h3 class="font-semibold text-gray-800 text-sm">{{ txt.by_category }}</h3>
                <select v-model="analyticsCat" class="border border-gray-300 rounded-lg px-3 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
                  <option value="all">{{ txt.cat_all }}</option>
                  <option value="vip">{{ txt.cat_vip }}</option>
                  <option value="doctor">{{ txt.cat_doctor }}</option>
                  <option value="marketolog">{{ txt.cat_marketolog }}</option>
                </select>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-3">
                <div v-for="cat in shownCats" :key="cat.key" class="bg-gray-50 rounded-xl p-4">
                  <div class="flex items-center justify-between mb-2">
                    <span class="text-sm font-semibold" :class="cat.color">{{ cat.label }}</span>
                    <span class="text-xs text-gray-400">{{ catData(cat.key).orders }} {{ txt.a_orders.toLowerCase() }}</span>
                  </div>
                  <div class="flex flex-wrap gap-x-4 gap-y-1 text-sm text-gray-600">
                    <span>{{ txt.pack }}: <b>{{ catData(cat.key).capsules }}</b></span>
                    <span>{{ txt.piece }}: <b>{{ catData(cat.key).pieces }}</b></span>
                    <span class="text-emerald-600">{{ formatPrice(catData(cat.key).revenue) }} {{ txt.sum }}</span>
                  </div>
                </div>
              </div>

              <!-- Per-doctor -->
              <div v-if="(analyticsCat === 'all' || analyticsCat === 'doctor') && (analyticsData.by_doctor || []).length" class="mt-4 overflow-x-auto rounded-xl border border-gray-100">
                <table class="w-full text-sm">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-gray-500 uppercase">{{ txt.cat_doctor }}</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-gray-500 uppercase">{{ txt.pack }}</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-gray-500 uppercase">{{ txt.piece }}</th>
                      <th class="text-right px-4 py-2 text-xs font-semibold text-gray-500 uppercase">{{ txt.a_revenue }}</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100">
                    <tr v-for="(d, i) in analyticsData.by_doctor" :key="i">
                      <td class="px-4 py-2 font-medium text-gray-800">{{ d.name }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ d.capsules }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ d.pieces }}</td>
                      <td class="px-4 py-2 text-right font-bold text-gray-700">{{ formatPrice(d.revenue) }} {{ txt.sum }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Per-marketolog -->
              <div v-if="(analyticsCat === 'all' || analyticsCat === 'marketolog') && (analyticsData.by_marketolog || []).length" class="mt-4 overflow-x-auto rounded-xl border border-purple-100">
                <table class="w-full text-sm">
                  <thead class="bg-purple-50">
                    <tr>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">{{ txt.cat_marketolog }}</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">{{ txt.pack }}</th>
                      <th class="text-left px-4 py-2 text-xs font-semibold text-purple-700 uppercase">{{ txt.piece }}</th>
                      <th class="text-right px-4 py-2 text-xs font-semibold text-purple-700 uppercase">{{ txt.a_revenue }}</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-purple-50">
                    <tr v-for="(m, i) in analyticsData.by_marketolog" :key="i">
                      <td class="px-4 py-2 font-medium text-gray-800">{{ m.name }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ m.capsules }}</td>
                      <td class="px-4 py-2 text-gray-600">{{ m.pieces }}</td>
                      <td class="px-4 py-2 text-right font-bold text-purple-700">{{ formatPrice(m.revenue) }} {{ txt.sum }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </template>
        </div>
      </div>

      <!-- ===== Warehouse / Склад ===== -->
      <div v-show="tab === 'stock'" class="space-y-6">
        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b"><h2 class="font-bold text-gray-800">{{ txt.stock_in }}</h2></div>
          <div class="px-6 py-5">
            <div class="flex gap-2 flex-wrap items-end">
              <div class="flex-1 min-w-[180px]">
                <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.product }}</label>
                <select v-model="stockProductId" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="">{{ txt.select_product }}</option>
                  <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
                </select>
              </div>
              <div class="w-24">
                <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.stock_qty }}</label>
                <input v-model.number="stockQty" type="number" min="1" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500" />
              </div>
              <div class="w-28">
                <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.unit }}</label>
                <select v-model="stockUnit" class="w-full border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="pack">{{ txt.pack }}</option>
                  <option value="piece">{{ txt.piece }}</option>
                </select>
              </div>
              <button @click="addStock" :disabled="!stockProductId || stockQty < 1 || addingStock"
                class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm font-medium disabled:opacity-40">+ {{ txt.stock_add }}</button>
            </div>
          </div>
        </div>
        <div class="bg-white rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b"><h2 class="font-bold text-gray-800">{{ txt.my_stock }}</h2></div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50 border-b">
                <tr>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">{{ txt.product }}</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">{{ txt.unit }}</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">{{ txt.total }}</th>
                  <th class="text-right px-5 py-3 text-xs font-semibold text-gray-500 uppercase tracking-wider">{{ txt.my_stock }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="s in stock" :key="s.id" class="hover:bg-gray-50 transition">
                  <td class="px-5 py-3">
                    <div class="flex items-center gap-3">
                      <div class="w-10 h-10 bg-gray-100 rounded-lg overflow-hidden flex items-center justify-center flex-shrink-0">
                        <img v-if="s.product?.image_path" :src="s.product.image_path" class="w-full h-full object-cover" />
                        <svg v-else class="w-5 h-5 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                      </div>
                      <span class="font-medium text-gray-800">{{ s.product?.name }}</span>
                    </div>
                  </td>
                  <td class="px-5 py-3 text-gray-500 text-sm">{{ s.product?.quantity_per_pack }} {{ txt.piece }}</td>
                  <td class="px-5 py-3 text-gray-600 text-sm">{{ formatPrice(s.product?.price_per_pack) }} {{ txt.sum }}</td>
                  <td class="px-5 py-3 text-right">
                    <span :class="stockOf(s.product_id) > 0 ? 'text-green-700 bg-green-50' : 'text-red-600 bg-red-50'" class="px-2.5 py-1 rounded-lg text-sm font-bold">
                      {{ capsulesOf(s.product_id) }} {{ txt.pack }} / {{ stockOf(s.product_id) }} {{ txt.piece }}
                    </span>
                  </td>
                </tr>
                <tr v-if="stock.length === 0">
                  <td colspan="4" class="px-5 py-10 text-center text-gray-400 text-sm">{{ txt.stock_empty }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- ===== Online order search by 6-digit code ===== -->
      <div v-show="tab === 'online'" class="bg-white rounded-xl shadow-sm p-6">
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

      <!-- ===== Orders list (shared by online / offline / history) ===== -->
      <div v-show="tab !== 'analytics'">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-bold text-gray-800">{{ listHeading }}</h2>
          <button @click="loadOrders" class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
            {{ txt.refresh }}
          </button>
        </div>

        <!-- History filters: type + status + period (dropdowns) -->
        <div v-if="tab === 'history'" class="flex gap-3 flex-wrap mb-4">
          <div>
            <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.type_all }}</label>
            <select v-model="historyType" class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 min-w-[150px]">
              <option value="all">{{ txt.type_all }}</option>
              <option value="online">{{ txt.type_online }}</option>
              <option value="offline">{{ txt.type_offline }}</option>
              <option value="vip">{{ txt.own_patient }}</option>
            </select>
          </div>
          <div>
            <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.status_pending }}</label>
            <select v-model="historyStatus" class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 min-w-[150px]">
              <option v-for="f in statusFilters" :key="f.value" :value="f.value">{{ f.label }}</option>
            </select>
          </div>
          <div>
            <label class="block text-xs font-medium text-gray-500 mb-1">{{ txt.a_date }}</label>
            <select v-model="historyPeriod" class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 min-w-[130px]">
              <option value="all">{{ txt.period_all }}</option>
              <option value="daily">{{ txt.a_today }}</option>
              <option value="weekly">{{ txt.a_week }}</option>
              <option value="monthly">{{ txt.a_month }}</option>
              <option value="custom">{{ txt.a_date }}</option>
            </select>
          </div>
          <div v-if="historyPeriod === 'custom'">
            <label class="block text-xs font-medium text-gray-500 mb-1">&nbsp;</label>
            <input v-model="historyDate" type="date" class="border border-gray-300 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
          </div>
        </div>

        <div class="space-y-3">
          <div v-for="order in displayedOrders" :key="order.id"
            class="bg-white rounded-xl shadow-sm p-5 hover:shadow-md transition-shadow">
            <div class="flex flex-col sm:flex-row justify-between items-start gap-3">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-xl font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                  <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                  <span v-if="order.is_offline" class="text-xs font-medium px-2 py-0.5 rounded bg-gray-100 text-gray-600">{{ txt.offline_badge }}</span>
                  <span v-if="paymentLabel(order)" :class="paymentBadgeClass(order)" class="text-xs font-medium px-2 py-0.5 rounded">{{ paymentLabel(order) }}</span>
                </div>
                <p class="font-semibold text-gray-800">
                  {{ order.is_offline ? order.offline_note : (order.user?.first_name + ' ' + order.user?.last_name) }}
                </p>
                <p class="text-sm text-gray-500" v-if="!order.is_offline">+{{ order.phone }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                <div v-if="order.referred_by" class="mt-1 flex items-center gap-1">
                  <svg class="w-3.5 h-3.5 text-purple-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                  <span class="text-xs text-purple-700">{{ txt.referred_by }}: {{ order.referred_by }}</span>
                </div>
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

            <!-- Items: view or edit mode -->
            <div class="mt-3 pt-3 border-t border-gray-100">
              <!-- VIEW MODE -->
              <div v-if="!listEdit[order.id]?.editing" class="space-y-1">
                <div v-for="item in boughtItems(order)" :key="item.id" class="flex justify-between text-sm text-gray-600">
                  <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ txt.pack }}</span></span>
                  <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                </div>
              </div>

              <!-- EDIT MODE -->
              <div v-else class="space-y-2">
                <!-- Existing items with +/- -->
                <div v-for="(ei, idx) in listEdit[order.id].items" :key="idx"
                  class="flex items-center gap-2 bg-gray-50 border border-gray-200 rounded-xl px-3 py-2">
                  <span class="flex-1 text-sm font-medium text-gray-800 truncate">{{ ei.name }}</span>
                  <div class="flex items-center gap-1">
                    <button @click="listEditDec(order.id, idx)"
                      class="w-7 h-7 bg-gray-200 hover:bg-gray-300 rounded-lg text-gray-700 font-bold text-sm flex items-center justify-center transition">−</button>
                    <span class="w-8 text-center text-sm font-bold">{{ ei.quantity }}</span>
                    <button @click="listEditInc(order.id, idx)"
                      class="w-7 h-7 bg-blue-100 hover:bg-blue-200 rounded-lg text-blue-700 font-bold text-sm flex items-center justify-center transition">+</button>
                  </div>
                  <span class="text-xs text-gray-500 w-14 text-center">{{ txt.pack }}</span>
                  <button @click="listEdit[order.id].items.splice(idx, 1)" class="text-red-400 hover:text-red-600 flex-shrink-0">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                  </button>
                </div>

                <!-- Add new product row -->
                <div class="flex gap-2 items-end flex-wrap pt-1">
                  <select v-model="listEdit[order.id].addProductId"
                    class="flex-1 min-w-[140px] border border-gray-300 rounded-lg px-2 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500">
                    <option value="">{{ txt.select_product }}</option>
                    <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
                  </select>
                  <input v-model.number="listEdit[order.id].addQty" type="number" min="1"
                    class="w-16 border border-gray-300 rounded-lg px-2 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500" />
                  <span class="text-xs text-gray-500 px-1">{{ txt.pack }}</span>
                  <button @click="listEditAddItem(order.id)" :disabled="!listEdit[order.id].addProductId || listEdit[order.id].addQty < 1"
                    class="bg-teal-600 text-white px-3 py-2 rounded-lg text-sm font-medium hover:bg-teal-700 transition disabled:opacity-40">
                    + {{ txt.add }}
                  </button>
                </div>

                <!-- Save / Cancel -->
                <div class="flex gap-2 pt-1">
                  <button @click="cancelListEdit(order.id)"
                    class="flex-1 border border-gray-300 py-2 rounded-lg text-sm font-medium hover:bg-gray-50 transition">
                    {{ txt.cancel_edit }}
                  </button>
                  <button @click="saveListEdit(order)" :disabled="listEdit[order.id].saving || listEdit[order.id].items.length === 0"
                    class="flex-1 bg-teal-600 text-white py-2 rounded-lg text-sm font-bold hover:bg-teal-700 transition disabled:opacity-40">
                    {{ listEdit[order.id].saving ? txt.saving_items : txt.save_items }}
                  </button>
                </div>
              </div>
            </div>

            <div class="mt-3 flex gap-2 flex-wrap">
              <!-- Edit items: offline orders (incl. delivered, which becomes a return) -->
              <button v-if="order.is_offline && !order.marketolog_id && order.status !== 'cancelled' && !listEdit[order.id]?.editing"
                @click="startListEdit(order)"
                :class="order.status === 'delivered' ? 'bg-red-50 text-red-600 border-red-200 hover:bg-red-100' : 'bg-gray-100 text-gray-700 border-gray-200 hover:bg-gray-200'"
                class="border px-4 py-1.5 rounded-lg transition text-sm font-medium flex items-center gap-1.5">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                {{ order.status === 'delivered' ? txt.return_edit : txt.edit_items }}
              </button>
              <button v-if="order.is_offline && !order.marketolog_id && order.status === 'delivered' && !listEdit[order.id]?.editing"
                @click="fullReturn(order)"
                class="bg-red-600 text-white border border-red-600 px-4 py-1.5 rounded-lg hover:bg-red-700 transition text-sm font-medium flex items-center gap-1.5">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"/></svg>
                {{ txt.full_return }}
              </button>
              <template v-if="!listEdit[order.id]?.editing">
                <!-- Offline order: one-tap confirm → delivered immediately -->
                <button v-if="order.is_offline && order.status === 'pending'" @click="askPayment(order, 'list')"
                  class="bg-green-600 text-white px-4 py-1.5 rounded-lg hover:bg-green-700 transition text-sm font-medium">✓ {{ txt.confirm }}</button>
                <!-- Online order: normal status flow -->
                <button v-if="!order.is_offline && order.status === 'pending'" @click="updateStatus(order, 'confirmed')"
                  class="bg-blue-600 text-white px-4 py-1.5 rounded-lg hover:bg-blue-700 transition text-sm font-medium">{{ txt.confirm }}</button>
                <button v-if="!order.is_offline && (order.status === 'confirmed' || order.status === 'shipped')" @click="updateStatus(order, 'in_transit')"
                  class="bg-orange-500 text-white px-4 py-1.5 rounded-lg hover:bg-orange-600 transition text-sm font-medium">🚚 {{ txt.in_transit }}</button>
                <button v-if="!order.is_offline && order.status === 'in_transit'" @click="updateStatus(order, 'delivered')"
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
              </template>
            </div>
          </div>

          <div v-if="displayedOrders.length === 0" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">
            <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            {{ txt.no_orders }}
          </div>
        </div>
      </div>
      </div>
    </div>
  </div>

  <!-- Payment method modal (offline order confirmation) -->
  <div v-if="showPayModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4" :class="{ 'night-mode': night }" @click.self="closePayModal">
    <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
    <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6">
      <div class="text-center mb-5">
        <div class="w-12 h-12 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/></svg>
        </div>
        <h3 class="text-lg font-bold text-gray-800">{{ txt.choose_payment }}</h3>
        <p class="text-sm text-gray-500 mt-1">{{ txt.payment_hint }}</p>
        <p v-if="payOrder" class="text-emerald-700 font-bold mt-2">{{ formatPrice(orderTotal(payOrder)) }} {{ txt.sum }}</p>
      </div>
      <div class="grid gap-2">
        <button v-for="pm in paymentMethods" :key="pm.value" @click="confirmPayment(pm.value)" :disabled="paySubmitting"
          class="w-full py-3 rounded-xl border-2 border-emerald-200 text-emerald-700 font-semibold hover:bg-emerald-50 transition disabled:opacity-40">
          {{ pm.label }}
        </button>
      </div>
      <button @click="closePayModal" :disabled="paySubmitting"
        class="w-full mt-3 py-2 text-sm text-gray-400 hover:text-gray-600 transition disabled:opacity-40">
        {{ txt.cancel }}
      </button>
    </div>
  </div>

  <!-- Chat panel -->
  <div v-if="chatOpen" class="fixed inset-0 z-50 flex items-end sm:items-center justify-center sm:justify-end" :class="{ 'night-mode': night }" @click.self="chatOpen = false">
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
import { ref, computed, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'
import { useNight } from '../stores/night'
import { useStockSocket, displayStock, realtime } from '../stores/stock'
import LineChart from '../components/LineChart.vue'

const authStore = useAuthStore()
useStockSocket()
const router = useRouter()
const { night, toggle: toggleNight } = useNight()

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
    pack: 'капс.',
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
    return_edit: 'Возврат',
    full_return: 'Полный возврат',
    save_items: 'Сохранить изменения',
    saving_items: 'Сохранение...',
    cancel_edit: 'Отмена',
    view_on_map: 'На карте',
    vip_free: 'Свой пациент (бесплатно)',
    sale_type: 'Тип продажи',
    sale_regular: 'Обычная',
    sale_marketolog: 'Маркетолог (долг)',
    choose_marketolog: 'Выберите маркетолога',
    no_marketologs: 'Нет маркетологов. Добавьте работника с ролью «Менеджер» в админ-панели.',
    payment_method: 'Способ оплаты',
    pay_cash: 'Наличные',
    pay_terminal: 'Терминал',
    pay_card: 'Карта',
    pay_online: 'Онлайн (карта)',
    payment_label: 'Оплата',
    vip_badge: 'Свой пациент · Бесплатно',
    choose_payment: 'Выберите способ оплаты',
    payment_hint: 'Как клиент оплатил этот заказ?',
    referral_label: 'Откуда пациент / доктор',
    referral_ph: 'Самостоятельно или имя доктора',
    referred_by: 'Рекомендовал',
    analytics: 'Моя аналитика',
    a_today: 'Сегодня',
    a_week: 'Неделя',
    a_month: 'Месяц',
    a_date: 'Дата',
    a_orders: 'Заказов',
    a_revenue: 'Выручка',
    a_created: 'Создано',
    a_confirmed: 'Подтверждено',
    by_category: 'По категориям',
    cat_all: 'Все категории',
    cat_vip: 'Свой пациент',
    cat_doctor: 'От доктора',
    cat_marketolog: 'Маркетолог',
    cat_regular: 'Обычные',
    sold_ok: 'Продано ✓ Продажа записана',
    nav_online: 'Онлайн',
    nav_offline: 'Офлайн',
    nav_analytics: 'Аналитика',
    nav_history: 'История',
    nav_stock: 'Склад',
    stock_in: 'Приход товара',
    my_stock: 'Мой склад',
    stock_qty: 'Кол-во (капс.)',
    stock_add: 'Пополнить',
    stock_empty: 'Склад пуст',
    pending_online: 'Ожидаемые онлайн-заказы',
    pending_offline: 'Ожидаемые офлайн-заказы',
    history_title: 'История заказов',
    type_all: 'Все',
    type_online: 'Онлайн',
    type_offline: 'Офлайн',
    own_patient: 'Свой пациент',
    period_all: 'Всё время',
    night_mode: 'Ночной',
    day_mode: 'Дневной',
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
    pack: 'kapsula',
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
    return_edit: 'Qaytarish',
    full_return: "To'liq qaytarish",
    save_items: "O'zgarishlarni saqlash",
    saving_items: "Saqlanmoqda...",
    cancel_edit: "Bekor qilish",
    view_on_map: "Xaritada ko'rish",
    vip_free: "O'z bemori (bepul)",
    sale_type: 'Sotuv turi',
    sale_regular: 'Oddiy',
    sale_marketolog: 'Marketolog (qarz)',
    choose_marketolog: 'Marketologni tanlang',
    no_marketologs: "Marketolog yo'q. Admin-panelda «Menejer» rolidagi xodim qo'shing.",
    payment_method: "To'lov usuli",
    pay_cash: 'Naqd',
    pay_terminal: 'Terminal',
    pay_card: 'Karta',
    pay_online: 'Onlayn (karta)',
    payment_label: "To'lov",
    vip_badge: "O'z bemori · Bepul",
    choose_payment: "To'lov usulini tanlang",
    payment_hint: "Mijoz buyurtmani qanday to'ladi?",
    referral_label: 'Bemor qayerdan / shifokor',
    referral_ph: 'Mustaqil yoki shifokor ismi',
    referred_by: 'Tavsiya qildi',
    analytics: 'Mening tahlilim',
    a_today: 'Bugun',
    a_week: 'Hafta',
    a_month: 'Oy',
    a_date: 'Sana',
    a_orders: 'Buyurtmalar',
    a_revenue: 'Tushum',
    a_created: 'Yaratilgan',
    a_confirmed: 'Tasdiqlangan',
    by_category: 'Toifalar bo\'yicha',
    cat_all: 'Barcha toifalar',
    cat_vip: 'O\'z bemori',
    cat_doctor: 'Shifokordan',
    cat_marketolog: 'Marketolog',
    cat_regular: 'Oddiy',
    sold_ok: 'Sotildi ✓ Sotuv yozildi',
    nav_online: 'Onlayn',
    nav_offline: 'Oflayn',
    nav_analytics: 'Tahlil',
    nav_history: 'Tarix',
    nav_stock: 'Ombor',
    stock_in: 'Tovar kirimi',
    my_stock: 'Mening omborim',
    stock_qty: 'Soni (kaps.)',
    stock_add: "To'ldirish",
    stock_empty: "Ombor bo'sh",
    pending_online: 'Kutilayotgan onlayn buyurtmalar',
    pending_offline: 'Kutilayotgan oflayn buyurtmalar',
    history_title: 'Buyurtmalar tarixi',
    type_all: 'Hammasi',
    type_online: 'Onlayn',
    type_offline: 'Oflayn',
    own_patient: "O'z bemori",
    period_all: 'Butun davr',
    night_mode: 'Tungi',
    day_mode: 'Kunduzgi',
  }
}

const txt = computed(() => {
  watchLang()
  return texts[lang.value] || texts.ru
})

// ===== Orders & sections =====
const orders = ref([])
const tab = ref('online') // online | offline | analytics | history

// History filters
const historyType = ref('all')   // all | online | offline
const historyStatus = ref('all') // all | pending | confirmed | shipped | in_transit | delivered | cancelled
const historyPeriod = ref('all') // all | daily | weekly | monthly | custom
const historyDate = ref('')

const statusFilters = computed(() => [
  { label: txt.value.type_all, value: 'all' },
  { label: txt.value.status_pending, value: 'pending' },
  { label: txt.value.status_confirmed, value: 'confirmed' },
  { label: txt.value.status_in_transit, value: 'in_transit' },
  { label: txt.value.status_delivered, value: 'delivered' },
  { label: txt.value.status_cancelled, value: 'cancelled' },
])

function inPeriod(order) {
  if (historyPeriod.value === 'all') return true
  const d = new Date(order.created_at)
  const now = new Date()
  if (historyPeriod.value === 'daily') return d.toDateString() === now.toDateString()
  if (historyPeriod.value === 'weekly') return d >= new Date(now.getTime() - 7 * 864e5)
  if (historyPeriod.value === 'monthly') return d >= new Date(now.getTime() - 30 * 864e5)
  if (historyPeriod.value === 'custom') return historyDate.value ? d.toDateString() === new Date(historyDate.value).toDateString() : true
  return true
}

const displayedOrders = computed(() => {
  if (tab.value === 'online') {
    return orders.value.filter(o => !o.is_offline && o.status !== 'delivered' && o.status !== 'cancelled')
  }
  if (tab.value === 'offline') {
    return orders.value.filter(o => o.is_offline && o.status !== 'delivered' && o.status !== 'cancelled')
  }
  // history
  let list = orders.value
  if (historyType.value === 'online') list = list.filter(o => !o.is_offline)
  else if (historyType.value === 'offline') list = list.filter(o => o.is_offline)
  else if (historyType.value === 'vip') list = list.filter(o => o.is_vip) // own patients — free
  if (historyStatus.value !== 'all') list = list.filter(o => o.status === historyStatus.value)
  return list.filter(inPeriod)
})

const listHeading = computed(() => {
  if (tab.value === 'online') return txt.value.pending_online
  if (tab.value === 'offline') return txt.value.pending_offline
  return txt.value.history_title
})

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

// Only the capsules actually bought (quantity > 0); removed ones are kept with quantity 0.
function boughtItems(order) {
  return (order.items || []).filter(i => i.quantity > 0)
}

function paymentLabel(order) {
  if (order.is_vip) return txt.value.vip_badge
  const t = txt.value
  const m = {
    cash: t.pay_cash,
    terminal: t.pay_terminal,
    card: t.pay_card,
    online: t.pay_online,
  }
  return m[order.payment_method] || ''
}

function paymentBadgeClass(order) {
  if (order.is_vip) return 'bg-amber-100 text-amber-700'
  if (order.payment_method === 'cash') return 'bg-green-100 text-green-700'
  if (order.payment_method === 'online') return 'bg-indigo-100 text-indigo-700'
  return 'bg-blue-100 text-blue-700'
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
  editItems.value = (nurseOrder.value.items || []).filter(item => item.quantity > 0).map(item => ({
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
    unit_type: 'pack',
  })
  editAddProductId.value = ''
  editAddQty.value = 1
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

function confirmNurseOrder() {
  if (!nurseOrder.value || !verifyName.value.trim()) return
  verifyError.value = ''

  const expected = (nurseOrder.value.patient_first_name + ' ' + nurseOrder.value.patient_last_name).toLowerCase()
  const entered = verifyName.value.trim().toLowerCase()
  if (!expected.includes(entered) && !entered.includes(expected.split(' ')[0])) {
    verifyError.value = txt.value.verify_error
    return
  }

  // Require choosing a payment method before the order is finalized.
  askPayment(nurseOrder.value, 'nurse')
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
  const payload = { status }
  if (status === 'cancelled') {
    const reason = prompt('Укажите причину отмены заказа:')
    if (reason === null) return
    const trimmed = reason.trim()
    if (!trimmed) {
      alert('Причина обязательна')
      return
    }
    payload.cancellation_reason = trimmed
  }
  try {
    const res = await api.put(`/pickup/orders/${order.id}/status`, payload)
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx !== -1) orders.value[idx] = res.data
    if (foundOrder.value?.id === order.id) foundOrder.value = res.data
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при обновлении статуса')
  }
}

// ===== Payment method modal (shown when confirming an offline order) =====
const showPayModal = ref(false)
const payOrder = ref(null)
const payContext = ref('list')
const paySubmitting = ref(false)

function askPayment(order, context) {
  payOrder.value = order
  payContext.value = context || 'list'
  showPayModal.value = true
}

function closePayModal() {
  if (paySubmitting.value) return
  showPayModal.value = false
  payOrder.value = null
}

async function confirmPayment(method) {
  if (!payOrder.value || paySubmitting.value) return
  paySubmitting.value = true
  try {
    const res = await api.put(`/pickup/orders/${payOrder.value.id}/status`, { status: 'delivered', payment_method: method })
    const idx = orders.value.findIndex(o => o.id === payOrder.value.id)
    if (idx !== -1) orders.value[idx] = res.data
    if (foundOrder.value?.id === payOrder.value.id) foundOrder.value = res.data
    if (payContext.value === 'nurse') nurseConfirmed.value = true
    showPayModal.value = false
    payOrder.value = null
    loadOrders()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при подтверждении')
  } finally {
    paySubmitting.value = false
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

// ===== Inline edit for orders list =====
const listEdit = ref({}) // { [orderId]: { editing, items, addProductId, addQty, addUnit, saving } }

function startListEdit(order) {
  listEdit.value[order.id] = {
    editing: true,
    items: (order.items || []).filter(item => item.quantity > 0).map(item => ({
      product_id: item.product_id,
      name: item.product?.name || '',
      quantity: item.quantity,
      unit_type: item.unit_type,
    })),
    addProductId: '',
    addQty: 1,
    addUnit: 'pack',
    saving: false,
  }
}

function cancelListEdit(orderId) {
  delete listEdit.value[orderId]
}

function listEditInc(orderId, idx) {
  listEdit.value[orderId].items[idx].quantity++
}

function listEditDec(orderId, idx) {
  const item = listEdit.value[orderId].items[idx]
  if (item.quantity > 1) {
    item.quantity--
  } else {
    listEdit.value[orderId].items.splice(idx, 1)
  }
}

function listEditAddItem(orderId) {
  const state = listEdit.value[orderId]
  if (!state.addProductId || state.addQty < 1) return
  const product = allProducts.value.find(p => p.id === state.addProductId)
  if (!product) return
  const existing = state.items.find(i => i.product_id === product.id && i.unit_type === state.addUnit)
  if (existing) {
    existing.quantity += state.addQty
  } else {
    state.items.push({
      product_id: product.id,
      name: product.name,
      quantity: state.addQty,
      unit_type: state.addUnit,
    })
  }
  state.addProductId = ''
  state.addQty = 1
  state.addUnit = 'pack'
}

async function fullReturn(order) {
  const r = prompt('Полный возврат — укажите причину:')
  if (r === null) return
  const reason = r.trim()
  if (!reason) { alert('Причина возврата обязательна'); return }
  try {
    const res = await api.post(`/pickup/orders/${order.id}/return`, { return_reason: reason })
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx !== -1) orders.value[idx] = res.data
    loadStock()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при возврате')
  }
}

async function saveListEdit(order) {
  const state = listEdit.value[order.id]
  if (!state || state.items.length === 0) return
  // Editing a delivered order is a return — ask the reason.
  let returnReason = ''
  if (order.status === 'delivered') {
    const r = prompt('Причина возврата:')
    if (r === null) return
    returnReason = r.trim()
    if (!returnReason) { alert('Причина возврата обязательна'); return }
  }
  state.saving = true
  try {
    const res = await api.put(`/pickup/orders/${order.id}/items`, {
      items: state.items.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_type: i.unit_type,
      })),
      return_reason: returnReason,
    })
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx !== -1) orders.value[idx] = res.data
    delete listEdit.value[order.id]
    loadStock()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при сохранении')
    state.saving = false
  }
}

// ===== Direct Offline Sale =====
const offlineOpen = ref(true)
const allProducts = ref([])
const offlineProductId = ref('')
const offlineQty = ref(1)
const offlineItems = ref([])
const offlineNote = ref('')
const offlineSubmitting = ref(false)
const offlineSuccess = ref('')
const saleType = ref('regular') // 'regular' | 'vip' | 'marketolog'
const offlineMarketolog = ref(null)
const marketologs = ref([])
const offlinePaymentMethod = ref('cash')
const offlineReferral = ref('')
const offlineUnit = ref('pack')
const allDoctors = ref([])

async function loadMarketologs() {
  try { marketologs.value = (await api.get('/pickup/marketologs')).data || [] } catch (e) { console.error(e) }
}

// Capsules available = pieces in stock / pieces per capsule.
function qppOf(productId) {
  const p = allProducts.value.find(x => x.id === productId)
  return p && p.quantity_per_pack > 0 ? p.quantity_per_pack : 1
}
function capsulesOf(productId) { return Math.floor(stockOf(productId) / qppOf(productId)) }

const paymentMethods = computed(() => [
  { value: 'cash', label: txt.value.pay_cash },
  { value: 'terminal', label: txt.value.pay_terminal },
  { value: 'card', label: txt.value.pay_card },
])

const offlineTotal = computed(() => offlineItems.value.reduce((s, i) => s + i.price, 0))

const offlineCanSubmit = computed(() => {
  if (offlineItems.value.length === 0) return false
  if (saleType.value === 'marketolog' && !offlineMarketolog.value) return false
  return true
})

async function loadProducts() {
  try {
    const res = await api.get('/products')
    allProducts.value = res.data || []
    for (const p of allProducts.value) p.price_per_pack = p.price_per_pill * p.quantity_per_pack
  } catch (e) { console.error(e) }
}

async function loadDoctors() {
  try {
    const res = await api.get('/doctors')
    allDoctors.value = res.data || []
  } catch (e) { console.error(e) }
}

function addOfflineItem() {
  if (!offlineProductId.value || offlineQty.value < 1) return
  const product = allProducts.value.find(p => p.id === offlineProductId.value)
  if (!product) return
  const unit = offlineUnit.value === 'piece' ? 'piece' : 'pack'
  const qpp = qppOf(product.id)
  // pieces required by this addition + already in cart for this product
  const piecesNeeded = unit === 'piece' ? offlineQty.value : offlineQty.value * qpp
  const alreadyPieces = offlineItems.value
    .filter(i => i.product_id === product.id)
    .reduce((s, i) => s + (i.unit_type === 'piece' ? i.quantity : i.quantity * qpp), 0)
  if (alreadyPieces + piecesNeeded > stockOf(product.id)) {
    alert(`${product.name}: ${txt.value.my_stock} ${stockOf(product.id)} ${txt.value.piece}`)
    return
  }
  const price = (unit === 'piece' ? product.price_per_pill : product.price_per_pack) * offlineQty.value
  offlineItems.value.push({ product_id: product.id, name: product.name, quantity: offlineQty.value, unit_type: unit, price })
  offlineProductId.value = ''
  offlineQty.value = 1
}

function resetOfflineSale() {
  offlineItems.value = []
  offlineNote.value = ''
  saleType.value = 'regular'
  offlineMarketolog.value = null
  offlinePaymentMethod.value = 'cash'
  offlineReferral.value = ''
  offlineUnit.value = 'pack'
}

async function submitOfflineSale() {
  if (!offlineCanSubmit.value) return
  offlineSubmitting.value = true
  offlineSuccess.value = ''
  try {
    const isVip = saleType.value === 'vip'
    const isMkt = saleType.value === 'marketolog'
    const res = await api.post('/pickup/offline-sale', {
      items: offlineItems.value.map(i => ({ product_id: i.product_id, quantity: i.quantity, unit_type: i.unit_type })),
      offline_note: offlineNote.value,
      is_vip: isVip,
      marketolog_id: isMkt ? offlineMarketolog.value : null,
      payment_method: (isVip || isMkt) ? '' : offlinePaymentMethod.value,
      referred_by: offlineReferral.value.trim(),
    })
    offlineSuccess.value = true
    resetOfflineSale()
    loadOrders()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при записи')
  } finally { offlineSubmitting.value = false }
}

// ===== My analytics =====
const analyticsPeriod = ref('daily')
const analyticsDate = ref('')
const analyticsData = ref(null)
const analyticsLoading = ref(false)

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const params = { period: analyticsPeriod.value }
    if (analyticsPeriod.value === 'custom') params.date = analyticsDate.value
    const res = await api.get('/pickup/analytics', { params })
    analyticsData.value = res.data
  } catch (e) { console.error(e) } finally { analyticsLoading.value = false }
}

function selectAnalyticsPeriod(p) {
  analyticsPeriod.value = p
  if (p !== 'custom') loadAnalytics()
}

const analyticsCat = ref('all')
const emptyCat = { orders: 0, capsules: 0, pieces: 0, revenue: 0 }
function catData(key) { return (analyticsData.value?.breakdown && analyticsData.value.breakdown[key]) || emptyCat }
const allCats = computed(() => [
  { key: 'vip', label: txt.value.cat_vip, color: 'text-amber-600' },
  { key: 'doctor', label: txt.value.cat_doctor, color: 'text-purple-600' },
  { key: 'marketolog', label: txt.value.cat_marketolog, color: 'text-indigo-600' },
  { key: 'regular', label: txt.value.cat_regular, color: 'text-emerald-600' },
])
const shownCats = computed(() => analyticsCat.value === 'all' ? allCats.value : allCats.value.filter(c => c.key === analyticsCat.value))

// ===== Warehouse (personal stock) =====
const stock = ref([])
const stockProductId = ref('')
const stockQty = ref(1)
const stockUnit = ref('pack')
const addingStock = ref(false)

const stockMap = computed(() => {
  const m = {}
  for (const s of stock.value) m[s.product_id] = s.quantity
  return m
})
function stockOf(productId) { return displayStock(productId, stockMap.value[productId] || 0) }

async function loadStock() {
  try { stock.value = (await api.get('/pickup/stock')).data || [] } catch (e) { console.error(e) }
}

async function addStock() {
  if (!stockProductId.value || stockQty.value < 1) return
  addingStock.value = true
  try {
    await api.post('/pickup/stock', { product_id: stockProductId.value, quantity: stockQty.value, unit_type: stockUnit.value })
    stockProductId.value = ''
    stockQty.value = 1
    stockUnit.value = 'pack'
    loadStock()
  } catch (e) { alert(e.response?.data?.error || 'Ошибка') } finally { addingStock.value = false }
}

watch(tab, (t) => {
  if (t === 'analytics' && !analyticsData.value) loadAnalytics()
  if (t === 'stock') loadStock()
  if (t === 'offline') loadStock()
})

// Real-time: refresh the visible data the moment any order changes anywhere.
watch(() => realtime.ordersVersion, () => {
  loadOrders()
  if (tab.value === 'analytics') loadAnalytics()
  if (tab.value === 'stock' || tab.value === 'offline') loadStock()
})

let stockPoll = null
onMounted(() => {
  loadStock()
  loadOrders()
  loadProducts()
  loadDoctors()
  loadMarketologs()
  // Keep stock fresh (near real-time) while selling or managing the warehouse.
  stockPoll = setInterval(() => {
    if (tab.value === 'offline' || tab.value === 'stock') loadStock()
  }, 7000)
})
onUnmounted(() => { if (stockPoll) clearInterval(stockPoll) })
</script>
