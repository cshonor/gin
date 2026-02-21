package routes

import (
	"gin/handlers"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(r *gin.Engine) {
	// 应用全局中间件
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// 基础路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "欢迎使用 Gin 框架！",
			"version": "1.0.0",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 用户路由（公开访问）
		users := v1.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.GET("/:id", handlers.GetUser)
			users.POST("", handlers.CreateUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		// 需要认证的路由示例
		protected := v1.Group("/protected")
		protected.Use(middleware.Auth())
		{
			protected.GET("/info", func(c *gin.Context) {
				token, _ := c.Get("token")
				c.JSON(200, gin.H{
					"message": "这是受保护的路由",
					"token":   token,
				})
			})
		}
	}
}


