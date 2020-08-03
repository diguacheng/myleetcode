package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	cur:=root
	list:=make([]*TreeNode, 0)
	for cur!=nil||len(list)!=0{
		for cur!=nil{
			list = append(list, cur)
			cur=cur.Left
		}
		cur=list[len(list)-1]
		list=list[:len(list)-1]
		res=append(res,cur.Val)
		cur=cur.Right

	}
	return res
}

func main() {

}
