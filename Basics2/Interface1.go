package main

import "fmt"

type Values interface{}

type SLICES []interface{}

type ARRAY [5]interface{}

func main() {

	var v1 Values

	v1 = "Hi"

	fmt.Println(v1)

	v1 = 100

	fmt.Println(v1)

	var arr = [5]Values{}

	fmt.Println(arr)

	arr[0] = "hi"
	arr[1] = 'a'
	arr[2] = 45
	arr[3] = 15.5
	arr[4] = true
	fmt.Println(arr)

	////////////////////////////////////////////////////////////

	var v2 SLICES
	v2 = append(v2, 15, "Kalai", 45.5, true, false)
	fmt.Println(v2)

	var v3 ARRAY
	v3[0] = "Hello"
	v3[1] = 10
	v3[2] = 15.55
	v3[3] = true
	v3[4] = "Sai"
	// v3[5] = "Sai"
	fmt.Println(v3)

}
