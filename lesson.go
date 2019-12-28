package main

import (
	"fmt"
)

func add(x int, y int) (int, int) {
	return x + y, x - y
}
// 引数は第一引数と第二引数を用いる必要がある

// resultは変数名
func cal(price, item int) (result int) {
	result = price * item
	return result
}

func converter(price int) float64 {
	return float64(price)
}

func main() {
	// add(10, 20)
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func() {
		fmt.Println("inner func")
	}
	f(1)
}