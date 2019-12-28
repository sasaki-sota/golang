package main

import (
	"fmt"
)
// boolearn型の指定の仕方
func main() {
	var t , f bool = true,false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	// 比較をしてくれる
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	// こっちはどちらかがtrueであれば
}