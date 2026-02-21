// 对应视频: 10.查询参数和动态参数
// 学习目标: Gin 中查询参数、路径参数的获取与使用
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 路径参数 - :param 必填
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "用户ID: %s", id)
	})

	// 2. 路径参数 - *path 通配符
	r.GET("/files/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		c.String(http.StatusOK, "文件路径: %s", filepath)
	})

	// 3. 查询参数 - Query
	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")           // 必取，无则空
		page := c.DefaultQuery("page", "1")     // 默认值
		size := c.DefaultQuery("page_size", "10")

		c.JSON(http.StatusOK, gin.H{
			"keyword":   keyword,
			"page":      page,
			"page_size": size,
		})
	})

	// 4. 获取多个同名查询参数 - QueryArray
	r.GET("/tags", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	})

	// 5. 查询参数 Map - QueryMap
	r.GET("/filters", func(c *gin.Context) {
		filters := c.QueryMap("filter")
		c.JSON(http.StatusOK, gin.H{"filters": filters})
	})

	// 6. 组合使用
	r.GET("/post/:id/comments", func(c *gin.Context) {
		id := c.Param("id")
		sort := c.DefaultQuery("sort", "desc")
		limit := c.DefaultQuery("limit", "20")
		c.JSON(http.StatusOK, gin.H{
			"post_id": id,
			"sort":    sort,
			"limit":   limit,
		})
	})

	r.Run(":8080")
}

