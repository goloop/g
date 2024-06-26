package g

import (
	"testing"

	"github.com/goloop/trit"
)

// TestIf tests the If function.
func TestIf(t *testing.T) {
	testsBool := []struct {
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

	for _, tt := range testsBool {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.cond, tt.trueVal, tt.falseVal); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}

	testsTrit := []struct {
		name     string
		cond     trit.Trit
		trueVal  interface{}
		falseVal interface{}
		want     interface{}
	}{
		{"Trit true, returns trueVal", trit.True, "pass", "fail", "pass"},
		{"Trit false, returns falseVal", trit.False, "pass", "fail", "fail"},
		{"Trit unknown, returns falseVal", trit.False, "pass", "fail", "fail"},
	}

	for _, tt := range testsTrit {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.cond, tt.trueVal, tt.falseVal); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIfTrit tests the If function for Trit type.
func TestIfTrit(t *testing.T) {
	tests := []struct {
		name       string
		cond       trit.Trit
		trueVal    interface{}
		falseVal   interface{}
		unknownVal []interface{}
		want       interface{}
	}{
		{
			name:       "Trit unknown, returns unknownVal",
			cond:       trit.Unknown,
			trueVal:    "pass",
			falseVal:   "fail",
			unknownVal: []interface{}{"unknown"},
			want:       "unknown",
		},
		{
			name:       "Trit unknown, returns falseVal",
			cond:       trit.Unknown,
			trueVal:    "pass",
			falseVal:   "fail",
			unknownVal: []interface{}{}, // hasn't unknownVal
			want:       "fail",
		},
		{
			name:       "Trit unknown, returns falseVal",
			cond:       trit.Unknown,
			trueVal:    "pass",
			falseVal:   "fail",
			unknownVal: []interface{}{}, // hasn't unknownVal
			want:       "fail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(
				tt.cond,
				tt.trueVal,
				tt.falseVal,
				tt.unknownVal...,
			); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAllBool tests the All function with bool.
func TestAllBool(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	tests := []struct {
		name string
		v    []bool
		want bool
	}{
		{
			name: "Empty",
			v:    []bool{},
			want: false,
		},
		{
			name: "One true",
			v:    []bool{true},
			want: true,
		},
		{
			name: "Many true",
			v:    []bool{true, true, true},
			want: true,
		},
		{
			name: "Mix true and false",
			v:    []bool{true, false, true},
			want: false,
		},
		{
			name: "Many false",
			v:    []bool{false, false, false},
			want: false,
		},
		{
			name: "One false",
			v:    []bool{false},
			want: false,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `All` and check if the output is as expected.
			if got := All(tt.v...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAllBoolList tests the All function with bool lits.
func TestAllBoolList(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	tests := []struct {
		name string
		v    [][]bool
		want bool
	}{
		{
			name: "Empty list",
			v:    [][]bool{},
			want: false,
		},
		{
			name: "One true list",
			v:    [][]bool{{true, false}},
			want: true,
		},
		{
			name: "Many true list",
			v:    [][]bool{{true, true}, {true}},
			want: true,
		},
		{
			name: "Mix true and false",
			v:    [][]bool{{}, {true, false, true}},
			want: false,
		},
		{
			name: "Many false",
			v:    [][]bool{{}, {}},
			want: false,
		},
		{
			name: "One false",
			v:    [][]bool{{}},
			want: false,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `All` and check if the output is as expected.
			if got := All(tt.v...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAll tests the All function.
func TestAll(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	minLoadPerGoroutine = 20
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"Trit as True", []interface{}{trit.True, trit.Define(1)}, true},
		{"Trit as False", []interface{}{trit.True, trit.Define(-1)}, false},
		{"All non-zero values", []interface{}{1, true, "test"}, true},
		{"One zero value", []interface{}{1, 0, "test"}, false},
		{"All zero values", []interface{}{0, false, ""}, false},
		{"Empty input", []interface{}{}, false},
		{"Big size with ctx cancel", []interface{}{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, trit.Unknown, 1, 1, 1, 1, 1, 1, // zerro here
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, false},
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

// TestAnyBool tests the Any function with bool.
func TestAnyBool(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	minLoadPerGoroutine = 20
	tests := []struct {
		name string
		v    []bool
		want bool
	}{
		{
			name: "Empty",
			v:    []bool{},
			want: false,
		},
		{
			name: "One true",
			v:    []bool{true},
			want: true,
		},
		{
			name: "Many true",
			v:    []bool{true, true, true},
			want: true,
		},
		{
			name: "Mix true and false",
			v:    []bool{true, false, true},
			want: true,
		},
		{
			name: "Many false",
			v:    []bool{false, false, false},
			want: false,
		},
		{
			name: "One false",
			v:    []bool{false},
			want: false,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `Any` and check if the output is as expected.
			if got := Any(tt.v...); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAnyBoolList tests the Any function with bool lits.
func TestAnyBoolList(t *testing.T) {
	// Define a list of test cases as anonymous struct.
	tests := []struct {
		name string
		v    [][]bool
		want bool
	}{
		{
			name: "Empty list",
			v:    [][]bool{},
			want: false,
		},
		{
			name: "One true list",
			v:    [][]bool{{true, false}},
			want: true,
		},
		{
			name: "Many true list",
			v:    [][]bool{{true, true}, {true}},
			want: true,
		},
		{
			name: "Mix true and false",
			v:    [][]bool{{}, {true, false, true}},
			want: true,
		},
		{
			name: "Many false",
			v:    [][]bool{{}, {}},
			want: false,
		},
		{
			name: "One false",
			v:    [][]bool{{}},
			want: false,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `Any` and check if the output is as expected.
			if got := Any(tt.v...); got != tt.want {
				t.Errorf("any() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAny tests the Any function.
func TestAny(t *testing.T) {
	// Define test cases.
	minLoadPerGoroutine = 20
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"Trit as True", []interface{}{trit.False, trit.Define(1)}, true},
		{"Trit as False", []interface{}{trit.False, trit.Define(-1)}, false},
		{"All non-zero values", []interface{}{1, true, "test"}, true},
		{"One non-zero value", []interface{}{0, false, "test"}, true},
		{"All zero values", []interface{}{0, false, ""}, false},
		{"Empty input", []interface{}{}, false},
		{"Big size with ctx cancel", []interface{}{
			trit.False, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, trit.Unknown, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // no zero here
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}, true},
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

// TestIsEmpty tests the IsEmpty function.
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"Trint True", trit.True, false},
		{"Trint False", trit.False, false},
		{"Trint Unknown", trit.Unknown, true}, // only trit.Unknown is empty
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
		{"Empty slice", []bool{}, true},
		{"Not empty slice", []bool{false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.v); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIsFalse tests the IsFalse function.
func TestIsFalse(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "Value is nil",
			v:    nil,
			want: true,
		},
		{
			name: "Value is zero",
			v:    0,
			want: true,
		},
		{
			name: "Value is Trit with true state",
			v:    trit.True,
			want: false,
		},
		{
			name: "Value is Trit with false state",
			v:    trit.False,
			want: true,
		},
		{
			name: "Value is Trit with unknown state",
			v:    trit.Unknown,
			want: true,
		},
		{
			name: "Value is empty slice",
			v:    []bool{},
			want: true,
		},
		{
			name: "Value is not empty slice",
			v:    []bool{false},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFalse(tt.v); got != tt.want {
				t.Errorf("IsFalse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIsTrue tests the IsTrue function.
func TestIsTrue(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "value is nil",
			v:    nil,
			want: false,
		},
		{
			name: "Value is zero",
			v:    0,
			want: false,
		},
		{
			name: "Value is Trit with true state",
			v:    trit.True,
			want: true,
		},
		{
			name: "Value is Trit with false state",
			v:    trit.False,
			want: false,
		},
		{
			name: "Value is Trit with unknown state",
			v:    trit.Unknown,
			want: false,
		},
		{
			name: "Value is empty slice",
			v:    []bool{},
			want: false,
		},
		{
			name: "Value is not empty slice",
			v:    []bool{false},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrue(tt.v); got != tt.want {
				t.Errorf("IsTrue() = %v, want %v", got, tt.want)
			}
		})
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

// TestIsNumber tests the IsNumber function.
func TestIsNumber(t *testing.T) {
	// Numeric values.
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

// TestAnyList tests the AnyList function.
func TestAnyList(t *testing.T) {
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"All false", []interface{}{false, false, false}, false},
		{"One true", []interface{}{false, true, false}, true},
		{"All true", []interface{}{true, true, true}, true},
		{"All false and true", []interface{}{true, false, true}, true},
		{"Empty slice", []interface{}{}, false},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `AnyList` and check
			// if the output is as expected.
			if got := AnyList(tt.v); got != tt.want {
				t.Errorf("AnyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestAllList tests the AllList function.
func TestAllList(t *testing.T) {
	tests := []struct {
		name string
		v    []interface{}
		want bool
	}{
		{"All false", []interface{}{false, false, false}, false},
		{"One true", []interface{}{false, true, false}, false},
		{"All true", []interface{}{true, true, true}, true},
		{"All false and true", []interface{}{true, false, true}, false},
		{"Empty slice", []interface{}{}, false},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function `AllList` and check
			// if the output is as expected.
			if got := AllList(tt.v); got != tt.want {
				t.Errorf("AllList() = %v, want %v", got, tt.want)
			}
		})
	}
}
