package main

import "fmt"

func main() {

	//Arrays ,slices and ranges

	fmt.Println("Enter the Quantity of an Bikes you want to store : ")
	var quantity int
	fmt.Scan(&quantity)

	Bikes := [5]string{}

	for index := 0; index < len(Bikes); index++ {
		fmt.Println("Enter the name of the Bike : ")
		fmt.Scan(&Bikes[index])
		fmt.Println()
	}

	newBikes := arrayCheck(Bikes)
	fmt.Println("New Bikes : ", newBikes)

}

func arrayCheck(Bikes [5]string) (copyBikes []string) {

	// copyBikes := []string{}

	for index, value := range Bikes {

		copyBikes = append(copyBikes, value)
		fmt.Println(index, " : ", value)

	}

	return copyBikes
}
