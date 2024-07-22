package main

import (
	"math"
)

func float2int(num float64) int32 {
	return int32(math.Round(num))
}
