<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import ZoomView from '../components/ZoomView.vue';

const router = useRouter();
const showWebView = ref(false);
const dialogVisible = ref(false);
const newRoom = ref({
  title: '',
  description: '',
  maxMembers: 10,
  isPrivate: false
});

const rooms = ref([
  { id: 1, title: '自习室 1', description: '安静舒适的学习环境', status: '空闲', onlineCount: 0 },
  { id: 2, title: '自习室 2', description: '适合小组讨论的空间', status: '空闲', onlineCount: 2 },
  { id: 3, title: '自习室 3', description: '24小时开放的自习室', status: '空闲', onlineCount: 1 },
  { id: 4, title: '自习室 4', description: '配备完善设备的学习室', status: '空闲', onlineCount: 3 }
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

const handleCreateRoom = () => {
  router.push('/main/create-study-room');
};

const handleCancel = () => {
  dialogVisible.value = false;
  newRoom.value = {
    title: '',
    description: '',
    maxMembers: 10,
    isPrivate: false
  };
};

const handleConfirm = () => {
  // TODO: 调用后端API创建自习室
  dialogVisible.value = false;
  newRoom.value = {
    title: '',
    description: '',
    maxMembers: 10,
    isPrivate: false
  };
};
</script>

<template>
  <div class="study-room-container">
    <div class="room-header">
      <h2>自习室列表</h2>
      <el-button type="primary" @click="handleCreateRoom">创建自习室</el-button>
    </div>
    <div v-if="error" class="error-message">
      {{ error }}
      <el-button type="text" @click="error = null">关闭</el-button>
    </div>
    <div v-if="!showWebView" class="room-grid">
      <div v-for="room in rooms" :key="room.id" class="room-card" @click="handleRoomClick(room)">
        <div class="room-info">
          <h3>{{ room.title }}</h3>
          <p class="description">{{ room.description }}</p>
          <div class="room-status">
            <span class="status">{{ room.status }}</span>
            <span class="online-count">在线: {{ room.onlineCount }}人</span>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="webview-container">
      <ZoomView/>
    </div>

    <el-dialog
      v-model="dialogVisible"
      title="创建自习室"
      width="30%"
      :close-on-click-modal="false"
    >
      <el-form :model="newRoom" label-width="80px">
        <el-form-item label="名称">
          <el-input v-model="newRoom.title" placeholder="请输入自习室名称"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input
            v-model="newRoom.description"
            type="textarea"
            placeholder="请输入自习室描述"
          ></el-input>
        </el-form-item>
        <el-form-item label="最大人数">
          <el-input-number
            v-model="newRoom.maxMembers"
            :min="1"
            :max="50"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="私密房间">
          <el-switch v-model="newRoom.isPrivate"></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleCancel">取消</el-button>
          <el-button type="primary" @click="handleConfirm">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.study-room-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  padding: 20px;
}

.webview-container {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
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
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.room-card {
  background: white;
  border-radius: 16px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  cursor: pointer;
  transition: all 0.3s ease;
  aspect-ratio: 1.2;
  display: flex;
  flex-direction: column;
}

.room-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.12);
}

.room-info {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.room-card h3 {
  margin: 0;
  color: #333;
  font-size: 1.5rem;
  margin-bottom: 12px;
}

.description {
  color: #666;
  font-size: 1rem;
  margin: 0;
  flex-grow: 1;
}

.room-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 16px;
}

.status {
  color: #10b981;
  font-weight: 500;
}

.online-count {
  color: #6b7280;
  font-size: 0.9rem;
}
.room-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.room-header h2 {
  margin: 0;
  font-size: 1.5rem;
  color: #333;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>