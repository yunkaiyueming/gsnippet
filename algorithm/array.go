//输入一个整数数组，调整数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。要求时间复杂度为O(n)。

package main

import (
	"fmt"
)

func main() {
	output := ([]int{5, 62, 7, 8, 95, 65, 3, 2, 44, 2, 76, 4, 5, 7, 8, 7})
	fmt.Println(oddEvenSort2(output))
}

//开辟一个新切片，偶数从末尾到头开始放，奇数从头到末尾开始放
func oddEvenSort1(data []int) []int {
	newData := make([]int, len(data))
	num := len(data)
	last := num
	first := 0
	for i := 0; i < num; i++ {
		if data[i]%2 == 0 {
			last--
			newData[last] = data[i]
		} else {
			newData[first] = data[i]
			first++
		}
	}
	return newData
}

//用1头指针，1尾指针，头指针发现偶数&尾指针发现奇数，交换之。
func oddEvenSort2(data []int) []int {
	first, last := 0, len(data)-1
	for first < last {
		if data[first]%2 == 1 {
			first++
			continue
		}

		if data[last]%2 == 0 {
			last--
			continue
		}

		data[first], data[last] = data[last], data[first]
		first++
		last--
	}
	return data
}
