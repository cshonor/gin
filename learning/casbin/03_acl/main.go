// 对应视频: 3.ACL访问控制
// ACL: 直接定义 用户->资源->操作 的权限，适合简单场景
package main

import (
	"fmt"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

func main() {
	model := filepath.Join("learning", "casbin", "03_acl", "model.conf")
	e, _ := casbin.NewEnforcer(model)

	// ACL 策略: 每个用户对每个资源单独授权
	e.AddPolicy("user1", "/api/users", "GET")
	e.AddPolicy("user1", "/api/users/1", "GET")
	e.AddPolicy("admin", "/api/users", "GET")
	e.AddPolicy("admin", "/api/users", "POST")
	e.AddPolicy("admin", "/api/users", "DELETE")

	fmt.Println("user1 GET /api/users:", mustEnforce(e, "user1", "/api/users", "GET"))       // true
	fmt.Println("user1 POST /api/users:", mustEnforce(e, "user1", "/api/users", "POST"))     // false
	fmt.Println("admin DELETE /api/users:", mustEnforce(e, "admin", "/api/users", "DELETE")) // true
}

func mustEnforce(e *casbin.Enforcer, sub, obj, act string) bool {
	ok, _ := e.Enforce(sub, obj, act)
	return ok
}

