package main

import (
	"fmt"
)

func main() {
	test1()
}

func test1() {
	dataCh := make(chan int, 5)

	go func() {
		dataCh <- 1
		dataCh <- 2
		dataCh <- 3
		close(dataCh) //在发送端关闭
	}()

	for i := 0; i < 10; i++ { //在一个已关闭 channel 上执行接收操作(<-ch)总是能够立即返回，返回值是对应类型的零值
		fmt.Println(i, <-dataCh) //output:1,2,3,0,0
	}

	fmt.Println("end")
}

func lenChan() {
	dataCh := make(chan int, 5)
	fmt.Println(len(dataCh))

	dataCh <- 1
	dataCh <- 2
	dataCh <- 3

	close(dataCh)
	fmt.Println(len(dataCh))
	for len(dataCh) > 0 {
		fmt.Println(len(dataCh), "receive:", <-dataCh)
	}
}

func test2() {
	dataCh := make(chan int, 5)

	go func() {
		dataCh <- 1
		dataCh <- 2
		dataCh <- 3
		close(dataCh) //在发送端关闭
	}()

	for d := range dataCh { //关闭后，只接受到里面存的值
		fmt.Println(d)
	}
	fmt.Println("end")
}
