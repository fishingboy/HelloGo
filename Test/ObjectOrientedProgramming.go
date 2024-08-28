package main

import (
	"fmt"
)

// 定義一個 Animal 結構體，模擬 "父類"
type Animal struct {
	Name string
}

// Animal 結構體的方法
func (a *Animal) Speak() {
	fmt.Printf("%s makes a sound.\n", a.Name)
}

// 定義一個 Dog 結構體，模擬 "子類"，並嵌入 Animal 結構體來實現類似繼承的行為
type Dog struct {
	Animal // 這裡的嵌入使 Dog "繼承" 了 Animal 的屬性和方法
	Breed  string
}

// 覆蓋 Animal 的 Speak 方法
func (d *Dog) Speak() {
	fmt.Printf("%s says Woof!\n", d.Name)
}

// 定義一個 Cat 結構體，模擬另一個 "子類"
type Cat struct {
	Animal // 類似於 php trait 的用法，直接繼承了可用的 function
	Color  string
}

// 覆蓋 Animal 的 Speak 方法
func (c *Cat) Speak() {
	fmt.Printf("%s says Meow!\n", c.Name)
}

// 定義一個 Speaker 接口，這在 Go 中用來實現多態性
type Speaker interface {
	Speak()
}

func main() {
	// 創建 Animal, Dog, Cat 的實例
	animal := Animal{Name: "Generic Animal"}
	dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
	cat := Cat{Animal: Animal{Name: "Whiskers"}, Color: "Black"}

	// 使用 Animal 的方法
	animal.Speak()

	// 使用 Dog 的方法
	dog.Speak()

	// 使用 Cat 的方法
	cat.Speak()

	// 使用 Speaker 接口實現多態性
	var speaker Speaker

	speaker = &dog
	speaker.Speak() // 呼叫 Dog 的 Speak 方法

	speaker = &cat
	speaker.Speak() // 呼叫 Cat 的 Speak 方法
}
