# NewENest

NewENest 是一个在线自习室和学习社区平台，旨在为用户提供专注学习环境和社交互动功能。本项目是对原 ENest 项目的全面重构，采用现代化的技术栈和架构设计。

## 项目结构

项目采用前后端分离架构：

- 后端：基于 Go + Fiber 开发的 RESTful API 服务
- 前端：基于 Vue 3 + TypeScript 开发的单页应用

### 后端目录结构

```
go-server/
├── config/             # 配置管理
├── database/           # 数据库连接和迁移
│   └── migrations/     # 数据库迁移文件
├── handlers/           # HTTP 控制器
├── middleware/         # HTTP 中间件
├── models/             # 数据模型
├── repositories/       # 数据访问层
├── services/           # 业务逻辑层
├── utils/              # 工具函数
├── main.go             # 应用入口
└── go.mod              # Go 模块定义
```

### 前端目录结构

```
web-app/
├── public/             # 静态资源
├── src/
│   ├── api/            # API 请求封装
│   ├── assets/         # 静态资源
│   ├── components/     # 通用组件
│   ├── composables/    # 可复用逻辑
│   ├── router/         # 路由配置
│   ├── stores/         # 状态管理
│   ├── views/          # 页面组件
│   ├── App.vue         # 根组件
│   └── main.ts         # 入口文件
├── package.json        # 依赖管理
└── vite.config.ts      # 构建配置
```

## 核心功能

- 用户管理：注册、登录、个人信息管理
- 自习室：创建、加入、管理自习室
- 学习记录：记录和统计学习时间数据
- 社交功能：好友管理、学习契约、私聊
- 成就系统：学习成就和奖励机制

## 技术栈

### 后端

- **Go**：核心编程语言
- **Fiber**：高性能 Web 框架
- **PostgreSQL**：关系型数据库
- **Redis**：缓存和会话管理
- **sqlx**：数据库操作库
- **JWT**：身份验证

### 前端

- **Vue 3**：渐进式 JavaScript 框架
- **TypeScript**：类型安全的 JavaScript 超集
- **Pinia**：状态管理
- **Vue Router**：路由管理
- **Element Plus**：UI 组件库
- **Axios**：HTTP 客户端

## 开发环境设置

### 后端

1. 安装 Go 1.20 或更高版本
2. 安装 PostgreSQL 14 或更高版本
3. 安装 Redis 6 或更高版本
4. 克隆本仓库
5. 配置环境变量或创建 `.env` 文件
6. 运行以下命令：

```bash
cd go-server
go mod download
go run main.go
```

### 前端

1. 安装 Node.js 18 或更高版本
2. 克隆本仓库
3. 运行以下命令：

```bash
cd web-app
npm install
npm run dev
```

## API 文档

API 文档基于 OpenAPI 规范，详见 [API.md](API.md)。

## 数据库设计

数据库设计详见 [模型设计文档.md](模型设计文档.md)。

## 贡献指南

欢迎贡献代码，请遵循以下步骤：

1. Fork 本仓库
2. 创建功能分支：`git checkout -b feature/your-feature-name`
3. 提交更改：`git commit -m 'Add some feature'`
4. 推送到分支：`git push origin feature/your-feature-name`
5. 提交 Pull Request

## 许可证

本项目采用 MIT 许可证。 