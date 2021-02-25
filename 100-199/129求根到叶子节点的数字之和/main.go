package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var help func(root *TreeNode, sum int)
	res := 0
	help = func(root *TreeNode, sum int) {
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			res += sum
		}
		if root.Left != nil {
			help(root.Left, sum)
		}
		if root.Right != nil {
			help(root.Right, sum)
		}
	}
	help(root, 0)
	return res

}

//广度优先搜索
func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	res := 0
	for len(list) != 0 {
		node := list[0]
		list = list[1:]
		if node.Left == nil && node.Right == nil {
			res += node.Val
		}
		if node.Left != nil {
			node.Left.Val += node.Val * 10
			list = append(list, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val * 10
			list = append(list, node.Right)
		}
	}
	return res

}

func main() {

}
