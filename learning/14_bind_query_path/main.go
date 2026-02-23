// 对应视频: 14.bind参数绑定-查询参数,路径参数
// 学习目标: ShouldBindQuery、ShouldBindUri 参数绑定
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryParams 查询参数绑定结构体
// form:"keyword" 表示 URL 里 ?keyword=xxx 会赋给 Keyword 字段
// 类型会自动转换: page=2 -> int, 无则零值(string="", int=0)
type QueryParams struct {
	Keyword  string `form:"keyword"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Sort     string `form:"sort"`
}

// PathParams 路径参数绑定
type PathParams struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name"`
}

// QueryAndPath 组合绑定
type PaginationQuery struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderBy  string `form:"order_by"`
}

func main() {
	r := gin.Default()

	// 1. ShouldBindQuery - 将 ?keyword=xx&page=2 自动绑定到结构体
	// 例: /search?keyword=gin&page=2&page_size=20 -> params={Keyword:"gin", Page:2, PageSize:20}
	r.GET("/search", func(c *gin.Context) {
		var params QueryParams
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if params.Page == 0 {
			params.Page = 1
		}
		if params.PageSize == 0 {
			params.PageSize = 10
		}
		c.JSON(http.StatusOK, params)
	})

	// 2. ShouldBindUri - 绑定路径参数
	r.GET("/user/:id", func(c *gin.Context) {
		var params PathParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, params)
	})

	// 3. ShouldBindUri - 多个路径参数
	r.GET("/post/:id/comment/:cid", func(c *gin.Context) {
		var params struct {
			ID  string `uri:"id" binding:"required"`
			Cid string `uri:"cid" binding:"required"`
		}
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, params)
	})

	// 4. 组合使用 Query + Uri
	r.GET("/article/:id", func(c *gin.Context) {
		var uriParams struct {
			ID string `uri:"id" binding:"required"`
		}
		var queryParams PaginationQuery
		if err := c.ShouldBindUri(&uriParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "路径参数错误"})
			return
		}
		c.ShouldBindQuery(&queryParams)
		c.JSON(http.StatusOK, gin.H{
			"id":        uriParams.ID,
			"page":      queryParams.Page,
			"page_size": queryParams.PageSize,
			"order_by":  queryParams.OrderBy,
		})
	})

	r.Run(":8080")
}

