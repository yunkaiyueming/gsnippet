package main

import (
	"container/list"
	"fmt"
)

//list包实现了双向链表
//对双向链表的操作
var l = list.New()

func main() {
	test_1()
}

func test_1() {
	e4 := l.PushBack(4)  //写入队尾
	e1 := l.PushFront(1) //写入队首
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	//1,2,3,4

	l.PushBack(5)
	l.PushBack(6)
	e7 := l.PushBack(7)
	e_1 := l.PushBack(-1) // 1->7->-1
	l.PushFront(-2)

	l.MoveAfter(e7, e4)   //e7移到e4的后面
	l.MoveBefore(e_1, e1) //-1移到1的前面
	GetAllList()

	l.MoveToBack(e4)
	l.Remove(e1)
	fmt.Println("===========")
	GetAllList()
}

func GetAllList() {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
