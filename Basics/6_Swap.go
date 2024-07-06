package main

import "fmt"

func main() {
	// 「:=」 簡潔賦值語句在明確類型的地方，可以替代 var 定義。
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x string, y string) (string, string) {
	return y, x
}
