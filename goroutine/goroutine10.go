package main

import (
	"fmt"
	"time"
)

func main() {
	ret1 := make(chan []string, 3)
	ret2 := make(chan []string)

	go send1(ret1)
	go send2(ret2)

	rets := make([]string, 0)
	time.Sleep(2 * time.Second)
	for _, ret := range <-ret1 {
		fmt.Println("ret1===>", ret)
		rets = append(rets, ret)
	}

	for _, ret := range <-ret2 {
		fmt.Println("ret1===>", ret)
		rets = append(rets, ret)
	}
	fmt.Println(rets)
}

func send1(ret chan []string) {
	ret <- []string{"1", "2", "3", "4"}
	ret <- []string{"1", "2", "3", "4"}
	//close(ret)
}

func send2(ret chan []string) {
	time.Sleep(3 * time.Second)
	ret <- []string{"1", "2", "3"}
}
