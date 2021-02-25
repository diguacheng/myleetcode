package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type pair struct {
		x string
		y string
	}
	m := len(equations)
	table := map[pair]float64{}
	last := make(map[string]string)
	for i := 0; i < m; i++ {
		table[pair{equations[i][0], equations[i][1]}] = values[i]
		last[equations[i][i]] = equations[i][0]
	}
	n := len(queries)
	res := make([]float64, n)
	help := func(x, y string) float64 {
		ans := 0.0
		var mid string
		mid = last[y]
		for mid != "" && mid != x {
			ans += table[pair{mid, y}]
			y = mid
			mid = last[mid]
		}
		if mid == x {
			return ans + table[pair{mid, y}]
		}
		return -1.0

	}

	for i := 0; i < n; i++ {
		x, y := equations[i][0], equations[i][1]
		res[i] = help(x, y)
	}
	return res
}

func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	m := len(values)
	n := len(queries)
	res := make([]float64, n)
	var DFS func(x, y string, isVisited []int) float64
	DFS = func(x, y string, isVisited []int) float64 {
		flag := false
		var a, b string
		for i := 0; i < m; i++ {
			if isVisited[i] == 1 {
				continue
			}
			a, b = equations[i][0], equations[i][1]
			if a == x && b == y {
				isVisited[i] = 1
				flag = true
				return values[i]
			}
			if a == x {
				flag = true
				isVisited[i] = 1
				return DFS(b, y, isVisited) * values[i]
			}
		}
		if flag == false {
			return -1.0
		}
		return -1.0

	}

	for i := 0; i < n; i++ {
		x, y := queries[i][0], queries[i][1]
		res[i] = DFS(x, y, make([]int, m))
	}
	return res
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	fmt.Println(calcEquation1(equations, values, queries))

}
