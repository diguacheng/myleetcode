package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 反转
	c := slow.Next   // curr
	// l := &ListNode{} // last
	// n := &ListNode{} // next
	// 注意，这里不要个l和n复制空节点，只赋值空指针即可
	var l,n *ListNode
	for c != nil {
		n = c.Next
		c.Next = l
		l = c
		c = n
	}
	p := head
	for p != nil && l != nil {
		if p.Val != l.Val {
			fmt.Println(p.Val,l.Val)
			return false
		}
		p = p.Next
		l = l.Next
	}
	return true

}
func main() {
	in := &ListNode{Val: 1}
	p := in
	for i := 1; i <= 1; i++ {
		node := &ListNode{Val: i}
		p.Next = node
		p = p.Next
	}
	fmt.Println(isPalindrome(in))
}
