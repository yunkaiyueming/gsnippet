package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	time.Sleep(time.Second)
	// 通知任务已完成
	done <- true
}
func main() {
	done := make(chan bool)
	go worker(done)
	// 等待任务完成
	<-done
	fmt.Println("finish")
}
