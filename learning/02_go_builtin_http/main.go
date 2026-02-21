// 对应视频: 2.go内置http库
// 学习目标: 了解 Go 标准库 net/http 的基本使用
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

