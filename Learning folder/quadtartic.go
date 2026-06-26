package main

import (
	"errors"
	"fmt"
	"math"
)

// proccess
// getting the value of a,b, and c from the CLI
// factoring that the input must be an int or float and not a string
// calculating the input gotten by the from the CLI
// factoring if the discriminant is a negative number

func main() {
	var a, b, c float64

	fmt.Println("enter an int or float")
	fmt.Println("enter the first number")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("enter a valid int or float")
		return
	}
	fmt.Println("enter the second number")
	fmt.Scan(&b)
	fmt.Println("enter the third number")
	fmt.Scan(&c)

	x, y, err := (quadratic(a, b, c))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("x ==> %f, y ==> %f\n", x, y)

}

func quadratic(a, b, c float64) (float64, float64, error) {
	if a == 0 {
		fmt.Print("a = 0 ")
		return 0, 0, errors.New("this is a linear equation")
	}
	cal := b*b - 4*a*c
	if cal < 0 {
		return 0, 0, errors.New("cant get the squareroot of a negative number")
	}
	root := math.Sqrt(cal)
	div := -b / (2 * a)
	x := div + root
	y := div - root

	return x, y, nil

}
