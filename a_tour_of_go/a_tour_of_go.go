package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	return
}

type ipaddr [4]byte

// fmt.Stringer (methods 18)
func (is ipaddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", is[0], is[1], is[2], is[3])
}

type vertex struct {
	X, Y float64
}

// Scale (methods 4)
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs (methods 1)
func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

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

// WordCount (moretypes 23)
func WordCount(s string) map[string]int {
	rc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		rc[w]++
	}
	return rc
}

// Pic (moretypes 18)
// @example
//   ```
//   pic.Show(Pic)
//   ```
func Pic(dx, dy int) [][]uint8 {
	rc := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(y)
		}
		rc[y] = row
	}
	return rc
}

// PrintSlice (moretypes 11)
func PrintSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
