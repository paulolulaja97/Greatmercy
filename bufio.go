package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	file, err := os.OpenFile("logs/app.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		fmt.Println("Failed to open fill:", err)
		return
	}

	_, err = file.WriteString("Hello!\n")
	if err != nil {
		fmt.Println("error writing:", err)
		file.Close()
		return
	}
	file.Close()

	file, err = os.Open("logs/app.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

}
