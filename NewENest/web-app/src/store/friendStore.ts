import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import friendApi, { FriendInfo, FriendRequest } from '../api/friendApi'

export const useFriendStore = defineStore('friend', () => {
  // 状态
  const friends = ref<FriendInfo[]>([])
  const friendRequests = ref<FriendRequest[]>([])
  const loading = ref(false)
  const searchResults = ref<any[]>([])
  const searchLoading = ref(false)
  
  // 计算属性
  const friendCount = computed(() => friends.value.length)
  const pendingRequestCount = computed(() => friendRequests.value.length)
  const hasFriends = computed(() => friendCount.value > 0)
  const hasPendingRequests = computed(() => pendingRequestCount.value > 0)
  
  // 方法
  // 加载好友列表
  async function loadFriends() {
    try {
      loading.value = true
      const response = await friendApi.getFriendList()
      
      // 确保我们正确处理API返回格式
      if (Array.isArray(response)) {
        friends.value = response
      } else if (response && response.data && Array.isArray(response.data)) {
        // 如果API返回格式是 {data: [...]}
        friends.value = response.data
      } else {
        console.error('好友列表API返回了意外的格式:', response)
        friends.value = [] // 保证是数组
      }
    } catch (error) {
      console.error('加载好友列表失败', error)
      friends.value = [] // 出错时设置为空数组
    } finally {
      loading.value = false
    }
  }
  
  // 加载好友请求
  async function loadFriendRequests() {
    try {
      loading.value = true
      const response = await friendApi.getFriendRequests()
      
      // 确保我们正确处理API返回格式
      if (Array.isArray(response)) {
        friendRequests.value = response
      } else if (response && response.data && Array.isArray(response.data)) {
        // 如果API返回格式是 {data: [...]}
        friendRequests.value = response.data
      } else {
        console.error('好友请求API返回了意外的格式:', response)
        friendRequests.value = [] // 保证是数组
      }
    } catch (error) {
      console.error('加载好友请求失败', error)
      friendRequests.value = [] // 出错时设置为空数组
    } finally {
      loading.value = false
    }
  }
  
  // 发送好友请求
  async function sendFriendRequest(receiverId: number) {
    return await friendApi.sendFriendRequest(receiverId)
  }
  
  // 接受好友请求
  async function acceptFriendRequest(requestId: number) {
    const result = await friendApi.acceptFriendRequest(requestId)
    if (result.success) {
      // 从请求列表中移除该请求
      friendRequests.value = friendRequests.value.filter(req => req.id !== requestId)
      // 重新加载好友列表
      await loadFriends()
    }
    return result
  }
  
  // 拒绝好友请求
  async function rejectFriendRequest(requestId: number) {
    const result = await friendApi.rejectFriendRequest(requestId)
    if (result.success) {
      // 从请求列表中移除该请求
      friendRequests.value = friendRequests.value.filter(req => req.id !== requestId)
    }
    return result
  }
  
  // 删除好友
  async function deleteFriend(friendshipId: number) {
    const result = await friendApi.deleteFriend(friendshipId)
    if (result.success) {
      // 从好友列表中移除该好友
      friends.value = friends.value.filter(friend => friend.friendship_id !== friendshipId)
    }
    return result
  }
  
  // 搜索用户
  async function searchUsers(keyword: string) {
    if (!keyword.trim()) {
      searchResults.value = []
      return
    }
    
    try {
      searchLoading.value = true
      const response = await friendApi.searchUsers(keyword)
      
      // 确保我们正确处理API返回格式
      if (Array.isArray(response)) {
        searchResults.value = response
      } else if (response && response.data && Array.isArray(response.data)) {
        // 如果API返回格式是 {data: [...]}
        searchResults.value = response.data
      } else {
        console.error('搜索用户API返回了意外的格式:', response)
        searchResults.value = [] // 保证是数组
      }
    } catch (error) {
      console.error('搜索用户失败', error)
      searchResults.value = [] // 出错时设置为空数组
    } finally {
      searchLoading.value = false
    }
  }
  
  // 初始化加载
  async function initialize() {
    await Promise.all([loadFriends(), loadFriendRequests()])
  }
  
  return {
    friends,
    friendRequests,
    loading,
    searchResults,
    searchLoading,
    friendCount,
    pendingRequestCount,
    hasFriends,
    hasPendingRequests,
    loadFriends,
    loadFriendRequests,
    sendFriendRequest,
    acceptFriendRequest,
    rejectFriendRequest,
    deleteFriend,
    searchUsers,
    initialize
  }
}) 