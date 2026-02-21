// 对应视频: 19.自定义验证器
// 学习目标: 注册自定义 validator 验证函数
package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义验证: 手机号
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// 简单规则: 11位数字
	matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, phone)
	return matched
}

// 自定义验证: 不能包含空格
func validateNoSpace(fl validator.FieldLevel) bool {
	return !strings.Contains(fl.Field().String(), " ")
}

// 自定义验证: 密码强度
func validateStrongPassword(fl validator.FieldLevel) bool {
	pwd := fl.Field().String()
	if len(pwd) < 8 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pwd)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pwd)
	hasDigit := regexp.MustCompile(`\d`).MatchString(pwd)
	return hasUpper && hasLower && hasDigit
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,no_space"`
	Phone    string `json:"phone" binding:"required,phone"`
	Password string `json:"password" binding:"required,strong_password"`
}

func main() {
	r := gin.Default()

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", validatePhone)
		v.RegisterValidation("no_space", validateNoSpace)
		v.RegisterValidation("strong_password", validateStrongPassword)
	}

	r.POST("/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"user":    req,
		})
	})

	r.Run(":8080")
}

