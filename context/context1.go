package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("end==>", ctx.Err())
					return // returning not to leak the goroutine
				case dst <- n:
					fmt.Println("send", n)
					n++
				}
			}
		}()

		return dst
	}

	//WithCancel 和 WithTimeout 函数 会返回继承的 Context 对象， 这些对象可以比它们的父 Context 更早地取消。
	ctx, cancel := context.WithCancel(context.Background()) //Background 是所有 Context 对象树的根，它不能被取消
	defer cancel()                                          // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
