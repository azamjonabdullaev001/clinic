import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

// Create axios instance OUTSIDE of Pinia store to avoid reactive proxy wrapping
const api = axios.create({ baseURL: '/api' })

api.interceptors.request.use(config => {
  const adminToken = localStorage.getItem('adminToken') || ''
  const userToken = localStorage.getItem('userToken') || ''
  if (config.url?.startsWith('/admin') && adminToken) {
    config.headers.Authorization = `Bearer ${adminToken}`
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

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => !!adminToken.value)

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
    adminToken.value = res.data.token
    admin.value = res.data.admin
    localStorage.setItem('adminToken', res.data.token)
    localStorage.setItem('admin', JSON.stringify(res.data.admin))
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

  return {
    token, user, adminToken, admin,
    isLoggedIn, isAdmin,
    register, login, adminLogin, logout, adminLogout
  }
})
