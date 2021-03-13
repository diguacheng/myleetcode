package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sl1 := []int{}
	for l1 != nil {
		sl1 = append(sl1, l1.Val)
		l1 = l1.Next
	}
	sl2 := []int{}
	for l2 != nil {
		sl2 = append(sl2, l2.Val)
		l2 = l2.Next
	}
	var res *ListNode
	carry := 0
	for len(sl1) > 0 || len(sl2) > 0 || carry != 0 {
		a, b := 0, 0
		if len(sl1) == 0 {
			a = 0
		} else {
			a = sl1[len(sl1)-1]
			sl1 = sl1[:len(sl1)-1]
		}
		if len(sl2) == 0 {
			b = 0
		} else {
			b = sl2[len(sl2)-1]
			sl2 = sl2[:len(sl2)-1]
		}
		val := a + b + carry
		carry = val / 10
		val = val % 10
		l := new(ListNode)
		l.Val = val
		l.Next = res
		res = l
	}
	return res

}

func main() {

}
