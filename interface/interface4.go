package main

import (
	"fmt"
)

type Header interface {
	toString() string
}

type Wrapper interface {
	Header
	//toHome() string 没有这个函数，实现了Header接口，也就实现了Wrapper接口
}

type Apply struct {
	url       string
	host      string
	userAgent string
}

func (a *Apply) toString() string {
	return fmt.Sprintf("url:%s,host:%s,userAgent:%s \n", a.url, a.host, a.userAgent)
}

func CreateHeader() Header {
	return &Apply{"/get/Header", "www.baidu.com", "firefox-v1.1"}
}

func CreateWrapper() Wrapper {
	return &Apply{"/get/Wrapper", "www.baidu.com", "firefox-v1.2"}
}

func main() {
	println(CreateHeader().toString())
	println(CreateWrapper().toString())

}
