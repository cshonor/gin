// 对应视频: 4.初始gin框架
// 学习目标: 掌握 Gin 框架的基本初始化和路由注册
//
// c.String vs c.JSON 区别:
//   c.String(200, "文本")  - 原样输出，Content-Type: text/plain，适合纯文本
//   c.JSON(200, obj)      - 序列化 obj 为 JSON 再输出，Content-Type: application/json，适合 API
//
// c.JSON vs json.Marshal 区别:
//   c.JSON  - Gin 封装，一步完成序列化+写响应+设 Content-Type
//   json.Marshal - 标准库，只序列化得 []byte，需手动 w.Write、设 Header
//
// gin.H 与 c.JSON 配合使用: gin.H 是数据，c.JSON 负责序列化并返回
// gin.H = map[string]any，可嵌套。c.JSON 的 obj 也可以是结构体、普通 map

/*
gin.H：本质是 map[string]interface{} 的别名，
是 Gin 框架提供的便捷工具，
用来快速构造键值对形式的 JSON 数据结构（数据载体）。

c.JSON：Gin 上下文（Context）的方法，
作用是设置 HTTP 响应头（Content-Type 为 application/json），
并将传入的数据序列化为 JSON 字符串返回给客户端。

*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode 设置运行模式，需在 gin.Default() 之前调用
	// gin.DebugMode(默认) / gin.ReleaseMode(生产) / gin.TestMode(测试)
	// 生产环境用 ReleaseMode 可减少日志输出、提升性能
	// gin.SetMode(gin.ReleaseMode)

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

