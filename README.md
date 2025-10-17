# 论坛666 - 轻量代码讨论论坛模板

包含：
- frontend/ (Vue 3 + Vite)
- backend/ (Golang + Gin)

## 快速运行

### 后端（Go）
1. 进入 backend 目录：
   ```
   cd backend
   ```
2. 安装依赖并运行：
   ```
   go mod tidy
   go run main.go
   ```
   后端监听 http://localhost:8080

默认内存账户：
- 管理员：admin / 123456
- 普通用户：user1 / 123456

### 前端（Vue + Vite）
1. 进入 frontend：
   ```
   cd frontend
   ```
2. 安装依赖并启动：
   ```
   npm install
   npm run dev
   ```
3. 打开浏览器访问 Vite 提供的地址（通常 http://localhost:5173）。

前端通过 Authorization: Bearer <token> 来传递登录后返回的 token。
示例登录后 token 为 "<username>-token"（简化实现，生产请使用 JWT / 安全存储）。
