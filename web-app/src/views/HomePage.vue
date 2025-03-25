<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import TypewriterText from '../components/TypewriterText.vue'
import { wsClient } from '../utils/websocket'
import { MessageType } from '../proto/message'
import { WS_CONFIG } from '../config/config'

const router = useRouter()
const showHoneycomb = ref(false)
const username = ref<string | null>(null)
const isLoggedIn = ref(false)

const initWebSocket = () => {
  wsClient.connect(WS_CONFIG.SERVER_URL)
  wsClient.registerHandler(MessageType.ERROR, (payload: Uint8Array) => {
    try {
      const errorMessage = JSON.parse(new TextDecoder().decode(payload))
      console.error('WebSocket error:', errorMessage.message || '未知错误')
    } catch (error) {
      console.error('WebSocket error: 解析错误消息失败', error)
    }
  })
}

const checkLoginStatus = async () => {
  const userData = localStorage.getItem('user_data')
  const token = localStorage.getItem('session_token')
  if (userData && token) {
    const user = JSON.parse(userData)
    username.value = user.username
    isLoggedIn.value = true
    router.push('/main')
  }
}

const handleLogout = async () => {
  localStorage.removeItem('user_data')
  localStorage.removeItem('session_token')
  username.value = null
  isLoggedIn.value = false
}

const onTypingFinished = () => {
  showHoneycomb.value = true
}

onMounted(() => {
  initWebSocket()
  checkLoginStatus()
})
</script>

<template>
  <el-container>
    <el-header class="header">
      <div class="logo">
        <h1>E - StudyRoom</h1>
      </div>
      <div class="nav-actions">
        <template v-if="isLoggedIn">
          <span class="username">{{ username }}</span>
          <el-button @click="handleLogout">Logout</el-button>
        </template>
        <el-button v-else @click="router.push('/login')">Log in</el-button>
      </div>
    </el-header>

    <el-main class="main-content">
      <div class="gradient-background">
        <div class="content-container">
          <TypewriterText
            text="Build Your Future, One Study Session at a Time."
            @typing-finished="onTypingFinished"
          />
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>
.el-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 40px;
  background-color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  height: 60px;
  position: relative;
  z-index: 1;
}

.logo h1 {
  font-size: 24px;
  font-weight: bold;
  margin: 0;
  text-rendering: optimizeLegibility;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.username {
  color: #666;
}

.main-content {
  height: calc(100vh - 60px);
  padding: 0;
  overflow: hidden;
  position: relative;
}

.gradient-background {
  height: 100%;
  width: 100%;
  background: linear-gradient(
    to bottom,
    rgba(16, 163, 127, 0.1),
    rgba(16, 163, 127, 0.2),
    rgba(16, 163, 127, 0.1)
  );
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-y: auto;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.content-container {
  max-width: min(100vh, 1200px);
  text-align: center;
  padding: 20px;
  margin: auto;
  user-select: none;
  -webkit-user-select: none;
}

::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: rgba(16, 163, 127, 0.2);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(16, 163, 127, 0.4);
}
</style>