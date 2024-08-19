package main

import "fmt"

func sum(x int, c chan int) {
	c <- x
	c <- x + 1
	c <- x * 2
}

func main() {
	c := make(chan int)
	go sum(5, c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
