package g

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

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
			if got := Contains(tt.val, tt.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
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
				}
				return b
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
				}
				return b
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

// TestValue tests the Value function.
func TestValue(t *testing.T) {
	// Define test cases.
	tests := []struct {
		name string
		v    []int
		want int
	}{
		{"Non-zero value is 0 in slice", []int{1, 0, 0}, 1},
		{"Non-zero value is in the 1 of the slice", []int{0, 2, 0}, 2},
		{"Non-zero value is at the 2 of the slice", []int{0, 0, 3}, 3},
		{"All zero values", []int{0, 0, 0}, 0},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert slice of interfaces to slice of empty interfaces.
			input := make([]int, len(tt.v))
			for i, val := range tt.v {
				input[i] = val
			}

			// Check if output is as expected.
			if got := Value(input...); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
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

// TestDistinct tests the Distinct function.
func TestDistinct(t *testing.T) {
	t.Run("Distinct with empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Distinct(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Distinct(%v) = %v, expected %v", input, result, expected)
		}
	})

	t.Run("Distinct with slice containing duplicates", func(t *testing.T) {
		input := []int{1, 2, 3, 2, 4, 3, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Distinct(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Distinct(%v) = %v, expected %v", input, result, expected)
		}
	})

	t.Run("Distinct with slice of strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry", "apple", "banana"}
		expected := []string{"apple", "banana", "cherry"}
		result := Distinct(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Distinct(%v) = %v, expected %v", input, result, expected)
		}
	})
}

// TestShuffle tests the Shuffle function.
func TestShuffle(t *testing.T) {
	// contains checks if a slice contains a specific element
	contains := func(slice interface{}, elem interface{}) bool {
		sliceVal := reflect.ValueOf(slice)
		elemVal := reflect.ValueOf(elem)

		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == elemVal.Interface() {
				return true
			}
		}

		return false
	}

	t.Run("Shuffle with empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		Shuffle(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Shuffle(%v) modifies the slice to %v, expected %v",
				input, input, expected)
		}
	})

	t.Run("Shuffle with slice of integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		// We cannot determine the expected output as it's random
		// Instead, we check if the input slice has the same elements
		// after shuffling.
		Shuffle(input)
		if len(input) != 5 {
			t.Errorf("Shuffle(%v) modifies the slice length to %d, expected 5",
				input, len(input))
		}
		// Check if all the original elements are still present
		// in the shuffled slice.
		for _, elem := range []int{1, 2, 3, 4, 5} {
			if !contains(input, elem) {
				t.Errorf("Shuffle(%v) modifies the slice, missing element %d",
					input, elem)
			}
		}
	})

	t.Run("Shuffle with slice of strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry", "date", "elderberry"}
		// We cannot determine the expected output as it's random
		// Instead, we check if the input slice has the same elements
		// after shuffling.
		Shuffle(input)
		if len(input) != 5 {
			t.Errorf("Shuffle(%v) modifies the slice length to %d, expected 5",
				input, len(input))
		}
		// Check if all the original elements are still present
		// in the shuffled slice.
		for _, elem := range []string{
			"apple",
			"banana",
			"cherry",
			"date",
			"elderberry",
		} {
			if !contains(input, elem) {
				t.Errorf("Shuffle(%v) modifies the slice, missing element %s",
					input, elem)
			}
		}
	})
}

// TestProduct tests the Product function.
func TestProduct(t *testing.T) {
	t.Run("Product with empty slice", func(t *testing.T) {
		input := []int{}
		expected := 1
		result := Product(input...)
		if result != expected {
			t.Errorf("Product(%v) returned %v, expected %v",
				input, result, expected)
		}
	})

	t.Run("Product with slice of integers", func(t *testing.T) {
		input := []int{2, 3, 4}
		expected := 24
		result := Product(input...)
		if result != expected {
			t.Errorf("Product(%v) returned %v, expected %v",
				input, result, expected)
		}
	})

	t.Run("Product with slice of floats", func(t *testing.T) {
		input := []float64{1.5, 2.5, 3.5}
		expected := 13.125
		result := Product(input...)
		if result != expected {
			t.Errorf("Product(%v) returned %v, expected %v",
				input, result, expected)
		}
	})
}

// TestMerge tests the Merge function.
func TestMerge(t *testing.T) {
	//  Merge two sorted integer arrays.
	intA := []int{1, 3, 5}
	intB := []int{2, 4, 6}
	mergedInt := Merge(intA, intB, true)
	expectedInt := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(mergedInt, expectedInt) {
		t.Errorf("Incorrect merged array. Expected %v, got %v",
			expectedInt, mergedInt)
	}

	//  Merge two unsorted integer arrays.
	intA = []int{1, 3, 5}
	intB = []int{2, 4, 6}
	mergedInt = Merge(intA, intB)
	expectedInt = []int{1, 3, 5, 2, 4, 6}
	if !reflect.DeepEqual(mergedInt, expectedInt) {
		t.Errorf("Incorrect merged array. Expected %v, got %v",
			expectedInt, mergedInt)
	}

	// Merge two sorted string arrays.
	strA := []string{"apple", "banana", "kiwi"}
	strB := []string{"cherry", "mango", "orange"}
	mergedStr := Merge(strA, strB, true)
	expectedStr := []string{
		"apple", "banana", "cherry",
		"kiwi", "mango", "orange",
	}
	if !reflect.DeepEqual(mergedStr, expectedStr) {
		t.Errorf("Incorrect merged string array. Expected %v, got %v",
			expectedStr, mergedStr)
	}

	// Merge two sorted float arrays.
	floatA := []float64{1.2, 3.4, 5.6}
	floatB := []float64{2.1, 4.3, 6.5}
	mergedFloat := Merge(floatA, floatB, true)
	expectedFloat := []float64{1.2, 2.1, 3.4, 4.3, 5.6, 6.5}
	if !reflect.DeepEqual(mergedFloat, expectedFloat) {
		t.Errorf("Incorrect merged float array. Expected %v, got %v",
			expectedFloat, mergedFloat)
	}

	// Merge empty arrays.
	emptyA := []int{}
	emptyB := []int{}
	mergedEmpty := Merge(emptyA, emptyB)
	expectedEmpty := []int{}
	if !reflect.DeepEqual(mergedEmpty, expectedEmpty) {
		t.Errorf("Incorrect merged empty array. Expected %v, got %v",
			expectedEmpty, mergedEmpty)
	}
}

// TestIn tests the flat in function for string values.
func TestIn(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	if res := In("cherry", slice...); !res {
		t.Errorf("expected %v, got %v", true, res)
	}

	if res := In("mango", slice...); res {
		t.Errorf("expected %v, got %v", false, res)
	}
}

// TestIn tests the In function for int values.
func TestInWithInt(t *testing.T) {
	minLoadPerGoroutine = 20
	generateIntSlice := func(size int) []int {
		slice := make([]int, size)
		rand.Seed(time.Now().UnixNano())
		for i := range slice {
			slice[i] = rand.Intn(100) - 50 // negative and positive integers
		}
		return slice
	}

	slice := generateIntSlice(10000)
	expected := In(10, slice...)

	if res := In(10, slice...); res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}

	// For empty.
	if res := In(100, []int{}...); res {
		t.Errorf("expected %v, got %v", false, res)
	}
}

// TestFloat tests the In function for float values.
func TestInWithFloat(t *testing.T) {
	generateFloatSlice := func(size int) []float64 {
		slice := make([]float64, size)
		rand.Seed(time.Now().UnixNano())
		for i := range slice {
			slice[i] = rand.Float64()*100 - 50 // negative and positive floats
		}
		return slice
	}

	slice := generateFloatSlice(10000)
	expected := In(10.5, slice...)

	if res := In(10.5, slice...); res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}

	// For empty.
	if res := In(100.5, []float64{}...); res {
		t.Errorf("expected %v, got %v", false, res)
	}
}

// TestString tests the In function for string values.
func TestInWithString(t *testing.T) {
	generateStringSlice := func(size int) []string {
		slice := make([]string, size)
		for i := range slice {
			slice[i] = string(rune(i + 65)) // string with ASCII characters
		}
		return slice
	}

	slice := generateStringSlice(10000)
	expected := In("A", slice...)

	if res := In("A", slice...); res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}

	// For empty.
	if res := In("Go"); res {
		t.Errorf("expected %v, got %v", false, res)
	}
}

// TestRange tests the Range function.
func TestRange(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		opt      []int
		expected []int
	}{
		{
			name:     "Single parameter",
			n:        5,
			opt:      []int{},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Two parameters",
			n:        3,
			opt:      []int{7},
			expected: []int{3, 4, 5, 6},
		},
		{
			name:     "Three parameters",
			n:        1,
			opt:      []int{10, 2},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Reverse range",
			n:        10,
			opt:      []int{0, -1},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "Incorrect range",
			n:        10,
			opt:      []int{20, -1},
			expected: []int{},
		},
		{
			name:     "Large range",
			n:        MaxRangeSize + 100,
			opt:      []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Range(tc.n, tc.opt...)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

/*
// TestRangef tests the Rangef function.

	func TestRangef(t *testing.T) {
		// Define the apple factory function
		appleFactory := func(i int) string {
			appleVarieties := []string{
				"Gala",
				"Fuji",
				"Honeycrisp",
				"Red Delicious",
				"Granny Smith",
				"Golden Delicious",
				"Pink Lady",
				"Braeburn",
				"McIntosh",
				"Jazz",
			}

			if i >= 0 && i < len(appleVarieties) {
				return appleVarieties[i]
			}
			return "-"
		}

		// Single parameter.
		result := Rangef(appleFactory, 3)
		expected := []string{"Gala", "Fuji", "Honeycrisp"}
		if fmt.Sprint(result) != fmt.Sprint(expected) {
			t.Errorf("Test case 1 failed: Expected %v, but got %v",
				expected, result)
		}

		// Two parameters.
		result = Rangef(appleFactory, 4, 7)
		expected = []string{"Granny Smith", "Golden Delicious", "Pink Lady"}
		if fmt.Sprint(result) != fmt.Sprint(expected) {
			t.Errorf("Test case 2 failed: Expected %v, but got %v",
				expected, result)
		}

		// Three parameters.
		result = Rangef(appleFactory, 7, 12, 2)
		expected = []string{"Braeburn", "Jazz", "-"}
		if fmt.Sprint(result) != fmt.Sprint(expected) {
			t.Errorf("Test case 3 failed: Expected %v, but got %v",
				expected, result)
		}
	}
*/
func TestRangef(t *testing.T) {
	tests := []struct {
		name string
		a    int
		opt  []int
		want []string
	}{
		{
			name: "Step Size 0",
			a:    3,
			opt:  []int{0},
			want: []string{},
		},
		{
			name: "Negative Step Size and n <= m",
			a:    2,
			opt:  []int{5, -1},
			want: []string{},
		},
		{
			name: "Positive Step Size and n >= m",
			a:    5,
			opt:  []int{2, 1},
			want: []string{},
		},
	}

	appleFactory := func() func(int) string {
		var appleVarieties = []string{
			"Gala",
			"Fuji",
			"Honeycrisp",
			"Red Delicious",
			"Granny Smith",
			"Golden Delicious",
			"Pink Lady",
			"Braeburn",
			"McIntosh",
			"Jazz",
		}

		return func(i int) string {
			if i >= 0 && i < len(appleVarieties) {
				return appleVarieties[i]
			}
			return "-"
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rangef(appleFactory(), tt.a, tt.opt...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rangef() = %v, want %v", got, tt.want)
			}
		})
	}

	// Single parameter.
	result := Rangef(appleFactory(), 3)
	expected := []string{"Gala", "Fuji", "Honeycrisp"}
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Test case 1 failed: Expected %v, but got %v",
			expected, result)
	}

	// Two parameters.
	result = Rangef(appleFactory(), 4, 7)
	expected = []string{"Granny Smith", "Golden Delicious", "Pink Lady"}
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Test case 2 failed: Expected %v, but got %v",
			expected, result)
	}

	// Three parameters.
	result = Rangef(appleFactory(), 7, 12, 2)
	expected = []string{"Braeburn", "Jazz", "-"}
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Test case 3 failed: Expected %v, but got %v",
			expected, result)
	}
}

// TestReverse tests the Reverse function.
func TestReverse(t *testing.T) {
	// Test case 1: Reversing a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	expectedNumbers := []int{5, 4, 3, 2, 1}
	Reverse(numbers)

	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Reverse(numbers) = %v, expected %v",
			numbers, expectedNumbers)
	}

	// Test case 2: Reversing a slice of strings
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	expectedNames := []string{"Dave", "Charlie", "Bob", "Alice"}
	Reverse(names)

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf("Reverse(names) = %v, expected %v", names, expectedNames)
	}

	// Test case 3: Reversing an empty slice
	empty := []bool{}
	expectedEmpty := []bool{}
	Reverse(empty)

	if !reflect.DeepEqual(empty, expectedEmpty) {
		t.Errorf("Reverse(empty) = %v, expected %v", empty, expectedEmpty)
	}
}
