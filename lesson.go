package main

import (
	"fmt"
)

func main() {
	var i int = 1
	var f64 = 1.2
	var s string = "test"
	var t bool = true
	var f bool = false
	fmt.Println(i,f64,s,t,f)

	xi := 1 
	// このような形で宣言することもできる
	// ショートバージョンは関数内でしか使うことができない
	// 型をしっかり定義する場合はvarを使うといい
}
// もし値を入れないとデフォルト値が表示される

// func main() {
// 	var (
// 		 i int = 1
// 		f64 = 1.2
// 		s string = "test"
// 		t bool = true
// 		f bool = false
// 	)
// 		fmt.Println(i,f64,s,t,f)
// } このような複数の宣言もすることができる