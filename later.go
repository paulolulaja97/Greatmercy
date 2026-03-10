package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	data := make([]byte, 1)
	var con []byte

	no_byte := 0

	for {
		n, err := file.Read(data)
		con = append(con, data...)
		if err != nil {
			break
		}
		no_byte += n
	}

	fmt.Println(string(con[:no_byte]))
}
