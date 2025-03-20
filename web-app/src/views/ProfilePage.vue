<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from '../utils/axios'

const router = useRouter()
const userData = ref<any>(null)
const loading = ref(false)

onMounted(() => {
  const storedUserData = localStorage.getItem('user_data')
  if (!storedUserData) {
    router.push('/login')
    return
  }
  userData.value = JSON.parse(storedUserData)
  fetchUserProfile()
})

const fetchUserProfile = async () => {
  loading.value = true
  try {
    const response = await axios.get('/profile')
    userData.value = { ...userData.value, ...response.data }
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '获取个人信息失败')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="profile-container">
    <div class="profile-header">
      <h2>个人信息</h2>
    </div>
    <div class="profile-content" v-loading="loading">
      <el-card class="profile-card" v-if="userData">
        <div class="user-info">
          <div class="avatar-section">
            <el-avatar :size="100" :src="userData.avatar">
              {{ userData.username?.charAt(0).toUpperCase() }}
            </el-avatar>
          </div>
          <div class="info-section">
            <div class="info-item">
              <span class="label">用户名：</span>
              <span class="value">{{ userData.username }}</span>
            </div>
            <div class="info-item">
              <span class="label">邮箱：</span>
              <span class="value">{{ userData.email }}</span>
            </div>
          </div>
        </div>
      </el-card>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  padding: 20px;
  height: 100%;
  overflow-y: auto;
}

.profile-header {
  margin-bottom: 20px;
}

.profile-header h2 {
  margin: 0;
  color: #333;
}

.profile-content {
  max-width: 800px;
  margin: 0 auto;
}

.profile-card {
  background: white;
  border-radius: 8px;
}

.user-info {
  display: flex;
  padding: 20px;
  gap: 40px;
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.info-section {
  flex: 1;
  padding: 20px 0;
}

.info-item {
  margin-bottom: 16px;
}

.info-item .label {
  color: #666;
  margin-right: 8px;
}

.info-item .value {
  color: #333;
  font-weight: 500;
}
</style>