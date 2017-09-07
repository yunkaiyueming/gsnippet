package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ts := r.FormValue("ts")
	tsInt, _ := strconv.Atoi(ts)

	notify := w.(http.CloseNotifier).CloseNotify() //捕获客户端abort的通知

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tsInt)*time.Second)
	defer cancel()

	dataCh := make(chan int)
	go func(ctx context.Context, dataCh chan int) {
		//do something
		time.Sleep(3 * time.Second)
		dataCh <- 1
	}(ctx, dataCh)

	select {
	case <-dataCh:
		w.Write([]byte("hello!"))
	case <-ctx.Done():
		w.Write([]byte(ctx.Err().Error()))
	case <-notify:
		fmt.Println("client abort")
		w.Write([]byte("client abort"))
	}
}

type worldHandler struct{}

func (wd *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("world"))
}

func main() {
	http.Handle("/h", &helloHandler{})
	http.Handle("/w", &worldHandler{})
	http.ListenAndServe(":8080", nil)
}
