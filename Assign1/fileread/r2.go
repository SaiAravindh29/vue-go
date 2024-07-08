package fileread

import (
	"encoding/base64"
	"fmt"
)

func Decode(encodeData string) {

	decodedData, err := base64.StdEncoding.DecodeString(encodeData)

	if err != nil {
		fmt.Println("Failed to decode base64:", err)
		return
	}

	fmt.Println("Decoded CSV data:")
	fmt.Println(string(decodedData))

}
