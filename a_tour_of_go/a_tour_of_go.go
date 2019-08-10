package main

import (
	"fmt"
	"math"
	"strings"
)

// Vertex struct
type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// Scale (methods 4)
// @example
//   ```
//   v := Vertex{3, 4}
//   v.Scale(10)
//   fmt.Println(v.Abs())
//   ```
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs (methods 1)
// @example
//   ```
//   v := Vertex{3, 4}
//   fmt.Println(v.abs())
//   ```
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Fibonacci (moretypes 26)
// @example
//   ```
//   f := Fibonacci()
//   for i := 0; i < 10; i++ {
//       fmt.Println(f(i))
//   }
//   ```
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
// @example
//   ```
//   wc.Test(WordCount)
//   ```
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
// @example
//   ```
//   s := []int{2, 3, 5, 7, 11, 13}
//   printSlice(s)
//   ```
func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
