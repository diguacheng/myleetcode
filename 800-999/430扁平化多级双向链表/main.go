package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	help(root)
	return root
}

func help(root *Node) (end *Node) {
	curr := root
	for curr != nil {
		if curr.Child != nil {
			e := help(curr.Child)
			e.Next = curr.Next
			if curr.Next != nil {
				curr.Next.Prev = e

			}
			curr.Next = curr.Child
			curr.Child.Prev = curr
			curr.Child = nil
		}
		end = curr
		curr = curr.Next
	}
	return end
}

func main() {

}
