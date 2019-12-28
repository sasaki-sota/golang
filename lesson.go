package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello" + "world")
	fmt.Println(string("Hello world"[0]))

	var s string = "Hello world"
	
	fmt.Println(strings.Replace(s,"H","X",1))
	
	// strings.relace(stringsはimportする必要がある)
	// 文字列を置き換えてコピーして出力している
}