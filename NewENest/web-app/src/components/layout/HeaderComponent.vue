<template>
  <div class="header-container">
    <div class="logo-container">
      <router-link to="/">
        <img src="../../assets/logo.svg" alt="NewENest Logo" class="logo" />
        <span class="logo-text">NewENest</span>
      </router-link>
    </div>
    
    <div class="search-container">
      <el-input
        v-model="searchQuery"
        placeholder="搜索自习室..."
        prefix-icon="el-icon-search"
        clearable
        @keyup.enter="handleSearch"
      />
    </div>
    
    <div class="action-container">
      <el-button type="primary" @click="handleCreateRoom">创建自习室</el-button>
      
      <el-dropdown trigger="click">
        <div class="user-dropdown">
          <el-avatar :src="userInfo?.avatar" :size="32">{{ userInfo?.username?.charAt(0) }}</el-avatar>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="goToProfile">个人中心</el-dropdown-item>
            <el-dropdown-item @click="goToFriends">好友列表</el-dropdown-item>
            <el-dropdown-item @click="goToDashboard">学习数据</el-dropdown-item>
            <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
  
  <el-dialog
    v-model="createRoomDialogVisible"
    title="创建自习室"
    width="500px"
  >
    <el-form :model="roomForm" label-width="80px">
      <el-form-item label="房间名称" required>
        <el-input v-model="roomForm.name" placeholder="请输入自习室名称" />
      </el-form-item>
      <el-form-item label="房间描述">
        <el-input v-model="roomForm.description" type="textarea" placeholder="请输入自习室描述" />
      </el-form-item>
      <el-form-item label="最大人数">
        <el-input-number v-model="roomForm.maxMembers" :min="1" :max="50" />
      </el-form-item>
      <el-form-item label="访问权限">
        <el-radio-group v-model="roomForm.isPrivate">
          <el-radio :label="false">公开</el-radio>
          <el-radio :label="true">私有</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="房间主题">
        <el-select v-model="roomForm.theme" placeholder="请选择主题">
          <el-option label="默认" value="default" />
          <el-option label="专注" value="focus" />
          <el-option label="森林" value="forest" />
          <el-option label="海洋" value="ocean" />
          <el-option label="咖啡厅" value="cafe" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="createRoomDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createRoom" :loading="creating">创建</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../store/userStore'
import { useStudyRoomStore } from '../../store/studyRoomStore'
import { ElMessage } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()
const studyRoomStore = useStudyRoomStore()

const searchQuery = ref('')
const createRoomDialogVisible = ref(false)
const creating = ref(false)

const userInfo = computed(() => userStore.userInfo)

const roomForm = ref({
  name: '',
  description: '',
  maxMembers: 20,
  isPrivate: false,
  theme: 'default'
})

const handleSearch = () => {
  if (!searchQuery.value.trim()) return
  router.push({ name: 'Home', query: { search: searchQuery.value } })
}

const handleCreateRoom = () => {
  createRoomDialogVisible.value = true
}

const createRoom = async () => {
  if (!roomForm.value.name.trim()) {
    ElMessage.warning('自习室名称不能为空')
    return
  }
  
  creating.value = true
  const result = await studyRoomStore.createRoom({
    name: roomForm.value.name,
    description: roomForm.value.description,
    maxMembers: roomForm.value.maxMembers,
    isPrivate: roomForm.value.isPrivate,
    theme: roomForm.value.theme
  })
  creating.value = false
  
  if (result.success) {
    ElMessage.success('自习室创建成功')
    createRoomDialogVisible.value = false
    router.push({ name: 'StudyRoom', params: { id: result.data.id } })
  } else {
    ElMessage.error(result.message)
  }
}

const goToProfile = () => {
  router.push({ name: 'Profile' })
}

const goToFriends = () => {
  router.push({ name: 'Friends' })
}

const goToDashboard = () => {
  router.push({ name: 'Dashboard' })
}

const handleLogout = () => {
  userStore.logout()
  router.push({ name: 'Login' })
  ElMessage.success('已成功退出登录')
}
</script>

<style scoped>
.header-container {
  height: 100%;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.logo-container {
  display: flex;
  align-items: center;
}

.logo-container a {
  display: flex;
  align-items: center;
  text-decoration: none;
}

.logo {
  width: 32px;
  height: 32px;
}

.logo-text {
  margin-left: 10px;
  font-size: 18px;
  font-weight: 600;
  color: var(--primary-color);
}

.search-container {
  flex: 1;
  max-width: 400px;
  margin: 0 20px;
}

.action-container {
  display: flex;
  align-items: center;
}

.user-dropdown {
  margin-left: 15px;
  cursor: pointer;
}

:deep(.el-input__wrapper) {
  padding-right: 12px;
}

:deep(.el-input__inner) {
  height: 36px;
  line-height: 36px;
}
</style> 