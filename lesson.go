package main

import (
	"fmt"
)

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++{
		fmt.Println(i, l[i])
	}

	for i, v := range l{
		fmt.Println(i, v)
	}
	// range文で取り出してくることができる
	// _で取り出したい時のみに変更することができるようになる

	m := map[string]int{"apple": 100, "banana": 200}

	for v := range m {
		fmt.Println(v)
	}
}