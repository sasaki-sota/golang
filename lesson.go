package main

// カッコを使うことで複数の指定をすることができるようになる
// userを使うときは階層が一つ下なのでその部分の指定もしてあげる必要がある
import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())
}

