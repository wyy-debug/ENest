import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { userApi, type UserProfileDTO } from '@/api/user';

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref<string>('');
  const user = ref<UserProfileDTO | null>(null);
  
  // 计算属性
  const isLoggedIn = computed(() => !!token.value);
  const username = computed(() => user.value?.username || '');
  
  // 从本地存储恢复状态
  const initFromStorage = () => {
    const storedToken = localStorage.getItem('token');
    const storedUser = localStorage.getItem('user');
    
    if (storedToken) {
      token.value = storedToken;
    }
    
    if (storedUser) {
      try {
        user.value = JSON.parse(storedUser);
      } catch (e) {
        console.error('Failed to parse stored user data');
      }
    }
  };
  
  // 初始化状态
  initFromStorage();
  
  // 操作方法
  const login = async (username: string, password: string) => {
    const result = await userApi.login({ username, password });
    token.value = result.token;
    user.value = result.user;
    
    // 保存到本地存储
    localStorage.setItem('token', result.token);
    localStorage.setItem('user', JSON.stringify(result.user));
  };
  
  const register = async (username: string, email: string, password: string) => {
    const newUser = await userApi.register({ username, email, password });
    user.value = newUser;
    
    // 注册后不自动登录，需要用户手动登录
  };
  
  const logout = async () => {
    try {
      // 只有当存在token时才调用登出API
      if (token.value) {
        await userApi.logout();
      }
    } catch (error) {
      console.error('Logout API error:', error);
    } finally {
      // 无论API是否成功都清除本地状态
      token.value = '';
      user.value = null;
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    }
  };
  
  const fetchUserProfile = async () => {
    try {
      const profile = await userApi.getProfile();
      user.value = profile;
      localStorage.setItem('user', JSON.stringify(profile));
    } catch (error) {
      console.error('Failed to fetch user profile:', error);
    }
  };
  
  const updateProfile = async (profileData: {
    username?: string;
    signature?: string;
    study_direction?: string;
    avatar?: string;
  }) => {
    const updatedProfile = await userApi.updateProfile(profileData);
    user.value = updatedProfile;
    localStorage.setItem('user', JSON.stringify(updatedProfile));
  };
  
  return {
    // 状态
    token,
    user,
    
    // 计算属性
    isLoggedIn,
    username,
    
    // 操作方法
    login,
    register,
    logout,
    fetchUserProfile,
    updateProfile
  };
}); 