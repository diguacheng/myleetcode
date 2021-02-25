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
	carry := 0
	var res *ListNode
	res = l1
	for l1 != nil {
		if l2 == nil {
			l2 = new(ListNode)
		}
		if l1.Val+l2.Val+carry >= 10 {
			l1.Val = l1.Val + l2.Val + carry - 10
			carry = 1
		} else {
			l1.Val = l1.Val + l2.Val + carry
			carry = 0
		}
		l2 = l2.Next
		if l1.Next == nil {
			break
		}
		l1 = l1.Next

	}
	for l2 != nil {
		if l2.Val+carry >= 10 {
			l2.Val = l2.Val + carry - 10
			carry = 1
		} else {
			l2.Val = l2.Val + carry
			carry = 0
		}
		l1.Next = l2
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry == 1 {
		temp := new(ListNode)
		temp.Val = 1
		l1.Next = temp
	}
	return res
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	curr := res
	carry := 0
	x, y := 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{carry, nil}
		curr.Next.Val = sum % 10
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return res.Next
}

func main() {

}
