package main

import (
	"container/ring"
	"fmt"
)

/*
func New(n int) *Ring //创建一个长度为n的环形链表
func (r *Ring) Do(f func(interface{}))  //对链表中任意元素执行f操作，如果f改变了r，则该操作造成的后果是不可预期的。
func (r *Ring) Len() int  //求环长度，返回环中元素数量
func (r *Ring) Link(s *Ring) *Ring  //Link连接r和s，并返回r原本的后继元素r.Next()。r不能为空。

如果r和s指向同一个环形链表，则会删除掉r和s之间的元素，删掉的元素构成一个子链表，返回指向该子链表的指针（r的原后继元素）；如果没有删除元素，则仍然返回r的原后继元素，而不是nil。如果r和s指向不同的链表，将创建一个单独的链表，将s指向的链表插入r后面，返回s原最后一个元素后面的元素（即r的原后继元素）。

func (r *Ring) Unlink(n int) *Ring //删除链表中n % r.Len()个元素，从r.Next()开始删除。如果n % r.Len() == 0，不修改r。返回删除的元素构成的链表，r不能为空。
func (r *Ring) Move(n int) *Ring  //返回移动n个位置（n>=0向前移动，n<0向后移动）后的元素，r不能为空。
func (r *Ring) Next() *Ring  //获取当前元素的下个元素
func (r *Ring) Prev() *Ring //获取当前元素的上个元素
*/

//ring实现了环形链表的操作。
func Test_Ring() {
	r := ring.New(10)
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 1; i <= r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

	r = r.Move(6)                      //移动到6
	fmt.Println("move to 6:", r.Value) //6
	deletedLink := r.Unlink(19)        //移除19%10=9个元素,返回删除以后的数据

	fmt.Println("======delete======")
	for i := 0; i < deletedLink.Len(); i++ {
		fmt.Println(deletedLink.Value)
		deletedLink = deletedLink.Next()
	}

	fmt.Println("=====last======")
	fmt.Println(r.Len())           //10-9=1
	fmt.Println(deletedLink.Len()) //9
	fmt.Println(r)
}
