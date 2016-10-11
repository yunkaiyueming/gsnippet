package main

import (
	"fmt"
)

func main() {
	//将函数赋值给变量,并对其调用
	a := func(a, b int) int {
		return a + b
	}

	sum1 := a(1, 2)
	fmt.Println(sum1)

	sum2 := a(2, 3)
	println(sum2)

	//直接调用匿名函数
	flag := make(chan int)
	go func(a, b int) {
		fmt.Println("go routine")
		fmt.Println(a * b)
		flag <- a * b
	}(5, 2)

	fmt.Println("run end")
	<-flag
}
