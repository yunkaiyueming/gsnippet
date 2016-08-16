package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(a[:cap(a)])

	s := a[1:3]
	fmt.Println(s)

	c := append(s, 6)
	fmt.Println(c)
	fmt.Println(c[:cap(c)])

	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("s ==", s[:cap(s)])
	fmt.Println("c ==", c[:cap(c)])

	d := append(a, 6)
	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("d ==", d[:cap(d)])
}
