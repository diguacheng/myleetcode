package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	// 中序遍历
	list := []int{root.Val}
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		help(root.Left)
		list = append(list, root.Val)
		help(root.Right)
	}
	help(root.Right)
	n := len(list)
	if n >= k {
		return list[n-k]
	}
	help(root.Left)
	return list[len(list)-(k-n)]
}

func kthLargest1(root *TreeNode, k int) int {
	// 右中左
	list := []int{}
	flag := false
	var help func(root *TreeNode, flag bool)
	help = func(root *TreeNode, flag bool) {
		if root == nil || flag {
			return
		}
		help(root.Right, flag)
		list = append(list, root.Val)
		if len(list) == k {
			flag = true
		}
		help(root.Left, flag)
	}
	help(root, flag)
	return list[k-1]

}

func kthLargest2(root *TreeNode, k int) int {
	// 右中左 8ms 最块
	res := 0
	flag := false
	var help func(root *TreeNode, flag bool)
	help = func(root *TreeNode, flag bool) {
		if root == nil || flag {
			return
		}
		help(root.Right, flag)
		k--
		if k == 0 {
			res = root.Val
			flag = true
		}

		help(root.Left, flag)
	}
	help(root, flag)
	return res

}

func kthLargest3(root *TreeNode, k int) int {
	// 右中左
	res := 0
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if k == 0 && root == nil {
			return
		}
		help(root.Right)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		help(root.Left)
	}
	help(root)
	return res

}

func main() {

}
