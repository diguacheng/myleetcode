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
			if node.Left != nil { // 注意这里不能加入空节点。
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
	curr := root
	var head *Node
	for curr.Left != nil {
		head = curr
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		curr = curr.Left
	}
	return root
}

func main() {

}
