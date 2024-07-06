package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var students = make(map[int]string)

	students[1] = "Sai Aravindh"
	students[2] = "Neelakandan"
	students[3] = "JeyaLakshmi"
	students[4] = "Infanta Jose Mary"

	fmt.Println(len(students))

	// keys := make([]int, 0, len(students))
	// for k := range students {
	// 	keys = append(keys, k)
	// }

	// for k := range students {
	// 	delete(students, k)
	// }

	numbers := 6
	res := math.Sqrt(float64(numbers))
	f := math.Round(res)
	fmt.Println(res)
	fmt.Printf("The output is %v\n", f)
	fmt.Printf("The output is %.2g", res)
	if res > 2 {
		fmt.Println("\n\nSuccess")
	} else if res < 2 {
		fmt.Println("\n\nmore skill needed...")
	} else {
		fmt.Println("\n\nFailed...")
	}

	switch f {

	case 1:
		{
			fmt.Println("yes no")

		}
	case 2:
		{
			fmt.Println("Yes")

		}
	case 3:
		{
			fmt.Println("hello")

		}

	}

	today := time.Now().Year()
	fmt.Println(today)

	// var i = 0

	// for i <= len(students) {

	// 	fmt.Println(students[i])
	// 	i++
	// }

}
