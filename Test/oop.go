package main

import "fmt"

type Campaign struct {
	N int
}

func (this Campaign) Double() int {
	return N * 2
}

type Marketing1 struct {
	N int
	Campaign
}

func main() {
	m := Marketing1{N: 100}
	fmt.Println(m, m.N, m.Campaign.N)
	fmt.Println("double -> ", m.Double())
}
