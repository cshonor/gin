// 对应视频: 10.查询参数和动态参数
// 学习目标: Gin 中查询参数、路径参数的获取与使用
//
// :param 与 *path 区别:
//   :param  - 只匹配一段，不含 /。如 /user/:id 中 id 匹配 "123" 不匹配 "a/b"
//   *path  - 匹配剩余全部，含 /。如 /files/*fp 中 fp 匹配 "/a/b/c"，注意带前导 /
//
// 路由竞争优先级: 字面 > :param > *path
// 如 /user/new、/user/:id、/user/*path 共存时，/user/new 走字面，/user/123 走 :id，/user/a/b 走 *path
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. :id 匹配路径中的一段，/user/123 时 c.Param("id") = "123"
	// c.Param 返回 string，需数字时用 strconv.Atoi(id) 转换
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "用户ID: %s", id)
	})

	// 2. *filepath 通配符，匹配后面所有，/files/a/b/c 时 c.Param("filepath") = "/a/b/c"
	r.GET("/files/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		c.String(http.StatusOK, "文件路径: %s", filepath)
	})

	// 3. 查询参数 - Query
	r.GET("/search", func(c *gin.Context) {
		// Query 无则返回 ""（空字符串），Go 的 string 不能为 nil
		// DefaultQuery 无则返回第二个参数
		keyword := c.Query("keyword")           // ?keyword=xxx，无则 ""
		page := c.DefaultQuery("page", "1")     // ?page=2，无则 "1"
		size := c.DefaultQuery("page_size", "10")

		c.JSON(http.StatusOK, gin.H{
			"keyword":   keyword,
			"page":      page,
			"page_size": size,
		})
	})

	// 4. QueryArray - 同名参数多个，如 ?tag=a&tag=b
	// 例: /tags?tag=go&tag=gin&tag=gorm -> tags=["go","gin","gorm"]
	r.GET("/tags", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	})

	// 5. QueryMap("filter") - 收集所有 filter[key]=value 形式的参数
	// "filter" 是约定俗成的命名，这类筛选用的 map 常叫 filter；可改成 search 等
	// 例: ?filter[name]=张三&filter[age]=25 -> map[name:张三 age:25]
	/*QueryMap 常用于需要多个可选筛选条件的场景：
后台管理列表：按状态、类型、部门等筛选，如 ?filter[status]=1&filter[type]=vip
搜索页面：动态组合多个筛选条件，如 ?filter[keyword]=xx&filter[category]=yy*/
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

