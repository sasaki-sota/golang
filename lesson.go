package main

import (
	"fmt"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logFile, _ := os.OpenFile(logFile, os)
}
// logで時間も一緒に出力される
func main() {
	file, err := os.Open("database")
	if err != nil {
		log.Fatalln("Exit")
	}

	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error")
	fmt.Println("okk")
}