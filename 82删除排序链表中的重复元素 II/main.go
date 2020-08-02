package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	d := make(map[int]int)
	p := head
	for p != nil {
		d[p.Val]++
		p = p.Next
	}
	p = head
	res := new(ListNode)
	q := res
	for p != nil {
		if d[p.Val] == 1 {
			q.Next = p
			q = q.Next
		}
		p = p.Next
	}
	q.Next=nil
	return res.Next
}

// 双指针
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res:=&ListNode{}
	res.Next=head
	pre,cur:=res,head
	for cur!=nil{
		for cur.Next != nil && cur.Val==cur.Next.Val{
			cur=cur.Next// cur节点与下一个值相同，就一直往后移
		}
		if pre.Next==cur{
			// cur 值没变，说明 cur 节点是非重读节点
			pre=pre.Next
		}else{
			pre.Next=cur.Next
			// 否则， cur节点指向的是重复节点的最后一个
			// 将pre指向cur的下一个节点，
		}
		cur=pre.Next
	}
	return res.Next

}

func main() {
	in := []int{1, 1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Println(deleteDuplicate1(in))

}
