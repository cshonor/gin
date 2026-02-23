// 对应视频: 2.casbin配置文件
// model.conf = 模型（如何判断权限）；policy.csv = 策略数据（谁对什么有什么操作）
package main

import (
	"fmt"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

func main() {
	// filepath.Join 将多个路径片段拼接成符合当前系统的完整路径
	// 例: Join("user","docs","report.txt") -> user/docs/report.txt (Linux) 或 user\docs\report.txt (Win)
	// 自动去重: Join("data//","logs/","app.log") -> data/logs/app.log
	// 路径简化: Join("a","b","..","c",".","d.txt") -> a/c/d.txt
	// 路径相对于运行目录(项目根)
	modelPath := filepath.Join("learning", "casbin", "02_config", "model.conf")
	policyPath := filepath.Join("learning", "casbin", "02_config", "policy.csv")

	// casbin.NewEnforcer 初始化权限执行器
	// modelPath: 模型文件(.conf)，定义权限规则结构(RBAC/ACL、matcher 等)
	// policyPath: 策略文件(.csv)，定义具体权限(谁对什么有什么操作)
	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}

	// e.Enforce(sub, obj, act) 检查是否有权限，返回 (bool, error)
	ok, _ := e.Enforce("admin", "data1", "write")
	fmt.Println("admin write data1:", ok) // true

	ok, _ = e.Enforce("alice", "data1", "write")
	fmt.Println("alice write data1:", ok) // false
}

