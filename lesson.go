package main

import (
	"fmt"
	"log"
	"os"
)

// エラー処理の文
func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatal("error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("error")
	}
}