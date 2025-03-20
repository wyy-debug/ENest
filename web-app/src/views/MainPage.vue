<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeIndex = ref('1')
const isCollapse = ref(false)

const handleSelect = (key: string) => {
  if (key === '1') {
    router.push('/main/study-room')
  } else if (key === '2') {
    router.push('/main/profile')
  } else if (key === '3') {
    router.push('/main/community')
  } else if (key === '4') {
    router.push('/main/friends')
  }
}

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}
</script>

<template>
  <el-container class="main-container">
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <el-menu
        :default-active="activeIndex"
        class="el-menu-vertical"
        :collapse="isCollapse"
        @select="handleSelect"
      >
        <el-menu-item index="1">
          <el-icon><img src="../assets/home.png" style="width: 20px; height: 20px;" /></el-icon>
          <span>自习室</span>
        </el-menu-item>
        <el-menu-item index="2">
          <el-icon><i-ep-user /></el-icon>
          <span>个人信息</span>
        </el-menu-item>
        <el-menu-item index="3">
          <el-icon><i-ep-message-box /></el-icon>
          <span>社区</span>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon><i-ep-user-filled /></el-icon>
          <span>好友</span>
        </el-menu-item>
      </el-menu>
      <div class="collapse-btn" @click="toggleCollapse">
        <el-icon>
          <i-ep-arrow-left v-if="!isCollapse" />
          <i-ep-arrow-right v-else />
        </el-icon>
      </div>
    </el-aside>
    <el-main class="main-content">
      <router-view />
    </el-main>
  </el-container>
</template>

<style scoped>
.main-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  display: flex;
}

.sidebar {
  background-color: #fff;
  border-right: 1px solid #e6e6e6;
  position: relative;
  transition: width 0.3s;
  overflow: hidden;
}

.el-menu-vertical {
  border-right: none;
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
}

.el-menu--collapse {
  width: 64px;
}

.el-menu-item {
  transition: padding-left 0.3s;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 0 20px;
}

.el-menu--collapse .el-menu-item {
  padding: 0 20px !important;
  text-align: center;
  display: flex;
  justify-content: center;
}

.el-menu--collapse .el-menu-item .el-icon {
  margin: 0;
  font-size: 20px;
}

.el-menu-item .el-icon {
  margin-right: 16px;
  display: flex;
  align-items: center;
}

.el-menu--collapse .el-menu-item .el-icon {
  margin-right: 0;
}

.el-menu--collapse .el-menu-item span {
  display: none;
}

.main-content {
  padding: 0;
  background-color: #f5f7fa;
  display: flex;
  justify-content: center;
  align-items: center;
  flex: 1;
  overflow: hidden;
}

.collapse-btn {
  position: absolute;
  top: 50%;
  right: -16px;
  transform: translateY(-50%);
  width: 24px;
  height: 48px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f7fa;
  border-radius: 12px 0 0 12px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: -2px 0 4px rgba(0, 0, 0, 0.1);
  z-index: 1;
}

.collapse-btn:hover {
  background-color: #e6e6e6;
}
</style>