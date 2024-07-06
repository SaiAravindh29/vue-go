package main

import (
	"fmt"
	"slices"
)

func main() {

	arr := []int{}
	arr = append(arr, 10, 50, 40, 80, 90, 100, 5, 20, 1, 8, 4, 6)
	fmt.Println(arr)
	slices.Sort(arr)
	fmt.Println(arr)
	a := add(5, 10)
	fmt.Println(a)
}

func add(a, b int) int {
	return a + b
}
