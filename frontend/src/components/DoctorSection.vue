<template>
  <section id="doctor" class="py-24 bg-white relative overflow-hidden">
    <!-- Subtle background decoration -->
    <div class="absolute top-0 right-0 w-[600px] h-[600px] bg-brand-50 rounded-full -translate-y-1/2 translate-x-1/3 blur-3xl opacity-50"></div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-2 gap-16 lg:gap-24 items-center">
        <!-- Image side -->
        <div class="relative">
          <div class="absolute -inset-4 bg-gradient-to-br from-brand-100 to-brand-50 rounded-3xl -rotate-2"></div>
          <div class="absolute -bottom-6 -right-6 w-32 h-32 bg-brand-100 rounded-2xl rotate-6 hidden lg:block"></div>
          <img
            src="/images/doctor.png"
            alt="Директор клиники"
            class="relative rounded-2xl w-full max-w-lg mx-auto object-cover shadow-2xl shadow-stone-900/10"
          />
          <!-- Experience badge -->
          <div class="absolute -bottom-4 -left-4 md:-bottom-6 md:-left-6 bg-white rounded-2xl shadow-xl shadow-stone-900/10 p-4 border border-stone-100">
            <div class="text-2xl font-serif text-brand-700">10+</div>
            <div class="text-xs font-medium text-stone-500 mt-0.5">лет опыта</div>
          </div>
        </div>

        <!-- Text side -->
        <div>
          <div class="section-label mb-5">
            <span class="w-8 h-px bg-brand-400"></span>
            Ведущий специалист
          </div>

          <h2 class="heading-lg mb-6">
            Наш специалист по
            <span class="text-brand-600">лечению волос</span>
          </h2>

          <p class="text-stone-500 text-lg leading-relaxed mb-5">
            Опытный специалист-трихолог с многолетним стажем работы в области
            диагностики и лечения заболеваний волос и кожи головы.
          </p>
          <p class="text-stone-500 text-lg leading-relaxed mb-10">
            Наша клиника использует современные препараты и индивидуальный подход
            к каждому пациенту для достижения максимального результата лечения.
          </p>

          <!-- Features grid -->
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
            <div class="bg-stone-50 rounded-xl p-4 border border-stone-100">
              <div class="w-10 h-10 bg-brand-100 rounded-lg flex items-center justify-center mb-3">
                <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">Индивидуальный</div>
              <div class="text-xs text-stone-400 mt-0.5">подход</div>
            </div>
            <div class="bg-stone-50 rounded-xl p-4 border border-stone-100">
              <div class="w-10 h-10 bg-amber-50 rounded-lg flex items-center justify-center mb-3">
                <svg class="w-5 h-5 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9.75 3.104v5.714a2.25 2.25 0 01-.659 1.591L5 14.5M9.75 3.104c-.251.023-.501.05-.75.082m.75-.082a24.301 24.301 0 014.5 0m0 0v5.714c0 .597.237 1.17.659 1.591L19.8 15.3M14.25 3.104c.251.023.501.05.75.082M19.8 15.3l-1.57.393A9.065 9.065 0 0112 15a9.065 9.065 0 00-6.23.693L5 14.5m14.8.8l1.402 1.402c1.232 1.232.65 3.318-1.067 3.611A48.309 48.309 0 0112 21c-2.773 0-5.491-.235-8.135-.687-1.718-.293-2.3-2.379-1.067-3.61L5 14.5" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">Современные</div>
              <div class="text-xs text-stone-400 mt-0.5">методики</div>
            </div>
            <div class="bg-stone-50 rounded-xl p-4 border border-stone-100">
              <div class="w-10 h-10 bg-brand-100 rounded-lg flex items-center justify-center mb-3">
                <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285z" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">Проверенные</div>
              <div class="text-xs text-stone-400 mt-0.5">препараты</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
