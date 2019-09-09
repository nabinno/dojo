package main

// SortArrayByParity ....
func SortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	firstA := A[l]

	for l < r {
		for l < r && A[r]%2 == 1 {
			r--

		}
		A[l] = A[r]
		for l < r && A[l]%2 == 0 {
			l++

		}
		A[r] = A[l]
	}
	A[l] = firstA

	return A
}
