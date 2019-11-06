package main

// BalancedStringSplit ....
func BalancedStringSplit(s string) int {
	rc, count := 0, 0
	for _, b := range s {
		if b == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			rc++
		}
	}
	return rc
}
