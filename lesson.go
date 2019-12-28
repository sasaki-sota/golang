package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0]= 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// これで初期値を設定することができる
	fmt.Println(b)


}