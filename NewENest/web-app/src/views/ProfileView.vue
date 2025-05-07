<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <template #header>
        <div class="card-header">
          <h2>个人资料</h2>
          <el-button type="primary" @click="enableEdit" v-if="!isEditing">编辑资料</el-button>
          <div v-else>
            <el-button type="success" @click="saveProfile" :loading="loading">保存</el-button>
            <el-button @click="cancelEdit">取消</el-button>
          </div>
        </div>
      </template>
      
      <div class="profile-content">
        <div class="avatar-container">
          <el-avatar :size="100" :src="profileForm.avatar">
            {{ profileForm.username?.charAt(0) }}
          </el-avatar>
          <el-upload
            v-if="isEditing"
            class="avatar-uploader"
            action="/api/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <el-button size="small" type="primary" class="upload-btn">更换头像</el-button>
          </el-upload>
        </div>
        
        <div class="info-container">
          <el-form :model="profileForm" label-position="top" :disabled="!isEditing">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="用户名">
                  <el-input v-model="profileForm.username" />
                </el-form-item>
              </el-col>
              
              <el-col :span="12">
                <el-form-item label="邮箱">
                  <el-input v-model="profileForm.email" disabled />
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-form-item label="个性签名">
              <el-input v-model="profileForm.signature" type="textarea" rows="3" />
            </el-form-item>
            
            <el-form-item label="学习方向">
              <el-select v-model="profileForm.studyDirection" style="width: 100%">
                <el-option label="编程开发" value="programming" />
                <el-option label="人工智能" value="ai" />
                <el-option label="数据科学" value="data_science" />
                <el-option label="语言学习" value="language" />
                <el-option label="数学" value="math" />
                <el-option label="物理" value="physics" />
                <el-option label="艺术" value="art" />
                <el-option label="其他" value="other" />
              </el-select>
            </el-form-item>
          </el-form>
        </div>
      </div>
    </el-card>
    
    <el-row :gutter="20" class="stats-row">
      <el-col :span="12">
        <el-card class="stats-card">
          <template #header>
            <div class="card-header">
              <h3>学习统计</h3>
            </div>
          </template>
          
          <div class="stats-content">
            <div class="stat-item">
              <div class="stat-label">累计学习时间</div>
              <div class="stat-value">{{ formatTime(userInfo.totalStudyTime) }}</div>
            </div>
            
            <div class="stat-item">
              <div class="stat-label">本周学习时间</div>
              <div class="stat-value">{{ formatTime(weeklyStudyTime) }}</div>
            </div>
            
            <div class="stat-item">
              <div class="stat-label">平均每日学习</div>
              <div class="stat-value">{{ formatTime(dailyAverage) }}</div>
            </div>
            
            <div class="stat-item">
              <div class="stat-label">学习天数</div>
              <div class="stat-value">{{ studyDays }} 天</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="12">
        <el-card class="achievement-card">
          <template #header>
            <div class="card-header">
              <h3>成就徽章</h3>
            </div>
          </template>
          
          <div class="achievements-content">
            <div v-if="achievements.length > 0" class="achievement-list">
              <div v-for="achievement in achievements" :key="achievement.id" class="achievement-item">
                <el-tooltip
                  :content="achievement.description"
                  placement="top"
                >
                  <div class="achievement-icon">
                    <i :class="achievement.icon"></i>
                  </div>
                </el-tooltip>
                <div class="achievement-name">{{ achievement.name }}</div>
              </div>
            </div>
            
            <el-empty v-else description="暂无成就" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useUserStore } from '../store/userStore'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()
const isEditing = ref(false)
const loading = ref(false)

// 模拟数据
const weeklyStudyTime = ref(1800) // 本周学习时间（分钟）
const studyDays = ref(45) // 学习天数
const dailyAverage = computed(() => Math.round(userStore.userInfo?.totalStudyTime / studyDays.value || 0)) // 平均每日学习时间

const achievements = ref([
  {
    id: 1,
    name: '初学者',
    description: '累计学习时间达到10小时',
    icon: 'el-icon-star-on'
  },
  {
    id: 2,
    name: '坚持不懈',
    description: '连续学习7天',
    icon: 'el-icon-trophy'
  },
  {
    id: 3,
    name: '专注达人',
    description: '单次学习时间超过2小时',
    icon: 'el-icon-time'
  }
])

// 用户信息
const userInfo = computed(() => userStore.userInfo || {})

// 编辑表单
const profileForm = reactive({
  username: '',
  email: '',
  avatar: '',
  signature: '',
  studyDirection: ''
})

onMounted(async () => {
  // 获取用户信息
  if (!userStore.userInfo) {
    await userStore.getUserInfo()
  }
  
  // 初始化表单数据
  Object.assign(profileForm, {
    username: userInfo.value.username || '',
    email: userInfo.value.email || '',
    avatar: userInfo.value.avatar || '',
    signature: userInfo.value.signature || '',
    studyDirection: userInfo.value.studyDirection || 'other'
  })
})

// 开始编辑
const enableEdit = () => {
  isEditing.value = true
}

// 取消编辑
const cancelEdit = () => {
  isEditing.value = false
  
  // 重置表单数据
  Object.assign(profileForm, {
    username: userInfo.value.username || '',
    email: userInfo.value.email || '',
    avatar: userInfo.value.avatar || '',
    signature: userInfo.value.signature || '',
    studyDirection: userInfo.value.studyDirection || 'other'
  })
}

// 保存资料
const saveProfile = async () => {
  if (!profileForm.username.trim()) {
    ElMessage.warning('用户名不能为空')
    return
  }
  
  loading.value = true
  const result = await userStore.updateProfile({
    username: profileForm.username,
    signature: profileForm.signature,
    studyDirection: profileForm.studyDirection
  })
  loading.value = false
  
  if (result.success) {
    ElMessage.success('个人资料更新成功')
    isEditing.value = false
  } else {
    ElMessage.error(result.message)
  }
}

// 头像上传相关函数
const handleAvatarSuccess = (res: any) => {
  profileForm.avatar = res.url
}

const beforeAvatarUpload = (file: File) => {
  const isJPG = file.type === 'image/jpeg'
  const isPNG = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJPG && !isPNG) {
    ElMessage.error('头像只能是 JPG 或 PNG 格式!')
  }
  
  if (!isLt2M) {
    ElMessage.error('头像大小不能超过 2MB!')
  }
  
  return (isJPG || isPNG) && isLt2M
}

// 格式化时间
const formatTime = (minutes: number) => {
  if (!minutes) return '0分钟'
  
  const hours = Math.floor(minutes / 60)
  const mins = minutes % 60
  
  if (hours > 0) {
    return `${hours}小时${mins > 0 ? `${mins}分钟` : ''}`
  }
  
  return `${mins}分钟`
}
</script>

<style scoped>
.profile-container {
  padding: 20px;
}

.profile-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h2, .card-header h3 {
  margin: 0;
}

.profile-content {
  display: flex;
  gap: 30px;
}

.avatar-container {
  text-align: center;
}

.avatar-uploader {
  margin-top: 15px;
}

.upload-btn {
  width: 100%;
}

.info-container {
  flex: 1;
}

.stats-row {
  margin-bottom: 20px;
}

.stats-content, .achievements-content {
  min-height: 200px;
}

.stats-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.stat-item {
  background-color: #f5f7fa;
  border-radius: 8px;
  padding: 15px;
  text-align: center;
}

.stat-label {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin-bottom: 8px;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: var(--primary-color);
}

.achievement-list {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
}

.achievement-item {
  width: 100px;
  text-align: center;
}

.achievement-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: var(--primary-color);
  color: white;
  font-size: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 auto 10px;
}

.achievement-name {
  font-size: 14px;
  color: var(--text-color);
}
</style> 