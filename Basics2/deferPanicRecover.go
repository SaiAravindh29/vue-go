package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr []int
	// arr = append(arr, 850, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 5000, 45000, 78500, 1565000)

	sort.Ints(arr)
	fmt.Println(arr)

	// slices.Sort(arr)

	er := printArr(arr)

	for i := range 10 {
		fmt.Println(10 - i)
	}

	if er != nil {
		fmt.Println(er)
	} else {
	}
}

func printArr(arr []int) error {

	if arr != nil {
		for _, i := range arr {
			fmt.Println(i)
		}
		return nil
	} else {
		return fmt.Errorf("Eerrrrrooooooooorrrrrrrrrrrrrrrrrrrrzzzz")
	}

}
