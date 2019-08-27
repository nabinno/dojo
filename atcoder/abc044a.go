package main

// TakeAndHotels ....
func TakeAndHotels(n, k, x, y int) (rc int) {
	for i := 1; i <= n; i++ {
		if i <= k {
			rc += x
			continue
		}
		rc += y
	}
	return
}
