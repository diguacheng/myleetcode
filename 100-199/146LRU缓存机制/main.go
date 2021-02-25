package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	head, tail *DLinknode
	cache      map[int]*DLinknode
}

type DLinknode struct {
	key, value int
	Prev, Next *DLinknode
}

func NewDLinknode(key, value int) *DLinknode {

	return &DLinknode{key, value, nil, nil}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		head:     NewDLinknode(0, 0),
		tail:     NewDLinknode(0, 0),
		cache:    map[int]*DLinknode{},
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l

}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.movetohead(node)
		return node.value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		newnode := NewDLinknode(key, value)
		this.cache[key] = newnode
		this.addToHead(newnode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.movetohead(node)

	}

}

func (this *LRUCache) movetohead(node *DLinknode) {
	this.romoveNode(node)
	this.addToHead(node)
}

func (this *LRUCache) romoveNode(node *DLinknode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *DLinknode) {
	node.Next = this.head.Next
	this.head.Next.Prev = node
	node.Prev = this.head
	this.head.Next = node
}

func (this *LRUCache) removeTail() *DLinknode {
	node := this.tail.Prev
	this.romoveNode(node)
	return node
}

func main() {
	s := Constructor(2)
	s.Put(1, 1)
	s.Put(2, 2)
	fmt.Println(s.Get(1))

	s.Put(3, 3)
	fmt.Println(s.Get(2))
	s.Put(4, 4)
	fmt.Println(s.Get(1))
	fmt.Println(s.Get(3))
	fmt.Println(s.Get(4))

}
