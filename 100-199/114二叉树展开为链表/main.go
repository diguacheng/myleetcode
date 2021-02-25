package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var help func(root *TreeNode) *TreeNode
	help = func(root *TreeNode) *TreeNode {
		var leftSubTreeTail *TreeNode
		if root.Left != nil {
			leftSubTreeTail = help(root.Left)
			leftSubTreeTail.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		tail := root
		for tail.Right != nil {
			tail = tail.Right
		}
		return tail
	}
	curr := root
	var leftSubTreeTail *TreeNode
	for curr.Left != nil || curr.Right != nil {

		if curr.Left != nil {
			leftSubTreeTail = help(curr.Left)
			leftSubTreeTail.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right

	}
}

// 迭代解法
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	curr := root
	var preOrder *TreeNode
	for curr != nil {
		if curr.Left != nil {
			preOrder = curr.Left
			for preOrder.Right != nil {
				preOrder = preOrder.Right
			}
			preOrder.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}

}

func main() {

}
