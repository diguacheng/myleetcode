package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		if root.Left == nil {
			node := new(TreeNode)
			node.Val = val
			root.Left = node

		} else {
			insertIntoBST(root.Left, val)
		}

	}
	if val > root.Val {
		if root.Right == nil {
			node := new(TreeNode)
			node.Val = val
			root.Right = node

		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root

}
//迭代
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	curr := root
	for {
		if val < curr.Val {
			if curr.Left == nil {
				node := &TreeNode{Val: val}
				curr.Left = node
				break
			} else {
				curr = curr.Left
			}

		} else {
			if curr.Right == nil {
				node := &TreeNode{Val: val}
				curr.Right = node
				break
			} else {
				curr = curr.Right
			}

		}
	}
	return root

}

func main() {

}
