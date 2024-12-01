package g

import "testing"

// TestParallelTasks tests the ParallelTasks function.
func TestParallelTasks(t *testing.T) {
	maxParallelTasks = 10 // set to whatever value you need
	parallelTasks = 1     // initialize to a default value

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "No parameters",
			input:    nil,
			expected: 1, // should return the default value
		},
		{
			name:     "Zero sum",
			input:    []int{-5, 5},
			expected: 1, // should return minimum value of 1
		},
		{
			name:     "Positive sum less than max",
			input:    []int{3, 2},
			expected: 5, // should return the sum
		},
		{
			name:     "Positive sum greater than max",
			input:    []int{15, 10},
			expected: 10, // should return maxParallelTasks
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParallelTasks(tt.input...); got != tt.expected {
				t.Errorf("ParallelTasks() = %v, want %v", got, tt.expected)
			}
		})
	}
}
