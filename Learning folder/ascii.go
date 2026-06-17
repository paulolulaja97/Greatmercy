package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fonts, err := os.ReadFile("shadow.txt")
	fonts, err := os.ReadFile("thinkertoy.txt")
	// fonts, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
	}
	// use to deal with the thinkertoy and the \r is carriar return
	f := strings.ReplaceAll(string(fonts), "\r\n", "\n")
	// this where i split the character by new line
	chrs := strings.Split(f, "\n")
	// this in where the argument was passed and and used
	str := os.Args[1]

	//here is where we split our str by escape and new line
	splited_str := strings.Split(str, "\\n")
	// this is where we looped throgh our splited_str and establish that if our word is empty we should just print newline and continue
	for _, word := range splited_str {
		if word == "" {
			fmt.Println()
			continue
		}
		// this where our itration over out ascii code wanting it to loop over 8 times
		for i := 0; i < 8; i++ {
			// this is where we loop over the words we pass in our argument so it can be printed out
			for _, ch := range word {
				// this is where our ascii formular come to play ch is our main ascii and the 32 is the ascii code for space and the +1 is to skip an empty line
				if ch >= 32 && ch <= 126 {
					start_index := int(9*(ch-32) + 1)
					// this is to print our actual output
					fmt.Print(chrs[start_index+i])
				}
			}
			fmt.Println()
		}
	}
}
