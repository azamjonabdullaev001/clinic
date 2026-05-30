<template>
  <div class="bg-white">
    <Navbar />
    <CartDrawer ref="cartDrawerRef" />
    <OrdersDrawer />

    <!-- ======= HERO SECTION ======= -->
    <section class="relative pt-[72px] overflow-hidden" style="background:#fff; min-height:640px">

      <!-- Right side: full-bleed hero image with fade edges -->
      <div class="absolute inset-y-0 right-0 hidden md:block" style="left:42%">
        <!-- Image -->
        <img v-if="heroImageSrc"
             :src="heroImageSrc"
             alt="Doctor Jalilov mahsulotlari"
             class="w-full h-full object-cover object-center"
             style="filter: brightness(1.02)" />
        <!-- Fade: left edge -->
        <div class="absolute inset-y-0 left-0 pointer-events-none"
             style="width:42%; background:linear-gradient(to right, #ffffff 0%, rgba(255,255,255,0.85) 35%, rgba(255,255,255,0.4) 65%, transparent 100%)"></div>
        <!-- Fade: top edge -->
        <div class="absolute inset-x-0 top-0 h-28 pointer-events-none"
             style="background:linear-gradient(to bottom, #ffffff 0%, transparent 100%)"></div>
        <!-- Fade: bottom edge -->
        <div class="absolute inset-x-0 bottom-0 h-28 pointer-events-none"
             style="background:linear-gradient(to top, #ffffff 0%, transparent 100%)"></div>
        <!-- Floating badge 1 -->
        <div class="absolute top-8 left-8 bg-white rounded-2xl px-4 py-3 shadow-xl shadow-slate-200/60 float-badge z-10">
          <div class="text-2xl font-bold text-brand-700">40+</div>
          <div class="text-xs text-slate-400 font-medium">{{ t.hero_years }}</div>
        </div>
        <!-- Floating badge 2 -->
        <div class="absolute bottom-10 right-8 bg-white rounded-2xl px-4 py-3 shadow-xl shadow-slate-200/60 float-badge z-10" style="animation-delay:0.5s">
          <div class="text-2xl font-bold text-brand-700">500K+</div>
          <div class="text-xs text-slate-400 font-medium">{{ t.stats_clients_label }}</div>
        </div>
      </div>

      <!-- Left: text content -->
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-14 sm:py-20">
        <div class="md:max-w-[46%]">
          <h1 class="text-[36px] sm:text-[46px] md:text-[52px] font-bold text-slate-900 leading-tight mb-5 reveal-up">
            {{ t.hero_title_main }}
          </h1>
          <p class="text-slate-500 text-base sm:text-lg mb-8 leading-relaxed reveal-up" style="animation-delay:.12s">
            {{ t.hero_subtitle_main }}
          </p>

          <!-- Feature rows -->
          <div class="space-y-4 mb-10">
            <div v-for="(feat, i) in heroFeatures" :key="i"
                 class="flex items-center gap-4 reveal-up"
                 :style="`animation-delay:${0.2 + i * 0.1}s`">
              <div class="w-11 h-11 rounded-xl bg-blue-50 border border-blue-100 flex items-center justify-center shrink-0 shadow-sm">
                <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2" v-html="feat.icon"></svg>
              </div>
              <div>
                <div class="font-semibold text-slate-900 text-sm">{{ feat.title }}</div>
                <div class="text-xs text-slate-400 mt-0.5">{{ feat.sub }}</div>
              </div>
            </div>
          </div>

          <a href="#products"
             class="inline-flex items-center gap-2.5 bg-brand-700 text-white px-7 py-3.5 rounded-xl font-semibold hover:bg-brand-800 hover:shadow-lg hover:shadow-brand-700/30 hover:-translate-y-0.5 active:scale-[0.97] transition-all duration-300 text-sm reveal-up"
             style="animation-delay:.5s">
            {{ t.hero_show_products }}
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
            </svg>
          </a>
        </div>

        <!-- Mobile: product image below text -->
        <div class="md:hidden mt-8 relative">
          <img v-if="heroImageSrc" :src="heroImageSrc"
               alt="Doctor Jalilov" class="w-full h-[260px] sm:h-[320px] object-cover rounded-2xl shadow-xl" />
          <div class="absolute top-4 left-4 bg-white rounded-xl px-3 py-2 shadow-lg">
            <div class="text-lg font-bold text-brand-700">40+</div>
            <div class="text-[10px] text-slate-400">{{ t.hero_years }}</div>
          </div>
          <div class="absolute bottom-4 right-4 bg-white rounded-xl px-3 py-2 shadow-lg">
            <div class="text-lg font-bold text-brand-700">500K+</div>
            <div class="text-[10px] text-slate-400">{{ t.stats_clients_label }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= STATS BAR ======= -->
    <div class="relative z-10 -mt-10 sm:-mt-12 px-4 sm:px-6 lg:px-8">
      <div class="max-w-5xl mx-auto" ref="statsBarRef">
        <div class="bg-white rounded-2xl shadow-xl shadow-slate-200/60 border border-slate-100 px-5 sm:px-8 py-5 sm:py-6 grid grid-cols-2 md:grid-cols-4 gap-4 sm:gap-6">
          <div class="flex flex-col items-start md:border-r md:border-slate-100 md:pr-6">
            <span class="stat-num text-2xl sm:text-3xl font-bold text-brand-700" data-target="40" data-suffix="+">40+</span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.hero_years }}</span>
          </div>
          <div class="flex flex-col items-start md:border-r md:border-slate-100 md:pr-6">
            <span class="stat-num text-2xl sm:text-3xl font-bold text-brand-700" data-target="500000" data-suffix="+">500 000+</span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_clients_label }}</span>
          </div>
          <div class="flex flex-col items-start md:border-r md:border-slate-100 md:pr-6">
            <span class="stat-num text-2xl sm:text-3xl font-bold text-brand-700" data-target="20" data-suffix="+">20+</span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_products_types }}</span>
          </div>
          <div class="flex flex-col items-start">
            <span class="text-base sm:text-lg font-bold text-brand-700">{{ t.stats_quality_title }}</span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_quality_sub }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="h-12 sm:h-16 bg-white"></div>

    <!-- ======= PRODUCTS SECTION ======= -->
    <section id="products" class="pb-16 sm:pb-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-10 reveal-up">
          <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-bold tracking-widest uppercase mb-3">
            <span class="w-8 h-px bg-brand-300"></span>
            MAHSULOTLAR
            <span class="w-8 h-px bg-brand-300"></span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-bold text-slate-900">{{ t.products_section_title }}</h2>
        </div>

        <!-- Category filter -->
        <div class="flex flex-wrap gap-2 mb-8 sm:mb-10 reveal-up" style="animation-delay:.1s">
          <button v-for="cat in categories" :key="cat.key"
                  @click="selectedCategory = cat.key"
                  class="px-4 sm:px-5 py-2 rounded-full text-sm font-medium transition-all duration-200 whitespace-nowrap"
                  :class="selectedCategory === cat.key
                    ? 'bg-brand-700 text-white shadow-sm'
                    : 'border border-slate-200 text-slate-600 hover:border-brand-300 hover:text-brand-700 hover:bg-brand-50'">
            {{ cat.label }}
          </button>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="flex justify-center py-20">
          <div class="w-10 h-10 border-[3px] border-brand-100 border-t-brand-600 rounded-full animate-spin"></div>
        </div>

        <!-- Empty -->
        <div v-else-if="filteredProducts.length === 0" class="text-center py-16 text-slate-400">
          <p class="font-medium">{{ t.products_empty }}</p>
        </div>

        <!-- Products grid -->
        <div v-else class="grid grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3 sm:gap-5 products-grid" ref="productsGridRef">
          <div v-for="(product, index) in filteredProducts"
               :key="product.id"
               class="product-card bg-white rounded-2xl border border-slate-100 overflow-hidden flex flex-col group"
               :style="`--stagger:${index}`">

            <!-- Image (click → lightbox) -->
            <div class="aspect-square bg-slate-50 overflow-hidden relative cursor-zoom-in"
                 @click="openLightbox(product)">
              <img v-if="product.image_path" :src="product.image_path" :alt="product.name"
                   class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110" />
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-10 h-10 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
              </div>
              <!-- Zoom hint overlay -->
              <div class="absolute inset-0 bg-black/0 group-hover:bg-black/[0.04] transition-colors duration-300"></div>
            </div>

            <!-- Content (click → info modal) -->
            <div class="p-3 sm:p-4 flex flex-col flex-1 cursor-pointer" @click="openInfoModal(product)">
              <h3 class="font-semibold text-slate-900 text-sm sm:text-base mb-1 group-hover:text-brand-700 transition-colors line-clamp-1">
                {{ product.name }}
              </h3>
              <p v-if="product.description" class="text-slate-400 text-xs mb-3 line-clamp-2 leading-relaxed flex-1">
                {{ product.description }}
              </p>
              <div v-else class="flex-1 text-xs text-slate-300 italic mb-3">{{ t.product_view_desc }} →</div>
              <div class="flex items-center justify-between mt-auto">
                <span class="text-brand-700 font-bold text-sm sm:text-base">
                  {{ formatPrice(product.price_per_pack) }} {{ t.currency }}
                </span>
                <button @click.stop="handleAddToCart(product)"
                        class="cart-btn relative w-9 h-9 rounded-full border-2 border-brand-200 text-brand-600 flex items-center justify-center shrink-0 transition-all duration-200 hover:bg-brand-700 hover:border-brand-700 hover:text-white active:scale-90"
                        :class="{ 'cart-bounce': addingMap[product.id] }"
                        :title="t.add_to_cart">
                  <span v-if="addingMap[product.id]"
                        class="absolute -top-5 left-1/2 -translate-x-1/2 text-xs font-bold text-emerald-600 pop-up pointer-events-none">+1</span>
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <div v-if="!loading && products.length > 8" class="text-center mt-10">
          <button @click="selectedCategory = ''" class="inline-flex items-center gap-2 text-brand-700 font-semibold text-sm hover:text-brand-800 transition-colors group">
            {{ t.products_show_all }}
            <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
            </svg>
          </button>
        </div>
      </div>
    </section>

    <!-- ======= FEATURES BAR ======= -->
    <section id="about" class="bg-brand-950 py-12 sm:py-14">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6 sm:gap-8 reveal-section" ref="featureBarRef">
          <div v-for="(feat, i) in featureBarItems" :key="i"
               class="flex flex-col gap-3 sm:flex-row sm:items-start sm:gap-4 feature-item"
               :style="`--stagger:${i}`">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8" v-html="feat.icon"></svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">{{ feat.title }}</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed">{{ feat.desc }}</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= BOTTOM CTA ======= -->
    <section class="bg-white py-12 sm:py-16 border-t border-slate-100" id="contacts-quick">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid md:grid-cols-3 gap-8 sm:gap-10 items-start">

          <div class="reveal-up">
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-3">{{ t.cta_quick_title }}</h3>
            <p class="text-slate-400 text-sm leading-relaxed mb-6">{{ t.cta_quick_desc }}</p>
            <a href="#products" class="inline-flex items-center justify-center px-6 py-3 bg-slate-900 text-white rounded-xl font-semibold text-sm hover:bg-slate-800 hover:-translate-y-0.5 transition-all duration-200">
              {{ t.cta_order_btn }}
            </a>
          </div>

          <div class="border border-slate-100 rounded-2xl p-5 sm:p-6 bg-slate-50/50 hover:shadow-lg transition-shadow duration-300 text-center reveal-up" style="animation-delay:.1s">
            <div class="text-xs font-bold text-brand-600 uppercase tracking-widest mb-3">{{ t.cta_special_label }}</div>
            <div class="w-28 h-28 sm:w-36 sm:h-36 mx-auto mb-4">
              <img v-if="featuredProduct && featuredProduct.image_path"
                   :src="featuredProduct.image_path" :alt="featuredProduct.name"
                   class="w-full h-full object-contain drop-shadow-lg hover:scale-105 transition-transform duration-300" />
            </div>
            <h4 class="font-bold text-slate-900 text-base mb-1">{{ featuredProduct?.name }}</h4>
            <p class="text-slate-400 text-xs mb-4 leading-relaxed line-clamp-2">{{ featuredProduct?.description }}</p>
            <div class="flex items-center justify-between">
              <span class="text-brand-700 font-bold text-base">{{ featuredProduct ? formatPrice(featuredProduct.price_per_pack) + ' ' + t.currency : '' }}</span>
              <button v-if="featuredProduct" @click="openInfoModal(featuredProduct)"
                      class="flex items-center gap-1.5 text-brand-700 hover:text-brand-800 font-semibold text-sm transition-colors">
                {{ t.cta_view }}
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
                </svg>
              </button>
            </div>
          </div>

          <div class="reveal-up" style="animation-delay:.2s">
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-5">{{ t.cta_contact_title }}</h3>
            <div class="space-y-3">
              <a :href="`tel:${t.contacts_phone.replace(/\s/g,'')}`" class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0 group-hover:bg-brand-100 transition-colors">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                  </svg>
                </div>
                {{ t.contacts_phone }}
              </a>
              <a :href="`tel:${t.contacts_phone2.replace(/\s/g,'')}`" class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0 group-hover:bg-brand-100 transition-colors">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                  </svg>
                </div>
                {{ t.contacts_phone2 }}
              </a>
              <div class="flex items-center gap-3 text-sm text-slate-600">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"/>
                  </svg>
                </div>
                {{ t.cta_workdays_text }}
              </div>
              <div class="flex items-center gap-3 text-sm text-slate-500">
                <div class="w-9 h-9 rounded-xl bg-slate-50 flex items-center justify-center shrink-0">
                  <svg class="w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5"/>
                  </svg>
                </div>
                {{ t.cta_weekend_text }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= INFO MODAL (photo left + description right) ======= -->
    <Teleport to="body">
      <div v-if="infoProduct" class="fixed inset-0 z-[60] flex items-center justify-center p-3 sm:p-6">
        <div class="absolute inset-0 bg-black/55 backdrop-blur-sm modal-bg" @click="infoProduct = null"></div>
        <div class="relative bg-white rounded-2xl sm:rounded-3xl w-full max-w-2xl shadow-2xl overflow-hidden max-h-[90vh] flex flex-col sm:flex-row modal-panel">
          <!-- Left: photo -->
          <div class="w-full sm:w-[42%] flex-shrink-0 bg-slate-50 flex items-center justify-center p-4 sm:p-6 cursor-zoom-in min-h-[200px] sm:min-h-0"
               @click="openLightbox(infoProduct)">
            <img v-if="infoProduct.image_path" :src="infoProduct.image_path" :alt="infoProduct.name"
                 class="max-h-[200px] sm:max-h-[380px] w-full object-contain drop-shadow-lg hover:scale-105 transition-transform duration-300" />
            <div v-else class="text-slate-200">
              <svg class="w-16 h-16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
              </svg>
            </div>
            <p v-if="infoProduct.image_path" class="absolute bottom-2 left-0 right-0 text-center text-[10px] text-slate-400 hidden sm:block">
              {{ t.product_zoom }} 🔍
            </p>
          </div>

          <!-- Right: description -->
          <div class="flex-1 flex flex-col overflow-y-auto">
            <!-- Close btn -->
            <button @click="infoProduct = null"
                    class="absolute top-3 right-3 w-8 h-8 bg-slate-100 hover:bg-slate-200 rounded-full flex items-center justify-center transition-colors z-10">
              <svg class="w-4 h-4 text-slate-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>

            <div class="p-5 sm:p-7 flex flex-col gap-4 flex-1">
              <div>
                <h2 class="text-xl sm:text-2xl font-bold text-slate-900 mb-2 pr-8">{{ infoProduct.name }}</h2>
                <span v-if="infoProduct.category"
                      class="inline-block text-xs font-semibold text-brand-600 bg-brand-50 px-3 py-1 rounded-full">
                  {{ infoProduct.category }}
                </span>
              </div>

              <p v-if="infoProduct.description"
                 class="text-slate-600 text-sm sm:text-base leading-relaxed flex-1">
                {{ infoProduct.description }}
              </p>
              <p v-else class="text-slate-400 text-sm italic flex-1">—</p>

              <div class="pt-4 border-t border-slate-100 flex items-center justify-between gap-4">
                <div>
                  <div class="text-xs text-slate-400 mb-0.5">{{ t.product_pack }} · {{ infoProduct.quantity_per_pack }} {{ t.unit_piece }}</div>
                  <div class="text-2xl font-bold text-brand-700">{{ formatPrice(infoProduct.price_per_pack) }} {{ t.currency }}</div>
                </div>
                <button @click="handleAddToCart(infoProduct); infoProduct = null"
                        class="flex items-center gap-2 bg-brand-700 text-white px-5 sm:px-6 py-3 rounded-xl font-semibold hover:bg-brand-800 active:scale-95 transition-all text-sm shrink-0">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z"/>
                  </svg>
                  {{ t.modal_add_cart }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- ======= LIGHTBOX (photo only, fullscreen) ======= -->
    <Teleport to="body">
      <div v-if="lightboxSrc" class="fixed inset-0 z-[70] flex items-center justify-center bg-black/85 backdrop-blur-sm"
           @click="lightboxSrc = null">
        <img :src="lightboxSrc" alt="Kattarishtirilgan rasm"
             class="max-h-[92vh] max-w-[92vw] object-contain rounded-xl shadow-2xl lightbox-img"
             @click.stop />
        <button @click="lightboxSrc = null"
                class="absolute top-4 right-4 w-10 h-10 bg-white/10 hover:bg-white/20 text-white rounded-full flex items-center justify-center transition-colors">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>
    </Teleport>

    <SimpleFooter />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, reactive } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import CartDrawer from '../components/CartDrawer.vue'
import OrdersDrawer from '../components/OrdersDrawer.vue'
import SimpleFooter from '../components/SimpleFooter.vue'
import { useCartStore } from '../stores/cart'
import { useLangStore } from '../stores/lang'

const cartStore = useCartStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)

const products = ref([])
const loading = ref(true)
const selectedCategory = ref('')
const infoProduct = ref(null)   // info modal (photo + description)
const lightboxSrc = ref(null)   // lightbox (photo only)
const cartDrawerRef = ref(null)
const statsBarRef = ref(null)
const productsGridRef = ref(null)
const featureBarRef = ref(null)
const addingMap = reactive({})

function handleAddToCart(product) {
  cartStore.addItem(product)
  addingMap[product.id] = true
  setTimeout(() => { delete addingMap[product.id] }, 700)
}

function openInfoModal(product) { infoProduct.value = product }

function openLightbox(product) {
  if (product.image_path) lightboxSrc.value = product.image_path
}

const heroFeatures = computed(() => [
  { title: t.value.hero_feat1_title, sub: t.value.hero_feat1_sub, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M12 2C12 2 5 8 5 13a7 7 0 0014 0c0-5-7-11-7-11z"/>' },
  { title: t.value.hero_feat2_title, sub: t.value.hero_feat2_sub, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>' },
  { title: t.value.hero_feat3_title, sub: t.value.hero_feat3_sub, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>' },
])

const featureBarItems = computed(() => [
  { title: t.value.feat1_title, desc: t.value.feat1_desc, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M12 2C12 2 5 8 5 13a7 7 0 0014 0c0-5-7-11-7-11z"/>' },
  { title: t.value.feat2_title, desc: t.value.feat2_desc, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>' },
  { title: t.value.feat3_title, desc: t.value.feat3_desc, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>' },
  { title: t.value.feat4_title, desc: t.value.feat4_desc, icon: '<path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z"/>' },
])

const categories = computed(() => [
  { key: '', label: t.value.cat_all },
  { key: 'Immunitet', label: t.value.cat_immune },
  { key: 'Jigar uchun', label: t.value.cat_liver },
  { key: 'Asab tizimi', label: t.value.cat_nerve },
  { key: "Bo'g'imlar", label: t.value.cat_bones },
  { key: 'Yurak-qon tomir', label: t.value.cat_heart },
  { key: 'Boshqalar', label: t.value.cat_other },
])

const filteredProducts = computed(() => {
  if (!selectedCategory.value) return products.value
  return products.value.filter(p => p.category === selectedCategory.value)
})

const heroImageSrc = computed(() => '/images/tabletka.png')

const featuredProduct = computed(() => products.value.find(p => p.image_path) || null)

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

// ── Scroll reveal ────────────────────────────────────────────────────────────
let observer = null
function setupScrollReveal() {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) { entry.target.classList.add('visible'); observer.unobserve(entry.target) }
    })
  }, { threshold: 0.08, rootMargin: '0px 0px -40px 0px' })
  document.querySelectorAll('.reveal-up, .reveal-right, .reveal-section').forEach(el => observer.observe(el))
  if (productsGridRef.value) observer.observe(productsGridRef.value)
  if (featureBarRef.value) observer.observe(featureBarRef.value)
}

function animateCountUp(el) {
  const target = parseInt(el.dataset.target)
  const suffix = el.dataset.suffix || ''
  if (!target) return
  let start = 0
  const step = target / (1800 / 16)
  const timer = setInterval(() => {
    start += step
    if (start >= target) { clearInterval(timer); el.textContent = new Intl.NumberFormat('ru-RU').format(target) + suffix }
    else el.textContent = new Intl.NumberFormat('ru-RU').format(Math.floor(start)) + suffix
  }, 16)
}

function setupStats() {
  const statsObs = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) { entry.target.querySelectorAll('.stat-num').forEach(animateCountUp); statsObs.unobserve(entry.target) }
    })
  }, { threshold: 0.4 })
  if (statsBarRef.value) statsObs.observe(statsBarRef.value)
}

onMounted(async () => {
  try {
    const res = await axios.get('/api/products')
    products.value = res.data || []
  } catch (e) { console.error(e) }
  finally { loading.value = false }
  setTimeout(() => { setupScrollReveal(); setupStats() }, 120)
})

onUnmounted(() => { if (observer) observer.disconnect() })
</script>

<style scoped>
/* Reveal animations */
.reveal-up { opacity: 0; transform: translateY(28px); transition: opacity .65s cubic-bezier(.16,1,.3,1), transform .65s cubic-bezier(.16,1,.3,1); }
.reveal-up.visible, .reveal-right.visible, .reveal-section.visible { opacity: 1; transform: none; }
.reveal-right { opacity: 0; transform: translateX(32px); transition: opacity .7s cubic-bezier(.16,1,.3,1), transform .7s cubic-bezier(.16,1,.3,1); }

/* Product cards stagger */
.products-grid.visible .product-card { animation: cardEnter .55s cubic-bezier(.16,1,.3,1) both; animation-delay: calc(var(--stagger,0) * .07s); }
@keyframes cardEnter { from { opacity:0; transform: translateY(22px) scale(.97); } to { opacity:1; transform:none; } }

/* Card hover */
.product-card { transition: transform .25s ease, box-shadow .25s ease, border-color .25s ease; }
.product-card:hover { transform: translateY(-4px); box-shadow: 0 12px 32px rgba(30,64,175,.10); border-color: rgb(191,219,254); }

/* Feature bar stagger */
.reveal-section.visible .feature-item { animation: fadeSlide .6s cubic-bezier(.16,1,.3,1) both; animation-delay: calc(var(--stagger,0) * .1s); }
@keyframes fadeSlide { from { opacity:0; transform: translateY(18px); } to { opacity:1; transform:none; } }

/* Float badges */
.float-badge { animation: floatBadge 4s ease-in-out infinite; }
@keyframes floatBadge { 0%,100% { transform: translateY(0); } 50% { transform: translateY(-6px); } }

/* Cart bounce */
@keyframes cartBounce { 0% { transform: scale(1); } 25% { transform: scale(.75); } 60% { transform: scale(1.3); } 80% { transform: scale(.9); } 100% { transform: scale(1); } }
.cart-btn.cart-bounce { animation: cartBounce .6s cubic-bezier(.36,.07,.19,.97); }

/* +1 popup */
@keyframes popUp { 0% { opacity:0; transform: translateX(-50%) translateY(0); } 30% { opacity:1; transform: translateX(-50%) translateY(-6px); } 100% { opacity:0; transform: translateX(-50%) translateY(-14px); } }
.pop-up { animation: popUp .65s ease forwards; }

/* Modals */
@keyframes fadeIn { from { opacity:0; } to { opacity:1; } }
@keyframes slideUp { from { opacity:0; transform: translateY(20px) scale(.97); } to { opacity:1; transform:none; } }
@keyframes zoomIn { from { opacity:0; transform: scale(.9); } to { opacity:1; transform:scale(1); } }
.modal-bg { animation: fadeIn .2s ease forwards; }
.modal-panel { animation: slideUp .3s cubic-bezier(.16,1,.3,1) forwards; }
.lightbox-img { animation: zoomIn .25s cubic-bezier(.16,1,.3,1) forwards; }
</style>
