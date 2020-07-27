package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return 1+max(maxDepth(root.Left),maxDepth(root.Right))

}

func max(a,b int) int {
	if a<b{
		return b
	}
	return a
}

func main() {
	fmt.Println("111")

}