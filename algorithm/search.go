package main

import (
	"fmt"
)

func main() {
	output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//fmt.Println(binarySearch(output, 0, len(output)-1, 6))

	fmt.Println(binarySearch2(output, 10))
}

//问题：给定一个有序的数组，查找某个数是否在数组中，请编程实现。
//二分查找
func binarySearch(data []int, start, end, findV int) (int, int) {
	if start == end-1 {
		if data[start] == findV {
			return start, findV
		} else if data[end] == findV {
			return end, findV
		} else {
			return 0, 0
		}
	}

	middle := (end + start) / 2
	if data[middle] == findV {
		return middle, data[middle]
	} else if data[middle] > findV {
		return binarySearch(data, start, middle, findV)
	} else {
		return binarySearch(data, middle, end, findV)
	}
}

func binarySearch2(data []int, findV int) (int, int) {
	first, last := 0, len(data)-1
	for first < last {
		if data[first] == findV {
			return first, findV
		} else if data[last] == findV {
			return last, findV
		}

		middle := (first + last) / 2
		if data[middle] < findV {
			first = middle
		} else if data[middle] == findV {
			return middle, findV
		} else {
			last = middle
		}
	}
	return 0, 0
}
