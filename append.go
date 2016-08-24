package main

import "fmt"

func main() {
	test4()
}

func test() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)
	d := append(c, 7)
	fmt.Println(d)
}

func test2() {
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

func test3() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a[:3], a[4:], a[1:7]) //不包含第一个元素
	fmt.Println(a[len(a)/2:], a[:len(a)/2])
}

func test4() {
	a := make([]int, 10)
	fmt.Println(a)

	b := make(map[string]string)
	fmt.Println(b)
	c := map[string]string{"a": "b", "c": "d"}
	fmt.Println(c)

	d := make(chan int, 5)
	fmt.Println(d)
	var e chan int
	fmt.Println(e)
}
