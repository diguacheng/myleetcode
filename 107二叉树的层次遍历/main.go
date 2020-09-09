package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	list:=make([]*TreeNode,0)
	list= append(list, root)
	list = append(list, nil)
	temp:=[]int{}
	res:=[][]int{}
	var node *TreeNode
	for len(list)!=0{
		node=list[0]
		list =list[1:]
		if node==nil{
			res = append(res, temp)
			temp=[]int{}
			if len(list)!=0{
				list = append(list, nil)
			}
	
		}else{
			temp=append(temp,node.Val)
			if node.Left!=nil{
				list = append(list, node.Left)
			}
			if node.Right!=nil{
				list = append(list, node.Right)
			}
		}
	}
	// reverse the res 
	n:=len(res)
	for i:=0;i<n/2;i++{
		res[i],res[n-1-i] = res[n-1-i],res[i]
	}
	return res
}

func main() {

}
