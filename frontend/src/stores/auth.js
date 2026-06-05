import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

// Create axios instance OUTSIDE of Pinia store to avoid reactive proxy wrapping
const api = axios.create({ baseURL: '/api' })

// Staff sessions (admin / worker / doctor) are kept in sessionStorage so they are NOT
// cached across browser sessions — closing the tab logs them out and they must sign in
// again (security). Regular landing-page customers stay in localStorage (persistent).
const staff = sessionStorage

api.interceptors.request.use(config => {
  const adminToken = staff.getItem('adminToken') || ''
  const workerToken = staff.getItem('workerToken') || ''
  const doctorToken = staff.getItem('doctorToken') || ''
  const userToken = localStorage.getItem('userToken') || ''
  if (config.url?.startsWith('/admin') && adminToken) {
    config.headers.Authorization = `Bearer ${adminToken}`
  } else if ((config.url?.startsWith('/pickup') || config.url?.startsWith('/nurse') || config.url?.startsWith('/manager')) && workerToken) {
    config.headers.Authorization = `Bearer ${workerToken}`
  } else if (config.url?.startsWith('/doctor') && doctorToken) {
    config.headers.Authorization = `Bearer ${doctorToken}`
  } else if (userToken) {
    config.headers.Authorization = `Bearer ${userToken}`
  }
  return config
})

export { api }

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('userToken') || '')
  const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))
  const adminToken = ref(staff.getItem('adminToken') || '')
  const admin = ref(JSON.parse(staff.getItem('admin') || 'null'))
  const workerToken = ref(staff.getItem('workerToken') || '')
  const worker = ref(JSON.parse(staff.getItem('worker') || 'null'))
  const doctorToken = ref(staff.getItem('doctorToken') || '')
  const doctor = ref(JSON.parse(staff.getItem('doctor') || 'null'))

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => !!adminToken.value)
  const isWorker = computed(() => !!workerToken.value)
  const isDoctor = computed(() => !!doctorToken.value)

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

  async function updateDeliveryAddress(address) {
    if (user.value) {
      user.value = { ...user.value, delivery_address: address }
      localStorage.setItem('user', JSON.stringify(user.value))
    }
  }

  async function adminLogin(data) {
    const res = await api.post('/admin/login', data)
    if (res.data.role === 'worker') {
      workerToken.value = res.data.token
      worker.value = res.data.worker
      staff.setItem('workerToken', res.data.token)
      staff.setItem('worker', JSON.stringify(res.data.worker))
    } else if (res.data.role === 'doctor') {
      doctorToken.value = res.data.token
      doctor.value = res.data.doctor
      staff.setItem('doctorToken', res.data.token)
      staff.setItem('doctor', JSON.stringify(res.data.doctor))
    } else {
      adminToken.value = res.data.token
      admin.value = res.data.admin
      staff.setItem('adminToken', res.data.token)
      staff.setItem('admin', JSON.stringify(res.data.admin))
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
    staff.removeItem('adminToken')
    staff.removeItem('admin')
  }

  function workerLogout() {
    workerToken.value = ''
    worker.value = null
    staff.removeItem('workerToken')
    staff.removeItem('worker')
  }

  function doctorLogout() {
    doctorToken.value = ''
    doctor.value = null
    staff.removeItem('doctorToken')
    staff.removeItem('doctor')
  }

  return {
    token, user, adminToken, admin, workerToken, worker, doctorToken, doctor,
    isLoggedIn, isAdmin, isWorker, isDoctor,
    register, login, adminLogin, logout, adminLogout, workerLogout, doctorLogout, updateDeliveryAddress
  }
})
