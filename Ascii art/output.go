package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// this where i passed the argument in this code that if the argument is less than 3 or grater than4 it should just print incomplet arg and return without doing anythig
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("incomplete")
		return
	}
	// this is where i assign each arg to what i want it to work for passing arg 1 and 2
	option := os.Args[1]
	str := os.Args[2]
	// this is a container created to save our variables so we can later have access to it without going through the stress of doing it over and over again
	var banner string
	// this where a modification is been made incase my argument is 4 that banner should automatically be argumant 3
	if len(os.Args) == 4 {
		banner = os.Args[3]
		//this where i want to note that if the banner string i pass on my terminal dose not have a surffix(.txt) it should just add .txt to it
		if !strings.HasSuffix(banner, ".txt") {
			banner += ".txt"
		}
	} else {
		//here is where i make standard.txt as my default banner incase i did not add banner to the arrgument i passed in the terminal
		banner = "standard.txt"
	}
	// this is where i read my file (banner) and incase we cannot read the file properly we sould just print error and return
	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("error", err)
		return
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
