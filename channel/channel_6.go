package main

import (
	"fmt"
)
s
func rand01() chan int {
	ch := make(chan int)

	go func() {
		for {
			select { //select会尝试执行各个case, 如果都可以执行，那么随机选一个执行
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}

func main() {
	generator := rand01() //初始化一个01随机生成器

	//测试，打印10个随机01
	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}
