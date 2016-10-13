package main

import (
	"fmt"
	"time"
)

var isDone = make(chan int, 5)

func main() {
	fmt.Println("main start")
	Start()

	select {
	case <-isDone:
		fmt.Println("receive down finsh")
	}

	fmt.Println("main finsh")
}

func Start() {
	fmt.Println("start")
	PareseUrl()
}

func PareseUrl() {
	fmt.Println("ParseUrl")
	go DownImg()
}

func DownImg() {
	fmt.Println("down start")
	time.Sleep(time.Duration(150))
	fmt.Println("down finish")
	isDone <- 1
}
