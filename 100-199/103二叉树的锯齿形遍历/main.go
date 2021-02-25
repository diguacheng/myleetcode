package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	temp := []int{}
	l1 := []*TreeNode{}
	l2 := []*TreeNode{}
	l1 = append(l1, root)
	var node *TreeNode
	for len(l1) != 0 || len(l2) != 0 {
		for len(l1) != 0 {
			node = l1[0]
			l1 = l1[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				l2 = append(l2, node.Left)
			}
			if node.Right != nil {
				l2 = append(l2, node.Right)
			}
		}
		if len(temp) > 0 {
			res = append(res, temp)
			temp = []int{}
		}
		for len(l2) != 0 {
			node = l2[len(l2)-1]
			l2 = l2[:len(l2)-1]
			temp = append(temp, node.Val)
			if node.Right != nil {
				l1 = append(l1, node.Right)
			}
			if node.Left != nil {
				l1 = append(l1, node.Left)
			}
		}
		if len(temp) > 0 {
			res = append(res, temp)
			temp = []int{}
		}
		lenL1 := len(l1)
		for i := 0; i < lenL1/2; i++ {
			l1[i], l1[lenL1-i-1] = l1[lenL1-i-1], l1[i]
		}

	}
	return res

}

func main() {

	root := &TreeNode{Val: 3}
	cur := root
	cur.Left = &TreeNode{Val: 9}
	cur.Right = &TreeNode{Val: 20}
	cur = cur.Right
	cur.Left = &TreeNode{Val: 15}
	cur.Right = &TreeNode{Val: 7}
	fmt.Println(zigzagLevelOrder(root))
}
