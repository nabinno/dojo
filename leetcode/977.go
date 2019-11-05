package main

import "sort"

// SortedSquares ...
func SortedSquares(A []int) []int {
	for k, v := range A {
		A[k] = v * v
	}
	sort.Ints(A)
	return A
}
