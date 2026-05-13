package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	font, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error reading file")

	}
	if len(os.Args) < 2 {
		fmt.Println("incomplet args")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("too many args")
		return
	}
	str := os.Args[1]
	f := strings.ReplaceAll(string(font), "\r", "")
	g := strings.Split(f, "\n")

	splited := strings.Split(str, "\\n")

	for _, words := range splited {
		if words == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {

			for _, chrs := range words {
				if chrs >= 32 && chrs <= 126 {
					in := int(9*(chrs-32) + 1)
					fmt.Print(g[in+i])
				}
			}
			fmt.Println()
		}

	}
}
