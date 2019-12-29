package main

import (
	"fmt"
)

func by2(num int) string {
	if num % 2 == 0 {
		return "ok"
	}else{
		return "no"
	}
}

func main() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}

	// セミコロンでつなげることができる
	if result2 := by2(10); result2 == "ok"{
		fmt.Println("great 2")
	}

	num := 4
	if num % 2 == 0 {
		fmt.Println("BY 2")
	}else
	{
		fmt.Println("else")
	}

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	// この記述は両方正しい場合

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
	// どちらか正しい場合

}