package main

// Fibonacci (moretypes 26)
func Fibonacci() func() int {
	var i, a, b, c int
	b = 1
	c = 0
	return func() int {
		switch i {
		case 0:
			i++
		default:
			a = b
			b = c
			c = a + b
			i++
		}
		return c
	}
}
