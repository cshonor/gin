// 对应视频: 7.响应html
// 学习目标: Gin 渲染 HTML 模板
//
// c.HTML = 传统 SSR（服务端渲染），和后端拼好完整 HTML 再返回
// 与 Next.js SSR 原理相同，都是服务端渲染，只是技术栈不同（Go 模板 vs React）
// 当前 Go 后端多用于纯 API，SSR 多由 Next.js/Nuxt 等前端框架负责
//
// c.HTML(200, "模板名.html", gin.H{...}) - 渲染模板并返回，模板中用 {{.key}} 取值
//
// --- Go 后端返回类型一览 ---
// c.String  - 纯文本，Content-Type: text/plain
// c.JSON    - JSON，API 最常用
// c.HTML    - HTML 页面，传统 SSR
// c.File    - 文件下载/静态文件
// c.Redirect - 重定向
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

