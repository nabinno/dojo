package main

// GetLargestProductInASeries ....
func GetLargestProductInASeries(digit string) int {
	var adjacentDigits []int

	for _, d := range digit {
		if int(d) == 0 || int(d) == 1 {
			adjacentDigits = nil
			continue
		}

		adjacentDigits = append(adjacentDigits, int(d))

		if len(adjacentDigits) == 13 {
			break
		}
	}

	return 1
}
