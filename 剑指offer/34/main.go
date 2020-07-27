package main



type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	temp := make([]int, 0)
	helper(root, sum, temp, &res)
	return res

}

func helper(root *TreeNode, sum int, temp []int, res *[][]int) {
	x := sum - root.Val
	if x == 0 && root.Left == nil && root.Right == nil {
		temp = append(temp, root.Val)
		r := make([]int, len(temp))
		copy(r, temp)// 切片是引用类型，指向底层数组的指针结构体，为了防止底层数据改变引发问题，再添加的时候，复制一个新的
		*res = append(*res, r)
		temp = temp[:len(temp)-1]

	} else {
		temp = append(temp, root.Val)
		if root.Left != nil {
			helper(root.Left, x, temp, res)
		}
		if root.Right != nil {
			helper(root.Right, x, temp, res)
		}
		temp = temp[:len(temp)-1]

	}

}
func main() {

}
