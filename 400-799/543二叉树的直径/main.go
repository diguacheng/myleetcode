package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	depth(root, &ans)
	return ans - 1
}
func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	L := depth(root.Left, ans)
	R := depth(root.Right, ans)
	if L+R+1 > *ans {
		*ans = L + R + 1
	}
	if L > R {
		return L + 1
	}
	return R + 1
}

func main() {

}
