package main

import (
	"net/http"
	"time"
)

var R Router

type Router struct {
	Maping map[string]func(w http.ResponseWriter, r *http.Request)
}

type customHandler struct {
}

func (cb *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("统一处理"))
	w.Write([]byte("customHandler!!"))
}

type LoginConhandler struct {
}

func main() {
	var server *http.Server = &http.Server{
		Addr:           ":8080",
		Handler:        &customHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
