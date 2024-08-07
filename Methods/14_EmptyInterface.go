package main

import "fmt"

// empty interface 可以接受任何型態的值
func main() {
	var i interface{}
	describe5(i)

	i = 42
	describe5(i)

	i = "hello"
	describe5(i)
}

func describe5(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
