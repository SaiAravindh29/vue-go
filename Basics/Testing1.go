package main

import (
	"fmt"
	"sync"
)

var check int = 0

func main() {

	var w sync.WaitGroup

	ch := make(chan int, 10)

	w.Add(1)

	go getData(ch, &w)

	w.Wait()
	// time.Sleep(time.Second * 2)
	if check == 1 {
		fmt.Println("Done")
		close(ch)
	}

	for message := range ch {
		fmt.Println(message)
	}

	fmt.Println("End")

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
}

func getData(ch chan int, w *sync.WaitGroup) {

	defer w.Done()

	for i := 1; i <= 10; i++ {

		ch <- i
	}

	check = 1

	// close(ch)

}
