package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	l:=maxDepth(root.Left)
	r:=maxDepth(root.Right)
	if math.Abs(float64(l-r))<=1.0{
		return isBalanced(root.Left)&&isBalanced(root.Right)
	}
	return false

}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return 1+max(maxDepth(root.Left),maxDepth(root.Right))

}

func max(a,b int) int {
	if a<b{
		return b
	}
	return a
}


func isBalanced1(root *TreeNode) bool{
	if root==nil{
		return true
	}
	//m:=make(map[*TreeNode]int)
	var help func(root *TreeNode)(int,bool)
	help = func(root *TreeNode)(int,bool){
		if root==nil{
			return 0,true
		}
		l,flagl:=help(root.Left)
		r,flagr:=help(root.Right)
		if math.Abs(float64(l-r))<=1&&flagl&&flagr{
			return max(l,r),true
		}
		return max(l,r),false

	}
	_,flag:=help(root)
	return flag
}

func isBalanced2(root *TreeNode) bool{
	if root==nil{
		return true
	}
	var help func(root *TreeNode)(int,bool)
	help = func(root *TreeNode)(int,bool){
		if root==nil{
			return 0,true
		}
		l,flagl:=help(root.Left)
		r,flagr:=help(root.Right)
		t:=flagl&&flagr
		if !t{
			return 0,false
		}
		if math.Abs(float64(l-r))<=1&&t{
			return 1+max(l,r),true
		}
		return 1+max(l,r),false

	}
	_,flag:=help(root)
	return flag
}

func main() {

}
