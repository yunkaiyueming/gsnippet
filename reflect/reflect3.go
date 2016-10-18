package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	GetStructInfo()
	GetStructInfo2()
}

func GetStructInfo() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset")
	fmt.Println("t is now", t)
}

func GetStructInfo2() {
	t := T{23, "skidoo"}
	value := reflect.ValueOf(t)
	typ := reflect.TypeOf(t)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("%d:%v %v=%v\n", i, typ.Field(i).Type, typ.Field(i).Name, value.Field(i))
	}
}
