<template>
  <div class="study-room-container" :class="`theme-${room?.theme || 'default'}`">
    <div class="room-header">
      <div class="room-info">
        <h1>{{ room?.name }}</h1>
        <p>{{ room?.description }}</p>
      </div>
      
      <div class="room-actions">
        <el-button @click="handleLeaveRoom" type="danger" plain>离开自习室</el-button>
      </div>
    </div>
    
    <div class="room-content">
      <div class="main-area">
        <div class="study-timer">
          <div class="timer-display">
            <div class="time">{{ formatTime(studyTime) }}</div>
            <div class="timer-status">{{ timerStatus }}</div>
          </div>
          
          <div class="timer-controls">
            <el-button 
              type="primary" 
              @click="startTimer" 
              :disabled="isRunning"
              round
            >
              <el-icon><VideoPlay /></el-icon>
              开始
            </el-button>
            
            <el-button 
              type="warning" 
              @click="pauseTimer" 
              :disabled="!isRunning"
              round
            >
              <el-icon><VideoPause /></el-icon>
              暂停
            </el-button>
            
            <el-button 
              type="danger" 
              @click="stopTimer" 
              :disabled="studyTime === 0"
              round
            >
              <el-icon><CircleClose /></el-icon>
              结束
            </el-button>
          </div>
        </div>
        
        <div class="focus-area">
          <h2>今日专注记录</h2>
          <el-timeline>
            <el-timeline-item 
              v-for="(record, index) in studyRecords" 
              :key="index"
              :timestamp="record.time"
              :type="record.type"
            >
              {{ record.text }}
            </el-timeline-item>
            
            <el-empty v-if="studyRecords.length === 0" description="今日暂无学习记录" />
          </el-timeline>
        </div>
      </div>
      
      <div class="side-panel">
        <div class="members-panel">
          <h2>自习室成员 ({{ room?.members?.length || 0 }}/{{ room?.maxMembers || 0 }})</h2>
          
          <div class="members-list">
            <div 
              v-for="member in room?.members" 
              :key="member.id" 
              class="member-item"
              :class="{ 'online': member.status === 'online', 'away': member.status === 'away' }"
            >
              <el-avatar size="small" :src="member.avatar">
                {{ member.username?.charAt(0) }}
              </el-avatar>
              <div class="member-info">
                <div class="member-name">
                  {{ member.username }} 
                  <span v-if="member.role === 'owner'" class="member-role">(房主)</span>
                  <span v-else-if="member.role === 'admin'" class="member-role">(管理员)</span>
                </div>
                <div class="member-status">
                  {{ member.status === 'online' ? '在线' : 
                     member.status === 'away' ? '离开' : '离线' }}
                </div>
              </div>
            </div>
            
            <el-empty v-if="!room?.members || room.members.length === 0" description="暂无成员" />
          </div>
        </div>
        
        <div class="chat-panel">
          <h2>聊天</h2>
          
          <div class="chat-messages" ref="chatContainer">
            <div v-for="(msg, index) in chatMessages" :key="index" 
                class="message" 
                :class="{ 'message-self': msg.senderId === currentUserId }">
              <div class="message-sender" v-if="msg.senderId !== currentUserId">
                {{ getMemberName(msg.senderId) }}
              </div>
              <div class="message-content">{{ msg.content }}</div>
              <div class="message-time">{{ msg.time }}</div>
            </div>
            
            <div v-if="chatMessages.length === 0" class="empty-chat">
              <el-empty description="暂无消息" />
            </div>
          </div>
          
          <div class="chat-input">
            <el-input
              v-model="messageInput"
              placeholder="输入消息..."
              @keyup.enter="sendMessage"
            >
              <template #append>
                <el-button @click="sendMessage">发送</el-button>
              </template>
            </el-input>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStudyRoomStore } from '../store/studyRoomStore'
import { ElMessage, ElMessageBox } from 'element-plus'
import { VideoPlay, VideoPause, CircleClose } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const studyRoomStore = useStudyRoomStore()

const roomId = computed(() => Number(route.params.id))
const room = ref<any>(null)
const loading = ref(false)
const currentUserId = 1 // 模拟当前用户ID，实际应从用户存储获取

// 学习计时器
const studyTime = ref(0) // 单位：秒
const timerInterval = ref<any>(null)
const isRunning = ref(false)
const timerStatus = computed(() => {
  if (studyTime.value === 0) return '准备开始'
  if (isRunning.value) return '专注中...'
  return '已暂停'
})

const studyRecords = ref([
  {
    time: '09:30',
    type: 'success',
    text: '开始学习，专注30分钟'
  },
  {
    time: '10:15',
    type: 'warning',
    text: '暂停学习'
  },
  {
    time: '10:20',
    type: 'success',
    text: '继续学习，专注40分钟'
  }
])

// 聊天
const chatMessages = ref([
  {
    senderId: 2,
    content: '大家好，今天有什么学习计划？',
    time: '09:15'
  },
  {
    senderId: 1,
    content: '我计划学习Vue 3和TypeScript',
    time: '09:16'
  },
  {
    senderId: 3,
    content: '我正在复习数据结构与算法',
    time: '09:18'
  }
])
const messageInput = ref('')
const chatContainer = ref<HTMLElement | null>(null)

onMounted(async () => {
  await fetchRoomData()
  
  // 聊天窗口滚动到底部
  scrollToBottom()
})

onBeforeUnmount(() => {
  // 清除计时器
  if (timerInterval.value) {
    clearInterval(timerInterval.value)
  }
  
  // 如果有未保存的学习时间，则保存
  if (studyTime.value > 0) {
    saveStudyRecord()
  }
})

const fetchRoomData = async () => {
  loading.value = true
  try {
    const data = await studyRoomStore.fetchRoomById(roomId.value)
    if (data) {
      room.value = data
    } else {
      ElMessage.error('自习室不存在或已关闭')
      router.push('/')
    }
  } catch (error) {
    ElMessage.error('获取自习室信息失败')
  } finally {
    loading.value = false
  }
}

const handleLeaveRoom = () => {
  ElMessageBox.confirm(
    '确定要离开自习室吗？未保存的学习记录将会丢失。',
    '离开提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    try {
      await studyRoomStore.leaveRoom(roomId.value)
      ElMessage.success('已离开自习室')
      router.push('/')
    } catch (error) {
      ElMessage.error('离开自习室失败')
    }
  }).catch(() => {})
}

// 计时器控制
const startTimer = () => {
  if (isRunning.value) return
  
  isRunning.value = true
  timerInterval.value = setInterval(() => {
    studyTime.value++
  }, 1000)
  
  // 添加开始学习记录
  const now = new Date()
  const time = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
  
  studyRecords.value.push({
    time,
    type: 'success',
    text: '开始学习'
  })
}

const pauseTimer = () => {
  if (!isRunning.value) return
  
  isRunning.value = false
  clearInterval(timerInterval.value)
  
  // 添加暂停学习记录
  const now = new Date()
  const time = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
  
  studyRecords.value.push({
    time,
    type: 'warning',
    text: `暂停学习，已专注${formatTime(studyTime.value)}`
  })
}

const stopTimer = () => {
  isRunning.value = false
  clearInterval(timerInterval.value)
  
  // 保存学习记录
  saveStudyRecord()
  
  // 重置计时器
  studyTime.value = 0
}

const saveStudyRecord = () => {
  // 实际应该调用API保存学习记录
  console.log('保存学习记录:', formatTime(studyTime.value))
  
  // 添加结束学习记录
  const now = new Date()
  const time = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
  
  studyRecords.value.push({
    time,
    type: 'info',
    text: `结束学习，本次专注${formatTime(studyTime.value)}`
  })
}

const formatTime = (seconds: number) => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  
  let result = ''
  if (hours > 0) {
    result += `${hours}小时`
  }
  if (minutes > 0 || hours > 0) {
    result += `${minutes}分钟`
  }
  if (hours === 0) {
    result += `${secs}秒`
  }
  
  return result
}

// 聊天功能
const sendMessage = () => {
  if (!messageInput.value.trim()) return
  
  const now = new Date()
  const time = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
  
  chatMessages.value.push({
    senderId: currentUserId,
    content: messageInput.value,
    time
  })
  
  messageInput.value = ''
  
  // 滚动到底部
  nextTick(() => {
    scrollToBottom()
  })
}

const scrollToBottom = () => {
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}

const getMemberName = (id: number) => {
  const member = room.value?.members?.find((m: any) => m.userId === id)
  return member ? member.username : '未知用户'
}
</script>

<style scoped>
.study-room-container {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.room-header {
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid var(--border-color-light);
  background-color: #fff;
}

.room-info h1 {
  margin: 0 0 5px;
  font-size: 22px;
}

.room-info p {
  margin: 0;
  color: var(--text-color-secondary);
}

.room-content {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.main-area {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

.study-timer {
  background-color: white;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
  padding: 30px;
  margin-bottom: 20px;
  text-align: center;
}

.timer-display {
  margin-bottom: 30px;
}

.timer-display .time {
  font-size: 48px;
  font-weight: 600;
  margin-bottom: 10px;
  color: var(--primary-color);
}

.timer-display .timer-status {
  font-size: 16px;
  color: var(--text-color-secondary);
}

.timer-controls {
  display: flex;
  justify-content: center;
  gap: 20px;
}

.focus-area {
  background-color: white;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
  padding: 20px;
}

.focus-area h2 {
  margin-top: 0;
  margin-bottom: 20px;
  font-size: 18px;
}

.side-panel {
  width: 300px;
  border-left: 1px solid var(--border-color-light);
  display: flex;
  flex-direction: column;
  background-color: white;
}

.members-panel, .chat-panel {
  padding: 20px;
}

.members-panel {
  border-bottom: 1px solid var(--border-color-light);
}

.members-panel h2, .chat-panel h2 {
  margin-top: 0;
  margin-bottom: 15px;
  font-size: 16px;
}

.members-list {
  max-height: 200px;
  overflow-y: auto;
}

.member-item {
  display: flex;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid var(--border-color-light);
}

.member-item:last-child {
  border-bottom: none;
}

.member-info {
  margin-left: 10px;
  flex: 1;
}

.member-name {
  font-size: 14px;
  margin-bottom: 2px;
}

.member-role {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.member-status {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.member-item.online .member-status {
  color: var(--success-color);
}

.member-item.away .member-status {
  color: var(--warning-color);
}

.chat-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  margin-bottom: 15px;
  max-height: 300px;
}

.message {
  margin-bottom: 12px;
  max-width: 80%;
}

.message-self {
  margin-left: auto;
  text-align: right;
}

.message-sender {
  font-size: 12px;
  color: var(--text-color-secondary);
  margin-bottom: 2px;
}

.message-content {
  background-color: #f5f7fa;
  padding: 8px 12px;
  border-radius: 8px;
  display: inline-block;
  word-break: break-word;
}

.message-self .message-content {
  background-color: var(--primary-color);
  color: white;
}

.message-time {
  font-size: 12px;
  color: var(--text-color-secondary);
  margin-top: 2px;
}

.empty-chat {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.theme-default {
  background-color: #f5f7fa;
}

.theme-focus {
  background-color: #f0f9ff;
}

.theme-forest {
  background-color: #f0f9eb;
}

.theme-ocean {
  background-color: #ecf5ff;
}

.theme-cafe {
  background-color: #fdf6ec;
}
</style> 