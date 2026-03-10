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

	data := make([]byte, 1000)
	n, err := file.Read(data)
	fmt.Println(string(data[:n]))
}
