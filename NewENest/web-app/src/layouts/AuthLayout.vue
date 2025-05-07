<template>
  <div class="auth-layout">
    <div class="auth-container">
      <div class="auth-logo">
        <img src="../assets/logo.svg" alt="NewENest Logo" />
        <h1>NewENest</h1>
      </div>
      <div class="auth-card">
        <router-view v-slot="{ Component }">
          <transition name="fade">
            <component :is="Component" />
          </transition>
        </router-view>
      </div>
      <div class="auth-footer">
        <p>© {{ currentYear }} NewENest - 探索学习的无限可能</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
.auth-layout {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--background-color);
  background-image: linear-gradient(135deg, rgba(64, 158, 255, 0.1) 0%, rgba(103, 194, 58, 0.1) 100%);
}

.auth-container {
  width: 100%;
  max-width: 420px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.auth-logo {
  margin-bottom: 40px;
  text-align: center;
}

.auth-logo img {
  width: 80px;
  height: 80px;
}

.auth-logo h1 {
  margin-top: 16px;
  font-size: 28px;
  color: var(--primary-color);
}

.auth-card {
  width: 100%;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
  padding: 30px;
  margin-bottom: 24px;
}

.auth-footer {
  text-align: center;
  color: var(--text-color-secondary);
  font-size: 14px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style> 