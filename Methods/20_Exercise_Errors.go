package main

import (
	"fmt"
	"math"
	"time"
)

/**
 * 要自訂錯誤，要先建立 struct
 * 並且要加個 method Error 來輸出 error string
 */

type SqrtError struct {
	When time.Time
	What string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return math.Sqrt(x), nil
	} else {
		return 0, &SqrtError{
			time.Now(),
			"it didn't work",
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
