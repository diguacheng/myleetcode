package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	curr := root
	var prev *TreeNode
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			// 后序遍历，就沿着该节点左边，依次入栈，但不访问
			stack = append(stack, curr)
			curr = curr.Left
		}
		node := stack[len(stack)-1]
		// 这里判断一下弹出的node节点的右子树是不是访问前一个节点
		// 或者node没有右子树  5
		if node.Right == nil || node.Right == prev {
			result = append(result, node.Val) // 访问
			prev = node
			stack = stack[:len(stack)-1]
			curr = nil
		} else {
			// 右节点还没访问，先访问右节点 4
			curr = node.Right
		}
	}
	return result
}

func main() {

}
