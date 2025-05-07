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
      const response = await axios.get('/api/study-rooms')
      rooms.value = response.data
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
    } finally {
      loading.value = false
    }
  }
  
  async function fetchRoomById(id: number) {
    try {
      loading.value = true
      const response = await axios.get(`/api/study-rooms/${id}`)
      currentRoom.value = response.data
      return response.data
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
      const response = await axios.post('/api/study-rooms', roomData)
      rooms.value.push(response.data)
      return { success: true, data: response.data }
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
      const response = await axios.put(`/api/study-rooms/${id}`, roomData)
      
      // 更新本地数据
      const index = rooms.value.findIndex(room => room.id === id)
      if (index !== -1) {
        rooms.value[index] = { ...rooms.value[index], ...response.data }
      }
      
      if (currentRoom.value?.id === id) {
        currentRoom.value = { ...currentRoom.value, ...response.data }
      }
      
      return { success: true, data: response.data }
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
      const response = await axios.post(`/api/study-rooms/${id}/join`, { isAnonymous })
      
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
      await axios.post(`/api/study-rooms/${id}/leave`)
      
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