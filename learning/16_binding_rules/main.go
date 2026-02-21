// 对应视频: 16.binding内置规则
// 学习目标: validator 内置验证规则
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 常用 binding 标签规则:
// required - 必填
// omitempty - 空值忽略
// min=N, max=N - 最小/最大长度或数值
// len=N - 固定长度
// oneof=a b c - 枚举值
// email, url, uri - 格式验证
// gte=N, lte=N - 大于等于/小于等于
// gt=N, lt=N - 大于/小于
// dive - 深入验证嵌套
// dive,required - 嵌套必填

type User struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Age      int    `json:"age" binding:"gte=0,lte=120"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"oneof=admin user guest"`
	Website  string `json:"website" binding:"omitempty,url"`
}

type PageQuery struct {
	Page     int    `form:"page" binding:"gte=1"`
	PageSize int    `form:"page_size" binding:"gte=1,lte=100"`
	Sort     string `form:"sort" binding:"omitempty,oneof=asc desc"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, u)
	})

	r.GET("/list", func(c *gin.Context) {
		var q PageQuery
		if err := c.ShouldBindQuery(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if q.Page == 0 {
			q.Page = 1
		}
		if q.PageSize == 0 {
			q.PageSize = 10
		}
		c.JSON(http.StatusOK, q)
	})

	r.Run(":8080")
}

