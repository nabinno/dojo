package main

// SumCandies ...
func SumCandies(n int) (rc int) {
	if n%2 == 0 {
		rc = (n / 2) * (n + 1)
	} else {
		rc = (((n - 1) / 2) * n) + n
	}
	return
}
