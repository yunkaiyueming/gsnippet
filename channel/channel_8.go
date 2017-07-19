package main

import (
	"fmt"
)

func main() {
	retCh := make(chan int, 0) //不管后面的长度为多少，取数据之前都需要往chan里面放数据
	for i := 0; i < 10; i++ {
		go send(i, retCh)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-retCh)
	}
	fmt.Println("run end")
}

func send(i int, retCh chan int) {
	retCh <- i
}
