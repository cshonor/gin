// 对应视频: 8.响应文件
// 学习目标: Gin 的文件响应 - 下载、静态文件服务
//
// r.Static(URL前缀, 本地目录) - 目录下所有文件可通过 URL 直接访问，如 /static/example.txt
package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 静态文件服务 - 从项目根目录运行
	// 访问: http://localhost:8080/static/example.txt
	r.Static("/static", "learning/08_response_file/assets")

	// 2. 文件下载 - 触发浏览器「另存为」而非直接预览
	// Content-Disposition: attachment 表示下载；filename 指定保存时的文件名
	// filepath.Base 从路径中提取文件名，如 /a/b/c.txt -> c.txt
	// c.File 读取文件并写入响应体
	/*没有 Content-Disposition: attachment 时，浏览器通常会直接显示（如 txt、pdf），而不是下载。*/
	r.GET("/download", func(c *gin.Context) {
		filePath := "learning/08_response_file/assets/example.txt"
		fileName := filepath.Base(filePath)
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.File(filePath)
	})

	// 3. 文件下载 - 自定义文件名
	r.GET("/download-as", func(c *gin.Context) {
		filePath := "learning/08_response_file/assets/example.txt"
		c.Header("Content-Disposition", "attachment; filename=自定义文件名.txt")
		c.File(filePath)
	})

	// 4. FileAttachment - 带自定义文件名的下载
	r.GET("/read", func(c *gin.Context) {
		c.FileAttachment("learning/08_response_file/assets/example.txt", "下载文件.txt")
	})

	r.Run(":8080")
}

