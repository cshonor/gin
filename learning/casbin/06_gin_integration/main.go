// 06. Casbin 与 Gin 配合
// 用中间件在请求前做权限校验：sub=用户 obj=路径 act=方法
package main

import (
	"net/http"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	model := filepath.Join("learning", "casbin", "06_gin_integration", "model.conf")
	e, _ := casbin.NewEnforcer(model)

	// 策略: 角色->路径->方法。keyMatch2 支持 :id 通配
	e.AddPolicy("admin", "/api/*", "*")
	e.AddPolicy("user", "/api/users", "GET")
	e.AddPolicy("user", "/api/users/:id", "GET")
	e.AddGroupingPolicy("alice", "admin")
	e.AddGroupingPolicy("bob", "user")

	r := gin.Default()

	// Casbin 中间件：从 Header 取用户，校验 sub+obj+act
	r.Use(func(c *gin.Context) {
		user := c.GetHeader("X-User")
		if user == "" {
			user = "anonymous"
		}
		path := c.Request.URL.Path
		act := c.Request.Method

		ok, _ := e.Enforce(user, path, act)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无权限"})
			return
		}
		c.Next()
	})

	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "用户列表"})
	})
	r.GET("/api/users/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "用户详情"})
	})
	r.POST("/api/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "创建用户"})
	})

	r.Run(":8080")
}

