package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	val := postorder[n-1]
	var i int
	for i = n - 2; i >= 0; i-- {
		if inorder[i] == val {
			break
		}
	}
	root := &TreeNode{Val: val}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:n-1])
	return root

}

func main() {

}
