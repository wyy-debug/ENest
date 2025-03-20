<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from '../utils/axios';
import { ElMessage } from 'element-plus';

const router = useRouter();
const loading = ref(false);

const form = ref({
  title: '',
  description: '',
  maxMembers: 10,
  isPrivate: false
});

const handleCancel = () => {
  router.back();
};

const handleConfirm = async () => {
  if (!form.value.title || !form.value.description) {
    ElMessage.warning('请填写完整信息');
    return;
  }

  loading.value = true;
  try {
    await axios.post('/study-room/create', {
      name: form.value.title,
      description: form.value.description,
      max_members: form.value.maxMembers,
      is_private: form.value.isPrivate,
      duration: 24 * 60 * 60 * 1000 // 24小时
    });
    ElMessage.success('创建成功');
    router.back();
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '创建失败');
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="create-room-container">
    <div class="header">
      <el-button @click="handleCancel">
        <el-icon><i-ep-arrow-left /></el-icon>
        返回
      </el-button>
      <h2>创建自习室</h2>
    </div>

    <div class="form-container">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称" required>
          <el-input v-model="form.title" placeholder="请输入自习室名称"></el-input>
        </el-form-item>
        <el-form-item label="描述" required>
          <el-input
            v-model="form.description"
            type="textarea"
            placeholder="请输入自习室描述"
          ></el-input>
        </el-form-item>
        <el-form-item label="最大人数">
          <el-input-number
            v-model="form.maxMembers"
            :min="1"
            :max="50"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="私密房间">
          <el-switch v-model="form.isPrivate"></el-switch>
        </el-form-item>

        <el-form-item>
          <div class="button-group">
            <el-button @click="handleCancel">取消</el-button>
            <el-button type="primary" @click="handleConfirm" :loading="loading">创建</el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped>
.create-room-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.header {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
}

.header h2 {
  margin: 0;
  margin-left: 20px;
}

.form-container {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.button-group {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>