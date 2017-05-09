package main

import (
	"fmt"
)

//自定义错误
type MyError struct {
	ErrorNum int
	ErrorMsg string
}

func New(code int, msg string) *MyError {
	return &MyError{code, msg}
}

func (this *MyError) Error() { //实现了error接口
	fmt.Println(this.ErrorNum, this.ErrorMsg)
}

func main() {
	var f interface{}
	f = "10"
	if _, ok := f.(int); !ok {
		New(500, "类型断言int失败").Error()
	} else {
		New(200, "类型断言int成功").Error()
	}

	if _, ok := f.(string); !ok {
		New(500, "类型断言string失败").Error()
	} else {
		New(200, "类型断言string成功").Error()
	}
}
