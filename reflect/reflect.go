package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string "the student name"
	age  int    "the student age"
}

func main() {
	s := student{"twd", 12}
	PrintTypeVal(s)
	//GetAllField(s)

	i := 10
	PrintTypeVal(i)

	a := "xhcabc"
	PrintTypeVal(a)

	Test()
}

func PrintTypeVal(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Printf("%T", i)
}

func GetAllField(s student) {
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
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
