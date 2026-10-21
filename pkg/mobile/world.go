package mobile

import (
	"fmt"

	"diasim/pkg/core"
)

// World stores the positions of mobile nodes.
type World struct {
	positions map[core.NodeID]Position
}

// NewWorld creates an empty world.
func NewWorld() *World {
	return &World{
		positions: make(map[core.NodeID]Position),
	}
}

// SetPosition assigns a position to a node.
func (w *World) SetPosition(id core.NodeID, p Position) {
	w.positions[id] = p
}

// MoveTo updates a node's position.
func (w *World) MoveTo(id core.NodeID, x, y float64) {
	w.SetPosition(id, Position{X: x, Y: y})
}

// Position returns a node's position and whether it exists.
func (w *World) Position(id core.NodeID) (Position, bool) {
	p, ok := w.positions[id]
	return p, ok
}

// Distance calculates the distance between two nodes.
func (w *World) Distance(a, b core.NodeID) (float64, error) {
	posA, okA := w.positions[a]
	posB, okB := w.positions[b]

	if !okA || !okB {
		return 0, fmt.Errorf(
			"position not found for both nodes %q and %q",
			a, b,
		)
	}

	return posA.DistanceTo(posB), nil
}
