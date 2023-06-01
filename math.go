package do

import (
	"math"
	"reflect"
)

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

// Max returns the largest value among all input values.
//
// The function iterates through all the passed values
// and returns the largest one. The type must be Verifiable
// and support the greater than (>) operator.
//
// Example usage:
//
//	n := []int{3,5,7,1,9,2}
//	m0 := do.Max(n...)
//	m1 := do.Max(3, 5, 7, 1, 9, 2)
//
// This function is generic and can work with any type T.
func Max[T Verifiable](v ...T) T {
	// Return zero if no values are provided.
	if len(v) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}

	max := v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
	}

	return max
}

// MaxList returns the largest value among all input values in a list.
//
// This function requires a list of values of a type that satisfies
// the Verifiable interface. It also accepts optional default values,
// which are used when the input list is empty.
//
// If the input list is empty:
//   - If defaults are provided, the maximum value among
//     the defaults is returned.
//   - If no defaults are provided, the function returns
//     the minimal value for the Verifiable type.
//
// Example usage:
//
//	n := []int{3, 5, 7, 1, 9, 2}
//	m0 := do.MaxList(n, 10, 20)       // 20
//	m1 := do.MaxList([]int{})         //  0
//	m2 := do.MaxList([]int{}, 20, 10) // 20
//
// This function is generic and can work with any type T that satisfies
// the Verifiable interface.
func MaxList[T Verifiable](v []T, defaults ...T) T {
	return If(len(v) != 0, Max(v...), Max(defaults...))
}

// Min returns the smallest value among all input values.
//
// The function iterates through all the passed values
// and returns the smallest one. The type must be Verifiable
// and support the less than (<) operator.
//
// Example usage:
//
//	n := []int{3,5,7,1,9,2}
//	m0 := do.Min(n...)
//	m1 := do.Min(3, 5, 7, 1, 9, 2)
//
// This function is generic and can work with any type T.
func Min[T Verifiable](v ...T) T {
	if len(v) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}

	min := v[0]
	for _, val := range v {
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
//     the minimum value for the Verifiable type.
//
// Example usage:
//
//	n := []int{3, 5, 7, 1, 9, 2}
//	m0 := do.MinList(n, 20, 10)       // 1
//	m1 := do.MinList([]int{})         // 0
//	m2 := do.MinList([]int{}, 20, 10) // 10
//
// This function is generic and can work with any type T that satisfies
// the Verifiable interface.
func MinList[T Verifiable](v []T, defaults ...T) T {
	return If(len(v) != 0, Min(v...), Min(defaults...))
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
//	do.IsEven(4.2, true) // true
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

// Random generates a random value of type T based on provided arguments:
//
//   - When called without any arguments, it returns 0.
//   - When called with one argument, it returns a random value from 0 to n-1.
//   - When called with two arguments, it returns a random value from a to b-1.
//   - When called with more than two arguments, it returns a randomly selected
//     value from the provided arguments.
//
// The function uses the time in nanoseconds as a seed for the random
// number generator.
//
// Example:
//
//	Random[int]() => returns 0
//	Random[int](5) => returns a random int from 0 to 4
//	Random[int](1, 5) => returns a random int from 1 to 4
//	Random[int](1, 2, 3) => returns 1, 2, or 3
func Random[T Numerable](v ...T) T {
	switch len(v) {
	case 0:
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	case 1:
		return randomValue(0, v[0])
	case 2:
		min := v[0]
		max := v[1]
		if min == max {
			return min
		} else if min > max {
			min, max = max, min
		}
		return randomValue(min, max)
	default:
		return v[randomGenerator.Intn(len(v))]
	}
}

// The randomValue function generates a random value of type T
// based on provided arguments:
func randomValue[T Numerable](min, max T) T {
	var t interface{} = min
	_, ok32 := t.(float32)
	_, ok64 := t.(float64)
	if ok32 || ok64 {
		return T(float64(min) + randomGenerator.Float64()*float64(max-min))
	}

	return T(randomGenerator.Intn(int(max-min)) + int(min))
}

// RandomList returns a random element from the given list.
// If the list is empty, it returns the zero value of type T.
//
// Example:
//
//	list := []int{1, 2, 3, 4, 5}
//	value := RandomList(list) // returns a random element from the list
//
//	emptyList := []string{}
//	value := RandomList(emptyList) // returns the zero value of string type
func RandomList[T any](v []T) T {
	if len(v) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}
	return v[randomGenerator.Intn(len(v))]
}

// RandomMap returns a random value from the given map.
// If the map is empty, it returns the zero value of type T.
//
// Example:
//
//	myMap := map[string]int{
//	    "apple":  1,
//	    "banana": 2,
//	    "cherry": 3,
//	}
//	value := RandomMap(myMap) // returns a random value from the map
//
//	emptyMap := map[string]bool{}
//	value := RandomMap(emptyMap) // returns the zero value of bool type (false)
func RandomMap[K comparable, T any](m map[K]T) T {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}

	if len(keys) != 0 {
		if v, ok := m[keys[randomGenerator.Intn(len(keys))]]; ok {
			return v
		}
	}

	return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
}

// RandomListPlural returns a slice of n random elements from the given list v.
// If n is less than or equal to zero, it returns an empty slice.
//
// Example:
//
//	list := []int{1, 2, 3, 4, 5}
//	values := RandomListPlural(3, list) // returns a slice of 3 random
//	                                    // elements from the list
//	emptyList := []string{}
//	values := RandomListPlural(2, emptyList) // returns an empty slice
//
//	values := RandomListPlural(0, list) // returns an empty slice
func RandomListPlural[T any](n int, v []T) []T {
	if n <= 0 || len(v) == 0 {
		return make([]T, 0)
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = RandomList(v)
	}

	return result
}

// RandomMapPlural returns a slice of n random values from the given map m.
// If n is less than or equal to zero, it returns an empty slice.
//
// Example:
//
//	myMap := map[string]int{
//	    "apple":  1,
//	    "banana": 2,
//	    "cherry": 3,
//	}
//	values := RandomMapPlural(2, myMap) // returns a slice of 2 random
//		                                // values from the map
//	emptyMap := map[string]bool{}
//	values := RandomMapPlural(3, emptyMap) // returns an empty slice ([]bool{})
//
//	values := RandomMapPlural(0, myMap) // returns an empty slice ([]int{})
func RandomMapPlural[K comparable, T any](n int, m map[K]T) []T {
	if n <= 0 || len(m) == 0 {
		return make([]T, 0)
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = RandomMap(m)
	}

	return result
}
