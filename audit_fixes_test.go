package g_test

import (
	"errors"
	"math"
	"testing"

	g "github.com/goloop/g/v2"
)

// TestRandomWideIntegerRanges guards BUG-01: full-width integer ranges must not
// collapse to the lower bound.
func TestRandomWideIntegerRanges(t *testing.T) {
	seenU := map[uint64]struct{}{}
	seenI := map[int64]struct{}{}
	for i := 0; i < 64; i++ {
		seenU[g.Random[uint64](0, math.MaxUint64)] = struct{}{}
		seenI[g.Random[int64](math.MinInt64, math.MaxInt64)] = struct{}{}
	}
	if len(seenU) < 2 {
		t.Errorf("Random[uint64](0, MaxUint64) produced %d unique values, want many", len(seenU))
	}
	if len(seenI) < 2 {
		t.Errorf("Random[int64](MinInt64, MaxInt64) produced %d unique values, want many", len(seenI))
	}
}

// TestRandomFloatExtremesFinite guards BUG-09: extreme float bounds must stay
// finite and within range.
func TestRandomFloatExtremesFinite(t *testing.T) {
	for i := 0; i < 128; i++ {
		v := g.Random[float64](-math.MaxFloat64, math.MaxFloat64)
		if math.IsInf(v, 0) || math.IsNaN(v) {
			t.Fatalf("Random float extreme returned non-finite %v", v)
		}
	}
}

// TestRangeExtremesNoPanic guards BUG-02: extreme bounds must return nil, not
// panic or a wrong slice.
func TestRangeExtremesNoPanic(t *testing.T) {
	if r := g.Range(-1, math.MaxInt); r != nil {
		t.Errorf("Range(-1, MaxInt) = %d-element slice, want nil", len(r))
	}
	if r := g.Range(math.MinInt, math.MaxInt); r != nil {
		t.Errorf("Range(MinInt, MaxInt) = %v, want nil", r)
	}
	// A small range still works.
	if got := g.Range(1, 5); len(got) != 4 {
		t.Errorf("Range(1,5) = %v, want 4 elements", got)
	}
}

// TestParityBeyondMantissa guards BUG-03: parity stays correct past 2^53.
func TestParityBeyondMantissa(t *testing.T) {
	if g.IsEven(int64(math.MaxInt64)) {
		t.Error("IsEven(MaxInt64) = true, want false (MaxInt64 is odd)")
	}
	if !g.IsOdd(int64(math.MaxInt64)) {
		t.Error("IsOdd(MaxInt64) = false, want true")
	}
	if g.IsEven(int64(1)<<53 + 1) {
		t.Error("IsEven(2^53+1) = true, want false")
	}
	if !g.IsEven(uint64(math.MaxUint64) - 1) {
		t.Error("IsEven(MaxUint64-1) = false, want true")
	}
}

// TestFloatToStringNamedFloat32 guards BUG-04: a named ~float32 formats at
// 32-bit width.
func TestFloatToStringNamedFloat32(t *testing.T) {
	type Celsius float32
	if got := g.FloatToString(Celsius(3.14)); got != "3.14" {
		t.Errorf("FloatToString(Celsius(3.14)) = %q, want %q", got, "3.14")
	}
	if got := g.FloatToString(float32(3.14)); got != "3.14" {
		t.Errorf("FloatToString(float32(3.14)) = %q, want %q", got, "3.14")
	}
}

// TestLookupShortResult guards BUG-05: a match beyond the result slice returns
// the default instead of panicking.
func TestLookupShortResult(t *testing.T) {
	if got := g.HLookup("C", []string{"A", "B", "C"}, []int{1}, -1); got != -1 {
		t.Errorf("HLookup short result = %d, want -1", got)
	}
	if got := g.VLookup("C", []string{"A", "B", "C"}, []int{1}, -1); got != -1 {
		t.Errorf("VLookup short result = %d, want -1", got)
	}
}

// TestIsEmptyContainers guards BUG-06: consistent length semantics for slice,
// map and channel; arrays use zero-value semantics.
func TestIsEmptyContainers(t *testing.T) {
	if !g.IsEmpty([]int{}) {
		t.Error("IsEmpty([]int{}) = false, want true")
	}
	if !g.IsEmpty(map[string]int{}) {
		t.Error("IsEmpty(empty map) = false, want true")
	}
	if !g.IsEmpty([3]int{0, 0, 0}) {
		t.Error("IsEmpty([3]int zero) = false, want true")
	}
	if g.IsEmpty([]int{1}) || g.IsEmpty(map[string]int{"a": 1}) || g.IsEmpty([3]int{1}) {
		t.Error("a non-empty container reported as empty")
	}
}

// TestStringToEmptyError guards BUG-07: the empty-input error is ErrEmptyString
// regardless of a supplied default.
func TestStringToEmptyError(t *testing.T) {
	if _, err := g.StringToBool("", true); !errors.Is(err, g.ErrEmptyString) {
		t.Errorf("StringToBool empty err = %v, want ErrEmptyString", err)
	}
	if _, err := g.StringToInt("", 7); !errors.Is(err, g.ErrEmptyString) {
		t.Errorf("StringToInt empty err = %v, want ErrEmptyString", err)
	}
	if _, err := g.StringToFloat("", 1.5); !errors.Is(err, g.ErrEmptyString) {
		t.Errorf("StringToFloat empty err = %v, want ErrEmptyString", err)
	}
	if _, err := g.StringToBool("nope"); !errors.Is(err, g.ErrInvalidBool) {
		t.Errorf("StringToBool invalid err = %v, want ErrInvalidBool", err)
	}
}

// TestAbsOk guards BUG-08: AbsOk reports overflow for the most negative value.
func TestAbsOk(t *testing.T) {
	if v, ok := g.AbsOk(-10); v != 10 || !ok {
		t.Errorf("AbsOk(-10) = (%d, %v), want (10, true)", v, ok)
	}
	if v, ok := g.AbsOk(int64(math.MinInt64)); v != math.MinInt64 || ok {
		t.Errorf("AbsOk(MinInt64) = (%d, %v), want (MinInt64, false)", v, ok)
	}
	if v, ok := g.AbsOk(uint(20)); v != 20 || !ok {
		t.Errorf("AbsOk(uint 20) = (%d, %v), want (20, true)", v, ok)
	}
}
