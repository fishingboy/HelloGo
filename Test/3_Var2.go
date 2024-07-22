package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := float64(9999995)
	fmt.Println(a)

	before, err := strconv.Atoi(fmt.Sprint(a))
	if err != nil {
		fmt.Println("err!!")
	}

	fmt.Printf("before = %v %T", before, before)
}
