package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2 //若开启会阻塞主线程main的执行
	go f1(out)

	time.Sleep(time.Second * 2)
}
