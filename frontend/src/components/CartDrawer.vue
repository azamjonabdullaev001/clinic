<template>
  <Teleport to="body">
    <div v-if="cartStore.isOpen" class="fixed inset-0 z-50">
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm" @click="cartStore.toggle()"></div>
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white/95 backdrop-blur-2xl shadow-2xl flex flex-col animate-slide-right">

        <!-- Header -->
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100">
          <h2 class="text-xl font-serif text-stone-900">
            {{ activeTab === 'cart' ? t.cart_title : t.orders_title }}
          </h2>
          <button @click="cartStore.toggle()" class="p-2 hover:bg-stone-100 rounded-xl hover:rotate-90 transition-all duration-300">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Tabs -->
        <div class="flex border-b border-stone-100 bg-stone-50/60">
          <button
            @click="switchTab('cart')"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-200"
            :class="activeTab === 'cart' ? 'text-brand-700 border-b-2 border-brand-600 bg-white' : 'text-stone-400 hover:text-stone-600'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z" />
            </svg>
            {{ t.cart_title }}
            <span v-if="cartStore.totalItems > 0" class="bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
              {{ cartStore.totalItems }}
            </span>
          </button>
          <button
            v-if="authStore.isLoggedIn"
            @click="switchTab('orders')"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-200"
            :class="activeTab === 'orders' ? 'text-brand-700 border-b-2 border-brand-600 bg-white' : 'text-stone-400 hover:text-stone-600'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            {{ t.orders_title }}
          </button>
        </div>

        <!-- ===== TAB: CART ===== -->
        <template v-if="activeTab === 'cart'">
          <!-- Empty state -->
          <div v-if="cartStore.items.length === 0" class="flex-1 flex flex-col items-center justify-center px-8">
            <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
              <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007zM8.625 10.5a.375.375 0 11-.75 0 .375.375 0 01.75 0zm7.5 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
              </svg>
            </div>
            <p class="text-stone-800 text-lg font-medium mb-1">{{ t.cart_empty }}</p>
            <p class="text-stone-400 text-sm">{{ t.cart_empty_sub }}</p>
          </div>

          <!-- Items -->
          <div v-else class="flex-1 overflow-y-auto px-6 py-5 space-y-4">
            <div
              v-for="item in cartStore.items"
              :key="item.product_id"
              class="bg-stone-50 rounded-2xl p-4 border border-stone-100 hover:border-brand-100 transition-colors duration-300"
            >
              <div class="flex gap-3">
                <div class="w-14 h-14 bg-stone-200 rounded-xl overflow-hidden flex-shrink-0 flex items-center justify-center">
                  <img v-if="item.image_path" :src="item.image_path" class="w-full h-full object-cover" />
                  <svg v-else class="w-6 h-6 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                  </svg>
                </div>
                <div class="flex-1 min-w-0">
                  <div class="flex items-start justify-between gap-2">
                    <h4 class="font-semibold text-stone-900 text-sm truncate">{{ item.name }}</h4>
                    <button @click="cartStore.removeItem(item.product_id)" class="text-stone-300 hover:text-red-400 transition-colors p-0.5 flex-shrink-0">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </button>
                  </div>
                  <p class="text-brand-600 font-semibold text-sm mt-0.5">
                    {{ formatPrice(item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack) }} {{ t.currency }}
                    <span class="text-stone-400 font-normal text-xs">/ {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</span>
                  </p>
                </div>
              </div>

              <div class="flex items-center justify-between mt-3 pt-3 border-t border-stone-200/60">
                <div class="flex items-center gap-1">
                  <button
                    @click="cartStore.updateUnitType(item.product_id, 'piece')"
                    :class="item.unit_type === 'piece' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                    class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                  >{{ t.unit_piece }}</button>
                  <button
                    @click="cartStore.updateUnitType(item.product_id, 'pack')"
                    :class="(item.unit_type || 'pack') === 'pack' ? 'bg-brand-700 text-white shadow-sm' : 'bg-white text-stone-500 border border-stone-200 hover:border-stone-300'"
                    class="px-2.5 py-1 rounded-lg text-xs font-medium transition-all"
                  >{{ t.unit_pack }}</button>
                </div>
                <div class="flex items-center gap-2">
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity - 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600">−</button>
                  <span class="text-sm font-bold w-6 text-center text-stone-900">{{ item.quantity }}</span>
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity + 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center text-sm font-medium transition-colors text-stone-600">+</button>
                </div>
              </div>

              <div class="text-right mt-2">
                <span class="text-xs font-semibold text-stone-500">{{ t.cart_item_total }} <span class="text-brand-700">{{ formatPrice(itemTotal(item)) }} {{ t.currency }}</span></span>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div v-if="cartStore.items.length > 0" class="border-t border-stone-100 px-6 py-5 bg-stone-50/50">
            <div class="flex justify-between items-baseline mb-4">
              <span class="text-sm font-medium text-stone-500">{{ t.cart_total }}</span>
              <span class="text-2xl font-serif text-brand-700">{{ formatPrice(cartStore.totalPrice) }} <span class="text-sm font-sans">{{ t.currency }}</span></span>
            </div>
            <button @click="openCheckout" class="w-full btn-primary py-4 text-base rounded-xl">
              {{ t.cart_checkout }}
            </button>
          </div>
        </template>

        <!-- ===== TAB: MY ORDERS ===== -->
        <template v-else>
          <div v-if="ordersLoading" class="flex-1 flex items-center justify-center">
            <div class="w-8 h-8 border-2 border-brand-600 border-t-transparent rounded-full animate-spin"></div>
          </div>

          <div v-else-if="orders.length === 0" class="flex-1 flex flex-col items-center justify-center px-8 text-center">
            <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
              <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
            <p class="text-stone-800 text-lg font-medium">{{ t.orders_empty }}</p>
          </div>

          <div v-else class="flex-1 overflow-y-auto px-6 py-5 space-y-4">
            <div
              v-for="order in activeOrders"
              :key="order.id"
              class="rounded-2xl border border-stone-100 bg-stone-50 overflow-hidden"
            >
              <div class="flex items-center justify-between px-4 py-3 bg-white border-b border-stone-100">
                <div>
                  <p class="text-[10px] text-stone-400 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                  <p class="text-2xl font-bold text-brand-700 tracking-widest">{{ order.order_code }}</p>
                </div>
                <span class="px-3 py-1.5 rounded-xl text-xs font-semibold" :class="statusClass(order.status)">
                  {{ statusLabel(order.status) }}
                </span>
              </div>

              <div class="px-4 py-3 space-y-2.5">
                <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2">
                  <div class="flex items-center gap-2 min-w-0">
                    <div class="w-9 h-9 rounded-lg bg-stone-200 overflow-hidden flex-shrink-0 flex items-center justify-center">
                      <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover" />
                      <svg v-else class="w-4 h-4 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                      </svg>
                    </div>
                    <div class="min-w-0">
                      <p class="font-medium text-stone-900 text-sm truncate">{{ item.product?.name }}</p>
                      <p class="text-stone-400 text-xs">{{ item.quantity }} {{ item.unit_type === 'piece' ? t.unit_piece : t.unit_pack }}</p>
                    </div>
                  </div>
                  <p class="font-semibold text-stone-700 text-sm flex-shrink-0">{{ formatPrice(item.price) }} {{ t.currency }}</p>
                </div>
              </div>

              <div class="px-4 py-2.5 border-t border-stone-100 flex justify-between items-center">
                <span class="text-xs text-stone-400">{{ t.cart_total }}</span>
                <span class="font-bold text-brand-700">{{ formatPrice(orderTotal(order)) }} {{ t.currency }}</span>
              </div>
            </div>

            <template v-if="cancelledOrders.length > 0">
              <div class="flex items-center gap-2 pt-2">
                <div class="flex-1 h-px bg-red-200"></div>
                <span class="text-xs font-semibold text-red-400 uppercase tracking-wider">{{ t.status_cancelled }}</span>
                <div class="flex-1 h-px bg-red-200"></div>
              </div>
              <div
                v-for="order in cancelledOrders"
                :key="order.id"
                class="rounded-2xl border border-red-200 bg-red-50/60 overflow-hidden"
              >
                <div class="flex items-center justify-between px-4 py-3 bg-red-50 border-b border-red-100">
                  <div>
                    <p class="text-[10px] text-red-300 uppercase tracking-wider mb-0.5">{{ t.orders_code }}</p>
                    <p class="text-2xl font-bold text-red-400 tracking-widest">{{ order.order_code }}</p>
                  </div>
                  <span class="px-3 py-1.5 rounded-xl text-xs font-semibold bg-red-100 text-red-600">{{ t.status_cancelled }}</span>
                </div>
                <div class="px-4 py-3 space-y-2">
                  <div v-for="item in order.items" :key="item.id" class="flex items-center justify-between gap-2 opacity-60">
                    <p class="font-medium text-stone-600 text-sm truncate line-through">{{ item.product?.name }}</p>
                    <p class="font-semibold text-stone-400 text-sm flex-shrink-0 line-through">{{ formatPrice(item.price) }} {{ t.currency }}</p>
                  </div>
                </div>
              </div>
            </template>
          </div>
        </template>

      </div>
    </div>

    <!-- ===== CHECKOUT MODAL ===== -->
    <div v-if="showCheckout" class="fixed inset-0 z-[60] flex items-end sm:items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="showCheckout = false"></div>
      <div class="relative bg-white w-full sm:max-w-lg sm:mx-4 rounded-t-3xl sm:rounded-3xl shadow-2xl max-h-[92vh] overflow-y-auto">

        <!-- Modal Header -->
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100 sticky top-0 bg-white rounded-t-3xl z-10">
          <h3 class="text-lg font-bold text-stone-900">Оформление заказа</h3>
          <button @click="showCheckout = false" class="p-2 hover:bg-stone-100 rounded-xl transition-colors">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="px-6 py-5 space-y-5">

          <!-- Location section -->
          <div>
            <label class="block text-sm font-semibold text-stone-700 mb-2">
              Адрес доставки <span class="text-red-400">*</span>
            </label>

            <!-- Confirmed location banner -->
            <div v-if="locationConfirmed" class="flex items-center gap-3 bg-green-50 border border-green-200 rounded-xl px-4 py-3 mb-3">
              <svg class="w-5 h-5 text-green-600 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/></svg>
              <div class="flex-1 min-w-0">
                <p class="text-xs font-semibold text-green-700 mb-0.5">Локация подтверждена</p>
                <p class="text-xs text-green-600 truncate">{{ checkoutForm.address || `${checkoutForm.lat.toFixed(5)}, ${checkoutForm.lng.toFixed(5)}` }}</p>
              </div>
              <button type="button" @click="locationConfirmed = false; showMap = true; initMap()" class="text-xs text-green-600 hover:text-green-800 font-medium flex-shrink-0">Изменить</button>
            </div>

            <template v-else>
              <!-- Text input with GPS button -->
              <div class="relative mb-3">
                <input
                  v-model="checkoutForm.address"
                  type="text"
                  placeholder="Например: Андижанская область, г. Карасу, ул. Навои 15"
                  class="w-full border rounded-xl px-4 py-3 pr-12 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900 text-sm"
                  :class="locationError ? 'border-red-400 bg-red-50/30' : 'border-stone-200'"
                />
                <button
                  type="button"
                  @click="detectGPS"
                  :disabled="gpsLoading"
                  class="absolute right-2 top-1/2 -translate-y-1/2 p-2 rounded-lg text-stone-400 hover:text-brand-600 hover:bg-brand-50 transition-all disabled:opacity-40"
                  title="Определить моё местоположение"
                >
                  <svg v-if="!gpsLoading" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                  </svg>
                  <svg v-else class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </button>
              </div>

              <!-- Map toggle button -->
              <button
                type="button"
                @click="toggleMap"
                class="flex items-center gap-2 text-sm text-brand-600 hover:text-brand-800 font-medium transition-colors mb-2"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
                </svg>
                {{ showMap ? 'Скрыть карту' : 'Выбрать точку на карте' }}
              </button>

              <!-- Leaflet Map -->
              <div v-show="showMap" class="rounded-2xl overflow-hidden border border-stone-200 mb-2" style="height: 260px">
                <div id="checkout-map" style="height: 100%; width: 100%;"></div>
              </div>

              <!-- Confirm location button (visible when map is open and coords are set) -->
              <button
                v-if="showMap && checkoutForm.lat && checkoutForm.lng"
                type="button"
                @click="confirmLocation"
                class="w-full flex items-center justify-center gap-2 bg-brand-600 text-white py-2.5 rounded-xl font-semibold hover:bg-brand-700 transition-all mb-2 text-sm"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/></svg>
                Подтвердить локацию
              </button>

              <p v-if="locationError" class="text-xs text-red-500 mt-1">{{ locationError }}</p>
              <p v-else-if="checkoutForm.lat && checkoutForm.lng && !showMap" class="text-xs text-green-600 mt-1 flex items-center gap-1">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/></svg>
                Координаты: {{ checkoutForm.lat.toFixed(5) }}, {{ checkoutForm.lng.toFixed(5) }}
              </p>
              <p v-else-if="!showMap" class="text-xs text-stone-400 mt-1">
                Введите адрес вручную, нажмите
                <svg class="inline w-3.5 h-3.5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"/></svg>
                для автоопределения или выберите точку на карте
              </p>
            </template>
          </div>

          <!-- Referral section (REQUIRED) -->
          <div>
            <label class="block text-sm font-semibold text-stone-700 mb-2">
              Откуда вы узнали о нас? <span class="text-red-400">*</span>
            </label>
            <div class="relative">
              <input
                v-model="referralInput"
                type="text"
                placeholder="Имя доктора, 'Самостоятельно' или 'Из рекламы'"
                class="w-full border rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 transition-all text-stone-900 text-sm"
                :class="referralError ? 'border-red-400 bg-red-50/30' : 'border-stone-200'"
                @input="onReferralInput"
                @focus="showReferralDropdown = true"
                @blur="hideReferralDropdown"
              />
              <!-- Autocomplete dropdown -->
              <div
                v-if="showReferralDropdown && (filteredDoctors.length > 0 || quickOptions.length > 0)"
                class="absolute left-0 right-0 top-full mt-1 bg-white border border-stone-200 rounded-xl shadow-lg z-20 overflow-hidden max-h-52 overflow-y-auto"
              >
                <button
                  v-for="opt in quickOptions"
                  :key="opt"
                  type="button"
                  @mousedown.prevent="selectReferral(opt)"
                  class="w-full text-left px-4 py-2.5 text-sm text-stone-600 hover:bg-brand-50 hover:text-brand-700 transition-colors flex items-center gap-2"
                >
                  <svg class="w-3.5 h-3.5 text-stone-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
                  {{ opt }}
                </button>
                <div v-if="filteredDoctors.length > 0 && quickOptions.length > 0" class="border-t border-stone-100"></div>
                <button
                  v-for="doc in filteredDoctors"
                  :key="doc.id"
                  type="button"
                  @mousedown.prevent="selectReferral(doc.name + (doc.specialty ? ' (' + doc.specialty + ')' : ''), true)"
                  class="w-full text-left px-4 py-2.5 hover:bg-brand-50 transition-colors flex items-center gap-3"
                >
                  <div class="w-7 h-7 rounded-full bg-brand-100 flex items-center justify-center flex-shrink-0">
                    <svg class="w-3.5 h-3.5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                  </div>
                  <div class="min-w-0">
                    <p class="text-sm font-semibold text-stone-800">{{ doc.name }}</p>
                    <p v-if="doc.specialty" class="text-xs text-brand-500">{{ doc.specialty }}</p>
                    <p v-else class="text-xs text-stone-400">Врач</p>
                  </div>
                </button>
              </div>
            </div>
            <p v-if="referralError" class="text-xs text-red-500 mt-1">{{ referralError }}</p>
            <p v-else class="text-xs text-stone-400 mt-1">Укажите, кто порекомендовал наши препараты</p>
          </div>

          <!-- Error -->
          <div v-if="checkoutError" class="bg-red-50 text-red-600 text-sm p-3.5 rounded-xl border border-red-100">
            {{ checkoutError }}
          </div>

          <!-- Order summary -->
          <div class="bg-stone-50 rounded-2xl p-4 border border-stone-100">
            <p class="text-xs font-semibold text-stone-500 uppercase tracking-wider mb-3">Ваш заказ</p>
            <div class="space-y-1.5">
              <div v-for="item in cartStore.items" :key="item.product_id" class="flex justify-between text-sm">
                <span class="text-stone-600 truncate mr-2">{{ item.name }} × {{ item.quantity }}</span>
                <span class="font-medium text-stone-700 flex-shrink-0">{{ formatPrice(itemTotal(item)) }} {{ t.currency }}</span>
              </div>
            </div>
            <div class="border-t border-stone-200 mt-3 pt-3 flex justify-between font-bold">
              <span>Итого:</span>
              <span class="text-brand-700">{{ formatPrice(cartStore.totalPrice) }} {{ t.currency }}</span>
            </div>
          </div>

          <!-- Submit button -->
          <button
            @click="submitOrder"
            :disabled="checkoutLoading"
            class="w-full btn-primary py-4 text-base rounded-xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ checkoutLoading ? 'Оформление...' : 'Подтвердить заказ' }}
          </button>
        </div>
      </div>
    </div>

    <!-- ===== ORDER SUCCESS MODAL ===== -->
    <div v-if="showOrderSuccess" class="fixed inset-0 z-[70] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-md" @click="closeOrderSuccess"></div>
      <div class="relative bg-white rounded-3xl p-8 max-w-sm w-full mx-4 text-center shadow-2xl">
        <!-- Success icon -->
        <div class="w-20 h-20 bg-brand-50 rounded-2xl flex items-center justify-center mx-auto mb-5">
          <svg class="w-10 h-10 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-stone-900 mb-2">Заказ оформлен!</h3>
        <p class="text-stone-500 mb-5 leading-relaxed text-sm">Мы свяжемся с вами в ближайшее время для подтверждения и доставки</p>

        <!-- Order code -->
        <div class="bg-gradient-to-br from-brand-50 to-blue-50 border border-brand-100 rounded-2xl p-5 mb-6">
          <p class="text-xs text-brand-500 font-semibold uppercase tracking-wider mb-2">Ваш код заказа</p>
          <p class="text-5xl font-bold text-brand-700 tracking-[0.15em] mb-2">{{ orderSuccessCode }}</p>
          <p class="text-xs text-stone-400">Сообщите этот код при получении заказа</p>
        </div>

        <button @click="closeOrderSuccess" class="w-full btn-primary py-3.5 rounded-xl text-base font-semibold">
          Отлично, понял!
        </button>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, computed, watch, nextTick, onUnmounted } from 'vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useLangStore } from '../stores/lang'
import { useRouter } from 'vue-router'
import { api } from '../stores/auth'
import axios from 'axios'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()
const langStore = useLangStore()
const t = computed(() => langStore.t)

// Tabs
const activeTab = ref('cart')

function switchTab(tab) {
  activeTab.value = tab
  if (tab === 'orders') loadOrders()
}

watch(() => cartStore.isOpen, (val) => {
  if (!val) {
    activeTab.value = 'cart'
    showCheckout.value = false
  }
})

// Orders
const orders = ref([])
const ordersLoading = ref(false)

async function loadOrders() {
  ordersLoading.value = true
  try {
    const res = await api.get('/orders')
    orders.value = res.data || []
  } catch {
    orders.value = []
  } finally {
    ordersLoading.value = false
  }
}

const activeOrders = computed(() => orders.value.filter(o => o.status !== 'cancelled'))
const cancelledOrders = computed(() => orders.value.filter(o => o.status === 'cancelled'))

function orderTotal(order) {
  return (order.items || []).reduce((sum, i) => sum + i.price, 0)
}

function statusLabel(status) {
  const map = {
    pending: t.value.status_pending,
    confirmed: t.value.status_confirmed,
    shipped: t.value.status_confirmed,
    delivered: t.value.status_delivered,
    cancelled: t.value.status_cancelled,
  }
  return map[status] || status
}

function statusClass(status) {
  const map = {
    pending: 'bg-yellow-100 text-yellow-700',
    confirmed: 'bg-blue-100 text-blue-700',
    shipped: 'bg-blue-100 text-blue-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-600',
  }
  return map[status] || 'bg-stone-100 text-stone-600'
}

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

function itemTotal(item) {
  const unitPrice = item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack
  return unitPrice * item.quantity
}

// ===== CHECKOUT =====
const showCheckout = ref(false)
const checkoutLoading = ref(false)
const checkoutError = ref('')
const locationError = ref('')
const referralError = ref('')
const gpsLoading = ref(false)
const showMap = ref(false)
const locationConfirmed = ref(false)

// Order success
const showOrderSuccess = ref(false)
const orderSuccessCode = ref('')

const checkoutForm = ref({
  address: '',
  lat: 0,
  lng: 0,
})

// Referral autocomplete
const referralInput = ref('')
const showReferralDropdown = ref(false)
const doctors = ref([])

const quickOptions = computed(() => {
  const q = referralInput.value.toLowerCase()
  const opts = ['Самостоятельно', 'Из рекламы']
  if (!q) return opts
  return opts.filter(o => o.toLowerCase().includes(q))
})

const filteredDoctors = computed(() => {
  const q = referralInput.value.toLowerCase()
  if (!q) return doctors.value.slice(0, 5)
  return doctors.value.filter(d => d.name.toLowerCase().includes(q)).slice(0, 5)
})

function onReferralInput() {
  showReferralDropdown.value = true
}

function hideReferralDropdown() {
  setTimeout(() => { showReferralDropdown.value = false }, 150)
}

function selectReferral(val, isDoctor = false) {
  referralInput.value = isDoctor ? ('Доктор: ' + val) : val
  referralError.value = ''
  showReferralDropdown.value = false
}

function confirmLocation() {
  if (!checkoutForm.value.lat || !checkoutForm.value.lng) {
    locationError.value = 'Выберите точку на карте'
    return
  }
  locationConfirmed.value = true
  locationError.value = ''
  destroyMap()
  showMap.value = false
}

async function loadDoctors() {
  try {
    const res = await axios.get('/api/doctors')
    doctors.value = res.data || []
  } catch {
    doctors.value = []
  }
}

// Leaflet map instance
let leafletMap = null
let leafletMarker = null

async function initMap() {
  await nextTick()
  const L = await import('leaflet')
  await import('leaflet/dist/leaflet.css')

  // Fix default icon paths
  delete L.Icon.Default.prototype._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })

  const mapEl = document.getElementById('checkout-map')
  if (!mapEl || leafletMap) return

  const defaultLat = checkoutForm.value.lat || 40.9983
  const defaultLng = checkoutForm.value.lng || 71.6726

  leafletMap = L.map('checkout-map').setView([defaultLat, defaultLng], checkoutForm.value.lat ? 15 : 13)
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
  }).addTo(leafletMap)

  if (checkoutForm.value.lat && checkoutForm.value.lng) {
    leafletMarker = L.marker([checkoutForm.value.lat, checkoutForm.value.lng]).addTo(leafletMap)
  }

  leafletMap.on('click', async (e) => {
    const { lat, lng } = e.latlng
    checkoutForm.value.lat = lat
    checkoutForm.value.lng = lng

    if (leafletMarker) {
      leafletMarker.setLatLng([lat, lng])
    } else {
      leafletMarker = L.marker([lat, lng]).addTo(leafletMap)
    }

    // Reverse geocode
    try {
      const res = await fetch(
        `https://nominatim.openstreetmap.org/reverse?lat=${lat}&lon=${lng}&format=json`,
        { headers: { 'Accept-Language': 'ru' } }
      )
      const data = await res.json()
      if (data?.display_name) {
        checkoutForm.value.address = data.display_name
        locationError.value = ''
      }
    } catch { /* ignore */ }
  })
}

function destroyMap() {
  if (leafletMap) {
    leafletMap.remove()
    leafletMap = null
    leafletMarker = null
  }
}

async function toggleMap() {
  showMap.value = !showMap.value
  if (showMap.value) {
    await nextTick()
    await initMap()
  } else {
    destroyMap()
  }
}

async function detectGPS() {
  if (!navigator.geolocation) {
    locationError.value = 'Геолокация не поддерживается браузером'
    return
  }
  gpsLoading.value = true
  locationError.value = ''
  navigator.geolocation.getCurrentPosition(
    async (pos) => {
      const { latitude, longitude } = pos.coords
      checkoutForm.value.lat = latitude
      checkoutForm.value.lng = longitude
      try {
        const res = await fetch(
          `https://nominatim.openstreetmap.org/reverse?lat=${latitude}&lon=${longitude}&format=json`,
          { headers: { 'Accept-Language': 'ru' } }
        )
        const data = await res.json()
        if (data?.address) {
          const a = data.address
          const parts = []
          if (a.state) parts.push(a.state)
          if (a.county || a.district) parts.push(a.county || a.district)
          if (a.city || a.town || a.village) parts.push(a.city || a.town || a.village)
          if (a.road) parts.push(a.road)
          checkoutForm.value.address = parts.join(', ') || data.display_name
        }
      } catch { /* ignore */ }
      // Update map if open
      if (showMap.value && leafletMap) {
        const L = await import('leaflet')
        leafletMap.setView([latitude, longitude], 15)
        if (leafletMarker) {
          leafletMarker.setLatLng([latitude, longitude])
        } else {
          leafletMarker = L.marker([latitude, longitude]).addTo(leafletMap)
        }
      }
      gpsLoading.value = false
    },
    (err) => {
      gpsLoading.value = false
      if (err.code === 1) locationError.value = 'Доступ к геолокации запрещён'
      else locationError.value = 'Не удалось получить местоположение'
    },
    { timeout: 10000 }
  )
}

function openCheckout() {
  if (!authStore.isLoggedIn) {
    cartStore.toggle()
    router.push('/login')
    return
  }
  checkoutError.value = ''
  locationError.value = ''
  referralError.value = ''
  checkoutForm.value = { address: '', lat: 0, lng: 0 }
  referralInput.value = ''
  showMap.value = false
  locationConfirmed.value = false
  destroyMap()
  loadDoctors()
  showCheckout.value = true
}

function closeOrderSuccess() {
  showOrderSuccess.value = false
  orderSuccessCode.value = ''
  switchTab('orders')
  loadOrders()
}

async function submitOrder() {
  checkoutError.value = ''
  locationError.value = ''
  referralError.value = ''

  const hasAddress = checkoutForm.value.address.trim().length > 0
  const hasCoords = checkoutForm.value.lat !== 0 && checkoutForm.value.lng !== 0

  if (!hasAddress && !hasCoords) {
    locationError.value = 'Укажите адрес доставки или выберите точку на карте'
    return
  }

  if (!referralInput.value.trim()) {
    referralError.value = 'Укажите, откуда вы узнали о нас — это обязательное поле'
    return
  }

  checkoutLoading.value = true
  try {
    const items = cartStore.items.map(item => ({
      product_id: item.product_id,
      quantity: item.quantity,
      unit_type: item.unit_type || 'pack'
    }))

    const res = await api.post('/orders', {
      items,
      phone: authStore.user.phone,
      delivery_address: checkoutForm.value.address.trim(),
      latitude: checkoutForm.value.lat,
      longitude: checkoutForm.value.lng,
      referred_by: referralInput.value.trim(),
    })

    cartStore.clear()
    showCheckout.value = false
    destroyMap()
    orderSuccessCode.value = res.data?.order_code || ''
    showOrderSuccess.value = true
  } catch (e) {
    checkoutError.value = e.response?.data?.error || 'Ошибка при оформлении заказа'
  } finally {
    checkoutLoading.value = false
  }
}

onUnmounted(() => {
  destroyMap()
})

function showOrders() {
  activeTab.value = 'orders'
  loadOrders()
}

defineExpose({ showOrders })
</script>
