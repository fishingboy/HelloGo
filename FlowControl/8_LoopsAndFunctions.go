package main

import "fmt"

// 要實作這個猜想
func Sqrt(x float64) float64 {
	var up float64 = x
	var down float64 = 1
	var sqrt_guest float64
	var mid float64
	for i := 0; i < 100; i++ {
		mid = (up + down) / 2
		sqrt_guest = mid * mid
		if sqrt_guest == x {
			return mid
		} else if sqrt_guest > x {
			up = mid
		} else {
			down = mid
		}
	}
	return mid
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(10))
	fmt.Println(Sqrt(11))
	fmt.Println(Sqrt(12))
	fmt.Println(Sqrt(13))
	fmt.Println(Sqrt(14))
	fmt.Println(Sqrt(15))
	fmt.Println(Sqrt(16))
}
