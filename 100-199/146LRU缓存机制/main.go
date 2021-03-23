package main

import "fmt"

// type LRUCache struct {
// 	size       int
// 	capacity   int
// 	head, tail *DLinknode
// 	cache      map[int]*DLinknode
// }

// type DLinknode struct {
// 	key, value int
// 	Prev, Next *DLinknode
// }

// func NewDLinknode(key, value int) *DLinknode {

// 	return &DLinknode{key, value, nil, nil}
// }

// func Constructor(capacity int) LRUCache {
// 	l := LRUCache{
// 		size:     0,
// 		capacity: capacity,
// 		head:     NewDLinknode(0, 0),
// 		tail:     NewDLinknode(0, 0),
// 		cache:    map[int]*DLinknode{},
// 	}
// 	l.head.Next = l.tail
// 	l.tail.Prev = l.head
// 	return l

// }

// func (this *LRUCache) Get(key int) int {
// 	if node, ok := this.cache[key]; ok {
// 		this.movetohead(node)
// 		return node.value
// 	}
// 	return -1

// }

// func (this *LRUCache) Put(key int, value int) {
// 	if _, ok := this.cache[key]; !ok {
// 		newnode := NewDLinknode(key, value)
// 		this.cache[key] = newnode
// 		this.addToHead(newnode)
// 		this.size++
// 		if this.size > this.capacity {
// 			removed := this.removeTail()
// 			delete(this.cache, removed.key)
// 			this.size--
// 		}
// 	} else {
// 		node := this.cache[key]
// 		node.value = value
// 		this.movetohead(node)

// 	}

// }

// func (this *LRUCache) movetohead(node *DLinknode) {
// 	this.romoveNode(node)
// 	this.addToHead(node)
// }

// func (this *LRUCache) romoveNode(node *DLinknode) {
// 	node.Prev.Next = node.Next
// 	node.Next.Prev = node.Prev
// }

// func (this *LRUCache) addToHead(node *DLinknode) {
// 	node.Next = this.head.Next
// 	this.head.Next.Prev = node
// 	node.Prev = this.head
// 	this.head.Next = node
// }

// func (this *LRUCache) removeTail() *DLinknode {
// 	node := this.tail.Prev
// 	this.romoveNode(node)
// 	return node
// }

type LRUCache struct {
	Head     *DLink
	Tail     *DLink
	Capacity int
	Size     int
	DLinks   map[int]*DLink
}

type DLink struct {
	Key  int
	Val  int
	Pre  *DLink
	Next *DLink
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		Head:     &DLink{},
		Tail:     &DLink{},
		Capacity: capacity,
		Size:     0,
		DLinks:   make(map[int]*DLink),
	}
	c.Head.Next = c.Tail
	c.Tail.Pre = c.Head
	return c
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.DLinks[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1

}

func (this *LRUCache) moveToHead(node *DLink) {
	this.Remove(node)
	this.appendToHead(node)
}

func (this *LRUCache) Remove(node *DLink) {
	fmt.Println("romove", node.Val)
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) appendToHead(node *DLink) {
	node.Next = this.Head.Next
	this.Head.Next.Pre = node
	this.Head.Next = node
	node.Pre = this.Head

}

func (this *LRUCache) removeTail() *DLink {
	node := this.Tail.Pre
	this.Remove(node)
	return node
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.DLinks[key]; ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	newNode := &DLink{key, value, nil, nil}
	this.DLinks[key] = newNode
	this.appendToHead(newNode)
	this.Size++
	if this.Size > this.Capacity {
		node := this.removeTail()
		delete(this.DLinks, node.Key)
		this.Size--
	}
	return
}

func main() {
	s := Constructor(2)
	s.Put(1, 0)
	s.Put(2, 2)
	fmt.Println(s.Get(1))

	s.Put(3, 3)
	fmt.Println(s.Get(2))
	s.Put(4, 4)
	fmt.Println(s.Get(1))
	fmt.Println(s.Get(3))
	fmt.Println(s.Get(4))

}
