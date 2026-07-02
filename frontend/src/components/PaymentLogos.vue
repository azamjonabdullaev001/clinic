<template>
  <!--
    Логотипы платёжных систем (ЛПС и МПС) + значок защищённого соединения.
    Требование банка для приёма онлайн-оплаты картами.
    variant="light" — для тёмного фона (футер), "dark" — для светлого.
  -->
  <div class="flex flex-wrap items-center gap-2.5">
    <!-- Visa -->
    <span class="inline-flex items-center justify-center h-8 min-w-[54px] px-3 rounded-md bg-white border border-stone-200 shadow-sm">
      <span class="font-black italic tracking-tight text-[15px] text-[#1a1f71]">VISA</span>
    </span>

    <!-- Mastercard -->
    <span class="inline-flex items-center justify-center h-8 min-w-[54px] px-3 rounded-md bg-white border border-stone-200 shadow-sm">
      <span class="relative inline-flex items-center">
        <span class="w-4 h-4 rounded-full bg-[#eb001b]"></span>
        <span class="w-4 h-4 rounded-full bg-[#f79e1b] -ml-1.5 mix-blend-multiply"></span>
      </span>
    </span>

    <!-- Humo -->
    <span class="inline-flex items-center justify-center h-8 min-w-[54px] px-3 rounded-md bg-white border border-stone-200 shadow-sm">
      <span class="font-extrabold text-[13px] tracking-tight text-[#00a3e0]">HUMO</span>
    </span>

    <!-- Uzcard -->
    <span class="inline-flex items-center justify-center h-8 min-w-[54px] px-3 rounded-md bg-white border border-stone-200 shadow-sm">
      <span class="font-extrabold text-[13px] tracking-tight text-[#0d5eaf]">UzCard</span>
    </span>

    <!-- SSL / безопасное соединение -->
    <span class="inline-flex items-center gap-1.5 h-8 px-3 rounded-md bg-emerald-50 border border-emerald-200">
      <svg class="w-3.5 h-3.5 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
        <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 0h10.5a2.25 2.25 0 012.25 2.25v6a2.25 2.25 0 01-2.25 2.25H6.75A2.25 2.25 0 014.5 18.75v-6a2.25 2.25 0 012.25-2.25z" />
      </svg>
      <span class="text-[11px] font-semibold text-emerald-700">SSL</span>
    </span>
  </div>
</template>

<script setup>
// Чисто презентационный компонент, пропсов нет.
</script>
