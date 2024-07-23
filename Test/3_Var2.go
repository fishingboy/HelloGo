package main

import (
	"fmt"
	"math"
)

func main() {
	a := float64(9999995)
	var before, before2 int32
	for i := 0; i < 10; i++ {
		a++
		fmt.Println(a)
		before = int32(math.Round(a))
		fmt.Printf("before = %v %T\n", before, before)
		before2 = int32(a)
		fmt.Printf("before2 = %v %T\n", before2, before2)
	}

	//before, err := strconv.Atoi(fmt.Sprint(a))
	//if err != nil {
	//	fmt.Println("err!!")
	//}

}
