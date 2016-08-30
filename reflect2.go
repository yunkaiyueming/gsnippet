// reflect.go
package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, S2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.S2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	//secret := NotknownType{"Ada", "Go", "Oberon"}
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("value %d: %v\n", i, value.Field(i))
	}

	for i := 0; i < typ.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, typ.Field(i).Name)
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

	for i := 0; i < typ.NumMethod(); i++ {
		res := value.Method(i).Call(nil)
		fmt.Printf("method %d:%s\n", i, res)
	}
}
