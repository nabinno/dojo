package main

import "fmt"

// WelcomeAdd (ABS practice 1)
func WelcomeAdd(a, b, c int, s string) string {
	return fmt.Sprintf("%d %s", a+b+c, s)
}
