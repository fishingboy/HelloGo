package UnitTest

import (
	"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除數不能為 0")
	}

	return a / b, nil
}
