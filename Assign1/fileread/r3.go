package fileread

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Read2() {

	file, err := os.Open("file2.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, er1 := reader.ReadAll()

	if er1 != nil {
		fmt.Println(er1)
		return
	}
	fmt.Println(records)

	// currentTime := time.Now()
	// col_time := currentTime.Format("02-01-2006 15")

	// for r := range records {
	// 	if r == 0 {
	// 		records[r] = append(records[r], "Date & Time")
	// 	} else {
	// 		records[r] = append(records[r], col_time)
	// 	}

	// }

}
