package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	rc := DetectHaikuLength2()
	fmt.Println(rc)
}

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

// DetectEvenOrOdd (ABS ABC086A)
func DetectEvenOrOdd() string {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if math.Mod(float64(a*b), 2) == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Add (ABS practice 1)
func Add() {
	var a, b, c int
	var s string
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d %d", &b, &c)
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}
