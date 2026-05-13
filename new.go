package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("stndard.txt")
	// if err != nil {
	// 	fmt.Println("error reading file")

	// }

	splitedData := strings.Split(string(data), "\n")

	if len(splitedData) == 856 {
		for _, i := range splitedData {
			for _, j := range i {

				if j < 32 || j > 126 {

					fmt.Println("invalid")
					return
				}
			}
		}
		fmt.Println("VALID")

	} else {
		fmt.Println("INVALID")
	}

	//fmt.Println(len(splitedData))

	// cont := 0
	// for i, _ := range splitedData {
	// 	cont = i
	// }
	// fmt.Println(cont)

}
