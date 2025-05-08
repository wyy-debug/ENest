<template>
  <div class="not-found">
    <div class="container">
      <h1>404</h1>
      <h2>页面不存在</h2>
      <p>您尝试访问的页面可能已经被移除、名称已更改或暂时不可用</p>
      <el-button type="primary" @click="goHome">返回首页</el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.not-found {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  background-color: var(--el-bg-color);
}

.container {
  max-width: 500px;
  padding: 2rem;
}

h1 {
  font-size: 120px;
  color: var(--el-color-primary);
  margin: 0;
  line-height: 1;
}

h2 {
  font-size: 32px;
  margin-bottom: 20px;
}

p {
  font-size: 18px;
  color: var(--el-text-color-secondary);
  margin-bottom: 30px;
}
</style> 