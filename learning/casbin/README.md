# Casbin 权限控制学习

对应视频教程 1-5 集。

| 目录 | 视频 | 内容 |
|------|------|------|
| 01_intro | 1.casbin基本介绍 | 基本用法、Enforce、AddPolicy |
| 02_config | 2.casbin配置文件 | model.conf、policy.csv |
| 03_acl | 3.ACL访问控制 | 主体对资源直接授权 |
| 04_rbac | 4.RBAC访问控制 | 用户->角色->权限 |
| 05_gorm_adapter | 5.gorm接入casbin | 策略持久化到数据库 |
| 06_gin_integration | Casbin+Gin 配合 | 中间件校验权限，Header X-User 传用户 |
| 07_abac | ABAC 访问控制 | 基于属性的权限，sub/obj 传结构体，matcher 用 r.sub.xxx |

## 运行

```bash
go run ./learning/casbin/01_intro
go run ./learning/casbin/02_config
go run ./learning/casbin/03_acl
go run ./learning/casbin/04_rbac
go run ./learning/casbin/05_gorm_adapter
go run ./learning/casbin/06_gin_integration
go run ./learning/casbin/07_abac
```

### 06 测试
```bash
# 无权限
curl http://localhost:8080/api/users

# alice(admin) 有权限
curl -H "X-User: alice" http://localhost:8080/api/users

# bob(user) 只能 GET
curl -H "X-User: bob" http://localhost:8080/api/users
curl -H "X-User: bob" -X POST http://localhost:8080/api/users  # 403
```

## 依赖

需添加：
- github.com/casbin/casbin/v2
- github.com/casbin/gorm-adapter/v3

运行 `go mod tidy` 自动下载。

