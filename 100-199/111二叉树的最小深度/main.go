package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return 1 + max(minDepth(root.Left), minDepth(root.Right))
	}

	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mind := 1 << 31
	var DFS func(root *TreeNode, depth int)
	DFS = func(root *TreeNode, depth int) {
		if root.Left == nil && root.Right == nil {
			depth++
			if depth < mind {
				mind = depth
			}
			return
		}
		if root.Left != nil {
			DFS(root.Left, depth+1)
		}
		if root.Right != nil {
			DFS(root.Right, depth+1)
		}
	}
	DFS(root, 0)
	return mind
}

func main() {

}
