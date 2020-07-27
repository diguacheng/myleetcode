package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 正常打印
func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	L := make([]*TreeNode, 0)
	if root != nil {
		L = append(L, root)
	}
	i := 0
	for i < len(L) {
		res = append(res, L[i].Val)
		if L[i].Left != nil {
			L = append(L, L[i].Left)

		}
		if L[i].Right != nil {
			L = append(L, L[i].Right)
		}
		i++

	}

	return res
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([]int, 0)
	res := make([][]int, 0)
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	nextLevel := 0 // 计数下一层需要打印的节点
	tobePrint := 1 // 代表这一层还需要打印的点
	for len(nodes) > 0 {
		x := nodes[0]
		temp = append(temp, x.Val)
		if x.Left != nil {
			nodes = append(nodes, x.Left)
			nextLevel++
		}
		if x.Right != nil {
			nodes = append(nodes, x.Right)
			nextLevel++
		}
		nodes = nodes[1:]
		tobePrint--
		if tobePrint == 0 {
			res = append(res, temp)
			temp = make([]int, 0)
			tobePrint = nextLevel
			nextLevel = 0
		}
	}
	return res

}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([]int, 0)
	res := make([][]int, 0)
	nodes := [2][]*TreeNode{}
	current := 0
	next := 1
	nodes[current] = append(nodes[current], root)
	for len(nodes[0]) > 0 || len(nodes[1]) > 0 {
		x := nodes[current][len(nodes[current])-1]
		nodes[current] = nodes[current][:len(nodes[current])-1]
		temp = append(temp, x.Val)
		if current == 0 {
			if x.Left != nil {
				nodes[next] = append(nodes[next], x.Left)
			}
			if x.Right != nil {
				nodes[next] = append(nodes[next], x.Right)
			}
		} else {
			if x.Right != nil {
				nodes[next] = append(nodes[next], x.Right)
			}
			if x.Left != nil {
				nodes[next] = append(nodes[next], x.Left)
			}

		}
		if len(nodes[current]) == 0 {
			res = append(res, temp)
			temp = make([]int, 0)
			current = 1 - current
			next = 1 - next
		}

	}
	return res

}

func main() {

}
