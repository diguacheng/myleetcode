package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {
	//迭代
	if head == nil || head.Next == nil {
		return head
	}
	var p, q = head, head.Next
	head = head.Next
	p.Next = q.Next
	q.Next = p
	q = p.Next
	for q != nil && q.Next != nil {
		p.Next = q.Next
		q.Next = p.Next.Next
		p.Next.Next = q
		p = q
		q = p.Next
	}
	return head

}

func swapPairs(head *ListNode) *ListNode {
	//递归
	if head == nil || head.Next == nil {
		return head
	}
	var p, q = head, head.Next
	head = head.Next

	p.Next = swapPairs(q.Next)
	q.Next = p
	return head

}

func main() {

}
