package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("incomplete")
		return
	}
	option := os.Args[1]
	str := os.Args[2]

	var banner string

	if len(os.Args) == 4 {
		banner = os.Args[3]
		if !strings.HasSuffix(banner, ".txt") {
			banner += ".txt"
		}
	} else {
		banner = "standard.txt"
	}

	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("error", err)
	}
	// --output=res.txt
	file := strings.Split(option, "--output=")[1]

	fmt.Println(file)
	f := strings.ReplaceAll(string(data), "\r\n", "\n")
	t := strings.Split(f, "\n")
	g := strings.Split(str, "\n")

	// var art strings.Builder
	var art []byte

	for _, words := range g {
		if words == "" {
			return
		}

		for i := 0; i < 8; i++ {
			for _, ch := range words {
				start_index := int(9*(ch-32) + 1)
				art = append(art, t[start_index+i]...)
			}
			art = append(art, byte('\n'))
		}
	}

	os.WriteFile(file, art, 0664)
}
