package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	return 0

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
