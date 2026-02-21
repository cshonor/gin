// 对应视频: 15.binding参数-json参数和header参数
// 学习目标: ShouldBindJSON、Header 绑定
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 请求体
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"gte=0,lte=120"`
}

// Header 绑定
type AuthHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
	Token         string `header:"X-Token"`
}

// 组合 JSON + Header
type ApiRequest struct {
	TraceID string `header:"X-Trace-Id"`
}

func main() {
	r := gin.Default()

	// 1. ShouldBindJSON - 绑定 JSON 请求体
	r.POST("/user", func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "参数错误",
				"detail": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "创建成功",
			"data":    req,
		})
	})

	// 2. ShouldBindHeader - 绑定请求头
	r.GET("/auth", func(c *gin.Context) {
		var h AuthHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 Authorization"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "认证成功",
			"token":   h.Token,
		})
	})

	// 3. 同时绑定 JSON + Header
	r.POST("/api/request", func(c *gin.Context) {
		var header ApiRequest
		var body CreateUserRequest
		c.ShouldBindHeader(&header)
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"trace_id": header.TraceID,
			"user":     body,
		})
	})

	// 4. ShouldBind - 自动根据 Content-Type 选择绑定方式
	r.POST("/auto", func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, req)
	})

	r.Run(":8080")
}

