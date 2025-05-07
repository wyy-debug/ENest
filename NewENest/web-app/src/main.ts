import { createApp } from 'vue';
import { createPinia } from 'pinia';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import zhCn from 'element-plus/dist/locale/zh-cn.mjs';
import App from './App.vue';
import router from './router';

import './style.css';

const app = createApp(App);

// 状态管理
app.use(createPinia());

// 路由
app.use(router);

// UI组件库
app.use(ElementPlus, {
  locale: zhCn
});

// 挂载应用
app.mount('#app'); 