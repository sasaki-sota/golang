package main

import (
	"fmt"
)

// *を入れることによってバイトのポインタに関係することができるようになる
func main() {
	// var p *int = new(int)
	// fmt.Println(*p)
	// // アドレスは帰ってくる
	// *p++
	// fmt.Println(*p)

	// var p2 *int
	// fmt.Println(p2)

	// newとmakeの違い
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int = new(int)
	fmt.Println("%T\n", p)
}
