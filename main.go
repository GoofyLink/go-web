package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 创建返回index数据的结构体
type indexData struct {
	Title string
	Age   int
}

func index(w http.ResponseWriter, r *http.Request) {
	//返回内容
	//设置请求头
	w.Header().Set("Content-Type", "application/json")
	//设置返回数据
	index := indexData{
		Title: "成为golang大神",
		Age:   20,
	}
	res, _ := json.Marshal(index)
	w.Write(res)

}

func main() {

	// 1.创建请求路径
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//3.创建处理函数
	http.HandleFunc("/index", index)

	// 2. 请求启动
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("http server failed!!!%v", err)
	}
}
