package main

import (
	"fmt"
)

var num = make(chan int)

func main() {
	//test1()
	test2()
	for i := 0; i < 10; i++ {
		fmt.Println(<-num)
	}
	fmt.Println("end")
}

func test1() {
	for i := 0; i < 10; i++ {
		go func() {
			num <- i // //这里的i很有可能全部是10, 因为每个并发执行到这一步的时候，i循环可能已经循环完了。
		}()
	}
}

func test2() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			num <- i //这里的i是具体的穿过来的，跟闭包是一个概念，变量和他的环境是一体的。
		}(i)
	}
}
