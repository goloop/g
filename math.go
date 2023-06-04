package do

import (
	"math"
	"reflect"
	"sync"
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
// Example usage:
//
//	var n int = -10
//	fmt.Println(do.Abs(n))  // Output: 10
//
//	var f float64 = -15.5
//	fmt.Println(do.Abs(f))  // Output: 15.5
//
//	var u uint = 20
//	fmt.Println(do.Abs(u))  // Output: 20
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
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	avg := do.Average(values...)
//	fmt.Println(avg)  // Output: 4.5
//
//	floats := []float64{1.1, 2.2, 3.3}
//	avg = do.Average(floats...)
//	fmt.Println(avg)  // Output: 2.2
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
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	median := do.Median(values...)
//	fmt.Println(median)  // Output: 4.0
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	median = do.Median(floats...)
//	fmt.Println(median)  // Output: 3.3
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

// The doMiniMax function is used by the Min and Max functions
// to calculate the minimum and maximum values.
func doMiniMax[T Verifiable](m bool, v ...T) T {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	// Return zero if no values are provided.
	if len(v) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}

	p := parallelTasks
	chunkSize := len(v) / p
	resultChan := make(chan T, p)
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(v)
		}

		go func(c []T) {
			defer wg.Done()
			if len(c) == 0 {
				return
			}

			r := c[0]
			for _, val := range c[1:] {
				if m {
					if val > r {
						r = val
					}
				} else {
					if val < r {
						r = val
					}
				}
			}

			mu.Lock()
			resultChan <- r
			mu.Unlock()
		}(v[start:end])
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	r := v[0]
	for val := range resultChan {
		if m {
			if val > r {
				r = val
			}
		} else {
			if val < r {
				r = val
			}
		}
	}

	return r
}

// Max returns the largest value among all input values.
//
// The function iterates through all the passed values
// and returns the largest one. The type must be Verifiable
// and support the greater than (>) operator.
//
// Example usage:
//
//	values := []int{3, 5, 7, 1, 9, 2}
//	maxI := do.Max(values...)
//	fmt.Println(maxI)  // Output: 9
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	maxF: = do.Max(floats...)
//	fmt.Println(maxF)  // Output: 5.5
func Max[T Verifiable](v ...T) T {
	return doMiniMax(true, v...)
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
//	values := []int{3, 5, 7, 1, 9, 2}
//	max := do.MaxList(values)
//	fmt.Println(max)  // Output: 9
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	max = do.MaxList(floats)
//	fmt.Println(max)  // Output: 5.5
//
//	empty := []int{}
//	defaults := []int{4, 5, 6}
//	max = do.MaxList(empty, defaults...)
//	fmt.Println(max)  // Output: 6
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
//	values := []int{3, 5, 7, 1, 9, 2}
//	minI := do.Min(values...)
//	fmt.Println(minI)  // Output: 1
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	minF = do.Min(floats...)
//	fmt.Println(minF)  // Output: 1.1
//
//	strings := []string{"z", "a", "m", "c", "y"}
//	minS = do.Min(strings...)
//	fmt.Println(minS)  // Output: a
func Min[T Verifiable](v ...T) T {
	return doMiniMax(false, v...)
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
//	values := []int{3, 5, 7, 1, 9, 2}
//	minI := do.MinList(values)
//	fmt.Println(minI)  // Output: 1
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	minF = do.MinList(floats)
//	fmt.Println(minF)  // Output: 1.1
//
//	empty := []int{}
//	defaults := []int{4, 5, 6}
//	minD := do.MinList(empty, defaults...)
//	fmt.Println(minD)  // Output: 4
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
//	values := []int{3, 5, 7, 1, 9, 2}
//	sum := do.Sum(values...)
//	fmt.Println(sum)  // Output: 27
//
//	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//	sum = do.Sum(floats...)
//	fmt.Println(sum)  // Output: 16.5
func Sum[T Numerable](v ...T) T {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	if len(v) == 0 {
		return T(0)
	}

	p := parallelTasks
	chunkSize := len(v) / p
	resultChan := make(chan T, p)
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(v)
		}

		go func(c []T) {
			defer wg.Done()

			r := T(0)
			for _, val := range c {
				r += val
			}

			mu.Lock()
			resultChan <- r
			mu.Unlock()
		}(v[start:end])
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	r := T(0)
	for val := range resultChan {
		r += val
	}

	return r
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
// Example usage:
//
//	even := do.IsEven(6)
//	fmt.Println(even)  // Output: true
//
//	odd := do.IsEven(7)
//	fmt.Println(odd)  // Output: false
//
//	floatingPoint := do.IsEven(6.6)
//	fmt.Println(floatingPoint)  // Output: false
//
//	floatingPoint = do.IsEven(6.6, true)
//	fmt.Println(floatingPoint)  // Output: true
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
// Example usage:
//
//	odd := do.IsOdd(7)
//	fmt.Println(odd)  // Output: true
//
//	even := do.IsOdd(6)
//	fmt.Println(even)  // Output: false
//
//	floatingPoint := do.IsOdd(7.7)
//	fmt.Println(floatingPoint)  // Output: false
//
//	floatingPoint = do.IsOdd(7.7, true)
//	fmt.Println(floatingPoint)  // Output: true
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
// Example usage:
//
//	whole := do.IsWhole(5)
//	fmt.Println(whole)  // Output: true
//
//	notWhole := do.IsWhole(5.5)
//	fmt.Println(notWhole)  // Output: false
//
//	zero := do.IsWhole(0)
//	fmt.Println(zero)  // Output: true
//
//	negative := do.IsWhole(-3)
//	fmt.Println(negative)  // Output: true
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
// Example usage:
//
//	rand0 := do.Random[int]()
//	fmt.Println(rand0)  // Output: 0
//
//	rand1 := do.Random[int](5)
//	fmt.Println(rand1)  // Output: a random int from 0 to 4
//
//	rand2 := do.Random[int](1, 5)
//	fmt.Println(rand2)  // Output: a random int from 1 to 4
//
//	rand3 := do.Random[int](1, 2, 3)
//	fmt.Println(rand3)  // Output: 1, 2, or 3
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
// Example usage:
//
//	list := []int{1, 2, 3, 4, 5}
//	value := do.RandomList(list)
//	fmt.Println(value)  // Output: a random element from the list
//
//	emptyList := []string{}
//	value := do.RandomList(emptyList)
//	fmt.Println(value)  // Output: ""
func RandomList[T any](v []T) T {
	if len(v) == 0 {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}
	return v[randomGenerator.Intn(len(v))]
}

// RandomMap returns a random value from the given map.
// If the map is empty, it returns the zero value of type T.
//
// Example usage:
//
//	myMap := map[string]int{
//	    "apple":  1,
//	    "banana": 2,
//	    "cherry": 3,
//	}
//	value := do.RandomMap(myMap)
//	fmt.Println(value)  // Output: a random value from the map
//
//	emptyMap := map[string]bool{}
//	value := do.RandomMap(emptyMap)
//	fmt.Println(value)  // Output: zero value for T type (false)
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
// Example usage:
//
//	list := []int{1, 2, 3, 4, 5}
//	values := do.RandomListPlural(3, list)
//	fmt.Println(values)  // Output: a slice of 3 random elements from the list
//
//	emptyList := []string{}
//	values := do.RandomListPlural(2, emptyList)
//	fmt.Println(values)  // Output: []
//
//	values := do.RandomListPlural(0, list)
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
//	myMap := map[string]int{
//	    "apple":  1,
//	    "banana": 2,
//	    "cherry": 3,
//	}
//	values := do.RandomMapPlural(2, myMap)
//	fmt.Println(values)  // Output: a slice of 2 random values from the map
//
//	emptyMap := map[string]bool{}
//	values := do.RandomMapPlural(3, emptyMap)
//	fmt.Println(values)  // Output: []
//
//	values := do.RandomMapPlural(0, myMap)
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
