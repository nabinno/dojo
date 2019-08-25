package main

// CartesianNeighbor ....
func CartesianNeighbor(x, y int) (res [][]int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) {
				res = append(res, []int{i, j})
			}
		}
	}
	return
}
