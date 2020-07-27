package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	res := false
	if A != nil && B != nil {
		if A.Val == B.Val {
			res = hasSubstructure(A, B)
		}
		if !res {
			res = isSubStructure(A.Left, B)

		}
		if !res {
			res = isSubStructure(A.Right, B)
		}
	}
	return res

}

func hasSubstructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return hasSubstructure(A.Left, B.Left) && hasSubstructure(A.Right, B.Right)

}

func main() {

}
