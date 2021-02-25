package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	// 返回选中根节点的最大值，不选中根节点的最大值
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]               // 如果是选中 那 左右子树就要加上不选中的值
	notSelected := max(l[0], l[1]) + max(r[0], r[1]) // 如果不选根节点，子树的根 就可选可不选
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
