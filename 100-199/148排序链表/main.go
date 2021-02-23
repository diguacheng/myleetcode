package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 使用了递归 空间复杂度不是0（1）
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	res := &ListNode{}
	p := res
	left, right := sortList(head), sortList(mid)
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p=p.Next
	}
	if left != nil {
		p.Next = left
	}
	if right != nil {
		p.Next = right
	}
	return res.Next

}


func sortList1(head *ListNode) *ListNode {
	// 归并 排序 自底向上 
	h:=head
	length:=0
	intv:=1
	for h!=nil{
		h=h.Next
		length++
	}
	res:=&ListNode{}
	res.Next=head
	for intv<length {
		pre:=res
		h:=res.Next
		for h!=nil{
			h1:=h
			i:=intv
			for i!=0&&h!=nil{
				h=h.Next
				i--
			}
			if i>0{
				break
			}
			h2:=h
			i=intv
			for i!=0&&h!=nil{
				h=h.Next
				i--
			}
			c1,c2:=intv, intv-i
			for c1!=0&&c2!=0{
				if h1.Val<h2.Val{
					pre.Next=h1
					h1=h1.Next
					c1--
				}else{
					pre.Next=h2
					h2=h2.Next
					c2--
				}
				pre=pre.Next
			}
			if c1>0{
				pre.Next=h1
				pre=pre.Next
			}else{
				pre.Next=h2
				pre=pre.Next
			}
			pre.Next=h
		}
		intv*=2
	}
	return res.Next
}

func main() {
	in := &ListNode{Val: 0}
	p := in
	ll := []int{4,2, 1,3}
	for _, i := range ll {
		node := &ListNode{Val: i}
		p.Next = node
		p = p.Next
	}
	x := sortList1(in.Next)
	fmt.Printf("%#v",x)

}
