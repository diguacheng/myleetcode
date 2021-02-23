package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var last *ListNode
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		last = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	if slow == fast {
		root.Left = nil
		if fast.Next != nil {
			root.Right = &TreeNode{Val: fast.Next.Val}
		} else {
			root.Right = nil
		}
		return root
	}
	last.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}

func createlist(arr []int) *ListNode {
	head := &ListNode{}
	curr := head
	for i := 0; i < len(arr); i++ {
		temp := &ListNode{Val: arr[i]}
		curr.Next = temp
		curr = curr.Next
	}
	return head.Next

}
func main() {
	arr := []int{-10, -3, 0, 5, 9}
	head := createlist(arr)
	s := sortedListToBST(head)
	fmt.Println(s)
}
