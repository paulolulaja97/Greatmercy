package main

import "fmt"

func main() {
	a := "this is"
	b := "good"
	res := addStr(a, b)
	fmt.Println(res)

}

func addStr(a, b string) string {
	c := a + " " + b
	return c
}
