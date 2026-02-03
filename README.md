# Gin Web 框架项目

这是一个使用 Gin 框架创建的 Go Web 应用项目，包含完整的项目结构和示例代码。

## 功能特性

- ✅ 使用 Gin 框架构建 RESTful API
- ✅ **GORM ORM 集成**（支持 SQLite、MySQL、PostgreSQL）
- ✅ 模块化的项目结构
- ✅ 中间件支持（日志、CORS、认证）
- ✅ 路由组管理
- ✅ 数据模型和处理器
- ✅ 配置管理
- ✅ JSON 响应支持
- ✅ 请求验证
- ✅ 数据库自动迁移
- ✅ 软删除支持
- ✅ 分页查询

## 快速开始

### 安装依赖

```bash
go mod download
```

### 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

### 测试接口

#### 基础接口
- `GET http://localhost:8080/` - 返回欢迎消息
- `GET http://localhost:8080/ping` - 返回 pong 消息

#### 用户管理接口（CRUD）
- `GET http://localhost:8080/api/v1/users` - 获取所有用户（支持分页：?page=1&page_size=10）
- `GET http://localhost:8080/api/v1/users/:id` - 获取指定用户
- `POST http://localhost:8080/api/v1/users` - 创建新用户
  ```json
  {
    "username": "testuser",
    "email": "test@example.com",
    "age": 25
  }
  ```
- `PUT http://localhost:8080/api/v1/users/:id` - 更新用户信息
- `DELETE http://localhost:8080/api/v1/users/:id` - 删除用户

#### 受保护接口（需要认证）
- `GET http://localhost:8080/api/v1/protected/info` - 需要 Bearer Token
  ```
  Header: Authorization: Bearer your-token-here
  ```

## 项目结构

```
.
├── config/              # 配置管理
│   └── config.go       # 配置加载
├── database/           # 数据库连接
│   └── database.go     # GORM 数据库初始化和迁移
├── handlers/           # 请求处理器
│   └── user_handler.go # 用户相关处理器
├── middleware/         # 中间件
│   ├── auth.go        # 认证中间件
│   ├── cors.go        # 跨域中间件
│   └── logger.go      # 日志中间件
├── models/             # 数据模型
│   └── user.go        # 用户模型（GORM）
├── routes/             # 路由配置
│   └── routes.go      # 路由设置
├── go.mod             # Go 模块依赖文件
├── main.go            # 主程序入口
├── .gitignore         # Git 忽略文件
└── README.md          # 项目说明文档
```

## 项目说明

### 中间件

- **Logger**: 自定义请求日志格式
- **CORS**: 处理跨域请求
- **Auth**: Bearer Token 认证示例

### 数据模型

项目包含用户模型示例，展示了：
- 结构体定义
- 请求验证标签（binding）
- 请求/响应模型分离

### 路由组织

- 使用路由组 (`/api/v1`) 进行 API 版本管理
- 公开路由和受保护路由分离
- 清晰的 RESTful API 设计

## 数据库配置

项目支持多种数据库，通过 `DATABASE_URL` 环境变量配置：

### SQLite（默认，开发环境）
```bash
# 不设置 DATABASE_URL 或设置为 SQLite 文件路径
DATABASE_URL=gin.db go run main.go
```

### MySQL
```bash
DATABASE_URL="user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" go run main.go
```

### PostgreSQL
```bash
DATABASE_URL="host=localhost user=postgres password=postgres dbname=gin port=5432 sslmode=disable" go run main.go
```

或者使用 PostgreSQL URL 格式：
```bash
DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable" go run main.go
```

## 环境变量

可以通过环境变量配置应用：

- `PORT`: 服务器端口（默认: 8080）
- `DATABASE_URL`: 数据库连接字符串（不设置则默认使用 SQLite: gin.db）
- `JWT_SECRET`: JWT 密钥
- `ENVIRONMENT`: 运行环境（development/production）

示例：
```bash
PORT=3000 DATABASE_URL=gin.db go run main.go
```

## GORM 特性

- ✅ 自动数据库迁移
- ✅ 软删除支持
- ✅ 模型关联（可扩展）
- ✅ 事务支持
- ✅ 查询构建器
- ✅ 预加载（Eager Loading）

## 下一步

你可以：
- ✅ ~~连接真实的数据库~~（已完成，支持 SQLite/MySQL/PostgreSQL）
- 实现 JWT 认证和用户登录
- 添加单元测试
- 集成 Swagger 文档
- 添加 Docker 支持
- 实现更复杂的业务逻辑
- 添加数据库关联（一对多、多对多等）
- 实现数据验证和业务规则
- 添加缓存层（Redis）

