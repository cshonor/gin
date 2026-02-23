// 对应视频: 21.局部中间件
// 学习目标: 路由组/单路由的局部中间件
//
// gin.HandlerFunc = func(*gin.Context)，Gin 的处理器/中间件类型
// 返回 HandlerFunc 的工厂函数（如 Logger()、Auth()）可被 r.Use() 或单路由使用
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 返回 gin.HandlerFunc，作为中间件仅对挂载它的路由生效
/*gin.HandlerFunc 是 Gin 对 func(*gin.Context) 的类型别名：  
 • 作为路由处理函数：r.GET("/path", func(c *gin.Context) { ... })  
 • 作为中间件：r.Use(Logger()) 或 r.GET("/path", Logger(), handler)*/

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		fmt.Printf("[%s] %d | %13v | %15s | %-7s %s\n",
			time.Now().Format("2006/01/02 15:04:05"),
			statusCode, latency, clientIP, method, path,
		)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "需要登录"})
			c.Abort()
			return
		}
		c.Set("user_id", token)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 公开路由 - 无中间件
	r.GET("/public", func(c *gin.Context) {
		c.String(200, "公开接口")
	})

	// 单路由使用中间件
	r.GET("/with-logger", Logger(), func(c *gin.Context) {
		c.String(200, "带日志的路由")
	})

	// 路由组 - 局部中间件（仅该组生效）
	private := r.Group("/api")
	private.Use(Auth())
	{
		private.GET("/profile", func(c *gin.Context) {
			uid, _ := c.Get("user_id")
			c.JSON(200, gin.H{"user_id": uid})
		})
		private.GET("/settings", func(c *gin.Context) {
			c.String(200, "设置页面")
		})
	}

	// 路由组 - 多个中间件
	admin := r.Group("/admin")
	admin.Use(Logger(), Auth())
	{
		admin.GET("/users", func(c *gin.Context) {
			c.String(200, "管理员-用户列表")
		})
	}

	// 链式：单路由多中间件
	r.GET("/chain", Logger(), Auth(), func(c *gin.Context) {
		c.String(200, "链式中间件")
	})

	r.Run(":8080")
}

