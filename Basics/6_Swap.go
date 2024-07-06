package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x string, y string) (string, string) {
	return y, x
}
