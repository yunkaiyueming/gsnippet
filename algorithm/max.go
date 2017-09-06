package main

import (
	"fmt"
)

func main2() {
	maxCommonDivisor(144, 63)
}

//求最大值
func maxNum() {
	var max int
	data := []int{87, 1, 2, 22, 99, 66, 5, 11, 6, 7, 8, 9}
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}

//求a,b的最大公约数
func maxCommonDivisor(a, b int) {
	if a < b {
		a, b = b, a
	}

	if a%b != 0 {
		maxCommonDivisor(b, a%b)
	} else {
		fmt.Println("最大公约数:", b)
	}
}
