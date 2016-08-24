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
	for i := 0; i < 2; i++ {
		GetFieldByReflect(s, i)
	}

	GetAllField(s)
}

func Test2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func GetFieldByReflect(s student, i int) {
	ts := reflect.TypeOf(s)
	f := ts.Field(i)
	fmt.Println(f.Name, f.Type, f.Tag, f.PkgPath)
}

func GetAllField(s student) {
	for i := 0; i < 2; i++ {
		fieldInfo := reflect.TypeOf(s).Field(i)
		fmt.Println(fieldInfo.Name, fieldInfo.Type, fieldInfo.Tag, fieldInfo.PkgPath)
	}
}
