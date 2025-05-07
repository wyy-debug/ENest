import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

export interface UserInfo {
  id: number
  username: string
  email: string
  avatar?: string
  signature?: string
  studyDirection?: string
  totalStudyTime: number
  achievementPoints: number
}

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref<string | null>(localStorage.getItem('token'))
  const userInfo = ref<UserInfo | null>(null)
  const loading = ref(false)
  
  // 计算属性
  const isLoggedIn = computed(() => !!token.value)
  
  // 方法
  async function login(email: string, password: string) {
    try {
      loading.value = true
      const response = await axios.post('/api/auth/login', { email, password })
      const { token: newToken, user } = response.data
      
      token.value = newToken
      userInfo.value = user
      localStorage.setItem('token', newToken)
      
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '登录失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  async function register(username: string, email: string, password: string) {
    try {
      loading.value = true
      const response = await axios.post('/api/auth/register', {
        username,
        email,
        password
      })
      
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '注册失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  async function getUserInfo() {
    if (!token.value || userInfo.value) return
    
    try {
      loading.value = true
      const response = await axios.get('/api/users/me')
      userInfo.value = response.data
    } catch (error) {
      logout()
    } finally {
      loading.value = false
    }
  }
  
  function logout() {
    token.value = null
    userInfo.value = null
    localStorage.removeItem('token')
  }
  
  async function updateProfile(data: Partial<UserInfo>) {
    try {
      loading.value = true
      const response = await axios.put('/api/users/me', data)
      userInfo.value = { ...userInfo.value, ...response.data } as UserInfo
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '更新失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  return {
    token,
    userInfo,
    loading,
    isLoggedIn,
    login,
    register,
    getUserInfo,
    logout,
    updateProfile
  }
}) 