<template>
  <div class="bg-white">
    <Navbar />
    <CartDrawer ref="cartDrawerRef" />
    <OrdersDrawer />

    <!-- ======= HERO SECTION ======= -->
    <section class="relative bg-white pt-[72px]">
      <div class="absolute inset-0 bg-gradient-to-br from-slate-50 via-white to-blue-50/30 pointer-events-none"></div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-12 pb-24 sm:pb-32">
        <div class="grid md:grid-cols-2 gap-8 lg:gap-16 items-center">

          <!-- LEFT: Text Content -->
          <div class="order-2 md:order-1">
            <h1 class="text-[36px] sm:text-[46px] md:text-[52px] font-bold text-slate-900 leading-tight mb-5">
              {{ t.hero_title_main }}
            </h1>
            <p class="text-slate-500 text-base sm:text-lg mb-8 leading-relaxed">
              {{ t.hero_subtitle_main }}
            </p>

            <!-- Feature rows -->
            <div class="space-y-4 mb-10">
              <div class="flex items-center gap-4">
                <div class="w-11 h-11 rounded-xl bg-blue-50 flex items-center justify-center shrink-0">
                  <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 2C12 2 5 8 5 13a7 7 0 0014 0c0-5-7-11-7-11z"/>
                  </svg>
                </div>
                <div>
                  <div class="font-semibold text-slate-900 text-sm">{{ t.hero_feat1_title }}</div>
                  <div class="text-xs text-slate-400 mt-0.5">{{ t.hero_feat1_sub }}</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-11 h-11 rounded-xl bg-blue-50 flex items-center justify-center shrink-0">
                  <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                  </svg>
                </div>
                <div>
                  <div class="font-semibold text-slate-900 text-sm">{{ t.hero_feat2_title }}</div>
                  <div class="text-xs text-slate-400 mt-0.5">{{ t.hero_feat2_sub }}</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-11 h-11 rounded-xl bg-blue-50 flex items-center justify-center shrink-0">
                  <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>
                  </svg>
                </div>
                <div>
                  <div class="font-semibold text-slate-900 text-sm">{{ t.hero_feat3_title }}</div>
                  <div class="text-xs text-slate-400 mt-0.5">{{ t.hero_feat3_sub }}</div>
                </div>
              </div>
            </div>

            <a href="#products"
               class="inline-flex items-center gap-2.5 bg-brand-700 text-white px-7 py-3.5 rounded-xl font-semibold hover:bg-brand-800 hover:shadow-lg hover:shadow-brand-700/25 hover:-translate-y-0.5 transition-all duration-300 text-sm">
              {{ t.hero_show_products }}
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
              </svg>
            </a>
          </div>

          <!-- RIGHT: 3 product images or fallback -->
          <div class="order-1 md:order-2 flex justify-center md:justify-end">
            <div class="relative w-full max-w-[480px]">

              <!-- Products arrangement (when products with images are loaded) -->
              <div v-if="heroProducts.length > 0"
                   class="relative h-[300px] sm:h-[400px] md:h-[460px] rounded-3xl overflow-hidden"
                   style="background: linear-gradient(145deg, #f0f4ff 0%, #e8f0fe 50%, #f5f7ff 100%);">
                <!-- Decorative circles -->
                <div class="absolute top-8 right-8 w-32 h-32 rounded-full bg-white/40"></div>
                <div class="absolute bottom-4 left-4 w-20 h-20 rounded-full bg-white/30"></div>

                <!-- Product images row -->
                <div class="absolute inset-0 flex items-end justify-center pb-4 sm:pb-8 gap-2 sm:gap-4 px-4">
                  <!-- Left product (shorter, offset down) -->
                  <div v-if="heroProducts.length >= 2" class="relative flex-shrink-0 self-end pb-2 sm:pb-4">
                    <img :src="heroProducts[0].image_path" :alt="heroProducts[0].name"
                         class="h-[140px] sm:h-[185px] md:h-[210px] w-auto object-contain drop-shadow-xl" />
                  </div>
                  <!-- Center product (tallest) -->
                  <div class="relative flex-shrink-0 self-end">
                    <img :src="heroProducts.length >= 2 ? heroProducts[1].image_path : heroProducts[0].image_path"
                         :alt="heroProducts.length >= 2 ? heroProducts[1].name : heroProducts[0].name"
                         class="h-[180px] sm:h-[235px] md:h-[270px] w-auto object-contain drop-shadow-2xl" />
                  </div>
                  <!-- Right product (shorter, offset down) -->
                  <div v-if="heroProducts.length >= 3" class="relative flex-shrink-0 self-end pb-4 sm:pb-6">
                    <img :src="heroProducts[2].image_path" :alt="heroProducts[2].name"
                         class="h-[120px] sm:h-[160px] md:h-[185px] w-auto object-contain drop-shadow-xl" />
                  </div>
                </div>

                <!-- Floating badge top-left -->
                <div class="absolute top-4 sm:top-6 left-4 sm:left-6 bg-white rounded-2xl px-3 sm:px-4 py-2 sm:py-3 shadow-xl shadow-slate-200/60">
                  <div class="text-xl sm:text-2xl font-bold text-brand-700">40+</div>
                  <div class="text-[10px] sm:text-xs text-slate-400 font-medium">{{ t.hero_years }}</div>
                </div>
                <!-- Floating badge bottom-right -->
                <div class="absolute bottom-4 sm:bottom-6 right-4 sm:right-6 bg-white rounded-2xl px-3 sm:px-4 py-2 sm:py-3 shadow-xl shadow-slate-200/60">
                  <div class="text-xl sm:text-2xl font-bold text-brand-700">500K+</div>
                  <div class="text-[10px] sm:text-xs text-slate-400 font-medium">{{ t.stats_clients_label }}</div>
                </div>
              </div>

              <!-- Fallback image (when no products have images) -->
              <div v-else class="relative">
                <img src="/images/patients/3%20landing.jpg" alt="Doctor Jalilov"
                     class="w-full h-[300px] sm:h-[400px] md:h-[460px] object-cover object-top rounded-3xl shadow-2xl shadow-slate-200" />
                <div class="absolute top-4 sm:top-6 left-4 sm:left-6 bg-white rounded-2xl px-3 sm:px-4 py-2 sm:py-3 shadow-xl">
                  <div class="text-xl sm:text-2xl font-bold text-brand-700">40+</div>
                  <div class="text-[10px] sm:text-xs text-slate-400 font-medium">{{ t.hero_years }}</div>
                </div>
                <div class="absolute bottom-4 sm:bottom-6 right-4 sm:right-6 bg-white rounded-2xl px-3 sm:px-4 py-2 sm:py-3 shadow-xl">
                  <div class="text-xl sm:text-2xl font-bold text-brand-700">500K+</div>
                  <div class="text-[10px] sm:text-xs text-slate-400 font-medium">{{ t.stats_clients_label }}</div>
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= STATS BAR (outside hero, overlapping with -mt) ======= -->
    <div class="relative z-10 -mt-10 sm:-mt-12 px-4 sm:px-6 lg:px-8">
      <div class="max-w-5xl mx-auto">
        <div class="bg-white rounded-2xl shadow-xl shadow-slate-200/60 border border-slate-100 px-5 sm:px-8 py-5 sm:py-6 grid grid-cols-2 md:grid-cols-4 gap-4 sm:gap-6">
          <div class="flex flex-col items-start border-r border-slate-100 pr-4 sm:pr-6 last:border-r-0">
            <span class="text-2xl sm:text-3xl font-bold text-brand-700">40<span class="text-brand-400">+</span></span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.hero_years }}</span>
          </div>
          <div class="flex flex-col items-start md:border-r md:border-slate-100 md:pr-6">
            <span class="text-2xl sm:text-3xl font-bold text-brand-700">500 000<span class="text-brand-400">+</span></span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_clients_label }}</span>
          </div>
          <div class="flex flex-col items-start border-r border-slate-100 pr-4 sm:pr-6 last:border-r-0">
            <span class="text-2xl sm:text-3xl font-bold text-brand-700">20<span class="text-brand-400">+</span></span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_products_types }}</span>
          </div>
          <div class="flex flex-col items-start">
            <span class="text-base sm:text-lg font-bold text-brand-700">{{ t.stats_quality_title }}</span>
            <span class="text-xs sm:text-sm text-slate-400 mt-1">{{ t.stats_quality_sub }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Spacer between stats and products -->
    <div class="h-12 sm:h-16 bg-white"></div>

    <!-- ======= PRODUCTS SECTION ======= -->
    <section id="products" class="pb-16 sm:pb-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <!-- Section header -->
        <div class="text-center mb-10">
          <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-bold tracking-widest uppercase mb-3">
            <span class="w-8 h-px bg-brand-300"></span>
            MAHSULOTLAR
            <span class="w-8 h-px bg-brand-300"></span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-bold text-slate-900">{{ t.products_section_title }}</h2>
        </div>

        <!-- Category filter -->
        <div class="flex flex-wrap gap-2 mb-8 sm:mb-10">
          <button
            v-for="cat in categories"
            :key="cat.key"
            @click="selectedCategory = cat.key"
            class="px-4 sm:px-5 py-2 rounded-full text-sm font-medium transition-all duration-200 whitespace-nowrap"
            :class="selectedCategory === cat.key
              ? 'bg-brand-700 text-white shadow-sm shadow-brand-700/20'
              : 'border border-slate-200 text-slate-600 hover:border-brand-300 hover:text-brand-700'"
          >{{ cat.label }}</button>
        </div>

        <!-- Loading -->
        <div v-if="loading" class="flex justify-center py-20">
          <div class="w-10 h-10 border-[3px] border-brand-100 border-t-brand-600 rounded-full animate-spin"></div>
        </div>

        <!-- Empty -->
        <div v-else-if="filteredProducts.length === 0" class="text-center py-16 text-slate-400">
          <svg class="w-12 h-12 mx-auto mb-3 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <p class="font-medium">{{ t.products_empty }}</p>
        </div>

        <!-- Products grid -->
        <div v-else class="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3 sm:gap-5">
          <div
            v-for="product in filteredProducts"
            :key="product.id"
            class="bg-white rounded-2xl border border-slate-100 overflow-hidden hover:shadow-lg hover:shadow-slate-200/60 hover:border-slate-200 transition-all duration-300 group flex flex-col cursor-pointer"
            @click="openProductModal(product)"
          >
            <!-- Image -->
            <div class="aspect-square bg-slate-50 overflow-hidden relative">
              <img v-if="product.image_path" :src="product.image_path" :alt="product.name"
                   class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
              <div v-else class="w-full h-full flex flex-col items-center justify-center">
                <svg class="w-10 h-10 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
              </div>
            </div>
            <!-- Content -->
            <div class="p-3 sm:p-4 flex flex-col flex-1">
              <h3 class="font-semibold text-slate-900 text-sm sm:text-base mb-1 group-hover:text-brand-700 transition-colors line-clamp-1">{{ product.name }}</h3>
              <p v-if="product.description" class="text-slate-400 text-xs mb-3 line-clamp-2 leading-relaxed flex-1">{{ product.description }}</p>
              <div v-else class="flex-1"></div>
              <div class="flex items-center justify-between mt-2">
                <span class="text-brand-700 font-bold text-sm sm:text-base">
                  {{ formatPrice(product.price_per_pack) }} {{ t.currency }}
                </span>
                <button
                  @click.stop="cartStore.addItem(product)"
                  class="w-9 h-9 rounded-full border-2 border-brand-200 text-brand-600 hover:bg-brand-700 hover:border-brand-700 hover:text-white transition-all duration-200 flex items-center justify-center shrink-0"
                  :title="t.add_to_cart"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Show all link -->
        <div v-if="!loading && products.length > 8" class="text-center mt-10">
          <button @click="selectedCategory = ''"
             class="inline-flex items-center gap-2 text-brand-700 font-semibold text-sm hover:text-brand-800 transition-colors group">
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
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6 sm:gap-8">
          <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 2C12 2 5 8 5 13a7 7 0 0014 0c0-5-7-11-7-11z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">{{ t.feat1_title }}</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed">{{ t.feat1_desc }}</div>
            </div>
          </div>
          <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">{{ t.feat2_title }}</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed">{{ t.feat2_desc }}</div>
            </div>
          </div>
          <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">{{ t.feat3_title }}</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed">{{ t.feat3_desc }}</div>
            </div>
          </div>
          <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">{{ t.feat4_title }}</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed">{{ t.feat4_desc }}</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= BOTTOM CTA SECTION ======= -->
    <section class="bg-white py-12 sm:py-16 border-t border-slate-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid md:grid-cols-3 gap-8 sm:gap-10 items-start">

          <!-- Col 1: Quick order -->
          <div>
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-3">{{ t.cta_quick_title }}</h3>
            <p class="text-slate-400 text-sm leading-relaxed mb-6">{{ t.cta_quick_desc }}</p>
            <a href="#products"
               class="inline-flex items-center justify-center px-6 py-3 bg-slate-900 text-white rounded-xl font-semibold text-sm hover:bg-slate-800 transition-colors duration-200">
              {{ t.cta_order_btn }}
            </a>
          </div>

          <!-- Col 2: Featured product -->
          <div class="border border-slate-100 rounded-2xl p-5 sm:p-6 bg-slate-50/50 hover:shadow-md transition-shadow duration-300 text-center">
            <div class="text-xs font-bold text-brand-600 uppercase tracking-widest mb-3">{{ t.cta_special_label }}</div>
            <div class="w-28 h-28 sm:w-36 sm:h-36 mx-auto mb-4">
              <img v-if="featuredProduct && featuredProduct.image_path"
                   :src="featuredProduct.image_path" :alt="featuredProduct.name"
                   class="w-full h-full object-contain drop-shadow-lg" />
              <div v-else class="w-full h-full bg-white rounded-2xl flex items-center justify-center shadow-inner">
                <svg class="w-10 h-10 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
              </div>
            </div>
            <h4 class="font-bold text-slate-900 text-base sm:text-lg mb-1">
              {{ featuredProduct ? featuredProduct.name : 'Immunorayd' }}
            </h4>
            <p class="text-slate-400 text-xs sm:text-sm mb-4 leading-relaxed line-clamp-2">
              {{ featuredProduct ? (featuredProduct.description || '—') : '—' }}
            </p>
            <div class="flex items-center justify-between">
              <span class="text-brand-700 font-bold text-base sm:text-lg">
                {{ featuredProduct ? formatPrice(featuredProduct.price_per_pack) + ' ' + t.currency : '' }}
              </span>
              <button v-if="featuredProduct" @click="openProductModal(featuredProduct)"
                      class="flex items-center gap-1.5 text-brand-700 hover:text-brand-800 font-semibold text-sm transition-colors">
                {{ t.cta_view }}
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
                </svg>
              </button>
            </div>
          </div>

          <!-- Col 3: Contact info -->
          <div id="contacts-quick">
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-5">{{ t.cta_contact_title }}</h3>
            <div class="space-y-3">
              <a :href="`tel:${t.contacts_phone.replace(/\s/g,'')}`"
                 class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0 group-hover:bg-brand-100 transition-colors">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                  </svg>
                </div>
                {{ t.contacts_phone }}
              </a>
              <a :href="`tel:${t.contacts_phone2.replace(/\s/g,'')}`"
                 class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
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

    <!-- ======= PRODUCT DETAIL MODAL ======= -->
    <Teleport to="body">
      <div v-if="selectedProduct" class="fixed inset-0 z-[60] flex items-center justify-center p-4" @click.self="selectedProduct = null">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="selectedProduct = null"></div>
        <div class="relative bg-white rounded-3xl max-w-lg w-full shadow-2xl overflow-hidden max-h-[90vh] flex flex-col">
          <div class="aspect-[16/9] bg-slate-50 flex-shrink-0 relative">
            <img v-if="selectedProduct.image_path" :src="selectedProduct.image_path" :alt="selectedProduct.name" class="w-full h-full object-cover" />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-16 h-16 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
              </svg>
            </div>
            <button @click="selectedProduct = null"
                    class="absolute top-3 right-3 w-9 h-9 bg-black/40 hover:bg-black/60 text-white rounded-full flex items-center justify-center transition-colors">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>
          <div class="p-6 flex flex-col gap-4 overflow-y-auto">
            <div>
              <h3 class="text-2xl font-bold text-slate-900 mb-1">{{ selectedProduct.name }}</h3>
              <div v-if="selectedProduct.category" class="inline-block text-xs font-semibold text-brand-600 bg-brand-50 px-3 py-1 rounded-full mb-2">
                {{ selectedProduct.category }}
              </div>
              <p v-if="selectedProduct.description" class="text-slate-500 text-sm leading-relaxed">{{ selectedProduct.description }}</p>
            </div>
            <div class="flex items-center justify-between py-4 border-t border-slate-100">
              <div>
                <div class="text-xs text-slate-400 mb-0.5">{{ t.product_pack }} ({{ selectedProduct.quantity_per_pack }} {{ t.unit_piece }})</div>
                <div class="text-2xl font-bold text-brand-700">{{ formatPrice(selectedProduct.price_per_pack) }} {{ t.currency }}</div>
              </div>
              <button
                @click="cartStore.addItem(selectedProduct); selectedProduct = null"
                class="flex items-center gap-2 bg-brand-700 text-white px-6 py-3 rounded-xl font-semibold hover:bg-brand-800 transition-colors text-sm">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z"/>
                </svg>
                {{ t.modal_add_cart }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <SimpleFooter />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import CartDrawer from '../components/CartDrawer.vue'
import OrdersDrawer from '../components/OrdersDrawer.vue'
import SimpleFooter from '../components/SimpleFooter.vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useLangStore } from '../stores/lang'

const cartStore = useCartStore()
const authStore = useAuthStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)

const products = ref([])
const loading = ref(true)
const selectedCategory = ref('')
const selectedProduct = ref(null)
const cartDrawerRef = ref(null)

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

const heroProducts = computed(() => {
  return products.value.filter(p => p.image_path).slice(0, 3)
})

const featuredProduct = computed(() => {
  if (!products.value.length) return null
  return products.value.find(p => p.image_path) || products.value[0]
})

function formatPrice(price) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(price || 0))
}

function openProductModal(product) {
  selectedProduct.value = product
}

onMounted(async () => {
  try {
    const res = await axios.get('/api/products')
    products.value = res.data || []
  } catch (e) {
    console.error('Failed to load products:', e)
  } finally {
    loading.value = false
  }
})
</script>
