package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Golang 的 switch 不用寫 break
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
