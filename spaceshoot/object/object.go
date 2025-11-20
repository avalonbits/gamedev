package object

import "math"

type vector struct {
	X float64
	Y float64
}

func (v vector) Normalize() vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return vector{
		X: v.X / magnitude,
		Y: v.Y / magnitude,
	}
}
