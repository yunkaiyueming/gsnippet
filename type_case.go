package main

import f "fmt"

func main() {
	var x interface{}
	test_type(x)

	i := 1
	test_type(i)

	i2r := string(i) //显式类型转换
	f.Println(i2r)
	f.Printf("%T", i)

	test_type2(i)

	classifier("xhc", "abd")
	classifier([2]string{"xhc", "abd"})
}

func test_type(x interface{}) { //类型断言
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

func test_type2(x interface{}) { //类型断言
	if _, ok := x.(int); ok {
		f.Println("x is int")
	}

	if _, ok := x.(string); ok {
		f.Println("x is int")
	}

	if _, ok := x.(float64); ok {
		f.Println("x is int")
	}

	if _, ok := x.(bool); ok {
		f.Println("x is int")
	}
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			f.Printf("Param #%d is a bool\n", i)
		case float64:
			f.Printf("Param #%d is a float64\n", i)
		case int, int64:
			f.Printf("Param #%d is a int\n", i)
		case nil:
			f.Printf("Param #%d is a nil\n", i)
		case string:
			f.Printf("Param #%d is a string\n", i)
		case []string:
			f.Printf("Param #%d is a []string\n", i)
		default:
			f.Printf("Param #%d is unknown\n", i)
		}
	}
}
