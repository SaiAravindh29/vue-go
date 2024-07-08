package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	number1 := 10
	number2 := 50

	var v2 float32 = 14.7
	fmt.Print(v2)
	fmt.Println(v2)
	// fmt.Printf("the value is %.2d", v2)
	// var addition int = number1 + number2
	// fmt.Println(addition)
	fmt.Println(addition(number1, number2)) // adding numbers with the help of function
	fmt.Println("\n")

	numberValue := 10    //  := can be used without entering the datatype
	stringValue := "Sai" // no need for using  var, const

	// without using fmt package
	println(numberValue, stringValue)
	print(numberValue, stringValue)

	var value1 float32 = 15.51
	fmt.Println(value1)
	println(value1)

	// declaration of arrays
	array1 := [15]int{}
	// array1[]

	for i := 0; i < len(array1); i++ {
		fmt.Println("Enter the value : ")
		fmt.Scan(&array1[i])
	}
	// fmt.Println(array1)

	var array2 = [10]int{10, 30, 50, 90, 20, 60, 80, 100, 40, 70}
	// array1 = append(array1, array2)

	var arr = [15]int{}
	fmt.Println("arr", arr)

	fmt.Println("array 1 : ", array1)
	fmt.Println(array2)

	var hello string = "hello"
	hello2 := "hello"

	fmt.Println(hello)
	fmt.Println(hello2)

	/*************************************************************************************/

	//slices

	slice1 := []int{}
	fmt.Println(slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50) // adding delements
	fmt.Println(slice1)
	fmt.Println("\n\n")

	content1 := []int{}
	fmt.Println("Content 1 : ", content1)

	//adding 10 elements
	content1 = append(content1, 10, 20, 30, 40, 50, 6, 7, 8, 9, 10)
	fmt.Println("Content 1 : ", content1)

	// Holding 5 elements from the start and remove other elements
	content1 = content1[:5]
	fmt.Println("Content 1 removed : ", content1)

	/// removing 2 elements from the start
	content1 = content1[2:]
	fmt.Println("Content 1 : ", content1)

	/*************************************************************************************/

	/* sorting array */

	var temporary = 0

	fmt.Println("Brfore Sorted :", array2)

	for index := 0; index < len(array2)-1; index++ {
		// fmt.Println(array2[index])
		for index2 := 0; index2 < len(array2)-1; index2++ {

			if array2[index2] > array2[index2+1] {

				temporary = array2[index2]
				array2[index2] = array2[index2+1]
				array2[index2+1] = temporary
			}

		}
	}

	fmt.Println("After Sorted :", array2)
	fmt.Println("\n\n")

	/*************************************************************************************/

	/* duplicate identify */

	duplicates := []int{5, 10, 120, 15, 15, 20, 5, 30, 45, 85, 75, 20, 120}

	fmt.Println("Slice having duplicate elements : ", duplicates)
	fmt.Print("The duplicates are : ")
	odd := 0
	even := 0
	oddValues := []int{}
	evenValues := []int{}

	for index3 := 0; index3 < len(duplicates); index3++ {

		for index4 := 0; index4 < len(duplicates); index4++ {

			if duplicates[index3] == duplicates[index4] && index3 != index4 && duplicates[index4] != -5 {

				if duplicates[index4]%2 == 0 {
					fmt.Print("Even : ")
					even++
					evenValues = append(evenValues, duplicates[index4])

				} else {
					fmt.Print("Odd : ")
					odd++
					oddValues = append(oddValues, duplicates[index4])

				}

				fmt.Print(duplicates[index4], " , ")
				duplicates[index4] = -5

			}
		}

	}

	fmt.Println("Duplicate Odd Number count = ", odd)
	fmt.Println("Odd Elements are : ", oddValues)
	fmt.Println("Duplicate Even Number count = ", even)
	fmt.Println("Even Elements are : ", evenValues)

	fmt.Println("\n\n")

	/*************************************************************************************/

	//Struct

	type UserDetails struct {
		userName string
		age      uint
		role     string
	}

	// var a userDetails

	var user1 UserDetails
	var user2 UserDetails

	user1.userName = "sai"
	user1.age = 24
	user1.role = "Developer"

	user2.userName = "Aravindh"
	user2.age = 23
	user2.role = "Developer"

	fmt.Println(user1)
	fmt.Println(user1.userName)
	fmt.Println(user1.age)
	fmt.Println(user1.role)

	fmt.Println(user2)
	fmt.Println(user2.userName)
	fmt.Println(user2.age)
	fmt.Println(user2.role)

	/****************************************************************************************/

}

func addition(num1 int, num2 int) int {
	var addResult = num1 + num2
	return addResult
}

func getDetails(userName string, age uint64, salary float64) {

	fmt.Printf(" The Name of the User is %s and the Age is %d and his salary is %g", userName, age, salary)

}

func userInput() {

	var userName string
	var age uint64
	var salary float64

	fmt.Println("Enter the Name : ")
	fmt.Scan(&userName)
	fmt.Println("Enter the Age : ")
	fmt.Scan(&age)
	fmt.Println("Enter the Salary : ")
	fmt.Scan(&salary)

	getDetails(userName, age, salary)

}

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
