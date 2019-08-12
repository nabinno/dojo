package main

func main() {
	return
}

// GetLargestPrimeFactorOf (problem 3)
func GetLargestPrimeFactorOf(n int64) int64 {
	var pfs []int64
	var i int64

	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i = 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs[len(pfs)-1]
}
