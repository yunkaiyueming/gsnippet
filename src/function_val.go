package main

import (
	"fmt"
	"math"
)

func main() {
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
