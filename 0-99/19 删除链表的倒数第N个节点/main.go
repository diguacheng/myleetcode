package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p,q *ListNode
	p=head
	q=head
	for n>0&&p.Next!=nil{
		p=p.Next
		n--
	}
    if n>0{
        return head.Next
    }
	for p.Next != nil {
		p=p.Next
		q=q.Next
	}
	q.Next=q.Next.Next
	return head

}

func main() {

}
