package do

import (
	"math"
	"reflect"
	"testing"
)

// TestSortAscending sorts a slice of integers in ascending order.
func TestSortAscending(t *testing.T) {
	// Test case 1: Unsorted slice.
	numbers1 := []int{3, 5, 1, 9, 2}
	Sort(numbers1)
	expected1 := []int{1, 2, 3, 5, 9}
	if !reflect.DeepEqual(numbers1, expected1) {
		t.Errorf("Sort ascending: expected %v, but got %v",
			expected1, numbers1)
	}

	// Test case 2: Sorted slice.
	numbers2 := []int{1, 2, 3, 4, 5}
	Sort(numbers2)
	expected2 := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(numbers2, expected2) {
		t.Errorf("Sort ascending: expected %v, but got %v",
			expected2, numbers2)
	}

	// Test case 3: Empty slice.
	numbers3 := []int{}
	Sort(numbers3)
	expected3 := []int{}
	if !reflect.DeepEqual(numbers3, expected3) {
		t.Errorf("Sort ascending: expected %v, but got %v",
			expected3, numbers3)
	}
}

// TestSortDescending sorts a slice of integers in descending order.
func TestSortDescending(t *testing.T) {
	// Test case 1: Unsorted slice.
	numbers1 := []int{3, 5, 1, 9, 2}
	Sort(numbers1, true)
	expected1 := []int{9, 5, 3, 2, 1}
	if !reflect.DeepEqual(numbers1, expected1) {
		t.Errorf("Sort descending: expected %v, but got %v",
			expected1, numbers1)
	}

	// Test case 2: Sorted slice.
	numbers2 := []int{5, 4, 3, 2, 1}
	Sort(numbers2, true)
	expected2 := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(numbers2, expected2) {
		t.Errorf("Sort descending: expected %v, but got %v",
			expected2, numbers2)
	}

	// Test case 3: Empty slice.
	numbers3 := []int{}
	Sort(numbers3, true)
	expected3 := []int{}
	if !reflect.DeepEqual(numbers3, expected3) {
		t.Errorf("Sort descending: expected %v, but got %v",
			expected3, numbers3)
	}
}

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

	// Iterate over each test case.
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
	// Define test cases.
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

	// Iterate over each test case.
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

// TestMinList tests the MinList function.
func TestMinList(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name     string
		list     []int
		defaults []int
		want     int
	}{
		{
			name:     "Non-empty list, no defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: nil,
			want:     1,
		},
		{
			name:     "Empty list, no defaults",
			list:     []int{},
			defaults: nil,
			want:     0,
		},
		{
			name:     "Empty list, with defaults",
			list:     []int{},
			defaults: []int{20, 10},
			want:     10,
		},
		{
			name:     "Non-empty list with defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: []int{20, 10},
			want:     1,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the MinList function.
			got := MinList(tt.list, tt.defaults...)

			// Check if the result matches the expected value.
			if got != tt.want {
				t.Errorf("MinList() = %v, want %v", got, tt.want)
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

// TestMaxList tests the MaxList function.
func TestMaxList(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name     string
		list     []int
		defaults []int
		want     int
	}{
		{
			name:     "Non-empty list, no defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: nil,
			want:     9,
		},
		{
			name:     "Empty list, no defaults",
			list:     []int{},
			defaults: nil,
			want:     0,
		},
		{
			name:     "Empty list, with defaults",
			list:     []int{},
			defaults: []int{20, 10},
			want:     20,
		},
		{
			name:     "Non-empty list with defaults",
			list:     []int{3, 5, 7, 1, 9, 2},
			defaults: []int{20, 10},
			want:     9,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the MaxList function.
			got := MaxList(tt.list, tt.defaults...)

			// Check if the result matches the expected value.
			if got != tt.want {
				t.Errorf("MaxList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	// Define test cases.
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

	// Iterate over each test case.
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

// TestMedian calculates the median of a list of integers.
func TestMedian(t *testing.T) {
	// Test case 1: Odd number of values.
	numbers1 := []int{3, 5, 7, 1, 9, 2}
	median1 := Median(numbers1...)
	expected1 := 4.0
	if median1 != expected1 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers1, expected1, median1)
	}

	// Test case 2: Even number of values.
	numbers2 := []int{4, 6, 2, 8}
	median2 := Median(numbers2...)
	expected2 := 5.0
	if median2 != expected2 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers2, expected2, median2)
	}

	// Test case 3: Empty slice.
	var numbers3 []int
	median3 := Median(numbers3...)
	expected3 := 0.0
	if median3 != expected3 {
		t.Errorf("Median of %v: expected %f, but got %f",
			numbers3, expected3, median3)
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

	// Iterate over each test case.
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

	// Iterate over each test case.
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

	// Iterate over each test case.
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
	// Define test cases.
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

	// Iterate over each test case.
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
	// Define test cases.
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
	// Define test cases.
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

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if output is as expected.
			if got := Zip(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAbs tests the Abs function.
func TestAbs(t *testing.T) {
	// Test positive integer.
	n1 := 5
	expected1 := 5
	result1 := Abs(n1)
	if result1 != expected1 {
		t.Errorf("Abs of positive integer: expected %v, but got %v",
			expected1, result1)
	}

	// Test negative integer.
	n2 := -8
	expected2 := 8
	result2 := Abs(n2)
	if result2 != expected2 {
		t.Errorf("Abs of negative integer: expected %v, but got %v",
			expected2, result2)
	}

	// Test positive floating-point number.
	n3 := 3.14
	expected3 := 3.14
	result3 := Abs(n3)
	if result3 != expected3 {
		t.Errorf("Psitive floating-point number: expected %v, but got %v",
			expected3, result3)
	}

	// Test negative floating-point number.
	n4 := -2.718
	expected4 := 2.718
	result4 := Abs(n4)
	if result4 != expected4 {
		t.Errorf("Negative floating-point number: expected %v, but got %v",
			expected4, result4)
	}

	// Test zero value.
	n5 := 0
	expected5 := 0
	result5 := Abs(n5)
	if result5 != expected5 {
		t.Errorf("Abs of zero value: expected %v, but got %v",
			expected5, result5)
	}
}

func TestAbs_Generic(t *testing.T) {
	// Test positive integer.
	n1 := 5
	expected1 := 5
	result1 := Abs(n1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Abs of positive integer (G): expected %v, but got %v",
			expected1, result1)
	}

	// Test negative integer.
	n2 := -8
	expected2 := 8
	result2 := Abs(n2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Abs of negative integer (G): expected %v, but got %v",
			expected2, result2)
	}

	// Test positive floating-point number.
	n3 := 3.14
	expected3 := 3.14
	result3 := Abs(n3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Positive floating-point number (G): expected %v, but got %v",
			expected3, result3)
	}

	// Test negative floating-point number.
	n4 := -2.718
	expected4 := 2.718
	result4 := Abs(n4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Negative floating-point number (G): expected %v, but got %v",
			expected4, result4)
	}

	// Test zero value.
	n5 := 0
	expected5 := 0
	result5 := Abs(n5)
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Abs of zero value (G): expected %v, but got %v",
			expected5, result5)
	}
}

func TestIsEven(t *testing.T) {
	// Test cases for integer values.
	// Expected results: true if the value is even, false otherwise.
	tests := []struct {
		value    float64
		expected bool
	}{
		{4, true},
		{3, false},
		{-6, true},
		{-5, false},
	}

	for _, test := range tests {
		result := IsEven(test.value)
		if result != test.expected {
			t.Errorf("IsEven(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	// Expected results: true if the integer part is even, false otherwise.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{4.2, false, []bool{}},
		{4.2, true, []bool{true}},
		{3.8, false, []bool{true, false}},
		{-5.5, false, []bool{}},
		{-6.5, true, []bool{true}},
	}

	for _, test := range testsFloat {
		result := IsEven(test.value, test.f...)
		if result != test.expected {
			t.Errorf("IsEven(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

func TestIsOdd(t *testing.T) {
	// Test cases for integer values.
	// Expected results: true if the value is odd, false otherwise.
	tests := []struct {
		value    int
		expected bool
	}{
		{4, false},   // even number
		{3, true},    // odd number
		{-6, false},  // negative even number
		{-5, true},   // negative odd number
		{-10, false}, // negative even number
	}

	for _, test := range tests {
		result := IsOdd(test.value)
		if result != test.expected {
			t.Errorf("IsOdd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values.
	// Expected results: true if the integer part is odd, false otherwise.
	testsFloat := []struct {
		value    float64
		expected bool
		f        []bool
	}{
		{3.2, false, []bool{}},
		{3.2, true, []bool{true}},
		{3.0, true, []bool{}},
		{-5.5, true, []bool{true}},
		{-6.5, false, []bool{true, false}},
		{-11.0, true, []bool{false}},
	}

	for _, test := range testsFloat {
		result := IsOdd(test.value, test.f...)
		if result != test.expected {
			t.Errorf("IsOdd(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

func TestIsWhole(t *testing.T) {
	// Test cases for integer values
	// Expected results: true if the value has no fractional part,
	// false otherwise
	tests := []struct {
		value    int
		expected bool
	}{
		{4, true},
		{3, true},
		{-6, true},
		{-5, true},
		{-10, true},
		{0, true},
		{100, true},
		{5, true},
		{5, true},
		{-2, true},
	}

	for _, test := range tests {
		result := IsWhole(test.value)
		if result != test.expected {
			t.Errorf("IsWhole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}

	// Test cases for floating-point values
	// Expected results: true if the value has no fractional part,
	// false otherwise
	testsFloat := []struct {
		value    float64
		expected bool
	}{
		{4.2, false},
		{3.8, false},
		{-5.5, false},
		{-11.0, true},
		{0.0, true},
		{100.0, true},
		{5.0, true},
		{-2.0, true},
		{-2.1, false},
	}

	for _, test := range testsFloat {
		result := IsWhole(test.value)
		if result != test.expected {
			t.Errorf("IsWhole(%v): expected %v, but got %v",
				test.value, test.expected, result)
		}
	}
}

// TestRank tests the Rank function.
func TestRank(t *testing.T) {
	// Test case 1: Ascending order, value found
	rank := Rank(7, []int{1, 5, 2, 3, 7, 8})
	if rank != 5 {
		t.Errorf("Expected rank 5, but got %d", rank)
	}

	// Test case 2: Ascending order, value not found
	rank = Rank(9, []int{1, 5, 2, 3, 7, 8})
	if rank != 0 {
		t.Errorf("Expected rank 0, but got %d", rank)
	}

	// Test case 3: Descending order, value found
	rank = Rank(7, []int{1, 5, 2, 3, 7, 8}, true)
	if rank != 2 {
		t.Errorf("Expected rank 2, but got %d", rank)
	}

	// Test case 4: Descending order, value not found
	rank = Rank(9, []int{1, 5, 2, 3, 7, 8}, true)
	if rank != 0 {
		t.Errorf("Expected rank 0, but got %d", rank)
	}

	// Test case 5: Ascending order, float values
	rank = Rank(4.5, []float64{1.2, 3.1, 4.5, 2.8, 4.5, 6.7})
	if rank != 3 {
		t.Errorf("Expected rank 3, but got %d", rank)
	}

	// Test case 6: Descending order, float values
	rank = Rank(4.5, []float64{1.2, 3.1, 4.5, 2.8, 4.5, 6.7}, true)
	if rank != 2 {
		t.Errorf("Expected rank 2, but got %d", rank)
	}
}

// TestIsNumber tests the IsNumber function.
func TestIsNumber(t *testing.T) {
	// Numeric values
	num := 10
	if !IsNumber(num) {
		t.Errorf("%v should be a number", num)
	}

	floatNum := 3.14
	if !IsNumber(floatNum) {
		t.Errorf("%v should be a number", floatNum)
	}

	// Non-numeric values
	str := "hello"
	if IsNumber(str) {
		t.Errorf("%v should not be a number", str)
	}

	boolean := true
	if IsNumber(boolean) {
		t.Errorf("%v should not be a number", boolean)
	}

	slice := []int{1}
	if IsNumber(slice) {
		t.Errorf("%v should not be a number", slice)
	}
}

// TestIsPointer tests the IsPointer function.
func TestIsPointer(t *testing.T) {
	str := "hello"
	ptr := &str
	if !IsPointer(ptr) {
		t.Errorf("%v should be a pointer", ptr)
	}

	if IsPointer(str) {
		t.Errorf("%v should not be a pointer", str)
	}

	num := 10
	if IsPointer(num) {
		t.Errorf("%v should not be a pointer", num)
	}

	var nilPtr *string
	if !IsPointer(nilPtr) {
		t.Errorf("%v should be a pointer", nilPtr)
	}
}

// TestHLookup tests the HLookup function.
func TestHLookup(t *testing.T) {
	lookup := []string{"A", "B", "C"}
	result := []int{1, 2, 3}

	// Test case 1: Lookup value exists in the table.
	val := HLookup("B", lookup, result, -1)
	expected := 2
	if val != expected {
		t.Errorf("HLookup test case 1 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 2: Lookup value does not exist in the table.
	val = HLookup("D", lookup, result, -1)
	expected = -1
	if val != expected {
		t.Errorf("HLookup test case 2 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 3: Lookup value is the first element in the table.
	val = HLookup("A", lookup, result, -1)
	expected = 1
	if val != expected {
		t.Errorf("HLookup test case 3 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 4: Lookup value is the last element in the table.
	val = HLookup("C", lookup, result, -1)
	expected = 3
	if val != expected {
		t.Errorf("HLookup test case 4 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 5: Lookup value is not found,
	// with a non-default default value.
	val = HLookup("E", lookup, result, 0)
	expected = 0
	if val != expected {
		t.Errorf("HLookup test case 5 failed: expected %d, got %d",
			expected, val)
	}
}

// TestVLookup tests the VLookup function.
func TestVLookup(t *testing.T) {
	lookup := []string{"A", "B", "C"}
	result := []int{1, 2, 3}

	// Test case 1: Lookup value exists in the table.
	val := VLookup("B", lookup, result, -1)
	expected := 2
	if val != expected {
		t.Errorf("VLookup test case 1 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 2: Lookup value does not exist in the table.
	val = VLookup("D", lookup, result, -1)
	expected = -1
	if val != expected {
		t.Errorf("VLookup test case 2 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 3: Lookup value is the first element in the table.
	val = VLookup("A", lookup, result, -1)
	expected = 1
	if val != expected {
		t.Errorf("VLookup test case 3 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 4: Lookup value is the last element in the table.
	val = VLookup("C", lookup, result, -1)
	expected = 3
	if val != expected {
		t.Errorf("VLookup test case 4 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 5: Lookup value is not found,
	// with a non-default default value.
	val = VLookup("E", lookup, result, 0)
	expected = 0
	if val != expected {
		t.Errorf("VLookup test case 5 failed: expected %d, got %d",
			expected, val)
	}
}
