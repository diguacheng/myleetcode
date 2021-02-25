package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head != nil {
		res := reversePrint(head.Next)
		return append(res, head.Val)
	}
	return []int{}

}

func main() {

}
