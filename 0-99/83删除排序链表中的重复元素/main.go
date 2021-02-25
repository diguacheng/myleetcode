package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p1, p2 *ListNode
	p1 = head
	p2 = p1.Next
	for p1 != nil && p2 != nil {
		if p1.Val == p2.Val {
			if p1.Next == p2 {
				p1.Next = p2.Next
				p2 = p1.Next
			} else {
				p2.Next = p1.Next
				p1 = p2.Next
			}
		} else {
			if p1.Next == p2 {
				p1 = p2.Next
			} else {
				p2 = p1.Next
			}
		}

	}
	return head

}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p *ListNode
	p = head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

func main() {

}
