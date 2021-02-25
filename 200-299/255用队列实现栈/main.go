package main

import "fmt"

type MyStack struct {
	list []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		list: make([]int, 0),
	}

}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	temp := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return temp

}

/** Get the top element. */
func (this *MyStack) Top() int {
	temp := this.list[len(this.list)-1]
	return temp

}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.list) == 0 {
		return true
	}
	return false

}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}
