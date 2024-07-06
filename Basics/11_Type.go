package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// go lang 基本型態
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // uint8 的别名
// rune // int32 的别名
//      // 表示一个 Unicode 码位
// float32 float64
// complex64 complex128

func main() {
	fmt.Printf("Type: %T, Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
