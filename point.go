package lof

import "math"

type Point struct {
	X float32
	Y float32
}

func (p *Point) Equals(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

func (p *Point) Distance(point Point) float32 {
	x := p.X - point.X
	y := p.Y - point.Y
	return float32(math.Sqrt(float64(x*x + y*y)))
}
