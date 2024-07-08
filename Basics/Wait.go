package main

import (
	"fmt"
	"sync"
	"time"
)

// func worker(id int) {
// 	fmt.Printf("Worker %d starting\n", id)

// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d done\n", id)
// }

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)
		fmt.Println("MT : ", i)
		// fmt.Println("Add : ", wg)

		go func() {
			defer wg.Done()
			// fmt.Println("Done : ", wg)
			fmt.Println("ST : ", i)

		}()

		if i < 5 {
			time.Sleep(time.Second)
		} else {
			time.Sleep(time.Millisecond * 0)

		}

	}

	wg.Wait()

}
