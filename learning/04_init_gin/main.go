// 对应视频: 4.初始gin框架
// 学习目标: 掌握 Gin 框架的基本初始化和路由注册
//
// c.String(200, "文本")  - 返回纯文本，Content-Type: text/plain
// c.JSON(200, obj)      - 返回 JSON，自动设置 Content-Type 并序列化结构体/map
// gin.H 是 map[string]any 的简写，用于快速构建 JSON 对象，如 gin.H{"code":0,"msg":"ok"}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 引擎（包含 Logger 和 Recovery 中间件）
	r := gin.Default()

	// 注册路由: GET 请求
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎使用 Gin 框架！")
	})

	// c.JSON: 返回 JSON，自动序列化
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "成功",
			"data": gin.H{"name": "张三"},
		})
	})

	// 路径参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})

	// 通配符参数
	r.GET("/view/*path", func(c *gin.Context) {
		path := c.Param("path")
		c.String(http.StatusOK, "查看路径: %s", path)
	})

	// 查询参数
	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")
		page := c.DefaultQuery("page", "1")
		c.String(http.StatusOK, "搜索: %s, 页码: %s", keyword, page)
	})

	// 启动服务
	r.Run(":8080") // 默认 0.0.0.0:8080
}

