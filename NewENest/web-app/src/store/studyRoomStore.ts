import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/config/api'
import { studyRoomApi, type StudyRoomDTO, type StudyRoomDetailDTO } from '@/api/studyRoom'

export interface RoomMember {
  id: number
  userId: number
  username: string
  avatar?: string
  isAnonymous: boolean
  role: 'owner' | 'admin' | 'member'
  status: 'online' | 'away' | 'offline'
  joinedAt: string
}

export interface StudyRoom {
  id: number
  ownerId: number
  name: string
  description?: string
  shareLink?: string
  maxMembers: number
  isPrivate: boolean
  theme: string
  backgroundImage?: string
  createdAt: string
  expiresAt: string
  members?: RoomMember[]
  currentMembers?: number
}

// 将StudyRoomDTO转换为StudyRoom的辅助函数
function convertDtoToStudyRoom(dto: StudyRoomDTO): StudyRoom {
  return {
    id: dto.id,
    ownerId: dto.owner?.id || 0,
    name: dto.name,
    description: dto.description,
    shareLink: dto.share_link,
    maxMembers: dto.max_members,
    isPrivate: dto.is_private,
    theme: dto.theme,
    backgroundImage: dto.background_image,
    createdAt: dto.created_at,
    expiresAt: dto.expires_at,
    currentMembers: dto.member_count
  }
}

// 将StudyRoomDetailDTO转换为StudyRoom的辅助函数
function convertDetailDtoToStudyRoom(dto: StudyRoomDetailDTO): StudyRoom {
  const room = convertDtoToStudyRoom(dto);
  if (dto.members) {
    room.members = dto.members.map(m => ({
      id: m.user_id,
      userId: m.user_id,
      username: m.username || 'Anonymous',
      avatar: m.avatar,
      isAnonymous: m.is_anonymous,
      role: m.role as 'owner' | 'admin' | 'member',
      status: m.status as 'online' | 'away' | 'offline',
      joinedAt: m.joined_at
    }));
  }
  return room;
}

export const useStudyRoomStore = defineStore('studyRoom', () => {
  // 状态
  const rooms = ref<StudyRoom[]>([])
  const currentRoom = ref<StudyRoom | null>(null)
  const loading = ref(false)
  
  // 计算属性
  const myRooms = computed(() => rooms.value.filter(room => room.ownerId === 1)) // 暂时硬编码用户ID
  const joinedRooms = computed(() => rooms.value.filter(room => room.ownerId !== 1))
  
  // 方法
  async function fetchRooms() {
    try {
      loading.value = true
      // 使用studyRoomApi代替axios
      const response = await studyRoomApi.getStudyRooms()
      // 处理不同的响应格式
      if (response && response.rooms) {
        // 转换DTO为本地格式
        rooms.value = response.rooms.map(convertDtoToStudyRoom)
      } else if (Array.isArray(response)) {
        // 如果直接返回数组，假设它们是StudyRoomDTO[]
        rooms.value = response.map(convertDtoToStudyRoom)
      } else {
        console.error('Unexpected API response format:', response)
        rooms.value = [] // 始终确保是数组
      }
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
      rooms.value = [] // 出错时设置为空数组
    } finally {
      loading.value = false
    }
  }
  
  async function fetchRoomById(id: number) {
    try {
      loading.value = true
      // 使用studyRoomApi代替axios
      const response = await studyRoomApi.getStudyRoomDetail(id)
      // 确保API返回格式正确
      if (response && response.id) {
        // 转换为本地StudyRoom格式
        const studyRoom = convertDetailDtoToStudyRoom(response)
        currentRoom.value = studyRoom
        return studyRoom
      } else {
        console.error(`Unexpected API response format for room ${id}:`, response)
        return null
      }
    } catch (error) {
      console.error(`Failed to fetch room ${id}:`, error)
      return null
    } finally {
      loading.value = false
    }
  }
  
  async function createRoom(roomData: Partial<StudyRoom>) {
    try {
      loading.value = true
      // 转换为CreateStudyRoomDTO格式
      const createDto = {
        name: roomData.name || '',
        description: roomData.description || '',
        max_members: roomData.maxMembers || 10,
        is_private: roomData.isPrivate || false,
        theme: roomData.theme,
        background_image: roomData.backgroundImage,
        expires_in: 24 // 默认24小时
      }
      
      // 使用studyRoomApi代替axios
      const newRoom = await studyRoomApi.createStudyRoom(createDto)
      
      if (newRoom && newRoom.id) {
        // 转换回StudyRoom格式
        const formattedRoom = convertDtoToStudyRoom(newRoom)
        rooms.value.push(formattedRoom)
        return { success: true, data: formattedRoom }
      } else {
        return {
          success: false,
          message: '创建自习室失败，服务器返回格式异常'
        }
      }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '创建自习室失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  async function updateRoom(id: number, roomData: Partial<StudyRoom>) {
    try {
      loading.value = true
      // 转换为UpdateStudyRoomDTO格式
      const updateDto = {
        name: roomData.name,
        description: roomData.description,
        max_members: roomData.maxMembers,
        is_private: roomData.isPrivate,
        theme: roomData.theme,
        background_image: roomData.backgroundImage
      }
      
      // 使用studyRoomApi代替axios
      const updatedRoom = await studyRoomApi.updateStudyRoom(id, updateDto)
      
      if (updatedRoom && updatedRoom.id) {
        // 转换回StudyRoom格式
        const formattedRoom = convertDtoToStudyRoom(updatedRoom)
        
        // 更新本地数据
        const index = rooms.value.findIndex(room => room.id === id)
        if (index !== -1) {
          rooms.value[index] = { ...rooms.value[index], ...formattedRoom }
        }
        
        if (currentRoom.value?.id === id) {
          currentRoom.value = { ...currentRoom.value, ...formattedRoom }
        }
        
        return { success: true, data: formattedRoom }
      } else {
        return {
          success: false,
          message: '更新自习室失败，服务器返回格式异常'
        }
      }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '更新自习室失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  async function joinRoom(id: number, isAnonymous = false) {
    try {
      loading.value = true
      // 使用studyRoomApi代替axios
      await studyRoomApi.joinStudyRoom(id, isAnonymous)
      
      // 更新本地数据
      await fetchRoomById(id)
      
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '加入自习室失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  async function leaveRoom(id: number) {
    try {
      loading.value = true
      // 使用studyRoomApi代替axios
      await studyRoomApi.leaveStudyRoom(id)
      
      // 更新本地数据
      if (currentRoom.value?.id === id) {
        currentRoom.value = null
      }
      
      return { success: true }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.message || '离开自习室失败，请稍后再试'
      }
    } finally {
      loading.value = false
    }
  }
  
  return {
    rooms,
    currentRoom,
    loading,
    myRooms,
    joinedRooms,
    fetchRooms,
    fetchRoomById,
    createRoom,
    updateRoom,
    joinRoom,
    leaveRoom
  }
}) 