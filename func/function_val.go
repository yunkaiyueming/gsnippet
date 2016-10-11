package main

import (
	"fmt"
	"math"
)

func main() {
	sumFunc := Init_func()
	fmt.Println(sumFunc(5, 6))
}

func Init_func() func(x, y int) int {
	reduceFun := func(x, y int) int {
		fmt.Println("run hele") //赋值的时候不会调用
		return x + y
	}

	return reduceFun
}

func func_val() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	test_func_sql2 := my_sqrt2(9)
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
	fmt.Println(my_sqrt2(9))
	fmt.Println(test_func_sql2)
}

func my_sqrt2(x float64) float64 {
	return math.Sqrt(x)
}
