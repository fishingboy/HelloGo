package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

}

// 看不太懂，改天再研究好啦！
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Slice the slice to give it zero length.
	// 取得元素 0 以前，所以是取 0 個長度，會出錯
	//s = s[:0]
	//printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}
