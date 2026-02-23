// ABAC = Attribute-Based Access Control，基于属性的访问控制
// sub、obj 传结构体，matcher 通过 r.sub.xxx、r.obj.xxx 访问属性判断权限
package main

import (
	"fmt"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

// Subject 主体属性
type Subject struct {
	Role       string // admin | user
	Department string // 部门
}

// Object 资源属性
type Object struct {
	Name       string
	Department string // 所属部门
}

func main() {
	model := filepath.Join("learning", "casbin", "07_abac", "model.conf")
	e, _ := casbin.NewEnforcer(model)

	// 纯 ABAC 需至少一条占位策略，matcher 主要靠属性判断
	e.AddPolicy("_", "_", "_")

	// admin 可做任何操作
	admin := Subject{Role: "admin", Department: "it"}
	doc1 := Object{Name: "doc1", Department: "sales"}
	fmt.Println("admin read sales文档:", mustEnforce(e, admin, doc1, "read"))   // true
	fmt.Println("admin write sales文档:", mustEnforce(e, admin, doc1, "write")) // true

	// user 只能读同部门
	user := Subject{Role: "user", Department: "it"}
	docIt := Object{Name: "docIt", Department: "it"}
	docSales := Object{Name: "docSales", Department: "sales"}
	fmt.Println("user(it) read it文档:", mustEnforce(e, user, docIt, "read"))     // true
	fmt.Println("user(it) read sales文档:", mustEnforce(e, user, docSales, "read")) // false
	fmt.Println("user(it) write it文档:", mustEnforce(e, user, docIt, "write"))   // false
}

func mustEnforce(e *casbin.Enforcer, sub Subject, obj Object, act string) bool {
	ok, _ := e.Enforce(sub, obj, act)
	return ok
}

