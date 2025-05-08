import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '../config/api'

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

// 登录响应类型
interface LoginResponse {
  token: string
  user: UserInfo
}

// 登录结果类型
interface LoginResult {
  success: boolean
  message?: string
  data?: LoginResponse
  token?: string
  user?: UserInfo
}

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref<string | null>(localStorage.getItem('token'))
  const userInfo = ref<UserInfo | null>(null)
  const loading = ref(false)
  
  // 计算属性
  const isLoggedIn = computed(() => !!token.value)
  
  // 方法
  async function login(email: string, password: string): Promise<LoginResult> {
    try {
      loading.value = true
      
      // 通过API发起登录请求
      // 响应拦截器已经提取了data，返回的是data内容而不是完整的AxiosResponse
      const response = await api.post('/auth/login', { email, password }) as LoginResponse
      
      if (!response || !response.token || !response.user) {
        console.error('登录响应格式不正确:', response)
        return {
          success: false,
          message: '登录响应格式不正确'
        }
      }
      
      // 更新状态
      token.value = response.token
      userInfo.value = response.user
      
      // 存储token (确保添加Bearer前缀)
      const formattedToken = response.token.startsWith('Bearer ') 
        ? response.token 
        : `Bearer ${response.token}`
      localStorage.setItem('token', formattedToken)
      
      // 返回成功结果
      return {
        success: true,
        data: response,
        token: response.token,
        user: response.user
      }
    } catch (error: any) {
      console.error('登录失败:', error)
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
      await api.post('/auth/register', {
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
      const response = await api.get('/user') as UserInfo
      userInfo.value = response
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
      const response = await api.put('/user', data) as UserInfo
      userInfo.value = { ...userInfo.value, ...response } as UserInfo
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