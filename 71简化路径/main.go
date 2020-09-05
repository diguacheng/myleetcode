package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) == 0 {
		return "/"
	}

	SplitedPath := strings.Split(path, "/")
	res := make([]string, 0)

	for _, v := range SplitedPath {
		if v == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			continue
		}
		if v == "." || v == "" {
			continue
		}
		res = append(res, v)

	}
	return "/" + strings.Join(res, "/")
}



func simplifyPath1(path string) string {
	return filepath.Clean(path)
}



func main() {
	ss := "/../"
	fmt.Println(simplifyPath(ss))

}
