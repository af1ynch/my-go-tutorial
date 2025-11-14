package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func insertNode(n0 *ListNode, p *ListNode) {
	n1 := n0.Next
	p.Next = n1
	n0.Next = p
}

func main() {

	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(5)
	n3 := NewListNode(4)
	n4 := NewListNode(2)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	fmt.Printf("n0：%v, n0地址：%p\n", n0, &n0)

	// 插入节点
	p := NewListNode(6)
	insertNode(n0, p)

	fmt.Printf("n0: %v\n", n0)
}
