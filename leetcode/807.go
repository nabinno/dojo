package main

import "math"

// MaxIncreaseKeepingSkyline ....
func MaxIncreaseKeepingSkyline(grid [][]int) (rc int) {
	rowMaxes := make(map[int]int)
	colMaxes := make(map[int]int)

	for row, currRow := range grid {
		for col, item := range currRow {
			rowMaxes[row] = int(math.Max(float64(rowMaxes[row]), float64(item)))
			colMaxes[col] = int(math.Max(float64(colMaxes[col]), float64(item)))
		}
	}

	for row, currRow := range grid {
		for col, item := range currRow {
			rc += int(math.Min(float64(rowMaxes[row]), float64(colMaxes[col]))) - item
		}
	}

	return
}
