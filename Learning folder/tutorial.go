package main

import (
	"fmt"
	"math"
)

// func isValidTemplate(filename string) bool {
// 	new := strings.ToLower(filename)
// 	if strings.Contains(new, ".html") || strings.Contains(new, ".tmpl") {
// 		return true
// 	}
// 	return false
// 	// 	var name string = "paul"
// 	// 	name2 := "paul"
// 	// 	const name3 = "paul"

// 	// fmt.Println(name)
// 	// fmt.Println(name2)
// 	// fmt.Println(name3)
// }

// func ByeByeFirst(strings []string) []string {

// 	if len(strings) == 0 {
// 		return strings
// 	}
// 	return strings[1:]

// }
// getting user input from the CLI with fmt.scan
// validate if our input is a valid float
// calculate all our input
// square all our result
// return my quadratic output
func main() {

	// if validateInput(a, b, c) {
	// 	fmt.Println("all valid")
	// } else {
	// 	fmt.Println("invalid ")
	// }
	// return

	// 	// fmt.Println(ByeByeFirst([]string{}))
	// 	fmt.Println(HalfSlice([]int{1, 2, 3}))
	// 	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	// }

	// // func HalfSlice(slice []int) []int {
	// // 	half := len(slice) / 2
	// // 	if len(slice)%2 == 1 {
	// // 		return slice[:half+1]

	// // 	}
	// // 	return slice[:half]
	// // }

	// func swapLast (s []int)[]int {
	// 	return s[:len(s)-3] + []int{s[len(s)-2]} + []int{s[len(s)-1]}

}

func validateInput(n ...float64) bool {

	for _, g := range n {
		if g < 0.0 {
			fmt.Printf("error: %s is not a float", g)
			return false
		}

	}

	return true
}

func quadratic(a, b, c float64) (float64, float64) {
	cal := b*b - 4*a*c
	root := math.Sqrt(cal)
	div := -b / 2 * a
	x := div + root
	y := div - root

	return x, y

}
