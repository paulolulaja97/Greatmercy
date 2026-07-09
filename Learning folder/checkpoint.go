package main

import (
	"fmt"
)

func main() {
	// fmt.Println(FindPrevPrime(5))
	// fmt.Println(FindPrevPrime(4))

	// fmt.Print(FromTo(1, 10))
	// fmt.Print(FromTo(10, 1))
	// fmt.Print(FromTo(10, 10))
	// fmt.Print(FromTo(100, 10))

	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

// func FindPrevPrime(nb int) int {
// 	for nb >= 1 {
// 		if isValid(nb) {
// 			return nb
// 		}
// 		nb--
// 	}
// 	return nb
// }

// func isValid(n int) bool {
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

//

// func FromTo(from int, to int) string {
// 	if from > 99 || from < 0 || to > 99 || to < 0 {
// 		return "Invalid\n"
// 	}
// 	if from == to {
// 		j := strconv.Itoa(from)
// 		return j + "\n"
// 	}
// 	add := ""
// 	if to > from {
// 		for i := from; i <= to; i++ {
// 			j := strconv.Itoa(i)
// 			if i < 10 {
// 				add += "0"
// 			}
// 			add += j
// 			if i != to {
// 				add += ", "
// 			}
// 		}
// 	} else {
// 		for i := from; i >= to; i-- {
// 			j := strconv.Itoa(i)
// 			if i < 10 {
// 				add += "0"
// 			}
// 			add += j
// 			if i != to {
// 				add += ", "
// 			}
// 		}
// 	}
// 	return add + "\n"
// }

func IsCapitalized(s string) bool {

}
