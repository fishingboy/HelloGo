package main

import (
	"fmt"
	"math"
)

func main() {
	a := float64(9999995)
	var before int32
	for i := 0; i < 10; i++ {
		a++
		fmt.Println(a)
		before = int32(math.Round(a))
		fmt.Printf("before = %v %T\n", before, before)
	}

	//before, err := strconv.Atoi(fmt.Sprint(a))
	//if err != nil {
	//	fmt.Println("err!!")
	//}

}
