package g

import (
	"errors"
	"math"
	"slices"
	"sync"
	"testing"
)

// TestRandomNeverPanics covers BUG-01: Random must not panic on degenerate
// ranges and must keep every result inside the requested half-open interval.
func TestRandomNeverPanics(t *testing.T) {
	// Single argument n <= 0 is an empty range and must return the zero value
	// instead of panicking on Intn(0)/Intn(negative).
	if got := Random(0); got != 0 {
		t.Errorf("Random(0) = %d, want 0", got)
	}
	if got := Random(-1); got != 0 {
		t.Errorf("Random(-1) = %d, want 0", got)
	}
	if got := Random(-1000); got != 0 {
		t.Errorf("Random(-1000) = %d, want 0", got)
	}

	// Two equal bounds collapse to that bound.
	if got := Random(7, 7); got != 7 {
		t.Errorf("Random(7,7) = %d, want 7", got)
	}

	// Many iterations: bounds must always hold, in both argument orders.
	for i := 0; i < 100_000; i++ {
		if v := Random(5); v < 0 || v >= 5 {
			t.Fatalf("Random(5) = %d, out of [0,5)", v)
		}
		if v := Random(1, 5); v < 1 || v >= 5 {
			t.Fatalf("Random(1,5) = %d, out of [1,5)", v)
		}
		if v := Random(5, 1); v < 1 || v >= 5 { // reversed bounds
			t.Fatalf("Random(5,1) = %d, out of [1,5)", v)
		}
		if v := Random(0.0, 1.0); v < 0 || v >= 1 {
			t.Fatalf("Random(0.0,1.0) = %f, out of [0,1)", v)
		}
	}

	// More than two arguments: must always pick one of the provided values.
	allowed := map[int]bool{10: true, 20: true, 30: true}
	for i := 0; i < 10_000; i++ {
		if v := Random(10, 20, 30); !allowed[v] {
			t.Fatalf("Random(10,20,30) = %d, not one of the inputs", v)
		}
	}
}

// TestSafeSumInt8Exhaustive proves the generic two's-complement overflow
// detection for every possible pair of signed 8-bit values (BUG-02 contract
// for SafeSum). This is the strongest guarantee we can give cheaply.
func TestSafeSumInt8Exhaustive(t *testing.T) {
	for a := -128; a <= 127; a++ {
		for b := -128; b <= 127; b++ {
			got, err := SafeSum(int8(a), int8(b))
			sum := a + b
			overflow := sum < math.MinInt8 || sum > math.MaxInt8

			if overflow {
				if !errors.Is(err, ErrOverflow) {
					t.Fatalf("SafeSum(int8 %d,%d): want overflow, got %d (err=%v)",
						a, b, got, err)
				}
				continue
			}
			if err != nil {
				t.Fatalf("SafeSum(int8 %d,%d): unexpected err %v", a, b, err)
			}
			if int(got) != sum {
				t.Fatalf("SafeSum(int8 %d,%d) = %d, want %d", a, b, got, sum)
			}
		}
	}
}

// TestSafeSumUint8Exhaustive does the same exhaustive proof for unsigned 8-bit
// values, where only the upper bound can overflow.
func TestSafeSumUint8Exhaustive(t *testing.T) {
	for a := 0; a <= 255; a++ {
		for b := 0; b <= 255; b++ {
			got, err := SafeSum(uint8(a), uint8(b))
			sum := a + b
			overflow := sum > math.MaxUint8

			if overflow {
				if !errors.Is(err, ErrOverflow) {
					t.Fatalf("SafeSum(uint8 %d,%d): want overflow, got %d (err=%v)",
						a, b, got, err)
				}
				continue
			}
			if err != nil {
				t.Fatalf("SafeSum(uint8 %d,%d): unexpected err %v", a, b, err)
			}
			if int(got) != sum {
				t.Fatalf("SafeSum(uint8 %d,%d) = %d, want %d", a, b, got, sum)
			}
		}
	}
}

// TestSafeSumWideBoundaries checks overflow detection at the extremes of the
// wide integer types and the multi-value accumulation path.
func TestSafeSumWideBoundaries(t *testing.T) {
	// No overflow right up to the limit.
	if v, err := SafeSum(math.MaxInt64-1, int64(1)); err != nil || v != math.MaxInt64 {
		t.Errorf("SafeSum(MaxInt64-1,1) = (%d,%v), want (MaxInt64,nil)", v, err)
	}
	// Positive overflow.
	if _, err := SafeSum(int64(math.MaxInt64), 1); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(MaxInt64,1) err = %v, want ErrOverflow", err)
	}
	// Negative overflow.
	if _, err := SafeSum(int64(math.MinInt64), -1); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(MinInt64,-1) err = %v, want ErrOverflow", err)
	}
	// Unsigned 64-bit overflow.
	if _, err := SafeSum(uint64(math.MaxUint64), 1); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(MaxUint64,1) err = %v, want ErrOverflow", err)
	}
	// Mid-stream overflow must be detected even if later terms would cancel it.
	if _, err := SafeSum(int8(127), 1, -1); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(int8 127,1,-1) err = %v, want ErrOverflow", err)
	}
	// Float overflow to +Inf.
	if _, err := SafeSum(math.MaxFloat64, math.MaxFloat64); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(MaxFloat64,MaxFloat64) err = %v, want ErrOverflow", err)
	}
	// NaN input is reported as overflow (non-finite result).
	if _, err := SafeSum(math.NaN(), 1.0); !errors.Is(err, ErrOverflow) {
		t.Errorf("SafeSum(NaN,1) err = %v, want ErrOverflow", err)
	}
}

// TestIsEvenOddLargeAndNegative covers BUG-03: parity must stay correct for
// values that overflow int when converted, and for negatives and fractionals.
func TestIsEvenOddLargeAndNegative(t *testing.T) {
	// 1e19 overflows int64, but as a float it is an even integer.
	if !IsEven(1e19) {
		t.Error("IsEven(1e19) = false, want true")
	}
	if IsOdd(1e19) {
		t.Error("IsOdd(1e19) = true, want false")
	}

	// Negatives.
	if !IsEven(-4.0) || IsEven(-3.0) {
		t.Error("IsEven sign handling wrong")
	}
	if !IsOdd(-3.0) || IsOdd(-4.0) {
		t.Error("IsOdd sign handling wrong")
	}

	// Fractional parts: neither even nor odd unless the f flag ignores them.
	if IsEven(6.6) || IsOdd(7.7) {
		t.Error("fractional value reported as even/odd without f flag")
	}
	if !IsEven(6.6, true) {
		t.Error("IsEven(6.6, true) = false, want true")
	}
	if !IsOdd(7.7, true) {
		t.Error("IsOdd(7.7, true) = false, want true")
	}

	// Invariant: for any whole number exactly one of IsEven/IsOdd holds; for a
	// non-whole number (without f) neither holds.
	for _, v := range []float64{0, 1, -1, 2, -2, 1e6, -1e6, 1e19, 2.5, -3.5} {
		whole := v == math.Trunc(v)
		e, o := IsEven(v), IsOdd(v)
		if whole && e == o {
			t.Errorf("whole %v: IsEven==IsOdd==%v", v, e)
		}
		if !whole && (e || o) {
			t.Errorf("non-whole %v: IsEven=%v IsOdd=%v, want both false", v, e, o)
		}
	}
}

// TestMinMaxNaNDeterministic covers BUG-05: the result must not depend on the
// position of NaN in the input.
func TestMinMaxNaNDeterministic(t *testing.T) {
	nan := math.NaN()
	inputs := [][]float64{
		{1, nan, 3},
		{nan, 1, 3},
		{1, 3, nan},
		{nan, nan, 1},
	}
	for _, in := range inputs {
		if !math.IsNaN(Max(in...)) {
			t.Errorf("Max(%v) is not NaN — NaN handling is order-dependent", in)
		}
		if !math.IsNaN(Min(in...)) {
			t.Errorf("Min(%v) is not NaN — NaN handling is order-dependent", in)
		}
	}

	// Without NaN the result is the ordinary extremum.
	if Max(1.0, 3.0, 2.0) != 3.0 || Min(1.0, 3.0, 2.0) != 1.0 {
		t.Error("Min/Max wrong on clean float input")
	}
}

// TestMaxMinListEmptyAndOk covers BUG-04: the empty case is no longer an
// ambiguous zero — the *Ok variants report it explicitly.
func TestMaxMinListEmptyAndOk(t *testing.T) {
	var empty []int

	if MaxList(empty) != 0 || MinList(empty) != 0 {
		t.Error("MaxList/MinList on empty without defaults must be zero value")
	}

	if _, ok := MaxListOk(empty); ok {
		t.Error("MaxListOk(empty) ok = true, want false")
	}
	if _, ok := MinListOk(empty); ok {
		t.Error("MinListOk(empty) ok = true, want false")
	}

	// Defaults are used when the list is empty.
	if v, ok := MaxListOk(empty, 4, 5, 6); !ok || v != 6 {
		t.Errorf("MaxListOk(empty, 4,5,6) = (%d,%v), want (6,true)", v, ok)
	}
	if v, ok := MinListOk(empty, 4, 5, 6); !ok || v != 4 {
		t.Errorf("MinListOk(empty, 4,5,6) = (%d,%v), want (4,true)", v, ok)
	}

	// Non-empty list wins over defaults.
	if v, ok := MaxListOk([]int{1, 2, 3}, 99); !ok || v != 3 {
		t.Errorf("MaxListOk([1,2,3],99) = (%d,%v), want (3,true)", v, ok)
	}
}

// TestIfFuncIsLazy covers BUG-07: IfFunc must evaluate only the chosen branch.
func TestIfFuncIsLazy(t *testing.T) {
	calls := 0
	a := func() int { calls++; return 1 }
	b := func() int { calls++; return 2 }

	if got := IfFunc(true, a, b); got != 1 || calls != 1 {
		t.Errorf("IfFunc(true) = %d calls=%d, want 1 calls=1", got, calls)
	}

	calls = 0
	if got := IfFunc(false, a, b); got != 2 || calls != 1 {
		t.Errorf("IfFunc(false) = %d calls=%d, want 2 calls=1", got, calls)
	}

	// The motivating case: the false branch must not dereference a nil pointer.
	var p *int
	got := IfFunc(p != nil,
		func() int { return *p }, // would panic if eagerly evaluated
		func() int { return -1 },
	)
	if got != -1 {
		t.Errorf("IfFunc nil-guard = %d, want -1", got)
	}
}

// TestShuffleConcurrent covers BUG-06: Shuffle must be safe under concurrent
// use (run with -race) and must always produce a permutation of the input.
func TestShuffleConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for g := 0; g < 64; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			Shuffle(s)
			sum := 0
			for _, x := range s {
				sum += x
			}
			if sum != 55 {
				t.Errorf("Shuffle lost elements: sum = %d, want 55", sum)
			}
		}()
	}
	wg.Wait()
}

// TestShuffleRandomizes makes sure Shuffle actually reorders (it is not a
// deterministic no-op as a same-nanosecond reseed could cause).
func TestShuffleRandomizes(t *testing.T) {
	orig := Range(0, 1000)
	shuffled := slices.Clone(orig)
	Shuffle(shuffled)
	if slices.Equal(orig, shuffled) {
		t.Error("Shuffle produced an identical ordering for 1000 elements")
	}
	// Still a permutation.
	cp := slices.Clone(shuffled)
	slices.Sort(cp)
	if !slices.Equal(cp, orig) {
		t.Error("Shuffle did not preserve the multiset of elements")
	}
}

// TestSortLargePathological guards against the old quicksort's O(n²) blow-up
// on already-sorted and reverse-sorted input: with stdlib delegation this is
// fast and correct. A regression to naive quicksort would make this hang.
func TestSortLargePathological(t *testing.T) {
	const n = 200_000

	asc := make([]int, n)
	for i := range asc {
		asc[i] = i // already sorted ascending — worst case for tail-pivot
	}
	Sort(asc)
	if !slices.IsSorted(asc) {
		t.Error("Sort failed on already-sorted input")
	}

	desc := make([]int, n)
	for i := range desc {
		desc[i] = n - i // strictly descending
	}
	Sort(desc)
	if !slices.IsSorted(desc) {
		t.Error("Sort failed on reverse-sorted input")
	}

	// Descending mode.
	Sort(desc, true)
	if desc[0] != n || desc[n-1] != 1 {
		t.Errorf("Sort descending wrong ends: [%d ... %d]", desc[0], desc[n-1])
	}
}

// TestClamp verifies clamping, reversed bounds, and string ordering.
func TestClamp(t *testing.T) {
	if got := Clamp(15, 0, 10); got != 10 {
		t.Errorf("Clamp(15,0,10) = %d, want 10", got)
	}
	if got := Clamp(-3, 0, 10); got != 0 {
		t.Errorf("Clamp(-3,0,10) = %d, want 0", got)
	}
	if got := Clamp(7, 0, 10); got != 7 {
		t.Errorf("Clamp(7,0,10) = %d, want 7", got)
	}
	// Reversed bounds are swapped, not a source of wrong results.
	if got := Clamp(7, 10, 0); got != 7 {
		t.Errorf("Clamp(7,10,0) = %d, want 7", got)
	}
	if got := Clamp(15, 10, 0); got != 10 {
		t.Errorf("Clamp(15,10,0) = %d, want 10", got)
	}
	// Works for any Ordered type, e.g. strings.
	if got := Clamp("m", "a", "z"); got != "m" {
		t.Errorf("Clamp(m,a,z) = %q, want m", got)
	}
	if got := Clamp("Z", "a", "z"); got != "a" {
		t.Errorf("Clamp(Z,a,z) = %q, want a", got)
	}
}

// TestRankTies pins down the tie/skip semantics and ordering of Rank.
func TestRankTies(t *testing.T) {
	arr := []float64{1, 5, 2, 3, 7, 8}
	if got := Rank(7.0, arr); got != 1 { // only 8 is greater
		t.Errorf("Rank(7, desc) = %d, want 1", got)
	}
	if got := Rank(8.0, arr); got != 0 { // nothing greater
		t.Errorf("Rank(8, desc) = %d, want 0", got)
	}
	if got := Rank(7.0, arr, true); got != 4 { // 1,2,3,5 are smaller
		t.Errorf("Rank(7, asc) = %d, want 4", got)
	}
	if got := Rank(9.0, arr, true); got != -1 { // not present
		t.Errorf("Rank(9, asc) = %d, want -1", got)
	}

	// Duplicates share a rank and the next rank skips by the number of ties.
	ties := []int{10, 10, 5, 3}
	if got := Rank(10, ties); got != 0 {
		t.Errorf("Rank(10, ties) = %d, want 0", got)
	}
	if got := Rank(5, ties); got != 2 { // both 10s are greater
		t.Errorf("Rank(5, ties) = %d, want 2", got)
	}
	if got := Rank(3, ties); got != 3 { // 10,10,5 are greater
		t.Errorf("Rank(3, ties) = %d, want 3", got)
	}
}

// TestSetDedup verifies that the set operations never emit duplicates even
// when the input slices themselves contain duplicates.
func TestSetDedup(t *testing.T) {
	a := []int{1, 1, 2, 2, 3}
	b := []int{2, 2, 3, 4}

	inter := Intersection(a, b)
	slices.Sort(inter)
	if !slices.Equal(inter, []int{2, 3}) {
		t.Errorf("Intersection with dup inputs = %v, want [2 3]", inter)
	}

	diff := Difference(a, b)
	slices.Sort(diff)
	if !slices.Equal(diff, []int{1}) {
		t.Errorf("Difference with dup inputs = %v, want [1]", diff)
	}

	uni := Union(a, b)
	slices.Sort(uni)
	if !slices.Equal(uni, []int{1, 2, 3, 4}) {
		t.Errorf("Union with dup inputs = %v, want [1 2 3 4]", uni)
	}
}

// TestRangeLimits checks that both Range and Rangef return nil (not an empty
// non-nil slice) when the requested range would exceed MaxRangeSize, and that
// invalid steps are rejected.
func TestRangeLimits(t *testing.T) {
	if got := Range(0, MaxRangeSize+10); got != nil {
		t.Errorf("Range over MaxRangeSize = %v, want nil", got[:0])
	}
	if got := Rangef(func(i int) int { return i }, 0, MaxRangeSize+10); got != nil {
		t.Error("Rangef over MaxRangeSize should be nil")
	}
	if got := Range(1, 10, 0); got != nil {
		t.Error("Range with zero step should be nil")
	}
	if got := Rangef(func(i int) int { return i }, 10, 1, 1); got != nil {
		t.Error("Rangef decreasing range with positive step should be nil")
	}
}

// TestAverageNoOverflow makes sure Average accumulates in float64 and stays
// correct even when the integer sum of the inputs would overflow.
func TestAverageNoOverflow(t *testing.T) {
	// Two MaxInt64 values: their integer sum overflows, but the average is
	// exactly MaxInt64.
	got := Average(int64(math.MaxInt64), int64(math.MaxInt64))
	if want := float64(math.MaxInt64); got != want {
		t.Errorf("Average(MaxInt64,MaxInt64) = %g, want %g", got, want)
	}
}
