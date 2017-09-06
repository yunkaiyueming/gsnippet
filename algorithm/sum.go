package main

import (
	"fmt"
)

func main() {
	twoSum([]int{1, 2, 4, 7, 11, 15}, 5)
	twoSum([]int{1, 2, 4, 7, 11, 13, 13, 15}, 15)
}

//desc:寻找和为定值的两个数,没排序，先进行快排
//input:data为从小到大的切片
//out:和为sum的2个数
//算法：2端逼近算法
func twoSum(data []int, sum int) {
	start := 0
	end := len(data) - 1

	for start < end {
		tmp := data[start] + data[end]
		if tmp < sum {
			start++
		} else if tmp > sum {
			end--
		} else {
			fmt.Println(data[start], data[end])
			start++
		}
	}
}
