package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//知道接收消息多少次，直接接多少次
func receive_with_num() {
	t := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			t <- 1
		}
	}()
	receiver := 0
	for i := range t {
		fmt.Println("element is  ", i)
		receiver++
		if receiver == 5 {
			close(t)
		}
	}
}
func main() {
	var job sync.WaitGroup
	t := make(chan string)

	go func() {
		defer job.Done()
		job.Add(1)

		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			t <- "i:" + strconv.Itoa(i)
		}
	}()

	go func() {
		defer job.Done()
		job.Add(1)

		for j := 0; j < 6; j++ {
			time.Sleep(time.Second)
			t <- "j:" + strconv.Itoa(j)
		}
	}()

	go func() {
		time.Sleep(time.Second)
		job.Wait()
		fmt.Println("all job done")
		close(t)
	}()

	for i := range t {
		fmt.Println("element is ", i)
	}
	fmt.Println("main done")
}
