package main

import "fmt"

// lc
func f(a, b string) bool {
	if a == b {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	count := make([]int, 26)
	for _, v := range a {
		count[v-'a']++
	}
	for _, v := range b {
		count[v-'a']--
		if count[v-'a'] < 0 {
			return false
		}
	}
	return true
}
type LRUCache struct {
	size       int
	capacity   int
	head, tail *DLinknode
	cache      map[int]*DLinknode
}

type DLinknode struct {
	Key, Value int
	Prev, Next *DLinknode
}

func newNode(key, value int) *DLinknode {
	return &DLinknode{key, value, nil, nil}
}

func NewLRUCache(c int) *LRUCache {
	L := &LRUCache{size: 0, capacity: c, head: newNode(0, 0), tail: newNode(0, 0), cache: make(map[int]*DLinknode)}
	L.head.Next=L.tail
	L.tail.Prev=L.head
	return L
}

func (this *LRUCache)Get(key int)int{
	if node,ok:=this.cache[key];ok{
		this.movetohead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache)Set(key int, value int){
	if node,ok := this.cache[key];ok{
		node.Value=value
		this.movetohead(node)
	}else{
		node:=newNode(key,value)
		this.cache[key]=node
		this.size++
		if this.size>this.capacity{
			removed:=this.removetail()
			delete(this.cache,removed.Key) // 不能只在链表里 删除 
		}
		this.addtohead(node)
	}
}

func (this *LRUCache)movetohead(node *DLinknode){
	this.removenode(node)
	this.addtohead(node)
}

func (this *LRUCache)removenode(node *DLinknode){
	node.Prev.Next=node.Next
	node.Next.Prev=node.Prev
}

func (this *LRUCache)addtohead(node *DLinknode){
	node.Next=this.head.Next
	this.head.Next.Prev=node 
	this.head.Next=node 
	node.Prev=this.head
}

func (this *LRUCache)removetail()*DLinknode{
	node:=this.tail.Prev
	this.removenode(node)
	return node
}

func main() {
	lru:=NewLRUCache(3)
	lru.Set(1,10)
	lru.Set(2,200)
	lru.Set(3,11)
	lru.Set(4,2)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(1))


}
