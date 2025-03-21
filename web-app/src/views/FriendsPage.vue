<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { wsClient } from '../utils/websocket'
import { MessageType, FriendOperation, type FriendMessage } from '../proto/message'

interface Friend {
  id: number
  username: string
  avatar?: string
  online: boolean
}

const friends = ref<Friend[]>([])
const loading = ref(false)

onMounted(() => {
  initWebSocket()
  fetchFriends()
})

const initWebSocket = () => {
  wsClient.connect('ws://localhost:8080/ws')
  wsClient.registerHandler(MessageType.FRIEND, handleFriendResponse)
  wsClient.registerHandler(MessageType.ERROR, handleErrorResponse)
}

const handleFriendResponse = (payload: Uint8Array) => {
  const response = JSON.parse(new TextDecoder().decode(payload))
  if (Array.isArray(response)) {
    friends.value = response
  }
  loading.value = false
}

const handleErrorResponse = (payload: Uint8Array) => {
  const error = JSON.parse(new TextDecoder().decode(payload))
  ElMessage.error(error.message || '操作失败')
  loading.value = false
}

const fetchFriends = () => {
  loading.value = true
  const message: FriendMessage = {
    operation: FriendOperation.GET_LIST
  }
  const payload = new TextEncoder().encode(JSON.stringify(message))
  wsClient.sendMessage(MessageType.FRIEND, payload)
}

const handleAddFriend = () => {
  const message: FriendMessage = {
    operation: FriendOperation.SEND_REQUEST,
    friendId: 0 // TODO: 需要添加一个对话框让用户输入好友ID
  }
  const payload = new TextEncoder().encode(JSON.stringify(message))
  wsClient.sendMessage(MessageType.FRIEND, payload)
}

const handleDeleteFriend = (friendId: number) => {
  const message: FriendMessage = {
    operation: FriendOperation.DELETE,
    friendId
  }
  const payload = new TextEncoder().encode(JSON.stringify(message))
  wsClient.sendMessage(MessageType.FRIEND, payload)
}
</script>

<template>
  <div class="friends-container">
    <div class="friends-header">
      <h2>好友列表</h2>
      <el-button type="primary" size="large" @click="handleAddFriend">
        <el-icon><i-ep-plus /></el-icon>
        添加好友
      </el-button>
    </div>
    
    <div class="friends-content" v-loading="loading">
      <el-empty v-if="friends.length === 0" description="暂无好友" />
      <div v-else class="friends-list">
        <el-card v-for="friend in friends" :key="friend.id" class="friend-card">
          <div class="friend-info">
            <el-avatar :size="50" :src="friend.avatar">
              {{ friend.username?.charAt(0).toUpperCase() }}
            </el-avatar>
            <div class="friend-details">
              <h3>{{ friend.username }}</h3>
              <p class="friend-status" :class="{ 'online': friend.online }">
                {{ friend.online ? '在线' : '离线' }}
              </p>
            </div>
          </div>
          <div class="friend-actions">
            <el-button type="danger" size="small" @click="handleDeleteFriend(friend.id)">
              <el-icon><i-ep-delete /></el-icon>
              删除好友
            </el-button>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<style scoped>
.friends-container {
  padding: 20px;
  height: 100%;
  overflow-y: auto;
}

.friends-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.friends-header h2 {
  margin: 0;
  color: #333;
}

.friends-content {
  max-width: 800px;
  margin: 0 auto;
}

.friends-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.friend-card {
  background: white;
  border-radius: 8px;
  transition: all 0.3s;
}

.friend-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.friend-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.friend-details {
  flex: 1;
}

.friend-details h3 {
  margin: 0 0 4px 0;
  color: #333;
}

.friend-status {
  margin: 0;
  font-size: 14px;
  color: #999;
}

.friend-status.online {
  color: #67C23A;
}

.friend-actions {
  display: flex;
  gap: 8px;
  margin-top: 12px;
}
</style>