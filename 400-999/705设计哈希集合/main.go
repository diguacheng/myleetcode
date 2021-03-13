package main

import "fmt"

type MyHashSet struct {
	Nodes []*Node
	L     int
}

type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	l := 10000
	nodes := make([]*Node, l)
	return MyHashSet{
		Nodes: nodes,
		L:     l,
	}
}

func (this *MyHashSet) Add(key int) {
	idx := key % this.L
	if this.Nodes[idx] == nil {
		this.Nodes[idx] = &Node{
			Val: key,
		}
		return
	}
	node := this.Nodes[idx]
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &Node{
		Val: key,
	}
	return
}

func (this *MyHashSet) Remove(key int) {
	idx := key % this.L
	if this.Nodes[idx] == nil {
		return // error
	}
	if this.Nodes[idx].Val == key {
		this.Nodes[idx] = this.Nodes[idx].Next
		return
	}
	node := this.Nodes[idx]
	for node.Next != nil && node.Next.Val != key {
		node = node.Next
	}
	node.Next = node.Next.Next
	return

}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	idx := key % this.L
	if this.Nodes[idx] == nil {
		return false
	}
	node := this.Nodes[idx]
	for node != nil {
		if node.Val == key {
			return true
		}
		node = node.Next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1))
	fmt.Println(obj.Contains(3))
	fmt.Println(obj.Contains(2))
	obj.Remove(2)
	fmt.Println(obj.Contains(2))

}