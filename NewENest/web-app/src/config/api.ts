import axios from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  }
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 从本地存储获取token
    const token = localStorage.getItem('token');
    
    // 详细调试信息
    console.log('===== API请求信息 =====');
    console.log('请求URL:', config.url);
    console.log('请求方法:', config.method);
    console.log('本地存储中的token:', token);
    
    // 如果存在token，添加到请求头
    if (token) {
      // 确保token格式正确，添加Bearer前缀
      const authToken = token.startsWith('Bearer ') ? token : `Bearer ${token}`;
      config.headers.Authorization = authToken;
      
      // 输出调试信息
      console.log('添加授权头:', authToken);
      console.log('token长度:', token.length);
      console.log('token前10个字符:', token.substring(0, 10));
      console.log('token是否包含Bearer前缀:', token.startsWith('Bearer '));
    } else {
      console.warn('警告：未找到token，请求未添加授权头');
    }
    
    console.log('完整请求头:', config.headers);
    console.log('======================');
    
    return config;
  },
  (error) => {
    console.error('请求拦截器错误:', error);
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    // 调试响应信息
    console.log('===== API响应信息 =====');
    console.log('响应状态:', response.status);
    console.log('响应URL:', response.config.url);
    console.log('响应数据:', response.data);
    console.log('======================');
    
    // 直接返回响应数据
    return response.data;
  },
  (error) => {
    console.error('API请求错误:', error);
    console.log('错误详情:', {
      status: error.response?.status,
      url: error.config?.url,
      method: error.config?.method,
      data: error.response?.data
    });
    
    // 处理401未授权错误
    if (error.response && error.response.status === 401) {
      console.warn('收到401未授权响应，准备重定向到登录页');
      
      // 清除token
      localStorage.removeItem('token');
      
      // 如果不在登录页，重定向到登录页
      if (window.location.pathname !== '/login') {
        console.log('重定向到登录页...');
        window.location.href = '/login';
      }
    }
    
    return Promise.reject(error);
  }
);

export default api; 