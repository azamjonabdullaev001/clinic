<template>
  <div class="bg-white">
    <Navbar />
    <CartDrawer ref="cartDrawerRef" />
    <OrdersDrawer />

    <!-- ======= HERO SECTION ======= -->
    <section class="relative bg-white pt-[72px] overflow-hidden">
      <!-- Subtle background gradient -->
      <div class="absolute inset-0 bg-gradient-to-br from-slate-50 via-white to-blue-50/30 pointer-events-none"></div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-14 pb-28 md:pb-36">
        <div class="grid md:grid-cols-2 gap-8 lg:gap-16 items-center">

          <!-- LEFT: Text Content -->
          <div class="order-2 md:order-1">
            <h1 class="text-[38px] sm:text-[46px] md:text-[52px] font-bold text-slate-900 leading-tight mb-5">
              Sifatli mahsulotlar sog'lom hayot uchun
            </h1>
            <p class="text-slate-500 text-base sm:text-lg mb-8 leading-relaxed">
              Tabiiy va samarali komponentlar asosida yaratilgan mahsulotlar.
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
                  <div class="font-semibold text-slate-900 text-sm">Tabiiy tarkib</div>
                  <div class="text-xs text-slate-400 mt-0.5">To'liq xavfsiz</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-11 h-11 rounded-xl bg-blue-50 flex items-center justify-center shrink-0">
                  <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                  </svg>
                </div>
                <div>
                  <div class="font-semibold text-slate-900 text-sm">Sifat kafolati</div>
                  <div class="text-xs text-slate-400 mt-0.5">Sertifikatlangan</div>
                </div>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-11 h-11 rounded-xl bg-blue-50 flex items-center justify-center shrink-0">
                  <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>
                  </svg>
                </div>
                <div>
                  <div class="font-semibold text-slate-900 text-sm">Tez yetkazib berish</div>
                  <div class="text-xs text-slate-400 mt-0.5">O'z vaqtida</div>
                </div>
              </div>
            </div>

            <a href="#products"
               class="inline-flex items-center gap-2.5 bg-brand-700 text-white px-7 py-3.5 rounded-xl font-semibold hover:bg-brand-800 hover:shadow-lg hover:shadow-brand-700/25 hover:-translate-y-0.5 transition-all duration-300 text-sm">
              Mahsulotlarni ko'rish
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
              </svg>
            </a>
          </div>

          <!-- RIGHT: Product image -->
          <div class="order-1 md:order-2 flex justify-center md:justify-end">
            <div class="relative w-full max-w-[480px]">
              <img
                src="/images/patients/3%20landing.jpg"
                alt="Doctor Jalilov mahsulotlari"
                class="w-full h-[320px] sm:h-[420px] md:h-[480px] object-cover object-center rounded-3xl shadow-2xl shadow-slate-200"
              />
              <!-- Floating badge -->
              <div class="absolute top-6 left-6 bg-white rounded-2xl px-4 py-3 shadow-xl shadow-slate-200/60">
                <div class="text-2xl font-bold text-brand-700">40+</div>
                <div class="text-xs text-slate-400 font-medium">Yillik tajriba</div>
              </div>
              <div class="absolute bottom-6 right-6 bg-white rounded-2xl px-4 py-3 shadow-xl shadow-slate-200/60">
                <div class="text-2xl font-bold text-brand-700">500K+</div>
                <div class="text-xs text-slate-400 font-medium">Mamnun mijozlar</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats bar (overlapping) -->
      <div class="absolute -bottom-[52px] sm:-bottom-[56px] left-0 right-0 px-4 sm:px-6 lg:px-8 z-10">
        <div class="max-w-5xl mx-auto">
          <div class="bg-white rounded-2xl shadow-xl shadow-slate-200/60 border border-slate-100 px-6 sm:px-10 py-5 sm:py-6 grid grid-cols-2 md:grid-cols-4 gap-4 sm:gap-6">
            <div class="flex flex-col items-center text-center md:items-start md:text-left md:border-r md:border-slate-100 md:pr-6">
              <span class="text-2xl sm:text-3xl font-bold text-brand-700">40<span class="text-brand-400">+</span></span>
              <span class="text-xs sm:text-sm text-slate-400 mt-0.5">Yillik tajriba</span>
            </div>
            <div class="flex flex-col items-center text-center md:items-start md:text-left md:border-r md:border-slate-100 md:pr-6">
              <span class="text-2xl sm:text-3xl font-bold text-brand-700">500 000<span class="text-brand-400">+</span></span>
              <span class="text-xs sm:text-sm text-slate-400 mt-0.5">Mamnun mijozlar</span>
            </div>
            <div class="flex flex-col items-center text-center md:items-start md:text-left md:border-r md:border-slate-100 md:pr-6">
              <span class="text-2xl sm:text-3xl font-bold text-brand-700">20<span class="text-brand-400">+</span></span>
              <span class="text-xs sm:text-sm text-slate-400 mt-0.5">Mahsulot turi</span>
            </div>
            <div class="flex flex-col items-center text-center md:items-start md:text-left">
              <span class="text-base sm:text-lg font-bold text-brand-700">Sifat va ishonch</span>
              <span class="text-xs sm:text-sm text-slate-400 mt-0.5">Sog'ligingiz biz uchun muhim</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Spacer for stats bar overlap -->
    <div class="h-[80px] sm:h-[88px] bg-white"></div>

    <!-- ======= PRODUCTS SECTION ======= -->
    <section id="products" class="py-16 sm:py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <!-- Section header -->
        <div class="text-center mb-10">
          <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-bold tracking-widest uppercase mb-3">
            <span class="w-8 h-px bg-brand-300"></span>
            MAHSULOTLAR
            <span class="w-8 h-px bg-brand-300"></span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-bold text-slate-900">Mahsulotlarimiz</h2>
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
          <p class="font-medium">Mahsulotlar topilmadi</p>
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
              <img
                v-if="product.image_path"
                :src="product.image_path"
                :alt="product.name"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div v-else class="w-full h-full flex flex-col items-center justify-center">
                <svg class="w-10 h-10 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
              </div>
            </div>
            <!-- Content -->
            <div class="p-3 sm:p-4 flex flex-col flex-1">
              <h3 class="font-semibold text-slate-900 text-sm sm:text-base mb-1 group-hover:text-brand-700 transition-colors">{{ product.name }}</h3>
              <p v-if="product.description" class="text-slate-400 text-xs mb-3 line-clamp-2 leading-relaxed flex-1">{{ product.description }}</p>
              <div v-else class="flex-1"></div>
              <div class="flex items-center justify-between mt-2">
                <span class="text-brand-700 font-bold text-sm sm:text-base">
                  {{ formatPrice(product.price_per_pack) }} so'm
                </span>
                <button
                  @click.stop="cartStore.addItem(product)"
                  class="w-9 h-9 rounded-full border-2 border-brand-200 text-brand-600 hover:bg-brand-700 hover:border-brand-700 hover:text-white transition-all duration-200 flex items-center justify-center shrink-0"
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
          <a href="#products" @click="selectedCategory = ''"
             class="inline-flex items-center gap-2 text-brand-700 font-semibold text-sm hover:text-brand-800 transition-colors group">
            Barcha mahsulotlar
            <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
            </svg>
          </a>
        </div>
      </div>
    </section>

    <!-- ======= FEATURES BAR ======= -->
    <section id="about" class="bg-brand-950 py-12 sm:py-14">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6 sm:gap-8">
          <!-- Feature 1 -->
          <div class="flex flex-col items-center text-center sm:flex-row sm:items-start sm:text-left sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0 mb-3 sm:mb-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 2C12 2 5 8 5 13a7 7 0 0014 0c0-5-7-11-7-11z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">Tabiiy tarkib</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed hidden sm:block">Mahsulotlarimiz faqat tabiiy komponentlardan iborat.</div>
            </div>
          </div>
          <!-- Feature 2 -->
          <div class="flex flex-col items-center text-center sm:flex-row sm:items-start sm:text-left sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0 mb-3 sm:mb-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">Sifat nazorati</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed hidden sm:block">Har bir mahsulot sertifikatlangan va sinovdan o'tgan.</div>
            </div>
          </div>
          <!-- Feature 3 -->
          <div class="flex flex-col items-center text-center sm:flex-row sm:items-start sm:text-left sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0 mb-3 sm:mb-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">Tez yetkazib berish</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed hidden sm:block">Buyurtmalaringizni qisqa muddatda yetkazib beramiz.</div>
            </div>
          </div>
          <!-- Feature 4 -->
          <div class="flex flex-col items-center text-center sm:flex-row sm:items-start sm:text-left sm:gap-4">
            <div class="w-11 h-11 rounded-xl bg-white/10 flex items-center justify-center shrink-0 mb-3 sm:mb-0">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z"/>
              </svg>
            </div>
            <div>
              <div class="font-semibold text-white text-sm sm:text-base">Ishonch va tajriba</div>
              <div class="text-white/50 text-xs mt-1 leading-relaxed hidden sm:block">40+ yillik tajriba va yuz minglab mamnun mijozlar.</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ======= BOTTOM CTA SECTION ======= -->
    <section class="bg-white py-12 sm:py-16 border-t border-slate-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid md:grid-cols-3 gap-8 sm:gap-10 items-start">

          <!-- Column 1: Quick order -->
          <div class="flex flex-col">
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-3">Tez buyurtma berish</h3>
            <p class="text-slate-400 text-sm leading-relaxed mb-6">
              Kerakli mahsulotni tanlang va buyurtma bering.
            </p>
            <a href="#products"
               class="inline-flex items-center justify-center px-6 py-3 bg-slate-900 text-white rounded-xl font-semibold text-sm hover:bg-slate-800 transition-colors duration-200 w-fit">
              Buyurtma berish
            </a>
          </div>

          <!-- Column 2: Featured product -->
          <div class="flex flex-col items-center text-center border border-slate-100 rounded-2xl p-6 bg-slate-50/50 hover:shadow-md transition-shadow duration-300">
            <div class="text-xs font-bold text-brand-600 uppercase tracking-widest mb-3">Maxsus taklif</div>
            <div class="w-32 h-32 sm:w-40 sm:h-40 mb-4 flex items-center justify-center">
              <img
                v-if="featuredProduct && featuredProduct.image_path"
                :src="featuredProduct.image_path"
                :alt="featuredProduct.name"
                class="w-full h-full object-contain drop-shadow-lg"
              />
              <div v-else class="w-32 h-32 bg-white rounded-2xl flex items-center justify-center shadow-inner">
                <svg class="w-12 h-12 text-slate-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                </svg>
              </div>
            </div>
            <h4 class="font-bold text-slate-900 text-base sm:text-lg mb-1">
              {{ featuredProduct ? featuredProduct.name : 'Immunorayd' }}
            </h4>
            <p class="text-slate-400 text-xs sm:text-sm mb-4 leading-relaxed">
              {{ featuredProduct ? (featuredProduct.description || 'Immunitetni mustahkamlovchi kompleks formula') : 'Immunitetni mustahkamlovchi kompleks formula' }}
            </p>
            <div class="flex items-center justify-between w-full">
              <span class="text-brand-700 font-bold text-base sm:text-lg">
                {{ featuredProduct ? formatPrice(featuredProduct.price_per_pack) + " so'm" : '300 000 so\'m' }}
              </span>
              <button
                v-if="featuredProduct"
                @click="cartStore.addItem(featuredProduct)"
                class="flex items-center gap-1.5 text-brand-700 hover:text-brand-800 font-semibold text-sm transition-colors"
              >
                Ko'rish
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
                </svg>
              </button>
            </div>
          </div>

          <!-- Column 3: Contact info -->
          <div class="flex flex-col" id="contacts-quick">
            <h3 class="text-xl sm:text-2xl font-bold text-slate-900 mb-5">Biz bilan bog'laning</h3>
            <div class="space-y-3">
              <a href="tel:+998993251740" class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0 group-hover:bg-brand-100 transition-colors">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                  </svg>
                </div>
                +998 99 325 17 40
              </a>
              <a href="tel:+998937805097" class="flex items-center gap-3 text-sm text-slate-600 hover:text-brand-700 transition-colors group">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0 group-hover:bg-brand-100 transition-colors">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                  </svg>
                </div>
                +998 93 780 50 97
              </a>
              <div class="flex items-center gap-3 text-sm text-slate-600">
                <div class="w-9 h-9 rounded-xl bg-brand-50 flex items-center justify-center shrink-0">
                  <svg class="w-4 h-4 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"/>
                  </svg>
                </div>
                Du-Ju 08:00 – 17:00
              </div>
              <div class="flex items-center gap-3 text-sm text-slate-600">
                <div class="w-9 h-9 rounded-xl bg-slate-50 flex items-center justify-center shrink-0">
                  <svg class="w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5"/>
                  </svg>
                </div>
                Shanba – Yakshanba dam olish kuni
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
          <!-- Image -->
          <div class="aspect-[16/9] bg-slate-50 flex-shrink-0 relative">
            <img v-if="selectedProduct.image_path" :src="selectedProduct.image_path" :alt="selectedProduct.name"
                 class="w-full h-full object-cover" />
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
          <!-- Content -->
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
                <div class="text-xs text-slate-400 mb-0.5">Narx ({{ selectedProduct.quantity_per_pack }} kapsul)</div>
                <div class="text-2xl font-bold text-brand-700">{{ formatPrice(selectedProduct.price_per_pack) }} so'm</div>
              </div>
              <button
                @click="cartStore.addItem(selectedProduct); selectedProduct = null"
                class="flex items-center gap-2 bg-brand-700 text-white px-6 py-3 rounded-xl font-semibold hover:bg-brand-800 transition-colors text-sm">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 10-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 01-1.12-1.243l1.264-12A1.125 1.125 0 015.513 7.5h12.974c.576 0 1.059.435 1.119 1.007z"/>
                </svg>
                Savatga qo'shish
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>

    <Footer />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import CartDrawer from '../components/CartDrawer.vue'
import OrdersDrawer from '../components/OrdersDrawer.vue'
import Footer from '../components/Footer.vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'

const cartStore = useCartStore()
const authStore = useAuthStore()

const products = ref([])
const loading = ref(true)
const selectedCategory = ref('')
const selectedProduct = ref(null)
const cartDrawerRef = ref(null)

const categories = [
  { key: '', label: 'Barchasi' },
  { key: 'Immunitet', label: 'Immunitet' },
  { key: 'Jigar uchun', label: 'Jigar uchun' },
  { key: 'Asab tizimi', label: 'Asab tizimi' },
  { key: "Bo'g'imlar", label: "Bo'g'imlar" },
  { key: 'Yurak-qon tomir', label: 'Yurak-qon tomir' },
  { key: 'Boshqalar', label: 'Boshqalar' },
]

const filteredProducts = computed(() => {
  if (!selectedCategory.value) return products.value
  return products.value.filter(p => p.category === selectedCategory.value)
})

const featuredProduct = computed(() => {
  if (!products.value.length) return null
  const withImage = products.value.find(p => p.image_path)
  return withImage || products.value[0]
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
