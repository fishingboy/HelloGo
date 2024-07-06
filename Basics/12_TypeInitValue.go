package main

import (
	"fmt"
)

// 沒有給定初值的變數
// 會有預設的初始值
// bool = false
// int = 0
// string = ""

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
