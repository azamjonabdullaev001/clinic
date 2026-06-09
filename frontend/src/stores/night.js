import { ref, watch } from 'vue'

// Shared (singleton) night-mode state for all worker panels.
const night = ref(localStorage.getItem('nightMode') === '1')

watch(night, (v) => {
  localStorage.setItem('nightMode', v ? '1' : '0')
})

export function useNight() {
  function toggle() {
    night.value = !night.value
  }
  return { night, toggle }
}
