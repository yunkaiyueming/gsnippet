package main

import (
	"fmt"
)

func main() {
	more_string := []string{"a", "b", "c"}
	SetArgs("hello", "world", "ok")
	SetArgs("hello", "world", more_string...)
	SetArgs("a", "b", "c", "d")
}

//变参函数
func SetArgs(a, b string, c ...string) {
	fmt.Println(a)
	fmt.Println(b)
	for i, v := range c {
		fmt.Println(i, v)
	}
}

//不能确定类型的变长参数
func SetMoreArgs(a int, b string, values ...interface{}) {
	for _, value := range values {
      	switch v := value.(type) {
           case int: …
           case float32: …
           case string: …
           case bool: …
           default: …
       	}
   }
}
