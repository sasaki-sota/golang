package main

import (
	"os"
	"fmt"
)
// main関数の処理が終わった後に実行される
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
	// deferは下のものから表示される

	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
	// 出力してキャストする
}

