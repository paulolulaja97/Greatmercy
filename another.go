package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	data := strings.Fields(string(file))
	// for _, w := range data {
	// 	fmt.Println(w)
	// }

	puncs := map[byte]bool{'.': true, ',': true, '?': true, '!': true, ':': true, ';': true}
	seen := false
	var start int
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	for i := 0; i < len(data); i++ {
		if data[i] == "(hex)" {
			v := data[i-1]
			v_i, _ := strconv.ParseInt(v, 16, 64)
			v_s := strconv.Itoa(int(v_i))
			data[i-1] = v_s
			data = slices.Delete(data, i, i+1)
			i--

		}
		if data[i] == "(bin)" {
			v := data[i-1]
			v_i, _ := strconv.ParseInt(v, 2, 64)
			v_S := strconv.Itoa(int(v_i))
			data[i-1] = v_S
			data = slices.Delete(data, i, i+1)
			i--
		}
		if data[i] == "(up)" {
			v := data[i-1]
			v_u := strings.ToUpper(v)
			data[i-1] = v_u
			data = slices.Delete(data, i, i+1)
			i--
		}
		if data[i] == "(low)" {
			v := data[i-1]
			v_l := strings.ToLower(v)
			data[i-1] = v_l
			data = slices.Delete(data, i, i+1)
			i--
		}
		if data[i] == "(cap)" {
			v := data[i-1][0]
			v_c := strings.ToUpper(string(v))
			data[i-1] = v_c + data[i-1][1:]
			data = slices.Delete(data, i, i+1)
			i--

		}
		if data[i] == "(low," {
			v := strings.TrimRight(data[i+1], ")")
			v_i, _ := strconv.Atoi(string(v))
			for a := v_i; a > 0; a-- {
				data[i-a] = strings.ToLower(data[i-a])
			}
			data = slices.Delete(data, i, i+2)
			i--

		}
		if data[i] == "(cap," {
			v := ""
			// for j := 0; j < len(data[i+1]); j++ {
			// 	if data[i+1][j] >= '0' && data[i+1][j] <= '9' {
			// 		v += string(data[i+1][j])
			// 	}
			// }
			v = strings.TrimRight(data[i+1], ")")
			v_t, _ := strconv.Atoi(string(v))
			for b := v_t; b > 0; b-- {
				v := data[i-b][0]
				v_c := strings.ToUpper(string(v))
				data[i-b] = v_c + data[i-b][1:]
			}
			data = slices.Delete(data, i, i+2)
			i--
		}
		if data[i] == "(up," {
			v := data[i+1][0]
			v_r, _ := strconv.Atoi(string(v))
			for t := v_r; t > 0; t-- {
				data[i-t] = strings.ToUpper(data[i-t])
			}
			data = slices.Delete(data, i, i+2)
			i--
		}
		if puncs[data[i][0]] {
			for _, p := range data[i] {
				if !puncs[byte(p)] {
					break
				}
				data[i-1] += string(p)
				data[i] = data[i][1:]
			}
			if data[i] == "" {
				data = slices.Delete(data, i, i+1)
				i--

			}

		}
		if data[i] == "'" {
			if !seen {
				seen = true
				start = i
			} else {
				data[i] += strings.Join(data[start+1:i], " ") + "'"
				data = slices.Delete(data, start, i)
				i = start

			}

		}
		if data[i] == "a" {
			if vowels[data[i+1][0]] {
				data[i] += "n"
			}
		}

	}

	value := strings.Join(data, " ")
	os.WriteFile(output, []byte(value), 0664)

}

// mapping_value := map[string]int{"ade": 20: "bola": 30}

// mapping_value[ade]
