<template>
  <footer class="bg-slate-900 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-12">
      <div class="flex flex-col md:flex-row justify-between gap-8">
        <!-- Brand -->
        <div class="max-w-xs">
          <div class="flex items-center gap-3 mb-4">
            <div class="w-10 h-10 rounded-xl overflow-hidden bg-white flex items-center justify-center shrink-0">
              <img src="/images/patients/Jalilov.jpg" alt="Doctor Jalilov" class="w-full h-full object-contain" />
            </div>
            <div>
              <div class="font-bold text-white">Doctor Jalilov</div>
              <div class="text-xs text-white/40">{{ t.footer_trich }}</div>
            </div>
          </div>
          <p class="text-white/40 text-sm leading-relaxed">{{ t.footer_desc }}</p>
        </div>

        <!-- Contacts -->
        <div class="flex flex-col sm:flex-row gap-8">
          <div>
            <div class="text-xs font-semibold tracking-widest uppercase text-white/50 mb-4">{{ t.footer_contacts }}</div>
            <div class="space-y-2">
              <a :href="`tel:${t.contacts_phone2.replace(/\s/g,'')}`"
                 class="flex items-center gap-2 text-sm text-white/50 hover:text-white transition-colors duration-200">
                <svg class="w-4 h-4 text-brand-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z"/>
                </svg>
                {{ t.contacts_phone2 }}
              </a>
              <div class="flex items-center gap-2 text-sm text-white/40">
                <svg class="w-4 h-4 text-brand-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"/>
                </svg>
                {{ t.contacts_address }}
              </div>
            </div>
          </div>

          <!-- Hours -->
          <div>
            <div class="text-xs font-semibold tracking-widest uppercase text-white/50 mb-4">{{ t.contacts_hours_label }}</div>
            <div class="space-y-2 text-sm text-white/40">
              <div class="flex justify-between gap-6">
                <span>{{ t.footer_weekdays }}</span>
                <span class="text-white/60 font-medium">08:00 — 17:00</span>
              </div>
              <div class="flex justify-between gap-6">
                <span>{{ t.footer_weekends }}</span>
                <span class="text-white/60 font-medium">{{ t.footer_holiday }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Legal + payment -->
      <div class="border-t border-white/10 mt-8 pt-6 flex flex-col lg:flex-row lg:items-center lg:justify-between gap-5">
        <div class="flex flex-wrap gap-x-5 gap-y-2 text-sm text-white/40">
          <router-link to="/offer" class="hover:text-white transition-colors duration-200">{{ t.footer_offer }}</router-link>
          <router-link to="/terms" class="hover:text-white transition-colors duration-200">{{ t.footer_terms }}</router-link>
          <router-link to="/privacy" class="hover:text-white transition-colors duration-200">{{ t.footer_privacy }}</router-link>
          <a :href="`mailto:${org.email}`" class="hover:text-white transition-colors duration-200">{{ org.email }}</a>
        </div>
        <PaymentLogos />
      </div>

      <div class="border-t border-white/10 mt-6 pt-6 flex flex-col sm:flex-row justify-between items-center gap-3">
        <p class="text-xs text-white/30">{{ t.footer_copy }}</p>
        <router-link to="/admin/login" class="text-xs text-white/20 hover:text-white/40 transition-colors duration-300">
          {{ t.footer_admin }}
        </router-link>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'
import { useLangStore } from '../stores/lang'
import PaymentLogos from './PaymentLogos.vue'
import { ORG } from '../config/org'

const langStore = useLangStore()
const t = computed(() => langStore.t)
const org = ORG
</script>
