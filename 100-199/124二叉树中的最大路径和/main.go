package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Mcount := -1 << 31 //不能设置为0
	var help func(root *TreeNode) int
	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		m1 := max(help(root.Left), 0)
		m2 := max(help(root.Right), 0) // 表示该节点的最大贡献值
		tmp := m1 + m2 + root.Val
		Mcount = max(Mcount, tmp)
		return max(m1, m2) + root.Val
	}
	help(root)
	return Mcount

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
