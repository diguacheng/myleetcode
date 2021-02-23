package main


type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 空间换时间
	l:=make([]*ListNode,0)
	for head!=nil{
		l=append(l,head)
		head=head.Next
	}
	n:=len(l)/2
	return l[n]

}

func middleNode1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 时间换空间
	l:=0
	var a *ListNode
	a=head
	for head!=nil{
		l++
		head=head.Next
	}
	n:=l/2
	for n>0{
		a=a.Next
		n--
	}
	return a
}


func middleNode2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 双指针 快指针是慢指针的双倍。
	slow,fast :=head,head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	return slow
	
}

func main() {

}