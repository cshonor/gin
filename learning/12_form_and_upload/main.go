// 对应视频: 12.表单参数与文件上传
// 学习目标: 表单数据解析、文件上传
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 表单参数 - PostForm
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"message":  "登录成功",
		})
	})

	// 2. 表单默认值 - DefaultPostForm
	r.POST("/register", func(c *gin.Context) {
		username := c.DefaultPostForm("username", "guest")
		email := c.PostForm("email")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
		})
	})

	// 3. 单文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请选择文件"})
			return
		}

		// 保存文件
		dst := filepath.Join("learning/12_form_and_upload/uploads", file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "上传成功",
			"filename": file.Filename,
			"size":     file.Size,
		})
	})

	// 4. 多文件上传
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		var filenames []string
		for _, file := range files {
			dst := filepath.Join("learning/12_form_and_upload/uploads", file.Filename)
			c.SaveUploadedFile(file, dst)
			filenames = append(filenames, file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "批量上传成功",
			"filenames": filenames,
		})
	})

	// 5. 表单 + 文件 混合
	r.POST("/profile", func(c *gin.Context) {
		name := c.PostForm("name")
		file, _ := c.FormFile("avatar")

		msg := fmt.Sprintf("用户 %s 提交了资料", name)
		if file != nil {
			dst := filepath.Join("learning/12_form_and_upload/uploads", file.Filename)
			c.SaveUploadedFile(file, dst)
			msg += fmt.Sprintf(", 头像: %s", file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{"message": msg})
	})

	r.Run(":8080")
}

