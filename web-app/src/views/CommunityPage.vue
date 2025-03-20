<script setup lang="ts">
import { ref } from 'vue'

const posts = ref([
  {
    id: 1,
    title: '分享一下我的学习方法',
    content: '最近在使用番茄工作法学习，效果很好...',
    author: '学习达人',
    likes: 15,
    comments: 5,
    createTime: '2024-01-20'
  },
  {
    id: 2,
    title: '推荐几个不错的学习资源',
    content: '整理了一些很棒的学习网站和工具...',
    author: '资源达人',
    likes: 20,
    comments: 8,
    createTime: '2024-01-19'
  }
])
</script>

<template>
  <div class="community-container">
    <div class="community-header">
      <h2>学习社区</h2>
      <el-button type="primary" size="large">
        <el-icon><i-ep-plus /></el-icon>
        发布帖子
      </el-button>
    </div>
    
    <div class="posts-container">
      <el-card v-for="post in posts" :key="post.id" class="post-card">
        <div class="post-header">
          <h3>{{ post.title }}</h3>
          <span class="post-author">{{ post.author }}</span>
        </div>
        <p class="post-content">{{ post.content }}</p>
        <div class="post-footer">
          <div class="post-stats">
            <span class="stat-item">
              <el-icon><i-ep-thumb /></el-icon>
              {{ post.likes }}
            </span>
            <span class="stat-item">
              <el-icon><i-ep-chat-dot-round /></el-icon>
              {{ post.comments }}
            </span>
          </div>
          <span class="post-time">{{ post.createTime }}</span>
        </div>
      </el-card>
    </div>
  </div>
</template>

<style scoped>
.community-container {
  padding: 20px;
  height: 100%;
  overflow-y: auto;
}

.community-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.community-header h2 {
  margin: 0;
  color: #333;
}

.posts-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.post-card {
  background: white;
  border-radius: 8px;
  transition: all 0.3s;
}

.post-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.post-header h3 {
  margin: 0;
  font-size: 18px;
  color: #333;
}

.post-author {
  color: #666;
  font-size: 14px;
}

.post-content {
  color: #666;
  margin: 12px 0;
  line-height: 1.6;
}

.post-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid #eee;
}

.post-stats {
  display: flex;
  gap: 16px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #666;
}

.post-time {
  color: #999;
  font-size: 14px;
}
</style>