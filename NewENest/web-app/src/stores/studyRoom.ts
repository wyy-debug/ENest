import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { studyRoomApi, type StudyRoomDTO, type StudyRoomDetailDTO } from '@/api/studyRoom';

export const useStudyRoomStore = defineStore('studyRoom', () => {
  // 状态
  const currentRoom = ref<StudyRoomDetailDTO | null>(null);
  const myRooms = ref<StudyRoomDTO[]>([]);
  const publicRooms = ref<StudyRoomDTO[]>([]);
  const totalPublicRooms = ref(0);
  const loading = ref(false);
  const error = ref('');
  
  // 计算属性
  const isInRoom = computed(() => !!currentRoom.value);
  const roomName = computed(() => currentRoom.value?.name || '');
  const memberCount = computed(() => currentRoom.value?.member_count || 0);
  
  // 操作方法
  const fetchPublicRooms = async (page: number = 1, pageSize: number = 10) => {
    loading.value = true;
    error.value = '';
    
    try {
      const result = await studyRoomApi.getPublicStudyRooms(page, pageSize);
      publicRooms.value = result.rooms;
      totalPublicRooms.value = result.total;
    } catch (err: any) {
      error.value = err.message || '获取公开自习室失败';
      console.error('Failed to fetch public rooms:', err);
    } finally {
      loading.value = false;
    }
  };
  
  const fetchMyRooms = async () => {
    loading.value = true;
    error.value = '';
    
    try {
      myRooms.value = await studyRoomApi.getMyStudyRooms();
    } catch (err: any) {
      error.value = err.message || '获取我的自习室失败';
      console.error('Failed to fetch my rooms:', err);
    } finally {
      loading.value = false;
    }
  };
  
  const fetchRoomDetail = async (roomId: number) => {
    loading.value = true;
    error.value = '';
    
    try {
      currentRoom.value = await studyRoomApi.getStudyRoomDetail(roomId);
    } catch (err: any) {
      error.value = err.message || '获取自习室详情失败';
      console.error('Failed to fetch room detail:', err);
    } finally {
      loading.value = false;
    }
  };
  
  const createRoom = async (roomData: {
    name: string;
    description: string;
    max_members: number;
    is_private: boolean;
    theme?: string;
    background_image?: string;
    expires_in: number;
  }) => {
    loading.value = true;
    error.value = '';
    
    try {
      const room = await studyRoomApi.createStudyRoom(roomData);
      await fetchMyRooms();
      return room;
    } catch (err: any) {
      error.value = err.message || '创建自习室失败';
      console.error('Failed to create room:', err);
      throw err;
    } finally {
      loading.value = false;
    }
  };
  
  const joinRoom = async (roomId: number, isAnonymous: boolean = false) => {
    loading.value = true;
    error.value = '';
    
    try {
      await studyRoomApi.joinStudyRoom(roomId, isAnonymous);
      await fetchRoomDetail(roomId);
      await fetchMyRooms();
    } catch (err: any) {
      error.value = err.message || '加入自习室失败';
      console.error('Failed to join room:', err);
      throw err;
    } finally {
      loading.value = false;
    }
  };
  
  const joinRoomByShareLink = async (shareLink: string, isAnonymous: boolean = false) => {
    loading.value = true;
    error.value = '';
    
    try {
      const room = await studyRoomApi.joinByShareLink(shareLink, isAnonymous);
      await fetchMyRooms();
      return room;
    } catch (err: any) {
      error.value = err.message || '通过分享链接加入自习室失败';
      console.error('Failed to join room by share link:', err);
      throw err;
    } finally {
      loading.value = false;
    }
  };
  
  const leaveRoom = async (roomId: number) => {
    if (!roomId && currentRoom.value) {
      roomId = currentRoom.value.id;
    }
    
    loading.value = true;
    error.value = '';
    
    try {
      await studyRoomApi.leaveStudyRoom(roomId);
      if (currentRoom.value && currentRoom.value.id === roomId) {
        currentRoom.value = null;
      }
      await fetchMyRooms();
    } catch (err: any) {
      error.value = err.message || '离开自习室失败';
      console.error('Failed to leave room:', err);
      throw err;
    } finally {
      loading.value = false;
    }
  };
  
  const updateMemberStatus = async (status: string) => {
    if (!currentRoom.value) {
      return;
    }
    
    try {
      await studyRoomApi.updateMemberStatus(currentRoom.value.id, status);
      // 重新获取房间信息以获取最新的成员状态
      await fetchRoomDetail(currentRoom.value.id);
    } catch (err: any) {
      console.error('Failed to update status:', err);
      throw err;
    }
  };
  
  const clearCurrentRoom = () => {
    currentRoom.value = null;
  };
  
  return {
    // 状态
    currentRoom,
    myRooms,
    publicRooms,
    totalPublicRooms,
    loading,
    error,
    
    // 计算属性
    isInRoom,
    roomName,
    memberCount,
    
    // 操作方法
    fetchPublicRooms,
    fetchMyRooms,
    fetchRoomDetail,
    createRoom,
    joinRoom,
    joinRoomByShareLink,
    leaveRoom,
    updateMemberStatus,
    clearCurrentRoom
  };
}); 