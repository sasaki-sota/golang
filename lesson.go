package main

import (
	"fmt"
)

func main() {
	// iは０で定義してセミコロンで続けて１０以下の場合はfor文で回すことができる
	for i := 0; i<10; i++ {
		if i ==3 {
			fmt.Println("contineu")
			continue
		}
		if i > 5{
			fmt.Println("break")
			break
			// break文を入れると終了する
		}
		fmt.Println(i)
	}
}