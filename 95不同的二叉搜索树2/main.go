package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start,end int)[]*TreeNode {
	if start>end{
		return []*TreeNode{nil}
	}
	res:=make([]*TreeNode,0)
	for i:=start; i <=end; i++{
		leftTrees:=helper(start,i-1)
		rightTrees:=helper(i+1,end)
		for _,left:= range leftTrees{
			for _,right:= range rightTrees{
				root:=&TreeNode{i,nil,nil}
				root.Left=left
				root.Right=right
				res=append(res,root)
			}
		}
	}
	return res
}

// 这道题 应该选用递归做，二叉搜索树的性质就是左子树的值都比根节点小 右子树的值 都比根节点大 
// 因此 不同的值做根节点，就会有不同的结果，
// 可以选定根节点i，在递归地求左子树0..i-1 和右子树i+1..n 
//  我没有做出来的原因是 求解递归函数的时候，没有考虑清楚 节点范围为0时的情况。
// 上面解法中，当start>end 说明最左侧的点时根节点， 当start=end 时 说明该子树只有一个点 。

// 失败的代码 
// func generateTrees(n int) []*TreeNode {
// 	trees := make([]*TreeNode, n)
// 	res := make([]*TreeNode, 0)
// 	for i := 1; i < n; i++ {
// 		trees[i] = &TreeNode{Val: i}
// 	}
// 	for i := 0; i < n; i++ {
// 		res = append(res, help(trees[0:i], trees[i+1:n], trees[i])...)

// 	}
// 	fmt.Println(len(res))
// 	return res
// }

// func help(left, right []*TreeNode, root *TreeNode) []*TreeNode {
// 	res := make([]*TreeNode, 0)
// 	resl := make([]*TreeNode, 0)
// 	resr := make([]*TreeNode, 0)
// 	l := len(left)
// 	r := len(right)

// 	for i := 0; i < l; i++ {
// 		resl = append(resl, help(left[0:i], left[i+1:l], left[i])...)
// 	}
// 	for i := 0; i < r; i++ {
// 		resr = append(resr, help(right[0:i], right[i+1:l], right[i])...)
// 	}
// 	rl, rr := len(resl), len(resr)
// 	if rl == 0 && rr == 0 {
// 		res = append(res, root)
// 		return res
// 	}
// 	if rl == 0 && rr != 0 {
// 		root.Left = nil
// 		for j := 0; j < len(resr); j++ {
// 			root.Right = resr[j]
// 			res = append(res, root)

// 		}
// 		return res

// 	}
// 	if rl != 0 && rr == 0 {
// 		root.Right = nil
// 		for j := 0; j < len(resl); j++ {
// 			root.Right = resl[j]
// 			res = append(res, root)

// 		}
// 		return res
// 	}
// 	for i := 0; i < len(resl); i++ {
// 		root.Left = resl[i]
// 		for j := 0; j < len(resr); j++ {
// 			root.Right = resr[j]
// 			res = append(res, root)

// 		}
// 	}
// 	return res

// }

func main() {
	fmt.Println(generateTrees(3))

}
