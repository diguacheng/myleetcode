package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var help func(root *TreeNode) (int, bool)
	help = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		l, flagl := help(root.Left)
		r, flagr := help(root.Right)
		t := flagl && flagr
		if !t {
			return 0, false
		}
		if math.Abs(float64(l-r)) <= 1 && t {
			return 1 + max(l, r), true
		}
		return 1 + max(l, r), false

	}
	_, flag := help(root)
	return flag
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}
