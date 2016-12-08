package main

import "fmt"

func modify(array []int) {
	for i, v := range array {
		array[i] = v + 1
	}
	fmt.Println("In modify(), array values:", array)
}

func Nomodify(array []int) {
	t := []int{1, 2, 3, 4, 5}
	for _, v := range t {
		array = append(array, v)
	}

	fmt.Println("In Nomodify(), array values:", array, "但不改变main里的array")
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	modify(array)
	Nomodify(array)
	fmt.Println("In main(), array values:", array)
}
