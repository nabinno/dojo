package main

import "math"

// Diagonal ....
func Diagonal(n, p int) int {
	rc := math.Round(float64(fac(n+1) / (fac(n-p) * fac(p+1))))
	return int(rc)
}

func fac(num int) int {
	if num == 0 {
		return 1
	}
	return num * fac(num-1)
}
