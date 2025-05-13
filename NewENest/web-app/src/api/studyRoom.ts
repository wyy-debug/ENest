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

// 使用从config/api.ts导入的全局api实例，删除这里的重复定义
// const api = axios.create({
//   baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
//   timeout: 10000,
//   headers: {
//     'Content-Type': 'application/json',
//     'Accept': 'application/json',
//   }
// });

// 请求和响应拦截器已在config/api.ts中定义，此处无需重复
// 删除拦截器代码
// ...

export const studyRoomApi = {
  // 获取自习室列表
  getStudyRooms: async (page: number = 1, pageSize: number = 10): Promise<{rooms: StudyRoomDTO[], total: number}> => {
    const response = await api.get('/study-rooms', { params: { page, pageSize } });
    return response.data || { rooms: [], total: 0 };
  },
  
  // 获取公开自习室列表
  getPublicStudyRooms: async (page: number = 1, pageSize: number = 10): Promise<{rooms: StudyRoomDTO[], total: number}> => {
    const response = await api.get('/study-rooms/public', { params: { page, pageSize } });
    return response.data || { rooms: [], total: 0 };
  },
  
  // 获取我的自习室列表
  getMyStudyRooms: async (): Promise<StudyRoomDTO[]> => {
    const response = await api.get('/study-rooms/mine');
    return response.data || [];
  },
  
  // 获取自习室详情
  getStudyRoomDetail: async (roomId: number): Promise<StudyRoomDetailDTO> => {
    const response = await api.get(`/study-rooms/${roomId}`);
    return response.data;
  },
  
  // 创建自习室
  createStudyRoom: async (roomData: CreateStudyRoomDTO): Promise<StudyRoomDTO> => {
    const response = await api.post('/study-rooms', roomData);
    return response.data;
  },
  
  // 更新自习室
  updateStudyRoom: async (roomId: number, roomData: UpdateStudyRoomDTO): Promise<StudyRoomDTO> => {
    const response = await api.put(`/study-rooms/${roomId}`, roomData);
    return response.data;
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
    const response = await api.post(`/study-rooms/join/${shareLink}`, { is_anonymous: isAnonymous });
    return response.data;
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