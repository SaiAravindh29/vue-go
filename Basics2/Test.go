package main

import (
	"fmt"
)

func main() {

	arr := []int{}

	tempres, ef := giveData(arr)

	if ef != nil {

		fmt.Print(ef)
		fmt.Println("Exiting....\n")

	} else {
		res, e := multiplybysame(tempres)

		if e != nil {
			fmt.Println(e)
		} else {

			for v := range res {
				fmt.Println(res[v])
			}
			fmt.Println("Done")
		}

	}

	// var c byte = 'N'
	// fmt.Println(c)

}

func giveData(arr []int) ([]int, error) {

	er2 := fmt.Errorf("Invalid data")

	fmt.Println("Do you want to enter some values ? (y/n)")
	var cc string
	fmt.Scan(&cc)

	if cc == "y" || cc == "Y" {

		var length int
		fmt.Println("Enter the no of data you want to store : ")
		fmt.Scan(&length)

		for i := 0; i < length; i++ {

			var tempData int

			fmt.Println("Enter a number : ")
			fmt.Scan(&tempData)

			arr = append(arr, tempData)

		}

	} else if cc == "n" || cc == "N" {
		fmt.Println("Done...")
		return arr, nil

	} else {
		return arr, er2
	}

	return arr, nil

}

func multiplybysame(arr []int) ([]int, error) {

	var er1 = fmt.Errorf("Slice is empty")

	if len(arr) == 0 {
		return arr, er1
	} else {
		for i := range arr {
			arr[i] = arr[i] * arr[i]
		}
		return arr, nil
	}

}
