package main

// RepeatNTimes ....
func RepeatNTimes(A []int) (rc int) {
	hasSeen := [10000]bool{}

	for _, rc = range A {
		if hasSeen[rc] {
			break
		}
		hasSeen[rc] = true
	}
	return
}
