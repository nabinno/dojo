package main

type vertex struct {
	X, Y float64
}

// Scale (methods 4)
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
