package main

import (
	"fmt"
)

func thirdParty() {
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdParty()
}

func main() {
	save()
	fmt.Println("OK?")
}