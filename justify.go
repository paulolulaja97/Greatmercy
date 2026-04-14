package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	option := os.Args[1]
	str := os.Args[2]
	banner := os.Args[3]
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Incomplet")
	}
	words, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("error reading file")
	}
	data := strings.Split(option, "--align=")[1]

	fmt.Println(data)
	terminalWidth := getTermainlWidth()
	// words := "go lang is good"
	// fmt.Println(terminalWidth)
	// fmt.Println(remaningSpaces)
	// fmt.Println(spacesBetweenWords)
	// remander := remaningSpaces % noOfSpaces
	// fmt.Println(remander)
	if data == "right" || data == "left" || data == "center" {
		art := isAscii(words, str)
		splitted_art := strings.Split(art, "\n")
		if data == "right" {
			for _, line := range splitted_art {
				fmt.Printf("%*s", terminalWidth, line)
			}
		} else if data == "left" {
			fmt.Printf("%-*s", terminalWidth, art)

		} else if data == "center" {
			for _, line := range splitted_art {
				fmt.Printf("%*s\n", (terminalWidth-len(line))/2+len(line), line)

			}
		}
	} else if data == "justify" {
		texts := strings.Fields(str)
		totalLenArt := 0
		var rowsArt [][]string

		for _, word := range texts {
			art := isAscii(words, word)
			rows := strings.Split(art, "\n")
			rowsArt = append(rowsArt, rows)
			totalLenArt += len(rows[0])
		}

		noOfSpaces := len(rowsArt) - 1
		remaningSpaces := terminalWidth - totalLenArt
		// splitedWord := strings.Fields(str)
		// noOfSpaces = len(splitedWord) - 1
		var spacesBetweenWords int
		var remainder int
		if noOfSpaces != 0 {
			spacesBetweenWords = remaningSpaces / noOfSpaces
			remainder = remaningSpaces % noOfSpaces
		}

		for i := 0; i < 8; i++ {
			for j, rows := range rowsArt {
				if j == 0 {
					fmt.Print(rows[i])
				} else {
					temp := remainder
					extra := 0
					if temp > 0 {
						extra = 1
					}
					fmt.Printf("%*s", spacesBetweenWords+len(rows[i])+extra, rows[i])
				}
			}
			fmt.Println()
		}
		// for i, word := range splitted_art {
		// 	if i == 0 {
		// 		fmt.Print(word)
		// 	} else if i == len(splitedWord)-1 {
		// 		fmt.Printf("%*s\n", spacesBetweenWords+len(word), word)
		// 	} else {
		// 		fmt.Printf("%*s", spacesBetweenWords+len(word), word)
		// 	}

		// }
	}
}

func getTermainlWidth() int {
	output := exec.Command("tput", "cols")
	output.Stdin = os.Stdin
	out, _ := output.Output()
	i, _ := strconv.Atoi(strings.TrimSuffix(string(out), "\n"))
	return i
}

func isAscii(words []byte, str string) string {

	f := strings.ReplaceAll(string(words), "\r\n", "\n")

	var container strings.Builder
	chrs := strings.Split(f, "\n")
	splited_str := strings.Split(str, "\\n")

	for _, word := range splited_str {
		if word == "" {
			container.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				start_index := int(9*(ch-32) + 1)
				container.WriteString(chrs[start_index+i])
			}
			container.WriteString("\n")
		}
	}

	return container.String()
}
