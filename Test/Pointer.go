package main

import "fmt"

type Name struct {
	First string
	Last  string
}

type User3 struct {
	Name *Name // 使用指針指向 Name 結構體
}

func main() {
	var leo = User3{Name: &Name{First: "Leo", Last: "Kuo"}}
	fmt.Println(leo, leo.Name, leo.Name.First)
}
