package main

import (
	"fmt"
)

func main() {
	a, b := 4, 5
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)
}
