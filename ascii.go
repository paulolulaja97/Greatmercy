package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fonts, err := os.ReadFile("shadow.txt")
	// fonts, err := os.ReadFile("thinkertoy.txt")
	// fonts, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
	}
	f := strings.ReplaceAll(string(fonts), "\r\n", "\n")

	chrs := strings.Split(f, "\n")
	str := os.Args[1]
	splited_str := strings.Split(str, "\\n")

	for _, word := range splited_str {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				start_index := int(9*(ch-32) + 1)
				fmt.Print(chrs[start_index+i])
			}
			fmt.Println()
		}
	}
}
