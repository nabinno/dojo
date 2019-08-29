package main

import "math"

// SumSquareDifference ....
func SumSquareDifference(n int) int {
	var rc1, rc2 int
	for i := 1; i <= n; i++ {
		rc1 += int(math.Pow(float64(i), 2))
		rc2 += i
	}
	return int(math.Pow(float64(rc2), 2)) - rc1
}
