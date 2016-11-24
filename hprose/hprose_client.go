package main

import (
	"github.com/hprose/hprose-golang/rpc"
)

type HelloService struct {
	Hello func(string) string
}

func main() {
	c := rpc.NewHTTPClient("http://127.0.0.1:8080")
	var hello *HelloService
	c.UseService(&hello)
	result := hello.Hello("World")
	println(result)
}
