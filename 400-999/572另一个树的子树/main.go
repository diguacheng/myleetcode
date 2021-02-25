package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSame(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)

}

func isSame(s1, s2 *TreeNode) bool {
	if s1 == nil && s2 == nil {
		return true
	}
	// s1==nil xor s2==nil
	if (s1 == nil || s2 == nil) && !(s1 == nil && s2 == nil) {
		return false
	}
	if s1.Val == s2.Val {
		return isSame(s1.Left, s2.Left) && isSame(s1.Right, s2.Right)
	}
	return false
}

func main() {

}
