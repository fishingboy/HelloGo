package main

import "fmt"

func main() {
	/**
		  a := 1
	      就等同於
		  var a int
		  a = 1
	*/
	fmt.Println(var_type1())
	fmt.Println(var_type2())
}

func var_type1() int {
	var a int
	a = 1
	return a
}

func var_type2() int {
	a := 1
	return a
}
