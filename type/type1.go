package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1
	b := ""
	c := true
	d := []string{"a", "b"}
	var e interface{}
	f := make(map[string]interface{})

	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b).Kind())
	fmt.Println(reflect.TypeOf(c), reflect.ValueOf(c).Kind())
	fmt.Println(reflect.TypeOf(d), reflect.ValueOf(d).Kind())
	fmt.Println(reflect.TypeOf(e), reflect.ValueOf(e).Kind())
	fmt.Println(reflect.TypeOf(f), reflect.ValueOf(f).Kind())
}
