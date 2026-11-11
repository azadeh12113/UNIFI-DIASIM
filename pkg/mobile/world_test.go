package mobile

import (
	"math"
	"testing"

	"diasim/pkg/core"
)

func TestWorldGetInitialPosition(t *testing.T) {
	id := core.NodeID("A")
	expected := Position{X: 2, Y: 3}

	world := NewWorld(map[core.NodeID]Position{
		id: expected,
	})

	got := world.Position(id)
	if got != expected {
		t.Fatalf("Position() = %v, want %v", got, expected)
	}
}

func TestWorldDistance(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
		core.NodeID("B"): {X: 3, Y: 4},
	})

	got := world.Distance(core.NodeID("A"), core.NodeID("B"))
	if math.Abs(got-5) > 1e-9 {
		t.Fatalf("Distance() = %v, want 5", got)
	}
}

func TestWorldPositionPanicsWhenPositionIsMissing(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
	})

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("Position() did not panic when a node position was missing")
		}
	}()

	world.Position(core.NodeID("B"))
}

func TestWorldDistancePanicsWhenPositionIsMissing(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
	})

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("Distance() did not panic when a node position was missing")
		}
	}()

	world.Distance(core.NodeID("A"), core.NodeID("B"))
}

func TestWorldCopiesInitialPositions(t *testing.T) {
	initial := map[core.NodeID]Position{
		core.NodeID("A"): {X: 1, Y: 2},
	}
	world := NewWorld(initial)

	initial[core.NodeID("A")] = Position{X: 9, Y: 9}

	got := world.Position(core.NodeID("A"))
	want := Position{X: 1, Y: 2}
	if got != want {
		t.Fatalf("Position() = %v after changing input map, want %v", got, want)
	}
}
