// 对应视频: 2.casbin配置文件
// 从 model.conf 和 policy.csv 加载，配置文件定义权限模型和初始策略
package main

import (
	"fmt"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 从配置文件加载，路径相对于运行目录
	modelPath := filepath.Join("learning", "casbin", "02_config", "model.conf")
	policyPath := filepath.Join("learning", "casbin", "02_config", "policy.csv")

	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}

	ok, _ := e.Enforce("admin", "data1", "write")
	fmt.Println("admin write data1:", ok) // true

	ok, _ = e.Enforce("alice", "data1", "write")
	fmt.Println("alice write data1:", ok) // false
}

