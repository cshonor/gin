// 对应视频: 2.go内置http库
// 学习目标: 了解 Go 标准库 net/http 的基本使用
//
// --- fmt.Fprintf 和 fmt.Printf 的区别 ---
// Printf: 固定输出到控制台
// Fprintf(w, ...): 输出到指定的 io.Writer，如 w(HTTP响应体)、文件等
// 这里必须用 Fprintf(w, ...)，否则内容只会打到服务器控制台，客户端收不到
//
// --- if err := xxx; err != nil 是什么写法？---
// Go 的 if 前置语句：if 前置语句; 条件 { }
// 1. 前置语句: err := ListenAndServe(...)  先执行，ListenAndServe 阻塞，只有出错才返回
// 2. 分号 ; 分隔前置语句和条件
// 3. 条件: err != nil  有错误则执行花括号
// 等价于: err := xxx; if err != nil { }
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册路由: 路径 -> 处理函数
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// 启动 HTTP 服务器，监听 8080 端口
	fmt.Println("服务器启动在 http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("启动失败: %v\n", err)
	}
}

// 处理根路径请求
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎！你访问的是: %s\n", r.URL.Path)
	fmt.Fprintf(w, "请求方法: %s\n", r.Method)
}

// 处理 /hello 路径请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 查询参数获取 name，默认为 "World"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

