package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 55
	v2, ok2 := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)

	m["Answer"] = 0
	v3, ok3 := m["Answer"]
	fmt.Println("The value:", v3, "Present?", ok3)

	delete(m, "Answer")
	v4, ok4 := m["Answer"]
	fmt.Println("The value:", v4, "Present?", ok4)

	// ok 得到的值，是表示該 map 的 key 到底有沒有指定
	// 應該是類似 php 的 isset
	// 但是就算沒有指定也會回傳預設值 0
}
