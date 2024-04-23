package g

import "testing"

// TestPtr ensures the Ptr function works correctly for all intended use cases.
func TestPtr(t *testing.T) {
	equalIntSlices := func(a, b []int) bool {
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

	// Test integer pointer creation from literal.
	if got := *Ptr(42); got != 42 {
		t.Errorf("Ptr(42) = %v, want 42", got)
	}

	// Test creation of pointer to default value.
	if got := *Ptr[int](); got != 0 {
		t.Errorf("Ptr[int]() = %v, want 0", got)
	}

	// Test float pointer creation from literal.
	if got := *Ptr(42.5); got != 42.5 {
		t.Errorf("Ptr(42.5) = %v, want 42.5", got)
	}

	// Test string pointer creation from literal.
	if got := *Ptr("hello"); got != "hello" {
		t.Errorf("Ptr(\"hello\") = %v, want \"hello\"", got)
	}

	// Test boolean pointer creation from literal.
	if got := *Ptr(true); !got {
		t.Errorf("Ptr(true) = %v, want true", got)
	}

	// Test creating a pointer from a complex type (slice).
	slice := []int{1, 2, 3}
	if got := *Ptr(slice); !equalIntSlices(got, slice) {
		t.Errorf("Ptr([]int{1, 2, 3}) = %v, want %v", got, slice)
	}
}
