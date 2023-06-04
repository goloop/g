package g

import (
	"reflect"
	"sort"
	"testing"
)

// TestUnion tests the Union function.
func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2, 3},
			inputB:   []int{3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "test 2",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 3",
			inputA:   []int{1, 2, 3},
			inputB:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 4",
			inputA:   []int{},
			inputB:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 5",
			inputA:   []int{},
			inputB:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Union(tt.inputA, tt.inputB)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestIntersection tests the Intersection function.
func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2, 3},
			inputB:   []int{3, 4, 5},
			expected: []int{3},
		},
		{
			name:     "test 2",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 3",
			inputA:   []int{1, 2, 3},
			inputB:   []int{},
			expected: []int{},
		},
		{
			name:     "test 4",
			inputA:   []int{},
			inputB:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "test 5",
			inputA:   []int{},
			inputB:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.inputA, tt.inputB)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestDiff tests the Diff and Difference function.
func TestDiff(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2, 3},
			inputB:   []int{3, 4, 5},
			expected: []int{1, 2},
		},
		{
			name:     "test 2",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "test 3",
			inputA:   []int{1, 2, 3},
			inputB:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 4",
			inputA:   []int{},
			inputB:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "test 5",
			inputA:   []int{},
			inputB:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Diff(tt.inputA, tt.inputB)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestSdiff tests the Sdiff and SymmetricDifference function.
func TestSdiff(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2, 3},
			inputB:   []int{3, 4, 5},
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "test 2",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "test 3",
			inputA:   []int{1, 2, 3},
			inputB:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 4",
			inputA:   []int{},
			inputB:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 5",
			inputA:   []int{},
			inputB:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sdiff(tt.inputA, tt.inputB)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComplement tests the Complement function.
func TestComplement(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected []int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3, 4, 5},
			expected: []int{4, 5},
		},
		{
			name:     "test 2",
			inputA:   []int{1, 2, 3},
			inputB:   []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "test 3",
			inputA:   []int{},
			inputB:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "test 4",
			inputA:   []int{1, 2, 3},
			inputB:   []int{},
			expected: []int{},
		},
		{
			name:     "test 5",
			inputA:   []int{},
			inputB:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Complement(tt.inputA, tt.inputB)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestCartesianProduct tests the CartesianProduct function.
func TestCartesianProduct(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   []int
		expected [][2]int
	}{
		{
			name:     "test 1",
			inputA:   []int{1, 2},
			inputB:   []int{3, 4},
			expected: [][2]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}},
		},
		{
			name:     "test 2",
			inputA:   []int{},
			inputB:   []int{3, 4},
			expected: [][2]int{},
		},
		{
			name:     "test 3",
			inputA:   []int{1, 2},
			inputB:   []int{},
			expected: [][2]int{},
		},
		{
			name:     "test 4",
			inputA:   []int{},
			inputB:   []int{},
			expected: [][2]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CartesianProduct(tt.inputA, tt.inputB)

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			} else if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
