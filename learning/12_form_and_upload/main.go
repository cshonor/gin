// 对应视频: 12.表单参数与文件上传
// 学习目标: 表单数据解析、文件上传
//
// c.PostForm("key") - 从表单 body 取字段，无则 ""
// c.DefaultPostForm("key", "默认") - 无则返回第二个参数
// c.FormFile("key") - 取上传的文件，返回 *multipart.FileHeader
// c.SaveUploadedFile(file, dst) - 将上传文件保存到本地路径 dst
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. PostForm - 从表单 body 取 username、password
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"message":  "登录成功",
		})
	})

	// 2. DefaultPostForm - 无 username 时返回 "guest"
	r.POST("/register", func(c *gin.Context) {
		username := c.DefaultPostForm("username", "guest")
		email := c.PostForm("email")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
		})
	})

	// 3. FormFile("file") - 取表单里 name="file" 的上传文件
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请选择文件"})
			return
		}

		// SaveUploadedFile 将文件写入本地
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
	// c.MultipartForm() 解析 multipart/form-data，返回 *multipart.Form
	// form.File 存文件，form.Value 存普通表单字段
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"] // 对应 <input name="files" type="file" multiple>

		// form.File["files"] 是多个 name="files" 的上传文件
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

		// fmt.Sprintf 格式化字符串，%s 被 name 替换，返回 string 不打印
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

