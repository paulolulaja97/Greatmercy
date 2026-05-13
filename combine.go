package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 || len(os.Args) > 7 {
		fmt.Println("inaprioprate agrs line")
		return
	}
	color := os.Args[1]
	justify := os.Args[2]
	str := os.Args[3]
	banner := os.Args[len(os.Args)-1]
	bannerList := map[string]bool{"standard": true, "thinkertoy": true, "shadow": true}

	if !bannerList[banner] {
		str = banner

	}
	text, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("can't open file", err)
		return
	}
	substr := ""
	if len(os.Args) == 6 {
		substr = os.Args[len(os.Args)-3]
	}
	j := strings.TrimPrefix(justify, "--justify=")
	c := strings.TrimPrefix(color, "--color=")

	var artC string
	if c != "" {
		artC = findColor(text, c, str, substr)
		if j == "" {
			fmt.Println(artC)
			return
		}
	}
	if c != "" && j != "" {
		isJustify(j, str, "", "", []int{}, text, true, artC)
	}

}

func isJustify(data, str, substr, code string, indexToColour []int, text []byte, js bool, artC string) {
	terminalWidth := getTermainlWidthNew()
	art := ""
	if data == "right" || data == "left" || data == "center" {
		if js {
			art = artC
		} else {
			art = isAsciiArt(text, str, substr, code, indexToColour)
		}
		splitted_art := strings.Split(strings.TrimSuffix(art, "\n"), "\n")
		if data == "right" {
			for _, line := range splitted_art {
				fmt.Printf("%*s\n", terminalWidth, line)
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
			art := isAsciiArt(text, word, substr, code, indexToColour)
			rows := strings.Split(art, "\n")
			rowsArt = append(rowsArt, rows)
			totalLenArt += len(rows[0])
		}

		noOfSpaces := len(rowsArt) - 1
		remaningSpaces := terminalWidth - totalLenArt

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

	}
}

func getTermainlWidthNew() int {
	output := exec.Command("tput", "cols")
	output.Stdin = os.Stdin
	out, _ := output.Output()
	i, _ := strconv.Atoi(strings.TrimSuffix(string(out), "\n"))
	return i
}

func findColor(data []byte, color, str, substr string) string {

	colorCode := map[string]string{
		"red":    "\033[38;5;1m",
		"orange": "\033[38;5;208m",
		"green":  "\033[38;5;2m",
		"blue":   "\033[34m",
	}
	code := colorCode[color]

	indexToColour := []int{}

	for i := 0; i+len(substr) <= len(str); i++ {
		if str[i:i+len(substr)] == substr {
			for j := i; j < i+len(substr); j++ {
				indexToColour = append(indexToColour, j)
			}
		}
	}

	coloredArt := isAsciiArt(data, str, substr, code, indexToColour)
	return coloredArt

}

func isAsciiArt(words []byte, str, substr, code string, indexToColour []int) string {

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
			for j, ch := range word {
				start_index := int(9*(ch-32) + 1)
				if substr == "" || slices.Contains(indexToColour, j) {
					container.WriteString(code + chrs[start_index+i] + "\033[35m")
				} else {
					container.WriteString(chrs[start_index+i])
				}
			}
			container.WriteString("\n")
		}
	}

	return container.String()
}
