package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	p := make([]int, n+1)
	for i := 0; i <= n; i++ {
		p[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if p[i] != i {
			return find(p[i])
		}
		return i
	}

	var union func(i, j int)
	union = func(i, j int) {
		p[find(i)] = find(j)
	}
	for i := 0; i < n; i++ {
		s, e := edges[i][0], edges[i][1]
		if find(s) != find(e) {
			union(s, e)
		} else {
			return edges[i]
		}
	}
	return nil

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
