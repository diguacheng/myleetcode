package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 先遍历一遍，求出两个链表的长度，然后较长的向后移动，使两个链表从头指针到末尾的距离相同
	if headA == nil || headB == nil {
		return nil
	}
	lA, lB := 0, 0
	p := headA
	for p != nil {
		p = p.Next
		lA++
	}
	p = headB
	for p != nil {
		p = p.Next
		lB++
	}
	// 求出长度
	if lA > lB {
		lA, lB = lB, lA
		headA, headB = headB, headA
	}
	for lA < lB {
		lB--
		headB = headB.Next
	}
	// 将较长的链表头指针往后移动
	for lA > 0 {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	// 找相同的节点
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	//  双指针
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB
	var enda, endb *ListNode
	for p != q {
		if p.Next == nil {
			enda = p
			if endb != nil && enda != endb {
				return nil
			}
			p = headB
		} else {
			p = p.Next
		}
		if q.Next == nil {
			endb = q
			if enda != nil && enda != endb {
				return nil
			}
			q = headA
		} else {
			q = q.Next
		}

	}
	return p

}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	//用一个哈希表保存链表a， 检查b中有无相同的

	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]int)
	for headA != nil {
		m[headA]++
		headA = headA.Next
	}
	for headB != nil {
		if m[headB] == 1 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
func main() {

}
