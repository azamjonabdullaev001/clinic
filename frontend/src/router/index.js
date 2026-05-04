import { createRouter, createWebHistory } from 'vue-router'
import Landing from '../views/Landing.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import AdminLogin from '../views/AdminLogin.vue'
import AdminPanel from '../views/AdminPanel.vue'
import PickupPanel from '../views/PickupPanel.vue'
import Support from '../views/Support.vue'

const routes = [
  { path: '/', name: 'Landing', component: Landing },
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  { path: '/admin/login', name: 'AdminLogin', component: AdminLogin },
  { path: '/admin', name: 'AdminPanel', component: AdminPanel, meta: { requiresAdmin: true } },
  { path: '/pickup', name: 'PickupPanel', component: PickupPanel, meta: { requiresWorker: true } },
  { path: '/support', name: 'Support', component: Support, meta: { requiresUser: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAdmin) {
    const token = localStorage.getItem('adminToken')
    if (!token) {
      next('/admin/login')
      return
    }
  }
  if (to.meta.requiresWorker) {
    const token = localStorage.getItem('workerToken')
    if (!token) {
      next('/admin/login')
      return
    }
  }
  if (to.meta.requiresUser) {
    const token = localStorage.getItem('userToken')
    if (!token) {
      next('/login')
      return
    }
  }
  next()
})

export default router
