package main

import (
	"fmt"
)

type T struct{ name string }
type I interface{}

func main() {
	var t *T
	if t == nil {
		fmt.Println("true", t)
	} else {
		fmt.Println("false", t)
	}

	var t1 *T
	if t1 == nil {
		fmt.Println("true", t1)
	} else {
		fmt.Println("false", t1)
	}

	var m map[string]string
	if m == nil {
		fmt.Println("true", m)
	} else {
		fmt.Println("false", m)
	}

	var c chan int
	if c == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
