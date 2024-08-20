package main

import "fmt"

type User struct {
	name string
	age  int
}
type UserData struct {
	email   string
	address string
}

var m2 map[User]UserData

func main() {
	// 看起來好像 map 就是 key => value 的陣列
	m2 = make(map[User]UserData)

	user := User{name: "Leo", age: 40}
	user_data := UserData{email: "et282523@hotmail.com", address: "大同路三段"}

	m2[user] = user_data
	fmt.Println(m2)
}
