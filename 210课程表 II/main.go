package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 本质是广度优先搜索 
	G := make([][]int, numCourses)// G[i]表示从i进入的点
	for _, pre := range prerequisites {
		G[pre[0]] = append(G[pre[0]], pre[1])

	}

	res := []int{}
	count := make([]int, numCourses)
	var flag int
	for len(res) < numCourses {
		flag = -1
		for i := 0; i < numCourses; i++ {
			if len(G[i]) == 0 && count[i] == 0 {
				// 如果某个点入度为0，且未访问过。
				flag = i
				res = append(res, i)
				count[i] = 1
				break
			}
		}
		if flag > -1 {
			// 将flag点删除，即删除图中以flag为起点的线段
			for i := 0; i < numCourses; i++ {
				for j := 0; j < len(G[i]); j++ {
					if G[i][j] == flag {
						G[i] = append(G[i][:j], G[i][j+1:]...)
						continue
					}
				}
			}
		} else {
			return []int{}
		}

	}
	return res

}

func findOrder1(numCourses int, prerequisites [][]int) []int {
	// 深度优先搜素
	res := []int{}
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	flag := false

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if flag {
					return
				}
			} else if visited[v] == 1 {
				flag = true
				return
			}
		}
		visited[u] = 2
		res = append(res, u)
	}
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
	}
	for i := 0; i < numCourses && !flag; i++ {
		if visited[i] == 0 {
			dfs(i)
		}

	}
	if flag {
		return []int{}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
	}
	return res

}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	// 深度优先搜索
	res := []int{}
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	flag := false

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if flag {
					return
				}
			} else if visited[v] == 1 {
				flag = true
				return
			}
		}
		visited[u] = 2
		res = append([]int{u}, res...)
	}
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
	}
	for i := 0; i < numCourses && !flag; i++ {
		if visited[i] == 0 {
			dfs(i)
		}

	}
	if flag {
		return []int{}
	}

	return res

}


func findOrder3(numCourses int, prerequisites [][]int) []int {
	// 广度优先搜索
	res := []int{}
	edges := make([][]int, numCourses)
	indeg:= make([]int, numCourses)
	
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		indeg[pre[0]]++
	}
	q:=[]int{}
	for i:=0;i<numCourses;i++ {
		if indeg[i]==0{
			q=append(q,i)
		}
	}
	for len(q)>0{
		u:=q[0]
		q=q[1:]
		res=append(res,u)
		for _,v:=range edges[u]{
			indeg[v]--
			if indeg[v]==0{
				q=append(q,v)
			}
		}
	}
	if len(res)!=numCourses{
		return []int{}
	}
	return res

	

}
func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(findOrder3(numCourses, prerequisites))

}
