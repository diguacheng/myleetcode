package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for k > 0 {
		if p.Next == nil {
			p = head
		} else {
			p = p.Next
		}

		k--

	}
	q := head
	for p.Next != nil {
		q = q.Next
		p = p.Next
	}
	p.Next = head
	head = q.Next
	q.Next = nil
	return head
}

func main() {

}
