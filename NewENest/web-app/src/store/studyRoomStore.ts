import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

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
      const response = await axios.get('/api/v1/study-rooms')
      // 确保response.data是数组
      if (Array.isArray(response.data)) {
        rooms.value = response.data
      } else if (response.data && Array.isArray(response.data.data)) {
        // 如果API返回格式是 {data: [...]}
        rooms.value = response.data.data
      } else {
        console.error('Unexpected API response format:', response.data)
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
      const response = await axios.get(`/api/v1/study-rooms/${id}`)
      // 确保API返回格式正确
      if (response.data && response.data.id) {
        currentRoom.value = response.data
        return response.data
      } else if (response.data && response.data.data && response.data.data.id) {
        // 如果API返回格式是 {data: {...}}
        currentRoom.value = response.data.data
        return response.data.data
      } else {
        console.error(`Unexpected API response format for room ${id}:`, response.data)
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
      const response = await axios.post('/api/v1/study-rooms', roomData)
      const newRoom = response.data && response.data.data ? response.data.data : response.data
      
      if (newRoom && newRoom.id) {
        rooms.value.push(newRoom)
        return { success: true, data: newRoom }
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
      const response = await axios.put(`/api/v1/study-rooms/${id}`, roomData)
      
      // 提取正确的响应数据
      const updatedRoom = response.data && response.data.data ? response.data.data : response.data
      
      if (updatedRoom && updatedRoom.id) {
        // 更新本地数据
        const index = rooms.value.findIndex(room => room.id === id)
        if (index !== -1) {
          rooms.value[index] = { ...rooms.value[index], ...updatedRoom }
        }
        
        if (currentRoom.value?.id === id) {
          currentRoom.value = { ...currentRoom.value, ...updatedRoom }
        }
        
        return { success: true, data: updatedRoom }
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
      const response = await axios.post(`/api/v1/study-rooms/join`, { 
        room_id: id,
        is_anonymous: isAnonymous 
      })
      
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
      await axios.post(`/api/v1/study-rooms/${id}/leave`)
      
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