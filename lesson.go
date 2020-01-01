package main

import (
	"fmt"
)
// メソッド名のみを定義している
type Human interface {
	say()
}

type Person struct {
	Name string
}

func (p person) say() {
	fmt.Println(p.name)
}

func main() {
	var mike Human = Person{"Mike"}
	mike.say()
}