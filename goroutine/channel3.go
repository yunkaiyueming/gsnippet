package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan int)
	go Create(s)
	Consumer(s)
	time.Sleep(time.Second * 1)
}

func Create(s chan int) {
	for i := 0; ; i++ {
		s <- i
	}
}

func Consumer(s chan int) {
	for i := 0; ; i++ {
		fmt.Println(<-s)
	}
}
