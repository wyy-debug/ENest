import axios from 'axios';
import type { AxiosResponse } from 'axios';
import { useUserStore } from '@/stores/user';

// API响应类型定义
interface ApiResponse<T> {
  code: number;
  message: string;
  data?: T;
  errors?: string[] | string;
}

// 用户相关DTO类型定义
export interface UserProfileDTO {
  id: number;
  username: string;
  email: string;
  avatar?: string;
  signature?: string;
  study_direction?: string;
  total_study_time: number;
  achievement_points: number;
  created_at: string;
}

export interface LoginDTO {
  username: string;
  password: string;
}

export interface RegisterDTO {
  username: string;
  email: string;
  password: string;
}

export interface UpdateProfileDTO {
  username?: string;
  signature?: string;
  study_direction?: string;
  avatar?: string;
}

export interface ChangePasswordDTO {
  old_password: string;
  new_password: string;
}

// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  }
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const userStore = useUserStore();
    if (userStore.token) {
      config.headers['Authorization'] = userStore.token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      // 未授权，清除用户状态并重定向到登录页
      const userStore = useUserStore();
      userStore.logout();
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export const userApi = {
  // 用户登录
  login: async (loginData: LoginDTO): Promise<{user: UserProfileDTO, token: string}> => {
    const response: AxiosResponse<ApiResponse<{user: UserProfileDTO, token: string}>> = 
      await api.post('/users/login', loginData);
    
    return response.data.data!;
  },
  
  // 用户注册
  register: async (registerData: RegisterDTO): Promise<UserProfileDTO> => {
    const response: AxiosResponse<ApiResponse<UserProfileDTO>> = 
      await api.post('/users/register', registerData);
    
    return response.data.data!;
  },
  
  // 用户登出
  logout: async (): Promise<void> => {
    await api.post('/users/logout');
  },
  
  // 获取用户资料
  getProfile: async (): Promise<UserProfileDTO> => {
    const response: AxiosResponse<ApiResponse<UserProfileDTO>> = 
      await api.get('/users/profile');
    
    return response.data.data!;
  },
  
  // 更新用户资料
  updateProfile: async (profileData: UpdateProfileDTO): Promise<UserProfileDTO> => {
    const response: AxiosResponse<ApiResponse<UserProfileDTO>> = 
      await api.put('/users/profile', profileData);
    
    return response.data.data!;
  },
  
  // 修改密码
  changePassword: async (passwordData: ChangePasswordDTO): Promise<void> => {
    await api.put('/users/password', passwordData);
  }
}; 