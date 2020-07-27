package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return nil
	}
	a:=new(TreeNode)
	/*
	if len(preorder)==1{
		a.Val=preorder[0]
		return a
		// 这段代码去掉的原因是在下面的for循环已经包含了 
	}
	*/
	for i,k:=range inorder{
		if k==preorder[0]{
			a.Val=k
			a.Left=buildTree(preorder[1:i+1],inorder[0:i])
			a.Right=buildTree(preorder[i+1:],inorder[i+1:])
			break
		}
	}
	return a
}

func main(){
	preorder:=[]int{3,9,20,15,7}
	inorder:=[]int{9,3,15,20,7}
	fmt.Println(buildTree(preorder,inorder))

}