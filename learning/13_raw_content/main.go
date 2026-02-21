// 对应视频: 13.原始内容
// 学习目标: 获取请求原始 Body、Request 对象
package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 获取原始 Body - c.Request.Body
	r.POST("/raw", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer c.Request.Body.Close()

		c.JSON(http.StatusOK, gin.H{
			"raw":       string(body),
			"content_type": c.GetHeader("Content-Type"),
		})
	})

	// 2. 获取 Request 相关信息
	r.Any("/request-info", func(c *gin.Context) {
		r := c.Request
		c.JSON(http.StatusOK, gin.H{
			"method":     r.Method,
			"url":        r.URL.String(),
			"host":       r.Host,
			"remote_addr": r.RemoteAddr,
			"header":     r.Header,
			"content_length": r.ContentLength,
		})
	})

	// 3. 获取特定 Header
	r.GET("/headers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user_agent":   c.GetHeader("User-Agent"),
			"content_type": c.GetHeader("Content-Type"),
			"authorization": c.GetHeader("Authorization"),
		})
	})

	// 4. 复制 Body 用于多次读取（如需解析多种格式）
	r.POST("/multi-parse", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()

		// 可以多次使用 body 进行不同解析
		raw := string(body)
		c.JSON(http.StatusOK, gin.H{
			"length": len(raw),
			"preview": func() string {
				if len(raw) > 100 {
					return raw[:100] + "..."
				}
				return raw
			}(),
		})
	})

	r.Run(":8080")
}

