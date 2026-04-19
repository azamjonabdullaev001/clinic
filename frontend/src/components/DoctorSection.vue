<template>
  <section id="doctor" class="py-20 bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-2 gap-12 lg:gap-16 items-center">
        <div class="relative">
          <div class="absolute -top-4 -left-4 w-72 h-72 bg-teal-100 rounded-full blur-3xl opacity-50"></div>
          <img
            src="/images/doctor.jpg"
            alt="Директор клиники"
            class="relative rounded-2xl shadow-2xl w-full max-w-lg mx-auto object-cover"
          />
        </div>
        <div>
          <div class="inline-flex items-center gap-2 bg-teal-50 text-teal-700 px-4 py-2 rounded-full text-sm font-medium mb-6">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
            Ведущий специалист
          </div>
          <h2 class="text-3xl md:text-4xl font-extrabold text-gray-800 mb-5 leading-tight">
            Наш специалист по<br /><span class="text-teal-600">лечению волос</span>
          </h2>
          <p class="text-gray-600 text-lg leading-relaxed mb-5">
            Опытный специалист-трихолог с многолетним стажем работы в области
            диагностики и лечения заболеваний волос и кожи головы.
          </p>
          <p class="text-gray-600 text-lg leading-relaxed mb-8">
            Наша клиника использует современные препараты и индивидуальный подход
            к каждому пациенту для достижения максимального результата лечения.
          </p>
          <div class="flex flex-wrap gap-4">
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-5 h-5 text-teal-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="text-sm font-medium">Индивидуальный подход</span>
            </div>
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-5 h-5 text-teal-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="text-sm font-medium">Современные методики</span>
            </div>
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-5 h-5 text-teal-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="text-sm font-medium">Проверенные препараты</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
