package main

import "fmt"

type MinStack struct {
	data []int
	list []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}

}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if x <= this.Min() || len(this.list) == 0 {
		this.list = append(this.list, x)
	}

}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		if this.data[len(this.data)-1] == this.list[len(this.list)-1] {
			this.list = this.list[:len(this.list)-1]
		}
		this.data = this.data[:len(this.data)-1]

	}

}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0

}

func (this *MinStack) Min() int {
	if len(this.list) > 0 {
		return this.list[len(this.list)-1]
	}
	return 0

}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.Min())

}
