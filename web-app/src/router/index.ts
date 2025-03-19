import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomePage.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginPage.vue')
  },
  {
    path: '/study-room',
    name: 'studyRoom',
    component: () => import('../views/StudyRoomPage.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/zoom',
    name: 'zoom',
    component: () => import('../views/ZoomView.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const token = localStorage.getItem('session_token')
  const userData = localStorage.getItem('user_data')

  if (requiresAuth && (!token || !userData)) {
    next('/login')
  } else if (to.path === '/login' && token && userData) {
    next('/')
  } else {
    next()
  }
})

export default router