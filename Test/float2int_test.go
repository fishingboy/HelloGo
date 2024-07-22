package main

import (
	"fmt"
	"testing"
)

func TestFloat2int(t *testing.T) {
	var val int32
	for i := 0; i < 99999999; i++ {
		var f = float64(i)
		val = float2int(f)
		fmt.Println(val)
	}
}
