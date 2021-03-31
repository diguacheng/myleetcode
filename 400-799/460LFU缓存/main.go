package main

import "fmt"

type LFUCache struct {
	Size     int
	Capacity int
	Items    []item
}

type item struct {
	Times int
	Key   int
	Value int
}

func Constructor(capacity int) LFUCache {
	L := LFUCache{
		Size:     0,
		Capacity: capacity,
		Items:    make([]item, 0),
	}
	return L
}

func (this *LFUCache) Get(key int) int {
	idx := this.find(key)
	if idx == -1 {
		return -1
	}
	v := this.Items[idx].Value
	this.Items[idx].Times++
	this.insertSort(idx)
	return v

}

// 查找key对应的索引 不排序
func (this *LFUCache) find(key int) int {
	for i := 0; i < this.Size; i++ {
		if this.Items[i].Key == key {
			return i
		}
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	idx := this.find(key)
	if idx == -1 {
		it := item{Times: 1, Key: key, Value: value}
		this.Size++
		if this.Size > this.Capacity {
			this.Items = this.Items[:this.Capacity-1]
			this.Size--
		}
		this.Items = append(this.Items, it)
		this.insertSort(this.Size-1)
		return
	}
	this.Items[idx].Times++
	this.Items[idx].Value = value
	this.insertSort(idx)
}

func (this *LFUCache) insertSort(idx int) {
	i := idx - 1
	for i >= 0 {
		if this.Items[i].Times <= this.Items[i+1].Times {
			this.Items[i], this.Items[i+1] = this.Items[i+1], this.Items[i]
		} else {
			break
		}
		i--
	}

}

func main() {
	l := Constructor(3)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)
	fmt.Println(l.Get(4))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(1))
	l.Put(5, 5)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
	fmt.Println(l.Get(5))


}