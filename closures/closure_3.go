package main

import (
	"fmt"
	"time"
)

//因为这个循环运行的比goroutine还快,goroutine可能还没有开始,可能运行了几个，for循环已经结束了,此时i的大部分值就是100了,所以闭包内的i自然就是大多是100.
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i) //并不是按照i的值从1-》100
		}()
	}

	time.Sleep(1e9)
}
