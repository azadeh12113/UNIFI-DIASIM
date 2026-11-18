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

func TestWorldNeighborsUnitDisk(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
		core.NodeID("B"): {X: 2, Y: 0},    // inside radius
		core.NodeID("C"): {X: 3, Y: 0},    // exactly at radius
		core.NodeID("D"): {X: 3.01, Y: 0}, // outside radius
	})

	got := world.Neighbors(core.NodeID("A"), 3)
	want := []core.NodeID{"B", "C"}
	if !equalNodeIDs(got, want) {
		t.Fatalf("Neighbors(A, 3) = %v, want %v", got, want)
	}
}

func TestWorldNeighborsAreSortedDeterministically(t *testing.T) {
	positions := []struct {
		id core.NodeID
		p  Position
	}{
		{"D", Position{X: 1, Y: 0}},
		{"B", Position{X: 1, Y: 0}},
		{"C", Position{X: 1, Y: 0}},
		{"A", Position{X: 0, Y: 0}},
	}

	first := make(map[core.NodeID]Position)
	second := make(map[core.NodeID]Position)
	for _, item := range positions {
		first[item.id] = item.p
	}
	for i := len(positions) - 1; i >= 0; i-- {
		item := positions[i]
		second[item.id] = item.p
	}

	gotFirst := NewWorld(first).Neighbors("A", 1)
	gotSecond := NewWorld(second).Neighbors("A", 1)
	want := []core.NodeID{"B", "C", "D"}

	if !equalNodeIDs(gotFirst, want) {
		t.Fatalf("Neighbors() = %v, want sorted %v", gotFirst, want)
	}
	if !equalNodeIDs(gotSecond, want) {
		t.Fatalf("Neighbors() with different insertion order = %v, want %v", gotSecond, want)
	}
}

func TestWorldNeighborsExcludeSelfAndReturnEmptyWhenNone(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
		core.NodeID("B"): {X: 10, Y: 10},
	})

	got := world.Neighbors("A", 0)
	if len(got) != 0 {
		t.Fatalf("Neighbors(A, 0) = %v, want empty result", got)
	}
}

func TestWorldNeighborsIncludeOtherNodeAtSamePositionWithZeroRadius(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 1, Y: 1},
		core.NodeID("B"): {X: 1, Y: 1},
	})

	got := world.Neighbors("A", 0)
	want := []core.NodeID{"B"}
	if !equalNodeIDs(got, want) {
		t.Fatalf("Neighbors(A, 0) = %v, want %v", got, want)
	}
}

func TestWorldNeighborsPanicsWhenNodeIsMissing(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
	})

	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatal("Neighbors() did not panic when the requested node was missing")
		}
	}()
	world.Neighbors("missing", 1)
}

func TestWorldNeighborsPanicsForInvalidRadius(t *testing.T) {
	world := NewWorld(map[core.NodeID]Position{
		core.NodeID("A"): {X: 0, Y: 0},
	})

	for _, radius := range []float64{-1, math.NaN(), math.Inf(1), math.Inf(-1)} {
		func() {
			defer func() {
				if recovered := recover(); recovered == nil {
					t.Errorf("Neighbors() did not panic for invalid radius %v", radius)
				}
			}()
			world.Neighbors("A", radius)
		}()
	}
}

func equalNodeIDs(a, b []core.NodeID) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
