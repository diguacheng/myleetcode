package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	res := &ListNode{}
	res.Next = head
	p, q := res, res
	count := 0
	var temp *ListNode
	for q.Next != nil {
		if q.Next.Val < x {
			if count == 0 {
				p.Next = q.Next
				p = p.Next
				q = q.Next
			} else {
				temp = q.Next
				q.Next = q.Next.Next
				temp.Next = p.Next
				p.Next = temp
				p = p.Next
			}
		} else {
			count++
			q = q.Next
		}
	}
	return res.Next

}

func partition1(head *ListNode, x int) *ListNode {
	before, after := &ListNode{}, &ListNode{}
	b, a := before, after
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = a
	return b.Next

}
func main() {

}
