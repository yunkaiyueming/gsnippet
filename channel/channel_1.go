package main

import (
	"fmt"
	"time"
)

//信道的应用场景
func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func main() {

	generator := xrange()
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ { // 我们生成10个自增的整数！只有当<-generator取数据时候，才ch<-
		fmt.Println(<-generator)
	}
}
