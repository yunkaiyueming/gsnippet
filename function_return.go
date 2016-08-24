package main

import (
	"fmt"
)

type strategyFunc func() int

func main() {
	fmt.Println(CreateFunc1(5)())

	fmt.Println(CreateFunc2(2, 5)())

	fmt.Println(CreateFunc3(2, 5)())
}

func CreateFunc1(a int) func() int {
	return func() int {
		return a
	}
}

func CreateFunc2(a, b int) strategyFunc {
	sum := a + b
	return func() int {
		return sum
	}
}

func CreateFunc3(a, b int) strategyFunc {
	x := a * b
	return func() int {
		return x
	}
}

//闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
func ApplyClosure() {
	var f = Adder()
	fmt.Print(f(1), " - ")  //1
	fmt.Print(f(20), " - ") // 21
	fmt.Print(f(300))       //321
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
