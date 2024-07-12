package main

import "fmt"

// golang 沒有 while, 只有 for
// 但其實功能一樣
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
