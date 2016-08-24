package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := NewPerson("aa", 20)
	fmt.Println(p1)

	p2 := NewPerson("bb", 20)
	fmt.Println(p2)
}

func NewPerson(name string, age int) *person {
	return &person{name, age}
}

func NewPerson2(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p
}

func (p person) SetName(name string) {
	p.name = name
}

func (p *person) SetName2(name string) {
	p.name = name
}

func CreateStructType() {
	//按照顺序提供初始化值 值类型
	p1 := person{"Bob", 20}

	//按照field.value初始化 引用类型
	p2 := &person{name: "Ann", age: 40}

	//使用new方式创建  引用类型
	p3 := new(person)

	//使用var声明  值类型
	var p4 person

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

	p4.SetName("a") //不改变值
	fmt.Println(p4)
	p4.SetName2("b") //改变值
	fmt.Println(p4)
}
