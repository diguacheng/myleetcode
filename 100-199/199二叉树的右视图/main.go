package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	l := make([][]*TreeNode, 2)

	l[0] = append(l[0], root)
	i := 0
	for len(l[i]) != 0 {

		if l[i][0].Left != nil {
			l[(i+1)%2] = append(l[(i+1)%2], l[i][0].Left)
		}

		if l[i][0].Right != nil {
			l[(i+1)%2] = append(l[(i+1)%2], l[i][0].Right)
		}

		v := l[i][0].Val
		l[i] = l[i][1:]
		if len(l[i]) == 0 {
			res = append(res, v)
			i = (i + 1) % 2
		}
	}
	return res

}

func main() {

}
