package main

import "fmt"

func main() {

	fmt.Println("Hello from file 2")

	// Maps

	// An Empty map
	// map[Key_Type]Value_Type{}

	// Map with key-value pair
	// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

	/* Map Declatation types */

	var a = map[int]string{}
	fmt.Println(a)
	var b = make(map[int]string)
	fmt.Println(b)

	c := map[int]string{}
	d := make(map[int]string)
	fmt.Println(c)
	fmt.Println(d)

	/**************************************************/

	collection := map[string]int{
		"One": 1, "Two": 2, "Three": 3}

	fmt.Println(collection)

	map1 := map[int]int{}

	fmt.Println("Map 1 : ", map1)

	map2 := map[int]string{

		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
	}

	// modification on maps

	fmt.Println("Map 2 : ", map2)

	map3 := map2

	fmt.Println("Map 3 : ", map3)

	fmt.Println("************************************************")

	map3[100] = "Six"
	map3[200] = "Seven"
	fmt.Println("Map 3 updated : ", map3)
	fmt.Println("Map 2 updated : ", map2)
	map3[100] = "Hundred"

	fmt.Println("************************************************")
	map2[8] = "Eight"
	fmt.Println("Map 3 updated 1: ", map3)
	fmt.Println("Map 2 updated 1: ", map2)
	fmt.Printf("map 3 : %p and map 2 : %p", map3, map2)

	var map4 = map[int]string{}
	fmt.Println("Map 4 : ", map4)

	map4 = map3
	fmt.Println("Map 4 : ", map4)
	map4[9] = "Nine"
	fmt.Println("Map 4 updated : ", map4)
	fmt.Println("Map 3 updated : ", map3)
	fmt.Println("Map 2 updated : ", map2)

	map5 := make(map[int]string)
	// copy(map5, map4)
	fmt.Println(map5)

}
