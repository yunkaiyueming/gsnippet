package main

import (
	"fmt"
)

func main() {
	dataCh := make(chan int, 3)
	done := make(chan int)
	//out := make(chan int, 6)

	dataCh <- 1
	dataCh <- 2
	dataCh <- 3
	close(dataCh)

	go func() {
		for {
			select {
			case i := <-dataCh:
				fmt.Println("receive:", i)
			case <-done:
				fmt.Println("done")
				break
			}
		}
	}()

	go sendDone(done)
	//	for n := range dataCh {
	//		select {
	//		case out <- n:
	//			fmt.Println(<-out)
	//		case <-done:
	//			fmt.Println("done over")
	//		}
	//	}
	for {
	}
	fmt.Println("end")
}

func sendDone(done chan int) {
	done <- 1
}
