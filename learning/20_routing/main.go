// 对应视频: 20.路由
// 学习目标: Gin 路由组、路由匹配、Any、NoRoute 等
//
// r.Group("/api/v1") - 路由分组，返回 *gin.RouterGroup
// 挂到分组上的路由自动带前缀，如 v1.GET("/users") 实际路径为 /api/v1/users
// 分组可加中间件: r.Group("/api/v1", AuthMiddleware()) 该组下所有接口先执行鉴权
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 路由组 - 统一前缀，实际路径 /api/v1/users、/api/v1/users/:id
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) { c.String(200, "用户列表") })
		v1.GET("/users/:id", func(c *gin.Context) { c.String(200, "用户详情") })
	}

	// 2. 路由组 - 嵌套
	admin := r.Group("/admin")
	{
		admin.GET("/", func(c *gin.Context) { c.String(200, "管理后台") })
		users := admin.Group("/users")
		{
			users.GET("/", func(c *gin.Context) { c.String(200, "管理员-用户列表") })
			users.POST("/", func(c *gin.Context) { c.String(200, "创建用户") })
		}
	}

	// 3. Any - 匹配所有 HTTP 方法
	r.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": c.Request.Method})
	})

	// 4. 匹配多个路径
	r.GET("/home", indexHandler)
	r.GET("/index", indexHandler)

	// 5. NoRoute - 404 处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "接口不存在",
			"path":    c.Request.URL.Path,
		})
	})

	// 6. 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://gin-gonic.com")
	})

	// 7. 路由转发
	r.GET("/go", func(c *gin.Context) {
		c.Request.URL.Path = "/api/v1/users"
		r.HandleContext(c)
	})

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.String(200, "首页")
}

