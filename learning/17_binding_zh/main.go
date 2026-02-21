// 对应视频: 17.binding错误信息显示中文
// 学习目标: 将 validator 错误信息翻译为中文
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 字段名中文映射
var fieldLabels = map[string]string{
	"Name":  "用户名",
	"Email": "邮箱",
	"Age":   "年龄",
}

// 验证规则中文映射
var ruleMessages = map[string]string{
	"required": "为必填项",
	"min":      "长度至少为 %s",
	"max":      "长度最多为 %s",
	"email":    "必须是有效的邮箱格式",
	"gte":      "必须大于等于 %s",
	"lte":      "必须小于等于 %s",
}

func getLabel(field string) string {
	if label, ok := fieldLabels[field]; ok {
		return label
	}
	return field
}

func translateError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var msgs []string
		for _, e := range errs {
			field := getLabel(e.Field())
			rule := e.Tag()
			if msg, ok := ruleMessages[rule]; ok {
				if strings.Contains(msg, "%s") {
					msgs = append(msgs, field+fmt.Sprintf(msg, e.Param()))
				} else {
					msgs = append(msgs, field+msg)
				}
			} else {
				msgs = append(msgs, field+": "+e.Error())
			}
		}
		return strings.Join(msgs, "; ")
	}
	return err.Error()
}

type User struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=0,lte=120"`
}

func main() {
	r := gin.Default()

	// 注册 TagNameFunc 使用 label 标签显示字段名
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			if label := fld.Tag.Get("label"); label != "" {
				return label
			}
			return fld.Name
		})
	}

	r.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": translateError(err),
			})
			return
		}
		c.JSON(http.StatusOK, u)
	})

	r.Run(":8080")
}

