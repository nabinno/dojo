package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(GetLargestPrimeFactorOf(600851475143))
}

// GetLargestPrimeFactorOf (problem 3)
// @todo 2019-08-08
func GetLargestPrimeFactorOf(i int64) int64 {
	var rc int64
	for j := sqrt(i); j >= 1; j-- {
		if i%j == 0 && isPrime(j) {
			rc = j
			break
		}
	}
	return rc
}

func isPrime(i int64) bool {
	var rc bool
	if i <= 1 {
		return false
	}
	for j := sqrt(i); j >= 1; j-- {
		if j == 1 {
			rc = true
		}
		if i%j == 0 {
			rc = false
		}
	}
	return rc
}

func sqrt(i int64) int64 {
	return int64(math.Sqrt(float64(i)))
}
