<template>
  <div class="w-full">
    <div v-if="pts.length > 1" :key="animKey" class="lc-reveal" style="height:230px">
      <svg :viewBox="`0 0 ${W} ${H}`" preserveAspectRatio="none" class="w-full h-full block">
        <defs>
          <linearGradient :id="gradId" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" :stop-color="color" stop-opacity="0.35" />
            <stop offset="100%" :stop-color="color" stop-opacity="0.02" />
          </linearGradient>
        </defs>
        <line v-for="g in 3" :key="'g' + g" :x1="0" :x2="W" :y1="(H / 4) * g" :y2="(H / 4) * g"
          stroke="rgba(148,163,184,0.18)" stroke-width="1" vector-effect="non-scaling-stroke" />
        <path :d="areaPath" :fill="`url(#${gradId})`" />
        <path :d="linePath" fill="none" :stroke="color" stroke-width="2.5" stroke-linejoin="round"
          stroke-linecap="round" vector-effect="non-scaling-stroke" />
      </svg>
    </div>
    <div v-else class="flex items-center justify-center text-sm text-gray-400" style="height:230px">Нет данных за период</div>
    <div v-if="pts.length > 1" class="flex justify-between text-[10px] text-gray-400 mt-1.5 px-0.5">
      <span v-for="(l, i) in xLabels" :key="i">{{ l }}</span>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

const props = defineProps({
  points: { type: Array, default: () => [] },
  color: { type: String, default: '#6366f1' },
})

const W = 600
const H = 220
const gradId = 'lcgrad-' + Math.random().toString(36).slice(2, 8)

const pts = computed(() => props.points || [])
const animKey = ref(0)
watch(() => props.points, () => { animKey.value++ }, { deep: true })

const max = computed(() => Math.max(1, ...pts.value.map(p => p.revenue || 0)))

function coords() {
  const n = pts.value.length
  return pts.value.map((p, i) => {
    const x = n > 1 ? (i / (n - 1)) * W : 0
    const y = H - 6 - ((p.revenue || 0) / max.value) * (H - 16)
    return [x, y]
  })
}

const linePath = computed(() => {
  const c = coords()
  if (!c.length) return ''
  return c.map(([x, y], i) => `${i === 0 ? 'M' : 'L'} ${x.toFixed(1)} ${y.toFixed(1)}`).join(' ')
})

const areaPath = computed(() => {
  const c = coords()
  if (!c.length) return ''
  return c.map(([x, y], i) => `${i === 0 ? 'M' : 'L'} ${x.toFixed(1)} ${y.toFixed(1)}`).join(' ') +
    ` L ${W} ${H} L 0 ${H} Z`
})

const xLabels = computed(() => {
  const n = pts.value.length
  if (n <= 1) return []
  const step = Math.max(1, Math.ceil(n / 6))
  const out = []
  for (let i = 0; i < n; i += step) out.push(pts.value[i].label)
  if (out[out.length - 1] !== pts.value[n - 1].label) out.push(pts.value[n - 1].label)
  return out
})
</script>

<style scoped>
.lc-reveal {
  animation: lcReveal 0.8s ease-out;
}
@keyframes lcReveal {
  from { clip-path: inset(0 100% 0 0); }
  to { clip-path: inset(0 0 0 0); }
}
</style>
