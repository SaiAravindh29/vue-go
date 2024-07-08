package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	log.Println("Main Starts..")

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Millisecond)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	c := make(chan int, 5)
// 	c <- 5
// 	c <- 15
// 	c <- 20
// 	c <- 25
// 	c <- 30
// 	fmt.Println("length : ", len(c))

// 	for i := 0; i < len(c); i++ {
// 		fmt.Println(<-c)
// 		c <- 100

// 	}
// 	close(c)
// 	for value := range c {

// 		fmt.Println(value)

// 	}

// 	// fmt.Println(len(c))

// 	// for
// 	// counting("Hi", c)

// 	// go counting("Hello", c)
// 	// go counting("Hi Hello", c)

// 	// finalResult := <-c

// 	// fmt.Println("Start")

// 	// fmt.Println(finalResult)

// 	// fmt.Println("End")

// 	// time.Sleep(time.Minute)

// 	// var

// }

// func counting(name string, c chan int) {

// 	ct := 1

// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(name, " : ", i)
// 		// time.Sleep(time.Millisecond * 500)
// 		ct++

// 	}

// 	c <- ct

// }
