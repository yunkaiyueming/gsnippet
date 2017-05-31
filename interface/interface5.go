package main

import (
	"fmt"
)

type I interface { // 有一个方法的接口 I
	Get() Int
}

type Int int // Int 类型实现了 I 接口
func (i Int) Get() Int {
	return i
}

func main() {
	var myint Int = 5
	var inter I = myint           // 变量赋值给接口
	val, ok := inter.(Int)        //接口变量.(实现接口的类型)
	fmt.Printf("%v, %v", val, ok) // 输出为：5，true

	//出错信息
	//	var myint2 Int = 6
	//	val, ok = myint2.(i)          //出错，接口才能断言，myint2不是接口
	//	fmt.Printf("%v, %v", val, ok) // 输出为：5，true
}
