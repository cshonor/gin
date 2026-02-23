// 对应视频: 18.binding显示错误字段和错误信息
// 学习目标: 返回结构化的验证错误，包含字段名和具体信息
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// FieldError 单个字段错误
// json:"value,omitempty" - omitempty：值为零值(空串/0/nil)时，序列化时省略该字段，不输出到 JSON
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// ValidationErrorResponse 验证错误响应
type ValidationErrorResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors"`
}

func parseValidationErrors(err error) ValidationErrorResponse {
	resp := ValidationErrorResponse{
		Code:    400,
		Message: "参数验证失败",
		Errors:  []FieldError{},
	}
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			resp.Errors = append(resp.Errors, FieldError{
				Field:   e.Field(),
				Message: getErrorMessage(e.Tag(), e.Param()),
				Value:   fmt.Sprint(e.Value()),
			})
		}
	}
	return resp
}

func getErrorMessage(tag, param string) string {
	messages := map[string]string{
		"required": "不能为空",
		"min":      "最小长度为 " + param,
		"max":      "最大长度为 " + param,
		"email":    "邮箱格式不正确",
		"gte":      "必须 >= " + param,
		"lte":      "必须 <= " + param,
		"oneof":    "必须是 " + param + " 中的一个",
	}
	if msg, ok := messages[tag]; ok {
		return msg
	}
	return tag
}

type User struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=0,lte=120"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			resp := parseValidationErrors(err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		c.JSON(http.StatusOK, u)
	})

	r.Run(":8080")
}

