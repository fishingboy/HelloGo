package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 如果給定初始值的話，類型會自動判斷出來
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
