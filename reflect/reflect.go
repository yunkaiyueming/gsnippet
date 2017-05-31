package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string `json:"the student name"`
	age  int    `json:"the student age"`
}

func main() {
	//	s := student{"twd", 12}
	//	PrintTypeVal(s)
	//	GetAllField(s)

	//	i := 10
	//	PrintTypeVal(i)

	//	a := "xhcabc"
	//	PrintTypeVal(a)

	Test()
}

func PrintTypeVal(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Printf("%T\n", i)
}

func GetAllField(s student) {
	fmt.Println("------field-------")
	for i := 0; i < 2; i++ {
		GetFieldByReflect(s, i)
	}
}

func GetFieldByReflect(s student, i int) {
	ts := reflect.TypeOf(s)
	f := ts.Field(i)
	fmt.Println(f.Name, f.Type, f.Tag, f.PkgPath)
}

func Test() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("type:", reflect.ValueOf(x).Type())
	fmt.Println("kind:", reflect.ValueOf(x).Kind())
	fmt.Println("value:", reflect.ValueOf(x).Float())
	fmt.Println(reflect.ValueOf(x).Interface())
	fmt.Printf("value is %5.2e\n", reflect.ValueOf(x).Interface())
	fmt.Println(reflect.ValueOf(x).Interface().(float64))
}
