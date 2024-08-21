package main

import "fmt"

func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = Swap(x, y)
	fmt.Println(x, y)

	s1, s2 := "hello", "world"
	fmt.Println(s1, s2)
	s1, s2 = Swap(s1, s2)
	fmt.Println(s1, s2)
}
