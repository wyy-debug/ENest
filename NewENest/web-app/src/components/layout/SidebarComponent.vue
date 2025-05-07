<template>
  <el-menu
    :default-active="activeMenu"
    class="sidebar-menu"
    router
  >
    <el-menu-item index="/">
      <el-icon><House /></el-icon>
      <span>首页</span>
    </el-menu-item>
    
    <el-menu-item index="/dashboard">
      <el-icon><DataLine /></el-icon>
      <span>学习数据</span>
    </el-menu-item>
    
    <el-sub-menu index="study-rooms">
      <template #title>
        <el-icon><Reading /></el-icon>
        <span>我的自习室</span>
      </template>
      
      <el-menu-item 
        v-for="room in myRooms" 
        :key="room.id" 
        :index="`/study-room/${room.id}`"
      >
        {{ room.name }}
      </el-menu-item>
      
      <el-menu-item v-if="myRooms.length === 0" disabled>
        暂无自习室
      </el-menu-item>
      
      <el-divider />
      
      <el-menu-item @click="handleCreateRoom">
        <el-icon><Plus /></el-icon>
        <span>创建自习室</span>
      </el-menu-item>
    </el-sub-menu>
    
    <el-menu-item index="/friends">
      <el-icon><User /></el-icon>
      <span>好友列表</span>
    </el-menu-item>
    
    <el-menu-item index="/profile">
      <el-icon><Setting /></el-icon>
      <span>个人设置</span>
    </el-menu-item>
  </el-menu>
  
  <div class="sidebar-footer">
    <div class="study-stats">
      <p>今日学习: {{ formatTime(todayStudyTime) }}</p>
      <p>本周学习: {{ formatTime(weekStudyTime) }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStudyRoomStore } from '../../store/studyRoomStore'
import { House, DataLine, Reading, User, Setting, Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const studyRoomStore = useStudyRoomStore()

// 模拟数据 - 实际项目中应该从 API 获取
const todayStudyTime = ref(120) // 分钟
const weekStudyTime = ref(720) // 分钟

const activeMenu = computed(() => route.path)
const myRooms = computed(() => studyRoomStore.myRooms || [])

onMounted(async () => {
  try {
    await studyRoomStore.fetchRooms()
  } catch (error) {
    ElMessage.error('获取自习室列表失败')
  }
})

const handleCreateRoom = () => {
  // 这里可以直接导航到创建页面，或者通过事件触发顶部组件的弹窗
  ElMessage.info('请使用顶部的创建自习室按钮')
}

const formatTime = (minutes: number) => {
  const hours = Math.floor(minutes / 60)
  const mins = minutes % 60
  
  if (hours > 0) {
    return `${hours}小时${mins > 0 ? `${mins}分钟` : ''}`
  }
  
  return `${mins}分钟`
}
</script>

<style scoped>
.sidebar-menu {
  height: calc(100% - 80px);
  border-right: none;
}

.sidebar-menu .el-menu-item.is-active {
  background-color: var(--primary-color);
  color: #fff;
}

.sidebar-footer {
  position: fixed;
  bottom: 0;
  width: var(--sidebar-width);
  padding: 15px;
  border-top: 1px solid var(--border-color-light);
  background-color: #fff;
}

.study-stats {
  font-size: 14px;
  color: var(--text-color-secondary);
}

.study-stats p {
  margin: 5px 0;
}
</style> 