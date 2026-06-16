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
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.263-12m7.5 0a2.25 2.25 0 00-1.359-2.081m0 0C9.75 3.75 9 3 8.25 3H5.25m0 0A2.25 2.25 0 003 5.25m0 0v10.5A2.25 2.25 0 005.25 18h13.5A2.25 2.25 0 0021 15.75V5.25m-21 0h21m0 0V3.75" />
            </svg>
            {{ t.cart_title }}
            <span v-if="cartStore.totalItems > 0" class="bg-brand-600 text-white text-[10px] font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
              {{ cartStore.totalItems }}
            </span>
          </button>
          <button
            v-if="authStore.isLoggedIn"
            @click="switchTab('orders')"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-200 relative"
            :class="activeTab === 'orders' ? 'text-brand-700 border-b-2 border-brand-600 bg-white' : 'text-stone-400 hover:text-stone-600'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v12a2 2 0 01-2 2h-2.5" />
            </svg>
            {{ t.orders_title }}
            <!-- Unread worker note badge -->
            <span v-if="hasUnreadNotes && activeTab !== 'orders'" class="absolute top-1.5 right-4 w-2.5 h-2.5 bg-red-500 rounded-full animate-pulse"></span>
          </button>
        </div>

        <!-- ===== TAB: CART ===== -->
        <template v-if="activeTab === 'cart'">
          <!-- Empty state -->
          <div v-if="cartStore.items.length === 0" class="flex-1 flex flex-col items-center justify-center px-8">
            <div class="w-20 h-20 bg-stone-50 rounded-2xl flex items-center justify-center mb-5">
              <svg class="w-9 h-9 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.263-12m7.5 0a2.25 2.25 0 00-1.359-2.081m0 0C9.75 3.75 9 3 8.25 3H5.25m0 0A2.25 2.25 0 003 5.25m0 0v10.5A2.25 2.25 0 005.25 18h13.5A2.25 2.25 0 0021 15.75V5.25m-21 0h21m0 0V3.75" />
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
                  <img v-if="item.image_path" :src="item.image_path" class="w-full h-full object-cover" alt="Product" />
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
                    {{ formatPrice(item.price_per_pack) }} {{ t.currency }}
                    <span class="text-stone-400 font-normal text-xs">/ {{ t.unit_pack }}</span>
                  </p>
                </div>
              </div>

              <div class="flex items-center justify-between mt-3 pt-3 border-t border-stone-200/60">
                <span class="text-xs font-medium text-stone-400">{{ item.quantity_per_pack }} {{ t.unit_piece }} / {{ t.unit_pack }}</span>
                <div class="flex items-center gap-2">
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity - 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center transition-colors disabled:opacity-50" :disabled="item.quantity <= 1">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M20 12H4" /></svg>
                  </button>
                  <span class="text-sm font-bold w-6 text-center text-stone-900">{{ item.quantity }}</span>
                  <button @click="cartStore.updateQuantity(item.product_id, item.quantity + 1)" class="w-8 h-8 rounded-lg bg-white border border-stone-200 hover:border-stone-300 flex items-center justify-center transition-colors">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" /></svg>
                  </button>
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
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v12a2 2 0 01-2 2h-2.5" />
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
                      <img v-if="item.product?.image_path" :src="item.product.image_path" class="w-full h-full object-cover" alt="Product" />
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

              <!-- Payment banner for awaiting_payment orders -->
              <div v-if="order.status === 'awaiting_payment'" class="mx-4 mb-3 bg-amber-50 border border-amber-200 rounded-xl px-3 py-2.5">
                <p class="text-sm font-semibold text-amber-800 mb-1">{{ t.status_awaiting_payment }}</p>
                <p class="text-xs text-amber-600 mb-3">{{ t.payment_scan_and_upload }}</p>
                <div class="flex justify-center mb-3">
                  <img
                    v-if="!ordersTabQrFailed"
                    src="/images/QRpayment.jpg"
                    class="w-48 h-48 object-contain rounded-xl border border-amber-200 bg-white p-1"
                    @error="ordersTabQrFailed = true"
                  />
                  <div v-else class="w-48 h-48 flex flex-col items-center justify-center bg-white rounded-xl border border-amber-200 text-center px-2">
                    <svg class="w-8 h-8 text-amber-300 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z"/></svg>
                    <p class="text-xs text-amber-500">{{ t.payment_qr_not_configured }}</p>
                    <button @click="ordersTabQrFailed = false" class="text-xs text-amber-600 font-medium mt-1 underline">{{ t.qr_retry }}</button>
                  </div>
                </div>
                <input :ref="el => ordersTabReceiptInputs[order.id] = el" type="file" accept="image/*" class="hidden" @change="e => uploadReceiptFromOrdersTab(order.id, e)"/>
                <button @click="ordersTabReceiptInputs[order.id]?.click()" :disabled="ordersTabUploading === order.id" class="w-full bg-amber-600 text-white text-sm px-4 py-2.5 rounded-lg hover:bg-amber-700 transition disabled:opacity-50 font-medium">
                  {{ ordersTabUploading === order.id ? t.uploading_text : t.upload_receipt_btn }}
                </button>
              </div>

              <div class="px-4 py-2.5 border-t border-stone-100 flex justify-between items-center">
                <span class="text-xs text-stone-400">{{ t.cart_total }}</span>
                <span class="font-bold text-brand-700">{{ formatPrice(orderTotal(order)) }} {{ t.currency }}</span>
              </div>

              <!-- Receipts + add more for pending online orders (split payment) -->
              <div v-if="order.status === 'pending' && !order.is_offline && !order.is_vip" class="mx-4 mb-3 border border-emerald-100 rounded-xl bg-emerald-50/60 px-3 py-2.5">
                <!-- Uploaded receipts thumbnails -->
                <div v-if="parseOrderReceiptPaths(order).length > 0" class="mb-3">
                  <p class="text-xs font-semibold text-emerald-700 mb-2">✓ {{ t.receipt_uploaded_count(parseOrderReceiptPaths(order).length) }}</p>
                  <div class="flex flex-wrap gap-2">
                    <div v-for="(path, i) in parseOrderReceiptPaths(order)" :key="i" class="relative">
                      <a :href="path" target="_blank" rel="noopener"><img :src="path" class="w-14 h-14 object-cover rounded-xl border-2 border-emerald-300 hover:opacity-80 transition"/></a>
                      <span v-if="parseOrderReceiptPaths(order).length > 1" class="absolute -top-1.5 -right-1.5 bg-emerald-500 text-white text-[9px] font-bold rounded-full w-4 h-4 flex items-center justify-center">{{ i+1 }}</span>
                    </div>
                  </div>
                </div>
                <!-- Add another receipt -->
                <p class="text-xs text-stone-500 mb-2">{{ t.split_payment_hint }}</p>
                <input :ref="el => ordersTabExtraInputs[order.id] = el" type="file" accept="image/*" class="hidden" @change="e => onOrdersTabExtraSelect(order.id, e)"/>
                <div v-if="ordersTabExtraPreviews[order.id]" class="mb-2"><img :src="ordersTabExtraPreviews[order.id]" class="w-full max-h-28 object-contain rounded-xl border border-stone-200"/></div>
                <button @click="ordersTabExtraInputs[order.id]?.click()" :disabled="(ordersTabCooldowns[order.id] || 0) > 0 || ordersTabAddingId === order.id" class="w-full border-2 border-dashed border-emerald-300 rounded-xl py-2 text-xs text-emerald-700 hover:border-emerald-500 transition disabled:opacity-40 disabled:cursor-not-allowed mb-1.5">
                  {{ t.add_another_receipt }}
                </button>
                <p v-if="(ordersTabCooldowns[order.id] || 0) > 0" class="text-xs text-amber-500 text-center mb-1.5">{{ t.receipt_cooldown(ordersTabCooldowns[order.id]) }}</p>
                <button v-if="ordersTabExtraFiles[order.id]" @click="sendOrdersTabExtra(order.id)" :disabled="ordersTabAddingId === order.id" class="w-full bg-emerald-600 text-white text-xs py-2 rounded-xl font-semibold disabled:opacity-50 mb-1 hover:bg-emerald-700 transition">
                  {{ ordersTabAddingId === order.id ? t.adding_receipt : t.payment_confirm }}
                </button>
                <p v-if="ordersTabExtraErrors[order.id]" class="text-xs text-red-500 text-center">{{ ordersTabExtraErrors[order.id] }}</p>
              </div>

              <!-- Delivery address text (always visible) -->
              <div v-if="order.delivery_address" class="px-4 py-2 border-t border-stone-100 flex items-start gap-2">
                <svg class="w-3.5 h-3.5 text-stone-400 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/><circle cx="12" cy="11" r="3"/></svg>
                <p class="text-xs text-stone-500 leading-relaxed">{{ order.delivery_address }}</p>
              </div>

              <!-- Delivery location map + change location for pending orders -->
              <div v-if="order.latitude && order.longitude || order.delivery_address" class="px-4 py-3 border-t border-stone-100 flex flex-col gap-2">
                <button @click="openOrderMap(order)" class="flex items-center justify-center gap-2 w-full bg-blue-600 hover:bg-blue-700 text-white text-sm font-semibold py-2.5 rounded-xl transition">
                  <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                  {{ t.delivery_map_link }}
                </button>
                <button v-if="order.status === 'pending' || order.status === 'awaiting_payment'" @click="openEditLocation(order)" class="flex items-center justify-center gap-2 w-full bg-amber-500 hover:bg-amber-600 text-white text-sm font-semibold py-2.5 rounded-xl transition">
                  <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
                  {{ t.change_location_btn }}
                </button>
              </div>

              <!-- Worker notes — highlighted when unread -->
              <div v-if="order.worker_notes" class="mx-4 my-2 rounded-xl px-3 py-2.5 border" :class="isNoteUnread(order) ? 'bg-red-50 border-red-300 animate-pulse-once' : 'bg-indigo-50 border-indigo-200'">
                <p class="text-xs font-semibold mb-1 flex items-center gap-1" :class="isNoteUnread(order) ? 'text-red-700' : 'text-indigo-700'">
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"/></svg>
                  {{ isNoteUnread(order) ? t.unread_note_label : t.worker_msg_title }}
                  <span v-if="isNoteUnread(order)" class="ml-1 bg-red-500 text-white text-[9px] font-bold rounded-full px-1.5 py-0.5">!</span>
                </p>
                <p class="text-sm whitespace-pre-wrap leading-relaxed" :class="isNoteUnread(order) ? 'text-red-900 font-medium' : 'text-indigo-900'">{{ order.worker_notes }}</p>
              </div>

              <!-- BTS cargo info -->
              <div v-if="order.bts_pickup_point || order.bts_tracking_number" class="mx-4 my-2 rounded-2xl overflow-hidden border border-blue-200 shadow-sm">
                <div class="bg-gradient-to-r from-blue-600 to-blue-500 px-3 py-2 flex items-center gap-2">
                  <svg class="w-4 h-4 text-white flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                  <span class="text-xs font-bold text-white tracking-wide">{{ t.bts_cargo_label }}</span>
                  <span class="ml-auto text-[10px] text-blue-200 font-medium">{{ t.bts_in_transit }}</span>
                </div>
                <div class="bg-blue-50 px-3 py-2.5 space-y-1.5">
                  <div v-if="order.bts_pickup_point" class="flex items-start gap-1.5">
                    <svg class="w-3.5 h-3.5 text-blue-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/><circle cx="12" cy="11" r="3"/></svg>
                    <p class="text-xs text-blue-900 font-semibold leading-snug">{{ order.bts_pickup_point }}</p>
                  </div>
                  <div v-if="order.bts_tracking_number" class="flex items-center gap-1.5">
                    <svg class="w-3.5 h-3.5 text-blue-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/></svg>
                    <p class="text-xs text-blue-700 font-mono">{{ order.bts_tracking_number }}</p>
                  </div>
                  <button
                    v-if="order.bts_branch_lat && order.bts_branch_lng"
                    @click="openBtsMap(order)"
                    class="w-full flex items-center justify-center gap-1.5 bg-blue-600 text-white text-xs font-semibold px-3 py-2 rounded-xl hover:bg-blue-700 transition mt-1"
                  >
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
                    {{ t.bts_map_btn }}
                  </button>
                </div>
              </div>

              <!-- Delete from history button (all orders) -->
              <div class="px-4 pb-3 pt-1 flex justify-end border-t border-stone-50">
                <button @click="confirmHide(order)" class="text-xs text-stone-300 hover:text-red-500 transition flex items-center gap-1">
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
                  {{ t.delete_from_history }}
                </button>
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
                  <!-- Worker notes remain visible even after cancellation (explain the reason) -->
                  <div v-if="order.worker_notes" class="mt-2 pt-2 border-t border-red-100">
                    <div class="bg-indigo-50 border border-indigo-200 rounded-xl px-3 py-2.5">
                      <p class="text-xs font-semibold text-indigo-700 mb-1 flex items-center gap-1">
                        <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"/></svg>
                        {{ t.worker_msg_title }}
                      </p>
                      <p class="text-sm text-indigo-900 whitespace-pre-wrap leading-relaxed">{{ order.worker_notes }}</p>
                    </div>
                  </div>
                  <!-- BTS info visible even after cancellation -->
                  <div v-if="order.bts_pickup_point || order.bts_tracking_number" class="mt-2 pt-2 border-t border-red-100">
                    <div class="rounded-xl overflow-hidden border border-blue-200 opacity-80">
                      <div class="bg-gradient-to-r from-blue-500 to-blue-400 px-3 py-1.5 flex items-center gap-2">
                        <svg class="w-3.5 h-3.5 text-white flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/></svg>
                        <span class="text-xs font-bold text-white">{{ t.bts_cargo_label }}</span>
                      </div>
                      <div class="bg-blue-50 px-3 py-2 space-y-1">
                        <p v-if="order.bts_pickup_point" class="text-xs text-blue-900 font-medium">📍 {{ order.bts_pickup_point }}</p>
                        <p v-if="order.bts_tracking_number" class="text-xs text-blue-600 font-mono">{{ order.bts_tracking_number }}</p>
                      </div>
                    </div>
                  </div>
                  <div v-if="order.cancellation_reason" class="mt-3 pt-3 border-t border-red-200">
                    <p class="text-xs font-semibold text-red-600">{{ t.cancellation_reason }}:</p>
                    <p class="text-xs text-red-500 mt-1">{{ order.cancellation_reason }}</p>
                  </div>
                </div>
              </div>
            </template>
          </div>
        </template>

      </div>
    </div>

    <!-- ===== ORDER MAP MODAL ===== -->
    <div v-if="showOrderMap" class="fixed inset-0 z-[75] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-md" @click="closeOrderMap"></div>
      <div class="relative bg-white rounded-3xl p-0 max-w-2xl w-full mx-4 shadow-2xl max-h-[90vh] overflow-hidden flex flex-col">
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100 bg-white">
          <div>
            <h3 class="text-lg font-bold text-stone-900">{{ t.delivery_order_title }}</h3>
            <p class="text-xs text-stone-400 mt-0.5">{{ selectedOrderForMap?.order_code }}</p>
          </div>
          <button @click="closeOrderMap" class="p-2 hover:bg-stone-100 rounded-xl transition-colors">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <iframe
          v-if="orderGoogleMapUrl"
          :src="orderGoogleMapUrl"
          class="flex-1 min-h-0"
          style="min-height:220px;width:100%;border:0"
          loading="lazy"
          referrerpolicy="no-referrer-when-downgrade"
          allowfullscreen
        ></iframe>
        <div v-else class="flex-1 flex items-center justify-center text-stone-400 text-sm min-h-[200px]">Адрес не указан</div>
        <div class="px-6 py-3 border-t border-stone-100 bg-stone-50 text-xs text-stone-500">
          <span v-if="selectedOrderForMap?.delivery_address">📍 {{ selectedOrderForMap.delivery_address }}</span>
        </div>
      </div>
    </div>

    <!-- ===== BTS MAP MODAL ===== -->
    <div v-if="showBtsMap" class="fixed inset-0 z-[76] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-md" @click="closeBtsMap"></div>
      <div class="relative bg-white rounded-3xl p-0 max-w-2xl w-full mx-4 shadow-2xl max-h-[90vh] overflow-hidden flex flex-col">
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100 bg-white">
          <div>
            <h3 class="text-lg font-bold text-stone-900">{{ t.bts_cargo_label }}</h3>
            <p class="text-xs text-stone-400 mt-0.5">{{ btsMapOrder?.order_code }}</p>
          </div>
          <button @click="closeBtsMap" class="p-2 hover:bg-stone-100 rounded-xl transition-colors">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <iframe
          v-if="btsGoogleMapUrl"
          :src="btsGoogleMapUrl"
          class="flex-1 min-h-0"
          style="min-height:220px;width:100%;border:0"
          loading="lazy"
          referrerpolicy="no-referrer-when-downgrade"
          allowfullscreen
        ></iframe>
        <div v-if="btsMapOrder" class="px-5 py-3 bg-blue-50 border-t border-blue-100 text-xs text-blue-800">
          📍 {{ btsMapOrder.bts_pickup_point || t.bts_pickup_point_label }}<span v-if="btsMapOrder.bts_tracking_number"> · {{ t.bts_track_info }}: {{ btsMapOrder.bts_tracking_number }}</span>
        </div>
      </div>
    </div>

    <!-- ===== CHECKOUT MODAL ===== -->
    <div v-if="showCheckout" class="fixed inset-0 z-[60] flex items-end sm:items-center justify-center">
      <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="closeCheckout()"></div>
      <div class="relative bg-white w-full sm:max-w-lg sm:mx-4 rounded-t-3xl sm:rounded-3xl shadow-2xl max-h-[92vh] overflow-y-auto">

        <!-- Modal Header -->
        <div class="flex items-center justify-between px-6 py-4 border-b border-stone-100 sticky top-0 bg-white rounded-t-3xl z-10">
          <h3 class="text-lg font-bold text-stone-900">{{ t.checkout_title }}</h3>
          <button @click="closeCheckout()" class="p-2 hover:bg-stone-100 rounded-xl transition-colors">
            <svg class="w-5 h-5 text-stone-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="px-6 py-5 space-y-5">

          <!-- Location section -->
          <div>
            <label class="block text-sm font-semibold text-stone-700 mb-2">
              {{ t.checkout_delivery_address }} <span class="text-red-400">*</span>
            </label>

            <!-- Confirmed location banner -->
            <div v-if="locationConfirmed" class="flex items-center gap-3 bg-green-50 border border-green-200 rounded-xl px-4 py-3 mb-3">
              <svg class="w-5 h-5 text-green-600 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" /></svg>
              <div class="flex-1 min-w-0">
                <p class="text-xs font-semibold text-green-700 mb-0.5">{{ t.checkout_location_confirmed }}</p>
                <p class="text-xs text-green-600 truncate">{{ checkoutForm.address || `${checkoutForm.lat.toFixed(5)}, ${checkoutForm.lng.toFixed(5)}` }}</p>
              </div>
              <button type="button" @click="locationConfirmed = false; showMap = true; initMap()" class="text-xs text-green-600 hover:text-green-800 font-medium flex-shrink-0">{{ t.checkout_change }}</button>
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
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2m0 0V3.75" />
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
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 7" />
                </svg>
                {{ showMap ? t.checkout_hide_map : t.checkout_select_on_map }}
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
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" /></svg>
                {{ t.checkout_confirm_location }}
              </button>

              <p v-if="locationError" class="text-xs text-red-500 mt-1">{{ locationError }}</p>
              <p v-else-if="checkoutForm.lat && checkoutForm.lng && !showMap" class="text-xs text-green-600 mt-1 flex items-center gap-1">
                <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" /></svg>
                Координаты: {{ checkoutForm.lat.toFixed(5) }}, {{ checkoutForm.lng.toFixed(5) }}
              </p>
              <p v-else-if="!showMap" class="text-xs text-stone-400 mt-1">
                Введите адрес вручную, нажмите
                <svg class="inline w-3.5 h-3.5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                для автоопределения или выберите точку на карте
              </p>
            </template>
          </div>

          <!-- Referral section (REQUIRED) -->
          <div>
            <label class="block text-sm font-semibold text-stone-700 mb-2">
              {{ t.checkout_referral }} <span class="text-red-400">*</span>
            </label>
            <div class="relative">
              <input
                v-model="referralInput"
                type="text"
                :placeholder="t.checkout_referral_placeholder"
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
                  <svg class="w-3.5 h-3.5 text-stone-400 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
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
                    <svg class="w-3.5 h-3.5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" /></svg>
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
            <p v-else class="text-xs text-stone-400 mt-1">{{ t.checkout_referral_placeholder }}</p>
          </div>

          <!-- Error -->
          <div v-if="checkoutError" class="bg-red-50 text-red-600 text-sm p-3.5 rounded-xl border border-red-100">
            {{ checkoutError }}
          </div>

          <!-- Order summary -->
          <div class="bg-stone-50 rounded-2xl p-4 border border-stone-100">
            <p class="text-xs font-semibold text-stone-500 uppercase tracking-wider mb-3">{{ t.checkout_order_summary }}</p>
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
            {{ checkoutLoading ? t.checkout_confirming : t.checkout_confirm_button }}
          </button>
        </div>
      </div>
    </div>

    <!-- ===== ORDER SUCCESS MODAL ===== -->
    <div v-if="showOrderSuccess" class="fixed inset-0 z-[70] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-md" @click="paymentStage==='done' && closeOrderSuccess()"></div>
      <div class="relative bg-white rounded-3xl p-7 max-w-sm w-full mx-4 text-center shadow-2xl max-h-[92vh] overflow-y-auto">

        <!-- Stage 1: pay via QR + upload receipt -->
        <template v-if="paymentStage === 'pay'">
          <h3 class="text-xl font-bold text-stone-900 mb-1">{{ t.payment_title }}</h3>
          <p class="text-stone-500 mb-4 text-sm">{{ t.payment_subtitle }}</p>

          <!-- QR Code — hide after first receipt sent -->
          <template v-if="uploadedReceipts.length === 0">
            <div class="bg-stone-50 border border-stone-100 rounded-2xl p-4 mb-3 flex items-center justify-center min-h-60">
              <img
                v-if="!successQrFailed"
                src="/images/QRpayment.jpg"
                :alt="t.payment_title"
                class="w-60 h-60 object-contain"
                @error="successQrFailed = true"
              />
              <div v-else class="w-60 h-60 flex flex-col items-center justify-center text-center px-3 gap-2">
                <p class="text-stone-400 text-sm">{{ t.payment_qr_not_configured }}</p>
                <button @click="successQrFailed = false" class="text-brand-600 hover:text-brand-700 text-sm font-medium">{{ t.qr_retry }}</button>
              </div>
            </div>
            <p class="text-sm font-semibold text-emerald-700 bg-emerald-50 border border-emerald-100 rounded-xl px-3 py-2 mb-3">{{ t.payment_qr_hint }}</p>
            <p class="text-xs text-stone-400 mb-4">{{ t.payment_order_code }}: <span class="font-bold text-brand-700 tracking-wider">{{ orderSuccessCode }}</span></p>

            <!-- Receipt upload (multiple) -->
            <div class="mb-4 text-left">
              <label class="text-sm font-medium text-stone-700 mb-2 block">{{ t.payment_upload_receipt }} <span class="text-red-400">*</span></label>
              <div v-if="receiptPreviews.length" class="flex flex-wrap gap-2 mb-3">
                <div v-for="(preview, i) in receiptPreviews" :key="i" class="relative">
                  <img :src="preview" class="w-16 h-16 object-cover rounded-xl border border-stone-200" />
                  <span class="absolute -top-1.5 -right-1.5 bg-brand-600 text-white text-[10px] font-bold rounded-full w-5 h-5 flex items-center justify-center">{{ i+1 }}</span>
                </div>
              </div>
              <button type="button" @click="$refs.receiptInput.click()" class="w-full border-2 border-dashed border-stone-300 rounded-xl py-3 text-sm text-stone-500 hover:border-brand-400 transition-colors">
                {{ receiptFiles.length ? t.payment_change_receipt : t.payment_select_receipt }}
              </button>
              <input ref="receiptInput" type="file" accept="image/*" multiple class="hidden" @change="onReceiptSelect" />
              <p v-if="receiptError" class="text-xs text-red-500 mt-2">{{ receiptError }}</p>
            </div>
            <p v-if="paymentError" class="text-sm text-red-500 mb-2">{{ paymentError }}</p>
            <button @click="confirmOrderPayment" :disabled="!receiptFiles.length || confirmingOrder" class="w-full btn-primary py-3.5 rounded-xl text-base font-semibold disabled:opacity-40 disabled:cursor-not-allowed">
              {{ confirmingOrder ? t.payment_confirming : t.payment_confirm }}
            </button>
          </template>

          <!-- After first receipt: split-payment section -->
          <template v-else>
            <p class="text-xs text-stone-400 mb-3">{{ t.payment_order_code }}: <span class="font-bold text-brand-700 tracking-wider">{{ orderSuccessCode }}</span></p>

            <!-- Uploaded receipts grid -->
            <div class="mb-3 text-left">
              <p class="text-xs font-semibold text-emerald-700 mb-2">{{ t.receipt_uploaded_count(uploadedReceipts.length) }}</p>
              <div class="flex flex-wrap gap-2">
                <div v-for="(path, i) in uploadedReceipts" :key="i" class="relative">
                  <a :href="path" target="_blank" rel="noopener">
                    <img :src="path" class="w-16 h-16 object-cover rounded-xl border-2 border-emerald-300 hover:opacity-80 transition" />
                  </a>
                  <span class="absolute -top-1.5 -right-1.5 bg-emerald-500 text-white text-[10px] font-bold rounded-full w-5 h-5 flex items-center justify-center">{{ i+1 }}</span>
                </div>
              </div>
            </div>

            <!-- Hint -->
            <p class="text-xs text-stone-400 bg-stone-50 rounded-xl p-3 mb-3 text-left">{{ t.split_payment_hint }}</p>

            <!-- Add another receipt -->
            <div class="mb-4 text-left">
              <div v-if="extraReceiptPreview" class="mb-2"><img :src="extraReceiptPreview" class="w-full max-h-36 object-contain rounded-xl border border-stone-200" /></div>
              <button type="button" :disabled="receiptCooldown > 0 || addingReceipt" @click="$refs.extraReceiptInput.click()" class="w-full border-2 border-dashed border-stone-300 rounded-xl py-3 text-sm text-stone-500 hover:border-brand-400 transition-colors disabled:opacity-40 disabled:cursor-not-allowed">
                {{ extraReceiptFile ? t.payment_change_receipt : t.add_another_receipt }}
              </button>
              <input ref="extraReceiptInput" type="file" accept="image/*" class="hidden" @change="onExtraReceiptSelect" />
              <p v-if="receiptCooldown > 0" class="text-xs text-amber-500 mt-1.5 text-center">{{ t.receipt_cooldown(receiptCooldown) }}</p>
              <p v-if="extraReceiptError" class="text-xs text-red-500 mt-1.5">{{ extraReceiptError }}</p>
            </div>
            <p v-if="addReceiptError" class="text-sm text-red-500 mb-2">{{ addReceiptError }}</p>

            <!-- Send extra receipt button -->
            <button v-if="extraReceiptFile" @click="addAnotherReceipt" :disabled="addingReceipt || receiptCooldown > 0" class="w-full bg-amber-500 text-white py-3 rounded-xl text-sm font-semibold mb-3 disabled:opacity-40 disabled:cursor-not-allowed hover:bg-amber-600 transition">
              {{ addingReceipt ? t.adding_receipt : t.payment_confirm }}
            </button>

            <!-- Finish payment -->
            <button @click="paymentStage = 'done'" class="w-full btn-primary py-3.5 rounded-xl text-base font-semibold">
              {{ t.finish_payment }}
            </button>
          </template>
        </template>

        <!-- Stage 2: thank you -->
        <template v-else>
          <div class="w-20 h-20 bg-brand-50 rounded-2xl flex items-center justify-center mx-auto mb-5">
            <svg class="w-10 h-10 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-stone-900 mb-2">{{ t.order_success_title }}</h3>
          <p class="text-stone-500 mb-5 leading-relaxed text-sm">{{ t.order_success_message }}</p>
          <div class="bg-gradient-to-br from-brand-50 to-blue-50 border border-brand-100 rounded-2xl p-5 mb-6">
            <p class="text-xs text-brand-500 font-semibold uppercase tracking-wider mb-2">{{ t.order_success_code_label }}</p>
            <p class="text-5xl font-bold text-brand-700 tracking-[0.15em] mb-2">{{ orderSuccessCode }}</p>
            <p class="text-xs text-stone-400">{{ t.order_success_code_note }}</p>
          </div>
          <button @click="closeOrderSuccess" class="w-full btn-primary py-3.5 rounded-xl text-base font-semibold">{{ t.order_success_button }}</button>
        </template>
      </div>
    </div>

    <!-- ===== EDIT DELIVERY LOCATION MODAL ===== -->
    <div v-if="showEditLocationModal" class="fixed inset-0 z-[80] flex items-center justify-center">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="closeEditLocation"></div>
      <div class="relative bg-white rounded-3xl p-5 max-w-sm w-full mx-4 shadow-2xl max-h-[92vh] overflow-y-auto">
        <h3 class="text-base font-bold text-stone-900 mb-1">{{ t.edit_location_title }}</h3>
        <p class="text-xs text-stone-400 mb-3">{{ t.edit_location_hint }}</p>
        <div id="edit-location-map" class="w-full rounded-2xl overflow-hidden border border-stone-200 mb-3" style="height: 280px;"></div>
        <p v-if="editLocError" class="text-xs text-red-500 mb-3">{{ editLocError }}</p>
        <div class="flex gap-2">
          <button @click="closeEditLocation" class="flex-1 border border-stone-200 text-stone-600 py-2.5 rounded-xl text-sm font-medium hover:bg-stone-50 transition">{{ t.checkout_cancel }}</button>
          <button @click="saveEditLocation" :disabled="editLocSaving || !editLocLat" class="flex-1 btn-primary py-2.5 rounded-xl text-sm font-semibold disabled:opacity-40 disabled:cursor-not-allowed">
            {{ editLocSaving ? t.save_location_saving : t.save_location_btn }}
          </button>
        </div>
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

let markNotesSeenTimer = null
function switchTab(tab) {
  activeTab.value = tab
  if (tab === 'orders') {
    loadOrders()
    ordersTabQrFailed.value = false
    clearTimeout(markNotesSeenTimer)
    markNotesSeenTimer = setTimeout(markNotesSeen, 1500)
  }
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
    awaiting_payment: t.value.status_awaiting_payment,
    pending: t.value.status_pending,
    in_transit: t.value.status_in_transit,
    delivered: t.value.status_delivered,
    cancelled: t.value.status_cancelled,
  }
  return map[status] || status
}

function statusClass(status) {
  const map = {
    awaiting_payment: 'bg-orange-100 text-orange-700',
    pending: 'bg-yellow-100 text-yellow-700',
    in_transit: 'bg-orange-100 text-orange-700',
    delivered: 'bg-green-100 text-green-700',
    cancelled: 'bg-red-100 text-red-600',
  }
  return map[status] || 'bg-stone-100 text-stone-600'
}

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price))
}

function itemTotal(item) {
  return item.price_per_pack * item.quantity
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
const successQrFailed = ref(false)

// Order Map
const showOrderMap = ref(false)
const selectedOrderForMap = ref(null)
const orderGoogleMapUrl = ref('')

// BTS Map
const showBtsMap = ref(false)
const btsMapOrder = ref(null)
const btsGoogleMapUrl = ref('')

async function confirmHide(order) {
  if (!confirm('Удалить этот заказ из истории?')) return
  try {
    await api.delete(`/orders/${order.id}/hide`)
    orders.value = orders.value.filter(o => o.id !== order.id)
  } catch (err) {
    alert(err.response?.data?.error || 'Ошибка')
  }
}

function openBtsMap(order) {
  btsMapOrder.value = order
  const btsLat = order.bts_branch_lat
  const btsLng = order.bts_branch_lng
  const customerLat = order.latitude
  const customerLng = order.longitude

  if (btsLat && btsLng && customerLat && customerLng) {
    // Route from customer delivery location → BTS pickup point
    btsGoogleMapUrl.value = `https://maps.google.com/maps?saddr=${customerLat},${customerLng}&daddr=${btsLat},${btsLng}&dirflg=d&output=embed`
  } else if (btsLat && btsLng) {
    btsGoogleMapUrl.value = `https://maps.google.com/maps?q=${btsLat},${btsLng}&z=16&output=embed`
  } else {
    btsGoogleMapUrl.value = `https://maps.google.com/maps?q=${encodeURIComponent(order.bts_pickup_point || 'Uzbekistan')}&output=embed`
  }
  showBtsMap.value = true
}

function closeBtsMap() {
  showBtsMap.value = false
  btsMapOrder.value = null
  btsGoogleMapUrl.value = ''
}


// Receipt upload from orders tab
const ordersTabReceiptInputs = ref({})
const ordersTabUploading = ref(null)
const ordersTabQrFailed = ref(false)

async function uploadReceiptFromOrdersTab(orderId, e) {
  const f = e.target.files[0]
  if (!f) return
  ordersTabUploading.value = orderId
  try {
    const fd = new FormData()
    fd.append('receipt', f)
    await api.post(`/orders/${orderId}/receipt`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    await loadOrders()
  } catch (err) {
    alert(err.response?.data?.error || 'Ошибка при отправке чека')
  } finally {
    ordersTabUploading.value = null
    e.target.value = ''
  }
}

function openOrderMap(order) {
  selectedOrderForMap.value = order
  const lat = order.latitude
  const lng = order.longitude
  if (lat && lng) {
    orderGoogleMapUrl.value = `https://maps.google.com/maps?q=${lat},${lng}&z=15&output=embed`
  } else if (order.delivery_address) {
    orderGoogleMapUrl.value = `https://maps.google.com/maps?q=${encodeURIComponent(order.delivery_address)}&output=embed`
  } else {
    orderGoogleMapUrl.value = ''
  }
  showOrderMap.value = true
}

function closeOrderMap() {
  showOrderMap.value = false
  selectedOrderForMap.value = null
  orderGoogleMapUrl.value = ''
}

// Online payment (QR + receipt) flow
const paymentStage = ref('pay')      // 'pay' -> 'done'
const currentOrderId = ref(null)
const receiptFiles = ref([])
const receiptPreviews = ref([])
const receiptError = ref('')
const confirmingOrder = ref(false)
const paymentError = ref('')

// Split-payment: additional receipts after first
const uploadedReceipts = ref([])        // paths returned by server (all receipts)
const extraReceiptFile = ref(null)
const extraReceiptPreview = ref(null)
const extraReceiptError = ref('')
const addingReceipt = ref(false)
const addReceiptError = ref('')
const receiptCooldown = ref(0)
let receiptCooldownTimer = null

function startReceiptCooldown() {
  receiptCooldown.value = 10
  clearInterval(receiptCooldownTimer)
  receiptCooldownTimer = setInterval(() => {
    receiptCooldown.value--
    if (receiptCooldown.value <= 0) clearInterval(receiptCooldownTimer)
  }, 1000)
}

function onExtraReceiptSelect(e) {
  const f = e.target.files[0]
  if (!f) return
  const validTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  if (!validTypes.includes(f.type)) { extraReceiptError.value = t.value.receipt_type_error || 'Только изображения'; e.target.value = ''; return }
  if (f.size > 10 * 1024 * 1024) { extraReceiptError.value = t.value.receipt_size_error || 'Максимум 10 МБ'; e.target.value = ''; return }
  extraReceiptError.value = ''
  extraReceiptFile.value = f
  extraReceiptPreview.value = URL.createObjectURL(f)
  e.target.value = ''
}

async function addAnotherReceipt() {
  if (!extraReceiptFile.value || !currentOrderId.value) return
  addReceiptError.value = ''
  addingReceipt.value = true
  try {
    const fd = new FormData()
    fd.append('receipt', extraReceiptFile.value)
    const res = await api.post(`/orders/${currentOrderId.value}/receipts`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    uploadedReceipts.value = res.data.receipt_paths || uploadedReceipts.value
    extraReceiptFile.value = null
    extraReceiptPreview.value = null
    startReceiptCooldown()
  } catch (err) {
    const data = err.response?.data
    addReceiptError.value = data?.error || 'Ошибка при отправке чека'
    if (data?.retry_after) {
      receiptCooldown.value = data.retry_after
      startReceiptCooldown()
    }
  } finally {
    addingReceipt.value = false
  }
}

// ===== PER-ORDER SPLIT PAYMENT IN ORDERS TAB =====
const ordersTabExtraInputs = ref({})
const ordersTabExtraFiles = ref({})
const ordersTabExtraPreviews = ref({})
const ordersTabExtraErrors = ref({})
const ordersTabAddingId = ref(null)
const ordersTabCooldowns = ref({})
const ordersTabCooldownTimers = {}

function startOrdersTabCooldown(orderId, seconds = 10) {
  ordersTabCooldowns.value[orderId] = seconds
  clearInterval(ordersTabCooldownTimers[orderId])
  ordersTabCooldownTimers[orderId] = setInterval(() => {
    ordersTabCooldowns.value[orderId]--
    if (ordersTabCooldowns.value[orderId] <= 0) clearInterval(ordersTabCooldownTimers[orderId])
  }, 1000)
}

function onOrdersTabExtraSelect(orderId, e) {
  const f = e.target.files[0]
  if (!f) return
  if (f.size > 10 * 1024 * 1024) { ordersTabExtraErrors.value[orderId] = 'Максимум 10 МБ'; e.target.value = ''; return }
  ordersTabExtraErrors.value[orderId] = ''
  ordersTabExtraFiles.value[orderId] = f
  ordersTabExtraPreviews.value[orderId] = URL.createObjectURL(f)
  e.target.value = ''
}

async function sendOrdersTabExtra(orderId) {
  const f = ordersTabExtraFiles.value[orderId]
  if (!f) return
  ordersTabAddingId.value = orderId
  ordersTabExtraErrors.value[orderId] = ''
  try {
    const fd = new FormData()
    fd.append('receipt', f)
    await api.post(`/orders/${orderId}/receipts`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    ordersTabExtraFiles.value[orderId] = null
    ordersTabExtraPreviews.value[orderId] = null
    startOrdersTabCooldown(orderId)
    await loadOrders()
  } catch (err) {
    const data = err.response?.data
    ordersTabExtraErrors.value[orderId] = data?.error || 'Ошибка при отправке чека'
    if (data?.retry_after) startOrdersTabCooldown(orderId, data.retry_after)
  } finally {
    ordersTabAddingId.value = null
  }
}

function parseOrderReceiptPaths(order) {
  if (order.receipt_paths) {
    try { const a = JSON.parse(order.receipt_paths); if (Array.isArray(a) && a.length > 0) return a } catch {}
  }
  if (order.receipt_path) return [order.receipt_path]
  return []
}

// ===== WORKER NOTES NOTIFICATION =====
const seenNoteKeys = ref(new Set(JSON.parse(localStorage.getItem('_seen_notes') || '[]')))

function noteKey(order) { return `${order.id}:${(order.worker_notes || '').length}` }
function isNoteUnread(order) { return order.worker_notes && !seenNoteKeys.value.has(noteKey(order)) }
const hasUnreadNotes = computed(() => orders.value.some(o => isNoteUnread(o)))

function markNotesSeen() {
  orders.value.forEach(o => { if (o.worker_notes) seenNoteKeys.value.add(noteKey(o)) })
  try { localStorage.setItem('_seen_notes', JSON.stringify([...seenNoteKeys.value])) } catch {}
}

// ===== EDIT DELIVERY LOCATION =====
const showEditLocationModal = ref(false)
const editLocationOrderId = ref(null)
const editLocLat = ref(0)
const editLocLng = ref(0)
const editLocSaving = ref(false)
const editLocError = ref('')
let editLocMap = null
let editLocMarker = null

async function openEditLocation(order) {
  editLocationOrderId.value = order.id
  editLocLat.value = order.latitude || 0
  editLocLng.value = order.longitude || 0
  editLocError.value = ''
  showEditLocationModal.value = true
  await nextTick()
  if (editLocMap) { editLocMap.remove(); editLocMap = null; editLocMarker = null }
  const L = await import('leaflet')
  await import('leaflet/dist/leaflet.css')
  delete L.Icon.Default.prototype._getIconUrl
  L.Icon.Default.mergeOptions({
    iconRetinaUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon-2x.png',
    iconUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png',
    shadowUrl: 'https://unpkg.com/leaflet@1.9.4/dist/images/marker-shadow.png',
  })
  const mapEl = document.getElementById('edit-location-map')
  if (!mapEl) return
  const lat = editLocLat.value || 40.9983
  const lng = editLocLng.value || 71.6726
  editLocMap = L.map('edit-location-map').setView([lat, lng], editLocLat.value ? 15 : 13)
  L.tileLayer('https://mt{s}.google.com/vt/lyrs=y&x={x}&y={y}&z={z}', {
    subdomains: ['0', '1', '2', '3'], maxZoom: 21, attribution: '© Google Maps'
  }).addTo(editLocMap)
  if (editLocLat.value && editLocLng.value) {
    editLocMarker = L.marker([editLocLat.value, editLocLng.value]).addTo(editLocMap)
  }
  editLocMap.on('click', (e) => {
    editLocLat.value = e.latlng.lat
    editLocLng.value = e.latlng.lng
    if (editLocMarker) editLocMarker.setLatLng([e.latlng.lat, e.latlng.lng])
    else editLocMarker = L.marker([e.latlng.lat, e.latlng.lng]).addTo(editLocMap)
  })
}

function closeEditLocation() {
  showEditLocationModal.value = false
  editLocationOrderId.value = null
  if (editLocMap) { editLocMap.remove(); editLocMap = null; editLocMarker = null }
}

async function saveEditLocation() {
  if (!editLocLat.value || !editLocLng.value) { editLocError.value = t.value.checkout_no_location || 'Выберите точку на карте'; return }
  editLocSaving.value = true
  editLocError.value = ''
  try {
    await api.put(`/orders/${editLocationOrderId.value}/location`, { latitude: editLocLat.value, longitude: editLocLng.value })
    closeEditLocation()
    await loadOrders()
  } catch (err) {
    editLocError.value = err.response?.data?.error || 'Ошибка при сохранении'
  } finally {
    editLocSaving.value = false
  }
}

// Validate receipt file
function onReceiptSelect(e) {
  const files = [...e.target.files]
  e.target.value = ''
  if (!files.length) return
  const validTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  const valid = files.filter(f => validTypes.includes(f.type) && f.size <= 10 * 1024 * 1024)
  if (valid.length !== files.length) {
    receiptError.value = 'Некоторые файлы отклонены (только изображения, макс. 10 МБ)'
  } else {
    receiptError.value = ''
  }
  if (!valid.length) return
  receiptFiles.value = [...receiptFiles.value, ...valid]
  receiptPreviews.value = receiptFiles.value.map(f => URL.createObjectURL(f))
}


async function confirmOrderPayment() {
  if (!receiptFiles.value.length || !currentOrderId.value) return
  paymentError.value = ''
  confirmingOrder.value = true  // set before try so finally always resets it
  try {
    const fd = new FormData()
    fd.append('receipt', receiptFiles.value[0])
    for (let i = 1; i < receiptFiles.value.length; i++) {
      fd.append('receipts', receiptFiles.value[i])
    }
    const res = await api.post(`/orders/${currentOrderId.value}/receipt`, fd, { headers: { 'Content-Type': 'multipart/form-data' } })
    uploadedReceipts.value = res.data.receipt_paths || [res.data.receipt_path]
    startReceiptCooldown()
  } catch (e) {
    paymentError.value = e.response?.data?.error || 'Ошибка при отправке чека'
  } finally {
    confirmingOrder.value = false
  }
}

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
  const opts = [t.value.referral_self, t.value.referral_ad]
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
  referralInput.value = isDoctor ? (t.value.referral_doctor_prefix + val) : val
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
  L.tileLayer('https://mt{s}.google.com/vt/lyrs=y&x={x}&y={y}&z={z}', {
    subdomains: ['0', '1', '2', '3'],
    maxZoom: 21,
    attribution: '© Google Maps'
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

  })
}

function destroyMap() {
  if (leafletMap) {
    leafletMap.remove()
    leafletMap = null
    leafletMarker = null
  }
}

function closeCheckout() {
  destroyMap()
  showCheckout.value = false
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
  currentOrderId.value = null
  receiptFiles.value = []
  receiptPreviews.value = []
  receiptError.value = ''
  paymentStage.value = 'pay'
  uploadedReceipts.value = []
  extraReceiptFile.value = null
  extraReceiptPreview.value = null
  extraReceiptError.value = ''
  addReceiptError.value = ''
  receiptCooldown.value = 0
  clearInterval(receiptCooldownTimer)
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
    currentOrderId.value = res.data?.id || null
    orderSuccessCode.value = res.data?.order_code || ''
    receiptFiles.value = []
    receiptPreviews.value = []
    receiptError.value = ''
    paymentError.value = ''
    paymentStage.value = 'pay'
    successQrFailed.value = false
    showOrderSuccess.value = true
  } catch (e) {
    checkoutError.value = e.response?.data?.error || 'Ошибка при оформлении заказа'
  } finally {
    checkoutLoading.value = false
  }
}

onUnmounted(() => {
  destroyMap()
  if (editLocMap) { editLocMap.remove(); editLocMap = null }
  clearTimeout(markNotesSeenTimer)
  clearInterval(receiptCooldownTimer)
  Object.values(ordersTabCooldownTimers).forEach(clearInterval)
})

function showOrders() {
  activeTab.value = 'orders'
  loadOrders()
}

defineExpose({ showOrders })
</script>
