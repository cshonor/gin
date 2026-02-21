// 对应视频: 3.原生http库的一些痛点
// 学习目标: 理解为什么需要 Gin 等 Web 框架
package main

import (
	"encoding/json"
	"net/http"
)

// 痛点1: 路由功能简单，不支持路径参数、路由组等
// 痛点2: 没有内置的 JSON 便捷方法，需要手动序列化
// 痛点3: 没有中间件机制
// 痛点4: 请求解析繁琐

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// 痛点: 需要手动设置 Content-Type
	w.Header().Set("Content-Type", "application/json")

	// 痛点: 需要根据 Method 手动分发
	switch r.Method {
	case "GET":
		// 痛点: 手动构造 JSON 响应
		user := User{ID: 1, Name: "张三"}
		json.NewEncoder(w).Encode(user)
	case "POST":
		// 痛点: 需要手动解析 JSON 请求体
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

