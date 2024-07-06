package main

import "fmt"

var c, python, java bool

func main() {
	// 原來宣告完就會給定初始值
	// bool 的初始值為 false
	// int 的初始值為 0
	var i int
	fmt.Println(i, c, python, java)
}
