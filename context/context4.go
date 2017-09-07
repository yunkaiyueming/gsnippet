package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	retCh := make(chan int, 1)

	pctx := context.Background()
	//ctx, cancel := context.WithCancel(pctx)
	ctx, cancel := context.WithTimeout(pctx, 3*time.Second)
	go doSomeThing(ctx, retCh)

	//	time.Sleep(3 * time.Second)
	//	fmt.Println("cancle send")
	//	cancel()

	time.Sleep(5 * time.Second)
	cancel()
}

func doSomeThing(ctx context.Context, retCh chan int) {
	count := 0
	tmpCh := make(chan int)
	for i := 0; i < 100; i++ {
		go func() {
			count++
			tmpCh <- count
		}()
	}

	total := 0
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case data := <-tmpCh:
			fmt.Println("recv:", data)
			total += data
		}
	}

	fmt.Println("send total", total)
	retCh <- total
}

func doSomeThing2(ctx context.Context) {
}
