// 对应视频: 22.全局中间件
// 学习目标: 全局中间件、gin.Default() 内置中间件、自定义全局中间件
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局中间件1: 请求日志
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		fmt.Printf("[全局] %s %s %v\n", method, path, time.Since(start))
	}
}

// 全局中间件2: CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 全局中间件3: 请求 ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// 全局中间件4: 恢复 panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "服务器内部错误",
					"msg":   fmt.Sprintf("%v", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func main() {
	// gin.Default() 已包含 Logger 和 Recovery 中间件
	// r := gin.Default()

	// gin.New() 不含任何中间件，可自定义
	r := gin.New()

	// 注册全局中间件（按顺序执行）
	r.Use(Recovery())   // 最先：捕获 panic
	r.Use(RequestID())  // 每个请求生成 ID
	r.Use(CORS())       // 跨域
	r.Use(RequestLogger()) // 请求日志

	r.GET("/", func(c *gin.Context) {
		rid, _ := c.Get("request_id")
		c.JSON(200, gin.H{
			"message":    "Hello",
			"request_id": rid,
		})
	})

	// 测试 panic 恢复
	r.GET("/panic", func(c *gin.Context) {
		panic("故意触发的 panic")
	})

	r.Run(":8080")
}

