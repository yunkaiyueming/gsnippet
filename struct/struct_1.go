package main

import (
	"fmt"
)

// go的组合，可以使用内部字段的方法，类似面向对象的继承，但没有java那种多重继承，也同样达到代码的封装和复用特点
type A struct {
	Number int
	FieldB *B
}

// new 是一个用于分配内存的内置函数。它接收一个类型作为参数，并返回一个指向该类型的零值的指针。
// new 函数分配的内存空间是被初始化为零值的。如果需要对字段进行更复杂的初始化操作，可以使用字面值结构体或构造函数。
func NewA() *A {
	return new(A)
}

func AConstruct(num int) *A {
	return &A{
		Number: num,
		FieldB: new(B),
	}
}

func (a *A) SayHello() {
	fmt.Println("a.hello", a.Number)
}

type B struct {
	Str    string
	FiledA *A
}

func (b *B) SayHello() {
	fmt.Println("b.hello", b.FiledA.Number) //b中a元素的number值
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
	b.FiledA.SayHello()
	b.SayHello()

	a1 := NewA()
	fmt.Println(a1)
}
