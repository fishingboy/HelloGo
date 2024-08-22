package main

import "fmt"

var a = "hello" // 包作用域

func main() {
	b := "world" // 函數作用域
	if true {
		c := "!"             // 區塊作用域
		fmt.Println(a, b, c) // 可以訪問 a, b, c
	}
	// fmt.Println(c)  // 這會導致編譯錯誤，因為 c 在這裡不可見
	fmt.Println(a, b) // 可以訪問 a 和 b
}

func anotherFunction() {
	fmt.Println(a) // 可以訪問 a
	// fmt.Println(b)  // 這會導致編譯錯誤，因為 b 在這裡不可見
}
