<template>
  <section id="doctor" class="py-24 bg-gradient-to-b from-surface-warm to-white relative overflow-hidden">
    <div class="absolute top-0 right-0 w-[600px] h-[600px] bg-brand-50 rounded-full -translate-y-1/2 translate-x-1/3 blur-3xl opacity-40 animate-pulse-soft"></div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid md:grid-cols-2 gap-16 lg:gap-24 items-center">
        <div class="relative reveal-left">
          <div class="absolute -inset-4 bg-gradient-to-br from-brand-100 to-brand-50 rounded-3xl -rotate-2 transition-transform duration-700 hover:rotate-0"></div>
          <div class="absolute -bottom-6 -right-6 w-32 h-32 bg-brand-100 rounded-2xl rotate-6 hidden lg:block animate-float" style="animation-delay: 1s"></div>

          <div class="doctor-tilt-wrap relative max-w-lg mx-auto">
            <div class="doctor-tilt-canvas">
              <div
                v-for="n in 25"
                :key="n"
                :class="`doctor-tracker doctor-tr-${n}`"
              ></div>

              <div class="doctor-tilt-card">
                <img
                  src="/images/patients/2%20landing.jpg"
                  alt="Директор клиники"
                  class="doctor-tilt-image"
                  draggable="false"
                />
                <div class="doctor-tilt-overlay"></div>
                <p class="doctor-tilt-prompt">HOVER</p>
                <p class="doctor-tilt-title">JALILOV</p>
              </div>
            </div>
          </div>

          <div class="absolute -bottom-4 -left-4 md:-bottom-6 md:-left-6 bg-white rounded-2xl shadow-xl shadow-stone-900/10 p-4 border border-stone-100 hover:shadow-2xl hover:-translate-y-1 transition-all duration-500">
            <div class="text-2xl font-serif text-brand-700">40<span class="text-brand-300">+</span></div>
            <div class="text-xs font-medium text-stone-500 mt-0.5">{{ t.hero_years }}</div>
          </div>
        </div>

        <div class="reveal-right">
          <div class="section-label mb-5">
            <span class="w-8 h-px bg-brand-400"></span>
            {{ t.doctor_label }}
          </div>

          <h2 class="heading-lg mb-6">
            {{ t.doctor_title1 }}
            <span class="text-brand-600"> {{ t.doctor_title2 }}</span>
          </h2>

          <p class="text-stone-500 text-lg leading-relaxed mb-5">{{ t.doctor_desc1 }}</p>
          <p class="text-stone-500 text-lg leading-relaxed mb-10">{{ t.doctor_desc2 }}</p>

          <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
            <div class="bg-surface-warm rounded-xl p-4 border border-stone-100/80 hover:border-brand-200/50 hover:bg-brand-50/30 hover:-translate-y-1 hover:shadow-lg hover:shadow-brand-100/50 transition-all duration-500 group/feature">
              <div class="w-10 h-10 bg-brand-100 rounded-lg flex items-center justify-center mb-3 group-hover/feature:bg-brand-200/70 group-hover/feature:scale-110 transition-all duration-500">
                <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">{{ t.doctor_individual }}</div>
              <div class="text-xs text-stone-400 mt-0.5">{{ t.doctor_approach }}</div>
            </div>
            <div class="bg-surface-warm rounded-xl p-4 border border-stone-100/80 hover:border-amber-200/50 hover:bg-amber-50/30 hover:-translate-y-1 hover:shadow-lg hover:shadow-amber-100/50 transition-all duration-500 group/feature">
              <div class="w-10 h-10 bg-amber-50 rounded-lg flex items-center justify-center mb-3 group-hover/feature:bg-amber-100/70 group-hover/feature:scale-110 transition-all duration-500">
                <svg class="w-5 h-5 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9.75 3.104v5.714a2.25 2.25 0 01-.659 1.591L5 14.5M9.75 3.104c-.251.023-.501.05-.75.082m.75-.082a24.301 24.301 0 014.5 0m0 0v5.714c0 .597.237 1.17.659 1.591L19.8 15.3M14.25 3.104c.251.023.501.05.75.082M19.8 15.3l-1.57.393A9.065 9.065 0 0112 15a9.065 9.065 0 00-6.23.693L5 14.5m14.8.8l1.402 1.402c1.232 1.232.65 3.318-1.067 3.611A48.309 48.309 0 0112 21c-2.773 0-5.491-.235-8.135-.687-1.718-.293-2.3-2.379-1.067-3.61L5 14.5" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">{{ t.doctor_modern }}</div>
              <div class="text-xs text-stone-400 mt-0.5">{{ t.doctor_methods }}</div>
            </div>
            <div class="bg-surface-warm rounded-xl p-4 border border-stone-100/80 hover:border-brand-200/50 hover:bg-brand-50/30 hover:-translate-y-1 hover:shadow-lg hover:shadow-brand-100/50 transition-all duration-500 group/feature">
              <div class="w-10 h-10 bg-brand-100 rounded-lg flex items-center justify-center mb-3 group-hover/feature:bg-brand-200/70 group-hover/feature:scale-110 transition-all duration-500">
                <svg class="w-5 h-5 text-brand-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.8">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285z" />
                </svg>
              </div>
              <div class="text-sm font-semibold text-stone-800">{{ t.doctor_proven }}</div>
              <div class="text-xs text-stone-400 mt-0.5">{{ t.doctor_meds }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { computed } from 'vue'
import { useLangStore } from '../stores/lang'

const langStore = useLangStore()
const t = computed(() => langStore.t)
</script>

<style scoped>
.doctor-tilt-wrap {
  width: min(100%, 420px);
  aspect-ratio: 5 / 6;
}

.doctor-tilt-canvas {
  position: relative;
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: repeat(5, 1fr);
  border-radius: 1rem;
  overflow: hidden;
}

.doctor-tracker {
  z-index: 3;
}

.doctor-tilt-card {
  position: absolute;
  inset: 0;
  z-index: 1;
  border-radius: 1rem;
  overflow: hidden;
  box-shadow: 0 24px 40px -20px rgba(2, 6, 23, 0.35);
  transition: transform 180ms ease, filter 180ms ease;
  transform: perspective(900px) rotateX(0deg) rotateY(0deg);
}

.doctor-tilt-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  user-select: none;
}

.doctor-tilt-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(15, 23, 42, 0.08) 40%, rgba(15, 23, 42, 0.5) 100%);
}

.doctor-tilt-title {
  position: absolute;
  left: 12px;
  bottom: 12px;
  font-size: 1.1rem;
  font-weight: 700;
  color: #fff;
  letter-spacing: 0.08em;
  opacity: 0;
  transition: opacity 180ms ease;
  z-index: 2;
}

.doctor-tilt-prompt {
  position: absolute;
  left: 12px;
  top: 12px;
  z-index: 2;
  padding: 4px 8px;
  border-radius: 9999px;
  background: rgba(255, 255, 255, 0.85);
  color: #1d4ed8;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  transition: opacity 180ms ease;
}

.doctor-tilt-canvas:hover .doctor-tilt-card {
  filter: brightness(1.06);
}

.doctor-tracker:hover ~ .doctor-tilt-card .doctor-tilt-title {
  opacity: 1;
}

.doctor-tracker:hover ~ .doctor-tilt-card .doctor-tilt-prompt {
  opacity: 0;
}

.doctor-tr-1:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(13deg) rotateY(-9deg); }
.doctor-tr-2:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(13deg) rotateY(-4deg); }
.doctor-tr-3:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(13deg) rotateY(0deg); }
.doctor-tr-4:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(13deg) rotateY(4deg); }
.doctor-tr-5:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(13deg) rotateY(9deg); }
.doctor-tr-6:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(7deg) rotateY(-9deg); }
.doctor-tr-7:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(7deg) rotateY(-4deg); }
.doctor-tr-8:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(7deg) rotateY(0deg); }
.doctor-tr-9:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(7deg) rotateY(4deg); }
.doctor-tr-10:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(7deg) rotateY(9deg); }
.doctor-tr-11:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(0deg) rotateY(-9deg); }
.doctor-tr-12:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(0deg) rotateY(-4deg); }
.doctor-tr-13:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(0deg) rotateY(0deg); }
.doctor-tr-14:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(0deg) rotateY(4deg); }
.doctor-tr-15:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(0deg) rotateY(9deg); }
.doctor-tr-16:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-7deg) rotateY(-9deg); }
.doctor-tr-17:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-7deg) rotateY(-4deg); }
.doctor-tr-18:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-7deg) rotateY(0deg); }
.doctor-tr-19:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-7deg) rotateY(4deg); }
.doctor-tr-20:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-7deg) rotateY(9deg); }
.doctor-tr-21:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-13deg) rotateY(-9deg); }
.doctor-tr-22:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-13deg) rotateY(-4deg); }
.doctor-tr-23:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-13deg) rotateY(0deg); }
.doctor-tr-24:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-13deg) rotateY(4deg); }
.doctor-tr-25:hover ~ .doctor-tilt-card { transform: perspective(900px) rotateX(-13deg) rotateY(9deg); }

@media (max-width: 768px) {
  .doctor-tilt-wrap {
    width: min(100%, 360px);
  }

  .doctor-tracker {
    display: none;
  }

  .doctor-tilt-card {
    transform: none !important;
  }

  .doctor-tilt-title {
    opacity: 1;
  }

  .doctor-tilt-prompt {
    display: none;
  }
}
</style>
