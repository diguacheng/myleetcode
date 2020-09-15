package main

import "fmt"

// 深度优先搜索 
func canFinish(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        valid = true
        dfs func(u int)
    )

    dfs = func(u int) {
        visited[u] = 1
        for _, v := range edges[u] {
            if visited[v] == 0 {
                dfs(v)
                if !valid {
                    return
                }
            } else if visited[v] == 1 {
                valid = false
                return
            }
        }
        visited[u] = 2

    }
	n:=len(prerequisites)
    for i:=0;i<n;i++ {

        edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
    }

    for i := 0; i < numCourses && valid; i++ {
        if visited[i] == 0 {
            dfs(i)
        }
	}
    return valid
}

// 广度优先搜索 
func canFinish1(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        indeg = make([]int, numCourses)
        result []int
    )

    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
        indeg[info[0]]++
    }

    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        result = append(result, u)
        for _, v := range edges[u] {
            indeg[v]--
            if indeg[v] == 0 {
                q = append(q, v)
            }
        }
    }
    return len(result) == numCourses
}




func main() {
	prerequisites := [][]int{
		{0,1},
		{1,2},
	
	}
	fmt.Println(canFinish(3,prerequisites))

}