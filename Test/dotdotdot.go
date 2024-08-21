package main

import "fmt"

type ItemType int

func AddItem(items ...ItemType) {
	for _, item := range items {
		fmt.Println("Adding item:", item)
	}
}

func main() {
	items := []ItemType{1, 2, 3}
	AddItem(items...) // 展開切片，將切片中的每個元素作為獨立參數傳遞

	AddItem(1, 2, 3, 4, 5, 6) // 展開切片，將切片中的每個元素作為獨立參數傳遞
}
