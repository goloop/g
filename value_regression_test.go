package g

import (
	"errors"
	"math"
	"testing"
)

func TestSafeSumNamedFloatOverflow(t *testing.T) {
	type amount float64
	for _, values := range [][]amount{{amount(math.MaxFloat64), amount(math.MaxFloat64)}, {amount(math.NaN())}, {amount(math.Inf(-1))}} {
		if _, err := SafeSum(values...); !errors.Is(err, ErrOverflow) {
			t.Fatalf("got %v for %v", err, values)
		}
	}
}

var intersectionResult []int

func BenchmarkIntersectionLarge(b *testing.B) {
	a, c := make([]int, 1024), make([]int, 1024)
	for i := range a {
		a[i] = i
		c[i] = i + 512
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectionResult = Intersection(a, c)
	}
}

func TestPtrNilInterface(t *testing.T) {
	p := Ptr[any](nil)
	if p == nil || *p != nil {
		t.Fatalf("got %#v", p)
	}
}
