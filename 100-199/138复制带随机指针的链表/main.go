package main



type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		temp := &Node{curr.Val, nil, nil}
		m[curr] = temp
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		m[curr].Next = m[curr.Next]
		m[curr].Random = m[curr.Random]
	}
	return m[head]

}

func copyRandomList(head *Node) *Node {
	// 回溯法 深度优先遍历 用哈希map 存储访问过的点
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	return help(head, m)

}

func help(head *Node, m map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if x, ok := m[head]; ok {
		return x
	}
	temp := &Node{head.Val, nil, nil}
	m[head] = temp
	temp.Next = help(head.Next, m)
	temp.Random = help(head.Random, m)
	return temp

}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	oldnode := head
	newnode := &Node{oldnode.Val, nil, nil}
	m[oldnode] = newnode
	for oldnode != nil {
		newnode.Random = getCloneNode(oldnode.Random, m)
		newnode.Next = getCloneNode(oldnode.Next, m)
		oldnode = oldnode.Next
		newnode = newnode.Next
	}
	return m[head]
}

func getCloneNode(head *Node, m map[*Node]*Node) *Node {
	if head != nil {
		if node, ok := m[head]; ok {
			return node
		}
		m[head] = &Node{head.Val, nil, nil}
		return m[head]
	}
	return nil
}

func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	ptr := head
	for ptr != nil {
		newnode := &Node{ptr.Val, nil, nil}
		newnode.Next = ptr.Next
		ptr.Next = newnode
		ptr = newnode.Next
	}
	ptr = head
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		}
		ptr = ptr.Next.Next
	}
	oldlist := head
	newlist := head.Next
	headold := head.Next
	for oldlist != nil {
		oldlist.Next = oldlist.Next.Next
		if newlist.Next != nil {
			newlist.Next = newlist.Next.Next
		} else {
			newlist.Next = nil
		}
		oldlist = oldlist.Next
		newlist = newlist.Next
	}
	return headold
}

func main() {

}
