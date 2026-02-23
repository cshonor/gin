// 对应视频: 13.原始内容
// 学习目标: 获取请求原始 Body、Request 对象
//
// io.ReadAll(r) - 从 io.Reader 读取全部内容直到 EOF，返回 ([]byte, error)
// c.Request.Body 实现了 io.Reader，ReadAll 一次读完整请求体
// 注意: 读完后 Body 被消费，无法再次读取
//
// defer 延迟执行: 等当前函数返回前再执行，常用于关闭资源、释放锁
// defer c.Request.Body.Close() 确保函数退出时关闭 Body
package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 获取原始 Body
	r.POST("/raw", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body) // 一次性读完整 body
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer c.Request.Body.Close() // 函数结束前执行，关闭 Body 释放连接
		/*
		defer 表示延迟执行：在当前函数即将返回前再执行这条语句。
        顺序：多个 defer 按后进先出执行，最后写的先执行。
        常见用途：关闭文件、关闭连接、释放锁等，保证函数退出时一定会执行。
       */

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

