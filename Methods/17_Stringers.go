package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Leo", 40}
	b := Person{"Evonne", 36}
	fmt.Println(a, b)
}
