package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData2(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) { //getData() 使用了无限循环：它随着 sendData() 的发送完成和 ch 变空也结束了
	var input string
	//time.Sleep(1e9) //若开启则main会直接执行完，getData还没有接收到数据
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func getData2(ch chan string) {
	var input string
	for {
		input = <-ch
		if input != "" {
			fmt.Printf("%s ", input)
		} else {
			break
		}
	}
}
