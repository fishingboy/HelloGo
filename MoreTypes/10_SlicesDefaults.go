package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s2 := []int{1, 2, 3, 4, 5}

	// 取到元素 2 之前的值 0, 1
	fmt.Println(s2[:2])
	// 取得元素 2 (含)之後的值 2, 3, 4
	fmt.Println(s2[2:])
}
