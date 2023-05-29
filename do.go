// The do package is a utility library for Go 1.20+ that provides generic
// and type-safe helper functions to simplify software development tasks.
// The package uses Go's generics extensively and offers an assortment of
// functions such as simplified conditionals, zero-value checks, numeric
// operations, functional programming helpers, element operations, and
// pairing functionality.
package do

import "reflect"

// Pair is a generic struct type with two fields: First and Second.
// Used as the result element of the Zip function and can contain
// a pair of values of any T and U types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Numerable is an interface type that is satisfied by all numeric types
// in Go, both integer and floating point. This includes int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, uint64, float32, and float64.
// It allows functions to operate on any of these types where numerical
// operations such as addition, subtraction, multiplication, and division
// are needed. It enables generic programming techniques for numeric types.
type Numerable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// IsEmpty checks if the  v of any type T is "zero value" for that type.
//
// Zero values in Go are values that the variables of respective types hold
// upon their declaration, if they do not have any explicit initialization.
//
// For example, zero value of type int is 0, for type float64 is 0.0,
// for a pointer is nil, for a string is "", for a boolean is false, etc.
//
// This function is generic and can work with any type T.
func IsEmpty[T any](v T) bool {
	t := reflect.TypeOf(v)
	if t == nil {
		return true
	}
	zero := reflect.Zero(t).Interface()
	return reflect.DeepEqual(v, zero)
}

// All returns true if all values in the provided slice
// are not zero values for their types.
//
// If at least one value is a zero value, it immediately returns false.
// If the slice is empty, it returns false.
//
// This function is generic and can work with any type T.
func All[T any](v ...T) bool {
	if len(v) == 0 {
		return false
	}

	for _, val := range v {
		if IsEmpty(val) {
			return false
		}
	}

	return true
}

// Any returns true if at least one value in the provided slice
// is not a zero value for its type.
//
// As soon as it finds a value that is not a zero value, it returns true.
// If all values in the slice are zero values or the slice is empty,
// it returns false.
//
// This function is generic and can work with any type T.
func Any[T any](v ...T) bool {
	for _, val := range v {
		if !IsEmpty(val) {
			return true
		}
	}

	return false
}

// Value returns the first non-zero value from the parameters.
//
// The function checks each parameter in the order they are passed to
// the function and returns the first parameter that is not a zero value.
// If all parameters are zero values, the function returns a zero value
// for the type T of the parameters.
//
// This function is generic and can work with any type T.
func Value[T any](v T, more ...T) T {
	if !IsEmpty(v) || len(more) == 0 {
		return v
	}

	for _, val := range more {
		if !IsEmpty(val) {
			return val
		}
	}

	return v
}

// If is a substitute for the ternary operator (?:)
// which is not available in Go.
//
// In languages like C/C++ and Python, you can use a ternary
// operator for a concise conditional expression:
//
// C/C++:  int max = (a > b) ? a : b;
// Python: max = a if a > b else b
//
// It takes three parameters: a boolean expression e, and two
// values of any type T (t and f). If the expression e is true,
// it returns t, otherwise it returns f.
//
// Example:
//
//	var a, b float64 // or int, or any other type
//	...
//	max := do.If(a > b, a, b)
//
// This function is generic and can work with any type T.
func If[T any](e bool, t, f T) T {
	if e {
		return t
	}

	return f
}

// Min returns the smallest value among all input values.
//
// This function requires at least one parameter of a type that
// satisfies the Numerable interface. Additional values can be
// passed using variadic arguments.
//
// The function iterates through all the passed values
// and returns the smallest one. The type must be Numerable
// and support the less than (<) operator.
//
// Example usage:
//
//	n := []int{3,5,7,1,9,2}
//	min := do.Min(n[0], n[1:]...)
//
// This function is generic and can work with any type T.
func Min[T Numerable](v T, more ...T) T {
	min := v
	for _, val := range more {
		if val < min {
			min = val
		}
	}

	return min
}

// Max returns the largest value among all input values.
//
// This function requires at least one parameter of a type that
// satisfies the Numerable interface. Additional values can be
// passed using variadic arguments.
//
// The function iterates through all the passed values
// and returns the largest one. The type must be Numerable
// and support the greater than (>) operator.
//
// Example usage:
//
//	n := []int{3,5,7,1,9,2}
//	max := do.Max(n[0], n[1:]...)
//
// This function is generic and can work with any type T.
func Max[T Numerable](v T, more ...T) T {
	max := v
	for _, val := range more {
		if val > max {
			max = val
		}
	}

	return max
}

// Sum returns the sum of all values.
//
// Note: this function does not handle overflow. If the sum of the input
// values exceeds the maximum value that can be stored in type T, the
// result will wrap around, due to how Go handles overflow.
//
// Example usage:
//
//	n := []int{3,5,7,1,9,2}
//	sum := do.Sum(n...)
//
// This function is generic and can work with any type T.
func Sum[T Numerable](v ...T) T {
	if len(v) == 0 {
		return 0
	}

	sum := v[0]
	for _, val := range v[1:] {
		sum += val
	}

	return sum
}

// Average calculates the average of a variable number
// of values of type Numerable.
//
// It first computes the sum of all the values, and then divides
// by the number of values to get the average.

// If no values are provided, it returns 0.
// Note: this function returns the average as a float64,
// regardless of the input type.
//
// Examples:
//
//	n := []int{3,5,7,1,9,2}
//	avg := do.Average(n...)
//
// This function is generic and can work with any type T.
func Average[T Numerable](v ...T) float64 {
	if len(v) == 0 {
		return 0
	}

	sum := float64(Sum(v...))
	return sum / float64(len(v))
}

// Map applies a function to all items in an input slice and returns
// a new slice with the transformed items.
//
// T is the type of items in the input slice, and U is the type of
// items in the output slice.
// The function f takes an item of type T and returns a new item of type U.
func Map[T any, U any](vs []T, f func(T) U) []U {
	result := make([]U, len(vs))
	for i, v := range vs {
		result[i] = f(v)
	}

	return result
}

// Filter applies a predicate function to all items in an input
// slice and returns a new slice with the items for which the
// predicate function returns true.
//
// T is the type of items in the input slice.
// The predicate function f takes an item of type T and returns a boolean.
func Filter[T any](vs []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

// Reduce takes a slice, an initial value and a fold function as input,
// and it returns a single value that results from applying the fold
// function to each item in the slice in order.
//
// T is the type of items in the slice.
// U is the type of the output and the initial value.
//
// The fold function f takes two parameters - an accumulator of type U
// and an item of type T, and returns a new accumulator.
func Reduce[T any, U any](vs []T, f func(U, T) U, init U) U {
	result := init
	for _, v := range vs {
		result = f(result, v)
	}

	return result
}

// Contains checks if a slice contains a specific element.
//
// It takes a slice and an element of the same type as input, and returns
// a boolean value indicating whether the element is found in the slice.
//
// T is the type of items in the slice and the element to be searched.
// The function returns true if the element is found, and false otherwise.
func Contains[T comparable](vs []T, v T) bool {
	for _, item := range vs {
		if item == v {
			return true
		}
	}

	return false
}

// Index returns the index of the first occurrence of a specific
// element in a slice, or -1 if the element is not present.
//
// T is the type of the items in the slice and the element to be searched.
// The function returns an integer indicating the position of the first
// occurrence of the element, or -1 if the element is not found.
func Index[T comparable](vs []T, v T) int {
	for i, item := range vs {
		if item == v {
			return i
		}
	}

	return -1
}

// Zip takes two slices and returns a slice of pairs.
// Each pair contains an item from each of the two slices,
// in the same position. The length of the returned slice is
// equal to the minimum of the lengths of the two input slices.
//
// If one slice is shorter than the other, the extra elements of
// the longer slice are ignored.
//
// The function returns a slice of struct where each struct has two fields:
// First and Second, representing the elements from the first and second
// slices respectively.
func Zip[T, U any](a []T, b []U) []Pair[T, U] {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	result := make([]Pair[T, U], minLength)
	for i := 0; i < minLength; i++ {
		result[i] = struct {
			First  T
			Second U
		}{a[i], b[i]}
	}

	return result
}
