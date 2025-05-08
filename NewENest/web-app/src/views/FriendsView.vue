<template>
  <div class="friends-view">
    <div class="page-header">
      <h1>好友管理</h1>
      <el-button type="primary" @click="showAddFriendDialog">
        <el-icon><Plus /></el-icon>
        添加好友
      </el-button>
    </div>

    <el-tabs v-model="activeTab" class="friends-tabs">
      <el-tab-pane label="好友列表" name="friends">
        <div class="friends-list">
          <el-skeleton :loading="loading" animated :count="3" v-if="loading">
            <template #template>
              <div class="friend-item-skeleton">
                <el-skeleton-item variant="circle" style="width: 40px; height: 40px;" />
                <div style="flex: 1; margin-left: 16px;">
                  <el-skeleton-item variant="text" style="width: 30%;" />
                  <el-skeleton-item variant="text" style="width: 50%; margin-top: 8px;" />
                </div>
              </div>
            </template>
          </el-skeleton>

          <template v-else>
            <div v-if="hasFriends" class="friend-list-content">
              <div v-for="friend in friends" :key="friend.id" class="friend-item">
                <el-avatar :src="friend.avatar" :size="40">
                  {{ friend && friend.username ? friend.username.charAt(0) : '?' }}
                </el-avatar>
                <div class="friend-info">
                  <div class="friend-name">{{ friend && friend.username ? friend.username : '未知用户' }}</div>
                  <div class="friend-meta">
                    <span class="study-time">累计学习：{{ formatTime(friend && friend.total_study_time ? friend.total_study_time : 0) }}</span>
                    <span v-if="friend && friend.study_direction">方向：{{ friend.study_direction }}</span>
                  </div>
                  <div class="friend-signature" v-if="friend && friend.signature">
                    {{ friend.signature }}
                  </div>
                </div>
                <div class="friend-actions">
                  <el-badge :value="friend.unread_messages" :hidden="friend.unread_messages === 0" type="danger">
                    <el-button size="small" type="primary" @click="openChat(friend)" plain>
                      <el-icon><ChatDotRound /></el-icon>
                      聊天
                    </el-button>
                  </el-badge>
                  <el-dropdown trigger="click" @command="handleFriendAction($event, friend)">
                    <el-button size="small" plain>
                      <el-icon><More /></el-icon>
                    </el-button>
                    <template #dropdown>
                      <el-dropdown-menu>
                        <el-dropdown-item command="study">一起学习</el-dropdown-item>
                        <el-dropdown-item command="contract">创建契约</el-dropdown-item>
                        <el-dropdown-item command="delete" divided>删除好友</el-dropdown-item>
                      </el-dropdown-menu>
                    </template>
                  </el-dropdown>
                </div>
              </div>
            </div>
            <el-empty v-else description="暂无好友" />
          </template>
        </div>
      </el-tab-pane>

      <el-tab-pane label="好友请求" name="requests" :disabled="!hasPendingRequests">
        <div class="friend-requests-list">
          <div v-for="request in friendRequests" :key="request.id" class="request-item">
            <el-avatar :src="request.sender?.avatar" :size="40">
              {{ request.sender && request.sender.username ? request.sender.username.charAt(0) : '?' }}
            </el-avatar>
            <div class="request-info">
              <div class="request-name">{{ request.sender && request.sender.username ? request.sender.username : '未知用户' }}</div>
              <div class="request-time">{{ formatDate(request.created_at) }}</div>
            </div>
            <div class="request-actions">
              <el-button size="small" type="success" @click="acceptRequest(request.id)" plain>接受</el-button>
              <el-button size="small" type="danger" @click="rejectRequest(request.id)" plain>拒绝</el-button>
            </div>
          </div>
          <el-empty v-if="friendRequests.length === 0" description="暂无好友请求" />
        </div>
      </el-tab-pane>
    </el-tabs>

    <!-- 添加好友对话框 -->
    <el-dialog v-model="addFriendDialogVisible" title="添加好友" width="500px">
      <el-form>
        <el-form-item>
          <el-input
            v-model="searchKeyword"
            placeholder="搜索用户名或邮箱"
            clearable
            @input="handleSearch"
            :suffix-icon="Search"
          >
          </el-input>
        </el-form-item>
      </el-form>

      <div class="search-results" v-loading="searchLoading">
        <div v-if="searchResults.length > 0">
          <div v-for="user in searchResults" :key="user.id" class="search-result-item">
            <el-avatar :src="user.avatar" :size="40">
              {{ user && user.username ? user.username.charAt(0) : '?' }}
            </el-avatar>
            <div class="user-info">
              <div class="user-name">{{ user && user.username ? user.username : '未知用户' }}</div>
              <div class="user-email">{{ user && user.email ? user.email : '未知邮箱' }}</div>
            </div>
            <div class="user-actions">
              <el-button
                size="small"
                type="primary"
                :disabled="user.is_friend"
                @click="sendRequest(user.id)"
                plain
              >
                {{ user.is_friend ? '已是好友' : '添加' }}
              </el-button>
            </div>
          </div>
        </div>
        <el-empty v-else-if="!searchLoading && searchKeyword" description="未找到相关用户" />
        <div v-else-if="!searchLoading && !searchKeyword" class="search-tip">
          请输入用户名或邮箱进行搜索
        </div>
      </div>
    </el-dialog>

    <!-- 聊天对话框 -->
    <el-dialog v-model="chatDialogVisible" :title="`与 ${selectedFriend?.username || ''} 的聊天`" width="600px">
      <div class="chat-container">
        <div class="chat-messages" ref="chatMessagesRef">
          <div v-for="(msg, index) in chatMessages" :key="index" 
               class="message-item" 
               :class="{ 'message-self': msg.sender_id === currentUserId }">
            <div class="message-content">{{ msg.content }}</div>
            <div class="message-time">{{ formatDate(msg.created_at) }}</div>
          </div>
          <div v-if="chatMessages.length === 0" class="empty-chat">
            <el-empty description="暂无消息记录" />
          </div>
        </div>
        <div class="chat-input">
          <el-input
            v-model="messageContent"
            type="textarea"
            :rows="2"
            placeholder="输入消息..."
            @keydown.enter.exact.prevent="sendMessage"
          ></el-input>
          <el-button type="primary" @click="sendMessage" :disabled="!messageContent.trim()">
            发送
          </el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { useFriendStore } from '../store/friendStore'
import { useUserStore } from '../store/userStore'
import { Plus, Search, ChatDotRound, More } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import friendApi from '../api/friendApi'
import { FriendInfo } from '../api/friendApi'

const userStore = useUserStore()
const friendStore = useFriendStore()

// 状态
const activeTab = ref('friends')
const addFriendDialogVisible = ref(false)
const chatDialogVisible = ref(false)
const searchKeyword = ref('')
const messageContent = ref('')
const selectedFriend = ref<FriendInfo | null>(null)
const chatMessages = ref<any[]>([])
const chatMessagesRef = ref<HTMLElement | null>(null)
const currentUserId = computed(() => userStore.userInfo?.id || 0)

// 计算属性
const loading = computed(() => friendStore.loading)
const friends = computed(() => friendStore.friends)
const friendRequests = computed(() => friendStore.friendRequests)
const hasFriends = computed(() => friendStore.hasFriends)
const hasPendingRequests = computed(() => friendStore.hasPendingRequests)
const searchResults = computed(() => friendStore.searchResults)
const searchLoading = computed(() => friendStore.searchLoading)

// 生命周期钩子
onMounted(async () => {
  // 检查用户是否已登录
  if (!userStore.isLoggedIn) {
    ElMessage.error('请先登录')
    // 跳转到登录页面
    window.location.href = '/login'
    return
  }
  
  // 检查localStorage中是否有token
  const token = localStorage.getItem('token')
  if (!token) {
    ElMessage.error('登录已过期，请重新登录')
    // 跳转到登录页面
    window.location.href = '/login'
    return
  }
  
  await friendStore.initialize()
  
  // 如果有待处理的好友请求，自动切换到请求标签页
  if (hasPendingRequests.value) {
    activeTab.value = 'requests'
  }
})

// 添加好友相关
const showAddFriendDialog = () => {
  addFriendDialogVisible.value = true
  searchKeyword.value = ''
  friendStore.searchResults = []
}

const handleSearch = () => {
  if (searchKeyword.value.trim()) {
    friendStore.searchUsers(searchKeyword.value)
  }
}

const sendRequest = async (userId: number) => {
  const result = await friendStore.sendFriendRequest(userId)
  if (result.success) {
    ElMessage.success('好友请求已发送')
    // 更新搜索结果中的状态
    const userIndex = searchResults.value.findIndex(user => user.id === userId)
    if (userIndex !== -1) {
      searchResults.value[userIndex].is_friend = true
    }
  } else {
    ElMessage.error(result.message)
  }
}

// 好友请求处理
const acceptRequest = async (requestId: number) => {
  const result = await friendStore.acceptFriendRequest(requestId)
  if (result.success) {
    ElMessage.success('已接受好友请求')
  } else {
    ElMessage.error(result.message)
  }
}

const rejectRequest = async (requestId: number) => {
  const result = await friendStore.rejectFriendRequest(requestId)
  if (result.success) {
    ElMessage.success('已拒绝好友请求')
  } else {
    ElMessage.error(result.message)
  }
}

// 好友操作处理
const handleFriendAction = (action: string, friend: FriendInfo) => {
  switch (action) {
    case 'study':
      ElMessage.info('邀请学习功能即将上线')
      break
    case 'contract':
      ElMessage.info('学习契约功能即将上线')
      break
    case 'delete':
      confirmDeleteFriend(friend)
      break
  }
}

const confirmDeleteFriend = (friend: FriendInfo) => {
  ElMessageBox.confirm(
    `确定要删除好友 ${friend.username} 吗？`,
    '删除好友',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    const result = await friendStore.deleteFriend(friend.friendship_id)
    if (result.success) {
      ElMessage.success('好友已删除')
    } else {
      ElMessage.error(result.message)
    }
  }).catch(() => {})
}

// 聊天相关
const openChat = async (friend: FriendInfo) => {
  selectedFriend.value = friend
  chatDialogVisible.value = true
  messageContent.value = ''
  chatMessages.value = []
  
  try {
    console.log('正在获取聊天记录，好友ID:', friend.id)
    const response = await friendApi.getChatHistory(friend.id)
    console.log('聊天记录API响应:', response)
    
    // 检查响应格式
    if (response && response.data) {
      chatMessages.value = response.data
      console.log('设置聊天消息为:', chatMessages.value)
    } else if (response && Array.isArray(response)) {
      // 如果响应本身就是数组
      chatMessages.value = response
      console.log('响应是数组，直接设置:', chatMessages.value)
    } else {
      console.error('无法识别的聊天记录响应格式:', response)
      ElMessage.warning('聊天记录格式不正确')
      chatMessages.value = []
    }
    
    // 滚动到最新消息
    await nextTick()
    scrollToBottom()
  } catch (error) {
    console.error('获取聊天记录失败', error)
    ElMessage.error('获取聊天记录失败')
  }
}

const sendMessage = async () => {
  if (!messageContent.value.trim() || !selectedFriend.value) return
  
  try {
    console.log('发送消息给好友ID:', selectedFriend.value.id, '内容:', messageContent.value)
    const response = await friendApi.sendMessage(selectedFriend.value.id, messageContent.value)
    console.log('发送消息API响应:', response)
    
    // 检查响应格式
    if (response && response.data) {
      chatMessages.value.push(response.data)
    } else if (response) {
      // 手动构造消息对象
      const newMessage = {
        id: Date.now(), // 临时ID
        sender_id: currentUserId.value,
        receiver_id: selectedFriend.value.id,
        content: messageContent.value,
        message_type: 'text',
        is_read: false,
        created_at: new Date().toISOString()
      }
      chatMessages.value.push(newMessage)
    }
    
    messageContent.value = ''
    
    // 滚动到最新消息
    await nextTick()
    scrollToBottom()
  } catch (error) {
    console.error('发送消息失败', error)
    ElMessage.error('发送消息失败')
  }
}

const scrollToBottom = () => {
  if (chatMessagesRef.value) {
    chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight
  }
}

// 工具函数
const formatTime = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  
  if (hours > 0) {
    return `${hours}小时${minutes}分钟`
  }
  return `${minutes}分钟`
}

const formatDate = (dateStr: string): string => {
  if (!dateStr || dateStr === "0001-01-01T08:05:43+08:05") {
    return '刚刚';
  }
  
  let date;
  try {
    date = new Date(dateStr);
    // 检查日期是否有效
    if (isNaN(date.getTime())) {
      return '时间未知';
    }
  } catch {
    return '时间未知';
  }
  
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date);
}
</script>

<style scoped>
.friends-view {
  margin-bottom: 40px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.friends-tabs {
  margin-top: 20px;
}

.friend-item, .request-item, .friend-item-skeleton {
  display: flex;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid var(--border-color-light);
}

.friend-item:last-child, .request-item:last-child {
  border-bottom: none;
}

.friend-info, .request-info {
  flex: 1;
  margin-left: 16px;
}

.friend-name, .request-name {
  font-weight: 500;
  font-size: 16px;
  margin-bottom: 4px;
}

.friend-meta {
  font-size: 13px;
  color: var(--text-color-secondary);
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 4px;
}

.friend-signature {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin-top: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 400px;
}

.friend-actions, .request-actions {
  display: flex;
  gap: 8px;
}

.request-time {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.search-results {
  max-height: 300px;
  overflow-y: auto;
  margin-top: 16px;
}

.search-result-item {
  display: flex;
  align-items: center;
  padding: 12px;
  border-bottom: 1px solid var(--border-color-light);
}

.search-result-item:last-child {
  border-bottom: none;
}

.user-info {
  flex: 1;
  margin-left: 12px;
}

.user-name {
  font-weight: 500;
}

.user-email {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.search-tip {
  padding: 20px;
  text-align: center;
  color: var(--text-color-secondary);
}

.chat-container {
  display: flex;
  flex-direction: column;
  height: 400px;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  background-color: var(--background-color-light);
  border-radius: 4px;
  margin-bottom: 16px;
}

.message-item {
  margin-bottom: 12px;
  max-width: 80%;
}

.message-self {
  margin-left: auto;
  text-align: right;
}

.message-content {
  display: inline-block;
  padding: 8px 12px;
  border-radius: 4px;
  background-color: #e8f4ff;
  word-break: break-word;
}

.message-self .message-content {
  background-color: #95ec69;
}

.message-time {
  font-size: 12px;
  color: var(--text-color-secondary);
  margin-top: 4px;
}

.chat-input {
  display: flex;
  gap: 8px;
}

.chat-input .el-button {
  align-self: flex-start;
}

.empty-chat {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style> 