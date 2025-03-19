<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { User } from '@element-plus/icons-vue'
import axios from '../utils/axios'

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
      validator: (rule: any, value: string, callback: Function) => {
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
  checkLoginStatus()
})

const checkLoginStatus = async () => {
  const token = localStorage.getItem('session_token')
  const userData = localStorage.getItem('user_data')
  if (token && userData) {
    router.push('/study-room')
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    loading.value = true

    const url = isLogin.value ? 'auth/login' : 'auth/register'
    const requestData = isLogin.value
      ? {
          username: form.value.username,
          password: form.value.password
        }
      : {
          username: form.value.username,
          password: form.value.password,
          email: form.value.email
        }

    const response = await axios.post(url, requestData)

    const { token, user } = response.data
    localStorage.setItem('session_token', token)
    localStorage.setItem('user_data', JSON.stringify(user))

    ElMessage.success(isLogin.value ? '登录成功' : '注册成功')
    router.push('/')
  } catch (error: any) {
    ElMessage.error(error.response?.data?.message || '操作失败')
  } finally {
    loading.value = false
  }
}

const toggleMode = () => {
  isLogin.value = !isLogin.value
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
      <h2>{{ isLogin ? '登录' : '注册' }}</h2>
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
      >
        <el-form-item v-if="!isLogin" prop="email">
            <el-input v-model="form.email" placeholder="请输入邮箱">
                <template #prefix>
                    <el-icon><User /></el-icon>
                </template>
            </el-input>
        </el-form-item>

        <el-form-item prop="username">
            <el-input v-model="form.username" placeholder="请输入用户名">
              <template #prefix>
                <el-icon><User /></el-icon>
              </template>
            </el-input>
        </el-form-item>

        <el-form-item prop="password">
            <el-input  v-model="form.password" type="password" placeholder="请输入密码">
                <template #prefix>
                    <el-icon><User /></el-icon>
                </template>
            </el-input>
        </el-form-item>

        <el-form-item v-if="!isLogin" prop="confirmPassword">
            <el-input v-model="form.confirmPassword" type="password" placeholder="请确认密码">
                <template #prefix>
                    <el-icon><User /></el-icon>
                </template>
            </el-input>
        </el-form-item>

        <el-form-item class="submit-button">
            <el-button type="primary" :loading="loading" @click="handleSubmit">
                {{ isLogin ? '登录' : '注册' }}
            </el-button>
        </el-form-item>

        <div class="toggle-mode">
            <el-button link @click="toggleMode">
                {{ isLogin ? '没有账号？立即注册' : '已有账号？立即登录' }}
            </el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.login-container {
  height: 100vh;
  width: 100vw;
  
  overflow: hidden;
}

.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(
    to bottom,
    rgba(16, 163, 127, 0.1),
    rgba(16, 163, 127, 0.2),
    rgba(16, 163, 127, 0.1)
  );
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 20px;
}

.toggle-mode {
  text-align: center;
  margin-top: 16px;
}

.submit-button {
  display: flex;
  justify-content: center;
  width: 100%;
}

.submit-button :deep(.el-form-item__content) {
  display: flex;
  justify-content: center;
  width: 100%;
}

.submit-button :deep(.el-button) {
  width: 100%;
  max-width: 100px;
}
</style>