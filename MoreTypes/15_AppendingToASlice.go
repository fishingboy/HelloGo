package main

import "fmt"

func main() {
	var slice []int
	printSlice3(slice)

	slice = append(slice, 0)
	printSlice3(slice)

	slice = append(slice, 1)
	printSlice3(slice)

	slice = append(slice, 2, 3, 4)
	printSlice3(slice)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
