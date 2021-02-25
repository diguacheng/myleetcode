package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var a, cursor *ListNode
	if l1.Val <= l2.Val {
		cursor = l1
		a = l1
		l1 = l1.Next
	} else {
		cursor = l2
		a = l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cursor.Next = l1
			cursor = cursor.Next
			l1 = l1.Next
		} else {
			cursor.Next = l2
			cursor = cursor.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cursor.Next = l1
	} else {
		cursor.Next = l2
	}

	return a

}

func main() {

}
