package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]bool)
	mapNode := make(map[int]*Node)
	var getNode func(int) *Node
	getNode = func(i int) *Node {
		node, ok := mapNode[i]
		if !ok {
			node = &Node{Val: i, Neighbors: make([]*Node, 0)}
			mapNode[i] = node
		}
		return node
	}

	var BFS func(node *Node) *Node
	BFS = func(node *Node) *Node {
		_, ok := visited[node]
		if !ok {
			newNode := getNode(node.Val)
			for _, n := range node.Neighbors {
				neighbor := getNode(n.Val)
				newNode.Neighbors = append(newNode.Neighbors, neighbor)
			}
			visited[node] = true
			for _, n := range node.Neighbors {
				BFS(n)
			}
			return newNode
		}
		return nil
	}
	res := BFS(node)
	return res
}

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]bool)
	mapNode := make(map[int]*Node)
	var getNode func(int) *Node
	getNode = func(i int) *Node {
		node, ok := mapNode[i]
		if !ok {
			node = &Node{Val: i, Neighbors: make([]*Node, 0)}
			mapNode[i] = node
		}
		return node
	}

	var DFS func(node *Node) *Node
	DFS = func(node *Node) *Node {
		_, ok := visited[node]
		if !ok {
			newNode := getNode(node.Val)
			for _, n := range node.Neighbors {
				neighbor := getNode(n.Val)
				newNode.Neighbors = append(newNode.Neighbors, neighbor)
				visited[node] = true
				DFS(n)
			}

			return newNode
		}
		return nil
	}
	res := DFS(node)
	return res
}
func main() {
	nodes := make([]*Node, 4)
	for i := 0; i < 4; i++ {
		nodes[i] = &Node{Val: i + 1, Neighbors: make([]*Node, 0)}
	}
	nodes[0].Neighbors = append(nodes[0].Neighbors, nodes[1], nodes[3])
	nodes[1].Neighbors = append(nodes[1].Neighbors, nodes[0], nodes[2])
	nodes[2].Neighbors = append(nodes[2].Neighbors, nodes[1], nodes[3])
	nodes[3].Neighbors = append(nodes[3].Neighbors, nodes[0], nodes[2])
	x := cloneGraph1(nodes[0])
	fmt.Println(x)

}
