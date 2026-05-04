import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

// Create axios instance OUTSIDE of Pinia store to avoid reactive proxy wrapping
const api = axios.create({ baseURL: '/api' })

api.interceptors.request.use(config => {
  const adminToken = localStorage.getItem('adminToken') || ''
  const workerToken = localStorage.getItem('workerToken') || ''
  const userToken = localStorage.getItem('userToken') || ''
  if (config.url?.startsWith('/admin') && adminToken) {
    config.headers.Authorization = `Bearer ${adminToken}`
  } else if (config.url?.startsWith('/pickup') && workerToken) {
    config.headers.Authorization = `Bearer ${workerToken}`
  } else if (userToken) {
    config.headers.Authorization = `Bearer ${userToken}`
  }
  return config
})

export { api }

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('userToken') || '')
  const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))
  const adminToken = ref(localStorage.getItem('adminToken') || '')
  const admin = ref(JSON.parse(localStorage.getItem('admin') || 'null'))
  const workerToken = ref(localStorage.getItem('workerToken') || '')
  const worker = ref(JSON.parse(localStorage.getItem('worker') || 'null'))

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => !!adminToken.value)
  const isWorker = computed(() => !!workerToken.value)

  async function register(data) {
    const res = await api.post('/auth/register', data)
    token.value = res.data.token
    user.value = res.data.user
    localStorage.setItem('userToken', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
    return res.data
  }

  async function login(data) {
    const res = await api.post('/auth/login', data)
    token.value = res.data.token
    user.value = res.data.user
    localStorage.setItem('userToken', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
    return res.data
  }

  async function adminLogin(data) {
    const res = await api.post('/admin/login', data)
    if (res.data.role === 'worker') {
      workerToken.value = res.data.token
      worker.value = res.data.worker
      localStorage.setItem('workerToken', res.data.token)
      localStorage.setItem('worker', JSON.stringify(res.data.worker))
    } else {
      adminToken.value = res.data.token
      admin.value = res.data.admin
      localStorage.setItem('adminToken', res.data.token)
      localStorage.setItem('admin', JSON.stringify(res.data.admin))
    }
    return res.data
  }

  function logout() {
    token.value = ''
    user.value = null
    localStorage.removeItem('userToken')
    localStorage.removeItem('user')
  }

  function adminLogout() {
    adminToken.value = ''
    admin.value = null
    localStorage.removeItem('adminToken')
    localStorage.removeItem('admin')
  }

  function workerLogout() {
    workerToken.value = ''
    worker.value = null
    localStorage.removeItem('workerToken')
    localStorage.removeItem('worker')
  }

  return {
    token, user, adminToken, admin, workerToken, worker,
    isLoggedIn, isAdmin, isWorker,
    register, login, adminLogin, logout, adminLogout, workerLogout
  }
})
