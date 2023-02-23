package main

import (
	"fmt"
)

func main(){
	main2()
}

func main1() {
	data := []string{"0", "1", "2", "3"}
	deleteIndex := 3

	//tmp := make([]string, 0)
	tmp := append(data[0:deleteIndex], data[deleteIndex+1:]...)
	fmt.Println(deleteIndex, tmp)

}

func main2() {
	array1 :=[4]int{10, 20, 30, 40}
	array := []int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v, type is %T, type is %T \n", array, array, array1)

	fmt.Printf("len(array1)=%v, cap(array1)=%v, len(array)=%v, cap(array)=%v ", len(array1), cap(array1),  len(array), cap(array))
}
