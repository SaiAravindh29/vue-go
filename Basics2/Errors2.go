package main

import "fmt"

func main() {

	ch := make(chan int)

	ch <- 10

	ch2 := <-ch

	fmt.Println(ch2)
}

// 	var user Details

// 	fmt.Println("Enter a name : ")
// 	fmt.Scan(&user[0])

// 	fmt.Println("Enter age : ")
// 	fmt.Scan(&user[1])

// 	checkDetails(user)

// }

// func checkDetails(user Details) error {

// 	for i := 0; i < len(user); i++ {
// 		if i == 0 {
// 			if user[i] < 0 || user[i] >= 0 {

// 				err := fmt.Errorf("The User Name is invalid")
// 				return err
// 			} else {
// 				fmt.Println("The Name of the User is : ", user[i])
// 			}
// 		} else if i == 1 {
// 			if user[i] <= 0 {

// 				err := fmt.Errorf("The Age is invalid")
// 				return err

// 			} else {
// 				fmt.Println("The Age of the User is : ", user[i])
// 			}
// 		}
// 	}

// }

// type Details [2]interface{}

/****************************************/
