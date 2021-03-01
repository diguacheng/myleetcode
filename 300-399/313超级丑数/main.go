package main

import (
	"container/heap"
	"fmt")



type MinHeap []int

func (m MinHeap)Len()int{ return len(m)}

func (m MinHeap)Less(i,j int)bool{ return m[i]<m[j] }

func (m MinHeap)Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}

func (m *MinHeap)Pop()interface{}{
	old:=*m 
	n:=len(old)
	x:=old[n-1]
	*m=old[:n-1]
	return x
}
func (m *MinHeap)Push(x interface{}){
	
	*m=append(*m,x.(int))
}
func nthSuperUglyNumber(n int, primes []int) int {
	if n==0||n==1{
		return n
	}
	h:=new(MinHeap)
	heap.Push(h,1)
	n-- 
	for n>0{
		x:=heap.Pop(h).(int)
		for h.Len()>0&&x==(*h)[0]{
			heap.Pop(h)
		}
		for i := 0; i < len(primes); i++ {
            heap.Push(h, x * primes[i])
        }
		n--
	}
	return heap.Pop(h).(int)
}



func main() {
	n := 12
	primes := []int{2, 7, 13, 19}
	fmt.Println(nthSuperUglyNumber(n,primes))

}