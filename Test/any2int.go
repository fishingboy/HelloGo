package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n any
	n = float64(9999999)
	fmt.Printf("%%T=%T, %%d=%d, %%f=%f, %%.0f=%.0f, %%v=%v \n", n, n, n, n, n)
	//n = any2int(9999999)
	//fmt.Printf("val = %v %T\n", n, n)
}

//func Num64(n interface{}) int64 {
//	s := fmt.Sprintf("%d", n)
//	i, err := strconv.ParseInt(s, 10, 64)
//	if err != nil {
//		return 0
//	} else {
//		return i
//	}
//}

func any2int(num any) int64 {
	str := fmt.Sprintf("%d", num)
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}
