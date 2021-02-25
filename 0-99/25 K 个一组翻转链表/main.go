package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	res := &ListNode{}
	var p *ListNode
	p = res
	var i int
	m := make(map[int]*ListNode)
	for {
		i = 0
		for i < k && head != nil {
			m[i] = head
			head = head.Next
			i++
		}
		if i == k {
			i--
			for i >= 0 {
				res.Next = m[i]
				res = res.Next
				i--

			}
			continue
		} else {
			for j := 0; j < i; j++ {
				res.Next = m[j]
				res = res.Next
			}
			res.Next = nil
			break

		}

	}
	return p.Next

}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	res := &ListNode{}
	var p *ListNode
	p = res
	var i int
	m := make([]*ListNode, k)
	for {
		i = 0
		for i < k && head != nil {
			m[i] = head
			head = head.Next
			i++
		}
		if i == k {
			i--
			for i >= 0 {
				res.Next = m[i]
				res = res.Next
				i--

			}
			continue
		} else {
			for j := 0; j < i; j++ {
				res.Next = m[j]
				res = res.Next
			}
			res.Next = nil
			break
		}

	}
	return p.Next

}

func main() {

}
