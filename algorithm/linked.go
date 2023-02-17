package main// 链表数据结构
type List struct {
    Head *Node
    Tail *Node
}

// 节点数据结构
type Node struct {
    Value int
    Prev *Node
    Next *Node
}

// 插入节点
func (l *List) Push(value int) {  
    // 创建新节点
    n := &Node{Value: value}

    // 如果当前节点为空，则将它作为头结点
    if l.Head == nil {
        l.Head = n
    // 否则添加新节点到尾部
    } else {
        l.Tail.Next = n
        n.Prev = l.Tail
    }

    // 将尾节点更新为n
    l.Tail = n
}

// 删除节点
func (l *List) Remove(n *Node) {
    // 如果当前节点不为空
    if n != nil {
        // 如果不是头节点，则删除当前节点前面的节点指向当前节点的下一个节点
        if n.Prev != nil {
            n.Prev.Next = n.Next
        }  
        // 如果不是尾节点，则删除当前节点的指向当前节点前面节点的下一节点
        if n.Next != nil {
            n.Next.Prev = n.Prev
        }
    }
}