package main

import "fmt"

func Plus(a int, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	fmt.Println("adf")
	postFile
}
