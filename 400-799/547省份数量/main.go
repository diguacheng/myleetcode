package main

func findCircleNum(isConnected [][]int) int {
	m := len(isConnected)
	isVisited := make([][]int, m)
	for i := 0; i < m; i++ {
		isVisited[i] = make([]int, m)
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if isVisited[i][j] == 0 && isConnected[i][j] == 1 {
				isVisited[i][j] = 1
				isVisited[j][i] = 1
				cnt++
				DFS(isVisited, isConnected, i, j, m)
			}
		}
	}
	return cnt
}

func DFS(isVisited, isConnected [][]int, i, j, m int) {
	for y := 0; y < m; y++ {
		if isVisited[j][y] == 0 && isConnected[j][y] == 1 {
			isVisited[j][y] = 1
			isVisited[y][j] = 1
			DFS(isVisited, isConnected, j, y, m)
		}
	}
	for y := j + 1; y < m; y++ {
		if isVisited[i][y] == 0 && isConnected[i][y] == 1 {
			isVisited[i][y] = 1
			isVisited[y][i] = 1
			DFS(isVisited, isConnected, i, y, m)
		}
	}
}

func findCircleNum1(isConnected [][]int) int {
	var ans int
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return ans

}

func main() {

}
