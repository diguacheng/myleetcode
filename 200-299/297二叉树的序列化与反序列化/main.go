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

	root, _ := rdeserialize(dataSlice, 0)
	return root
}

func rdeserialize(l []string, i int) (*TreeNode, int) {

	if l[i] == "nil" {
		i++
		return nil, i
	}
	num, _ := strconv.Atoi(l[i])

	root := &TreeNode{Val: num}

	var a, b int
	root.Left, a = rdeserialize(l, i+1)
	root.Right, b = rdeserialize(l, a)
	return root, b

}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {
	s := []string{"1", "2", "3", "nil", "nil", "4", "nil", "nil", "nil"}
	fmt.Println(rdeserialize(s, 0))
	de := "1,2,3,4,5,6,7,,"
	dd := strings.Split(de, ",")
	fmt.Println(dd, len(dd))

}
