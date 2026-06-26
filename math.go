package g

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
)

// ErrOverflow is returned by SafeSum when a numeric overflow occurs.
// It can be tested with errors.Is.
var ErrOverflow = errors.New("g: numeric overflow")

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
// Example usage:
//
//	var n int = -10
//	fmt.Println(g.Abs(n))  // Output: 10
//
//	var f float64 = -15.5
//	fmt.Println(g.Abs(f))  // Output: 15.5
//
//	var u uint = 20
//	fmt.Println(g.Abs(u))  // Output: 20
func Abs[T Numerable](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

// Average calculates the average of a variable number
// of values of type Numerable.
//
// The values are accumulated in float64, so the result is correct even
// when the integer sum of the inputs would overflow type T. If no values
// are provided, it returns 0. Note: this function returns the average as
// a float64, regardless of the input type.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	avg := g.Average(values...)
//	fmt.Println(avg)  // Output: 4.5
//
//	floats := []float64{1.1, 2.2, 3.3}
//	avg = g.Average(floats...)
//	fmt.Println(avg)  // Output: 2.2
func Average[T Numerable](v ...T) float64 {
	if len(v) == 0 {
		return 0
	}

	var sum float64
	for _, x := range v {
		sum += float64(x)
	}

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
// the two middle values. The averaging is performed in float64,
// so it does not overflow type T.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	median := g.Median(values...)
//	fmt.Println(median)  // Output: 4.0
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	median = g.Median(floats...)
//	fmt.Println(median)  // Output: 3.3
func Median[T Numerable](v ...T) float64 {
	if len(v) == 0 {
		return 0
	}

	// Sort a copy so the caller's slice is left intact.
	s := make([]T, len(v))
	copy(s, v)
	slices.Sort(s)

	middle := len(s) / 2
	if len(s)%2 == 0 {
		// Even number of values, average the two middle values
		// in float64 to avoid overflowing type T.
		return (float64(s[middle-1]) + float64(s[middle])) / 2.0
	}

	// Odd number of values, return the middle value.
	return float64(s[middle])
}

// Max returns the largest value among all input values.
//
// The type must be Verifiable and support the greater than (>) operator.
// If no values are provided, the zero value of type T is returned; use
// MaxOk to distinguish the empty case explicitly.
//
// For floating-point types NaN propagates: if any value is NaN the result
// is NaN, regardless of its position. This makes the result deterministic.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	maxI := g.Max(values...)
//	fmt.Println(maxI)  // Output: 9
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	maxF := g.Max(floats...)
//	fmt.Println(maxF)  // Output: 5.5
func Max[T Verifiable](v ...T) T {
	if len(v) == 0 {
		var zero T
		return zero
	}
	return slices.Max(v)
}

// MaxOk returns the largest value among all input values and a boolean
// that reports whether any value was provided.
//
// If the input is empty, it returns the zero value of T and false; otherwise
// it returns the maximum and true. This removes the ambiguity of Max, where
// the zero value may also be a legitimate result.
//
// Example usage:
//
//	if m, ok := g.MaxOk(values...); ok {
//	    fmt.Println(m)
//	}
func MaxOk[T Verifiable](v ...T) (T, bool) {
	if len(v) == 0 {
		var zero T
		return zero, false
	}
	return slices.Max(v), true
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
//   - If no defaults are provided, the zero value of type T is returned;
//     use MaxListOk to detect this case explicitly.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	max := g.MaxList(values)
//	fmt.Println(max)  // Output: 9
//
//	empty := []int{}
//	defaults := []int{4, 5, 6}
//	max = g.MaxList(empty, defaults...)
//	fmt.Println(max)  // Output: 6
func MaxList[T Verifiable](v []T, defaults ...T) T {
	m, _ := MaxListOk(v, defaults...)
	return m
}

// MaxListOk returns the largest value in the list (falling back to defaults
// when the list is empty) and a boolean reporting whether any value existed.
//
// It returns false only when both the list and the defaults are empty.
func MaxListOk[T Verifiable](v []T, defaults ...T) (T, bool) {
	if len(v) != 0 {
		return slices.Max(v), true
	}
	if len(defaults) != 0 {
		return slices.Max(defaults), true
	}
	var zero T
	return zero, false
}

// Min returns the smallest value among all input values.
//
// The type must be Verifiable and support the less than (<) operator.
// If no values are provided, the zero value of type T is returned; use
// MinOk to distinguish the empty case explicitly.
//
// For floating-point types NaN propagates: if any value is NaN the result
// is NaN, regardless of its position. This makes the result deterministic.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	minI := g.Min(values...)
//	fmt.Println(minI)  // Output: 1
//
//	strings := []string{"z", "a", "m", "c", "y"}
//	minS := g.Min(strings...)
//	fmt.Println(minS)  // Output: a
func Min[T Verifiable](v ...T) T {
	if len(v) == 0 {
		var zero T
		return zero
	}
	return slices.Min(v)
}

// MinOk returns the smallest value among all input values and a boolean
// that reports whether any value was provided.
//
// If the input is empty, it returns the zero value of T and false; otherwise
// it returns the minimum and true.
//
// Example usage:
//
//	if m, ok := g.MinOk(values...); ok {
//	    fmt.Println(m)
//	}
func MinOk[T Verifiable](v ...T) (T, bool) {
	if len(v) == 0 {
		var zero T
		return zero, false
	}
	return slices.Min(v), true
}

// MinList returns the smallest value among all input values in a list.
//
// This function requires a list of values of a type that satisfies
// the Verifiable interface. It also accepts optional default values,
// which are used when the input list is empty.
//
// If the input list is empty:
//   - If defaults are provided, the minimum value among
//     the defaults is returned.
//   - If no defaults are provided, the zero value of type T is returned;
//     use MinListOk to detect this case explicitly.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	minI := g.MinList(values)
//	fmt.Println(minI)  // Output: 1
//
//	empty := []int{}
//	defaults := []int{4, 5, 6}
//	minD := g.MinList(empty, defaults...)
//	fmt.Println(minD)  // Output: 4
func MinList[T Verifiable](v []T, defaults ...T) T {
	m, _ := MinListOk(v, defaults...)
	return m
}

// MinListOk returns the smallest value in the list (falling back to defaults
// when the list is empty) and a boolean reporting whether any value existed.
//
// It returns false only when both the list and the defaults are empty.
func MinListOk[T Verifiable](v []T, defaults ...T) (T, bool) {
	if len(v) != 0 {
		return slices.Min(v), true
	}
	if len(defaults) != 0 {
		return slices.Min(defaults), true
	}
	var zero T
	return zero, false
}

// Clamp constrains a value to the inclusive range [lo, hi].
//
// If v is less than lo, lo is returned; if v is greater than hi, hi is
// returned; otherwise v is returned unchanged. If the bounds are given in
// the wrong order (lo > hi) they are swapped, so Clamp never panics on
// reversed bounds.
//
// Example usage:
//
//	g.Clamp(15, 0, 10) // Output: 10
//	g.Clamp(-3, 0, 10) // Output: 0
//	g.Clamp(7, 0, 10)  // Output: 7
func Clamp[T Ordered](v, lo, hi T) T {
	if lo > hi {
		lo, hi = hi, lo
	}
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// SafeSum returns the sum of all values with overflow checking.
//
// For integer types it detects overflow (signed and unsigned alike) via
// two's-complement wraparound, which Go defines for all integer types, and
// returns an error wrapping ErrOverflow. For floating-point types it checks
// the result for Inf/NaN. On overflow the zero value and an error are
// returned; the error can be tested with errors.Is(err, g.ErrOverflow).
//
// Example usage:
//
//	// Safe sum with integers.
//	sum, err := g.SafeSum(math.MaxInt, 1)
//	if errors.Is(err, g.ErrOverflow) {
//	    fmt.Println("overflow")
//	} else {
//	    fmt.Println("Sum:", sum)
//	}
//
//	// Safe sum with floats.
//	sum, err = g.SafeSum(1.1, 2.2, 3.3)
//	fmt.Println(sum) // Output: 6.6
func SafeSum[T Numerable](v ...T) (T, error) {
	var s T

	switch any(s).(type) {
	case float32, float64:
		for _, x := range v {
			s += x
		}
		if f := float64(s); math.IsInf(f, 0) || math.IsNaN(f) {
			return 0, fmt.Errorf("g: floating-point overflow: %w", ErrOverflow)
		}
		return s, nil
	default:
		// Integer types. On two's-complement wraparound the running sum
		// crosses past the addend's direction: adding a positive value
		// must not make the result smaller, adding a negative value must
		// not make it larger. For unsigned types the second clause is
		// never reached (values are never negative).
		for _, x := range v {
			n := s + x
			if (x > 0 && n < s) || (x < 0 && n > s) {
				return 0, fmt.Errorf("g: integer overflow: %w", ErrOverflow)
			}
			s = n
		}
		return s, nil
	}
}

// Sum returns the sum of all values.
//
// Sum behaves exactly like the + operator applied in sequence: for integer
// types it wraps around on overflow (two's complement), and for
// floating-point types it follows IEEE-754. It never reports an error.
// When you need overflow detection, use SafeSum.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	sum := g.Sum(values...)
//	fmt.Println(sum)  // Output: 27
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	sum = g.Sum(floats...)
//	fmt.Println(sum)  // Output: 16.5
func Sum[T Numerable](v ...T) T {
	var s T
	for _, x := range v {
		s += x
	}
	return s
}

// IsEven checks if a value is an even number.
//
// The function accepts a value of any type T that satisfies the Numerable
// interface. The check is performed in float64 without converting to int,
// so it stays correct for very large or negative values that do not fit
// into int. For floating-point values it considers only the integer part:
// if the value has a non-zero fractional part it is neither even nor odd,
// so IsEven returns false. If the optional argument f is true, the
// fractional part is ignored and only the integer part's parity is checked.
//
// Example usage:
//
//	g.IsEven(6)        // Output: true
//	g.IsEven(7)        // Output: false
//	g.IsEven(6.6)      // Output: false
//	g.IsEven(6.6, true) // Output: true
func IsEven[T Numerable](v T, f ...bool) bool {
	ip, frac := math.Modf(float64(v))
	if frac != 0 && !All(f...) {
		return false
	}
	return math.Mod(ip, 2) == 0
}

// IsOdd checks if a value is an odd number.
//
// The function accepts a value of any type T that satisfies the Numerable
// interface. Like IsEven, the check is performed in float64 to stay correct
// for values outside the int range. For floating-point values it considers
// only the integer part: if the value has a non-zero fractional part it is
// neither even nor odd, so IsOdd returns false. If the optional argument f
// is true, the fractional part is ignored and only the integer part's
// parity is checked.
//
// Example usage:
//
//	g.IsOdd(7)        // Output: true
//	g.IsOdd(6)        // Output: false
//	g.IsOdd(7.7)      // Output: false
//	g.IsOdd(7.7, true) // Output: true
func IsOdd[T Numerable](v T, f ...bool) bool {
	ip, frac := math.Modf(float64(v))
	if frac != 0 && !All(f...) {
		return false
	}
	return math.Mod(ip, 2) != 0
}

// IsWhole checks if a value is a whole number.
//
// The function accepts a value of any type T that satisfies
// the Numerable interface. It first checks if the value has
// a non-zero fractional part. If it does, it returns false
// since a whole number cannot have a fractional part.
// If the value does not have a fractional part, it returns true.
//
// Example usage:
//
//	g.IsWhole(5)    // Output: true
//	g.IsWhole(5.5)  // Output: false
//	g.IsWhole(0)    // Output: true
//	g.IsWhole(-3)   // Output: true
func IsWhole[T Numerable](v T) bool {
	_, fraction := math.Modf(float64(v))
	return fraction == 0
}

// Random generates a random value of type T based on provided arguments:
//
//   - When called without any arguments, it returns the zero value of T.
//   - When called with one argument n, it returns a random value in [0, n).
//     If n <= 0 the range is empty and the zero value is returned.
//   - When called with two arguments a and b, it returns a random value in
//     [a, b); the bounds may be given in any order. If a == b that value is
//     returned.
//   - When called with more than two arguments, it returns a randomly
//     selected value from the provided arguments.
//
// The generator is the concurrency-safe, automatically seeded source from
// math/rand/v2, so Random never panics on degenerate ranges and needs no
// manual seeding.
//
// Example usage:
//
//	g.Random[int]()        // Output: 0
//	g.Random[int](5)       // Output: a random int in [0, 5)
//	g.Random[int](1, 5)    // Output: a random int in [1, 5)
//	g.Random[int](1, 2, 3) // Output: 1, 2, or 3
func Random[T Numerable](v ...T) T {
	switch len(v) {
	case 0:
		var zero T
		return zero
	case 1:
		return randomValue(0, v[0])
	case 2:
		lo, hi := v[0], v[1]
		if lo == hi {
			return lo
		} else if lo > hi {
			lo, hi = hi, lo
		}
		return randomValue(lo, hi)
	default:
		return v[rand.IntN(len(v))]
	}
}

// The randomValue function generates a random value of type T in the
// half-open range [lo, hi). For an empty or invalid span it returns lo,
// which keeps Random panic-free for degenerate inputs.
func randomValue[T Numerable](lo, hi T) T {
	switch any(lo).(type) {
	case float32, float64:
		return T(float64(lo) + rand.Float64()*float64(hi-lo))
	}

	span := int64(hi - lo)
	if span <= 0 {
		return lo
	}
	return T(rand.Int64N(span) + int64(lo))
}

// RandomList returns a random element from the given list.
// If the list is empty, it returns the zero value of type T.
//
// Example usage:
//
//	list := []int{1, 2, 3, 4, 5}
//	value := g.RandomList(list)
//	fmt.Println(value)  // Output: a random element from the list
//
//	emptyList := []string{}
//	value := g.RandomList(emptyList)
//	fmt.Println(value)  // Output: ""
func RandomList[T any](v []T) T {
	if len(v) == 0 {
		var zero T
		return zero
	}
	return v[rand.IntN(len(v))]
}

// RandomMap returns a random value from the given map.
// If the map is empty, it returns the zero value of type T.
//
// Example usage:
//
//	myMap := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
//	value := g.RandomMap(myMap)
//	fmt.Println(value)  // Output: a random value from the map
//
//	emptyMap := map[string]bool{}
//	value := g.RandomMap(emptyMap)
//	fmt.Println(value)  // Output: false
func RandomMap[K comparable, T any](m map[K]T) T {
	var zero T
	if len(m) == 0 {
		return zero
	}

	n := rand.IntN(len(m))
	i := 0
	for _, val := range m {
		if i == n {
			return val
		}
		i++
	}

	return zero
}

// RandomListPlural returns a slice of n random elements from the given list v.
// If n is less than or equal to zero, it returns an empty slice.
//
// Example usage:
//
//	list := []int{1, 2, 3, 4, 5}
//	values := g.RandomListPlural(3, list)
//	fmt.Println(values)  // Output: a slice of 3 random elements from the list
//
//	emptyList := []string{}
//	values := g.RandomListPlural(2, emptyList)
//	fmt.Println(values)  // Output: []
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
// Example usage:
//
//	myMap := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
//	values := g.RandomMapPlural(2, myMap)
//	fmt.Println(values)  // Output: a slice of 2 random values from the map
//
//	emptyMap := map[string]bool{}
//	values := g.RandomMapPlural(3, emptyMap)
//	fmt.Println(values)  // Output: []
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
