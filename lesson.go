package main

import (
	"fmt"
)

func main() {
	n := []int{1,2,3,4,5,6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])
	// これがスライスの方法
	// 指定した配列を出力することができるようになる
	var board = [][]int{
		[]int {0,1,2,},
		[]int{3,4,5},
		[]int{6,7,8},
	}
	fmt.Println(board)
	// 複数の配列を表示することができるようになる
}