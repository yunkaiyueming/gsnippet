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
	test()
}

func test() {
	data := Person{}
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i)                                              //取得结构体的第i个字段
		fmt.Printf("%s %s %s \n", f.Name, f.Tag.Get("json"), f.Type) //第i个字段的名称,类型,值
	}
}

func test1() {
	fmt.Println(reflect.TypeOf(Person{}).Field(0).Tag.Get("html"))
	fmt.Println(reflect.TypeOf(Person{}).Field(1).Tag.Get("html"))
	fmt.Println(reflect.TypeOf(Person{}).Field(2).Tag.Get("html"))
	fmt.Println("=======================================")

	fmt.Println(reflect.TypeOf(Person{}).Field(0).Tag.Get("json"))
	fmt.Println(reflect.TypeOf(Person{}).Field(1).Tag.Get("json"))
	fmt.Println(reflect.TypeOf(Person{}).Field(2).Tag.Get("json"))
}
