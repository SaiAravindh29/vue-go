package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	var m sync.Mutex

	var arr1 = [5]string{"Sai", "Aravindh", "Ezhil", "Dakshin", "Eniyan"}
	var arr2 = [5]string{"Developer", "Developer", "Geologist", "Govt", "Bank Emp"}

	name := make(chan string, 1)
	role := make(chan string, 1)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		name <- arr1[i]
		role <- arr2[i]

		go details(name, role, &m, &wg)

	}

	close(name)
	close(role)
	wg.Wait()

}

func details(name chan string, role chan string, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	m.Lock()

	fmt.Println("User Name = ", <-name)
	time.Sleep(time.Second * 2)
	fmt.Println("User Role = ", <-role)
	time.Sleep(time.Second * 2)
	fmt.Println("************************\n")

	m.Unlock()

}
