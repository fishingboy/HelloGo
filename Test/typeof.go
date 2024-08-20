package main

import (
	"fmt"
	"reflect"
)

type User2 struct {
	name string
	age  int
}

func main() {
	var a int = 10
	var b = User2{name: "leo", age: 40}
	value := reflect.ValueOf(a)

	fmt.Println("Type:", value.Type())
	fmt.Println("Kind:", value.Kind())
	fmt.Println("Value:", value.Interface())

	value = reflect.ValueOf(b)
	fmt.Println("Type:", value.Type())
	fmt.Println("Kind:", value.Kind())
	fmt.Println("Value:", value.Interface())
}
