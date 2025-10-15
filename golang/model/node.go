package model

import "math"

type Node struct {
	ID       int
	Position Position
	Load     int
}

type Position struct {
	X, Y int
}

func (p Position) CalculateDistance(n1 Node, n2 Node) float32 {
	p1 := n1.Position
	p2 := n2.Position
	return float32(math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2)))
}
