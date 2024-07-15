package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// 把陣列的某部份取下來，只是參考
	// 儲存空間還是一樣，所以如果元素有變更，會跟著一起變更
	b[0] = "XX"
	fmt.Println(a, b)
	fmt.Println(names)
}
