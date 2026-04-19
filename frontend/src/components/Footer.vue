<template>
  <footer id="contacts" class="bg-brand-950 text-white relative overflow-hidden grain">
    <!-- Decorative -->
    <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-px bg-gradient-to-r from-transparent via-brand-700/50 to-transparent"></div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-16 pb-10">
      <div class="grid md:grid-cols-4 gap-10 pb-12 border-b border-white/10">
        <!-- Brand -->
        <div class="md:col-span-2">
          <div class="flex items-center gap-3 mb-5">
            <div class="w-10 h-10 rounded-xl bg-brand-800 flex items-center justify-center">
              <svg class="w-5 h-5 text-gold-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 007.92 12.446A9 9 0 1112 2.992z" />
              </svg>
            </div>
            <div>
              <span class="text-lg font-bold tracking-tight">Hair Clinic</span>
              <span class="block text-[10px] font-medium tracking-widest uppercase text-gold-400">Трихология</span>
            </div>
          </div>
          <p class="text-white/40 text-sm leading-relaxed max-w-sm">
            Профессиональное лечение и восстановление волос с использованием проверенных препаратов и индивидуального подхода
          </p>
        </div>

        <!-- Contacts -->
        <div>
          <h4 class="text-xs font-semibold tracking-widest uppercase text-white/60 mb-5">Контакты</h4>
          <div class="space-y-3 text-sm text-white/40">
            <a href="tel:+998901475130" class="flex items-center gap-2.5 hover:text-white transition-colors">
              <svg class="w-4 h-4 text-gold-400/60" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8"><path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z" /></svg>
              +998 90 147 51 30
            </a>
            <div class="flex items-center gap-2.5">
              <svg class="w-4 h-4 text-gold-400/60" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8"><path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" /><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" /></svg>
              Узбекистан
            </div>
          </div>
        </div>

        <!-- Working hours -->
        <div>
          <h4 class="text-xs font-semibold tracking-widest uppercase text-white/60 mb-5">Режим работы</h4>
          <div class="space-y-3 text-sm text-white/40">
            <div class="flex justify-between">
              <span>Пн — Сб</span>
              <span class="text-white/60 font-medium">9:00 — 18:00</span>
            </div>
            <div class="flex justify-between">
              <span>Воскресенье</span>
              <span class="text-white/60 font-medium">Выходной</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Bottom bar -->
      <div class="mt-8 flex flex-col md:flex-row justify-between items-center gap-4">
        <p class="text-xs text-white/30">
          © {{ new Date().getFullYear() }} Hair Clinic. Все права защищены.
        </p>
        <router-link to="/admin/login" class="text-xs text-white/20 hover:text-white/40 transition-colors">
          Админ
        </router-link>
      </div>
    </div>
  </footer>
</template>
