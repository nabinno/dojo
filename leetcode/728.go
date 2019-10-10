package main

// SelfDividingNumbers ...
func SelfDividingNumbers(left int, right int) (rc []int) {
	for i := left; i <= right; i++ {
		flag := true
		for tmp := i; flag && tmp > 0; {
			if tmp%10 == 0 || i%(tmp%10) > 0 {
				flag = false
			}
			tmp /= 10
		}
		if flag {
			rc = append(rc, i)
		}
	}
	return
}
