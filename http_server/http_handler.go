package main

import (
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
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
