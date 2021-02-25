package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	l := len(postorder)
	root := postorder[l-1] //根元素的值
	flag := 0
	// 找到二叉搜索树比根的值大的点的第一个人索引
	for i := l - 2; i >= 0; i-- {
		if postorder[i] < root {
			flag = i + 1
			break
		}
	}
	// 索引左边的值都应该小于根节点的值，若有大于的 返回false
	for i := 0; i < flag; i++ {
		if postorder[i] > root {
			return false
		}
	}
	// 继续判断根的左子树和右子树是否满足二叉搜索树的要求
	return verifyPostorder(postorder[:flag]) && verifyPostorder(postorder[flag:l-1])
}

func main() {
	postorder := []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(postorder))

}
