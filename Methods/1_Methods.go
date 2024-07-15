package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) echo() {
	fmt.Println(v)
}

// golang 沒有 class (太像 C 語言了吧！？)
// 要增加 method 要定義好 struct
// 然後在 struct 上再新增 function
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.echo()
}
