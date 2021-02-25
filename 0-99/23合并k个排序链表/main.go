package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 二分法
	for len(lists) > 1 {
		temp := make([]*ListNode, 0)
		i := len(lists) - 1
		for ; i >= 1; i -= 2 {
			templ := merge2lists(lists[i], lists[i-1])
			temp = append(temp, templ)
		}
		if i == 0 {
			temp = append(temp, lists[0])
		}
		lists = temp
	}
	return lists[0]

}

func merge2lists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func main() {

}
