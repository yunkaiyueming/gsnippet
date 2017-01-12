package main

import (
	"fmt"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()
	for i := 0; i < 2; i++ {
		<-quit
	}
}
