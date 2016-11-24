package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func world(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
}

func main() {
	// 通过 HandlerFunc 把函数转换成 Handler 接口的实现对象
	wdHandler := http.HandlerFunc(world)
	http.Handle("/w", wdHandler)

	//HandleFunc 只是一个适配器
	http.HandleFunc("/h", hello)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
