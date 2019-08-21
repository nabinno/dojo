package main

// DetectHaikuLength ...
func DetectHaikuLength(is [3]int) (rc string) {
	var n5, n7 int
	for i := 0; i < 3; i++ {
		if is[i] == 5 {
			n5++
		} else if is[i] == 7 {
			n7++
		}
	}
	if n5 == 2 && n7 == 1 {
		rc = "YES"
	} else {
		rc = "NO"
	}
	return
}
