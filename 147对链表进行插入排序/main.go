package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	res:=&ListNode{0,nil} 
	node:=head
	head=head.Next
	node.Next=nil 
	res.Next=node
	var prev,curr *ListNode
	for head!=nil{
		node=head
		head=head.Next
		node.Next=nil 
		prev,curr=res,res.Next
		for curr != nil{
			if node.Val>=curr.Val{
				prev=curr
				curr=curr.Next
			}else{
				break
			}
		}
		node.Next=curr
		prev.Next=node
	}
	return res.Next
}

func insertionSortList1(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	res:=&ListNode{0,nil} 
	node:=head
	head=head.Next
	node.Next=nil 
	res.Next=node
	var curr *ListNode
	for head!=nil{
		node=head
		head=head.Next
		node.Next=nil 
		curr=res
		for curr.Next != nil&&node.Val>=curr.Next.Val{
			curr=curr.Next
		}
		node.Next=curr.Next
		curr.Next=node
	}
	return res.Next
}



func main() {

}
