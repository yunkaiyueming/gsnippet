package main

import (
	"fmt"
)

type A struct {
	Number int
	FieldB *B
}

type B struct {
	Str    string
	FiledA *A
}

func main() {
	a := A{Number: 5}
	fmt.Println(a)

	b := B{Str: "hehe"}
	fmt.Println(b)

	a.FieldB = &b
	fmt.Println(a, &a)

	b.FiledA = &a
	fmt.Println(b)
}
