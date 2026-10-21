package mobile

import "math"

// Position represents a point in a two-dimensional space.
type Position struct {
	X float64
	Y float64
}

// DistanceTo calculates the Euclidean distance to another position.
func (p Position) DistanceTo(other Position) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y

	return math.Sqrt(dx*dx + dy*dy)
}
