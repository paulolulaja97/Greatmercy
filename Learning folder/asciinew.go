package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error")
	}

	pass := os.Args[1]
	p := strings.Split(pass, "\\n")
	t := strings.ReplaceAll(string(text), "\r\n", "\n")
	f := strings.Split(t, "\n")

	for _, words := range p {
		if string(words) == "" {
			fmt.Println()
		}
		for i := 0; i < 8; i++ {
			for _, g := range words {
				ch := int(9*(g-32) + 1)
				fmt.Print(f[ch+i])
			}
			fmt.Println()
		}

	}

}
