<template>
  <div class="min-h-screen flex pp-root" :class="{ 'night-mode': night }">

    <!-- ===== SIDEBAR ===== -->
    <aside class="worker-sidebar pp-sidebar flex-shrink-0 flex flex-col"
      style="width:200px;min-height:100vh;position:sticky;top:0;height:100vh;overflow-y:auto;">
      <div class="px-4 py-5 flex items-center gap-3 pp-sidebar-border-b">
        <div class="w-9 h-9 rounded-xl bg-blue-600 flex items-center justify-center flex-shrink-0">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/>
          </svg>
        </div>
        <div class="min-w-0">
          <p class="pp-sidebar-text text-sm font-bold leading-tight">{{ txt.title }}</p>
          <p class="pp-sidebar-text-muted text-xs leading-tight truncate">{{ authStore.worker?.name }}</p>
        </div>
      </div>

      <nav class="flex-1 px-3 py-3 space-y-0.5">
        <button
          v-for="s in [
            {k:'online',    l:txt.nav_online,    icon:'M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064'},
            {k:'offline',   l:txt.nav_offline,   icon:'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'},
            {k:'stock',     l:txt.nav_stock,     icon:'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4'},
            {k:'analytics', l:txt.nav_analytics, icon:'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'},
            {k:'history',   l:txt.nav_history,   icon:'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'}
          ]"
          :key="s.k"
          @click="tab = s.k"
          :class="tab === s.k ? 'bg-blue-600 text-white' : 'pp-sidebar-nav-inactive'"
          class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-all text-left">
          <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" :d="s.icon"/>
          </svg>
          {{ s.l }}
        </button>
      </nav>

      <div class="px-4 pb-5 pt-3 space-y-3 pp-sidebar-border-t">
        <div class="flex items-center gap-1.5">
          <button @click="lang='ru'" :class="lang==='ru' ? 'pp-sidebar-lang-active font-bold' : 'pp-sidebar-lang-inactive'" class="text-sm transition">RU</button>
          <span class="pp-sidebar-lang-sep text-sm">|</span>
          <button @click="lang='uz'" :class="lang==='uz' ? 'pp-sidebar-lang-active font-bold' : 'pp-sidebar-lang-inactive'" class="text-sm transition">UZ</button>
        </div>
        <button @click="toggleNight" class="flex items-center gap-2 text-sm transition w-full" :class="night ? 'text-amber-400 hover:text-amber-300' : 'pp-sidebar-theme-btn'">
          <svg v-if="night" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
          <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>
          </svg>
          {{ night ? txt.day_mode : txt.night_mode }}
        </button>
        <button @click="logout" class="flex items-center gap-2 text-red-500 hover:text-red-400 text-sm font-medium transition w-full">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
          </svg>
          {{ txt.logout }}
        </button>
      </div>
    </aside>

    <!-- ===== MAIN ===== -->
    <div class="flex-1 flex flex-col min-h-screen pp-main">

      <!-- Top header -->
      <header class="flex-shrink-0 flex items-center justify-between px-6 py-3 pp-header">
        <h1 class="text-base font-semibold pp-text">{{ tabTitle }}</h1>
        <div class="flex items-center gap-4">
          <button class="relative pp-text-3 hover:pp-text-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>
            </svg>
          </button>
          <div class="flex items-center gap-2">
            <div class="w-8 h-8 rounded-full bg-blue-600 flex items-center justify-center text-white text-xs font-bold">
              {{ (authStore.worker?.name || 'A').charAt(0).toUpperCase() }}
            </div>
            <div class="text-sm">
              <p class="font-medium pp-text leading-tight">{{ txt.title }}</p>
              <p class="pp-text-3 text-xs leading-tight">{{ authStore.worker?.name }}</p>
            </div>
            <svg class="w-4 h-4 pp-text-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </div>
        </div>
      </header>

      <!-- ===== ONLINE TAB ===== -->
      <div v-show="tab === 'online'" class="flex-1 p-6 space-y-5">
        <!-- Search card -->
        <div class="pp-card rounded-xl shadow-sm p-6">
          <h2 class="text-base font-semibold pp-text mb-4">{{ txt.search_online }}</h2>
          <div class="flex gap-3">
            <input
              v-model="searchCode"
              type="text"
              maxlength="6"
              :placeholder="txt.enter_6_code"
              class="flex-1 border border-gray-300 rounded-lg px-4 py-3 text-xl font-bold tracking-widest text-center focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
              @keyup.enter="searchByCode"
            />
            <button
              @click="searchByCode"
              :disabled="searchCode.length < 6 || searching"
              class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition font-medium disabled:opacity-40 whitespace-nowrap">
              {{ searching ? txt.searching : txt.find_order }}
            </button>
          </div>
          <p v-if="searchError" class="mt-3 text-red-500 text-sm">{{ searchError }}</p>

          <!-- Found order -->
          <div v-if="foundOrder" class="mt-5 border-2 border-blue-200 rounded-xl p-5 bg-blue-50">
            <div class="flex items-start justify-between mb-4">
              <div>
                <div class="flex items-center gap-2 mb-1 flex-wrap">
                  <span class="text-xl font-bold text-blue-700 tracking-widest">{{ foundOrder.order_code }}</span>
                  <span :class="statusClass(foundOrder.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(foundOrder.status) }}</span>
                </div>
                <p class="font-semibold pp-text">{{ foundOrder.user?.first_name }} {{ foundOrder.user?.last_name }}</p>
                <p class="text-sm text-gray-500">+{{ foundOrder.phone }}</p>
                <p class="text-xs text-gray-400 mt-0.5">{{ new Date(foundOrder.created_at).toLocaleString('ru-RU') }}</p>
              </div>
              <div class="text-right">
                <p class="text-xs text-gray-400">{{ txt.total }}</p>
                <p class="text-xl font-bold text-blue-700">{{ formatPrice(orderTotal(foundOrder)) }} <span class="text-sm font-normal">{{ txt.sum }}</span></p>
              </div>
            </div>
            <div class="border-t border-blue-200 pt-3 mb-4 space-y-1">
              <div v-for="item in foundOrder.items" :key="item.id" class="flex justify-between text-sm text-gray-700">
                <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span></span>
                <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
              </div>
            </div>
            <div class="flex gap-2 flex-wrap">
              <button v-if="foundOrder.status === 'pending'" @click="updateStatus(foundOrder, 'confirmed')" class="flex-1 bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition font-medium text-sm">✓ {{ txt.confirm }}</button>
              <button v-if="foundOrder.status === 'confirmed' || foundOrder.status === 'shipped'" @click="updateStatus(foundOrder, 'in_transit')" class="flex-1 bg-orange-500 text-white py-2 rounded-lg hover:bg-orange-600 transition font-medium text-sm">🚚 {{ txt.in_transit }}</button>
              <button v-if="foundOrder.status === 'in_transit'" @click="updateStatus(foundOrder, 'delivered')" class="flex-1 bg-green-600 text-white py-2 rounded-lg hover:bg-green-700 transition font-medium text-sm">✓ {{ txt.deliver }}</button>
              <button v-if="foundOrder.status !== 'cancelled' && foundOrder.status !== 'delivered'" @click="updateStatus(foundOrder, 'cancelled')" class="flex-1 bg-red-50 text-red-600 border border-red-200 py-2 rounded-lg hover:bg-red-100 transition font-medium text-sm">{{ txt.cancel }}</button>
            </div>
          </div>
        </div>

        <!-- Pending orders -->
        <div class="pp-card rounded-xl shadow-sm overflow-hidden">
          <div class="flex items-center justify-between px-6 py-4 border-b">
            <h2 class="font-semibold pp-text">{{ txt.pending_online }}</h2>
            <button @click="loadOrders" class="flex items-center gap-1.5 text-sm text-blue-600 hover:text-blue-800 transition">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
              {{ txt.refresh }}
            </button>
          </div>
          <div v-if="onlineOrders.length === 0" class="flex flex-col items-center justify-center py-16 text-gray-400">
            <svg class="w-14 h-14 mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
            </svg>
            <p class="font-medium text-gray-500">{{ txt.no_orders }}</p>
            <p class="text-sm text-gray-400 mt-1">{{ txt.online_orders_hint }}</p>
          </div>
          <div v-else class="pp-divide">
            <div v-for="order in onlineOrders" :key="order.id" class="p-5 pp-row-hover transition">
              <div class="flex justify-between items-start gap-3">
                <div class="flex-1">
                  <div class="flex items-center gap-2 mb-1 flex-wrap">
                    <span class="text-lg font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                    <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                    <span v-if="paymentLabel(order)" :class="paymentBadgeClass(order)" class="text-xs font-medium px-2 py-0.5 rounded">{{ paymentLabel(order) }}</span>
                  </div>
                  <p class="font-semibold pp-text">{{ order.user?.first_name }} {{ order.user?.last_name }}</p>
                  <p class="text-sm text-gray-500">+{{ order.phone }}</p>
                  <p class="text-xs text-gray-400">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                  <p v-if="order.delivery_address" class="text-xs text-orange-700 mt-0.5">📍 {{ order.delivery_address }}</p>
                </div>
                <div class="text-right flex-shrink-0">
                  <p class="font-bold pp-text">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
                  <p class="text-xs text-gray-400">{{ order.items?.length }} {{ txt.positions }}</p>
                </div>
              </div>
              <div class="mt-2 pt-2 pp-border-t space-y-1">
                <div v-for="item in boughtItems(order)" :key="item.id" class="flex justify-between text-sm text-gray-600">
                  <span>{{ item.product?.name }} × {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span>
                  <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                </div>
              </div>
              <div class="mt-3 flex gap-2 flex-wrap">
                <button v-if="order.status === 'pending'" @click="updateStatus(order, 'confirmed')" class="bg-blue-600 text-white px-4 py-1.5 rounded-lg hover:bg-blue-700 transition text-sm font-medium">✓ {{ txt.confirm }}</button>
                <button v-if="order.status === 'confirmed' || order.status === 'shipped'" @click="updateStatus(order, 'in_transit')" class="bg-orange-500 text-white px-4 py-1.5 rounded-lg hover:bg-orange-600 transition text-sm font-medium">🚚 {{ txt.in_transit }}</button>
                <button v-if="order.status === 'in_transit'" @click="updateStatus(order, 'delivered')" class="bg-green-600 text-white px-4 py-1.5 rounded-lg hover:bg-green-700 transition text-sm font-medium">✓ {{ txt.deliver }}</button>
                <button v-if="order.status !== 'cancelled' && order.status !== 'delivered'" @click="updateStatus(order, 'cancelled')" class="bg-red-50 text-red-600 border border-red-200 px-4 py-1.5 rounded-lg hover:bg-red-100 transition text-sm font-medium">{{ txt.cancel }}</button>
                <button @click="openChat(order)" class="bg-indigo-50 text-indigo-600 border border-indigo-200 px-3 py-1.5 rounded-lg hover:bg-indigo-100 transition text-sm font-medium flex items-center gap-1">
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"/></svg>
                  {{ txt.write }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ===== OFFLINE TAB ===== -->
      <div v-show="tab === 'offline'" class="flex-1 p-6 space-y-5">
        <!-- Nurse order lookup -->
        <div class="pp-card rounded-xl shadow-sm p-6">
          <h2 class="text-base font-semibold pp-text mb-4">{{ txt.nurse_section }}</h2>
          <p class="text-sm text-gray-500 mb-4">{{ txt.nurse_desc }}</p>
          <div class="flex gap-3">
            <input
              v-model="nurseCode"
              type="text"
              maxlength="5"
              :placeholder="txt.nurse_placeholder"
              class="flex-1 border-2 border-gray-200 rounded-xl px-5 py-3 text-2xl font-bold tracking-[0.4em] text-center focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-teal-400 transition"
              @keyup.enter="searchNurseOrder"
            />
            <button @click="searchNurseOrder" :disabled="nurseCode.length < 5 || nurseSearching" class="bg-teal-600 text-white px-6 py-3 rounded-xl hover:bg-teal-700 transition font-semibold disabled:opacity-40">
              {{ nurseSearching ? txt.searching : txt.find }}
            </button>
          </div>
          <p v-if="nurseSearchError" class="mt-3 text-red-500 text-sm">{{ nurseSearchError }}</p>
          <div v-if="nurseOrder" class="mt-5 border-2 border-teal-200 rounded-xl p-5 bg-teal-50">
            <div v-if="nurseConfirmed" class="text-center py-4">
              <svg class="w-14 h-14 text-teal-500 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <p class="text-xl font-bold text-teal-700">{{ txt.payment_success }}</p>
              <button @click="resetNurse" class="mt-4 text-sm text-teal-600 underline">{{ txt.new_search }}</button>
            </div>
            <template v-else>
              <div class="flex items-start justify-between mb-4">
                <div>
                  <p class="text-2xl font-bold tracking-widest text-teal-700 mb-1">{{ nurseOrder.order_code }}</p>
                  <p class="font-semibold pp-text">{{ nurseOrder.patient_first_name }} {{ nurseOrder.patient_last_name }}</p>
                  <p class="text-xs text-gray-400 mt-0.5">{{ new Date(nurseOrder.created_at).toLocaleString('ru-RU') }}</p>
                </div>
                <p class="text-xl font-bold text-teal-700">{{ formatPrice(orderTotal(nurseOrder)) }} {{ txt.sum }}</p>
              </div>
              <div class="border-t border-teal-200 pt-3 mb-4 space-y-1">
                <div v-for="item in boughtItems(nurseOrder)" :key="item.id" class="flex justify-between text-sm text-gray-700">
                  <span>{{ item.product?.name }} × {{ item.quantity }} {{ txt.pack }}</span>
                  <span>{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                </div>
              </div>
              <div class="bg-white border border-teal-200 rounded-xl p-4 mb-4">
                <p class="text-sm font-medium text-gray-700 mb-2">{{ txt.verify_name }}</p>
                <input v-model="verifyName" type="text" class="w-full border-2 border-gray-200 rounded-xl px-4 py-3 text-base focus:outline-none focus:ring-2 focus:ring-teal-500" :placeholder="txt.enter_patient_name"/>
                <p v-if="verifyError" class="mt-2 text-red-500 text-sm">{{ verifyError }}</p>
              </div>
              <button @click="confirmNurseOrder" :disabled="!verifyName.trim() || nurseConfirming" class="w-full bg-teal-600 text-white py-3 rounded-xl hover:bg-teal-700 transition font-bold disabled:opacity-40">
                {{ nurseConfirming ? txt.confirming : txt.confirm_issue }}
              </button>
            </template>
          </div>
        </div>

        <!-- Direct offline sale: 3-step wizard -->
        <div class="pp-card rounded-xl shadow-sm p-6">
          <h2 class="text-base font-semibold pp-text mb-5">{{ txt.offline_sale }}</h2>

          <div class="grid grid-cols-3 gap-4 mb-5">
            <!-- Step 1: Product -->
            <div>
              <div class="flex items-center gap-2 mb-3">
                <span class="w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-bold flex items-center justify-center flex-shrink-0">1</span>
                <span class="text-sm font-semibold text-gray-700">{{ txt.select_product }}</span>
              </div>
              <div class="relative">
                <select v-model="offlineProductId" class="w-full border border-gray-300 rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 appearance-none pr-8">
                  <option value="">{{ txt.select_product }}</option>
                  <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
                </select>
                <svg class="w-4 h-4 text-gray-400 absolute right-2 top-3 pointer-events-none" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/></svg>
              </div>
              <!-- Search hint below select -->
              <div v-if="offlineProductId" class="mt-2 p-2 pp-inset rounded-lg">
                <p class="text-xs text-gray-500">{{ txt.in_stock }}: <span class="font-semibold text-gray-700">{{ capsulesOf(offlineProductId) }} {{ txt.pack }} / {{ stockOf(offlineProductId) }} {{ txt.piece }}</span></p>
              </div>
            </div>

            <!-- Step 2: Quantity -->
            <div>
              <div class="flex items-center gap-2 mb-3">
                <span class="w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-bold flex items-center justify-center flex-shrink-0">2</span>
                <span class="text-sm font-semibold text-gray-700">{{ txt.qty }}</span>
              </div>
              <div class="flex items-center gap-2 mb-3">
                <button @click="offlineQty = Math.max(1, offlineQty - 1)" class="w-9 h-9 rounded-lg bg-gray-100 hover:bg-gray-200 text-gray-700 font-bold text-lg flex items-center justify-center transition">−</button>
                <input v-model.number="offlineQty" type="number" min="1" class="flex-1 border border-gray-300 rounded-lg px-3 py-2 text-center text-base font-bold focus:outline-none focus:ring-2 focus:ring-blue-500"/>
                <button @click="offlineQty++" class="w-9 h-9 rounded-lg bg-blue-600 hover:bg-blue-700 text-white font-bold text-lg flex items-center justify-center transition">+</button>
              </div>
              <p class="text-xs text-gray-500 mb-2">{{ txt.quick_select }}:</p>
              <div class="grid grid-cols-3 gap-1">
                <button v-for="n in [1,2,5,10,20,50]" :key="n" @click="offlineQty=n"
                  :class="offlineQty===n ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
                  class="py-1 rounded text-xs font-medium transition">{{ n }}</button>
              </div>
            </div>

            <!-- Step 3: Unit -->
            <div>
              <div class="flex items-center gap-2 mb-3">
                <span class="w-6 h-6 rounded-full bg-blue-600 text-white text-xs font-bold flex items-center justify-center flex-shrink-0">3</span>
                <span class="text-sm font-semibold text-gray-700">{{ txt.unit }}</span>
              </div>
              <div class="grid grid-cols-2 gap-2">
                <button v-for="u in [{v:'pack',l:txt.pack},{v:'piece',l:txt.piece}]" :key="u.v"
                  @click="offlineUnit=u.v"
                  :class="offlineUnit===u.v ? 'bg-blue-600 text-white border-blue-600' : 'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                  class="border py-2 rounded-lg text-sm font-medium transition">{{ u.l }}</button>
              </div>
            </div>
          </div>

          <!-- Price + total + add button -->
          <div class="flex items-center gap-4 p-4 pp-inset rounded-xl mb-4">
            <div class="flex-1">
              <label class="text-xs pp-text-3 mb-1 block">{{ txt.price_per_unit }}</label>
              <p class="text-lg font-bold text-gray-800">
                {{ offlineProductId ? formatPrice(offlineUnit === 'piece' ? (allProducts.find(p=>p.id===offlineProductId)?.price_per_pill || 0) : (allProducts.find(p=>p.id===offlineProductId)?.price_per_pack || 0)) : '0' }} {{ txt.sum }}
              </p>
            </div>
            <div class="flex-1">
              <label class="text-xs pp-text-3 mb-1 block">{{ txt.total_sum }}:</label>
              <p class="text-lg font-bold text-emerald-600">
                {{ offlineProductId ? formatPrice((offlineUnit === 'piece' ? (allProducts.find(p=>p.id===offlineProductId)?.price_per_pill || 0) : (allProducts.find(p=>p.id===offlineProductId)?.price_per_pack || 0)) * offlineQty) : '0' }} {{ txt.sum }}
              </p>
            </div>
            <button @click="addOfflineItem" :disabled="!offlineProductId || offlineQty < 1" class="flex items-center gap-2 bg-blue-600 text-white px-5 py-3 rounded-xl hover:bg-blue-700 transition font-semibold disabled:opacity-40">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
              {{ txt.add_product }}
            </button>
          </div>

          <!-- Cart items -->
          <div v-if="offlineItems.length" class="pp-border-box rounded-xl overflow-hidden mb-4">
            <div v-for="(item, idx) in offlineItems" :key="idx" class="flex items-center justify-between px-4 py-3 pp-cart-row border-b last:border-0">
              <div>
                <span class="font-medium pp-text text-sm">{{ item.name }}</span>
                <span class="text-gray-500 text-sm ml-2">× {{ item.quantity }} {{ item.unit_type === 'piece' ? txt.piece : txt.pack }}</span>
              </div>
              <div class="flex items-center gap-3">
                <span class="font-semibold text-gray-700 text-sm">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                <button @click="offlineItems.splice(idx, 1)" class="text-red-400 hover:text-red-600 transition">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                </button>
              </div>
            </div>
            <div class="flex items-center justify-between px-4 py-2 bg-white border-t">
              <span class="text-sm font-semibold text-gray-700">{{ txt.total }}:</span>
              <span class="font-bold text-emerald-600">{{ formatPrice(offlineTotal) }} {{ txt.sum }}</span>
            </div>
          </div>

          <!-- Sale type + payment + referral + submit -->
          <div v-if="offlineItems.length" class="space-y-4">
            <div class="pp-border-box rounded-xl px-4 py-4 space-y-4">
              <div>
                <label class="text-xs font-medium text-gray-500 mb-2 block">{{ txt.sale_type }}</label>
                <div class="flex gap-2 flex-wrap">
                  <button @click="saleType='regular'" :class="saleType==='regular'?'bg-emerald-600 text-white border-emerald-600':'bg-white text-gray-600 border-gray-300 hover:border-gray-400'" class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.sale_regular }}</button>
                  <button @click="saleType='vip'" :class="saleType==='vip'?'bg-amber-500 text-white border-amber-500':'bg-white text-gray-600 border-gray-300 hover:border-gray-400'" class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.vip_free }}</button>
                  <button @click="saleType='marketolog'" :class="saleType==='marketolog'?'bg-purple-600 text-white border-purple-600':'bg-white text-gray-600 border-gray-300 hover:border-gray-400'" class="px-3 py-2 rounded-lg text-sm font-medium border transition">{{ txt.sale_marketolog }}</button>
                </div>
              </div>
              <div v-if="saleType==='marketolog'">
                <label class="text-xs font-medium text-gray-500 mb-1.5 block">{{ txt.choose_marketolog }}</label>
                <select v-model="offlineMarketolog" class="w-full pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-purple-500">
                  <option :value="null">{{ txt.choose_marketolog }}</option>
                  <option v-for="m in marketologs" :key="m.id" :value="m.id">{{ m.name }}</option>
                </select>
              </div>
              <div v-if="saleType==='regular'">
                <label class="text-xs font-medium text-gray-500 mb-1.5 block">{{ txt.payment_method }}</label>
                <div class="flex gap-2">
                  <button v-for="pm in paymentMethods" :key="pm.value" @click="offlinePaymentMethod=pm.value"
                    :class="offlinePaymentMethod===pm.value?'bg-emerald-600 text-white border-emerald-600':'bg-white text-gray-600 border-gray-300 hover:border-gray-400'"
                    class="flex-1 py-2 rounded-lg text-sm font-medium border transition">{{ pm.label }}</button>
                </div>
              </div>
            </div>
            <div v-if="saleType !== 'marketolog'" class="flex gap-3">
              <input v-model="offlineNote" :placeholder="txt.buyer_name" class="flex-1 pp-input rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500"/>
              <input v-model="offlineReferral" list="offline-doctors" :placeholder="txt.referral_ph" class="flex-1 pp-input rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-500"/>
              <datalist id="offline-doctors">
                <option value="Самостоятельно"></option>
                <option v-for="d in allDoctors" :key="d.id" :value="d.name + (d.specialty?' ('+d.specialty+')':'')"></option>
              </datalist>
            </div>
            <button @click="submitOfflineSale" :disabled="!offlineCanSubmit || offlineSubmitting" class="w-full bg-emerald-600 text-white py-3 rounded-xl hover:bg-emerald-700 transition font-bold disabled:opacity-40">
              {{ offlineSubmitting ? txt.saving : txt.record_sale }}
            </button>
            <div v-if="offlineSuccess" class="bg-emerald-50 border border-emerald-200 rounded-lg px-4 py-3 text-sm text-emerald-700 flex items-center gap-2">
              <svg class="w-5 h-5 text-emerald-500 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
              {{ txt.sold_ok }}
            </div>
          </div>
        </div>

        <!-- My Stock -->
        <div class="pp-card rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 pp-border-b">
            <h2 class="font-semibold pp-text">{{ txt.my_stock }}</h2>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="pp-table-head border-b">
                <tr>
                  <th class="text-left px-5 py-3 text-xs font-semibold pp-text-3 uppercase">{{ txt.product }}</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold pp-text-3 uppercase">{{ txt.unit }}</th>
                  <th class="text-left px-5 py-3 text-xs font-semibold pp-text-3 uppercase">{{ txt.total }}</th>
                  <th class="text-right px-5 py-3 text-xs font-semibold pp-text-3 uppercase">{{ txt.remainder }}</th>
                </tr>
              </thead>
              <tbody class="pp-divide">
                <tr v-for="s in stock" :key="s.id" class="pp-row-hover transition">
                  <td class="px-5 py-3">
                    <div class="flex items-center gap-3">
                      <div class="w-10 h-10 bg-gray-100 rounded-lg overflow-hidden flex items-center justify-center flex-shrink-0">
                        <img v-if="s.product?.image_path" :src="s.product.image_path" class="w-full h-full object-cover"/>
                        <svg v-else class="w-5 h-5 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                      </div>
                      <span class="font-medium pp-text text-sm">{{ s.product?.name }}</span>
                    </div>
                  </td>
                  <td class="px-5 py-3 text-gray-500 text-sm">{{ txt.pack }}</td>
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

        <!-- Pending offline orders -->
        <div class="pp-card rounded-xl shadow-sm overflow-hidden">
          <div class="flex items-center justify-between px-6 py-4 border-b">
            <h2 class="font-semibold pp-text">{{ txt.pending_offline }}</h2>
            <button @click="loadOrders" class="flex items-center gap-1.5 text-sm text-blue-600 hover:text-blue-800 transition">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
              {{ txt.refresh }}
            </button>
          </div>
          <div v-if="offlineOrders.length === 0" class="flex flex-col items-center justify-center py-12 text-gray-400">
            <svg class="w-12 h-12 mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
            </svg>
            <p class="font-medium text-gray-500">{{ txt.no_orders }}</p>
          </div>
          <div v-else class="pp-divide">
            <div v-for="order in offlineOrders" :key="order.id" class="p-5 pp-row-hover transition">
              <div class="flex justify-between items-start gap-3">
                <div class="flex-1">
                  <div class="flex items-center gap-2 mb-1 flex-wrap">
                    <span class="text-lg font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                    <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                    <span v-if="paymentLabel(order)" :class="paymentBadgeClass(order)" class="text-xs font-medium px-2 py-0.5 rounded">{{ paymentLabel(order) }}</span>
                  </div>
                  <p class="font-semibold pp-text">{{ order.offline_note || '—' }}</p>
                  <p class="text-xs text-gray-400">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                </div>
                <p class="font-bold text-gray-800 text-sm">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
              </div>
              <div class="mt-2 pt-2 pp-border-t space-y-1">
                <div v-for="item in boughtItems(order)" :key="item.id" class="flex justify-between text-sm text-gray-600">
                  <span>{{ item.product?.name }} × {{ item.quantity }} {{ txt.pack }}</span>
                  <span>{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                </div>
              </div>
              <div class="mt-3 flex gap-2 flex-wrap">
                <button v-if="order.status === 'pending'" @click="askPayment(order, 'list')" class="bg-green-600 text-white px-4 py-1.5 rounded-lg hover:bg-green-700 transition text-sm font-medium">✓ {{ txt.confirm }}</button>
                <button v-if="order.status !== 'cancelled' && order.status !== 'delivered'" @click="updateStatus(order, 'cancelled')" class="bg-red-50 text-red-600 border border-red-200 px-4 py-1.5 rounded-lg hover:bg-red-100 transition text-sm font-medium">{{ txt.cancel }}</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ===== STOCK TAB ===== -->
      <div v-show="tab === 'stock'" class="flex-1 p-6 space-y-5">
        <!-- Add stock -->
        <div class="pp-card rounded-xl shadow-sm p-6">
          <h2 class="text-base font-semibold pp-text mb-4">{{ txt.stock_in }}</h2>
          <div class="flex gap-3 flex-wrap items-end">
            <div class="flex-1 min-w-[180px]">
              <label class="text-xs font-medium text-gray-500 mb-1 block">{{ txt.product }}</label>
              <select v-model="stockProductId" class="w-full pp-input rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option value="">{{ txt.select_product }}</option>
                <option v-for="p in allProducts" :key="p.id" :value="p.id">{{ p.name }}</option>
              </select>
            </div>
            <div class="w-28">
              <label class="text-xs font-medium text-gray-500 mb-1 block">{{ txt.stock_qty }}</label>
              <input v-model.number="stockQty" type="number" min="1" class="w-full pp-input rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"/>
            </div>
            <div class="w-28">
              <label class="text-xs font-medium text-gray-500 mb-1 block">{{ txt.unit }}</label>
              <select v-model="stockUnit" class="w-full pp-input rounded-lg px-3 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option value="pack">{{ txt.pack }}</option>
                <option value="piece">{{ txt.piece }}</option>
              </select>
            </div>
            <button @click="addStock" :disabled="!stockProductId || stockQty < 1 || addingStock" class="bg-blue-600 text-white px-5 py-2.5 rounded-lg hover:bg-blue-700 transition text-sm font-medium disabled:opacity-40">+ {{ txt.stock_add }}</button>
          </div>
        </div>

        <!-- Search + filter -->
        <div class="pp-card rounded-xl shadow-sm p-4">
          <div class="flex gap-3 items-center flex-wrap">
            <div class="flex-1 relative min-w-[200px]">
              <svg class="w-4 h-4 text-gray-400 absolute left-3 top-2.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
              <input v-model="stockSearch" :placeholder="txt.stock_search" class="w-full border border-gray-300 rounded-lg pl-10 pr-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"/>
            </div>
          </div>
        </div>

        <!-- Product list -->
        <div class="pp-card rounded-xl shadow-sm overflow-hidden">
          <div class="px-6 py-4 pp-border-b">
            <h2 class="font-semibold pp-text">{{ txt.product_list }}</h2>
          </div>
          <div class="pp-divide">
            <div v-for="s in filteredStock" :key="s.id" class="flex items-center justify-between px-6 py-4 pp-row-hover transition">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 bg-gray-100 rounded-xl overflow-hidden flex items-center justify-center flex-shrink-0">
                  <img v-if="s.product?.image_path" :src="s.product.image_path" class="w-full h-full object-cover"/>
                  <svg v-else class="w-6 h-6 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                </div>
                <div>
                  <p class="font-semibold pp-text">{{ s.product?.name }}</p>
                  <p class="text-xs text-gray-500">{{ txt.pack }} · {{ txt.price }}: {{ formatPrice(s.product?.price_per_pack) }} {{ txt.sum }}</p>
                </div>
              </div>
              <div class="text-right">
                <p class="text-sm font-bold" :class="stockOf(s.product_id) > 0 ? 'text-green-700' : 'text-red-600'">
                  {{ capsulesOf(s.product_id) }} {{ txt.pack }} / {{ stockOf(s.product_id) }} {{ txt.piece }}
                </p>
                <p class="text-xs text-gray-400 mt-0.5">{{ txt.remainder }}</p>
              </div>
            </div>
            <div v-if="filteredStock.length === 0" class="py-12 text-center text-gray-400 text-sm">{{ txt.stock_empty }}</div>
          </div>
        </div>
      </div>

      <!-- ===== ANALYTICS TAB ===== -->
      <div v-show="tab === 'analytics'" class="flex-1 p-6 space-y-5">
        <!-- Date filter bar -->
        <div class="pp-card rounded-xl shadow-sm px-6 py-4 flex items-center gap-3 flex-wrap">
          <div class="flex gap-1">
            <button v-for="p in [{v:'daily',l:txt.a_today},{v:'weekly',l:txt.a_week},{v:'monthly',l:txt.a_month},{v:'yearly',l:txt.a_year}]" :key="p.v"
              @click="selectAnalyticsPeriod(p.v)"
              :class="analyticsPeriod===p.v ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'"
              class="px-4 py-2 rounded-lg text-sm font-medium transition">{{ p.l }}</button>
          </div>
          <div class="flex items-center gap-2 border border-gray-300 rounded-lg px-3 py-2">
            <input v-model="analyticsDate" type="date" @change="selectAnalyticsPeriod('custom')" class="text-sm text-gray-700 focus:outline-none bg-transparent"/>
            <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
          </div>
          <div class="ml-auto">
            <button @click="exportAnalyticsExcel" class="flex items-center gap-2 bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition text-sm font-medium">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>
              {{ txt.export_excel }}
            </button>
          </div>
        </div>

        <div v-if="analyticsLoading" class="bg-white rounded-xl shadow-sm p-12 text-center text-gray-400">{{ txt.loading }}</div>
        <template v-else-if="analyticsData">
          <!-- Stat cards -->
          <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
            <div class="pp-card rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-500 mb-2">{{ txt.a_orders }}</p>
              <p class="text-3xl font-bold text-gray-800">{{ analyticsData.total_orders }}</p>
            </div>
            <div class="pp-card rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-500 mb-2">{{ txt.a_revenue }}</p>
              <p class="text-2xl font-bold text-emerald-600">{{ formatPrice(analyticsData.total_revenue) }} {{ txt.sum }}</p>
            </div>
            <div class="pp-card rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-500 mb-2">{{ txt.a_created }}</p>
              <p class="text-3xl font-bold text-blue-600">{{ analyticsData.created_count }}</p>
            </div>
            <div class="pp-card rounded-xl shadow-sm p-5">
              <p class="text-xs text-gray-500 mb-2">{{ txt.a_confirmed }}</p>
              <p class="text-3xl font-bold text-teal-600">{{ analyticsData.confirmed_count }}</p>
            </div>
          </div>

          <!-- Chart -->
          <div class="pp-card rounded-xl shadow-sm p-6">
            <h3 class="font-semibold text-gray-800 mb-4">{{ txt.dynamics }}</h3>
            <LineChart :points="analyticsData.points || []" color="#3b82f6"/>
          </div>

          <!-- Category breakdown -->
          <div class="pp-card rounded-xl shadow-sm p-6">
            <h3 class="font-semibold text-gray-800 mb-4">{{ txt.by_category }}</h3>
            <div class="overflow-x-auto">
              <table class="w-full text-sm">
                <thead>
                  <tr class="border-b border-gray-100">
                    <th class="text-left py-2 font-semibold text-gray-500">{{ txt.category }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.a_orders }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.pack }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.piece }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.a_revenue }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                  <tr v-for="cat in allCats" :key="cat.key">
                    <td class="py-3 font-medium" :class="cat.color">{{ cat.label }}</td>
                    <td class="py-3 text-right text-gray-700">{{ catData(cat.key).orders }}</td>
                    <td class="py-3 text-right text-gray-700">{{ catData(cat.key).capsules }}</td>
                    <td class="py-3 text-right text-gray-700">{{ catData(cat.key).pieces }}</td>
                    <td class="py-3 text-right font-bold text-gray-800">{{ formatPrice(catData(cat.key).revenue) }} {{ txt.sum }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Per-product breakdown -->
          <div v-if="analyticsData.top_products && analyticsData.top_products.length" class="pp-card rounded-xl shadow-sm p-6">
            <h3 class="font-semibold text-gray-800 mb-4">{{ txt.top_products }}</h3>
            <div class="overflow-x-auto">
              <table class="w-full text-sm">
                <thead>
                  <tr class="border-b border-gray-100">
                    <th class="text-left py-2 font-semibold text-gray-500">#</th>
                    <th class="text-left py-2 font-semibold text-gray-500">{{ txt.product_name }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.a_orders }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.pack }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.piece }}</th>
                    <th class="text-right py-2 font-semibold text-gray-500">{{ txt.a_revenue }}</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                  <tr v-for="(prod, idx) in analyticsData.top_products" :key="prod.product_id" class="pp-row-hover">
                    <td class="py-3 text-gray-400 font-medium w-8">{{ idx + 1 }}</td>
                    <td class="py-3 font-semibold text-gray-800">{{ prod.product_name }}</td>
                    <td class="py-3 text-right text-gray-700">{{ prod.orders }}</td>
                    <td class="py-3 text-right text-gray-700">{{ prod.capsules }}</td>
                    <td class="py-3 text-right text-gray-700">{{ prod.pieces }}</td>
                    <td class="py-3 text-right font-bold text-blue-600">{{ formatPrice(prod.revenue) }} {{ txt.sum }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </template>
      </div>

      <!-- ===== HISTORY TAB ===== -->
      <div v-show="tab === 'history'" class="flex-1 p-6 space-y-5">
        <!-- Filters + export -->
        <div class="pp-card rounded-xl shadow-sm px-6 py-4">
          <div class="flex items-center gap-3 flex-wrap">
            <h2 class="font-semibold text-gray-800 mr-2">{{ txt.history_title }}</h2>
            <div class="flex items-center gap-2">
              <label class="text-xs text-gray-500">{{ txt.status_label }}:</label>
              <select v-model="historyStatus" class="pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option v-for="f in statusFilters" :key="f.value" :value="f.value">{{ f.label }}</option>
              </select>
            </div>
            <div class="flex items-center gap-2">
              <label class="text-xs text-gray-500">{{ txt.order_type }}:</label>
              <select v-model="historyType" class="pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option value="all">{{ txt.type_all }}</option>
                <option value="online">{{ txt.type_online }}</option>
                <option value="offline">{{ txt.type_offline }}</option>
                <option value="vip">{{ txt.own_patient }}</option>
              </select>
            </div>
            <div class="flex items-center gap-2">
              <label class="text-xs text-gray-500">{{ txt.period_label }}:</label>
              <select v-model="historyPeriod" class="pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option value="all">{{ txt.period_all }}</option>
                <option value="daily">{{ txt.a_today }}</option>
                <option value="weekly">{{ txt.a_week }}</option>
                <option value="monthly">{{ txt.a_month }}</option>
                <option value="custom">{{ txt.a_date }}</option>
              </select>
            </div>
            <input v-if="historyPeriod==='custom'" v-model="historyDate" type="date" class="pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"/>
            <button @click="loadOrders" class="flex items-center gap-1.5 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm font-medium">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
              {{ txt.refresh }}
            </button>
            <button @click="exportHistoryExcel" class="flex items-center gap-2 bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition text-sm font-medium ml-auto">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>
              {{ txt.export_excel }}
            </button>
          </div>
        </div>

        <!-- Orders list -->
        <div class="pp-card rounded-xl shadow-sm overflow-hidden">
          <div v-if="historyOrders.length === 0" class="flex flex-col items-center justify-center py-16 text-gray-400">
            <svg class="w-16 h-16 mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"/>
            </svg>
            <p class="font-semibold text-gray-500">{{ txt.no_orders }}</p>
            <p class="text-sm text-gray-400 mt-1">{{ txt.no_orders_filter }}</p>
          </div>
          <div v-else class="pp-divide">
            <div v-for="order in historyOrders" :key="order.id" class="p-5 pp-row-hover transition">
              <div class="flex justify-between items-start gap-3">
                <div class="flex-1">
                  <div class="flex items-center gap-2 mb-1 flex-wrap">
                    <span class="text-lg font-bold text-blue-700 tracking-widest">{{ order.order_code }}</span>
                    <span :class="statusClass(order.status)" class="text-xs font-medium px-2 py-0.5 rounded">{{ statusLabel(order.status) }}</span>
                    <span v-if="order.is_offline" class="text-xs bg-gray-100 text-gray-600 px-2 py-0.5 rounded">{{ txt.offline_badge }}</span>
                    <span v-if="paymentLabel(order)" :class="paymentBadgeClass(order)" class="text-xs font-medium px-2 py-0.5 rounded">{{ paymentLabel(order) }}</span>
                  </div>
                  <p class="font-semibold pp-text">{{ order.is_offline ? (order.offline_note || '—') : (order.user?.first_name + ' ' + order.user?.last_name) }}</p>
                  <p v-if="!order.is_offline" class="text-sm text-gray-500">+{{ order.phone }}</p>
                  <p class="text-xs text-gray-400">{{ new Date(order.created_at).toLocaleString('ru-RU') }}</p>
                  <p v-if="order.referred_by" class="text-xs text-purple-700 mt-0.5">{{ txt.referred_by }}: {{ order.referred_by }}</p>
                </div>
                <div class="text-right flex-shrink-0">
                  <p class="font-bold pp-text">{{ formatPrice(orderTotal(order)) }} {{ txt.sum }}</p>
                  <p class="text-xs text-gray-400">{{ order.items?.length }} {{ txt.positions }}</p>
                </div>
              </div>
              <div class="mt-3 pt-3 pp-border-t space-y-1">
                <div v-for="item in boughtItems(order)" :key="item.id" class="flex justify-between text-sm text-gray-600">
                  <span>{{ item.product?.name }} <span class="text-gray-400">× {{ item.quantity }} {{ txt.pack }}</span></span>
                  <span class="font-medium">{{ formatPrice(item.price) }} {{ txt.sum }}</span>
                </div>
              </div>
              <!-- Actions for non-delivered/cancelled history items -->
              <div class="mt-3 flex gap-2 flex-wrap">
                <button v-if="order.is_offline && !order.marketolog_id && order.status !== 'cancelled' && !listEdit[order.id]?.editing"
                  @click="startListEdit(order)"
                  :class="order.status === 'delivered' ? 'bg-red-50 text-red-600 border-red-200' : 'bg-gray-100 text-gray-700 border-gray-200'"
                  class="border px-4 py-1.5 rounded-lg transition text-sm font-medium">
                  {{ order.status === 'delivered' ? txt.return_edit : txt.edit_items }}
                </button>
                <button v-if="order.is_offline && !order.marketolog_id && order.status === 'delivered' && !listEdit[order.id]?.editing"
                  @click="fullReturn(order)" class="bg-red-600 text-white border border-red-600 px-4 py-1.5 rounded-lg hover:bg-red-700 transition text-sm font-medium">
                  {{ txt.full_return }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>

  <!-- Payment modal -->
  <div v-if="showPayModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4" @click.self="closePayModal">
    <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
    <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6">
      <div class="text-center mb-5">
        <div class="w-12 h-12 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-3">
          <svg class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/></svg>
        </div>
        <h3 class="text-lg font-bold text-gray-800">{{ txt.choose_payment }}</h3>
        <p v-if="payOrder" class="text-emerald-700 font-bold mt-2">{{ formatPrice(orderTotal(payOrder)) }} {{ txt.sum }}</p>
      </div>
      <div class="grid gap-2">
        <button v-for="pm in paymentMethods" :key="pm.value" @click="confirmPayment(pm.value)" :disabled="paySubmitting"
          class="w-full py-3 rounded-xl border-2 border-emerald-200 text-emerald-700 font-semibold hover:bg-emerald-50 transition disabled:opacity-40">
          {{ pm.label }}
        </button>
      </div>
      <button @click="closePayModal" :disabled="paySubmitting" class="w-full mt-3 py-2 text-sm text-gray-400 hover:text-gray-600 transition">{{ txt.cancel }}</button>
    </div>
  </div>

  <!-- Chat -->
  <div v-if="chatOpen" class="fixed inset-0 z-50 flex items-end sm:items-center justify-center sm:justify-end" @click.self="chatOpen = false">
    <div class="w-full sm:w-96 h-[55vh] sm:h-[70vh] bg-white sm:rounded-tl-2xl sm:rounded-bl-2xl shadow-2xl flex flex-col rounded-t-2xl">
      <div class="flex items-center justify-between px-4 py-3 border-b bg-indigo-600 rounded-t-2xl sm:rounded-tl-2xl">
        <div>
          <p class="text-white font-semibold text-sm">{{ chatUserName }}</p>
          <p class="text-indigo-200 text-xs">{{ txt.chat_with_client }}</p>
        </div>
        <button @click="chatOpen = false" class="text-white/70 hover:text-white">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </div>
      <div ref="chatMessagesEl" class="flex-1 overflow-y-auto p-4 space-y-2 pp-chat-bg">
        <div v-if="chatLoading" class="text-center pt-8 text-gray-400 text-sm">{{ txt.loading }}</div>
        <template v-else>
          <div v-if="chatMessages.length === 0" class="text-center text-gray-400 text-sm pt-8">{{ txt.no_messages }}</div>
          <div v-for="msg in chatMessages" :key="msg.id" class="flex" :class="msg.sender_role === 'user' ? 'justify-start' : 'justify-end'">
            <div class="max-w-[75%] px-3 py-2 rounded-xl text-sm" :class="msg.sender_role === 'user' ? 'bg-white text-gray-800 shadow-sm' : 'bg-indigo-600 text-white'">
              {{ msg.message }}
              <p class="text-[10px] mt-0.5 opacity-50 text-right">{{ new Date(msg.created_at).toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' }) }}</p>
            </div>
          </div>
        </template>
      </div>
      <div class="px-3 py-3 border-t flex gap-2">
        <input v-model="chatMsg" @keyup.enter="sendWorkerMessage" :placeholder="txt.type_message" class="flex-1 pp-input rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"/>
        <button @click="sendWorkerMessage" :disabled="!chatMsg.trim() || chatSending" class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition disabled:opacity-40">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"/></svg>
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
import * as XLSX from 'xlsx'

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
    find_order: 'Найти заказ',
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
    refresh: 'Обновить',
    product: 'Препарат',
    select_product: 'Выберите препарат',
    qty: 'Количество',
    unit: 'Единицы',
    pack: 'капс.',
    piece: 'шт',
    add: 'Добавить',
    add_product: '+ Добавить товар',
    total: 'Итого',
    total_sum: 'Общая сумма',
    price_per_unit: 'Цена закупки (за ед.)',
    sum: 'сум',
    buyer_name: 'Имя покупателя',
    saving: 'Запись...',
    record_sale: 'Записать продажу',
    confirm: 'Подтвердить',
    in_transit: 'В пути',
    deliver: 'Выдать',
    cancel: 'Отменить',
    write: 'Написать',
    offline_badge: 'Офлайн',
    positions: 'позиций',
    no_orders: 'Нет заказов',
    no_orders_filter: 'Заказы не найдены по выбранным фильтрам',
    online_orders_hint: 'Здесь появятся ожидаемые онлайн-заказы',
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
    edit_items: 'Редактировать',
    return_edit: 'Возврат',
    full_return: 'Полный возврат',
    save_items: 'Сохранить',
    saving_items: 'Сохранение...',
    cancel_edit: 'Отмена',
    view_on_map: 'На карте',
    vip_free: 'Свой пациент (бесплатно)',
    sale_type: 'Тип продажи',
    sale_regular: 'Обычная',
    sale_marketolog: 'Маркетолог (долг)',
    choose_marketolog: 'Выберите маркетолога',
    no_marketologs: 'Нет маркетологов',
    payment_method: 'Способ оплаты',
    pay_cash: 'Наличные',
    pay_terminal: 'Терминал',
    pay_card: 'Карта',
    pay_online: 'Онлайн',
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
    a_year: 'Год',
    a_date: 'Дата',
    a_orders: 'Заказов',
    a_revenue: 'Выручка',
    a_created: 'Создано',
    a_confirmed: 'Подтверждено',
    by_category: 'По категориям',
    top_products: 'По препаратам',
    product_name: 'Препарат',
    category: 'Категория',
    cat_all: 'Все категории',
    cat_vip: 'Свой пациент',
    cat_doctor: 'От доктора',
    cat_marketolog: 'Маркетолог',
    cat_regular: 'Обычные',
    sold_ok: 'Продажа записана ✓',
    nav_online: 'Онлайн',
    nav_offline: 'Офлайн',
    nav_analytics: 'Аналитика',
    nav_history: 'История',
    nav_stock: 'Склад',
    stock_in: 'Приход товара',
    my_stock: 'Мой склад',
    stock_qty: 'Кол-во',
    stock_add: 'Пополнить',
    stock_empty: 'Склад пуст',
    stock_search: 'Поиск по складу...',
    product_list: 'Список препаратов',
    price: 'Цена',
    remainder: 'Остаток',
    pending_online: 'Ожидаемые онлайн-заказы',
    pending_offline: 'Ожидаемые офлайн-заказы',
    history_title: 'История заказов',
    status_label: 'Статус',
    order_type: 'Тип заказа',
    period_label: 'Период',
    type_all: 'Все',
    type_online: 'Онлайн',
    type_offline: 'Офлайн',
    own_patient: 'Свой пациент',
    period_all: 'Всё время',
    night_mode: 'Ночной',
    day_mode: 'Дневной',
    in_stock: 'На складе',
    quick_select: 'Быстрый выбор',
    dynamics: 'Динамика за день',
    export_excel: 'Экспорт Excel',
  },
  uz: {
    title: 'Berish punkti',
    logout: 'Chiqish',
    nurse_section: "Hamshira kodi bo'yicha berish (5 raqam)",
    nurse_desc: "Bemor hamshiradan olgan 5 raqamli kodni kiriting",
    nurse_placeholder: 'XXXXX',
    find: 'Topish',
    find_order: 'Buyurtmani topish',
    searching: 'Qidirilmoqda...',
    verify_name: "Bemorning shaxsini tasdiqlang",
    enter_patient_name: "Bemorning ismini kiriting",
    confirm_issue: "✓ Berish va to'lovni tasdiqlash",
    confirming: 'Tasdiqlanmoqda...',
    payment_success: "To'lov muvaffaqiyatli!",
    new_search: 'Yangi qidiruv',
    offline_sale: "To'g'ridan-to'g'ri oflayn sotuv",
    search_online: "Onlayn buyurtmani kod bo'yicha qidirish",
    enter_6_code: '6 raqamli kodni kiriting',
    refresh: 'Yangilash',
    product: 'Dori',
    select_product: 'Dori tanlang',
    qty: 'Miqdor',
    unit: 'Birlik',
    pack: 'kapsula',
    piece: 'dona',
    add: "Qo'shish",
    add_product: "+ Tovar qo'shish",
    total: 'Jami',
    total_sum: 'Umumiy summa',
    price_per_unit: 'Narx (birlik uchun)',
    sum: "so'm",
    buyer_name: 'Xaridor ismi',
    saving: 'Saqlanmoqda...',
    record_sale: 'Sotuvni yozish',
    confirm: 'Tasdiqlash',
    in_transit: "Yo'lda",
    deliver: 'Berish',
    cancel: 'Bekor qilish',
    write: 'Yozish',
    offline_badge: 'Oflayn',
    positions: 'ta mahsulot',
    no_orders: "Buyurtmalar yo'q",
    no_orders_filter: "Tanlangan filtrlarga ko'ra buyurtmalar topilmadi",
    online_orders_hint: 'Kutilayotgan onlayn buyurtmalar shu yerda ko\'rinadi',
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
    verify_error: "Ism mos kelmadi.",
    edit_items: "Tahrirlash",
    return_edit: 'Qaytarish',
    full_return: "To'liq qaytarish",
    save_items: "Saqlash",
    saving_items: "Saqlanmoqda...",
    cancel_edit: "Bekor qilish",
    view_on_map: "Xaritada",
    vip_free: "O'z bemori (bepul)",
    sale_type: 'Sotuv turi',
    sale_regular: 'Oddiy',
    sale_marketolog: 'Marketolog (qarz)',
    choose_marketolog: 'Marketologni tanlang',
    no_marketologs: "Marketolog yo'q",
    payment_method: "To'lov usuli",
    pay_cash: 'Naqd',
    pay_terminal: 'Terminal',
    pay_card: 'Karta',
    pay_online: 'Onlayn',
    payment_label: "To'lov",
    vip_badge: "O'z bemori · Bepul",
    choose_payment: "To'lov usulini tanlang",
    payment_hint: "Mijoz qanday to'ladi?",
    referral_label: 'Bemor qayerdan',
    referral_ph: 'Mustaqil yoki shifokor ismi',
    referred_by: 'Tavsiya qildi',
    analytics: 'Mening tahlilim',
    a_today: 'Bugun',
    a_week: 'Hafta',
    a_month: 'Oy',
    a_year: 'Yil',
    a_date: 'Sana',
    a_orders: 'Buyurtmalar',
    a_revenue: 'Tushum',
    a_created: 'Yaratilgan',
    a_confirmed: 'Tasdiqlangan',
    by_category: "Toifalar bo'yicha",
    top_products: "Dorilar bo'yicha",
    product_name: 'Dori nomi',
    category: 'Toifa',
    cat_all: 'Barcha toifalar',
    cat_vip: "O'z bemori",
    cat_doctor: 'Shifokordan',
    cat_marketolog: 'Marketolog',
    cat_regular: 'Oddiy',
    sold_ok: 'Sotuv yozildi ✓',
    nav_online: 'Onlayn',
    nav_offline: 'Oflayn',
    nav_analytics: 'Tahlil',
    nav_history: 'Tarix',
    nav_stock: 'Ombor',
    stock_in: 'Tovar kirimi',
    my_stock: 'Mening omborim',
    stock_qty: 'Soni',
    stock_add: "To'ldirish",
    stock_empty: "Ombor bo'sh",
    stock_search: "Omborda qidirish...",
    product_list: "Dorilar ro'yxati",
    price: 'Narx',
    remainder: 'Qoldiq',
    pending_online: 'Kutilayotgan onlayn buyurtmalar',
    pending_offline: 'Kutilayotgan oflayn buyurtmalar',
    history_title: 'Buyurtmalar tarixi',
    status_label: 'Status',
    order_type: 'Buyurtma turi',
    period_label: 'Davr',
    type_all: 'Hammasi',
    type_online: 'Onlayn',
    type_offline: 'Oflayn',
    own_patient: "O'z bemori",
    period_all: 'Butun davr',
    night_mode: 'Tungi',
    day_mode: 'Kunduzgi',
    in_stock: 'Omborda',
    quick_select: 'Tezkor tanlov',
    dynamics: 'Kunlik dinamika',
    export_excel: 'Excel eksport',
  }
}

const txt = computed(() => { watchLang(); return texts[lang.value] || texts.ru })

const tabTitle = computed(() => {
  const t = txt.value
  const m = { online: t.nav_online, offline: t.nav_offline, stock: t.nav_stock, analytics: t.nav_analytics, history: t.history_title }
  return m[tab.value] || t.title
})

// ===== Orders & sections =====
const orders = ref([])
const tab = ref('online')

const historyType = ref('all')
const historyStatus = ref('all')
const historyPeriod = ref('all')
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

const onlineOrders = computed(() =>
  orders.value.filter(o => !o.is_offline && o.status !== 'delivered' && o.status !== 'cancelled')
)

const offlineOrders = computed(() =>
  orders.value.filter(o => o.is_offline && o.status !== 'delivered' && o.status !== 'cancelled')
)

const historyOrders = computed(() => {
  let list = orders.value
  if (historyType.value === 'online') list = list.filter(o => !o.is_offline)
  else if (historyType.value === 'offline') list = list.filter(o => o.is_offline)
  else if (historyType.value === 'vip') list = list.filter(o => o.is_vip)
  if (historyStatus.value !== 'all') list = list.filter(o => o.status === historyStatus.value)
  return list.filter(inPeriod)
})

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function orderTotal(order) {
  return order.items?.reduce((sum, item) => sum + item.price, 0) || 0
}

function boughtItems(order) {
  return (order.items || []).filter(i => i.quantity > 0)
}

function paymentLabel(order) {
  if (order.is_vip) return txt.value.vip_badge
  const t = txt.value
  const m = { cash: t.pay_cash, terminal: t.pay_terminal, card: t.pay_card, online: t.pay_online }
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
  const m = { pending: t.status_pending, confirmed: t.status_confirmed, shipped: t.status_shipped, in_transit: t.status_in_transit, delivered: t.status_delivered, cancelled: t.status_cancelled }
  return m[status] || status
}

function statusClass(status) {
  const m = { pending: 'bg-yellow-100 text-yellow-700', confirmed: 'bg-blue-100 text-blue-700', shipped: 'bg-purple-100 text-purple-700', in_transit: 'bg-orange-100 text-orange-700', delivered: 'bg-green-100 text-green-700', cancelled: 'bg-red-100 text-red-700' }
  return m[status] || 'bg-gray-100 text-gray-700'
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
const editingNurseItems = ref(false)
const editItems = ref([])
const editAddProductId = ref('')
const editAddQty = ref(1)
const savingEditItems = ref(false)

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
    nurseSearchError.value = e.response?.data?.error || 'Заказ не найден'
  } finally { nurseSearching.value = false }
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
    searchError.value = e.response?.data?.error || 'Заказ не найден'
  } finally { searching.value = false }
}

async function updateStatus(order, status) {
  const payload = { status }
  if (status === 'cancelled') {
    const reason = prompt('Укажите причину отмены заказа:')
    if (reason === null) return
    const trimmed = reason.trim()
    if (!trimmed) { alert('Причина обязательна'); return }
    payload.cancellation_reason = trimmed
  }
  try {
    const res = await api.put(`/pickup/orders/${order.id}/status`, payload)
    const idx = orders.value.findIndex(o => o.id === order.id)
    if (idx !== -1) orders.value[idx] = res.data
    if (foundOrder.value?.id === order.id) foundOrder.value = res.data
  } catch (e) { alert(e.response?.data?.error || 'Ошибка при обновлении статуса') }
}

// ===== Payment method modal =====
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
  } catch (e) { alert(e.response?.data?.error || 'Ошибка при подтверждении') }
  finally { paySubmitting.value = false }
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
    const thread = (res.data || []).find(t => t.user_id === order.user_id || t.user?.id === order.user_id)
    if (thread) {
      chatThreadId.value = thread.id
      const detail = await api.get(`/pickup/support/threads/${thread.id}`)
      chatMessages.value = detail.data.messages || []
    }
  } catch (e) { console.error(e) }
  finally {
    chatLoading.value = false
    await nextTick()
    if (chatMessagesEl.value) chatMessagesEl.value.scrollTop = chatMessagesEl.value.scrollHeight
  }
}

async function sendWorkerMessage() {
  if (!chatMsg.value.trim() || chatSending.value) return
  if (!chatThreadId.value) { alert('Пользователь ещё не начал переписку'); return }
  const text = chatMsg.value.trim()
  chatSending.value = true
  try {
    const res = await api.post(`/pickup/support/threads/${chatThreadId.value}/reply`, { message: text })
    chatMessages.value.push(res.data)
    chatMsg.value = ''
    await nextTick()
    if (chatMessagesEl.value) chatMessagesEl.value.scrollTop = chatMessagesEl.value.scrollHeight
  } catch (e) { alert(e.response?.data?.error || 'Ошибка при отправке') }
  finally { chatSending.value = false }
}

// ===== Inline edit for orders list =====
const listEdit = ref({})

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

function cancelListEdit(orderId) { delete listEdit.value[orderId] }
function listEditInc(orderId, idx) { listEdit.value[orderId].items[idx].quantity++ }
function listEditDec(orderId, idx) {
  const item = listEdit.value[orderId].items[idx]
  if (item.quantity > 1) item.quantity--
  else listEdit.value[orderId].items.splice(idx, 1)
}

function listEditAddItem(orderId) {
  const state = listEdit.value[orderId]
  if (!state.addProductId || state.addQty < 1) return
  const product = allProducts.value.find(p => p.id === state.addProductId)
  if (!product) return
  const existing = state.items.find(i => i.product_id === product.id && i.unit_type === state.addUnit)
  if (existing) existing.quantity += state.addQty
  else state.items.push({ product_id: product.id, name: product.name, quantity: state.addQty, unit_type: state.addUnit })
  state.addProductId = ''
  state.addQty = 1
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
  } catch (e) { alert(e.response?.data?.error || 'Ошибка при возврате') }
}

async function saveListEdit(order) {
  const state = listEdit.value[order.id]
  if (!state || state.items.length === 0) return
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
      items: state.items.map(i => ({ product_id: i.product_id, quantity: i.quantity, unit_type: i.unit_type })),
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
const allProducts = ref([])
const offlineProductId = ref('')
const offlineQty = ref(1)
const offlineItems = ref([])
const offlineNote = ref('')
const offlineSubmitting = ref(false)
const offlineSuccess = ref(false)
const saleType = ref('regular')
const offlineMarketolog = ref(null)
const marketologs = ref([])
const offlinePaymentMethod = ref('cash')
const offlineReferral = ref('')
const offlineUnit = ref('pack')
const allDoctors = ref([])

async function loadMarketologs() {
  try { marketologs.value = (await api.get('/pickup/marketologs')).data || [] } catch (e) { console.error(e) }
}

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
  try { allDoctors.value = (await api.get('/doctors')).data || [] } catch (e) { console.error(e) }
}

function addOfflineItem() {
  if (!offlineProductId.value || offlineQty.value < 1) return
  const product = allProducts.value.find(p => p.id === offlineProductId.value)
  if (!product) return
  const unit = offlineUnit.value === 'piece' ? 'piece' : 'pack'
  const qpp = qppOf(product.id)
  const piecesNeeded = unit === 'piece' ? offlineQty.value : offlineQty.value * qpp
  const alreadyPieces = offlineItems.value.filter(i => i.product_id === product.id)
    .reduce((s, i) => s + (i.unit_type === 'piece' ? i.quantity : i.quantity * qpp), 0)
  if (alreadyPieces + piecesNeeded > stockOf(product.id)) {
    alert(`${product.name}: недостаточно на складе`)
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
  offlineSuccess.value = false
  try {
    const isVip = saleType.value === 'vip'
    const isMkt = saleType.value === 'marketolog'
    await api.post('/pickup/offline-sale', {
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
    loadStock()
  } catch (e) { alert(e.response?.data?.error || 'Ошибка при записи') }
  finally { offlineSubmitting.value = false }
}

// ===== Analytics =====
const analyticsPeriod = ref('daily')
const analyticsDate = ref(new Date().toISOString().slice(0, 10))
const analyticsData = ref(null)
const analyticsLoading = ref(false)

async function loadAnalytics() {
  analyticsLoading.value = true
  try {
    const params = { period: analyticsPeriod.value }
    if (analyticsPeriod.value === 'custom') params.date = analyticsDate.value
    analyticsData.value = (await api.get('/pickup/analytics', { params })).data
  } catch (e) { console.error(e) } finally { analyticsLoading.value = false }
}

function selectAnalyticsPeriod(p) {
  analyticsPeriod.value = p
  if (p !== 'custom' || analyticsDate.value) loadAnalytics()
}

const analyticsCat = ref('all')
const emptyCat = { orders: 0, capsules: 0, pieces: 0, revenue: 0 }
function catData(key) { return analyticsData.value?.breakdown?.[key] || emptyCat }

const allCats = computed(() => [
  { key: 'vip', label: txt.value.cat_vip, color: 'text-amber-600' },
  { key: 'doctor', label: txt.value.cat_doctor, color: 'text-purple-600' },
  { key: 'marketolog', label: txt.value.cat_marketolog, color: 'text-indigo-600' },
  { key: 'regular', label: txt.value.cat_regular, color: 'text-emerald-600' },
])

// ===== Warehouse =====
const stock = ref([])
const stockProductId = ref('')
const stockQty = ref(1)
const stockUnit = ref('pack')
const addingStock = ref(false)
const stockSearch = ref('')

const stockMap = computed(() => {
  const m = {}
  for (const s of stock.value) m[s.product_id] = s.quantity
  return m
})
function stockOf(productId) { return displayStock(productId, stockMap.value[productId] || 0) }

const filteredStock = computed(() => {
  if (!stockSearch.value.trim()) return stock.value
  const q = stockSearch.value.toLowerCase()
  return stock.value.filter(s => s.product?.name?.toLowerCase().includes(q))
})

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

// ===== Excel Export =====
function exportHistoryExcel() {
  const rows = historyOrders.value.flatMap(order => {
    const buyerName = order.is_offline
      ? (order.offline_note || '—')
      : ((order.user?.first_name || '') + ' ' + (order.user?.last_name || '')).trim() || '—'
    return boughtItems(order).map(item => ({
      'Дата': new Date(order.created_at).toLocaleString('ru-RU'),
      'Код заказа': order.order_code,
      'Покупатель': buyerName,
      'Тип': order.is_offline ? 'Офлайн' : 'Онлайн',
      'Статус': statusLabel(order.status),
      'Препарат': item.product?.name || '—',
      'Единица': item.unit_type === 'piece' ? 'шт' : 'капс.',
      'Кол-во': item.quantity,
      'Цена за ед. (сум)': item.quantity > 0 ? Math.round(item.price / item.quantity) : 0,
      'Общая стоимость (сум)': Math.round(item.price),
      'Способ оплаты': paymentLabel(order) || '—',
      'Рекомендовал': order.referred_by || '—',
    }))
  })

  if (rows.length === 0) {
    alert('Нет данных для экспорта')
    return
  }

  const ws = XLSX.utils.json_to_sheet(rows)
  ws['!cols'] = [
    { wch: 18 }, { wch: 12 }, { wch: 20 }, { wch: 10 }, { wch: 12 },
    { wch: 22 }, { wch: 8 }, { wch: 8 }, { wch: 18 }, { wch: 22 }, { wch: 14 }, { wch: 18 },
  ]
  const wb = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(wb, ws, 'История заказов')
  XLSX.writeFile(wb, `история_заказов_${new Date().toISOString().slice(0,10)}.xlsx`)
}

function exportAnalyticsExcel() {
  const data = analyticsData.value
  if (!data) { alert('Нет данных для экспорта'); return }

  const summaryRows = [
    { 'Показатель': 'Всего заказов', 'Значение': data.total_orders },
    { 'Показатель': 'Выручка (сум)', 'Значение': data.total_revenue },
    { 'Показатель': 'Создано', 'Значение': data.created_count },
    { 'Показатель': 'Подтверждено', 'Значение': data.confirmed_count },
  ]

  const catRows = allCats.value.map(cat => {
    const d = catData(cat.key)
    return {
      'Категория': cat.label,
      'Заказов': d.orders,
      'Капс.': d.capsules,
      'Шт.': d.pieces,
      'Выручка (сум)': d.revenue,
    }
  })

  const wb = XLSX.utils.book_new()
  const wsSummary = XLSX.utils.json_to_sheet(summaryRows)
  wsSummary['!cols'] = [{ wch: 22 }, { wch: 18 }]
  XLSX.utils.book_append_sheet(wb, wsSummary, 'Итоги')

  const wsCat = XLSX.utils.json_to_sheet(catRows)
  wsCat['!cols'] = [{ wch: 18 }, { wch: 10 }, { wch: 10 }, { wch: 10 }, { wch: 18 }]
  XLSX.utils.book_append_sheet(wb, wsCat, 'По категориям')

  if ((data.by_doctor || []).length) {
    const wsDoc = XLSX.utils.json_to_sheet(data.by_doctor.map(d => ({
      'Доктор': d.name, 'Капс.': d.capsules, 'Шт.': d.pieces, 'Выручка (сум)': d.revenue,
    })))
    XLSX.utils.book_append_sheet(wb, wsDoc, 'По докторам')
  }

  XLSX.writeFile(wb, `аналитика_${new Date().toISOString().slice(0,10)}.xlsx`)
}

watch(tab, (t) => {
  if (t === 'analytics' && !analyticsData.value) loadAnalytics()
  if (t === 'stock') loadStock()
  if (t === 'offline') loadStock()
})

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
  stockPoll = setInterval(() => {
    if (tab.value === 'offline' || tab.value === 'stock') loadStock()
  }, 7000)
})
onUnmounted(() => { if (stockPoll) clearInterval(stockPoll) })
</script>

<style scoped>
/* ── Light mode (default) ── */
.pp-root      { background: #f8fafc; }
.pp-main      { background: #f1f5f9; }
.pp-header    { background: #ffffff; border-bottom: 1px solid #e2e8f0; }
.pp-card      { background: #ffffff; }
.pp-inset     { background: #f8fafc; }
.pp-border-b  { border-bottom: 1px solid #e2e8f0; }
.pp-border-t  { border-top: 1px solid #e2e8f0; }
.pp-border-box{ border: 1px solid #e2e8f0; }
.pp-table-head{ background: #f8fafc; }
.pp-divide    { divide-color: #f1f5f9; }
.pp-cart-row  { background: #f8fafc; border-color: #e2e8f0; }
.pp-chat-bg   { background: #f8fafc; }
.pp-text      { color: #1e293b; }
.pp-text-2    { color: #64748b; }
.pp-text-3    { color: #94a3b8; }

.pp-input {
  border: 1px solid #cbd5e1;
  background: #ffffff;
  color: #1e293b;
}
.pp-input::placeholder { color: #94a3b8; }

.pp-row-hover:hover { background: #f1f5f9; }

/* ── Sidebar – Light mode (default) ── */
.pp-sidebar            { background: #1e293b; border-right: 1px solid rgba(255,255,255,0.06); }
.pp-sidebar-border-b   { border-bottom: 1px solid rgba(255,255,255,0.08); }
.pp-sidebar-border-t   { border-top: 1px solid rgba(255,255,255,0.08); }
.pp-sidebar-text       { color: #f1f5f9; }
.pp-sidebar-text-muted { color: #94a3b8; }
.pp-sidebar-nav-inactive       { color: #94a3b8; }
.pp-sidebar-nav-inactive:hover { background: rgba(255,255,255,0.06); color: #e2e8f0; }
.pp-sidebar-lang-active   { color: #f1f5f9; }
.pp-sidebar-lang-inactive { color: #64748b; }
.pp-sidebar-lang-inactive:hover { color: #94a3b8; }
.pp-sidebar-lang-sep  { color: #334155; }
.pp-sidebar-theme-btn { color: #94a3b8; }
.pp-sidebar-theme-btn:hover { color: #cbd5e1; }

/* ── Sidebar – Dark mode (night-mode) ── */
.night-mode .pp-sidebar            { background: #0f172a; border-right: 1px solid rgba(255,255,255,0.06); }
.night-mode .pp-sidebar-border-b   { border-bottom: 1px solid rgba(255,255,255,0.06); }
.night-mode .pp-sidebar-border-t   { border-top: 1px solid rgba(255,255,255,0.06); }

/* ── Dark mode ── */
.night-mode .pp-root   { background: #0f172a; }
.night-mode .pp-main   { background: #0f172a; }
.night-mode .pp-header { background: #1e293b; border-bottom: 1px solid rgba(255,255,255,0.07); }
.night-mode .pp-card   { background: #1e293b; box-shadow: 0 1px 3px rgba(0,0,0,0.4); }
.night-mode .pp-inset  { background: #0f172a; }
.night-mode .pp-border-b   { border-bottom: 1px solid rgba(255,255,255,0.07); }
.night-mode .pp-border-t   { border-top: 1px solid rgba(255,255,255,0.07); }
.night-mode .pp-border-box { border: 1px solid rgba(255,255,255,0.07); }
.night-mode .pp-table-head { background: #0f172a; }
.night-mode .pp-cart-row   { background: rgba(255,255,255,0.03); border-color: rgba(255,255,255,0.07); }
.night-mode .pp-chat-bg    { background: #0f172a; }
.night-mode .pp-text       { color: #f1f5f9; }
.night-mode .pp-text-2     { color: #94a3b8; }
.night-mode .pp-text-3     { color: #64748b; }
.night-mode .pp-row-hover:hover { background: rgba(255,255,255,0.04); }

.night-mode .pp-input {
  border: 1px solid #334155;
  background: #0f172a;
  color: #f1f5f9;
}
.night-mode .pp-input::placeholder { color: #4b5563; }

/* Dark mode: override hardcoded Tailwind grays inside cards */
.night-mode .text-gray-800 { color: #f1f5f9 !important; }
.night-mode .text-gray-700 { color: #e2e8f0 !important; }
.night-mode .text-gray-600 { color: #cbd5e1 !important; }
.night-mode .text-gray-500 { color: #94a3b8 !important; }
.night-mode .text-gray-400 { color: #64748b !important; }
.night-mode .bg-gray-100   { background: #1e293b !important; }
.night-mode .bg-gray-50    { background: #0f172a !important; }
.night-mode .border-gray-100 { border-color: rgba(255,255,255,0.06) !important; }
.night-mode .border-gray-200 { border-color: rgba(255,255,255,0.08) !important; }
.night-mode .border-t { border-color: rgba(255,255,255,0.07) !important; }
.night-mode .border-b { border-color: rgba(255,255,255,0.07) !important; }
.night-mode .divide-y > * + * { border-color: rgba(255,255,255,0.07) !important; }

/* Modal dark */
.night-mode .relative.bg-white { background: #1e293b !important; }
.night-mode input[type="date"] { color-scheme: dark; }

/* Light mode: fix hardcoded gray text/bg inside cards */
.pp-root:not(.night-mode) .text-gray-800 { color: #1e293b; }
.pp-root:not(.night-mode) .text-gray-700 { color: #374151; }
.pp-root:not(.night-mode) .text-gray-500 { color: #6b7280; }
.pp-root:not(.night-mode) .text-gray-400 { color: #9ca3af; }
.pp-root:not(.night-mode) .bg-gray-100   { background: #f3f4f6; }
.pp-root:not(.night-mode) .bg-gray-50    { background: #f9fafb; }
</style>
