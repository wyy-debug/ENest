<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { User } from '@element-plus/icons-vue'
import { wsClient } from '../utils/websocket'
import { MessageType } from '../proto/message'
import { WS_CONFIG } from '../config/config'

const router = useRouter()
const formRef = ref()
const isLogin = ref(true)
const loading = ref(false)

const form = ref({
  email: '',
  username: '',
  password: '',
  confirmPassword: ''
})

const rules = {
  email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }],
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (_rule: any, value: string, callback: Function) => {
        if (value !== form.value.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

onMounted(() => {
  //checkLoginStatus()
  //initWebSocket()
})

const checkLoginStatus = async () => {
  const token = localStorage.getItem('session_token')
  const userData = localStorage.getItem('user_data')
  if (token && userData) {
    router.push('/main')
  }
}

const initWebSocket = () => {
  wsClient.connect(WS_CONFIG.SERVER_URL)
  wsClient.registerHandler(MessageType.AUTH, handleAuthResponse)
  wsClient.registerHandler(MessageType.ERROR, handleErrorResponse)
}

const handleAuthResponse = (payload: Uint8Array) => {
  const response = JSON.parse(new TextDecoder().decode(payload))
  const { token, user } = response
  if (token == 'undefined' || token == null || !user) {
    ElMessage.error('认证响应数据不完整')
    loading.value = false
    return
  }
  localStorage.setItem('session_token', token)
  localStorage.setItem('user_data', JSON.stringify(user))
  localStorage.setItem('session_id', user.id)
  ElMessage.success(isLogin.value ? '登录成功' : '注册成功')
  router.push('/main')
  loading.value = false
}

const handleErrorResponse = (payload: Uint8Array) => {
  const error = JSON.parse(new TextDecoder().decode(payload))
  ElMessage.error(error.message || '操作失败')
  loading.value = false
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    loading.value = true

    const authData = isLogin.value
      ? {
          username: form.value.username,
          password: form.value.password
        }
      : {
          username: form.value.username,
          password: form.value.password,
          email: form.value.email
        }

    const payload = new TextEncoder().encode(JSON.stringify(authData))
    initWebSocket()
    wsClient.sendMessage(MessageType.AUTH, payload)
  } catch (error: any) {
    ElMessage.error(error.message || '表单验证失败')
    loading.value = false
  }
}

const toggleMode = () => {
  isLogin.value = !isLogin.value
  formRef.value?.resetFields()
  form.value = {
    email: '',
    username: '',
    password: '',
    confirmPassword: ''
  }
}
</script>

<template>
  <div class="login-container">
    <el-card class="login-card">
      <div class="login-header">
        <h2>{{ isLogin ? '登录' : '注册' }}</h2>
      </div>
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
        @submit.prevent
      >
        <el-form-item
          v-if="!isLogin"
          label="邮箱"
          prop="email"
        >
          <el-input
            v-model="form.email"
            placeholder="请输入邮箱"
            :prefix-icon="User"
          />
        </el-form-item>

        <el-form-item
          label="用户名"
          prop="username"
        >
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            :prefix-icon="User"
          />
        </el-form-item>

        <el-form-item
          label="密码"
          prop="password"
        >
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            show-password
          />
        </el-form-item>

        <el-form-item
          v-if="!isLogin"
          label="确认密码"
          prop="confirmPassword"
        >
          <el-input
            v-model="form.confirmPassword"
            type="password"
            placeholder="请确认密码"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            class="submit-button"
            @click="handleSubmit"
          >
            {{ isLogin ? '登录' : '注册' }}
          </el-button>
        </el-form-item>

        <div class="toggle-mode">
          <el-button
            type="text"
            @click="toggleMode"
          >
            {{ isLogin ? '没有账号？立即注册' : '已有账号？立即登录' }}
          </el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fa;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 20px;
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-header h2 {
  margin: 0;
  color: #333;
}

.submit-button {
  width: 100%;
}

.toggle-mode {
  text-align: center;
  margin-top: 16px;
}
</style>