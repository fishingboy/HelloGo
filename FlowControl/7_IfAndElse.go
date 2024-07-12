package main

import (
	"fmt"
	"math"
)

func pow3(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow3(3, 2, 10),
		pow3(3, 3, 20),
	)
}
