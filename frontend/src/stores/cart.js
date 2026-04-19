import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useCartStore = defineStore('cart', () => {
  const items = ref(JSON.parse(localStorage.getItem('cart') || '[]'))
  const isOpen = ref(false)

  const totalItems = computed(() =>
    items.value.reduce((sum, item) => sum + item.quantity, 0)
  )

  const totalPrice = computed(() =>
    items.value.reduce((sum, item) => {
      const unitPrice = item.unit_type === 'piece' ? item.price_per_pill : item.price_per_pack
      return sum + unitPrice * item.quantity
    }, 0)
  )

  function save() {
    localStorage.setItem('cart', JSON.stringify(items.value))
  }

  function addItem(product) {
    const existing = items.value.find(i => i.product_id === product.id)
    if (existing) {
      existing.quantity++
    } else {
      items.value.push({
        product_id: product.id,
        name: product.name,
        price_per_pill: product.price_per_pill,
        price_per_pack: product.price_per_pack,
        quantity_per_pack: product.quantity_per_pack,
        image_path: product.image_path,
        quantity: 1,
        unit_type: 'pack'
      })
    }
    save()
  }

  function removeItem(productId) {
    items.value = items.value.filter(i => i.product_id !== productId)
    save()
  }

  function updateQuantity(productId, quantity) {
    const item = items.value.find(i => i.product_id === productId)
    if (item) {
      item.quantity = Math.max(1, quantity)
      save()
    }
  }

  function updateUnitType(productId, unitType) {
    const item = items.value.find(i => i.product_id === productId)
    if (item) {
      item.unit_type = unitType
      save()
    }
  }

  function clear() {
    items.value = []
    save()
  }

  function toggle() {
    isOpen.value = !isOpen.value
  }

  return {
    items, isOpen, totalItems, totalPrice,
    addItem, removeItem, updateQuantity, updateUnitType, clear, toggle
  }
})
