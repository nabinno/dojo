package main

// JudgeCircle ....
func JudgeCircle(moves string) bool {
	position := [2]int{0, 0}

	for _, move := range moves {
		switch string(move) {
		case "L":
			position[0]++
		case "R":
			position[0]--
		case "U":
			position[1]++
		case "D":
			position[1]--
		}
	}

	if position == [2]int{0, 0} {
		return true
	}
	return false
}
