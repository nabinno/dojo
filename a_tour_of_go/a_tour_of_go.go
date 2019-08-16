package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

func main() {
	Rot13("Lbh penpxrq gur pbqr!")
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		char := b[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b[i] += 13
		} else if (char > 'N' && char <= 'Z') || (char > 'm' && char <= 'z') {
			b[i] -= 13
		}
	}
	return
}

// Rot13 (methods 23)
func Rot13(s string) {
	sr := strings.NewReader(s)
	rot := rot13Reader{sr}
	_, _ = io.Copy(os.Stdout, &rot)
}

type myReader struct{}

func (r myReader) Read(b []byte) (n int, err error) {
	for n, err = 0, nil; n < len(b); n++ {
		b[n] = 'A'
	}
	return
}

// Validate (methods 22)
func Validate() {
	reader.Validate(myReader{})
}

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt (methods 20)
func Sqrt(x float64) (float64, error) {
	var f float64
	var err error
	if x > 0 {
		f = math.Sqrt(x)
	} else {
		f = math.Sqrt(x)
		err = errNegativeSqrt(x)
	}
	return f, err
}

// fmt.Stringer (methods 18)
func (is ipaddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", is[0], is[1], is[2], is[3])
}

type ipaddr [4]byte

// Scale (methods 4)
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type vertex struct {
	X, Y float64
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
