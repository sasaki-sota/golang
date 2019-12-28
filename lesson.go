package main

import (
	"strconv"
	"fmt"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	// このような形で肩の変換をすることができるようになる

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy,yy,yy)

	var s string = "12"
	i, _:= strconv.Atoi(s)
	// strconvは文字列のコンバージョン
	fmt.Printf("%T %v", i, i)
	// _を使うことでエラーを防ぐことができる
}