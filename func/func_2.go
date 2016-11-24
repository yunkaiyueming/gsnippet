package main

import "fmt"

type A func(int, int)

func (f A) Serve() {
	fmt.Println("serve2")
}

func serve(int, int) {
	fmt.Println("serve1")
}

func main() {
	a := A(serve)
	a(1, 2)   //这行输出的结果是serve1
	a.Serve() //这行输出的结果是serve2
}

/*
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

type serverHandler struct {
	srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}
