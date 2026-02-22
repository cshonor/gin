// 对应视频: 9.静态文件
// 学习目标: Gin 静态文件服务
//
// r.Static(URL前缀, 本地目录) - 将整个目录映射为静态资源
// 常用在提供 CSS、JS、图片等；只提供单个文件时用 r.StaticFile
// 例: r.Static("/assets", "./static") 表示 /assets/xxx 访问 ./static/xxx
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// c.Redirect(状态码, URL) - 重定向到新地址
	// 301 永久跳转，302 临时跳转；访问 / 时跳转到 /assets/index.html
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/assets/index.html")
	})

	// 1. Static - 静态目录映射
	// 访问: http://localhost:8080/assets/xxx
	r.Static("/assets", "learning/09_static_files/static")

	// 2. StaticFS - 使用 embed 或自定义 FileSystem（更灵活）
	// r.StaticFS("/files", http.Dir("./static"))

	// 3. StaticFile - 单个文件
	r.StaticFile("/index", "learning/09_static_files/static/index.html")

	// 4. 多组静态文件
	r.Static("/css", "learning/09_static_files/static/css")
	r.Static("/js", "learning/09_static_files/static/js")

	r.Run(":8080")
}

