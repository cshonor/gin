// 对应视频: 4.RBAC访问控制
// RBAC: 用户->角色->权限，角色继承用 g 的层级表示
package main

import (
	"fmt"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

func main() {
	model := filepath.Join("learning", "casbin", "04_rbac", "model.conf")
	e, _ := casbin.NewEnforcer(model)

	// 角色权限: 角色->资源->操作
	e.AddPolicy("admin", "/api/users", "GET")
	e.AddPolicy("admin", "/api/users", "POST")
	e.AddPolicy("user", "/api/users", "GET")

	// 用户角色: 用户->角色
	e.AddGroupingPolicy("alice", "admin")
	e.AddGroupingPolicy("bob", "user")

	fmt.Println("alice(admin) POST /api/users:", mustEnforce(e, "alice", "/api/users", "POST")) // true
	fmt.Println("bob(user) GET /api/users:", mustEnforce(e, "bob", "/api/users", "GET"))        // true
	fmt.Println("bob(user) POST /api/users:", mustEnforce(e, "bob", "/api/users", "POST"))      // false
}

func mustEnforce(e *casbin.Enforcer, sub, obj, act string) bool {
	ok, _ := e.Enforce(sub, obj, act)
	return ok
}

