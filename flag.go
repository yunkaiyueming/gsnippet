package main

import (
	"flag"
	"fmt"
)

func main() {
	id := flag.Int("id", 1, "id for user_id")
	name := flag.String("name", "aaa", "name for user_name")
	a := flag.Int("a", 20, "a value")
	b := flag.Int("b", 40, "b value")

	var is_ok bool
	flag.BoolVar(&is_ok, "is_ok", true, "is_ok for user is_ok")

	flag.Parse() //处理接受数据前，要先进行parse解析

	c := *a + *b
	fmt.Println(*id)
	fmt.Println(*name)
	fmt.Println(is_ok)
	fmt.Println(c)
}
