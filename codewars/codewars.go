package main

import (
	"strings"
)

func main() {

}

// Points a game in CodeWar
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
