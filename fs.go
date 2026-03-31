package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// passing an argumanet in our code for reading our banner
	data := os.Args[1]
	banner := os.Args[2]
	// if the arguent is not up to 3 just print incomplet line of argument
	if len(os.Args) != 3 {
		fmt.Println("incomplet line of argument")
	}
	// this is where we read our file
	file, err := os.ReadFile(banner)
	// if the file is not opening properly just pirnt error reading file
	if err != nil {
		fmt.Println("error reading file")
	}

	t := strings.ReplaceAll(string(file), "\r\n", "\n")
	f := strings.Split(t, "\n")
	if data == "" {
		return
	}
	d := strings.Split(data, "\\n")
	// this is where we loop through our work and split it by newline
	for _, word := range d {
		if word == "" {
			fmt.Println()
			continue
		}
		//
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				start_index := int(9*(ch-32) + 1)
				fmt.Print(f[start_index+i])
			}
			fmt.Println()
		}
	}

}
