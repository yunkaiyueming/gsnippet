package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go goWithC()
	go goWithSync()

	time.Sleep(time.Second * 15)
}

//使用channel阻塞goroutine
func goWithC() {
	t := make(chan int, 1)
	go func() {
		num := 1
		for {
			fmt.Println(num)
			time.Sleep(1e9)
			num++
			if num == 3 {
				break
			}
		}
		t <- 1
	}()

	select {
	case <-t:
		fmt.Println("end...")
	}
}

//使用sync阻塞goroutine
func goWithSync() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		num := 1
		for {
			fmt.Println(num)
			time.Sleep(1e9)
			num++
			if num == 3 {
				break
			}
		}
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("over")
}
