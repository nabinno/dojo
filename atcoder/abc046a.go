package main

// AtCoDeerAndPaintCans ....
func AtCoDeerAndPaintCans(a, b, c int) int {
	m := map[int]int{}
	m[a]++
	m[b]++
	m[c]++
	return len(m)
}
