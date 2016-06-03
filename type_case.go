package main

import f "fmt"

func main() {
	var x interface{}
	test_type(x)

	i := 20
	test_type(i)
}

func test_type(x interface{}) {
	switch x.(type) {
	case nil:
		f.Printf("x is nil")
	case int:
		f.Printf("x is int")
	case float64:
		f.Printf("x is float64")
	case bool, string:
		f.Printf("x is bool")
	default:
		f.Printf("x is unknown")
	}
}
