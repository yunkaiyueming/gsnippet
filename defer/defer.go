package main

import (
	"fmt"
)

func main() {
	//defer是栈，先进后出，只负责一个最新的
	defer func1().func2()
	defer func3().func4()
	//output: 1,3,4,2
}

type oneDo struct{}
type twoDo struct{}

func func1() *oneDo {
	fmt.Println("func1 call")
	return &oneDo{}
}

func (this *oneDo) func2() {
	fmt.Println("func2 call")
}

func func3() *twoDo {
	fmt.Println("func3 call")
	return &twoDo{}
}

func (t *twoDo) func4() {
	fmt.Println("func4 call")
}

func 我们的() {
	fmt.Println("我们")
}
