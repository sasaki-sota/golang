package main

import (
	"fmt"
)

func incrementGneration() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}
	
func main() {
	counter := incrementGneration()
	// functionのものが入っている
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}