<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import ZoomView from '../components/ZoomView.vue';
import { wsClient } from '../utils/websocket';
import { MessageType, StudyRoomOperation, type StudyRoomMessage } from '../proto/message';

const router = useRouter();
const showWebView = ref(false);
const dialogVisible = ref(false);
const newRoom = ref({
  title: '',
  description: '',
  maxMembers: 10,
  isPrivate: false,
  duration: '2h'
});

interface Room {
  id: number;
  name: string;
  description: string;
  status: string;
  currentMembers: number;
  maxMembers: number;
}

const rooms = ref<Room[]>([]);
const userData = ref<any>(null);
const error = ref<string | null>(null);
const loading = ref(false);

onMounted(() => {
  const storedUserData = localStorage.getItem('user_data');
  if (!storedUserData) {
    router.push('/login');
    return;
  }
  userData.value = JSON.parse(storedUserData);
  initWebSocket();
  fetchRooms();
});

const initWebSocket = () => {
  wsClient.connect('ws://localhost:8080/ws');
  wsClient.registerHandler(MessageType.STUDY_ROOM, handleStudyRoomResponse);
  wsClient.registerHandler(MessageType.ERROR, handleErrorResponse);
};

const handleStudyRoomResponse = (payload: Uint8Array) => {
  const response = JSON.parse(new TextDecoder().decode(payload));
  if (Array.isArray(response)) {
    rooms.value = response;
  } else if (response.roomId) {
    router.push(`/main/study-room/${response.roomId}`);
  }
  loading.value = false;
};

const handleErrorResponse = (payload: Uint8Array) => {
  const error = JSON.parse(new TextDecoder().decode(payload));
  ElMessage.error(error.message || '操作失败');
  loading.value = false;
};

const fetchRooms = () => {
  loading.value = true;
  const message: StudyRoomMessage = {
    operation: StudyRoomOperation.GET_DETAIL
  };
  const payload = new TextEncoder().encode(JSON.stringify(message));
  wsClient.sendMessage(MessageType.STUDY_ROOM, payload);
};

const handleRoomClick = (room: any) => {
  const message: StudyRoomMessage = {
    operation: StudyRoomOperation.JOIN,
    roomId: room.id
  };
  const payload = new TextEncoder().encode(JSON.stringify(message));
  wsClient.sendMessage(MessageType.STUDY_ROOM, payload);
  showWebView.value = true;
};

const handleCreateRoom = () => {
  dialogVisible.value = true;
};

const handleCancel = () => {
  dialogVisible.value = false;
  newRoom.value = {
    title: '',
    description: '',
    maxMembers: 10,
    isPrivate: false,
    duration: '2h'
  };
};

const handleConfirm = () => {
  loading.value = true;
  const message: StudyRoomMessage = {
    operation: StudyRoomOperation.CREATE,
    name: newRoom.value.title,
    maxMembers: newRoom.value.maxMembers,
    isPrivate: newRoom.value.isPrivate,
    duration: newRoom.value.duration
  };
  const payload = new TextEncoder().encode(JSON.stringify(message));
  wsClient.sendMessage(MessageType.STUDY_ROOM, payload);
  dialogVisible.value = false;
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

    <div v-if="!showWebView" class="room-grid" v-loading="loading">
      <div v-for="room in rooms" :key="room.id" class="room-card" @click="handleRoomClick(room)">
        <div class="room-info">
          <h3>{{ room.name }}</h3>
          <p class="description">{{ room.description }}</p>
          <div class="room-status">
            <span class="status">{{ room.status }}</span>
            <span class="online-count">在线: {{ room.currentMembers }}/{{ room.maxMembers }}人</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="webview-container">
      <ZoomView />
    </div>

    <el-dialog
      v-model="dialogVisible"
      title="创建自习室"
      width="500px"
    >
      <el-form :model="newRoom" label-width="100px">
        <el-form-item label="房间名称">
          <el-input v-model="newRoom.title" placeholder="请输入房间名称" />
        </el-form-item>
        <el-form-item label="最大人数">
          <el-input-number v-model="newRoom.maxMembers" :min="1" :max="20" />
        </el-form-item>
        <el-form-item label="持续时间">
          <el-select v-model="newRoom.duration">
            <el-option label="1小时" value="1h" />
            <el-option label="2小时" value="2h" />
            <el-option label="4小时" value="4h" />
          </el-select>
        </el-form-item>
        <el-form-item label="私密房间">
          <el-switch v-model="newRoom.isPrivate" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleCancel">取消</el-button>
          <el-button type="primary" @click="handleConfirm" :loading="loading">创建</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.study-room-container {
  padding: 20px;
  height: 100%;
  overflow-y: auto;
}

.room-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.room-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
  padding: 20px 0;
}

.room-card {
  background: white;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.room-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
}

.room-info h3 {
  margin: 0 0 10px 0;
  color: #333;
}

.description {
  color: #666;
  margin-bottom: 15px;
}

.room-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 14px;
}

.status {
  color: #409EFF;
}

.online-count {
  color: #67C23A;
}

.error-message {
  background-color: #FEF0F0;
  color: #F56C6C;
  padding: 10px;
  border-radius: 4px;
  margin-bottom: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.webview-container {
  height: calc(100vh - 120px);
}
</style>