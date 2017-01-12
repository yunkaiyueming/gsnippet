package main

import (
	"fmt"
	"sync"
)

func main() {
	onceDoFunc := func() {
		fmt.Println("do once only")
	}

	OnceObj := sync.Once{}
	for i := 0; i < 10; i++ {
		OnceObj.Do(onceDoFunc) //只执行1次
		fmt.Println(i, " do")
	}
}

func TestOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
