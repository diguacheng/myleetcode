package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root==nil{
		return false
	}
	var DFS func(root *TreeNode,sum int) bool
	DFS= func(root *TreeNode,sum int) bool{
		if root.Left==nil&& root.Right == nil{
			if sum-root.Val == 0 {
				return true
			}
			return false
		}
		sum=sum-root.Val
		
		if root.Left==nil{
			return DFS(root.Right, sum)
		}
		if root.Right == nil {
			return DFS(root.Left, sum)
		}
		return  DFS(root.Right, sum)||DFS(root.Left, sum)
	}
	return DFS(root,sum)
	
}

func main() {

}
