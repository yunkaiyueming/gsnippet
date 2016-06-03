package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) SetName(name string) {
	p.name = name
}

func (p *person) SetName2(name string) {
	p.name = name
}

func main() {
	//按照顺序提供初始化值
	p1 := person{"Bob", 20}

	//按照field.value初始化
	p2 := &person{name: "Ann", age: 40}

	//使用new方式创建
	p3 := new(person)

	p1.SetName("a") //不改变值
	fmt.Println(p1)
	p1.SetName2("b") //改变值
	fmt.Println(p1)

	p2.SetName("a") //不改变值
	fmt.Println(p2)
	p2.SetName2("b") //改变值
	fmt.Println(p2)

	p3.SetName("a") //不改变值
	fmt.Println(p3)
	p3.SetName2("b") //改变值
	fmt.Println(p3)
}
