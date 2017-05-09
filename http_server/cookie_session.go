package main

import (
	"fmt"
	"net/http"
	"time"
)

var cookies map[string]string

func init() {
	cookies = map[string]string{"go_version": "1.6", "web_server": "go", "age": "16"}
	//fmt.Println(cookies)
}

func main() {
	http.HandleFunc("/", webServer)   //设置访问的路由
	http.ListenAndServe(":8989", nil) //设置监听的端口
}

func webServer(w http.ResponseWriter, r *http.Request) {
	MySetCookie(w)
	MyGetCookie(r)
}

func MySetCookie(w http.ResponseWriter) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	for name, val := range cookies {
		cookie := http.Cookie{Name: name, Value: val, Expires: expiration}
		cookie := http.Cookie{Name: name, Value: val, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: 1}

		http.SetCookie(w, &cookie)
	}
}

func MyGetCookie(r *http.Request) {
	getCookies := r.Cookies()
	fmt.Println(getCookies)

	for index, cookie := range getCookies {
		fmt.Println(index, cookie.Name, cookie.Value, cookie.MaxAge)
	}

	//获取指定cookie
	//cookie, _ := r.Cookie("username")
	//fmt.Println(cookie)
}
