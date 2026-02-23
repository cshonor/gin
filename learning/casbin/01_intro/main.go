// 对应视频: 1.casbin基本介绍
// Casbin: 权限管理库，支持 ACL、RBAC、ABAC 等模型
// 核心: 定义「谁(subject)对什么(object)有什么操作(action)」
package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// NewEnforcer() 无参时: 使用内置 ACL 模型 + 内存适配器
	// 策略初始为空，没有默认策略，需通过 AddPolicy 手动添加
	e, err := casbin.NewEnforcer()
	if err != nil {
		panic(err)
	}

	// e.AddPolicy(sub, obj, act) - 添加一条策略
	// sub: subject 主体，如用户 alice
	// obj: object 客体/资源，如 data1、/api/user、订单等，不是前端 CSS/JS 资源
	// act: action 操作，如 read、write
	e.AddPolicy("alice", "data1", "read")
	e.AddPolicy("bob", "data2", "write")

	// e.Enforce(sub, obj, act) - 检查 (sub, obj, act) 是否在策略中，返回 (allowed bool, err error)
	// true=有权限，false=无权限
	ok, _ := e.Enforce("alice", "data1", "read")
	fmt.Println("alice read data1:", ok) // true

	ok, _ = e.Enforce("alice", "data2", "write")
	fmt.Println("alice write data2:", ok) // false

	ok, _ = e.Enforce("bob", "data2", "write")
	fmt.Println("bob write data2:", ok) // true
}

