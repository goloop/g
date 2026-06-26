package g

import "testing"

// TestPair verifies the generic Pair container holds both fields.
func TestPair(t *testing.T) {
	p := Pair[int, string]{First: 7, Second: "seven"}
	if p.First != 7 || p.Second != "seven" {
		t.Errorf("Pair = %+v, want {7 seven}", p)
	}
}

// TestMaxRangeSize guards the documented range guard constant.
func TestMaxRangeSize(t *testing.T) {
	if MaxRangeSize != 100_000_000 {
		t.Errorf("MaxRangeSize = %d, want 100000000", MaxRangeSize)
	}
}
