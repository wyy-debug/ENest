<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="title">登录</h1>
      <p class="subtitle">欢迎回到NewENest自习室</p>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="email">邮箱</label>
          <input
            type="email"
            id="email"
            v-model="email"
            placeholder="请输入邮箱"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="password">密码</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="请输入密码"
            required
          />
        </div>
        
        <div class="form-options">
          <div class="remember-me">
            <input type="checkbox" id="remember" v-model="remember" />
            <label for="remember">记住我</label>
          </div>
          <router-link to="/forgot-password" class="forgot-password">忘记密码?</router-link>
        </div>
        
        <button type="submit" class="login-btn" :disabled="isLoading">
          {{ isLoading ? '登录中...' : '登录' }}
        </button>
        
        <div v-if="error" class="error-message">{{ error }}</div>
      </form>
      
      <div class="divider">
        <span>或</span>
      </div>
      
      <div class="social-login">
        <button class="social-btn wechat">
          <i class="icon wechat-icon"></i>
          微信登录
        </button>
        <button class="social-btn github">
          <i class="icon github-icon"></i>
          GitHub登录
        </button>
      </div>
      
      <div class="register-link">
        还没有账号? <router-link to="/register">立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginView',
  data() {
    return {
      email: '',
      password: '',
      remember: false,
      isLoading: false,
      error: ''
    }
  },
  methods: {
    async handleLogin() {
      try {
        this.isLoading = true;
        this.error = '';
        
        // 调用登录API
        const response = await fetch('http://localhost:8080/api/v1/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
          }),
        });
        
        const data = await response.json();
        
        if (!response.ok) {
          throw new Error(data.message || '登录失败，请检查您的凭据');
        }
        
        // 保存token和用户信息
        const { token, user } = data.data;
        localStorage.setItem('token', token);
        localStorage.setItem('user', JSON.stringify(user));
        
        // 重定向到首页
        this.$router.push('/dashboard');
      } catch (err) {
        this.error = err.message || '登录时发生错误，请稍后再试';
        console.error('登录错误:', err);
      } finally {
        this.isLoading = false;
      }
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
  background-color: #f5f7fa;
}

.login-card {
  width: 100%;
  max-width: 450px;
  padding: 2.5rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.title {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
  margin-bottom: 0.5rem;
}

.subtitle {
  font-size: 1rem;
  color: #666;
  margin-bottom: 2rem;
}

.login-form {
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #333;
}

.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

.form-group input:focus {
  border-color: #4e54c8;
  outline: none;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.remember-me {
  display: flex;
  align-items: center;
}

.remember-me input {
  margin-right: 0.5rem;
}

.forgot-password {
  color: #4e54c8;
  text-decoration: none;
}

.login-btn {
  width: 100%;
  padding: 0.85rem;
  background: linear-gradient(90deg, #4e54c8, #8f94fb);
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.3s;
}

.login-btn:hover {
  opacity: 0.9;
}

.login-btn:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.error-message {
  color: #e74c3c;
  margin-top: 1rem;
  text-align: center;
}

.divider {
  display: flex;
  align-items: center;
  margin: 2rem 0;
}

.divider::before,
.divider::after {
  content: "";
  flex: 1;
  border-bottom: 1px solid #eee;
}

.divider span {
  padding: 0 1rem;
  color: #999;
}

.social-login {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.social-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem;
  border: 1px solid #eee;
  border-radius: 6px;
  background: white;
  cursor: pointer;
  transition: background-color 0.3s;
}

.social-btn:hover {
  background-color: #f5f7fa;
}

.wechat {
  color: #07c160;
}

.github {
  color: #333;
}

.icon {
  margin-right: 0.5rem;
  width: 1.2rem;
  height: 1.2rem;
  display: inline-block;
}

.register-link {
  text-align: center;
  color: #666;
}

.register-link a {
  color: #4e54c8;
  text-decoration: none;
  font-weight: 500;
}
</style> 