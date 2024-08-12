package main

// @todo: 看不懂，之後再回頭來做
import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
