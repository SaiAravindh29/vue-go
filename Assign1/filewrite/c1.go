package filewrite

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateFile(file string) {

	Nfile, err := os.Create("file2.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File Created Successfully ")
	defer Nfile.Close()

	write := csv.NewWriter(Nfile)

	f1 := []string{file}

	// currentTime := time.Now()
	// col_time := currentTime.Format("02-01-2006 15")

	// for r := range f1 {
	// 	if r == 0 {
	// 		f1 = append(f1, "Date & Time")
	// 	} else {
	// 		f1 = append(f1, col_time)
	// 	}

	// }

	write.WriteAll([][]string{f1})

}
