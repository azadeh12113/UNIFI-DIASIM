package mobile

import (
	"fmt"

	"diasim/pkg/core"
)

// World stores the positions of nodes in the simulated two-dimensional space.
type World struct {
	positions map[core.NodeID]Position
}

// NewWorld creates a world with initial positions for its nodes.
// The supplied map is copied so later changes to it do not modify the world.
func NewWorld(initialPositions map[core.NodeID]Position) *World {
	positions := make(map[core.NodeID]Position, len(initialPositions))

	for id, position := range initialPositions {
		positions[id] = position
	}

	return &World{
		positions: positions,
	}
}

// Position returns a node's position.
// It panics if the node has no position, because every node in the world
// is expected to have an initial position.
func (w *World) Position(id core.NodeID) Position {
	position, ok := w.positions[id]
	if !ok {
		panic(fmt.Sprintf("position missing for node %q", id))
	}
	return position
}

// Distance returns the distance between two nodes.
func (w *World) Distance(a, b core.NodeID) float64 {
	posA := w.Position(a)
	posB := w.Position(b)
	return posA.DistanceTo(posB)
}
