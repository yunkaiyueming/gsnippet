package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for _, element := range list {
		TestType(element)
		TestType2(element)
	}
}

func TestType(element interface{}) {
	if value, ok := element.(int); ok {
		fmt.Printf("%d\n", value)
	} else if value, ok := element.(string); ok {
		fmt.Printf("%s\n", value)
	} else if value, ok := element.(Person); ok {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Printf("a different type\n")
	}
}

func TestType2(element interface{}) {
	switch value := element.(type) {
	case int:
		fmt.Printf("%d\n", value)
	case string:
		fmt.Printf("%s\n", value)
	case Person:
		fmt.Printf("%s\n", value)
	default:
		fmt.Printf("a different type\n")
	}
}
