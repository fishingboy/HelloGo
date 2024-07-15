package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func main() {
	v := Vertex3{1, 2}
	p := &v
	// 這是針對 struct 的指標簡寫
	p.X = 1e9
	fmt.Println(v)
	// 這是原本完整的寫法
	(*p).X = 1e10
	fmt.Println(v)
}
