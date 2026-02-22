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
// 痛点5: 统一响应格式要自己封装，json.Marshal 后再 w.Write

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/resp", respHandler) // 演示 Response 封装
	http.ListenAndServe(":8080", nil)
}

// 统一响应结构
//
// any = interface{} 的别名(Go1.18+)
// interface{} 表示「空接口」：没有任何方法，所以任意类型都实现了它，可以接收任何值
// 例如: Data 可以是 int、string、map、[]User、User 等，运行时再断言取出具体类型
//
// `json:"code"` 是 struct tag，序列化时字段名用 "code"
// 这样 JSON 输出为 {"code":0,"msg":"成功",...}，符合小写命名习惯
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 演示 json.Marshal + w.Write 返回统一格式
// json.Marshal(v) 把 Go 结构体/map 等序列化成 JSON 字节 []byte，返回 ([]byte, error)
func respHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(Response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]string{"name": "test"},
	})
	w.Write(data)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// 痛点: 需要手动设置 Content-Type
	w.Header().Set("Content-Type", "application/json")

	// 痛点: 需要根据 Method 手动分发
	switch r.Method {
	case "GET":
		// 痛点: 手动构造 JSON。两种方式: Encode 直接写 w，或 Marshal 得 []byte 再 w.Write
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

