package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	res := 0
	DFS(root, sum, sum, &res)
	return res
}

func DFS(root *TreeNode, sum int, X int, res *int) {
	if root == nil {
		return
	}
	if X == 0 {
		*res += 1
	}
	DFS(root.Left, sum, X-root.Val, res)
	DFS(root.Right, sum, X-root.Val, res)
	DFS(root.Left, sum, sum, res)
	DFS(root.Right, sum, sum, res)

}
