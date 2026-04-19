<template>
  <footer class="bg-gray-900 text-gray-400 pt-12 pb-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-3 gap-8 pb-8 border-b border-gray-800">
        <div>
          <div class="flex items-center gap-2.5 mb-4">
            <div class="w-8 h-8 bg-teal-600 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </div>
            <span class="text-white font-bold text-lg">Hair Clinic</span>
          </div>
          <p class="text-sm leading-relaxed">
            Профессиональное лечение и восстановление волос с использованием проверенных препаратов
          </p>
        </div>

        <div>
          <h4 class="text-white font-semibold mb-4">Контакты</h4>
          <div class="space-y-2 text-sm">
            <p>📞 +998 90 147 51 30</p>
            <p>📍 Узбекистан</p>
          </div>
        </div>

        <div>
          <h4 class="text-white font-semibold mb-4">Режим работы</h4>
          <div class="space-y-2 text-sm">
            <p>Пн — Сб: 9:00 — 18:00</p>
            <p>Вс: выходной</p>
          </div>
        </div>
      </div>

      <div class="mt-8 flex flex-col md:flex-row justify-between items-center gap-4">
        <p class="text-sm text-gray-500">
          © {{ new Date().getFullYear() }} Hair Clinic. Все права защищены.
        </p>
        <router-link to="/admin/login" class="text-xs text-gray-600 hover:text-gray-400 transition">
          Админ
        </router-link>
      </div>
    </div>
  </footer>
</template>
