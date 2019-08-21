package main

import (
	"math"
)

// DetectEvenOrOdd ...
func DetectEvenOrOdd(a, b int) (rc string) {
	if math.Mod(float64(a*b), 2) == 0 {
		rc = "Even"
	} else {
		rc = "Odd"
	}
	return
}
