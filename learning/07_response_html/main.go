// 对应视频: 7.响应html
// 学习目标: Gin 渲染 HTML 模板
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载模板文件
	// 从项目根目录运行: go run ./learning/07_response_html
	r.LoadHTMLGlob("learning/07_response_html/templates/*")

	// 渲染 HTML 页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Gin 学习",
			"message": "欢迎学习 Gin 框架！",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.HTML(http.StatusOK, "user.html", gin.H{
			"name": name,
		})
	})

	// 使用自定义模板
	r.GET("/custom", func(c *gin.Context) {
		c.HTML(http.StatusOK, "custom.html", gin.H{
			"items": []string{"Go", "Gin", "GORM"},
		})
	})

	r.Run(":8080")
}

