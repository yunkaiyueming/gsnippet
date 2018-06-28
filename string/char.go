package main

import (
	"fmt"
)

//golang的字符称为rune，等价于C中的char，可直接与整数转换
//rune实际是整型，必需先将其转换为string才能打印出来，否则打印出来的是一个整数
func main() {
	var c rune = 'a' //字符
	fmt.Println("'a' asci to", int(c))

	var i int = 98
	fmt.Println("98 convert to char ", string(rune(i)))

	//string to rune
	for _, char := range []rune("abcdefg") {
		fmt.Println(string(char), int(char))
	}

	fmt.Println(rune(200), string(rune(200)), int(rune(200)))
}
