package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	go listen()
	robotgo.KeyTap("s")

	ch := make(chan int)
	<-ch
}

func listen() {
	ok := robotgo.AddEvent("s")
	if ok == 0 {
		fmt.Println("press s")
	}
}
