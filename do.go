// The do package is a utility library for Go 1.20+ that provides generic
// and type-safe helper functions to simplify software development tasks.
// The package uses Go's generics extensively and offers an assortment of
// functions such as simplified conditionals, zero-value checks, numeric
// operations, functional programming helpers, element operations, and
// pairing functionality.
package do

import (
	"math"
	"reflect"
)

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

// Abs returns the absolute value of a numeric input value.
//
// The function takes a value of a type that satisfies the Numerable interface
// and returns its absolute value as the same type.
//
// For numeric types that support the negation operator (-), the function
// uses the negation operator to calculate the absolute value.
// For unsigned integer types, the absolute value is equal to
// the original value.
//
// Example:
//
//	n := -5
//	abs := do.Abs(n) // 5
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func Abs[T Numerable](v T) T {
	if v < 0 {
		return -v
	}
	return v
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

// IsEven checks if a value is an even number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. If the `f` argument is provided
// and set to true, the function ignores the fractional part
// of the value when checking for evenness. For integer types,
// it checks if the value is divisible by 2 without a remainder.
// For floating-point types, it considers only the integer part
// of the value and determines the parity of the integer part.
// If the value has a non-zero fractional part and `f` is true,
// it returns false since an even number cannot have a fractional part.
//
// Examples:
//
//	do.IsEven(4)         // true
//	do.IsEven(3)         // false
//	do.IsEven(4.2)       // false
//	do.IsEven(4.2, true) // false
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func IsEven[T Numerable](v T, f ...bool) bool {
	if All(f...) {
		// Ignore the fact that the number is a float
		// and determine the parity of the left side only.
		return int(v)%2 == 0
	}

	return If(IsWhole(v), int(v)%2 == 0, false)
}

// IsNumber checks if a value is a numeric type.
//
// The function takes an interface{} value `v` and checks if it is
// of a numeric type, including integer and floating-point types.
// It returns true if `v` is a numeric type, and false otherwise.
//
// Example:
//
//	isNum := IsNumber(10)       // true
//	isNum = IsNumber(3.14)      // true
//	isNum = IsNumber("hello")   // false
//	isNum = IsNumber(true)      // false
//	isNum = IsNumber([]int{1})  // false
func IsNumber(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

// IsOdd checks if a value is an odd number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. If the `f` argument is provided
// and set to true, the function ignores the fractional part
// of the value when checking for oddness. For integer types,
// it checks if the value is not divisible by 2 without a remainder.
// For floating-point types, it considers only the integer part
// of the value and determines the parity of the integer part.
// If the value has a non-zero fractional part and `f` is true,
// it returns true since an odd number cannot have a fractional part.
// Otherwise, it returns the negation of the IsEven function.
//
// Examples:
//
//	do.IsOdd(3)         // true
//	do.IsOdd(4)         // false
//	do.IsOdd(5.5)       // false
//	do.IsOdd(5.5, true) // true
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func IsOdd[T Numerable](v T, f ...bool) bool {
	if All(f...) {
		// Ignore the fact that the number is a float
		// and determine the parity of the left side only.
		return int(v)%2 != 0
	}

	return If(IsWhole(v), int(v)%2 != 0, false)
}

// IsPointer checks if a value is a pointer.
//
// The function takes an interface{} value `v` and checks if it is
// a pointer type. It returns true if `v` is a pointer, and false
// otherwise.
//
// Example:
//
//	str := "hello"
//	isPtr := IsPointer(&str)  // true
//	isPtr = IsPointer(str)    // false
//	isPtr = IsPointer(10)     // false
//	isPtr = IsPointer(nil)    // false
func IsPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

// IsWhole checks if a value is a whole number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. It first checks if the value has
// a non-zero fractional part. If it does, it returns false
// since a whole number cannot have a fractional part.
// If the value does not have a fractional part, it returns true.
//
// Examples:
//
//	do.IsWhole(4)   // true
//	do.IsWhole(3.5) // false
//	do.IsWhole(5.0) // true
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func IsWhole[T Numerable](v T) bool {
	_, fraction := math.Modf(float64(v))
	return fraction == 0
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

// MaxList returns the largest value among all input values in a list.
//
// This function requires a list of values of a type that satisfies
// the Numerable interface. It also accepts optional default values,
// which are used when the input list is empty.
//
// If the input list is empty:
//   - If defaults are provided, the maximum value among
//     the defaults is returned.
//   - If no defaults are provided, the function returns the minimal value
//     for the Numerable type (0).
//
// Example usage:
//
//	n := []int{3, 5, 7, 1, 9, 2}
//	m0 := do.MaxList(n, 10, 20)       // 20
//	m1 := do.MaxList([]int{})         //  0
//	m2 := do.MaxList([]int{}, 20, 10) // 20
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func MaxList[T Numerable](v []T, defaults ...T) T {
	if len(v) == 0 {
		if len(defaults) > 0 {
			m := defaults[0]
			return Max(m, defaults[1:]...)
		}

		return 0
	}

	return Max(v[0], v[1:]...)
}

// Median calculates the median value of a variable number
// of values of type Numerable.
//
// It takes a slice of values of type T and returns the median
// value as a float64. The type T must satisfy the Numerable interface.
//
// The median is the middle value of a sorted list of values.
// If the number of values is odd, the median is the middle value.
// If the number of values is even, the median is the average of
// the two middle values.
//
// Example:
//
//	n := []int{3, 5, 7, 1, 9, 2}
//	median := do.Median(n...) // 4.0
//
// This function is generic and can work with any type T that
// satisfies the Numerable interface.
func Median[T Numerable](v ...T) float64 {
	if len(v) == 0 {
		return 0
	}

	// Sort the values
	s := make([]T, len(v))
	copy(s, v)
	Sort(s)

	// Calculate the median
	middle := len(s) / 2
	if len(s)%2 == 0 {
		// Even number of values, average the two middle values
		return float64(s[middle-1]+s[middle]) / 2.0
	} else {
		// Odd number of values, return the middle value
		return float64(s[middle])
	}
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

// MinList returns the smallest value among all input values in a list.
//
// This function requires a list of values of a type that satisfies
// the Numerable interface. It also accepts optional default values,
// which are used when the input list is empty.
//
// If the input list is empty:
//   - If defaults are provided, the minimum value among
//     the defaults is returned.
//   - If no defaults are provided, the function returns
//     the minimum value for the Numerable type (0).
//
// Example usage:
//
//	n := []int{3, 5, 7, 1, 9, 2}
//	m0 := do.MinList(n, 20, 10)       // 1
//	m1 := do.MinList([]int{})         // 0
//	m2 := do.MinList([]int{}, 20, 10) // 10
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func MinList[T Numerable](v []T, defaults ...T) T {
	if len(v) == 0 {
		if len(defaults) > 0 {
			m := defaults[0]
			return Min(m, defaults[1:]...)
		}

		return 0
	}

	return Min(v[0], v[1:]...)
}

// Rank function returns the rank of a numeric value when compared
// to a list of other numeric values.
//
// The function takes a value `v` of type T and a slice `values` of type []T,
// where T satisfies the Numerable interface. It also accepts an optional
// boolean parameter `desc` that specifies the ranking order. If `desc`
// is true, the function returns the rank in descending order, where
// the largest value has rank 1. If `desc` is false or omitted, the function
// returns the rank in ascending order, where the smallest value has rank 1.
//
// The function calculates the rank by comparing the value `v` with the values
// in the `values` slice. The rank represents the position of `v` in the sorted
// list of values. If `v` is not found in the list, the function returns 0.
//
// Example:
//
//	rank := do.Rank(7, []float64{1, 5, 2, 3, 7, 8})       // rank is 5
//	rank := do.Rank(7, []float64{1, 5, 2, 3, 7, 8}, true) // rank is 2
//	rank := do.Rank(9, []float64{1, 5, 2, 3, 7, 8}, true) // rank is 0
func Rank[T Numerable](v T, values []T, desc ...bool) int {
	// Check the order parameter.
	inverse := If(All(desc...), true, false)

	// Sort the values in the specified order.
	Sort(values, inverse)

	// Find the rank of the value.
	for i, val := range values {
		if val == v {
			return i + 1
		}
	}

	return 0
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

// Sort sorts the values in a slice in either ascending or descending order.
//
// This function modifies the original slice in-place and doesn't
// return a new slice. The type T must satisfy the Numerable interface.
//
// If the `inverse` argument is provided and set to true, the values
// are sorted in descending order. Otherwise, the values are sorted
// in ascending order.
//
// The function uses the quicksort algorithm to sort the values.
//
// Example:
//
//	n := []int{3, 5, 1, 9, 2}
//	do.Sort(n)
//	// n is now []int{1, 2, 3, 5, 9}
//
//	n := []int{3, 5, 1, 9, 2}
//	do.Sort(n, true)
//	// n is now []int{9, 5, 3, 2, 1}
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
func Sort[T Numerable](v []T, inverse ...bool) {
	inv := All(inverse...)

	if inv {
		quicksortDesc(v, 0, len(v)-1)
	} else {
		quicksortAsc(v, 0, len(v)-1)
	}
}

// The quicksortAsc is a helper function that performs the quicksort
// algorithm on a slice in ascending order.
func quicksortAsc[T Numerable](v []T, low, high int) {
	if low < high {
		pivotIndex := partitionAsc(v, low, high)
		quicksortAsc(v, low, pivotIndex-1)
		quicksortAsc(v, pivotIndex+1, high)
	}
}

// The quicksortDesc is a helper function that performs the quicksort
// algorithm on a slice in descending order.
func quicksortDesc[T Numerable](v []T, low, high int) {
	if low < high {
		pivotIndex := partitionDesc(v, low, high)
		quicksortDesc(v, low, pivotIndex-1)
		quicksortDesc(v, pivotIndex+1, high)
	}
}

// The partitionAsc is a helper function that selects a pivot and partitions
// the slice around it in ascending order.
func partitionAsc[T Numerable](v []T, low, high int) int {
	pivot := v[high]
	i := low - 1

	for j := low; j < high; j++ {
		if v[j] < pivot {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[high] = v[high], v[i+1]

	return i + 1
}

// The partitionDesc is a helper function that selects a pivot and partitions
// the slice around it in descending order.
func partitionDesc[T Numerable](v []T, low, high int) int {
	pivot := v[high]
	i := low - 1

	for j := low; j < high; j++ {
		if v[j] > pivot {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[high] = v[high], v[i+1]

	return i + 1
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

// HLookup looks up and retrieves data from a specific row in a table.
//
// The function takes a search value `v`, a slice of lookup values `lookup`,
// a slice of result values `result`, and an optional default value `def`.
// It searches for the first occurrence of `v` in the `lookup` slice and
// returns the corresponding value from the `result` slice. If `v` is not
// found in the `lookup` slice, it returns the default value `def`.
//
// Example:
//
//	lookup := []string{"A", "B", "C"}
//	result := []int{1, 2, 3}
//	val := do.HLookup("B", lookup, result, -1)  // val is 2
//	val = do.HLookup("D", lookup, result, -1)   // val is -1
//
// This function is generic and can work with any type T as the search value,
// and any type U as the lookup and result values.
func HLookup[T comparable, U any](v T, lookup []T, result []U, def U) U {
	for i, item := range lookup {
		if item == v {
			return result[i]
		}
	}

	return def
}

// VLookup looks up and retrieves data from a specific column in a table.
//
// The function takes a search value `v`, a slice of lookup values `lookup`,
// a slice of result values `result`, and an optional default value `def`.
// It searches for the first occurrence of `v` in the `lookup` slice and
// returns the corresponding value from the `result` slice. If `v` is not
// found in the `lookup` slice, it returns the default value `def`.
//
// Example:
//
//	lookup := []string{"A", "B", "C"}
//	result := []int{1, 2, 3}
//	val := do.VLookup("B", lookup, result, -1)  // val is 2
//	val = do.VLookup("D", lookup, result, -1)   // val is -1
//
// This function is generic and can work with any type T as the search value,
// and any type U as the lookup and result values.
func VLookup[T comparable, U any](v T, lookup []T, result []U, def U) U {
	for i, item := range lookup {
		if item == v {
			return result[i]
		}
	}

	return def
}
