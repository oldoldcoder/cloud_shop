# Cloud Shop Frontend

这是一个基于 Vue 3 + Element Plus 的前端项目模板。

## 技术栈

- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 下一代前端构建工具
- **Vue Router 4** - Vue.js 官方路由管理器
- **Element Plus** - 基于 Vue 3 的组件库
- **Axios** - HTTP 客户端

## 项目结构

```
front-code/
├── src/
│   ├── views/          # 页面组件
│   ├── router/         # 路由配置
│   ├── components/     # 公共组件
│   ├── assets/         # 静态资源
│   ├── utils/          # 工具函数
│   ├── api/            # API 接口
│   ├── App.vue         # 根组件
│   ├── main.js         # 入口文件
│   └── style.css       # 全局样式
├── public/             # 公共资源
├── index.html          # HTML 模板
├── vite.config.js      # Vite 配置
├── package.json        # 项目依赖
└── README.md           # 项目说明
```

## 快速开始

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

### 构建生产版本

```bash
npm run build
```

### 预览生产版本

```bash
npm run preview
```

## 开发说明

- 项目使用 Vite 作为构建工具，支持热重载
- 集成了 Element Plus UI 组件库
- 配置了代理，API 请求会转发到后端服务 (localhost:8081)
- 使用 Vue 3 Composition API 语法
- 支持 TypeScript（可选）

## 自定义配置

可以在 `vite.config.js` 中修改配置，比如：
- 修改开发服务器端口
- 配置 API 代理
- 添加构建优化选项

## 部署

构建完成后，将 `dist` 目录部署到 Web 服务器即可。
