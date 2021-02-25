package main

import "fmt"

import "reflect"

func isValid(s string) bool {
	fmt.Println("hello world")
	return true

}

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(x)

}
