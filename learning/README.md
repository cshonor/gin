# Gin 框架学习笔记

配合视频教程使用的代码示例，每个目录对应一个视频章节。

## 目录结构

| 目录 | 对应视频 | 内容 |
|------|----------|------|
| [02_go_builtin_http](./02_go_builtin_http/) | 2.go内置http库 | Go 标准库 net/http 基础 |
| [03_native_http_painpoints](./03_native_http_painpoints/) | 3.原生http库的一些痛点 | 为什么需要 Web 框架 |
| [04_init_gin](./04_init_gin/) | 4.初始gin框架 | Gin 初始化与路由 |
| [05_json_response](./05_json_response/) | 5.gin响应-json响应封装 | JSON 响应与统一封装 |
| [06_json_encapsulation2](./06_json_encapsulation2/) | 6.gin响应json-响应封装2 | 完善响应封装、分页 |
| [07_response_html](./07_response_html/) | 7.响应html | HTML 模板渲染 |
| [08_response_file](./08_response_file/) | 8.响应文件 | 静态文件与文件下载 |
| [09_static_files](./09_static_files/) | 9.静态文件 | 静态目录与文件服务 |
| [10_query_and_path_params](./10_query_and_path_params/) | 10.查询参数和动态参数 | Query、Param、QueryArray |
| [11_api_testing](./11_api_testing/) | 11.接口测试工具杂谈 | curl、Postman 等测试工具 |
| [12_form_and_upload](./12_form_and_upload/) | 12.表单参数与文件上传 | 表单解析、单/多文件上传 |
| [13_raw_content](./13_raw_content/) | 13.原始内容 | 获取原始 Body、Request |
| [14_bind_query_path](./14_bind_query_path/) | 14.bind参数绑定 | ShouldBindQuery、ShouldBindUri |
| [15_bind_json_header](./15_bind_json_header/) | 15.binding参数 | ShouldBindJSON、Header 绑定 |
| [16_binding_rules](./16_binding_rules/) | 16.binding内置规则 | required、min、max、email、oneof 等 |
| [17_binding_zh](./17_binding_zh/) | 17.binding错误信息显示中文 | 验证错误中文翻译 |
| [18_binding_error_detail](./18_binding_error_detail/) | 18.binding显示错误字段和错误信息 | 结构化验证错误响应 |
| [19_custom_validator](./19_custom_validator/) | 19.自定义验证器 | RegisterValidation 自定义规则 |
| [20_routing](./20_routing/) | 20.路由 | 路由组、NoRoute、重定向 |
| [21_local_middleware](./21_local_middleware/) | 21.局部中间件 | 路由组/单路由中间件 |
| [22_global_middleware](./22_global_middleware/) | 22.全局中间件 | Use、gin.New、Recovery |

## 运行方式

在项目根目录下运行对应示例：

```bash
# 02 - Go 内置 HTTP
go run ./learning/02_go_builtin_http

# 04 - 初始化 Gin
go run ./learning/04_init_gin

# 05 - JSON 响应
go run ./learning/05_json_response

# 07 - HTML 响应（需在 learning/07_response_html 目录下）
go run ./learning/07_response_html

# 08 - 文件响应
go run ./learning/08_response_file

# 09 - 静态文件
go run ./learning/09_static_files

# 10 - 查询参数和动态参数
go run ./learning/10_query_and_path_params

# 12 - 表单与文件上传
go run ./learning/12_form_and_upload

# 13 - 原始内容
go run ./learning/13_raw_content

# 14 - Bind 查询/路径参数
go run ./learning/14_bind_query_path

# 15 - Bind JSON/Header
go run ./learning/15_bind_json_header

# 16 - Binding 内置规则
go run ./learning/16_binding_rules

# 17 - Binding 错误中文
go run ./learning/17_binding_zh

# 18 - Binding 错误详情
go run ./learning/18_binding_error_detail

# 19 - 自定义验证器
go run ./learning/19_custom_validator

# 20 - 路由
go run ./learning/20_routing

# 21 - 局部中间件
go run ./learning/21_local_middleware

# 22 - 全局中间件
go run ./learning/22_global_middleware
```

注意：不同示例使用相同端口 8080，同时只能运行一个。运行前请先停止其他示例。

## 学习建议

1. 按视频顺序学习，从 02 开始
2. 先看视频理解概念，再看代码加深印象
3. 可以修改代码做实验，观察效果变化
4. 07、08 需要对应的模板和静态文件目录

## 测试示例

### 02 - 内置 HTTP
```bash
curl http://localhost:8080/
curl "http://localhost:8080/hello?name=Gin"
```

### 04 - Gin 初始化
```bash
curl http://localhost:8080/user/张三
curl "http://localhost:8080/search?keyword=gin&page=2"
```

### 05/06 - JSON 响应
```bash
curl http://localhost:8080/user
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"测试","age":20}'
```

### 07 - HTML 响应
浏览器访问: http://localhost:8080/

### 08 - 文件响应
```bash
curl -O http://localhost:8080/download
```

### 10 - 查询/路径参数
```bash
curl http://localhost:8080/user/123
curl "http://localhost:8080/search?keyword=gin&page=1&page_size=20"
curl "http://localhost:8080/tags?tag=a&tag=b"
```

### 12 - 表单与文件上传
```bash
curl -X POST -d "username=admin&password=123" http://localhost:8080/login
curl -X POST -F "file=@test.txt" http://localhost:8080/upload
```

### 13 - 原始内容
```bash
curl -X POST -d '{"key":"value"}' -H "Content-Type: application/json" http://localhost:8080/raw
```

### 14 - Bind 参数
```bash
curl "http://localhost:8080/search?keyword=gin&page=1"
curl http://localhost:8080/user/123
```

### 15 - Bind JSON/Header
```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"test","email":"test@a.com","age":20}' http://localhost:8080/user
curl -H "Authorization: Bearer xxx" -H "X-Token: 123" http://localhost:8080/auth
```

### 16 - Binding 内置规则
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"ab","age":25,"email":"test@a.com","password":"123456","role":"user"}' http://localhost:8080/user
```

### 17/18 - Binding 错误
```bash
# 故意传错参数看中文错误
curl -X POST -H "Content-Type: application/json" -d '{"name":"a","email":"invalid"}' http://localhost:8080/user
```

### 19 - 自定义验证器
```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"test123","phone":"13800138000","password":"Pass1234"}' http://localhost:8080/register
```

### 21 - 局部中间件
```bash
curl http://localhost:8080/public
curl -H "X-Token: abc" http://localhost:8080/api/profile
```

### 22 - 全局中间件
```bash
curl http://localhost:8080/
curl http://localhost:8080/panic
```

