package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 把變數初始化也寫在 if 裡面，這樣真的好用嗎？
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	// 正常寫法是這樣吧
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
