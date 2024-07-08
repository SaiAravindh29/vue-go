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
	var sno int = 1

	for index := 1; index <= 10; index++ {

		wg.Add(1)

		go func(index int) {

			m.Lock()

			fmt.Println("S.No : ", (sno))
			mainRun(&m, &wg)
			fmt.Println("\n************************\n")
			sno++
			m.Unlock()

		}(index)

	}

	wg.Wait()

}

func generateNumber(user chan int) {
	user <- rand.Intn(550-50) + 50
	user <- rand.Intn(350-25) + 25
	// close(user)
}

func mainRun(m *sync.Mutex, wg *sync.WaitGroup) {

	user := make(chan int, 2)

	generateNumber(user)

	var use Content
	use.num1 = <-user
	use.num2 = <-user

	resultChan := make(chan int, 3)

	go Addition(resultChan, use.num1, use.num2)

	use.Add = <-resultChan
	use.Sub = <-resultChan
	use.Mul = <-resultChan

	log.Println("Number 1 : ", use.num1)
	log.Println("Number 2 : ", use.num2)
	log.Println("Addition : ", use.Add)
	log.Println("Subtraction : ", use.Sub)
	log.Println("Multiplication : ", use.Mul)

	wg.Done()

}

func Addition(resultChan chan int, num1, num2 int) {
	sum := num1 + num2
	resultChan <- sum
	Subraction(resultChan, num1, num2, sum)
}

func Subraction(resultChan chan int, num1, num2, sum int) {
	minus := num1 - num2
	resultChan <- minus
	Multiplication(resultChan, num1, num2, sum, minus)
}

func Multiplication(resultChan chan int, num1, num2, sum, minus int) {
	multiply := num1 * num2
	resultChan <- multiply
	close(resultChan)
}
