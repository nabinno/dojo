package main

import (
	"strconv"
	"strings"
)

// NbDig (Count the Digit)
func NbDig(n int, d int) int {
	var cnt int
	dString := strconv.Itoa(d)
	for n >= 0 {
		squareN := strconv.Itoa(n * n)
		dCnt := len(strings.Split(squareN, dString)) - 1
		cnt += dCnt
		n--
	}
	return cnt
}
