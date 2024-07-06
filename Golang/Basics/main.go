package main

import "fmt"

func main() {

	companyName := "Qspiders" // also initialize a variable without using datatype.  SYNTAX [ variableName := values ]
	const companyVacancies = 5
	var remainingVacancies uint = 5  // explicitly specifying the variables from int -> uint to avoid getting negative values.
	fmt.Println(&remainingVacancies) // to print the address of the memory
	// we can explicitly declare the datatype like these also
	// var companyName string= "Qspiders"
	// const companyVacancies int = 100
	// var remainingVacancies int = 100

	/*
		unsigned ( positive , whole numbers )
			 uint8 = 0 - 255
			uint16 = 0 - 65535
			uint32 = 0 - 4294967295
			uint64 = 0 - 18446744073709551615

		signed ( whole numbers )
			int8 = byte
			int16 = short
			int32 = int
			int64 = long
	*/

	for {

		fmt.Printf("Company name is %T and Company vacancies is %T and Remaining vacancies is %T\n\n", companyName, companyVacancies, remainingVacancies)
		fmt.Printf("Welcome to %v institute\n", companyName)
		fmt.Printf("We have total of %v vacancies provided and currently %v vacancies available.\n", companyVacancies, remainingVacancies)

		var userName string
		var userVacancy int
		var emailId string
		// asking user for their name and using vacancy.

		fmt.Println("Enter the name of the Candidate : ")
		userVacancy = 1
		fmt.Scan(&userName)
		fmt.Println("Enter the email ID of the candidate : ")
		fmt.Scan(&emailId)
		remainingVacancies -= uint(userVacancy)

		fmt.Printf("User : %v \nemail id: %v \nOccupied : %v vacancy\nRemaining Vacancies : %v\n", userName, emailId, userVacancy, remainingVacancies)

		//! Arrays
		// var trainers [3]string
		// trainers[0] = "Tabrez"
		// trainers[1] = "Bharath"
		// trainers[2] = "Harish"

		// fmt.Printf("The Whole array :%v \n", trainers)
		// fmt.Printf("The first value :%v \n", trainers[0])
		// fmt.Printf("Array Type : %T\n", trainers)
		// fmt.Printf("Array length :%v ", len(trainers))
		//!-------------------------------------------------

		//^ Slice
		// students := []string{}
		// var students = []string{}    alternate ways to create slice

		var students []string
		students = append(students, userName+" "+emailId)

		fmt.Println(students)

		noVacancies := remainingVacancies == 0

		if noVacancies {
			fmt.Println(" No Vacancy \n Exiting....")
			break
		}

	}

	// var num1, num2 int

	// fmt.Println("Enter the first number : ")
	// fmt.Scan(&num1)
	// fmt.Println("Enter the Second number : ")
	// fmt.Scan(&num2)

	// result1, result2 := add(num1, num2)
	// fmt.Printf("The sum of %v + %v is : %v and subtraction of %v - %v is : %v", num1, num2, result1, num1, num2, result2)

}

// func add(num1 int, num2 int) (int, int) {
//  function name ( formal arguements ) ( return type )
//  or

//  func add(a, b int) (resPlus  , resMinus int ){

// resPlus := num1 + num2
// resMinus := num1 - num2
// return

// }

// 	resPlus := num1 + num2
// 	resMinus := num1 - num2
// 	return resPlus, resMinus

// }
