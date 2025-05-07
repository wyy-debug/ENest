<template>
  <div class="home-container">
    <div class="welcome-section">
      <h1>欢迎来到 NewENest</h1>
      <p>在这里，你可以创建或加入自习室，与志同道合的伙伴一起高效学习。</p>
      
      <div class="stats-card">
        <div class="stat-item">
          <div class="stat-value">{{ userCount }}</div>
          <div class="stat-label">活跃用户</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ roomCount }}</div>
          <div class="stat-label">自习室数量</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ formatTime(totalStudyTime) }}</div>
          <div class="stat-label">累计学习时间</div>
        </div>
      </div>
    </div>
    
    <el-row :gutter="20">
      <el-col :xs="24" :sm="24" :md="16" :lg="18">
        <div class="section-heading">
          <h2>公开自习室</h2>
          <el-button
            type="primary"
            plain
            @click="handleCreateRoom"
          >
            创建自习室
          </el-button>
        </div>
        
        <el-row :gutter="20" v-loading="loading">
          <el-col 
            v-for="room in publicRooms" 
            :key="room.id" 
            :xs="24" 
            :sm="12" 
            :md="8" 
            :lg="8"
            :xl="6"
            class="mb-10"
          >
            <el-card class="room-card" @click="handleJoinRoom(room)">
              <div class="room-theme" :class="`theme-${room.theme}`"></div>
              <div class="room-info">
                <h3>{{ room.name }}</h3>
                <p class="room-description">{{ room.description || '暂无描述' }}</p>
                <div class="room-meta">
                  <span class="room-member-count">
                    <el-icon><User /></el-icon>
                    {{ room.currentMembers }} / {{ room.maxMembers }}
                  </span>
                </div>
              </div>
              <el-button type="primary" size="small" class="join-btn">进入自习室</el-button>
            </el-card>
          </el-col>
          
          <el-empty v-if="publicRooms.length === 0" description="暂无公开自习室" />
        </el-row>
      </el-col>
      
      <el-col :xs="24" :sm="24" :md="8" :lg="6">
        <div class="side-section">
          <h2>我的自习室</h2>
          <div class="my-rooms" v-loading="loading">
            <div v-for="room in myRooms" :key="room.id" class="my-room-item">
              <div class="room-name">{{ room.name }}</div>
              <el-button
                type="primary"
                size="small"
                @click="handleJoinRoom(room)"
              >
                进入
              </el-button>
            </div>
            
            <el-empty v-if="myRooms.length === 0" description="暂无自习室" />
          </div>
        </div>
        
        <div class="side-section">
          <h2>学习排行榜</h2>
          <div class="leaderboard">
            <div v-for="(user, index) in leaderboard" :key="user.id" class="leaderboard-item">
              <div class="rank">{{ index + 1 }}</div>
              <el-avatar :size="32" :src="user.avatar">{{ user.username.charAt(0) }}</el-avatar>
              <div class="user-info">
                <div class="username">{{ user.username }}</div>
                <div class="study-time">{{ formatTime(user.studyTime) }}</div>
              </div>
            </div>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStudyRoomStore } from '../store/studyRoomStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { User } from '@element-plus/icons-vue'

const router = useRouter()
const studyRoomStore = useStudyRoomStore()

const loading = ref(false)
const userCount = ref(256) // 模拟数据
const roomCount = ref(32)  // 模拟数据
const totalStudyTime = ref(1209600) // 模拟数据：总学习时间（分钟）

// 模拟的用户排行榜数据
const leaderboard = ref([
  { id: 1, username: '学霸一号', avatar: '', studyTime: 18000 },
  { id: 2, username: '勤奋学子', avatar: '', studyTime: 15600 },
  { id: 3, username: '夜猫子', avatar: '', studyTime: 12800 },
  { id: 4, username: '早起鸟', avatar: '', studyTime: 10500 },
  { id: 5, username: '专注达人', avatar: '', studyTime: 9200 }
])

const publicRooms = computed(() => 
  studyRoomStore.rooms.filter(room => !room.isPrivate && room.ownerId !== 1)
)

const myRooms = computed(() => studyRoomStore.myRooms)

onMounted(async () => {
  await fetchRooms()
})

const fetchRooms = async () => {
  loading.value = true
  try {
    await studyRoomStore.fetchRooms()
  } catch (error) {
    ElMessage.error('获取自习室列表失败')
  } finally {
    loading.value = false
  }
}

const handleCreateRoom = () => {
  // 这里触发创建自习室的对话框，可以通过事件总线或其他方式与HeaderComponent通信
  // 简单起见，这里我们直接显示提示信息
  ElMessage.info('请使用顶部的创建自习室按钮')
}

const handleJoinRoom = (room: any) => {
  if (room.currentMembers >= room.maxMembers) {
    ElMessage.warning('该自习室已满员')
    return
  }
  
  router.push({ name: 'StudyRoom', params: { id: room.id } })
}

const formatTime = (minutes: number) => {
  const hours = Math.floor(minutes / 60)
  const days = Math.floor(hours / 24)
  
  if (days > 0) {
    return `${days}天${hours % 24}小时`
  } else if (hours > 0) {
    return `${hours}小时${minutes % 60}分钟`
  }
  
  return `${minutes}分钟`
}
</script>

<style scoped>
.home-container {
  padding: 20px;
}

.welcome-section {
  text-align: center;
  margin-bottom: 40px;
}

.welcome-section h1 {
  font-size: 28px;
  margin-bottom: 10px;
  color: var(--text-color);
}

.welcome-section p {
  font-size: 16px;
  color: var(--text-color-secondary);
  max-width: 600px;
  margin: 0 auto 30px;
}

.stats-card {
  display: flex;
  justify-content: center;
  gap: 40px;
  margin: 30px auto;
  max-width: 800px;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
}

.stat-item {
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 600;
  color: var(--primary-color);
  margin-bottom: 5px;
}

.stat-label {
  font-size: 14px;
  color: var(--text-color-secondary);
}

.section-heading {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.section-heading h2 {
  margin: 0;
  font-size: 20px;
  color: var(--text-color);
}

.room-card {
  height: 100%;
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.room-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.room-theme {
  height: 80px;
  border-radius: 4px 4px 0 0;
  background-color: #e0e0e0;
  background-size: cover;
  background-position: center;
}

.theme-default {
  background-color: #e0e0e0;
}

.theme-focus {
  background-color: #409EFF;
}

.theme-forest {
  background-color: #67C23A;
}

.theme-ocean {
  background-color: #409EFF;
}

.theme-cafe {
  background-color: #E6A23C;
}

.room-info {
  padding: 15px 0;
}

.room-info h3 {
  margin: 0 0 10px;
  font-size: 16px;
  color: var(--text-color);
}

.room-description {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin-bottom: 10px;
  height: 40px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.room-meta {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
  color: var(--text-color-secondary);
}

.join-btn {
  margin-top: 10px;
  width: 100%;
}

.side-section {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
  margin-bottom: 20px;
}

.side-section h2 {
  margin: 0 0 15px;
  font-size: 18px;
  color: var(--text-color);
}

.my-rooms {
  margin-bottom: 20px;
}

.my-room-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid var(--border-color-light);
}

.my-room-item:last-child {
  border-bottom: none;
}

.room-name {
  font-size: 14px;
  color: var(--text-color);
  max-width: 70%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.leaderboard-item {
  display: flex;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid var(--border-color-light);
}

.leaderboard-item:last-child {
  border-bottom: none;
}

.rank {
  width: 24px;
  font-weight: 600;
  color: var(--text-color);
}

.user-info {
  margin-left: 10px;
  flex: 1;
}

.username {
  font-size: 14px;
  color: var(--text-color);
  margin-bottom: 2px;
}

.study-time {
  font-size: 12px;
  color: var(--text-color-secondary);
}
</style> 