<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import ZoomView from '../components/ZoomView.vue';

const router = useRouter();
const showWebView = ref(false);
const rooms = ref([
  { id: 1, title: '自习室 1', status: '空闲' },
  { id: 2, title: '自习室 2', status: '空闲' },
  { id: 3, title: '自习室 3', status: '空闲' },
  { id: 4, title: '自习室 4', status: '空闲' }
]);

const userData = ref<any>(null);
const error = ref<string | null>(null);

onMounted(() => {
  const storedUserData = localStorage.getItem('user_data');
  if (!storedUserData) {
    router.push('/login');
    return;
  }
  userData.value = JSON.parse(storedUserData);
});

const handleRoomClick = (room: any) => {
  showWebView.value = true;
};

const handleSessionError = (errorMessage: string) => {
  error.value = errorMessage;
  showWebView.value = false;
};

const handleSessionLeft = () => {
  showWebView.value = false;
};
</script>

<template>
  <div class="study-room-container">
    <h1 class="title">自习室</h1>
    <div v-if="error" class="error-message">
      {{ error }}
      <el-button type="text" @click="error = null">关闭</el-button>
    </div>
    <div v-if="!showWebView" class="room-grid">
      <div v-for="room in rooms" :key="room.id" class="room-card" @click="handleRoomClick(room)">
        <h3>{{ room.title }}</h3>
        <p>状态: {{ room.status }}</p>
      </div>
    </div>
    <div v-else class="webview-container">
      <ZoomView/>
    </div>
  </div>
</template>

<style scoped>
.study-room-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.webview-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.title {
  color: #f97316;
  font-size: 2rem;
  font-weight: bold;
  margin-bottom: 24px;
}

.error-message {
  background-color: #fff2f0;
  border: 1px solid #ffccc7;
  padding: 8px 16px;
  border-radius: 4px;
  margin-bottom: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.room-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  padding: 16px 0;
}

.room-card {
  background: white;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: transform 0.2s;
  aspect-ratio: 1.5;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.room-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.room-card h3 {
  margin: 0 0 8px 0;
  color: #333;
}

.room-card p {
  margin: 0;
  color: #666;
}



.placeholder {
  color: #666;
  font-size: 1.2rem;
}
</style>