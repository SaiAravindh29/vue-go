package main

import (
	"fmt"
	"strings"
)

func main() {

	var s, t string

	s = "Apple"
	t = "Orange"
	// var a = [10]int{1, 2, 3, 5, 6, 4, 7, 8, 9, 10}
	fmt.Println(s, t)
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(t))
	fmt.Println(strings.HasPrefix(s, "A"))
	fmt.Println(strings.HasSuffix(t, "a"))
	fmt.Println(strings.Contains(s, "pa"))
	// fmt.Println(math.Max(9199999, 5))
}
