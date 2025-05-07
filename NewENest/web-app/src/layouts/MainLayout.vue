<template>
  <div class="main-layout">
    <el-container>
      <el-header height="60px">
        <header-component />
      </el-header>
      <el-container>
        <el-aside width="240px" v-if="showSidebar">
          <sidebar-component />
        </el-aside>
        <el-main>
          <router-view v-slot="{ Component }">
            <transition name="page">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import HeaderComponent from '../components/layout/HeaderComponent.vue'
import SidebarComponent from '../components/layout/SidebarComponent.vue'

const route = useRoute()

// 根据路由决定是否显示侧边栏
const showSidebar = computed(() => {
  // 在某些页面不显示侧边栏，例如自习室页面
  return route.name !== 'StudyRoom'
})
</script>

<style scoped>
.main-layout {
  height: 100%;
}

.el-header {
  padding: 0;
  border-bottom: 1px solid var(--border-color-light);
  background-color: #fff;
}

.el-aside {
  background-color: #fff;
  border-right: 1px solid var(--border-color-light);
}

.el-main {
  padding: 20px;
  background-color: var(--background-color);
}

.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style> 