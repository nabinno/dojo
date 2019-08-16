package main

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
