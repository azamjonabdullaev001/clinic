<template>
  <div class="w-full select-none">
    <div
      v-if="pts.length > 1"
      ref="wrap"
      class="relative"
      :style="{ height: height + 'px' }"
      @pointermove="onMove"
      @pointerleave="hover = null"
    >
      <!-- Y axis value labels -->
      <div class="absolute inset-y-0 left-0 w-full pointer-events-none">
        <div
          v-for="(g, i) in gridFracs"
          :key="'yl' + i"
          class="absolute left-0 -translate-y-1/2 text-[10px] text-gray-400 pl-0.5"
          :style="{ top: (g * 100) + '%' }"
        >{{ fmtCompact(max * (1 - g)) }}</div>
      </div>

      <svg :viewBox="`0 0 ${W} ${H}`" preserveAspectRatio="none" class="w-full h-full block overflow-visible">
        <defs>
          <linearGradient :id="gradId" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" :stop-color="color" stop-opacity="0.32" />
            <stop offset="100%" :stop-color="color" stop-opacity="0.01" />
          </linearGradient>
        </defs>
        <!-- grid -->
        <line
          v-for="(g, i) in gridFracs" :key="'g' + i"
          :x1="0" :x2="W" :y1="g * H" :y2="g * H"
          stroke="rgba(148,163,184,0.18)" stroke-width="1" vector-effect="non-scaling-stroke"
        />
        <!-- animated area + line -->
        <g :key="animKey" class="lc-reveal">
          <path :d="areaPath" :fill="`url(#${gradId})`" />
          <path
            :d="linePath" fill="none" :stroke="color" stroke-width="2.5"
            stroke-linejoin="round" stroke-linecap="round" vector-effect="non-scaling-stroke"
          />
        </g>
      </svg>

      <!-- hover guide + dot (HTML overlay, undistorted) -->
      <template v-if="hover !== null">
        <div
          class="absolute top-0 bottom-0 w-px bg-gray-300/70 pointer-events-none"
          :style="{ left: (coordsFrac[hover].x * 100) + '%' }"
        ></div>
        <div
          class="absolute w-3 h-3 rounded-full border-2 bg-white pointer-events-none lc-dot"
          :style="{
            left: (coordsFrac[hover].x * 100) + '%',
            top: (coordsFrac[hover].y * 100) + '%',
            borderColor: color,
            transform: 'translate(-50%, -50%)',
          }"
        ></div>
        <div
          class="absolute z-10 pointer-events-none bg-white/95 backdrop-blur shadow-lg rounded-lg border border-gray-100 px-3 py-2 text-xs whitespace-nowrap"
          :style="tooltipStyle"
        >
          <div class="text-gray-400 mb-0.5">{{ pts[hover].label }}</div>
          <div class="flex items-center gap-1.5 font-semibold" :style="{ color }">
            <span class="inline-block w-2 h-2 rounded-full" :style="{ background: color }"></span>
            {{ fmtFull(pts[hover].revenue || 0) }} {{ unit }}
          </div>
          <div v-if="hasOrders" class="text-gray-500 mt-0.5">{{ orderLabel }}: {{ pts[hover].orders || 0 }}</div>
        </div>
      </template>
    </div>
    <div v-else class="flex items-center justify-center text-sm text-gray-400" :style="{ height: height + 'px' }">
      {{ emptyText }}
    </div>

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
  height: { type: Number, default: 240 },
  unit: { type: String, default: 'сўм' },
  orderLabel: { type: String, default: 'Заказов' },
  emptyText: { type: String, default: 'Нет данных за период' },
})

const W = 600
const H = 240
const PAD_T = 14
const PAD_B = 14
const gradId = 'lcgrad-' + Math.random().toString(36).slice(2, 8)
const gridFracs = [0, 0.25, 0.5, 0.75, 1]

const wrap = ref(null)
const hover = ref(null)

const pts = computed(() => props.points || [])
const hasOrders = computed(() => pts.value.some(p => p && p.orders != null))
const animKey = ref(0)
watch(() => props.points, () => { animKey.value++; hover.value = null }, { deep: true })

const max = computed(() => Math.max(1, ...pts.value.map(p => p.revenue || 0)))

// Fractional coordinates in [0,1] for the HTML overlay; multiply by W/H for the SVG.
const coordsFrac = computed(() => {
  const n = pts.value.length
  const usable = H - PAD_T - PAD_B
  return pts.value.map((p, i) => {
    const x = n > 1 ? i / (n - 1) : 0
    const yPx = PAD_T + (1 - (p.revenue || 0) / max.value) * usable
    return { x, y: yPx / H }
  })
})

function svgCoords() {
  return coordsFrac.value.map(c => [c.x * W, c.y * H])
}

// Smooth curve via Catmull-Rom -> cubic bezier.
function buildSmooth(c) {
  if (c.length < 2) return ''
  const d = [`M ${c[0][0].toFixed(1)} ${c[0][1].toFixed(1)}`]
  for (let i = 0; i < c.length - 1; i++) {
    const p0 = c[i === 0 ? 0 : i - 1]
    const p1 = c[i]
    const p2 = c[i + 1]
    const p3 = c[i + 2] || p2
    const cp1x = p1[0] + (p2[0] - p0[0]) / 6
    const cp1y = p1[1] + (p2[1] - p0[1]) / 6
    const cp2x = p2[0] - (p3[0] - p1[0]) / 6
    const cp2y = p2[1] - (p3[1] - p1[1]) / 6
    d.push(`C ${cp1x.toFixed(1)} ${cp1y.toFixed(1)} ${cp2x.toFixed(1)} ${cp2y.toFixed(1)} ${p2[0].toFixed(1)} ${p2[1].toFixed(1)}`)
  }
  return d.join(' ')
}

const linePath = computed(() => buildSmooth(svgCoords()))
const areaPath = computed(() => {
  const c = svgCoords()
  if (c.length < 2) return ''
  return buildSmooth(c) + ` L ${W} ${H} L 0 ${H} Z`
})

function onMove(e) {
  const el = wrap.value
  if (!el) return
  const rect = el.getBoundingClientRect()
  const frac = Math.min(1, Math.max(0, (e.clientX - rect.left) / rect.width))
  hover.value = Math.round(frac * (pts.value.length - 1))
}

const tooltipStyle = computed(() => {
  if (hover.value === null) return {}
  const x = coordsFrac.value[hover.value].x
  let tx = '-50%'
  if (x < 0.15) tx = '0%'
  else if (x > 0.85) tx = '-100%'
  return {
    left: (x * 100) + '%',
    top: '8px',
    transform: `translateX(${tx})`,
  }
})

function fmtCompact(v) {
  return new Intl.NumberFormat('ru-RU', { notation: 'compact', maximumFractionDigits: 1 }).format(Math.round(v))
}
function fmtFull(v) {
  return new Intl.NumberFormat('ru-RU').format(Math.round(v))
}

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
  animation: lcReveal 0.85s cubic-bezier(0.22, 1, 0.36, 1);
}
@keyframes lcReveal {
  from { clip-path: inset(0 100% 0 0); }
  to { clip-path: inset(0 0 0 0); }
}
.lc-dot {
  animation: lcDot 0.18s ease-out;
}
@keyframes lcDot {
  from { transform: translate(-50%, -50%) scale(0.4); opacity: 0; }
  to { transform: translate(-50%, -50%) scale(1); opacity: 1; }
}
</style>
