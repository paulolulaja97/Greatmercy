package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// if the arguent is not up to 3 just print incomplet line of argument
	if len(os.Args) != 3 {
		fmt.Println("incomplet line of argument")
		return
	}
	// passing an argumanet in our code banner and data
	data := os.Args[1]
	banner := os.Args[2]
	// this is where we read our file and our os.readfile returns two things the file itself and error
	file, err := os.ReadFile(banner + ".txt")
	// if the file is not opening properly just pirnt error reading file
	if err != nil {
		fmt.Println("error reading file")
	}

	t := strings.ReplaceAll(string(file), "\r", "")
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
		// this deals with the ascii art looping it until the index is equal to 8
		for i := 0; i < 8; i++ {
			// this is the inner loop that deals with the rang of the words you want to print from the ascii table
			for _, ch := range word {
				// the CH is the word you want to print from the ascii table and the 32 is the ascii nember of space so the ascii pick your character and start it
				// from space and +1 → skips the separator/empty line at the top
				start_index := int(9*(ch-32) + 1)
				fmt.Print(f[start_index+i])
			}
			fmt.Println()
		}
	}

}
