package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return getlist(root)[k]

}

func getlist(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, getlist(root.Left)...)
	res = append(res, root.Val)
	res = append(res, getlist(root.Right)...)
	return res

}


func kthSmallest1(root *TreeNode, k int) int {
	res:=[]int{}
	stack:=make([]*TreeNode,0)
	curr:=root 
	for curr!=nil||len(stack)>0{
		for curr!=nil{
			stack=append(stack,curr)
			curr=curr.Left
		}
		node:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res=append(res,node.Val)
		if len(res)>=k{
			break
		}
		curr=node.Right
	}
	return res[k-1]
	

}


func main() {

}
