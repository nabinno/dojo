package main

import (
	"strconv"
	"strings"
)

func main() {
	return
}

// NbDig (7kyu: Count the Digit)
func NbDig(n int, d int) int {
	var cnt int
	dString := strconv.Itoa(d)
	for n >= 0 {
		squareN := strconv.Itoa(n * n)
		dCnt := len(strings.Split(squareN, dString)) - 1
		cnt += dCnt
		n--
	}
	return cnt
}

// Points a game (8kyu)
// @param {[]string} game - []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
func Points(game []string) int {
	var totalPoint int
	for _, g := range game {
		h := strings.Split(g, ":")
		if h[0] > h[1] {
			totalPoint += 3
		} else if h[0] < h[1] {
			totalPoint += 0
		} else if h[0] == h[1] {
			totalPoint++
		}
	}
	return totalPoint
}
