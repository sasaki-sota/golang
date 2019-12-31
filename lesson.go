package main

import (
	"fmt"
)

// structはnilがない
type Vertex struct {
	// x int
	// y int
	X, Y int
}
// 大文字でないとできないので気をつける必要がある

func main() {
	v := Vertex{x: 1, y: 2}
	fmt.Println(v)
	fmt.Println(v.x )
	// 一つだけの取り組みも可能

	v2 := Vertex{x: 1}
	fmt.Println(v2)

	v3 := Vertex{}
	fmt.Println(v3)

	// 何も入れないとできない
}