package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 失败了 
// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	if root==nil{
// 		return nil  
// 	}
// 	curr:=root 
// 	var pre *TreeNode
// 	for curr!=nil&&curr.Val!=key{
// 		pre=curr
// 		if curr.Val>key{
// 			curr=curr.Left
// 			continue
// 		}
// 		if curr.Val<key{
// 			curr=curr.Right
// 		}
// 	}
// 	if curr==nil{
// 		return root
// 	}
// 	del:=curr 
// 	for curr.Left!=nil||curr.Right!=nil{
// 		pre=curr
// 		if curr.Left!=nil{
// 			curr=curr.Left
// 			continue
// 		}
// 		if curr.Right!=nil{
// 			curr=curr.Right
// 		}
// 	}
// 	if del==curr&&del==root{
// 		return nil
// 	}
// 	curr.Val,del.Val=del.Val,curr.Val
// 	if pre.Left==curr{
// 		pre.Left=nil
// 	}else{
// 		pre.Right=nil
// 	}
// 	return root
	
// }

func main() {

}
