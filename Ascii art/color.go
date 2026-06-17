package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println("incomplet argument")
		return
	}
	option := os.Args[1]
	substr := ""
	banner := "standard.txt"
	bannerList := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
	str := os.Args[len(os.Args)-1]
	if len(os.Args) == 4 {
		if bannerList[str] {
			banner = str
			str = os.Args[len(os.Args)-2]
		} else {
			substr = os.Args[2]
		}

	} else if len(os.Args) == 5 {
		banner = str
		// str = os.Args[len(os.Args)-2]
		// substr = os.Args[len(os.Args)-3]
		str = os.Args[3]
		substr = os.Args[2]
	}

	words, _ := os.ReadFile(banner)

	color := strings.TrimPrefix(option, "--color=")
	// color := strings.Split(option, "--color=")[1]
	// this is for flag
	// color := flag.String("color", "orange", "this is a color")
	// flag.Parse()
	// this is for regexp
	// pattern := regexp.MustCompile(`^--color=(.+)$`)
	// color := pattern.FindStringSubmatch(option)[1]

	colorCode := map[string]string{
		"red":    "\033[38;5;1m",
		"orange": "\033[38;5;208m",
		"green":  "\033[38;5;2m",
		"blue":   "\033[34m",
	}
	code := colorCode[color]

	indexToColour := []int{}
	// s := 0
	// for {
	// 	start := strings.Index(str[s:], substr)
	// 	if start == -1 {
	// 		break
	// 	}
	// 	actualindex := start + s
	// 	end := actualindex + len(substr) - 1
	// 	for i := actualindex; i <= end; i++ {
	// 		indexToColour = append(indexToColour, i)
	// 	}
	// 	s = actualindex + 1
	// }

	for i := 0; i+len(substr) <= len(str); i++ {
		if str[i:i+len(substr)] == substr {
			for j := i; j < i+len(substr); j++ {
				indexToColour = append(indexToColour, j)
			}
		}
	}

	//str := s t r i n g s t r
	//	     0 1 2 3 4 5 6 7 8
	//substr := str
	// str = t r i n g s t r
	//       0 1 2 3 4 5 6 7
	// str = t r
	//       0 1

	// for i, chr := range str {
	// 	if slices.Contains(indexToColour, i) {
	// 		fmt.Print(code + string(chr) + "\033[0m")
	// 	} else {
	// 		fmt.Print(string(chr))
	// 	}
	// }
	// fmt.Println()
	coloredArt := isAscii(words, str, substr, code, indexToColour)
	fmt.Println(coloredArt)

	// fmt.Printf("color -> %s, substring -> %s, text -> %s, banner -> %s\n", color, substr, str, banner)

	// fmt.Println("bot")

	// for i := 0; i < 255; i++ {
	// 	fmt.Printf("\033[38;5;%dm \033[48;5;%dm %3d ", i, i, i)
	// }
	// fmt.Println("\033[0m")

	// fmt.Println("\033[38;5;208m" + "good" + "\033[0m")
}

 