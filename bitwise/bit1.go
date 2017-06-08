package main

import (
	"fmt"
)

var a, b = 42, 25 //这是有符号整数
var a1, b1 uint   //无符号整数
//int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。

// 0010 1010
// 0001 1001
func main() {
	bitAnd()        // &
	bitOr()         // |
	bitXor()        // ^
	bitComplement() //^

	bitLeft(3)  // <<
	bitRight(3) // >>
	swap(a, b)
}

func bitAnd() {
	//按位与:0000 1000
	c := a & b
	fmt.Println(c) //8
}

func bitOr() {
	//按位或 0011 1011
	c := a | b
	fmt.Println(c) //32+16+8+2+1=59
}

func bitXor() {
	//按位异或 0011 0011
	fmt.Println(a ^ b) //32+16+2+1=51
}

func bitComplement() {
	//-0000 0001 ^ 0010 1010
	//按位补码 0010 1011
	fmt.Println(^a) //-（32+8+2+1） = -43
}

func bitLeft(length uint) { //1 0101 0000
	fmt.Println(a << length) //256+64+16=336
}

func bitRight(length uint) { //0000 0011
	fmt.Println(b >> length) //3
}

func swap(a, b int) {
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
