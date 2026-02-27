package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		err := os.MkdirAll("logs/app.txt", 0755)
		if err != nil {
			fmt.Println("Error creating directory", err)
			return
		}

	}

	file, err := os.OpenFile(
		"logs/app.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Saved Successfuly\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return

	}
	fmt.Println("message saved successfuly")
}
