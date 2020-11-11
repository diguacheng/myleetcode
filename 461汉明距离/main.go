package main

import "fmt"

func hammingDistance(x int, y int) int {
	c:=0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			c++
		}
		x = x >> 1
		y = y >> 1
	}
	return c

}

func main() {
	fmt.Println(hammingDistance(1,4))


}