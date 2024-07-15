package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b)

	c := b[:2]
	fmt.Println("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
