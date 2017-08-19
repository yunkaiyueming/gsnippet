package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := `fjdsiafdsa`
	t := 'w' //字节类型
	w := "fdsafsafdsa"
	fmt.Println(reflect.TypeOf(s), reflect.TypeOf(t))
	fmt.Println([]byte(s), byte(t), []byte(w))
}
