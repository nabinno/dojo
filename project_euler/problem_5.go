package main

// GetSmalletsMultiple ....
func GetSmalletsMultiple(n int) (rc int) {
	rc = 1
	for j := 2; j <= n; j++ {
		rc = lcm(rc, j)
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
