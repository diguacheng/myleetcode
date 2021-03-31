package main

import "hash"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 没做出来

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	m := make(map[int]int) // 前缀树的思想 存储在这个节点之前的父节点  前缀和的个数
	m[0] = 1               // 前缀和为0的节点有1个 root节点
	var help func(root *TreeNode, curr int, sum int) int
	help = func(root *TreeNode, curr int, sum int) int {
		if root == nil {
			return 0
		}
		res := 0
		curr += root.Val                   // 计算当前节点的前缀和
		res += m[curr-sum]                 // 综合为sum 当前节点为curr 则curr-sum的个数 就是以当前节点为路径终点的个数
		m[curr]++                          // 更新哈希表
		res += help(root.Left, curr, sum)  // 递归左子树
		res += help(root.Right, curr, sum) // 递归右子树
		m[curr]--                          // 因为m只记录当前节点的父节点的信息 所以这里要减去
		return res
	}

	return help(root, sum, 0)
}

func main() {

}
