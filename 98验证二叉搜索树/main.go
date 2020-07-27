package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	//flag,_,_:=help(root)
	return help1(root,math.MinInt64,math.MaxInt64)

}

func help(root *TreeNode)(flag bool, min,max int ){
	if root==nil{
		return true,0,0
	}
	rootval := root.Val
	if root.Left==nil&&root.Right==nil{
		return true,rootval,rootval
	}	
	if root.Left!=nil&&root.Right==nil{
		flag,min,max=help(root.Right)
		if  flag!=true||rootval>=min{
			return false,0,0
		}
		min=rootval
		return flag,min,max
		
	}
	if root.Right!=nil&&root.Left==nil{
		flag,min,max=help(root.Left)
		if  flag!=true||rootval<=max{
			return false,0,0
		}
		max=rootval
		return flag,min,max
	}
	flagl,_,maxl:=help(root.Left)
	flagr,minr,_:=help(root.Right)
	if flagl!=true||flagr!=true||maxl>=rootval||minr<=rootval {
		return false,0,0
	}
	return true,maxl,minr
}

func help1(root *TreeNode, lower,higher int)bool{
	if root==nil{
		return true
	}
	if root.Val<=lower||root.Val>=higher{
		return false
	}
	return help1(root.Left,lower,root.Val)&&help1(root.Right,root.Val,higher)
}

func main() {

}
