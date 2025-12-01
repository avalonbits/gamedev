package objects

import (
	"math"
)

type bounds interface {
	Rect() rect
}

type rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, width, height float64) rect {
	return rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r rect) MaxX() float64 {
	return r.X + r.Width
}

func (r rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r rect) Intersects(other rect) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}

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
