package main

import "fmt"

func doubleCallByValue(x int) {
	x = x * 2
}

func doubleCallByReference(x *int) {
	*x = *x * 2
}

func (site *Site) DoublePrice() {
	site.price = site.price * 2
}

type Site struct {
	url   string
	price int
}

func main() {
	a := 10
	fmt.Println(a)
	doubleCallByValue(a)
	fmt.Println(a)

	b := 10
	fmt.Println(b)
	doubleCallByReference(&b)
	fmt.Println(b)

	site1 := Site{url: "leo-kuo.com", price: 100}
	site1.DoublePrice()
	fmt.Println(site1)

	site2 := &Site{url: "evonne-chen.com", price: 300}
	site2.DoublePrice()
	fmt.Println(site2)

}
