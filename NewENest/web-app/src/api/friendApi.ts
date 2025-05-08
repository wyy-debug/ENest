import api from '../config/api';

// 好友信息接口
export interface FriendInfo {
  id: number;
  friendship_id: number;
  username: string;
  avatar?: string;
  signature?: string;
  study_direction?: string;
  total_study_time: number;
  friend_since: string;
  has_active_contract: boolean;
  unread_messages: number;
}

// 好友请求接口
export interface FriendRequest {
  id: number;
  user_id: number;
  created_at: string;
  status: string;
  sender?: {
    id: number;
    username: string;
    avatar?: string;
    signature?: string;
    study_direction?: string;
  };
}

// API响应接口
export interface ApiResponse<T> {
  code?: number;
  message?: string;
  data: T;
  total?: number;
}

// 好友API服务
const friendApi = {
  // 获取好友列表
  getFriendList: async (): Promise<ApiResponse<FriendInfo[]> | FriendInfo[]> => {
    const response = await api.get('/friends');
    return response;
  },

  // 获取好友请求列表
  getFriendRequests: async (): Promise<ApiResponse<FriendRequest[]> | FriendRequest[]> => {
    const response = await api.get('/friends/requests');
    return response;
  },

  // 发送好友请求
  sendFriendRequest: async (receiverId: number): Promise<{ success: boolean; message?: string }> => {
    try {
      await api.post('/friends/requests', { receiver_id: receiverId });
      return { success: true };
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '发送好友请求失败，请稍后再试'
      };
    }
  },

  // 接受好友请求
  acceptFriendRequest: async (requestId: number): Promise<{ success: boolean; message?: string }> => {
    try {
      await api.post('/friends/requests/response', {
        request_id: requestId,
        action: 'accept'
      });
      return { success: true };
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '接受好友请求失败，请稍后再试'
      };
    }
  },

  // 拒绝好友请求
  rejectFriendRequest: async (requestId: number): Promise<{ success: boolean; message?: string }> => {
    try {
      await api.post('/friends/requests/response', {
        request_id: requestId,
        action: 'reject'
      });
      return { success: true };
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '拒绝好友请求失败，请稍后再试'
      };
    }
  },

  // 删除好友
  deleteFriend: async (friendshipId: number): Promise<{ success: boolean; message?: string }> => {
    try {
      await api.delete(`/friends/${friendshipId}`);
      return { success: true };
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '删除好友失败，请稍后再试'
      };
    }
  },

  // 搜索用户
  searchUsers: async (keyword: string): Promise<ApiResponse<any[]> | any[]> => {
    const response = await api.get(`/users/search?keyword=${encodeURIComponent(keyword)}`);
    return response;
  },

  // 获取与某个好友的聊天记录
  getChatHistory: async (friendId: number, page = 1, pageSize = 20): Promise<any> => {
    try {
      console.log(`调用API: /friends/${friendId}/messages, 参数:`, { page, pageSize })
      const response = await api.get(`/friends/${friendId}/messages`, {
        params: { page, pageSize }
      });
      console.log('原始API响应:', response)
      
      // 如果response是标准格式(有code,data等字段)
      if (response && typeof response === 'object' && 'data' in response) {
        return response;
      }
      
      // 如果response直接就是数据数组
      if (Array.isArray(response)) {
        return { data: response };
      }
      
      // 其他格式，包装一下返回
      return { data: response || [] };
    } catch (error: any) {
      console.error('获取聊天记录失败:', error);
      return { data: [] };
    }
  },

  // 发送消息给好友
  sendMessage: async (receiverId: number, content: string, messageType = 'text'): Promise<any> => {
    try {
      console.log('调用发送消息API, 参数:', { receiver_id: receiverId, message_type: messageType, content })
      const response = await api.post('/friends/messages', {
        receiver_id: receiverId,
        message_type: messageType,
        content
      });
      console.log('发送消息API响应:', response)
      
      // 如果response是标准格式(有code,data等字段)
      if (response && typeof response === 'object' && 'data' in response) {
        return response;
      }
      
      // 如果response直接就是消息对象
      if (response && typeof response === 'object' && 'content' in response) {
        return { data: response };
      }
      
      // 其他格式，返回null
      return { data: null };
    } catch (error: any) {
      console.error('发送消息失败:', error);
      throw error;
    }
  }
};

export default friendApi; 