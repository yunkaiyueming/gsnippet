package main

import (
	"fmt"
)

func main() {
	test2()
}

func test1() {
	c, quit := make(chan int), make(chan int)

	go func() {
		quit <- 1 // 发送完成信号
	}()

	go func() {
		c <- 2 // 添加数据
	}()

	for is_quit := false; !is_quit; {
		select { // 监视信道c的数据流出
		case v := <-c:
			fmt.Printf("received %d from c", v)
		case <-quit:
			is_quit = true // quit信道有输出，关闭for循环
		}
	}
}

func test2() {
	c, quit := make(chan int), make(chan int)

	go func() {
		quit <- 1 // 发送完成信号
	}()

	go func() {
		c <- 2 // 添加数据
	}()

loop:
	for {
		select {
		case v := <-c:
			fmt.Printf("received %d from c", v)
			fmt.Println("")
		case <-quit:
			fmt.Printf("quit")
			fmt.Println("")
			break loop
		}
	}
}
