package main

import (
	"fmt"
)

type callBackFun func(t, w int) int

func main() {
	a := Strategy(5, 2, callBackAdd)
	fmt.Println(a)

	a = Strategy(5, 2, callBackX)
	fmt.Println(a)
}

//函数作为参数传入
func Strategy(a, b int, f callBackFun) int {
	return f(a, b)
}

func callBackAdd(a, b int) int {
	return a + b
}

func callBackX(a, b int) int {
	return a * b
}
