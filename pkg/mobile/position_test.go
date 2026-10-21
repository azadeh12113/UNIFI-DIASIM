package mobile

import (
	"math"
	"testing"
)

func TestPositionDistanceTo(t *testing.T) {
	a := Position{X: 0, Y: 0}
	b := Position{X: 3, Y: 4}

	got := a.DistanceTo(b)

	if math.Abs(got-5) > 1e-9 {
		t.Fatalf("DistanceTo() = %v, want 5", got)
	}
}

func TestPositionDistanceToSamePosition(t *testing.T) {
	p := Position{X: 2, Y: 7}

	if got := p.DistanceTo(p); got != 0 {
		t.Fatalf("DistanceTo() = %v, want 0", got)
	}
}
