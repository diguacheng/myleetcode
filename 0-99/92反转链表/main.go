package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	i := 1
	curr := head
	var start, end *ListNode
	flag := false
	for i < m {
		start = curr
		curr = curr.Next
		i++
	}
	var last, next *ListNode
	for i <= n {
		if !flag {
			end = curr
			flag = true
		}
		next = curr.Next
		curr.Next = last
		last = curr
		curr = next
		i++
	}
	end.Next = curr
	if start == nil {
		return last
	}
	start.Next = last
	return head

}

func newL(arr []int) *ListNode {
	head := new(ListNode)
	curr := head
	for i := 1; i <= len(arr); i++ {
		tmp := new(ListNode)
		tmp.Val = i
		curr.Next = tmp
		curr = tmp
	}
	return head.Next
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := newL(arr)
	fmt.Println(reverseBetween(head, 2, 4))

}
