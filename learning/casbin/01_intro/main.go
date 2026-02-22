// 对应视频: 1.casbin基本介绍
// Casbin: 权限管理库，支持 ACL、RBAC、ABAC 等模型
// 核心: 定义「谁(subject)对什么(object)有什么操作(action)」
package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 使用内置模型和内存适配器，适合快速体验
	e, err := casbin.NewEnforcer()
	if err != nil {
		panic(err)
	}

	// 添加策略: 用户 alice 可以对 data1 进行 read
	e.AddPolicy("alice", "data1", "read")
	e.AddPolicy("bob", "data2", "write")

	// 检查权限: Enforce(sub, obj, act)
	ok, _ := e.Enforce("alice", "data1", "read")
	fmt.Println("alice read data1:", ok) // true

	ok, _ = e.Enforce("alice", "data2", "write")
	fmt.Println("alice write data2:", ok) // false

	ok, _ = e.Enforce("bob", "data2", "write")
	fmt.Println("bob write data2:", ok) // true
}

