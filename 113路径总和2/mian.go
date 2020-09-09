package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res:=[][]int{}
	temp:=[]int{}
	var DFS func(root *TreeNode,temp []int,sum int)
	DFS = func(root *TreeNode,temp []int, sum int) {
		if root.Left==nil && root.Right==nil{
			if sum-root.Val==0{
				temp= append(temp,sum)
				ans:= make([]int,len(temp))
				copy(ans,temp)
				res= append(res,ans)
			}
		}
		temp=append(temp,root.Val)
		sum=sum-root.Val
		if root.Left==nil{
			DFS(root.Right,temp,sum)
			return
		}
		if root.Right==nil{
			DFS(root.Left,temp,sum)
			return
		}
		DFS(root.Left,temp,sum)
		DFS(root.Right,temp,sum)

	}
	DFS(root,temp,sum)
	return res
}

func main() {

}
