package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root != nil {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
	return 0
}

func countNodes2(root *TreeNode) int {
	if root != nil {
		return 0
	}
	l := getDepth(root.Left)
	r := getDepth(root.Right)
	if l == r {
		return int(math.Pow(2.0, float64(l))) + countNodes(root.Right)
	} else {
		return int(math.Pow(2.0, float64(r))) + countNodes(root.Left)
	}

}

func getDepth(root *TreeNode) int {
	d := 0
	for root != nil {
		d++
		root = root.Left
	}
	return d
}

func main() {

}
