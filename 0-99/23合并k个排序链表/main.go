package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists1(lists []*ListNode) *ListNode {
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type minHeap []*ListNode

func (m minHeap) Len() int { return len(m) }

func (m minHeap) Less(i, j int) bool { return m[i].Val < m[j].Val }

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	l := minHeap{}
	heap.Init(&l)
	for _, node := range lists {
		if node != nil {
			heap.Push(&l, node)
		}
	}
	res := new(ListNode)
	curr := res
	for l.Len() > 0 {
		node := heap.Pop(&l).(*ListNode)
		if node.Next != nil {
			heap.Push(&l, node.Next)
		}
		node.Next = nil
		curr.Next = node
		curr = curr.Next
	}
	return res.Next

}

func main() {

}
