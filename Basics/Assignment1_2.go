package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

type Content struct {
	num1, num2, Add, Sub, Mul int
}

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex

	fmt.Println("Hello")

	// for index := 1; index <= 10; index++ {

	// 	wg.Add(1)

	// 	go func(index int) {

	// 		m.Lock()

	// 		fmt.Println("S.No : ", index)
	mainRun(&m, &wg)
	// 		fmt.Println("\n************************\n")

	// 		m.Unlock()

	// 	}(index)

	// }

	// wg.Wait()

}

func generateNumber(user chan int) {

	// for _, i := range user {
	// 	i <- rand.Intn(550-50) + 50
	// }
	user <- rand.Intn(550-50) + 50
	user <- rand.Intn(350-25) + 25
	close(user)
}

func mainRun(m *sync.Mutex, wg *sync.WaitGroup) {

	user := make(chan int, 2)

	generateNumber(user)

	resChan := make(chan Content, 5)

	Addition(user, resChan)

	// // var use Content
	// // use.num1 = <-user
	// // use.num2 = <-user

	// // resultChan := make(chan int, 3)

	// // go Addition(resultChan, use.num1, use.num2)

	// log.Println("Number 1 : ", use.num1)
	// log.Println("Number 2 : ", use.num2)
	// log.Println("Addition : ", use.Add)
	// log.Println("Subtraction : ", use.Sub)
	// log.Println("Multiplication : ", use.Mul)

	// wg.Done()

}

func Addition(user chan int, resChan chan Content) {
	var Values Content
	Values.num1 = <-user
	Values.num2 = <-user
	Values.Add = Values.num1 + Values.num2
	log.Println(Values.Add)

}

// func Subraction(resultChan chan int, num1, num2 int) {
// 	minus := num1 - num2
// 	resultChan <- minus
// 	Multiplication(resultChan, num1, num2)
// }

// func Multiplication(resultChan chan int, num1, num2 int) {
// 	multiply := num1 * num2
// 	resultChan <- multiply
// 	close(resultChan)

// }
