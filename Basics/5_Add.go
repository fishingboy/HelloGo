package main

import (
	"fmt"
)

func main() {
	fmt.Println(add2(42, 13))
}

func add2(x, y int) int {
	return x + y
}
