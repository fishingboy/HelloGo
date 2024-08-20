package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("defer 1")
	fmt.Println("2")
	defer fmt.Println("defer 2")
	fmt.Println("3")
	defer fmt.Println("defer 3")
	fmt.Println("4")
	defer fmt.Println("defer 4")
	fmt.Println("5")
	defer fmt.Println("defer 5")
}
