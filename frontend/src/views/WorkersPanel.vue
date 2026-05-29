<template>
  <div class="workers-panel min-h-screen flex flex-col" style="background: #0d1525; color: white;">

    <!-- ===== HEADER ===== -->
    <header class="flex-shrink-0 flex items-center justify-between px-5 h-16"
      style="background: #111827; border-bottom: 1px solid rgba(255,255,255,0.07);">
      <!-- Brand -->
      <div class="flex items-center gap-3">
        <div class="w-9 h-9 rounded-xl bg-blue-600 flex items-center justify-center flex-shrink-0 shadow-lg shadow-blue-900/50">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
        <div>
          <p class="text-white text-sm font-bold leading-tight">Панель работников</p>
          <p class="text-gray-500 text-xs leading-tight">{{ authStore.admin?.phone || 'Администратор' }}</p>
        </div>
      </div>

      <!-- Right side -->
      <div class="flex items-center gap-4">
        <button class="relative text-gray-500 hover:text-gray-300 transition p-1.5 rounded-lg hover:bg-white/5">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
          </svg>
        </button>
        <div class="flex items-center gap-2 cursor-pointer group select-none">
          <div class="w-8 h-8 rounded-full bg-blue-700 flex items-center justify-center text-white text-xs font-bold flex-shrink-0">А</div>
          <span class="text-white text-sm font-medium hidden sm:block">Администратор</span>
          <svg class="w-4 h-4 text-gray-500 group-hover:text-gray-300 transition hidden sm:block" fill="none" stroke="currentColor"
            viewBox="0 0 24 24" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </header>

    <!-- ===== BODY ===== -->
    <div class="flex flex-1">

      <!-- ===== SIDEBAR ===== -->
      <aside class="flex-shrink-0 flex flex-col"
        style="width: 196px; background: #111827; border-right: 1px solid rgba(255,255,255,0.06); min-height: calc(100vh - 64px);">
        <nav class="flex-1 p-3 space-y-0.5 mt-2">
          <button v-for="item in navItems" :key="item.key" @click="tab = item.key"
            :class="tab === item.key
              ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/40'
              : 'text-gray-400 hover:bg-white/5 hover:text-gray-200'"
            class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-all">
            <span class="w-4 h-4 flex-shrink-0" v-html="item.icon"></span>
            {{ item.label }}
          </button>
        </nav>

        <!-- Bottom: Lang + Logout -->
        <div class="p-4 space-y-4" style="border-top: 1px solid rgba(255,255,255,0.06);">
          <div class="flex items-center gap-1.5">
            <button @click="lang = 'ru'"
              :class="lang === 'ru' ? 'text-white font-bold' : 'text-gray-600 hover:text-gray-400'"
              class="text-sm transition">RU</button>
            <span class="text-gray-700 text-sm">|</span>
            <button @click="lang = 'uz'"
              :class="lang === 'uz' ? 'text-white font-bold' : 'text-gray-600 hover:text-gray-400'"
              class="text-sm transition">UZ</button>
          </div>
          <button @click="logout"
            class="flex items-center gap-2 text-red-400 hover:text-red-300 text-sm font-medium transition w-full">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
            Выйти
          </button>
        </div>
      </aside>

      <!-- ===== MAIN CONTENT ===== -->
      <main class="flex-1 p-6 overflow-auto">

        <!-- ===================================== -->
        <!-- TAB: РАБОТНИКИ                        -->
        <!-- ===================================== -->
        <div v-if="tab === 'workers'">
          <!-- Header -->
          <div class="flex flex-wrap items-center justify-between gap-4 mb-6">
            <div>
              <h1 class="text-white text-xl font-bold">Работники</h1>
              <p class="text-gray-500 text-sm mt-0.5">{{ workers.length }} сотрудников в системе</p>
            </div>
            <button @click="openAddModal"
              class="flex items-center gap-2 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2.5 rounded-xl font-medium text-sm transition-all shadow-lg shadow-blue-900/30">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              Добавить работника
            </button>
          </div>

          <!-- Filters -->
          <div class="flex flex-wrap gap-3 mb-5">
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-gray-600 pointer-events-none"
                fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input v-model="searchQuery" placeholder="Поиск по имени или телефону..."
                class="pl-9 pr-4 py-2.5 rounded-xl text-sm text-white placeholder-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition w-64"
                style="background: #1a2540; border: 1px solid rgba(255,255,255,0.08);" />
            </div>
            <select v-model="roleFilter"
              class="px-4 py-2.5 rounded-xl text-sm outline-none focus:ring-2 focus:ring-blue-500 transition"
              style="background: #1a2540; border: 1px solid rgba(255,255,255,0.08); color: #d1d5db;">
              <option value="">Все роли</option>
              <option value="pickup">Пункт выдачи</option>
              <option value="nurse">Медсестра</option>
              <option value="manager">Менеджер</option>
            </select>
          </div>

          <!-- Table card -->
          <div class="rounded-2xl overflow-hidden" style="background: #141e33; border: 1px solid rgba(255,255,255,0.07);">
            <!-- Table head -->
            <div class="hidden md:grid px-5 py-3 text-xs font-semibold tracking-wider uppercase"
              style="grid-template-columns: 2fr 1.5fr 1fr 120px; color: #6b7280; border-bottom: 1px solid rgba(255,255,255,0.06);">
              <span>Работник</span>
              <span>Телефон</span>
              <span>Роль</span>
              <span class="text-center">Действия</span>
            </div>

            <!-- Loading -->
            <div v-if="loading" class="flex items-center justify-center py-16">
              <svg class="w-6 h-6 text-blue-500 animate-spin mr-3" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
              </svg>
              <span class="text-gray-400 text-sm">Загрузка...</span>
            </div>

            <!-- Empty -->
            <div v-else-if="filteredWorkers.length === 0" class="flex flex-col items-center justify-center py-16 px-6">
              <div class="w-14 h-14 rounded-2xl flex items-center justify-center mb-4"
                style="background: rgba(255,255,255,0.04);">
                <svg class="w-7 h-7 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                  stroke-width="1.5">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <p class="text-gray-400 font-medium">Работники не найдены</p>
              <p class="text-gray-600 text-sm mt-1">
                {{ searchQuery || roleFilter ? 'Попробуйте изменить фильтры' : 'Добавьте первого работника' }}
              </p>
            </div>

            <!-- Rows -->
            <template v-else>
              <div v-for="(w, idx) in filteredWorkers" :key="w.id">
                <!-- Worker row -->
                <div class="grid px-5 py-4 items-center transition-colors hover:bg-white/2"
                  :class="{ 'border-b': idx < filteredWorkers.length - 1 || expandedWorkerId === w.id }"
                  style="grid-template-columns: 2fr 1.5fr 1fr 120px; border-color: rgba(255,255,255,0.05);">
                  <!-- Name + avatar -->
                  <div class="flex items-center gap-3 min-w-0">
                    <div class="w-9 h-9 rounded-xl flex items-center justify-center text-white text-sm font-bold flex-shrink-0"
                      :style="{ background: roleColor(w.role) }">
                      {{ (w.name || '?')[0].toUpperCase() }}
                    </div>
                    <span class="text-white text-sm font-medium truncate">{{ w.name }}</span>
                  </div>
                  <!-- Phone -->
                  <span class="text-gray-400 text-sm font-mono">+{{ w.phone }}</span>
                  <!-- Role badge -->
                  <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold w-fit"
                    :class="roleBadgeClass(w.role)">
                    {{ roleLabel(w.role) }}
                  </span>
                  <!-- Actions -->
                  <div class="flex items-center gap-1 justify-center">
                    <button @click="toggleStats(w)"
                      :class="expandedWorkerId === w.id
                        ? 'text-emerald-400 bg-emerald-500/15'
                        : 'text-gray-600 hover:text-emerald-400 hover:bg-emerald-500/10'"
                      class="p-2 rounded-lg transition" title="Статистика">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                          d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                      </svg>
                    </button>
                    <button @click="openEditModal(w)"
                      class="p-2 rounded-lg text-gray-600 hover:text-blue-400 hover:bg-blue-500/10 transition"
                      title="Редактировать">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                          d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                    </button>
                    <button @click="deleteWorker(w.id)"
                      class="p-2 rounded-lg text-gray-600 hover:text-red-400 hover:bg-red-500/10 transition"
                      title="Удалить">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                          d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                </div>

                <!-- Expanded stats panel -->
                <transition name="slide-down">
                  <div v-if="expandedWorkerId === w.id" class="px-5 pb-5 pt-4"
                    style="background: #0f1929; border-bottom: 1px solid rgba(255,255,255,0.05);">
                    <p class="text-xs font-semibold uppercase tracking-wider text-gray-600 mb-3">
                      Статистика — {{ w.name }}
                    </p>
                    <div v-if="statsLoading" class="flex items-center gap-2 text-sm text-gray-500">
                      <svg class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
                      </svg>
                      Загрузка статистики...
                    </div>
                    <div v-else-if="workerStats" class="grid grid-cols-1 sm:grid-cols-3 gap-3">
                      <div v-for="(period, key) in workerStats" :key="key" class="rounded-xl p-4"
                        style="background: #1a2540; border: 1px solid rgba(255,255,255,0.07);">
                        <p class="text-xs font-semibold uppercase tracking-wider mb-3"
                          :class="key === 'today' ? 'text-blue-400' : key === 'week' ? 'text-purple-400' : 'text-amber-400'">
                          {{ key === 'today' ? 'Сегодня' : key === 'week' ? 'Неделя' : 'Месяц' }}
                        </p>
                        <div class="space-y-2.5">
                          <div class="flex justify-between text-sm">
                            <span class="text-gray-500">Заказы</span>
                            <span class="text-white font-bold">{{ period.orders }}</span>
                          </div>
                          <div class="flex justify-between text-sm">
                            <span class="text-gray-500">Товаров</span>
                            <span class="text-white font-bold">{{ period.items }} шт</span>
                          </div>
                          <div class="flex justify-between text-sm pt-1.5"
                            style="border-top: 1px solid rgba(255,255,255,0.06);">
                            <span class="text-gray-500">Выручка</span>
                            <span class="text-emerald-400 font-bold">{{ formatPrice(period.revenue) }} сўм</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </transition>
              </div>
            </template>
          </div>
        </div>

        <!-- ===================================== -->
        <!-- TAB: СТАТИСТИКА                       -->
        <!-- ===================================== -->
        <div v-if="tab === 'stats'">
          <div class="flex flex-wrap items-center justify-between gap-4 mb-6">
            <div>
              <h1 class="text-white text-xl font-bold">Статистика работников</h1>
              <p class="text-gray-500 text-sm mt-0.5">Сравнительная аналитика по всем сотрудникам</p>
            </div>
            <button @click="loadAllStats"
              class="flex items-center gap-2 text-gray-400 hover:text-white text-sm transition px-3 py-2 rounded-xl hover:bg-white/5">
              <svg class="w-4 h-4" :class="{ 'animate-spin': allStatsLoading }" fill="none" stroke="currentColor"
                viewBox="0 0 24 24" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Обновить
            </button>
          </div>

          <!-- Period selector -->
          <div class="flex gap-2 mb-6">
            <button v-for="p in periods" :key="p.key" @click="statsPeriod = p.key"
              :class="statsPeriod === p.key
                ? 'bg-blue-600 text-white shadow-lg shadow-blue-900/30'
                : 'text-gray-500 hover:text-white hover:bg-white/5'"
              class="px-4 py-2 rounded-xl text-sm font-medium transition">
              {{ p.label }}
            </button>
          </div>

          <!-- Loading -->
          <div v-if="allStatsLoading" class="flex items-center justify-center py-16">
            <svg class="w-6 h-6 text-blue-500 animate-spin mr-3" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
            </svg>
            <span class="text-gray-400 text-sm">Загрузка статистики...</span>
          </div>

          <!-- Stats grid -->
          <div v-else-if="allWorkerStats.length > 0" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
            <div v-for="ws in allWorkerStats" :key="ws.worker.id" class="rounded-2xl p-5 transition-all hover:scale-[1.01]"
              style="background: #141e33; border: 1px solid rgba(255,255,255,0.07);">
              <!-- Worker info -->
              <div class="flex items-center gap-3 mb-4 pb-4"
                style="border-bottom: 1px solid rgba(255,255,255,0.06);">
                <div class="w-10 h-10 rounded-xl flex items-center justify-center text-white font-bold flex-shrink-0"
                  :style="{ background: roleColor(ws.worker.role) }">
                  {{ (ws.worker.name || '?')[0].toUpperCase() }}
                </div>
                <div class="min-w-0">
                  <p class="text-white font-semibold text-sm truncate">{{ ws.worker.name }}</p>
                  <span class="text-xs px-2 py-0.5 rounded-full inline-block mt-0.5"
                    :class="roleBadgeClass(ws.worker.role)">{{ roleLabel(ws.worker.role) }}</span>
                </div>
              </div>

              <!-- Period stats -->
              <div class="grid grid-cols-3 gap-2 text-center">
                <div>
                  <p class="text-gray-600 text-xs mb-1.5">Заказы</p>
                  <p class="text-white text-2xl font-bold leading-none">
                    {{ ws.stats[statsPeriod]?.orders || 0 }}
                  </p>
                </div>
                <div>
                  <p class="text-gray-600 text-xs mb-1.5">Товаров</p>
                  <p class="text-white text-2xl font-bold leading-none">
                    {{ ws.stats[statsPeriod]?.items || 0 }}
                  </p>
                </div>
                <div>
                  <p class="text-gray-600 text-xs mb-1.5">Выручка</p>
                  <p class="text-emerald-400 text-sm font-bold leading-none pt-1">
                    {{ formatPrice(ws.stats[statsPeriod]?.revenue || 0) }}
                  </p>
                  <p class="text-gray-600 text-xs">сўм</p>
                </div>
              </div>

              <!-- Mini progress bar (visual indicator) -->
              <div class="mt-4 pt-3" style="border-top: 1px solid rgba(255,255,255,0.05);">
                <div class="flex justify-between text-xs text-gray-600 mb-1.5">
                  <span>Активность</span>
                  <span>{{ ws.stats[statsPeriod]?.orders || 0 }} заказов</span>
                </div>
                <div class="h-1.5 rounded-full overflow-hidden" style="background: rgba(255,255,255,0.06);">
                  <div class="h-full rounded-full transition-all duration-500"
                    :style="{ width: getActivityWidth(ws) + '%', background: roleColor(ws.worker.role) }">
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty -->
          <div v-else class="flex flex-col items-center justify-center py-16">
            <div class="w-14 h-14 rounded-2xl flex items-center justify-center mb-4"
              style="background: rgba(255,255,255,0.04);">
              <svg class="w-7 h-7 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
            <p class="text-gray-400 font-medium">Нет данных</p>
            <p class="text-gray-600 text-sm mt-1">Добавьте работников для отображения статистики</p>
          </div>
        </div>

        <!-- ===================================== -->
        <!-- TAB: ИСТОРИЯ                          -->
        <!-- ===================================== -->
        <div v-if="tab === 'history'">
          <div class="mb-6">
            <h1 class="text-white text-xl font-bold">История действий</h1>
            <p class="text-gray-500 text-sm mt-0.5">Лог активности работников</p>
          </div>
          <div class="rounded-2xl p-10 flex flex-col items-center justify-center"
            style="background: #141e33; border: 1px solid rgba(255,255,255,0.07);">
            <div class="w-14 h-14 rounded-2xl flex items-center justify-center mb-4"
              style="background: rgba(255,255,255,0.04);">
              <svg class="w-7 h-7 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <p class="text-gray-400 font-medium">История пуста</p>
            <p class="text-gray-600 text-sm mt-1">Здесь будут отображаться действия работников</p>
          </div>
        </div>

      </main>
    </div>

    <!-- ===================================== -->
    <!-- ADD / EDIT MODAL                      -->
    <!-- ===================================== -->
    <transition name="fade">
      <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/70 backdrop-blur-sm" @click="showModal = false"></div>
        <transition name="scale-in">
          <div v-if="showModal" class="relative rounded-2xl p-6 w-full max-w-md shadow-2xl"
            style="background: #141e33; border: 1px solid rgba(255,255,255,0.1);">
            <!-- Modal header -->
            <div class="flex items-center justify-between mb-6">
              <h3 class="text-white text-lg font-bold">
                {{ editingWorker ? 'Редактировать работника' : 'Добавить работника' }}
              </h3>
              <button @click="showModal = false"
                class="p-1.5 rounded-lg text-gray-500 hover:text-gray-300 hover:bg-white/5 transition">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <form @submit.prevent="saveWorker" class="space-y-4">
              <!-- Name -->
              <div>
                <label class="block text-gray-500 text-xs font-semibold mb-1.5 uppercase tracking-wider">
                  Имя <span v-if="!editingWorker" class="text-red-400 normal-case font-normal">*</span>
                </label>
                <input v-model="form.name" type="text" :required="!editingWorker" placeholder="Введите имя работника"
                  class="w-full px-4 py-3 rounded-xl text-white text-sm placeholder-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition"
                  style="background: #1e2d4a; border: 1px solid rgba(255,255,255,0.1);" />
              </div>

              <!-- Phone -->
              <div>
                <label class="block text-gray-500 text-xs font-semibold mb-1.5 uppercase tracking-wider">
                  Телефон <span v-if="!editingWorker" class="text-red-400 normal-case font-normal">*</span>
                </label>
                <div class="flex items-center rounded-xl overflow-hidden"
                  style="background: #1e2d4a; border: 1px solid rgba(255,255,255,0.1);">
                  <span class="px-3 py-3 text-gray-500 text-sm font-mono flex-shrink-0"
                    style="border-right: 1px solid rgba(255,255,255,0.1);">+998</span>
                  <input v-model="form.phone" type="tel" maxlength="9" :required="!editingWorker"
                    placeholder="901234567"
                    class="flex-1 px-3 py-3 text-white text-sm placeholder-gray-600 outline-none font-mono"
                    style="background: transparent;" />
                </div>
              </div>

              <!-- Role -->
              <div>
                <label class="block text-gray-500 text-xs font-semibold mb-1.5 uppercase tracking-wider">Роль</label>
                <select v-model="form.role"
                  class="w-full px-4 py-3 rounded-xl text-white text-sm outline-none focus:ring-2 focus:ring-blue-500 transition"
                  style="background: #1e2d4a; border: 1px solid rgba(255,255,255,0.1);">
                  <option value="pickup">Пункт выдачи</option>
                  <option value="nurse">Медсестра</option>
                  <option value="manager">Менеджер</option>
                </select>
              </div>

              <!-- Password -->
              <div>
                <label class="block text-gray-500 text-xs font-semibold mb-1.5 uppercase tracking-wider">
                  Пароль
                  <span v-if="!editingWorker" class="text-red-400 normal-case font-normal"> *</span>
                  <span v-else class="text-gray-600 normal-case font-normal ml-1">(пусто = не менять)</span>
                </label>
                <input v-model="form.password" type="password" :required="!editingWorker" minlength="6"
                  placeholder="Минимум 6 символов"
                  class="w-full px-4 py-3 rounded-xl text-white text-sm placeholder-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition"
                  style="background: #1e2d4a; border: 1px solid rgba(255,255,255,0.1);" />
              </div>

              <!-- Error -->
              <div v-if="formError"
                class="flex items-center gap-2 px-4 py-3 rounded-xl text-red-400 text-sm"
                style="background: rgba(239,68,68,0.08); border: 1px solid rgba(239,68,68,0.2);">
                <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                  stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                {{ formError }}
              </div>

              <!-- Buttons -->
              <div class="flex gap-3 pt-1">
                <button type="button" @click="showModal = false"
                  class="flex-1 px-4 py-3 rounded-xl text-gray-400 text-sm font-medium hover:text-white transition"
                  style="border: 1px solid rgba(255,255,255,0.1);">
                  Отмена
                </button>
                <button type="submit" :disabled="saving"
                  class="flex-1 px-4 py-3 rounded-xl bg-blue-600 hover:bg-blue-700 text-white text-sm font-semibold transition disabled:opacity-50 flex items-center justify-center gap-2">
                  <svg v-if="saving" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
                  </svg>
                  {{ saving ? 'Сохранение...' : (editingWorker ? 'Сохранить' : 'Добавить') }}
                </button>
              </div>
            </form>
          </div>
        </transition>
      </div>
    </transition>

  </div>
</template>

<script setup>
import { ref, computed, reactive, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore, api } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const lang = ref('ru')
const tab = ref('workers')

// ---- Nav items ----
const navItems = [
  {
    key: 'workers',
    label: 'Работники',
    icon: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2" class="w-4 h-4">
      <path stroke-linecap="round" stroke-linejoin="round"
        d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z"/>
    </svg>`,
  },
  {
    key: 'stats',
    label: 'Статистика',
    icon: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2" class="w-4 h-4">
      <path stroke-linecap="round" stroke-linejoin="round"
        d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
    </svg>`,
  },
  {
    key: 'history',
    label: 'История',
    icon: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2" class="w-4 h-4">
      <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
    </svg>`,
  },
]

// ---- Workers list state ----
const workers = ref([])
const loading = ref(false)
const searchQuery = ref('')
const roleFilter = ref('')

// ---- Inline stats (expand row) ----
const expandedWorkerId = ref(null)
const workerStats = ref(null)
const statsLoading = ref(false)

// ---- All-worker stats tab ----
const allWorkerStats = ref([])
const allStatsLoading = ref(false)
const statsPeriod = ref('today')
const periods = [
  { key: 'today', label: 'Сегодня' },
  { key: 'week', label: 'Неделя' },
  { key: 'month', label: 'Месяц' },
]

// ---- Modal ----
const showModal = ref(false)
const editingWorker = ref(null)
const form = reactive({ name: '', phone: '', password: '', role: 'pickup' })
const formError = ref('')
const saving = ref(false)

// ---- Computed ----
const filteredWorkers = computed(() =>
  workers.value.filter(w => {
    const q = searchQuery.value.toLowerCase()
    const matchSearch = !q ||
      (w.name || '').toLowerCase().includes(q) ||
      (w.phone || '').includes(q)
    const matchRole = !roleFilter.value || w.role === roleFilter.value
    return matchSearch && matchRole
  })
)

// ---- Helpers ----
function roleLabel(role) {
  return { pickup: 'Пункт выдачи', nurse: 'Медсестра', manager: 'Менеджер' }[role] || role
}

function roleColor(role) {
  return { pickup: '#2563eb', nurse: '#059669', manager: '#7c3aed' }[role] || '#6b7280'
}

function roleBadgeClass(role) {
  return {
    pickup: 'bg-blue-500/15 text-blue-400',
    nurse: 'bg-emerald-500/15 text-emerald-400',
    manager: 'bg-purple-500/15 text-purple-400',
  }[role] || 'bg-gray-500/15 text-gray-400'
}

function formatPrice(n) {
  return Number(n || 0).toLocaleString('ru-RU')
}

function getActivityWidth(ws) {
  const maxOrders = Math.max(1, ...allWorkerStats.value.map(x => x.stats[statsPeriod.value]?.orders || 0))
  return Math.round(((ws.stats[statsPeriod.value]?.orders || 0) / maxOrders) * 100)
}

// ---- API calls ----
async function loadWorkers() {
  loading.value = true
  try {
    const res = await api.get('/admin/workers')
    workers.value = res.data || []
  } catch {
    workers.value = []
  } finally {
    loading.value = false
  }
}

async function toggleStats(w) {
  if (expandedWorkerId.value === w.id) {
    expandedWorkerId.value = null
    workerStats.value = null
    return
  }
  expandedWorkerId.value = w.id
  workerStats.value = null
  statsLoading.value = true
  try {
    const res = await api.get(`/admin/workers/${w.id}/stats`)
    workerStats.value = res.data
  } catch {
    workerStats.value = {
      today: { orders: 0, items: 0, revenue: 0 },
      week: { orders: 0, items: 0, revenue: 0 },
      month: { orders: 0, items: 0, revenue: 0 },
    }
  } finally {
    statsLoading.value = false
  }
}

async function loadAllStats() {
  allStatsLoading.value = true
  allWorkerStats.value = []
  try {
    const promises = workers.value.map(w =>
      api.get(`/admin/workers/${w.id}/stats`)
        .then(res => ({ worker: w, stats: res.data }))
        .catch(() => ({
          worker: w,
          stats: {
            today: { orders: 0, items: 0, revenue: 0 },
            week: { orders: 0, items: 0, revenue: 0 },
            month: { orders: 0, items: 0, revenue: 0 },
          },
        }))
    )
    allWorkerStats.value = await Promise.all(promises)
  } finally {
    allStatsLoading.value = false
  }
}

function openAddModal() {
  editingWorker.value = null
  form.name = ''
  form.phone = ''
  form.password = ''
  form.role = 'pickup'
  formError.value = ''
  showModal.value = true
}

function openEditModal(w) {
  editingWorker.value = w
  form.name = w.name || ''
  form.phone = (w.phone || '').replace(/^998/, '')
  form.password = ''
  form.role = w.role || 'pickup'
  formError.value = ''
  showModal.value = true
}

async function saveWorker() {
  formError.value = ''
  saving.value = true
  try {
    if (editingWorker.value) {
      const payload = { role: form.role }
      if (form.name) payload.name = form.name
      if (form.phone) payload.phone = '998' + form.phone
      if (form.password) payload.password = form.password
      await api.put(`/admin/workers/${editingWorker.value.id}`, payload)
    } else {
      await api.post('/admin/workers', {
        name: form.name,
        phone: '998' + form.phone,
        password: form.password,
        role: form.role,
      })
    }
    showModal.value = false
    await loadWorkers()
    if (tab.value === 'stats') await loadAllStats()
  } catch (e) {
    formError.value = e.response?.data?.error || 'Ошибка при сохранении'
  } finally {
    saving.value = false
  }
}

async function deleteWorker(id) {
  if (!confirm('Удалить этого работника?')) return
  try {
    await api.delete(`/admin/workers/${id}`)
    if (expandedWorkerId.value === id) {
      expandedWorkerId.value = null
      workerStats.value = null
    }
    await loadWorkers()
    if (tab.value === 'stats') await loadAllStats()
  } catch (e) {
    alert(e.response?.data?.error || 'Ошибка при удалении')
  }
}

function logout() {
  authStore.adminLogout()
  router.push('/admin/login')
}

// Auto-load stats when switching to stats tab
watch(tab, async newTab => {
  if (newTab === 'stats' && allWorkerStats.value.length === 0 && workers.value.length > 0) {
    await loadAllStats()
  }
})

onMounted(async () => {
  if (!authStore.adminToken) {
    router.push('/admin/login')
    return
  }
  await loadWorkers()
})
</script>

<style scoped>
.workers-panel {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.scale-in-enter-active {
  transition: all 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.scale-in-enter-from {
  opacity: 0;
  transform: scale(0.95) translateY(-8px);
}

.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.25s ease;
  overflow: hidden;
  max-height: 400px;
}
.slide-down-enter-from,
.slide-down-leave-to {
  max-height: 0;
  opacity: 0;
}

/* Subtle row hover */
.hover\:bg-white\/2:hover {
  background-color: rgba(255, 255, 255, 0.02);
}

/* Dark select styling */
select option {
  background: #1e2d4a;
  color: white;
}
</style>
