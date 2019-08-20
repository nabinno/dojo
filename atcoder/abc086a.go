package main

import (
	"fmt"
	"math"
)

// DetectEvenOrOdd (ABS ABC086A)
func DetectEvenOrOdd() string {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if math.Mod(float64(a*b), 2) == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
