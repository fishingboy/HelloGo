package main

import "fmt"

type I4 interface {
	M()
}

func describe4(i I4) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I4
	describe4(i)
	i.M()
}
