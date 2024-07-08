package main

import "fmt"

func main() {

	a := make(map[int]string)
	fmt.Println("a : ", a)

	a[1] = "a"
	a[2] = "b"
	a[3] = "c"
	a[4] = "d"
	a[5] = "e"
	fmt.Println("a : ", a)
	fmt.Println(a[8])

	value1, key1 := a[5]
	value2, key2 := a[6]
	fmt.Println("value1 ", value1)
	fmt.Println("key1 ", key1)
	fmt.Println("value2 ", value2)
	fmt.Println("key2 ", key2)

}
