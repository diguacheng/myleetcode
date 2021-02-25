package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	res := []int{}
	curr := root
	var node *TreeNode
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			res = append(res, curr.Val)
			curr = curr.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = node.Right
	}
	return res
}

func main() {

}
