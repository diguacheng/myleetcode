package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	i := 0
	for i < l-1 {
		if flowerbed[i] == 0 {
			if flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
				if n == 0 {
					return true
				}
				i += 2
			} else {
				i += 3
			}

		} else {
			if flowerbed[i+1] == 0 {
				i += 2
			} else {
				i += 3
			}
		}
	}
	if i == l-1 {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			n--
		}
	}
	return n <= 0

}

// func backtack(arr []int,n int)bool{
//     l:=len(arr)
// }

func main() {
	arr := []int{1, 0, 0, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(arr, 2))

}
