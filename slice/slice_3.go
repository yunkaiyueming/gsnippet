package main

import (
	"fmt"
)

func main() {
	array := [5]int{0, 1, 2, 3, 4}
	slice := array[:]
	fmt.Println(slice)
	fmt.Printf("slice in main1 %p \n", &slice)
	appendTest(slice)
	fmt.Printf("slice in main2 %p \n", &slice)
	fmt.Println(slice)
	fmt.Println("=============")
	fmt.Println(slice)
	fmt.Printf("slice in main1 %p \n", &slice)
	appendTestRef(&slice)
	fmt.Printf("slice in main2 %p \n", &slice)
	fmt.Println(slice)
}
func appendTest(slice []int) {
	fmt.Printf("slice appendTest1 %p \n", slice)
	slice = append(slice, 5)
	fmt.Println(slice)
	fmt.Printf("slice appendTest2 %p \n", slice)
}
func appendTestRef(slice *[]int) {
	fmt.Printf("slice appendTestRef1 %p \n", slice)
	*slice = append(*slice, 5)
	fmt.Println(slice)
	fmt.Printf("slice appendTestRef2 %p \n", slice)
}

