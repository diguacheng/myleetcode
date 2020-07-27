package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	Lorders(root, &l1)
	Rorders(root, &l2)
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func Lorders(root *TreeNode, l *[]int) {
	if root == nil {
		*l = append(*l, -1)

	} else {
		*l = append(*l, root.Val)
		Lorders(root.Left, l)
		Lorders(root.Right, l)

	}
}

func Rorders(root *TreeNode, l *[]int) {
	if root == nil {
		*l = append(*l, -1)

	} else {
		*l = append(*l, root.Val)
		Rorders(root.Right, l)
		Rorders(root.Left, l)

	}
}

func isSymmetric1(root *TreeNode) bool{
	if root == nil {
		return true
	}
	return isSame(root.Left,root.Right)
}

func isSame(l ,r *TreeNode) bool{
	if l== nil&&r==nil{
		return true
	}
	if l==nil||r==nil{
		return false
	}
	if l.Val!=r.Val{
		return false
	}
	return isSame(l.Left,r.Right)&&isSame(l.Right,r.Left)
}

func main() {

}
