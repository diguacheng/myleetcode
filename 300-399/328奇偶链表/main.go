package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd :=  head.Next // 奇数
	var pre *ListNode
	pe, po := head, head.Next
	for po.Next != nil {
		pre=pe
		pe.Next = po.Next
		pe = pe.Next
		if pe.Next!= nil {
			po.Next = pe.Next
			po = po.Next
		} else {
			po.Next=nil
			break
		}
	}
	if pe==nil{
		pre.Next=odd
	}else{
		pe.Next=odd
	}
	return head 

}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := &ListNode{}
	curr := head
	for i := 0; i < 5; i++ {
		temp := &ListNode{Val: arr[i]}
		curr.Next = temp
		curr = curr.Next
	}
	head = oddEvenList(head.Next)
	for i := 0; i < 5; i++ {
		fmt.Println(head.Val)
		head=head.Next
	}

}
