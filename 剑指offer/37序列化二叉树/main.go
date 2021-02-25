package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return rserialize(root, "")

}

func rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "nil,"
		return str
	}

	str += strconv.Itoa(root.Val) + ","

	str = rserialize(root.Left, str)
	str = rserialize(root.Right, str)
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dataSlice := strings.Split(data, ",")
	dataSlice = dataSlice[:len(dataSlice)-1]
	//i := 0
	//root, i := rdeserialize(dataSlice, i)
	root := rdeserialize1(&dataSlice)
	return root
}

func rdeserialize(l []string, i int) (*TreeNode, int) {

	if l[i] == "nil" {
		(i)++
		return nil, i
	}
	num, _ := strconv.Atoi(l[i])

	root := new(TreeNode)
	root.Val = num
	i++
	a, b := 0, 0
	root.Left, a = rdeserialize(l, i)
	root.Right, b = rdeserialize(l, a)
	return root, b

}

func rdeserialize1(l *[]string) *TreeNode {
	if (*l)[0] == "nil" {
		(*l) = (*l)[1:]
		return nil
	}
	num, _ := strconv.Atoi((*l)[0])
	root := new(TreeNode)
	root.Val = num
	(*l) = (*l)[1:]
	root.Left = rdeserialize1(l)
	root.Right = rdeserialize1(l)
	return root

}

func main() {

	obj := Constructor()
	str := "1,2,nil,nil,3,4,nil,nil,5,nil,nil,"
	ans := obj.deserialize(str)
	ssr := obj.serialize(ans)
	fmt.Println(ssr, "\n", str)

}
