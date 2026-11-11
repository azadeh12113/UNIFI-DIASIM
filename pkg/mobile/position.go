package mobile

import "math"

// Position represents a point or displacement in a two-dimensional space.
// The same type is used for locations and vectors because both have X and Y components.
type Position struct {
	X float64
	Y float64
}

// Add returns the sum of p and other.
func (p Position) Add(other Position) Position {
	return Position{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Sub returns the difference between p and other.
func (p Position) Sub(other Position) Position {
	return Position{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// Scale returns p multiplied by scalar.
func (p Position) Scale(scalar float64) Position {
	return Position{
		X: p.X * scalar,
		Y: p.Y * scalar,
	}
}

// Norm returns the length of p when it is interpreted as a vector.
func (p Position) Norm() float64 {
	return math.Hypot(p.X, p.Y)
}

// DistanceTo calculates the Euclidean distance to another position.
func (p Position) DistanceTo(other Position) float64 {
	return p.Sub(other).Norm()
}
