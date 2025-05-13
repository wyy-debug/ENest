import axios from 'axios';
import type { AxiosResponse } from 'axios';
import { useUserStore } from '@/stores/user';
import api from '@/config/api';

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

// 使用从config/api.ts导入的全局api实例
// 以下代码已删除
// const api = axios.create({
//   baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
//   timeout: 10000,
//   headers: {
//     'Content-Type': 'application/json',
//     'Accept': 'application/json',
//   }
// });

// 请求和响应拦截器已在config/api.ts中定义，此处删除
// api.interceptors.request.use...
// api.interceptors.response.use...

export const userApi = {
  // 用户登录
  login: async (loginData: LoginDTO): Promise<{user: UserProfileDTO, token: string}> => {
    const response = await api.post('/users/login', loginData);
    return response.data || { user: {} as UserProfileDTO, token: '' };
  },
  
  // 用户注册
  register: async (registerData: RegisterDTO): Promise<UserProfileDTO> => {
    const response = await api.post('/users/register', registerData);
    return response.data || {} as UserProfileDTO;
  },
  
  // 用户登出
  logout: async (): Promise<void> => {
    await api.post('/users/logout');
  },
  
  // 获取用户资料
  getProfile: async (): Promise<UserProfileDTO> => {
    const response = await api.get('/users/profile');
    return response.data || {} as UserProfileDTO;
  },
  
  // 更新用户资料
  updateProfile: async (profileData: UpdateProfileDTO): Promise<UserProfileDTO> => {
    const response = await api.put('/users/profile', profileData);
    return response.data || {} as UserProfileDTO;
  },
  
  // 修改密码
  changePassword: async (passwordData: ChangePasswordDTO): Promise<void> => {
    await api.put('/users/password', passwordData);
  }
}; 