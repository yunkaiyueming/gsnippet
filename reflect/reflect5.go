package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" html:"uname"`
	Age  int    `json:"age" html:"uage"`
	Pwd  string `json:"pwd" html:"upwd"`
}

func main() {
	fmt.Println(reflect.TypeOf(Person{}).Field(0).Tag.Get("html"))
	fmt.Println(reflect.TypeOf(Person{}).Field(1).Tag.Get("html"))
	fmt.Println(reflect.TypeOf(Person{}).Field(2).Tag.Get("html"))
	fmt.Println("=======================================")

	fmt.Println(reflect.TypeOf(Person{}).Field(0).Tag.Get("json"))
	fmt.Println(reflect.TypeOf(Person{}).Field(1).Tag.Get("json"))
	fmt.Println(reflect.TypeOf(Person{}).Field(2).Tag.Get("json"))
}
