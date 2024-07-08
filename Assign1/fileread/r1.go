package fileread

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func Read() string {

	filePath := "file1.csv"

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Failed to read file")
		return ""
	}

	encodeData := base64.StdEncoding.EncodeToString(data)
	return encodeData

}
