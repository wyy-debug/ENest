<template>
  <div class="login-container">
    <h2 class="title">登录</h2>
    
    <el-form
      ref="formRef"
      :model="loginForm"
      :rules="rules"
      label-position="top"
      size="large"
      @submit.prevent="handleLogin"
    >
      <el-form-item label="邮箱" prop="email">
        <el-input 
          v-model="loginForm.email" 
          placeholder="请输入邮箱"
          prefix-icon="Message"
        />
      </el-form-item>
      
      <el-form-item label="密码" prop="password">
        <el-input 
          v-model="loginForm.password" 
          type="password" 
          placeholder="请输入密码"
          prefix-icon="Lock"
          show-password
        />
      </el-form-item>
      
      <div class="remember-forgot">
        <el-checkbox v-model="loginForm.remember">记住我</el-checkbox>
        <el-link type="primary" :underline="false">忘记密码？</el-link>
      </div>
      
      <el-form-item>
        <el-button 
          type="primary" 
          native-type="submit" 
          :loading="loading"
          class="submit-btn"
        >
          登录
        </el-button>
      </el-form-item>
    </el-form>
    
    <div class="register-link">
      还没有账号？ <router-link to="/register">立即注册</router-link>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '../store/userStore'
import { ElMessage, FormInstance } from 'element-plus'
import { Lock, Message } from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const formRef = ref<FormInstance>()
const loading = ref(false)

const loginForm = reactive({
  email: '',
  password: '',
  remember: false
})

const rules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少为6个字符', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    loading.value = true
    
    try {
      // 使用userStore统一处理登录逻辑
      const result = await userStore.login(loginForm.email, loginForm.password)
      
      if (!result.success) {
        throw new Error(result.message || '登录失败')
      }
      
      ElMessage.success('登录成功')
      
      // 导航到指定页面
      const redirectPath = route.query.redirect as string || '/'
      router.push(redirectPath)
    } catch (error: any) {
      console.error('登录失败:', error)
      ElMessage.error(error.message || '登录失败，请稍后再试')
    } finally {
      loading.value = false
    }
  })
}
</script>

<style scoped>
.login-container {
  width: 100%;
}

.title {
  text-align: center;
  margin-bottom: 30px;
  color: var(--text-color);
}

.remember-forgot {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.submit-btn {
  width: 100%;
  margin-top: 10px;
}

.register-link {
  text-align: center;
  margin-top: 20px;
  font-size: 14px;
  color: var(--text-color-secondary);
}

.register-link a {
  color: var(--primary-color);
  text-decoration: none;
}

.register-link a:hover {
  text-decoration: underline;
}
</style> 