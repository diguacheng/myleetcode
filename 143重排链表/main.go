package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 空间复杂度为o(n)
func reorderList(head *ListNode) {
	//
	if head==nil{
		return 
	}

	list := make([]*ListNode, 0)
	p := head
	for p != nil {
		list = append(list, p)
		p = p.Next
	}
	n := (len(list) + 1) / 2
	p = head
	for i := len(list) - 1; i >=n; i-- {
		node := list[i]
		node.Next = p.Next
		p.Next = node
		p = node.Next
	}
	p.Next = nil

	return

}


//  找到链表中间，将后半段转置 然后拼接 空间复杂度为o(1)
func reorderList1(head *ListNode) {
	time.Sleep(1*time.Millisecond)
	if head==nil{
		return 
	}
	// 找到中点
	slow,fast :=head,head
	for fast.Next!=nil&&fast.Next.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	// 反转
	c:=slow.Next // curr
	slow.Next=nil
	// l:=&ListNode{}// last
	// n:=&ListNode{}// next
	l:=reserve(c)
	// 合并
	p:=head
	for l!=nil&&p.Next!=nil{
		node:=l // node
		l=l.Next
		node.Next=p.Next
		p.Next=node
		p=node.Next
	}
	if p.Next==nil&&l!=nil{
		p.Next=l

	}
	return

}

func reserve(c *ListNode)*ListNode{
	var l,n *ListNode
	for c!=nil{
		n=c.Next
		c.Next=l
		l=c
		c=n
	}
	return l

}


func main() {
	in := &ListNode{Val:1}
	p := in
	for i := 2; i <= 4; i++ {
		node := &ListNode{Val:i}
		p.Next = node
		p = p.Next
	}
	t:=time.Now()
	reorderList1(in)
	fmt.Println(time.Since(t))
	//r(in)
	for i := 1; i <= 5; i++ {
		fmt.Print(in.Val)
		in=in.Next
	}

}
