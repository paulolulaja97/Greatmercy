package main

import (
	"fmt"
)

func main() {

	names := []string{"John", "Paul", "Samuel"}
	ages := []int{10, 15, 20}
	var age int
	fmt.Print("Enter age ")
	fmt.Scanln(&age)
	for i := range len(ages) {
		ages[i] = ages[i] + age
	}
	fmt.Printf("ages,%v\n", ages)

	fmt.Println(names)
	fmt.Println(ages)

}
