package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	add5 := add(5)
	fmt.Println(add5(10))
}

func add(base int) func(i int) int {
	return func(i int) int {
		return i + base
	}
}
