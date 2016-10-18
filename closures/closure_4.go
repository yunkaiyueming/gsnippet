package main

import "fmt"

var ok = 1

/*
当我们调用分别由不同的参数调用ExFunc函数得到的函数时（myFunc()，myAnotherFunc()），得到的结果是隔离的，
也就是说每次调用ExFunc函数后都将生成并保存一个新的局部变量sum。其实这里ExFunc函数返回的就是闭包。
*/

//返回的是闭包
func ExFunc(n int) func() {
	sum := n
	return func() { //把匿名函数作为值赋给变量a (Go 不允许函数嵌套。//然而你可以利用匿名函数实现函数嵌套)
		fmt.Println(sum + 1) //调用本函数外的变量
		ok = sum + 1
		//sum++
	} //这里没有()匿名函数不会马上执行
}

func ExFunc2(n int) func() {
	sum := n
	return func() { //把匿名函数作为值赋给变量a (Go 不允许函数嵌套。//然而你可以利用匿名函数实现函数嵌套)
		fmt.Println(sum + 1) //调用本函数外的变量
		ok = sum + 1
		sum++
	} //这里没有()匿名函数不会马上执行
}

//myfunc与myAnotherFunc之间没有关系
////1.闭包中对外部数据的修改，外部不可见
//2.外部数据的值被保存到新建的静态变量中
func main() {
	myFunc := ExFunc(10)
	myFunc() // 11
	myAnotherFunc := ExFunc(20)
	myAnotherFunc() //21
	myFunc()        //11
	myAnotherFunc() //21

	fmt.Println(ok) //21,全局变量的值改变了
	myFunc()        //11

	fmt.Println("====================")
	myFunc = ExFunc2(10)
	myFunc() // 11
	myAnotherFunc = ExFunc2(20)
	myAnotherFunc() //21
	myFunc()        //12
	myAnotherFunc() //22

	fmt.Println(ok) //22,全局变量的值改变了
	myFunc()        //13
}
