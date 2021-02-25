package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	var res, prev *ListNode
	prev = head
	res = head
	for k > 0 {
		prev = prev.Next
		k--
	}
	for prev != nil {
		res = res.Next
		prev = prev.Next

	}
	return res

}

func main() {

}
