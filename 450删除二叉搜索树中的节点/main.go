package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil  
	}
	// find the target node
	curr:=root 
	var pre *TreeNode
	for curr!=nil&&curr.Val!=key{
		pre=curr
		if curr.Val>key{
			curr=curr.Left
			continue
		}
		if curr.Val<key{
			curr=curr.Right
		}
	}
	if curr==nil{
		// 没找到 返回根节点 
		return root
	}
	// delet node 
	node:=delnode(curr)
	if pre==nil{
		// 若 pre 为nil 说明删除的是根节点
		return node
	}
	// 判断是用前继还是后继节点 
	if pre.Left==curr{
		pre.Left=node
	}else{
		pre.Right=node
	}
	return root	
}

func delnode(node *TreeNode)*TreeNode{
	//  若为 叶子节点 则直接删除 返回nil 
	if node.Left==nil&&node.Right==nil{
		return nil
	}
	// 用前继节点补位 
	var curr,pre *TreeNode
	if node.Left!=nil{
		curr=node.Left
		for curr.Right!=nil{
			pre=curr 
			curr=curr.Right
		}
		if pre==nil{
			// 说明 被删除的节点的左节点 没有右子树  被删除的节点的左节点 为前继节点 
			node.Left.Right=node.Right
			node=node.Left
			return node
		}
		
		pre.Right=curr.Left
		curr.Right=node.Right
		curr.Left=node.Left
		return curr
	}
	if node.Right!=nil{
		curr=node.Right
		for curr.Left!=nil{
			// 要删除的右子节点 没有子树 
			pre=curr 
			curr=curr.Left
		}
		if pre==nil{
			node=node.Right
			return node
		}
		pre.Left=curr.Right
		curr.Right=node.Right 
		return curr
	}
	return nil
}

func main() {

}
