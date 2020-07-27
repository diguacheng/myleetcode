package main

import "fmt"

type CQueue struct {
	input  []int
	output []int
}

func Constructor() CQueue {
	return CQueue{
		make([]int, 0, 100),
		make([]int, 0, 100),
	}

}

func (this *CQueue) AppendTail(value int) {
	this.input = append(this.input, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.input) == 0 && len(this.output) == 0 {
		return -1
	}
	if len(this.output) != 0 {
		x := this.output[len(this.output)-1]
		this.output = this.output[:len(this.output)-1]
		return x
	}
	for len(this.input) > 0 {
		this.output = append(this.output, this.input[len(this.input)-1])
		this.input = this.input[:len(this.input)-1]
	}
	x := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return x
}

func main() {
	x := Constructor()
	x.AppendTail(5)
	x.AppendTail(3)
	x.AppendTail(2)
	fmt.Println(x.DeleteHead())
	fmt.Println(x.DeleteHead())
}
