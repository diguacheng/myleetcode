package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	// 方法一  20ms 6.3mb
	m := make(map[string]int, 0)
	// 去重
	for _, w := range words {
		if _, ok := m[w]; !ok {
			m[w] = 1
		}
	}
	// 去子串
	for _, w := range words {
		for i := 1; i < len(w); i++ {
			delete(m, w[i:])
		}
	}
	// 计算结果
	ans := 0
	for k := range m {
		ans += len(k) + 1
	}

	return ans
}

func minimumLengthEncoding1(words []string) int {
	// 方法2 72ms 16.2mb
	root := newNode()
	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}
	return findLeaf(root, 0)

}

type node struct {
	istrue bool
	next   []*node
}

func newNode() *node {
	return &node{
		istrue: true,
		next:   make([]*node, 26),
	}
}

func insert(root *node, word string) {
	for i := len(word) - 1; i >= 0; i-- {
		if root.next[word[i]-'a'] == nil {
			root.next[word[i]-'a'] = &node{
				istrue: i == 0,
				next:   make([]*node, 26),
			}
		}
		root = root.next[word[i]-'a']
	}
}

func findLeaf(root *node, length int) int {
	var (
		isleaf = true
		sum    = 0
	)
	nodes := root.next
	for i := 0; i < 26; i++ {
		if nodes[i] != nil {
			isleaf = false
			sum = sum + findLeaf(nodes[i], length+1)
		}
	}
	if isleaf {
		return length + 1
	}
	return sum
}

func minimumLengthEncoding2(words []string) int {
	// 28ms 6.1mb
	if words == nil || len(words) == 0 {
		return 0
	}
	for i, word := range words {
		words[i] = reservedWord(word)
	}
	fmt.Println(words)
	sort.Strings(words)
	fmt.Println(words)
	words = append(words, "*")
	res := 0
	for i := 0; i < len(words)-1; i++ {
		if strings.HasPrefix(words[i+1], words[i]) {
			continue
		}
		res += len(words[i]) + 1

	}
	return res

}

func reservedWord(word string) string {
	bytes := []byte(word)
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)

}

func main() {
	words := []string{"time", "atime", "btime"}
	fmt.Println(minimumLengthEncoding2(words))

}
