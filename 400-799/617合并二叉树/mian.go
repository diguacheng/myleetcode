package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
		return t1
	}
	if t1 != nil && t2 == nil {
		return t1
	}
	if t1 == nil && t2 != nil {
		return t2
	}
	return nil

}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	stack := make([][]*TreeNode, 0)
	stack = append(stack, []*TreeNode{t1, t2})
	var curr1, curr2 *TreeNode
	for len(stack) > 0 {
		curr1 = stack[len(stack)-1][0]
		curr2 = stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		curr1.Val += curr2.Val
		if curr1.Left != nil && curr2.Left != nil {
			stack = append(stack, []*TreeNode{curr1.Left, curr2.Left})
		}
		if curr1.Left == nil && curr2.Left != nil {
			curr1.Left = curr2.Left
		}
		if curr1.Right != nil && curr2.Right != nil {
			stack = append(stack, []*TreeNode{curr1.Right, curr2.Right})
		}
		if curr1.Right == nil && curr2.Right != nil {
			curr1.Right = curr2.Right
		}
	}
	return t1
}

func main() {

}
