package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0{
		return nil
	}
	if len(inorder)==1{
		return &TreeNode{preorder[0],nil,nil}
	}
	root:=new(TreeNode)
	root.Val=preorder[0]
	var i int
	for i=0;i<len(inorder);i++{
		if preorder[0]==inorder[i]{
			break
		}
	}
	root.Left=buildTree(preorder[1:1+i],inorder[:i])
	root.Right=buildTree(preorder[i+1:],inorder[i+1:])
	return root

}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	fmt.Println(buildTree(preorder, inorder))


}
