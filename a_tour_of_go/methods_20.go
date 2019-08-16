package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt (methods 20)
func Sqrt(x float64) (float64, error) {
	var f float64
	var err error
	if x > 0 {
		f = math.Sqrt(x)
	} else {
		f = math.Sqrt(x)
		err = errNegativeSqrt(x)
	}
	return f, err
}
