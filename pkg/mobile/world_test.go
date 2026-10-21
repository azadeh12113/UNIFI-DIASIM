package mobile

import (
	"math"
	"testing"

	"diasim/pkg/core"
)

func TestWorldSetAndGetPosition(t *testing.T) {
	world := NewWorld()
	id := core.NodeID("A")
	expected := Position{X: 2, Y: 3}

	world.SetPosition(id, expected)

	got, ok := world.Position(id)
	if !ok {
		t.Fatal("Position() reported that the node was not found")
	}

	if got != expected {
		t.Fatalf("Position() = %v, want %v", got, expected)
	}
}

func TestWorldDistance(t *testing.T) {
	world := NewWorld()

	world.SetPosition(core.NodeID("A"), Position{X: 0, Y: 0})
	world.SetPosition(core.NodeID("B"), Position{X: 3, Y: 4})

	got, err := world.Distance(core.NodeID("A"), core.NodeID("B"))
	if err != nil {
		t.Fatalf("Distance() returned an unexpected error: %v", err)
	}

	if math.Abs(got-5) > 1e-9 {
		t.Fatalf("Distance() = %v, want 5", got)
	}
}

func TestWorldDistanceMissingPosition(t *testing.T) {
	world := NewWorld()
	world.SetPosition(core.NodeID("A"), Position{X: 0, Y: 0})

	_, err := world.Distance(core.NodeID("A"), core.NodeID("B"))
	if err == nil {
		t.Fatal("Distance() expected an error for a missing position")
	}
}
