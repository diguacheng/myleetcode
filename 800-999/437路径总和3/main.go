package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 没做出来

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	var help func(root *TreeNode, curr int)
	help = func(root *TreeNode, curr int) {
		if root == nil {
			return
		}
		if curr-root.Val == 0 {
			res++
		}
		help(root.Left, curr-root.Val)
		help(root.Right, curr-root.Val)
		help(root.Left, sum)
		help(root.Right, sum)
	}

	help(root, sum)
	return res

}

func main() {

}
