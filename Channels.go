package main

import (
	"fmt"
	"time"
)

func main() {
	blockChanTest()
	BufferChanTest()
	fmt.Println("main run end")
}

func blockChanTest() {
	messages := make(chan string) //阻塞性通道（非缓冲型通道），发送必须接受，接收必须发送
	go func() {
		fmt.Println("send msg to nobuffer")
		messages <- "要使用我必须发送,不可存储，必须要接收，否则阻塞"
		fmt.Println("no buffer not run here")
	}()

	time.Sleep(2 * time.Second)

	//阻塞
	//fmt.Println("receive msg:", <-messages)
	fmt.Println("block run end")
}

func BufferChanTest() {
	chanMsg := make(chan string, 2)
	go func() {
		fmt.Println("send msg to buffer")
		chanMsg <- "我发送了，可以不用接收，我缓存存在channel中"
		fmt.Println("buffer run here")
	}()

	time.Sleep(2 * time.Second)
	//非阻塞，可以不接收
	//fmt.Println(<-chanMsg)
	fmt.Println("buffer run end")
}
