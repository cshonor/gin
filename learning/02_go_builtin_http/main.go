// 对应视频: 2.go内置http库
// 学习目标: 了解 Go 标准库 net/http 的基本使用
//
// --- fmt.Fprintf 和 fmt.Printf 的区别 ---
// Printf: 固定输出到控制台
// Fprintf(w, ...): 输出到指定的 io.Writer，如 w(HTTP响应体)、文件等
// 这里必须用 Fprintf(w, ...)，否则内容只会打到服务器控制台，客户端收不到
//
// 实际开发中: HTTP 响应多用 json.Encode / c.JSON 返回 JSON；Fprintf 常用于写文件、缓冲区
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
	"io"
	"net/http"
)

func main() {
	// 注册路由: 路径 -> 处理函数
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)   // 演示读取请求体
	http.HandleFunc("/info", infoHandler)   // 演示请求头、w.Write

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

// 处理 /echo：演示读取请求体（POST/PUT 等）
// io.ReadAll(r.Body) 读取整个 body，读完后 Body 会被消费，无法再次读取
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// 打印到控制台: 请求方法(GET/POST等) + 完整URL路径
	fmt.Println(r.Method, r.URL.String())
	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body)
		fmt.Println(string(byteData))
		fmt.Fprintf(w, "收到: %s\n", string(byteData))
	} else {
		fmt.Fprintf(w, "请用 POST 发送数据\n")
	}
}

// 处理 /info：演示 r.Header、w.Write
// r.Header 是 map[string][]string，可获取 Content-Type、User-Agent 等
// w.Write([]byte) 是底层写入方法，Fprintf 内部也是调它；Write([]byte("")) 表示空响应体
func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())
	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body)
		fmt.Println(string(byteData))
	}
	fmt.Println(r.Header)          // 打印请求头到控制台
	w.Write([]byte("ok"))          // 直接写入响应体，等价于 fmt.Fprintf(w, "ok")
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

