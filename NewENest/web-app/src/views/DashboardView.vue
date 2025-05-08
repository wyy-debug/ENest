<template>
  <div class="dashboard">
    <h1>控制台</h1>
    <p>欢迎回来！这里是你的个人学习控制台。</p>

    <el-row :gutter="20" class="dashboard-cards">
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>累计学习时间</span>
            </div>
          </template>
          <div class="stat-value">{{ formatTime(studyTime) }}</div>
        </el-card>
      </el-col>
      
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>完成契约数</span>
            </div>
          </template>
          <div class="stat-value">{{ contracts }}</div>
        </el-card>
      </el-col>
      
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>今日已学习</span>
            </div>
          </template>
          <div class="stat-value">{{ formatTime(todayStudyTime) }}</div>
        </el-card>
      </el-col>
      
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>好友数量</span>
            </div>
          </template>
          <div class="stat-value">{{ friendCount }}</div>
        </el-card>
      </el-col>
    </el-row>

    <h2>快速操作</h2>
    <el-row :gutter="20" class="quick-actions">
      <el-col :span="8">
        <el-button type="primary" @click="$router.push('/create-room')">
          <el-icon><Plus /></el-icon>
          创建自习室
        </el-button>
      </el-col>
      <el-col :span="8">
        <el-button type="success" @click="$router.push('/friends')">
          <el-icon><ChatDotRound /></el-icon>
          好友社交
        </el-button>
      </el-col>
      <el-col :span="8">
        <el-button type="warning" @click="$router.push('/profile')">
          <el-icon><User /></el-icon>
          个人中心
        </el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { Plus, ChatDotRound, User } from '@element-plus/icons-vue'

// 模拟数据
const studyTime = ref(5400) // 1小时30分钟
const contracts = ref(3)
const todayStudyTime = ref(1800) // 30分钟
const friendCount = ref(12)

// 工具函数
const formatTime = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  
  if (hours > 0) {
    return `${hours}小时${minutes}分钟`
  }
  return `${minutes}分钟`
}
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.dashboard-cards {
  margin-bottom: 30px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  text-align: center;
  color: var(--el-color-primary);
}

.quick-actions {
  margin-top: 20px;
}

.el-button {
  width: 100%;
  margin-bottom: 10px;
}

h2 {
  margin-top: 30px;
  margin-bottom: 20px;
  font-size: 18px;
  color: var(--el-text-color-primary);
}
</style> 