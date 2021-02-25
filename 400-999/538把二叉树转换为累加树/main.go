package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var x int
	if root.Left != nil {
		x = help(root.Right, 0)
	}
	root.Val += x
	if root.Right != nil {
		help(root.Left, root.Val)
	}
	return root
}

func help(root *TreeNode, k int) int {
	x := 0
	y := 0
	if root.Right != nil {
		x = help(root.Right, k)
	}
	root.Val += x
	if root.Left != nil {
		y = help(root.Left, root.Val)
		return y
	}
	return root.Val
}

func convertBST1(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
		return sum
	}
	dfs(root)
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

func main() {

}
