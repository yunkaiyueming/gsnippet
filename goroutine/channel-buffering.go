package main

import "fmt"
import "time"

func worker1(done chan bool) {
	fmt.Print("working1...")
	done <- true
	fmt.Println("done1")
}

func worker2(done chan bool) {
	fmt.Print("working2...")
	done <- true
	fmt.Println("done2")
}

func main() {
	testNoBuffer()
	testBuffer()
	fmt.Println("main run end")
}

func testNoBuffer() {
	done := make(chan bool)
	go worker1(done)
	time.Sleep(time.Second * 2)
	fmt.Println("nobuffer run end")
}

func testBuffer() {
	done := make(chan bool, 1) //可以存储值，可以不用接收
	go worker2(done)
	time.Sleep(time.Second * 2)
	fmt.Println("buffer run end")
}
