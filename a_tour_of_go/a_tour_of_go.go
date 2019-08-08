package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func main() {
	wc.Test(WordCount)
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
