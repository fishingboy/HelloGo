package main

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m map[string]Vertex5

func main() {
	// 看起來好像 map 就是 key => value 的陣列
	m = make(map[string]Vertex5)
	m["Bell Labs"] = Vertex5{
		40.68433, -74.39967,
	}
	fmt.Println(m)
}
