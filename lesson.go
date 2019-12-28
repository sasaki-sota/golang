package main

import (
	"fmt"
)

func main() {
	b := []byte{72, 73}
	// {}のカッコないで初期値を設定することができるようになる
	fmt.Println(b)
	fmt.Println(string(b))
	// バイト型で表示するようになる

	c := []byte("HI")
	fmt.Println(c)
	
}