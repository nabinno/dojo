package main

import (
	"fmt"
	"sort"
	"strings"
)

// DetectHaikuLength2 (ABC042A)
func DetectHaikuLength2() string {
	var a, b, c string
	_, _ = fmt.Scan(&a, &b, &c)
	if strings.Count(a+b+c, "5") == 2 && strings.Count(a+b+c, "7") == 1 {
		return "Yes"
	}
	return "No"
}

// DetectHaikuLength (ABC042A)
func DetectHaikuLength() string {
	checkEqualityOfSlices := func(a, b []int) bool {
		if (a == nil) != (b == nil) {
			return false
		}
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	var a, b, c int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	ints := []int{a, b, c}
	sort.Ints(ints)
	if checkEqualityOfSlices(ints, []int{5, 5, 7}) {
		return "Yes"
	}
	return "No"
}
