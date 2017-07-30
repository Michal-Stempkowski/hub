package framework

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	down := Round(0.49)
	up := Round(0.5)

	if math.Abs(down-0.) >= floatPrecision {
		t.Errorf("framework.Round should round %v down", down)
	}

	if math.Abs(up-1.) >= floatPrecision {
		t.Errorf("framework.Round should round %v up", up)
	}
}

func TestFloatEq(t *testing.T) {
	if FloatEq(0.5, 0.49) {
		t.Errorf("framework.FloatEq should say false")
	}

	if !FloatEq(0.5, 0.49+0.01) {
		t.Errorf("framework.FloatEq should say true")
	}
}

func TestFloatLs(t *testing.T) {
	if FloatLs(0.5, 0.5) {
		t.Errorf("framework.FloatLs should say false")
	}

	if !FloatLs(0.5, 0.51) {
		t.Errorf("framework.FloatLs should say true")
	}
}
