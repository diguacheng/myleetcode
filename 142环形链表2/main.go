package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for head != nil {
		m[head]++
		if m[head] == 2 {
			return head
		}
		head = head.Next

	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return nil
	}
	slow,fast :=head,head
	ans:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			for ans!=slow{
				ans=ans.Next
				slow=slow.Next
			}
			return ans
			
		}
	}
	return nil
	
}
func main() {

}
