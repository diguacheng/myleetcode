package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil || head.Val == val && head.Next == nil {
		return nil
	}
	if head.Val == val && head.Next != nil {
		head = head.Next
		return head
	}
	var p *ListNode = head
	for p.Next != nil {
		if val == p.Next.Val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
	return head

}
func main() {

}
