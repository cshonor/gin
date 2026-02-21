// 对应视频: 5.gin响应-json响应封装
// 学习目标: Gin 的 JSON 响应及统一响应格式封装
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// 失败响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		// 方式1: 直接返回 JSON
		// c.JSON(200, gin.H{"name": "张三", "age": 25})

		// 方式2: 使用封装的成功响应
		Success(c, gin.H{
			"id":   1,
			"name": "张三",
			"age":  25,
		})
	})

	r.GET("/error", func(c *gin.Context) {
		Error(c, 500, "服务器内部错误")
	})

	// 绑定 JSON 到结构体
	r.POST("/user", func(c *gin.Context) {
		var user struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			Error(c, 400, "参数错误: "+err.Error())
			return
		}
		Success(c, user)
	})

	r.Run(":8080")
}

