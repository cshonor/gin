# 对应视频: 11.接口测试工具杂谈

## 常用 API 测试工具

### 1. curl（命令行）
```bash
# GET 请求
curl http://localhost:8080/api/users

# POST JSON
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","age":25}'

# 带查询参数
curl "http://localhost:8080/search?keyword=gin&page=1"

# 带 Header
curl -H "Authorization: Bearer token123" http://localhost:8080/api/protected
```

### 2. Postman
- 图形界面，易于使用
- 支持环境变量、集合
- 可导出分享

### 3. Apifox / Apizza
- 国产工具，中文友好
- 支持 API 文档、Mock

### 4. VS Code 扩展
- REST Client
- Thunder Client

### 5. Go 代码测试
```go
// 使用 httptest
req := httptest.NewRequest("GET", "/user/1", nil)
w := httptest.NewRecorder()
router.ServeHTTP(w, req)
```

