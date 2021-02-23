package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归 
	if root==nil{
		return nil
	}
	if p.Val>q.Val{
		p,q=q,p
	}
	if q.Val<root.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}
	if p.Val>root.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root

}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 迭代
	if root==nil{
		return nil
	}
	for root!=nil{
		if p.Val<root.Val&&q.Val<root.Val{
			root=root.Left
			continue
		}
		if p.Val>root.Val&&q.Val>root.Val{
			root=root.Right
			continue
		}
		return root
	}
	return root

}

func main() {




}
