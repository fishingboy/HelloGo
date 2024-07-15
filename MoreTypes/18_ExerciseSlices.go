package main

// 不知道到底要幹嘛，先跳過
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	points := [][]uint8{
		[]uint8{1, 2, 3, 4, 5, 6, 7, 8},
		[]uint8{1, 2, 3, 4, 5, 6, 7, 8},
		[]uint8{1, 2, 3, 4, 5, 6, 7, 8},
	}
	return points
}

func main() {
	pic.Show(Pic(0, 0))
}
