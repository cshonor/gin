// 对应视频: 6.gin响应json-响应封装2
// 学习目标: 更完善的响应封装，支持分页等
//
// 序列化: Result 内部调用 c.JSON，Gin 将 Response 结构体序列化为 JSON 返回
// 例: Result(c,200,"success",data) -> {"code":200,"message":"success","data":{...}}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 基础响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageData 分页数据结构
type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// Result 统一响应，c.JSON 将 Response 序列化为 JSON 写入响应体
func Result(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Ok 成功
func Ok(c *gin.Context) {
	Result(c, 200, "success", nil)
}

// OkWithData 成功带数据
func OkWithData(c *gin.Context, data interface{}) {
	Result(c, 200, "success", data)
}

// OkWithPage 成功带分页数据
func OkWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	Result(c, 200, "success", PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// Fail 失败
func Fail(c *gin.Context, message string) {
	Result(c, 500, message, nil)
}

func main() {
	r := gin.Default()

	r.GET("/ok", func(c *gin.Context) {
		Ok(c)
	})

	r.GET("/user", func(c *gin.Context) {
		user := map[string]interface{}{
			"id":   1,
			"name": "李四",
		}
		OkWithData(c, user)
	})

	r.GET("/users", func(c *gin.Context) {
		// 模拟分页数据
		list := []map[string]interface{}{
			{"id": 1, "name": "用户1"},
			{"id": 2, "name": "用户2"},
		}
		OkWithPage(c, list, 100, 1, 10)
	})

	r.GET("/fail", func(c *gin.Context) {
		Fail(c, "操作失败")
	})

	r.Run(":8080")
}

