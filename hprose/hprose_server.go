package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

func Hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("Hello", Hello)
	http.ListenAndServe("127.0.0.1:8080", service)
}
