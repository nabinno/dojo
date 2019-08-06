package main

import (
	"fmt"
	"math"
)

func main() {
	rc := DetectEvenOrOdd()
	fmt.Println(rc)
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
