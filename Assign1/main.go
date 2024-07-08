package main

import (
	fr "Assign1/fileread"
	fw "Assign1/filewrite"
	"fmt"
	"time"
	// "fmt"
	// "io/ioutil"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)
	dmy := currentTime.Format("02-01-2006")
	t1 := currentTime.Format("15:04:05")
	fmt.Println(dmy + " " + t1)
	fmt.Println("Main File")

	d1 := fr.Read()
	fw.CreateFile(d1)
	fr.Read2()

	// filePath := "file2.csv"
	// data, err := ioutil.ReadFile(filePath)

	// if err != nil {
	// 	fmt.Println("Failed to read file")
	// }

	// encoded := string(data)

	// fr.Decode(encoded)

	// result := fr.Read2()

	// fr.Decode(result)

	// res1 := string(result)

	// fr.Decode(res1)

}
