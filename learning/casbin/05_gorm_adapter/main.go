// 对应视频: 5.gorm接入casbin
// 使用 gorm-adapter 将策略存到数据库，重启后策略持久化
package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 创建 SQLite 数据库
	db, _ := gorm.Open(sqlite.Open("learning/casbin/05_gorm_adapter/casbin.db"), &gorm.Config{})

	// 创建 gorm adapter，会自动建表 casbin_rule
	a, _ := gormadapter.NewAdapterByDB(db)

	// 使用内置 ACL 模型 + gorm adapter
	e, _ := casbin.NewEnforcer("learning/casbin/03_acl/model.conf", a)

	// 添加策略会写入数据库
	e.AddPolicy("admin", "/api/config", "write")
	e.SavePolicy()

	ok, _ := e.Enforce("admin", "/api/config", "write")
	fmt.Println("admin write /api/config:", ok)

	// 重启后从数据库加载，策略仍在
}

