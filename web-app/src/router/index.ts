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
    path: '/main',
    name: 'main',
    component: () => import('../views/MainPage.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: 'study-room',
        name: 'studyRoom',
        component: () => import('../views/StudyRoomPage.vue')
      },
      {
        path: 'create-study-room',
        name: 'createStudyRoom',
        component: () => import('../views/CreateStudyRoom.vue')
      },
      {
        path: 'profile',
        name: 'profile',
        component: () => import('../views/ProfilePage.vue')
      },
      {
        path: 'community',
        name: 'community',
        component: () => import('../views/CommunityPage.vue')
      },
      {
        path: 'friends',
        name: 'friends',
        component: () => import('../views/FriendsPage.vue')
      }
    ]
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