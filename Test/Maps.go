package main

import "fmt"

var m map[int]int

func main() {
	// 看起來好像 map 就是 key => value 的陣列
	m = make(map[int]int)
	m[1] = 2
	fmt.Println(m)
}
