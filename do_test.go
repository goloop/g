package do

import (
	"math"
	"reflect"
	"testing"
)

// TestIsEmpty tests the IsEmpty function.
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"Zero int", int(0), true},
		{"Non-zero int", int(1), false},
		{"Zero float", float64(0), true},
		{"Non-zero float", float64(1.0), false},
		{"Empty string", "", true},
		{"Non-empty string", "test", false},
		{"Zero bool", false, true},
		{"Non-zero bool", true, false},
		{"Nil pointer", nil, true},
		{"Non-nil pointer", new(int), false},
		{"Zero complex", complex(0, 0), true},
		{"Non-zero complex", complex(1, 1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.v); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAll tests the All function.
func TestAll(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"All non-zero values", []interface{}{1, true, "test"}, true},
		{"One zero value", []interface{}{1, 0, "test"}, false},
		{"All zero values", []interface{}{0, false, ""}, false},
		{"Empty input", []interface{}{}, false},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For each test case, convert the slice of interfaces to a slice
			// of empty interfaces as the input for the function `All`.
			input := make([]interface{}, len(tt.v))
			for i, val := range tt.v {
				input[i] = val
			}

			// Call the function `All` and check if the output is as expected.
			if got := All(input...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAny tests the Any function.
func TestAny(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"All non-zero values", []interface{}{1, true, "test"}, true},
		{"One non-zero value", []interface{}{0, false, "test"}, true},
		{"All zero values", []interface{}{0, false, ""}, false},
		{"Empty input", []interface{}{}, false},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert slice of interfaces to slice of empty interfaces.
			input := make([]interface{}, len(tt.v))
			for i, val := range tt.v {
				input[i] = val
			}

			// Check if output is as expected.
			if got := Any(input...); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestValue tests the Value function.
func TestValue(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []interface{}
		want interface{}
	}{
		{"Non-zero value is 0 in slice", []interface{}{1, 0, 0}, 1},
		{"Non-zero value is in the 1 of the slice", []interface{}{0, 2, 0}, 2},
		{"Non-zero value is at the 2 of the slice", []interface{}{0, 0, 3}, 3},
		{"All zero values", []interface{}{0, 0, 0}, 0},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert slice of interfaces to slice of empty interfaces.
			input := make([]interface{}, len(tt.v))
			for i, val := range tt.v {
				input[i] = val
			}

			// Check if output is as expected.
			if got := Value(input[0], input[1:]...); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIf tests the If function.
func TestIf(t *testing.T) {
	tests := []struct {
		name     string
		cond     bool
		trueVal  interface{}
		falseVal interface{}
		want     interface{}
	}{
		{"Condition true, returns trueVal", true, "pass", "fail", "pass"},
		{"Condition false, returns falseVal", false, "pass", "fail", "fail"},
		{"Values are integers", true, 5, 10, 5},
		{"Values are different types", true, 10, "ten", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.cond, tt.trueVal, tt.falseVal); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMin tests the Min function.
func TestMin(t *testing.T) {
	tests := []struct {
		name string
		v    int
		more []int
		want int
	}{
		{
			name: "All positive numbers",
			v:    5,
			more: []int{7, 2, 4, 9},
			want: 2,
		},
		{
			name: "Includes negative numbers",
			v:    0,
			more: []int{-7, 2, -3, 9},
			want: -7,
		},
		{
			name: "Single number",
			v:    1,
			more: []int{},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.v, tt.more...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMax tests the Max function.
func TestMax(t *testing.T) {
	tests := []struct {
		name string
		v    int
		more []int
		want int
	}{
		{
			name: "All positive numbers",
			v:    5,
			more: []int{7, 2, 4, 9},
			want: 9,
		},
		{
			name: "Includes negative numbers",
			v:    0,
			more: []int{-7, 2, -3, 9},
			want: 9,
		},
		{
			name: "Single number",
			v:    1,
			more: []int{},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.v, tt.more...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		want int
	}{
		{"Numbers including zero", []int{1, 0, 3}, 4},
		{"Non-zero numbers", []int{2, 2, 6}, 10},
		{"Numbers including negative", []int{-1, 0, 3}, 2},
		{"All zeros", []int{0, 0, 0}, 0},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Sum(tt.v...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAverage tests the Average function.
func TestAverage(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		want float64
	}{
		{"Numbers including zero", []int{1, 0, 3}, 1.3333333333333333},
		{"Non-zero numbers", []int{2, 2, 6}, 3.3333333333333335},
		{"Numbers including negative", []int{-1, 0, 3}, 0.6666666666666666},
		{"All zeros", []int{0, 0, 0}, 0},
		{"Single element", []int{5}, 5},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Average(tt.v...); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMap tests the Map function.
func TestMap(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		f    func(int) int
		want []int
	}{
		{
			"Map with multiplication",
			[]int{1, 2, 3},
			func(n int) int { return n * 2 },
			[]int{2, 4, 6},
		},
		{
			"Map with addition",
			[]int{1, 2, 3},
			func(n int) int { return n + 2 },
			[]int{3, 4, 5},
		},
		{
			"Map with subtraction",
			[]int{1, 2, 3},
			func(n int) int { return n - 1 },
			[]int{0, 1, 2},
		},
		{
			"Map with division",
			[]int{2, 4, 6},
			func(n int) int { return n / 2 },
			[]int{1, 2, 3},
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Map(tt.v, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFilter tests the Filter function.
func TestFilter(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		f    func(int) bool
		want []int
	}{
		{
			"Filter out odd numbers",
			[]int{1, 2, 3, 4, 5},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4},
		},
		{
			"Filter out even numbers",
			[]int{1, 2, 3, 4, 5},
			func(n int) bool { return n%2 != 0 },
			[]int{1, 3, 5},
		},
		{
			"Filter out numbers less than 3",
			[]int{1, 2, 3, 4, 5},
			func(n int) bool { return n >= 3 },
			[]int{3, 4, 5},
		},
		{
			"Filter out numbers greater than 3",
			[]int{1, 2, 3, 4, 5},
			func(n int) bool { return n <= 3 },
			[]int{1, 2, 3},
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Filter(tt.v, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestReduce tests the Reduce function.
func TestReduce(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		f    func(int, int) int
		init int
		want int
	}{
		{
			"Sum of numbers",
			[]int{1, 2, 3, 4, 5},
			func(a, b int) int { return a + b },
			0,
			15,
		},
		{
			"Product of numbers",
			[]int{1, 2, 3, 4, 5},
			func(a, b int) int { return a * b },
			1,
			120,
		},
		{
			"Maximum of numbers",
			[]int{1, 2, 3, 4, 5},
			func(a, b int) int {
				if a > b {
					return a
				} else {
					return b
				}
			},
			0,
			5,
		},
		{
			"Minimum of numbers",
			[]int{1, 2, 3, 4, 5},
			func(a, b int) int {
				if a < b {
					return a
				} else {
					return b
				}
			},
			5,
			1,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Reduce(tt.v, tt.f, tt.init); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestContains tests the Contains function.
func TestContains(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		val  int
		want bool
	}{
		{"Contains in array", []int{1, 2, 3, 4, 5}, 3, true},
		{"Does not contain in array", []int{1, 2, 3, 4, 5}, 6, false},
		{"Empty array", []int{}, 1, false},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Contains(tt.v, tt.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIndex tests the Index function.
func TestIndex(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		v    []int
		val  int
		want int
	}{
		{"Element exists in array", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Element does not exist in array", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty array", []int{}, 1, -1},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Index(tt.v, tt.val); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestZip tests the Zip function.
func TestZip(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		a    []int
		b    []string
		want []Pair[int, string]
	}{
		{
			"Equal length slices",
			[]int{1, 2, 3},
			[]string{"one", "two", "three"},
			[]Pair[int, string]{{1, "one"}, {2, "two"}, {3, "three"}},
		},
		{
			"First slice longer",
			[]int{1, 2, 3, 4, 5},
			[]string{"one", "two", "three"},
			[]Pair[int, string]{{1, "one"}, {2, "two"}, {3, "three"}},
		},
		{
			"Second slice longer",
			[]int{1, 2, 3},
			[]string{"one", "two", "three", "four", "five"},
			[]Pair[int, string]{{1, "one"}, {2, "two"}, {3, "three"}},
		},
		{
			"Empty slices",
			[]int{},
			[]string{},
			[]Pair[int, string]{},
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Zip(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}
