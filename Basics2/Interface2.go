package main

import "fmt"

func main() {

	var user UserDetail
	fmt.Println("Enter the first name : ")
	fmt.Scanln(&user.firstName)

	fmt.Println("Enter the Last name : ")
	fmt.Scanln(&user.lastName)

	fmt.Println("Enter the age: ")
	fmt.Scanln(&user.age)

	process(user)

}

func (user UserDetail) mergeName() string {
	return user.firstName + " " + user.lastName
}

func process(v Modify) {
	fmt.Println(v.mergeName())
}

type Modify interface {
	mergeName() string
}

type UserDetail struct {
	firstName string
	lastName  string
	age       uint
}
