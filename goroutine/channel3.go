package main

import (
	"fmt"
	"time"
)

var flag = 1

func main() {
	s := make(chan int)
	go Create(s)
	Consumer(s)
	time.Sleep(time.Second * 1)
}

func Create(s chan int) {
	for i := 0; i < 10; i++ {
		s <- i
	}
	flag = 0
}

func Consumer(s chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-s)
	}

	fmt.Println(flag)
}
