package main

import (
	"errors"
	"fmt"
)

type SortBy []Type

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	// age := -10

	// error := fmt.Errorf("Age Cannot Be Negative")

	// if age < 0 {
	// 	fmt.Println(error)
	// } else {
	// 	fmt.Println("age is ", age)
	// }
	er := errors.New("Message")
	res, errs := divide(500, 0)
	fmt.Println(res)
	fmt.Println(errs)

}

func divide(num1, num2 int) (int, error) {

	if num2 == 0 {
		err := fmt.Errorf("Cannot divide by zero")
		return num1, err
	}

	div := (num1 / num2)
	return div, nil
}
