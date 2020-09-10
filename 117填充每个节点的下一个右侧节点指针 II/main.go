package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	list := make([]*Node, 0)
	list = append(list, root)
	list = append(list, nil)
	var node *Node
	for len(list) != 0 {
		node = list[0]
		list = list[1:]
		if node != nil {
			node.Next = list[0]
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}

		} else {
			if len(list) > 0 {
				list = append(list, nil)
			}
		}

	}
	return root
}



func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	curr:=root
	mostLeft:=&Node{}
	for curr != nil{
		tail:=mostLeft
		for curr!=nil{
			if curr.Left!=nil{
				tail.Next=curr.Left
				tail=tail.Next
			}
			if curr.Right!=nil{
				tail.Next=curr.Right
				tail=tail.Next
			}
			curr=curr.Next
		}
		curr=mostLeft.Next
		if tail==mostLeft{
			break
		}
		
	}
	return root
}


func main() {

}
