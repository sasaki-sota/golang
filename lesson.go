package main

import (
	"fmt"
)

// 書き換える必要のないものを定義するときに必要
// cpnstはtypeを指定しない
const Pi = 3.14

var big int = 427312424343 + 1

const (
	username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, username, Password)
	fmt.Println(big)
}