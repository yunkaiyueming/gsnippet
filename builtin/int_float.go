package main

import (
	"fmt"
)

func main() {
	test_slice()
}

func int_2_float() {
	var a int
	var b float64
	a = 1005
	b = 13.56
	fmt.Println(a, float64(a), int64(b))
}

func test_slice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//fmt.Println(a[0:2])
	//fmt.Println(a[0:len(a)])

	//	for i := 0; i <= len(a); i++ {
	//		fmt.Println(a[0:i], a[i:len(a)])
	//	}
	fmt.Println(len(a), cap(a))
	fmt.Println(a[0], a[7], a[7:7], a[7:8], a[8:8])
	//delete
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], a[0:i], a[i+1:len(a)])
	}
}
