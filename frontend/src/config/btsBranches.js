// Known BTS Cargo pickup points in Uzbekistan.
// Keys are search terms (lowercase) found in delivery addresses.
// lat/lng are the branch coordinates for OSRM routing.
// Verify exact addresses with BTS before going live.
export const BTS_BRANCHES = [
  {
    keys: ['namangan', 'наманган', 'нamanган'],
    city: 'Наманган',
    name: 'БТС Карго — Наманган',
    address: 'Наманган, ул. Мустақиллик (уточните у БТС)',
    lat: 40.9983,
    lng: 71.6726,
  },
  {
    keys: ['andijan', 'андижан', 'andijon'],
    city: 'Андижан',
    name: 'БТС Карго — Андижан',
    address: 'Андижан, ул. Навои (уточните у БТС)',
    lat: 40.7829,
    lng: 72.3440,
  },
  {
    keys: ['tashkent', 'ташкент', 'toshkent'],
    city: 'Ташкент',
    name: 'БТС Карго — Ташкент',
    address: 'Ташкент, ул. Сергели (уточните у БТС)',
    lat: 41.2995,
    lng: 69.2401,
  },
  {
    keys: ['fergana', 'фергана', 'farg\'ona', 'farghona'],
    city: 'Фергана',
    name: 'БТС Карго — Фергана',
    address: 'Фергана, ул. Мустақиллик (уточните у БТС)',
    lat: 40.3736,
    lng: 71.7975,
  },
  {
    keys: ['kokand', 'коканд', 'qo\'qon'],
    city: 'Коканд',
    name: 'БТС Карго — Коканд',
    address: 'Коканд, ул. Навои (уточните у БТС)',
    lat: 40.5281,
    lng: 70.9419,
  },
  {
    keys: ['margilan', 'марғилон', 'margilon'],
    city: 'Маргилан',
    name: 'БТС Карго — Маргилан',
    address: 'Маргилан (уточните у БТС)',
    lat: 40.4726,
    lng: 71.7270,
  },
  {
    keys: ['samarkand', 'самарканд', 'samarqand'],
    city: 'Самарканд',
    name: 'БТС Карго — Самарканд',
    address: 'Самарканд (уточните у БТС)',
    lat: 39.6542,
    lng: 66.9597,
  },
  {
    keys: ['bukhara', 'бухара', 'buxoro'],
    city: 'Бухара',
    name: 'БТС Карго — Бухара',
    address: 'Бухара (уточните у БТС)',
    lat: 39.7747,
    lng: 64.4286,
  },
  {
    keys: ['urgench', 'ургенч', 'urganch'],
    city: 'Ургенч',
    name: 'БТС Карго — Ургенч',
    address: 'Ургенч (уточните у БТС)',
    lat: 41.5477,
    lng: 60.6342,
  },
  {
    keys: ['nukus', 'нукус'],
    city: 'Нукус',
    name: 'БТС Карго — Нукус',
    address: 'Нукус (уточните у БТС)',
    lat: 42.4526,
    lng: 59.6107,
  },
  {
    keys: ['termez', 'термез', 'termiz'],
    city: 'Термез',
    name: 'БТС Карго — Термез',
    address: 'Термез (уточните у БТС)',
    lat: 37.2242,
    lng: 67.2783,
  },
  {
    keys: ['karshi', 'карши', 'qarshi'],
    city: 'Карши',
    name: 'БТС Карго — Карши',
    address: 'Карши (уточните у БТС)',
    lat: 38.8600,
    lng: 65.7911,
  },
]

/**
 * Find a BTS branch by scanning the delivery address text for known city names.
 * Returns the branch object or null.
 */
export function detectBtsBranch(deliveryAddress) {
  if (!deliveryAddress) return null
  const lower = deliveryAddress.toLowerCase()
  for (const branch of BTS_BRANCHES) {
    if (branch.keys.some(k => lower.includes(k))) {
      return branch
    }
  }
  return null
}
