package main

import (
	"fmt"
	"log"
)

func main() {

	log.Println("Main Starts..")

	ub := make(chan int, 5)

	// go func() {
	// receiver := <-ub
	// log.Println("Received:", receiver)
	ub <- 51
	// }()

	rec := <-ub

	fmt.Println(rec)
	// fmt.Println(<-ub)
	// receiver := <-ub
	// close(ub)

	// fmt.Println(receiver)

	// var
	// fmt.Println(receiver)
	// go func() {
	//
	// }()

	// fmt.Println(<-ub)

	// done := make(chan bool, 5)

	// done <- true
	// done <- true
	// done <- true
	// done <- false
	// done <- false
	// // close(done)

	// fmt.Println("length : ", len(done))
	// fmt.Println("capacity : ", cap(done))

	// for j := 0; j < len(done); j++ {
	// 	fmt.Println(<-done)

	// 	if j%2 == 0 {
	// 		done <- true
	// 	} else {
	// 		done <- false
	// 	}

	// }

	// close(done)
	// fmt.Println("length : ", len(done))
	// fmt.Println("capacity : ", cap(done))

	// for i := range done {
	// 	fmt.Println(i)
	// 	// done <- false
	// }

	// go worker(done) // Launch the worker goroutine

	// <-done // Block until receiving a value from the channel
	// fmt.Println("length : ", len(done))
	// fmt.Println("capacity : ", cap(done))

	log.Println("Main ends..")

}

// func worker(done chan bool) {

// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- false // Send a value to notify that the work is done

// }
