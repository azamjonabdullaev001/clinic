<template>
  <div class="bg-surface">
    <Navbar />
    <CartDrawer ref="cartDrawerRef" />
    <OrdersDrawer />

    <!-- Hero -->
    <section class="relative min-h-screen flex items-center bg-gradient-to-br from-brand-700 via-brand-800 to-brand-900 overflow-hidden pt-[96px] md:pt-[112px]">
      <!-- Animated decorative blobs -->
      <div class="absolute top-0 right-0 w-[700px] h-[700px] bg-brand-500/20 rounded-full -translate-y-1/4 translate-x-1/4 animate-pulse-soft"></div>
      <div class="absolute bottom-0 left-0 w-[500px] h-[500px] bg-brand-600/20 rounded-full translate-y-1/4 -translate-x-1/4 animate-pulse-soft" style="animation-delay: 1.5s"></div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 pb-24 md:pt-24 md:pb-32 w-full">
        <div class="grid md:grid-cols-2 gap-12 lg:gap-20 items-center">
          <!-- LEFT: Text -->
          <div>
            <div class="flex items-center gap-3 mb-6 animate-fade-up text-brand-200 text-xs font-semibold tracking-widest uppercase" style="animation-delay: 0.1s; opacity: 0">
              <span class="w-8 h-px bg-brand-400"></span>
              {{ t.hero_label }}
            </div>

            <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight mb-6 animate-fade-up" style="animation-delay: 0.25s; opacity: 0">
              {{ t.hero_title1 }}
              <span class="text-brand-300">{{ t.hero_title2 }}</span>
            </h1>

            <p class="text-lg md:text-xl text-blue-100/80 leading-relaxed mb-10 max-w-xl animate-fade-up" style="animation-delay: 0.4s; opacity: 0">
              {{ t.hero_desc }}
            </p>

            <div class="flex flex-col sm:flex-row gap-4 animate-fade-up" style="animation-delay: 0.55s; opacity: 0">
              <a href="#products" class="inline-flex items-center justify-center gap-2 bg-white text-brand-700 px-8 py-3.5 rounded-full font-semibold hover:bg-brand-50 hover:-translate-y-0.5 hover:shadow-xl hover:shadow-white/20 active:translate-y-0 transition-all duration-300 shadow-lg group">
                {{ t.hero_catalog }}
                <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
              </a>
              <a href="#doctor" class="inline-flex items-center justify-center gap-2 border border-white/30 text-white px-8 py-3.5 rounded-full font-medium hover:bg-white/10 hover:border-white/50 transition-all duration-300">
                {{ t.hero_about }}
              </a>
            </div>

            <!-- Stats row -->
            <div class="grid grid-cols-2 gap-6 mt-14 pt-10 border-t border-white/20 animate-fade-up" style="animation-delay: 0.7s; opacity: 0">
              <div>
                <div class="text-3xl md:text-4xl font-bold text-white">40<span class="text-brand-300">+</span></div>
                <div class="text-sm text-blue-200 mt-1">{{ t.hero_years }}</div>
              </div>
              <div>
                <div class="text-3xl md:text-4xl font-bold text-white">500000<span class="text-brand-300">+</span></div>
                <div class="text-sm text-blue-200 mt-1">{{ t.hero_patients }}</div>
              </div>
            </div>
          </div>

          <!-- RIGHT: Doctor image -->
          <div class="relative hidden md:flex justify-center animate-fade-up" style="animation-delay: 0.35s; opacity: 0">
            <div class="relative w-full max-w-md">
              <!-- Decorative ring -->
              <div class="absolute -inset-4 rounded-3xl bg-white/10 backdrop-blur-sm border border-white/20 z-10"></div>
              <img
                src="/images/patients/3%20landing.jpg"
                alt="Специалист клиники"
                class="relative z-20 rounded-2xl w-full object-cover object-top shadow-2xl shadow-brand-900/40 scale-110"
                style="max-height: 550px;"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Trust bar -->
    <section class="bg-brand-50 py-5 border-y border-brand-100/60">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-wrap items-center justify-center gap-x-10 gap-y-3 text-sm text-brand-700">
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M16.403 12.652a3 3 0 000-5.304 3 3 0 00-3.75-3.751 3 3 0 00-5.305 0 3 3 0 00-3.751 3.75 3 3 0 000 5.305 3 3 0 003.75 3.751 3 3 0 005.305 0 3 3 0 003.751-3.75zm-2.546-4.46a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd" /></svg>
            {{ t.trust_certified }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" /></svg>
            {{ t.trust_payment }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path d="M8 16.25a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75z" /><path fill-rule="evenodd" d="M4 4a3 3 0 013-3h6a3 3 0 013 3v12a3 3 0 01-3 3H7a3 3 0 01-3-3V4zm3-1.5A1.5 1.5 0 005.5 4v12A1.5 1.5 0 007 17.5h6a1.5 1.5 0 001.5-1.5V4A1.5 1.5 0 0013 2.5H7z" clip-rule="evenodd" /></svg>
            {{ t.trust_consult }}
          </div>
          <div class="flex items-center gap-2 hover:text-brand-800 transition-colors duration-300 cursor-default">
            <svg class="w-4 h-4 text-brand-500" fill="currentColor" viewBox="0 0 20 20"><path d="M10 2a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 2zm0 13a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 15zm0-8a3 3 0 100 6 3 3 0 000-6zm-4.243-2.757a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.061zM2 10a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 012 10zm13 0a.75.75 0 01.75-.75h1.5a.75.75 0 010 1.5h-1.5A.75.75 0 0115 10zM4.343 4.343a.75.75 0 011.061 0l1.06 1.06a.75.75 0 01-1.06 1.061l-1.06-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.061z" /></svg>
            {{ t.trust_delivery }}
          </div>
        </div>
      </div>
    </section>

    <DoctorSection />

    <!-- Latest News Carousel -->
    <section id="news" class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12 reveal">
          <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-semibold tracking-widest uppercase mb-4">
            <span class="w-8 h-px bg-brand-400"></span>
            {{ t.nav_news || 'Новости' }}
            <span class="w-8 h-px bg-brand-400"></span>
          </div>
          <h2 class="text-3xl md:text-4xl font-bold text-stone-900 mb-3">{{ t.news_title || 'Последние новости' }}</h2>
          <p class="text-stone-500 text-lg max-w-xl mx-auto">{{ t.news_subtitle || 'Актуальные события и обновления' }}</p>
        </div>

        <!-- Loading -->
        <div v-if="newsLoading" class="flex justify-center py-10">
          <div class="w-8 h-8 border-[3px] border-brand-200 border-t-brand-600 rounded-full animate-spin"></div>
        </div>

        <!-- Empty -->
        <div v-else-if="latestNews.length === 0" class="text-center py-12 text-stone-400">
          <svg class="w-12 h-12 mx-auto mb-3 text-stone-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" />
          </svg>
          <p class="text-base font-medium">Новостей пока нет</p>
        </div>

        <!-- Carousel -->
        <div v-else class="relative select-none">
          <div class="relative overflow-hidden rounded-3xl bg-stone-900 shadow-2xl h-[360px] md:h-[420px]">

            <div
              v-for="(post, idx) in latestNews"
              :key="post.id"
              v-show="newsSlide === idx"
              class="absolute inset-0 flex flex-col md:flex-row"
            >
              <!-- === VIDEO SLIDE === -->
              <template v-if="post.video_url && getYouTubeId(post.video_url)">
                <div class="relative flex-shrink-0 w-full md:w-2/3 h-[55%] md:h-full overflow-hidden">
                  <iframe
                    v-if="newsSlide === idx"
                    :key="`yt-${post.id}-${newsMuted[idx] !== false}`"
                    :src="getNewsEmbedUrl(post.video_url, idx)"
                    class="w-full h-full"
                    frameborder="0"
                    allow="autoplay; encrypted-media; fullscreen; picture-in-picture"
                    allowfullscreen
                  ></iframe>
                  <button
                    @click.stop="toggleNewsMute(idx)"
                    class="absolute top-3 right-3 z-10 bg-black/50 hover:bg-black/70 text-white rounded-full w-9 h-9 flex items-center justify-center transition-colors"
                  >
                    <svg v-if="newsMuted[idx] !== false" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z"/>
                      <path stroke-linecap="round" stroke-linejoin="round" d="M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2"/>
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M15.536 8.464a5 5 0 010 7.072M12 6v12M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z"/>
                    </svg>
                  </button>
                </div>
                <div class="flex-1 p-5 md:p-7 flex flex-col justify-between bg-stone-900 cursor-pointer overflow-hidden" @click="openNewsUrl(post)">
                  <div>
                    <p class="text-xs text-stone-400 mb-2">{{ formatNewsDate(post.created_at) }}</p>
                    <h3 class="font-bold text-white text-base md:text-lg leading-snug mb-3 line-clamp-3">{{ post.title }}</h3>
                    <p v-if="post.description" class="text-stone-400 text-sm leading-relaxed line-clamp-4">{{ post.description }}</p>
                  </div>
                  <span class="mt-4 text-brand-400 text-xs font-semibold flex items-center gap-1">
                    {{ t.news_watch || 'Смотреть' }}
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
                  </span>
                </div>
              </template>

              <!-- === IMAGE / TEXT SLIDE === -->
              <template v-else>
                <div class="flex-shrink-0 w-full md:w-2/3 h-[55%] md:h-full overflow-hidden cursor-pointer" @click="openNewsUrl(post)">
                  <img v-if="post.images && post.images.length > 0" :src="post.images[0].image_path" :alt="post.title" class="w-full h-full object-cover hover:scale-105 transition-transform duration-700" />
                  <img v-else-if="post.image_path" :src="post.image_path" :alt="post.title" class="w-full h-full object-cover hover:scale-105 transition-transform duration-700" />
                  <div v-else class="w-full h-full bg-stone-800 flex flex-col items-center justify-center text-stone-500">
                    <svg class="w-12 h-12 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" /></svg>
                    <span class="text-sm">Новость</span>
                  </div>
                </div>
                <div class="flex-1 p-5 md:p-7 flex flex-col justify-between bg-stone-900 cursor-pointer overflow-hidden" @click="openNewsUrl(post)">
                  <div>
                    <p class="text-xs text-stone-400 mb-2">{{ formatNewsDate(post.created_at) }}</p>
                    <h3 class="font-bold text-white text-base md:text-lg leading-snug mb-3 line-clamp-3">{{ post.title }}</h3>
                    <p v-if="post.description" class="text-stone-400 text-sm leading-relaxed line-clamp-4">{{ post.description }}</p>
                  </div>
                  <span v-if="post.video_url" class="mt-4 text-brand-400 text-xs font-semibold flex items-center gap-1">
                    {{ t.news_watch || 'Открыть' }}
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"/></svg>
                  </span>
                </div>
              </template>
            </div>

            <!-- Left arrow -->
            <button
              v-if="latestNews.length > 1"
              @click="prevNewsSlide"
              class="absolute left-3 top-1/2 -translate-y-1/2 z-20 bg-black/40 hover:bg-black/60 text-white rounded-full w-9 h-9 flex items-center justify-center transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7"/></svg>
            </button>

            <!-- Right arrow -->
            <button
              v-if="latestNews.length > 1"
              @click="nextNewsSlide"
              class="absolute right-3 top-1/2 -translate-y-1/2 z-20 bg-black/40 hover:bg-black/60 text-white rounded-full w-9 h-9 flex items-center justify-center transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7"/></svg>
            </button>

            <!-- Progress bar -->
            <div class="absolute bottom-0 left-0 right-0 h-0.5 bg-white/10">
              <div class="h-full bg-brand-500/80 transition-all duration-500" :style="{ width: `${((newsSlide + 1) / latestNews.length) * 100}%` }"></div>
            </div>
          </div>

          <!-- Dot indicators -->
          <div v-if="latestNews.length > 1" class="flex items-center justify-center gap-2 mt-5">
            <button
              v-for="(_, i) in latestNews"
              :key="i"
              @click="goToNewsSlide(i)"
              class="rounded-full transition-all duration-300 focus:outline-none"
              :class="newsSlide === i ? 'w-6 h-2 bg-brand-600' : 'w-2 h-2 bg-stone-300 hover:bg-stone-400'"
            ></button>
          </div>
        </div>

        <div class="text-center mt-10">
          <router-link
            to="/news"
            class="inline-flex items-center gap-2 border border-brand-200 text-brand-700 px-7 py-3 rounded-full font-semibold hover:bg-brand-50 hover:border-brand-400 transition-all duration-300 group"
          >
            {{ t.news_all || 'Все новости' }}
            <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
          </router-link>
        </div>
      </div>
    </section>

    <!-- Products -->
    <section id="products" class="py-24 bg-gradient-to-b from-white via-surface-warm to-surface relative">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16 reveal" ref="productsHeader">
          <div class="section-label justify-center mb-4">
            <span class="w-8 h-px bg-brand-400"></span>
            {{ t.products_label }}
            <span class="w-8 h-px bg-brand-400"></span>
          </div>
          <h2 class="heading-lg mb-4">{{ t.products_title }}</h2>
          <p class="text-stone-500 text-lg max-w-2xl mx-auto leading-relaxed">{{ t.products_desc }}</p>
        </div>

        <div v-if="loading" class="text-center py-20">
          <div class="inline-block w-10 h-10 border-[3px] border-brand-200 border-t-brand-600 rounded-full animate-spin"></div>
          <p class="mt-4 text-stone-400 font-medium">{{ t.products_loading }}</p>
        </div>

        <div v-else-if="products.length === 0" class="text-center py-20">
          <div class="w-20 h-20 bg-stone-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
            <svg class="w-10 h-10 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <p class="text-stone-400 text-lg font-medium">{{ t.products_empty }}</p>
        </div>

        <div v-else class="grid grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3 sm:gap-6 stagger-children">
          <ProductCard
            v-for="(product, index) in products"
            :key="product.id"
            :product="product"
            :style="{ '--stagger': index }"
            class="reveal"
            :ref="el => { if (el) productRefs[index] = el.$el || el }"
            @add-to-cart="cartStore.addItem"
          />
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="relative overflow-hidden reveal" ref="ctaSection">
      <div class="absolute inset-0 bg-gradient-to-br from-brand-700 via-brand-800 to-brand-900"></div>
      <div class="absolute top-0 right-0 w-96 h-96 bg-brand-500/10 rounded-full blur-[100px] animate-pulse-soft"></div>
      <div class="absolute bottom-0 left-0 w-72 h-72 bg-brand-400/10 rounded-full blur-[80px] animate-pulse-soft" style="animation-delay: 2s"></div>
      <div class="relative max-w-4xl mx-auto px-4 py-24 text-center">
        <div class="section-label !text-brand-200 justify-center mb-6">
          <span class="w-8 h-px bg-brand-400/50"></span>
          {{ t.cta_label }}
          <span class="w-8 h-px bg-brand-400/50"></span>
        </div>
        <h2 class="heading-lg !text-white mb-5 text-balance">{{ t.cta_title }}</h2>
        <p class="text-brand-200/70 text-lg mb-10 max-w-lg mx-auto leading-relaxed">{{ t.cta_desc }}</p>
        <router-link
          v-if="!authStore.isLoggedIn"
          to="/register"
          class="inline-block bg-white text-brand-800 px-8 py-4 rounded-full font-semibold hover:bg-brand-50 hover:shadow-2xl hover:shadow-white/10 hover:-translate-y-1 active:scale-[0.97] active:translate-y-0 transition-all duration-300 shadow-lg group"
        >
          <span class="inline-flex items-center gap-2">
            {{ t.cta_register }}
            <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
          </span>
        </router-link>
      </div>
    </section>

    <Footer />

    <!-- Order success modal (kept for reference, not used) -->
    <Teleport to="body">
      <div v-if="showSuccess" class="fixed inset-0 z-[60] flex items-center justify-center">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-md animate-fade-in" @click="showSuccess = false"></div>
        <div class="relative bg-white rounded-3xl p-10 max-w-sm mx-4 text-center shadow-2xl animate-scale-in">
          <div class="w-20 h-20 bg-brand-50 rounded-2xl flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h3 class="text-2xl font-serif text-stone-900 mb-2">Заказ оформлен!</h3>
          <p class="text-stone-500 mb-4 leading-relaxed">Мы свяжемся с вами в ближайшее время для подтверждения</p>
          <div v-if="lastOrderCode" class="bg-blue-50 border border-blue-200 rounded-2xl p-4 mb-6">
            <p class="text-xs text-blue-500 font-medium uppercase tracking-wider mb-1">Ваш код заказа</p>
            <p class="text-4xl font-bold text-blue-700 tracking-widest">{{ lastOrderCode }}</p>
            <p class="text-xs text-blue-400 mt-1">Назовите этот код на пункте выдачи</p>
          </div>
          <button @click="showSuccess = false" class="btn-primary w-full">
            Отлично
          </button>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import DoctorSection from '../components/DoctorSection.vue'
import ProductCard from '../components/ProductCard.vue'
import CartDrawer from '../components/CartDrawer.vue'
import OrdersDrawer from '../components/OrdersDrawer.vue'
import Footer from '../components/Footer.vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore, api } from '../stores/auth'
import { useLangStore } from '../stores/lang'

const cartStore = useCartStore()
const authStore = useAuthStore()
const langStore = useLangStore()
const t = computed(() => langStore.t)
const cartDrawerRef = ref(null)
const products = ref([])
const loading = ref(true)
const showSuccess = ref(false)
const lastOrderCode = ref('')
const latestNews = ref([])
const newsLoading = ref(true)
const newsSlide = ref(0)
const newsMuted = ref({})
let newsTimer = null

function getYouTubeId(url) {
  if (!url) return null
  const m = url.match(/(?:youtube\.com\/(?:watch\?v=|embed\/)|youtu\.be\/)([^&?#\n]+)/)
  return m ? m[1] : null
}

function getNewsEmbedUrl(videoUrl, idx) {
  const id = getYouTubeId(videoUrl)
  if (!id) return ''
  const muted = newsMuted.value[idx] !== false ? 1 : 0
  return `https://www.youtube.com/embed/${id}?autoplay=1&mute=${muted}&enablejsapi=1&playsinline=1&rel=0`
}

function toggleNewsMute(idx) {
  const currentlyMuted = newsMuted.value[idx] !== false
  newsMuted.value = { ...newsMuted.value, [idx]: !currentlyMuted }
}

function openNewsUrl(post) {
  if (post.video_url) window.open(post.video_url, '_blank', 'noopener,noreferrer')
}

function formatNewsDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('ru-RU', { year: 'numeric', month: 'long', day: 'numeric' })
}

function nextNewsSlide() {
  if (!latestNews.value.length) return
  newsSlide.value = (newsSlide.value + 1) % latestNews.value.length
  restartNewsTimer()
}

function prevNewsSlide() {
  if (!latestNews.value.length) return
  newsSlide.value = (newsSlide.value - 1 + latestNews.value.length) % latestNews.value.length
  restartNewsTimer()
}

function goToNewsSlide(idx) {
  newsSlide.value = idx
  restartNewsTimer()
}

function startNewsTimer() {
  if (newsTimer) clearInterval(newsTimer)
  newsTimer = setInterval(() => {
    if (latestNews.value.length > 1) {
      newsSlide.value = (newsSlide.value + 1) % latestNews.value.length
    }
  }, 10000)
}

function restartNewsTimer() {
  startNewsTimer()
}

// Scroll reveal refs
const productsHeader = ref(null)
const ctaSection = ref(null)
const productRefs = ref({})
let observer = null

function setupObserver() {
  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible')
          observer.unobserve(entry.target)
        }
      })
    },
    { threshold: 0.1, rootMargin: '0px 0px -40px 0px' }
  )

  // Observe direct refs
  if (productsHeader.value) observer.observe(productsHeader.value)
  if (ctaSection.value) observer.observe(ctaSection.value)

  // Observe all .reveal elements
  document.querySelectorAll('.reveal, .reveal-left, .reveal-right, .reveal-scale').forEach((el) => {
    observer.observe(el)
  })
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

  try {
    const newsRes = await axios.get('/api/news')
    latestNews.value = newsRes.data || []
    if (latestNews.value.length > 1) startNewsTimer()
  } catch { /* ignore */ } finally {
    newsLoading.value = false
  }

  await nextTick()
  setupObserver()
})

onUnmounted(() => {
  if (observer) observer.disconnect()
  if (newsTimer) clearInterval(newsTimer)
})


</script>
