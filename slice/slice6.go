package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 10)
	fmt.Println(a) //output:0,0,0,0,0

	a = append(a, 123)
	fmt.Println(a)

	p := new([]int)
	fmt.Println(p)
	*p = append(*p, 123)
	fmt.Println(p)

	i := new(int)
	fmt.Println(*i)
}
