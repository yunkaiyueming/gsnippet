package main

import (
	"fmt"
)

func main() {
	output := insertSort([]int{5, 62, 7, 8, 95, 65, 3, 2, 44, 2, 76, 4, 5, 7, 8, 7})
	fmt.Println(output)
}

//快排
func quickSort(data []int) []int {
	if len(data) == 0 {
		return nil
	} else if len(data) == 1 {
		return []int{data[0]}
	} else if len(data) == 2 {
		if data[0] > data[1] {
			return []int{data[1], data[0]}
		}
		return []int{data[1], data[0]}
	}

	min, center, max := make([]int, 0), data[0], make([]int, 0)
	for i := 1; i < len(data); i++ {
		if data[i] < center {
			min = append(min, data[i])
		} else {
			max = append(max, data[i])
		}
	}

	maxSort := quickSort(max)
	minSort := quickSort(min)

	tmp := append(minSort, center)
	return append(tmp, maxSort...)
}

//冒泡排序
func bubbleSort(data []int) []int {
	for i := len(data) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

//选择排序
//算法原理：先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
func selectSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		swapIndex := i

		maxIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex > swapIndex {
			data[swapIndex], data[maxIndex] = data[maxIndex], data[swapIndex]
		}
	}
	return data
}

//插入排序：
//算法原理：将数据分为两部分，有序部分与无序部分，一开始有序部分包含第1个元素，依次将无序的元素插入到有序部分，直到所有元素有序。插入排序又分为直接插入排序、二分插入排序、链表插入等，这里只讨论直接插入排序。它是稳定的排序算法，时间复杂度为O(n^2)
func insertSort(data []int) []int {
	for i := 1; i < len(data); i++ {

		for j := i - 1; j > 0; j-- {
			if data[i] > data[j] {
				data[j] = data[j-1]
			} else {
				data[j] = data[i]
				break
			}
		}
	}
	return data
}
