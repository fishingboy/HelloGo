package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 選擇一個範圍，包括第一個元素，但不包括最後一個元素
	var s []int = primes[1:4]
	fmt.Println(s)
	// 可以忽略變數類型
	var s2 = primes[1:5]
	fmt.Println(s2)
}
