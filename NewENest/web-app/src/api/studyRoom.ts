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

// 自习室相关DTO类型定义
export interface StudyRoomDTO {
  id: number;
  name: string;
  description: string;
  share_link?: string;
  max_members: number;
  is_private: boolean;
  theme: string;
  background_image?: string;
  created_at: string;
  expires_at: string;
  owner?: {
    id: number;
    username: string;
    avatar?: string;
  };
  member_count: number;
}

export interface StudyRoomDetailDTO extends StudyRoomDTO {
  members?: RoomMemberDTO[];
}

export interface RoomMemberDTO {
  user_id: number;
  username?: string;
  avatar?: string;
  is_anonymous: boolean;
  role: string;
  status: string;
  joined_at: string;
}

export interface CreateStudyRoomDTO {
  name: string;
  description: string;
  max_members: number;
  is_private: boolean;
  theme?: string;
  background_image?: string;
  expires_in: number; // 过期时间（小时）
}

export interface UpdateStudyRoomDTO {
  name?: string;
  description?: string;
  max_members?: number;
  is_private?: boolean;
  theme?: string;
  background_image?: string;
}

// 使用与user.ts相同的axios实例
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

export const studyRoomApi = {
  // 获取自习室列表
  getStudyRooms: async (page: number = 1, pageSize: number = 10): Promise<{rooms: StudyRoomDTO[], total: number}> => {
    const response: AxiosResponse<ApiResponse<{rooms: StudyRoomDTO[], total: number}>> = 
      await api.get('/study-rooms', { params: { page, pageSize } });
    
    return response.data.data!;
  },
  
  // 获取公开自习室列表
  getPublicStudyRooms: async (page: number = 1, pageSize: number = 10): Promise<{rooms: StudyRoomDTO[], total: number}> => {
    const response: AxiosResponse<ApiResponse<{rooms: StudyRoomDTO[], total: number}>> = 
      await api.get('/study-rooms/public', { params: { page, pageSize } });
    
    return response.data.data!;
  },
  
  // 获取我的自习室列表
  getMyStudyRooms: async (): Promise<StudyRoomDTO[]> => {
    const response: AxiosResponse<ApiResponse<StudyRoomDTO[]>> = 
      await api.get('/study-rooms/mine');
    
    return response.data.data!;
  },
  
  // 获取自习室详情
  getStudyRoomDetail: async (roomId: number): Promise<StudyRoomDetailDTO> => {
    const response: AxiosResponse<ApiResponse<StudyRoomDetailDTO>> = 
      await api.get(`/study-rooms/${roomId}`);
    
    return response.data.data!;
  },
  
  // 创建自习室
  createStudyRoom: async (roomData: CreateStudyRoomDTO): Promise<StudyRoomDTO> => {
    const response: AxiosResponse<ApiResponse<StudyRoomDTO>> = 
      await api.post('/study-rooms', roomData);
    
    return response.data.data!;
  },
  
  // 更新自习室
  updateStudyRoom: async (roomId: number, roomData: UpdateStudyRoomDTO): Promise<StudyRoomDTO> => {
    const response: AxiosResponse<ApiResponse<StudyRoomDTO>> = 
      await api.put(`/study-rooms/${roomId}`, roomData);
    
    return response.data.data!;
  },
  
  // 删除自习室
  deleteStudyRoom: async (roomId: number): Promise<void> => {
    await api.delete(`/study-rooms/${roomId}`);
  },
  
  // 加入自习室
  joinStudyRoom: async (roomId: number, isAnonymous: boolean = false): Promise<void> => {
    await api.post(`/study-rooms/${roomId}/join`, { is_anonymous: isAnonymous });
  },
  
  // 通过分享链接加入自习室
  joinByShareLink: async (shareLink: string, isAnonymous: boolean = false): Promise<StudyRoomDTO> => {
    const response: AxiosResponse<ApiResponse<StudyRoomDTO>> = 
      await api.post(`/study-rooms/join/${shareLink}`, { is_anonymous: isAnonymous });
    
    return response.data.data!;
  },
  
  // 离开自习室
  leaveStudyRoom: async (roomId: number): Promise<void> => {
    await api.post(`/study-rooms/${roomId}/leave`);
  },
  
  // 更新成员状态
  updateMemberStatus: async (roomId: number, status: string): Promise<void> => {
    await api.put(`/study-rooms/${roomId}/status`, { status });
  }
}; 