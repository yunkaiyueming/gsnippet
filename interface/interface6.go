package main

import (
	"fmt"
)

type I interface {
	Get() int
	Put(v int)
}

// S实现了I接口
type S struct{ i int }

func (s *S) Get() int  { return s.i }
func (s *S) Put(v int) { s.i = v }

// 另一个实现了 I 接口的 R 类型
type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

//类型断言
func f(p interface{}) {
	switch t := p.(type) { // 判断传递给 p 的实际类型
	case *S: // 指向 S 的指针类型
		fmt.Println("*S", t.Get())

	case *R: // 指向 R 的指针类型
		fmt.Println("*R", t.Get())

	case S: // S 类型
		fmt.Println("S", t.Get())

	case R: // R 类型
		fmt.Println("R", t.Get())

	default: //实现了 I 接口的其他类型
		fmt.Println("not find")
	}
}

//make用于map ,slice,channel
func main() {
	//接口类型的断言使用
	s := S{}
	s.Put(1)

	s2 := new(S)
	s2.Put(2)

	r := R{}
	r.Put(3)

	r2 := new(R)
	r2.Put(4)

	f(s)
	f(s2)
	f(r)
	f(r2)
}
