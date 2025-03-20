<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import axios from '../utils/axios'

const friends = ref([])
const loading = ref(false)

onMounted(() => {
  fetchFriends()
})

const fetchFriends = async () => {
  loading.value = true
  try {
    const response = await axios.get('/friends')
    friends.value = response.data
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '获取好友列表失败')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="friends-container">
    <div class="friends-header">
      <h2>好友列表</h2>
      <el-button type="primary" size="large">
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
            <el-button type="primary" size="small">
              <el-icon><i-ep-message /></el-icon>
              发消息
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
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
  font-size: 16px;
  color: #333;
}

.friend-status {
  margin: 0;
  font-size: 14px;
  color: #999;
}

.friend-status.online {
  color: #67c23a;
}

.friend-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}
</style>