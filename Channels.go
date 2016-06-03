package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		messages <- "ping"
	}()

	//阻塞
	msg := <-messages
	fmt.Println(msg)
	fmt.Println("fasfs")
}
