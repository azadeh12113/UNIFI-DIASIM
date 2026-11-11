package mobile

import (
	"math"
	"math/rand"
	"testing"
)

const floatTolerance = 1e-9

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func positionsAlmostEqual(a, b Position, tolerance float64) bool {
	return almostEqual(a.X, b.X, tolerance) &&
		almostEqual(a.Y, b.Y, tolerance)
}

func TestPositionDistanceTo(t *testing.T) {
	a := Position{X: 0, Y: 0}
	b := Position{X: 3, Y: 4}

	got := a.DistanceTo(b)

	if !almostEqual(got, 5, floatTolerance) {
		t.Fatalf("DistanceTo() = %v, want 5", got)
	}
}

func TestPositionDistanceToSamePosition(t *testing.T) {
	p := Position{X: 2, Y: 7}

	if got := p.DistanceTo(p); got != 0 {
		t.Fatalf("DistanceTo() = %v, want 0", got)
	}
}

func TestPositionAdd(t *testing.T) {
	got := (Position{X: 2, Y: 3}).Add(Position{X: 4, Y: 1})
	want := Position{X: 6, Y: 4}

	if got != want {
		t.Fatalf("Add() = %v, want %v", got, want)
	}
}

func TestPositionSub(t *testing.T) {
	got := (Position{X: 5, Y: 7}).Sub(Position{X: 2, Y: 3})
	want := Position{X: 3, Y: 4}

	if got != want {
		t.Fatalf("Sub() = %v, want %v", got, want)
	}
}

func TestPositionScale(t *testing.T) {
	got := (Position{X: 3, Y: 4}).Scale(2)
	want := Position{X: 6, Y: 8}

	if got != want {
		t.Fatalf("Scale() = %v, want %v", got, want)
	}
}

func TestPositionNorm(t *testing.T) {
	got := (Position{X: 3, Y: 4}).Norm()

	if !almostEqual(got, 5, floatTolerance) {
		t.Fatalf("Norm() = %v, want 5", got)
	}
}

// Property tests use a fixed seed so failures are reproducible.
func TestDistancePropertiesForRandomPoints(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	const numberOfCases = 10000

	for i := 0; i < numberOfCases; i++ {
		a := randomPosition(rng)
		b := randomPosition(rng)
		c := randomPosition(rng)

		// Distance is symmetric: d(a,b) == d(b,a).
		distanceAB := a.DistanceTo(b)
		distanceBA := b.DistanceTo(a)
		if !almostEqual(distanceAB, distanceBA, floatTolerance) {
			t.Fatalf("case %d: distance symmetry failed: d(a,b)=%g, d(b,a)=%g", i, distanceAB, distanceBA)
		}

		// Triangle inequality: d(a,c) <= d(a,b) + d(b,c).
		distanceAC := a.DistanceTo(c)
		routeViaB := distanceAB + b.DistanceTo(c)
		if distanceAC > routeViaB+floatTolerance {
			t.Fatalf("case %d: triangle inequality failed: d(a,c)=%g, d(a,b)+d(b,c)=%g", i, distanceAC, routeViaB)
		}
	}
}

func TestPositionAddSubPropertyForRandomVectors(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	const numberOfCases = 10000

	for i := 0; i < numberOfCases; i++ {
		v := randomPosition(rng)
		w := randomPosition(rng)

		got := v.Add(w).Sub(w)
		if !positionsAlmostEqual(got, v, floatTolerance) {
			t.Fatalf("(v+w)-w property failed: v=%+v, w=%+v, got=%+v", v, w, got)
		}
	}
}

func randomPosition(rng *rand.Rand) Position {
	const coordinateLimit = 1000.0
	return Position{
		X: (rng.Float64()*2 - 1) * coordinateLimit,
		Y: (rng.Float64()*2 - 1) * coordinateLimit,
	}
}
