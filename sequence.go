package g

import (
	"context"
	"math"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// The logicFoundValue is a helper struct that holds a boolean value
// and a Mutex to protect it from concurrent access.
//
// They are used in the In function to detect the desired result
// in a separate goroutine.
type logicFoundValue struct {
	m     sync.Mutex
	value bool
}

// SetValue sets a new value for the Found. It locks the Mutex before
// changing the value and unlocks it after the change is complete.
func (f *logicFoundValue) SetValue(value bool) {
	f.m.Lock()
	defer f.m.Unlock()
	f.value = value
}

// GetValue retrieves the current value of the Found. It locks the Mutex
// before reading the value and unlocks it after the read is complete.
func (f *logicFoundValue) GetValue() bool {
	f.m.Lock()
	defer f.m.Unlock()
	return f.value
}

// Contains checks if a slice contains a specific element.
//
// It takes a slice and an element of the same type as input, and returns
// a boolean value indicating whether the element is found in the slice.
//
// T is the type of items in the slice and the element to be searched.
// The function returns true if the element is found, and false otherwise.
//
// Example usage:
//
//	numSlice := []int{1, 2, 3, 4, 5}
//	resultNum := g.Contains(3, numSlice)
//	fmt.Println(resultNum) // Output: true
//
//	strSlice := []string{"Hello", "World", "Golang"}
//	resultStr := g.Contains("Python", strSlice)
//	fmt.Println(resultStr) // Output: false
func Contains[T Verifiable](v T, vs []T) bool {
	return In(v, vs...)
}

// Filter applies a predicate function to all items in an input
// slice and returns a new slice with the items for which the
// predicate function returns true.
//
// T is the type of items in the input slice.
// The predicate function f takes an item of type T and returns a boolean.
//
// Example usage:
//
//	// Let's say you have a slice of integers and you want
//	// to filter out even numbers:
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	evens := g.Filter(nums, func(v int) bool { return v%2 == 0 })
//	fmt.Println(evens) // Output: [2 4 6 8 10]
//
//	// Or you may have a slice of strings and you want
//	// to filter out strings with length greater than 5:
//	strs := []string{"apple", "banana", "cherry", "date", "elderberry"}
//	longStrings := g.Filter(strs, func(v string) bool { return len(v) > 5 })
//	fmt.Println(longStrings) // Output: ["banana" "cherry" "elderberry"]
func Filter[T any](vs []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

// Index returns the index of the first occurrence of a specific
// element in a slice, or -1 if the element is not present.
//
// T is the type of the items in the slice and the element to be searched.
// The function returns an integer indicating the position of the first
// occurrence of the element, or -1 if the element is not found.
//
// Example usage:
//
//	// Let's say you have a slice of integers and you want
//	// to find the index of the number 7:
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	idx := g.Index(nums, 7)
//	fmt.Println(idx) // Output: 6
//
//	// Or you have a slice of strings and you want to find
//	// the index of the string "cherry":
//	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
//	idx := g.Index(fruits, "cherry")
//	fmt.Println(idx) // Output: 2
//
//	// In case the element is not in the slice, the function will return -1:
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	idx := g.Index(nums, 11)
//	fmt.Println(idx) // Output: -1
func Index[T comparable](vs []T, v T) int {
	for i, item := range vs {
		if item == v {
			return i
		}
	}

	return -1
}

// Map applies a function to all items in an input slice and returns
// a new slice with the transformed items.
//
// T is the type of items in the input slice, and U is the type of
// items in the output slice.
// The function f takes an item of type T and returns a new item of type U.
//
// Example usage:
//
//	// Let's say you have a slice of integers and you want to create a
//	// new slice where each element is the square of the original element:
//	nums := []int{1, 2, 3, 4, 5}
//	squares := g.Map(nums, func(n int) int {
//	   return n * n
//	})
//	fmt.Println(squares) // Output: [1 4 9 16 25]
//
//	// Or you have a slice of strings and you want to create a new slice
//	// where each element is the length of the original string:
//	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
//	lengths := g.Map(fruits, func(s string) int {
//	   return len(s)
//	})
//	fmt.Println(lengths) // Output: [5 6 6 4 10]
func Map[T any, U any](vs []T, f func(T) U) []U {
	result := make([]U, len(vs))
	for i, v := range vs {
		result[i] = f(v)
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
//
// Example usage:
//
//	// Let's say you have a slice of integers and you want to
//	// compute the sum of all elements:
//	nums := []int{1, 2, 3, 4, 5}
//	sum := go.Reduce(nums, func(acc int, n int) int {
//	   return acc + n
//	}, 0)
//	fmt.Println(sum) // Output: 15
//
//	// Or you have a slice of strings and you want to concatenate
//	// them all into a single string:
//	words := []string{"Hello", "World", "From", "Go"}
//	sentence := go.Reduce(words, func(acc string, s string) string {
//	   return acc + " " + s
//	}, "")
//	fmt.Println(sentence) // Output: "Hello World From Go"
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
// Example usage:
//
//	// If you have a slice of integers that you want
//	// to sort in ascending order:
//	nums := []int{5, 2, 7, 8, 1, 9}
//	g.Sort(nums)
//	fmt.Println(nums) // Output: [1 2 5 7 8 9]
//
//	// If you want to sort the same slice in descending order:
//	nums := []int{5, 2, 7, 8, 1, 9}
//	g.Sort(nums, true)
//	fmt.Println(nums) // Output: [9 8 7 5 2 1]
//
//	// This function is generic and can work with any type
//	// that satisfies the Numerable interface.
//	// For example, if you have a slice of floats:
//	floats := []float64{5.5, 2.2, 7.7, 8.8, 1.1, 9.9}
//	g.Sort(floats)
//	fmt.Println(floats) // Output: [1.1 2.2 5.5 7.7 8.8 9.9]
func Sort[T Verifiable](v []T, inverse ...bool) {
	inv := All(inverse...)

	if inv {
		quicksortDesc(v, 0, len(v)-1)
	} else {
		quicksortAsc(v, 0, len(v)-1)
	}
}

// The quicksortAsc is a helper function that performs the quicksort
// algorithm on a slice in ascending order.
func quicksortAsc[T Verifiable](v []T, low, high int) {
	if low < high {
		pivotIndex := partitionAsc(v, low, high)
		quicksortAsc(v, low, pivotIndex-1)
		quicksortAsc(v, pivotIndex+1, high)
	}
}

// The quicksortDesc is a helper function that performs the quicksort
// algorithm on a slice in descending order.
func quicksortDesc[T Verifiable](v []T, low, high int) {
	if low < high {
		pivotIndex := partitionDesc(v, low, high)
		quicksortDesc(v, low, pivotIndex-1)
		quicksortDesc(v, pivotIndex+1, high)
	}
}

// The partitionAsc is a helper function that selects a pivot and partitions
// the slice around it in ascending order.
func partitionAsc[T Verifiable](v []T, low, high int) int {
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
func partitionDesc[T Verifiable](v []T, low, high int) int {
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

// Value returns the first non-zero value from the parameters.
//
// The function checks each parameter in the order they are passed to
// the function and returns the first parameter that is not a zero value.
// If all parameters are zero values, the function returns a zero value
// for the type T of the parameters.
//
// This function is generic and can work with any type T.
//
// Example usage:
//
//	// If you want to find the first non-zero value among
//	// several integer variables:
//	a, b, c := 0, 0, 3
//	fmt.Println(g.Value(a, b, c)) // Output: 3
//
//	// If all values are zero, the function will return zero of the type T:
//	a, b, c := 0, 0, 0
//	fmt.Println(g.Value(a, b, c)) // Output: 0
//
//	// This function can work with any type. For example,
//	// if you have several string variables:
//	s1, s2, s3 := "", "Hello", "World"
//	fmt.Println(g.Value(s1, s2, s3)) // Output: Hello
//
//	// If all strings are empty (which is the zero value for strings),
//	// the function will return an empty string:
//	s1, s2, s3 := "", "", ""
//	fmt.Println(g.Value(s1, s2, s3)) // Output: ""
//
// Warning: the function checks the list as the whole object, that is:
//
//	l := []int{0, 1, 2, 3}
//	g.Value(l)    // Returns: []int{0, 1, 2, 3}, list is not empty object
//	g.Value(l...) // Returns: 1, because 1 is the first not empty value
func Value[T any](v ...T) T {
	if len(v) != 0 {
		for _, val := range v {
			if !IsEmpty(val) {
				return val
			}
		}
	}

	return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
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
//
// Example usage:
//
//	// If you have two slices of the same length:
//	a := []int{1, 2, 3}
//	b := []string{"one", "two", "three"}
//	pairs := g.Zip(a, b)
//	for _, pair := range pairs {
//	    fmt.Printf("(%d, %s)\n", pair.First, pair.Second)
//	}
//	// Output: (1, one) (2, two) (3, three)
//
//	// If one slice is shorter than the other:
//	a := []int{1, 2}
//	b := []string{"one", "two", "three"}
//	pairs := g.Zip(a, b)
//	for _, pair := range pairs {
//	    fmt.Printf("(%d, %s)\n", pair.First, pair.Second)
//	}
//	// Output: (1, one) (2, two)
//	// Note that the third element of the second slice is ignored.
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

// Distinct returns a new slice with unique values from the input slice.
//
// It removes duplicates from the input slice and returns a new slice
// with unique values in the same order as they appear in the input slice.
//
// This function is generic and can work with any type T.
//
// Example usage:
//
//	// If you have a slice with duplicates:
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
//	uniqueNumbers := g.Distinct(numbers)
//	fmt.Println(uniqueNumbers)  // Output: [1 2 3 4]
//
//	// This function also works with slices of other comparable types,
//	// like strings:
//	words := []string{"hello", "world", "hello", "gophers"}
//	uniqueWords := g.Distinct(words)
//	fmt.Println(uniqueWords)  // Output: ["hello" "world" "gophers"]
func Distinct[T comparable](v []T) []T {
	uniqueMap := make(map[T]bool)
	uniqueValues := make([]T, 0)

	for _, item := range v {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueValues = append(uniqueValues, item)
		}
	}

	return uniqueValues
}

// Shuffle randomly shuffles the elements in the input slice.
//
// It modifies the input slice in-place by rearranging the elements
// in a random order using the Fisher-Yates algorithm.
//
// This function is generic and can work with any type T.
//
// Example usage:
//
//	// If you have a slice and want to shuffle its elements:
//	numbers := []int{1, 2, 3, 4, 5}
//	g.Shuffle(numbers)
//	fmt.Println(numbers)
//	// Prints the slice numbers in a random order, e.g., [3 1 5 2 4]
//
//	// This function also works with slices of other types, like strings:
//	words := []string{"hello", "world", "gophers", "Go", "OpenAI"}
//	g.Shuffle(words)
//	fmt.Println(words)
//	// Prints the slice words in a random order, e.g.,
//	// ["Go" "gophers" "hello" "OpenAI" "world"]
func Shuffle[T any](v []T) {
	rand.Seed(time.Now().UnixNano())

	for i := len(v) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		v[i], v[j] = v[j], v[i]
	}
}

// Product calculates the product of all numeric values in the input slice.
//
// It iterates through the input slice and multiplies all the numeric values
// together to compute the product. The type T must satisfy the Numerable
// interface.
//
// If there are no numeric values in the slice, the function returns 1.
//
// Example usage:
//
//	// Calculate the product of a slice of integers:
//	nums := []int{2, 3, 4}
//	p := g.Product(nums)
//	fmt.Println(p)  // Outputs: 24
//
//	// Compute the product of a slice of floats:
//	nums := []float64{1.2, 3.4, 5.6}
//	p := g.Product(nums)
//	fmt.Println(p)  // Outputs: 22.848
//
//	// If the slice is empty, the function returns 1:
//	nums := []int{}
//	p := g.Product(nums)
//	fmt.Println(p)  // Outputs: 1
func Product[T Numerable](v ...T) T {
	if len(v) == 0 {
		return 1
	}

	prod := v[0]
	for _, val := range v[1:] {
		prod *= val
	}

	return prod
}

// Merge merges two sorted arrays into a single sorted array.
//
// It takes two sorted arrays a and b as input and returns a single
// sorted array that contains all the elements from both arrays.
//
// The function assumes that both input arrays are already sorted
// in ascending order.
//
// Example:
//
//	a := []int{1, 3, 5}
//	b := []int{2, 4, 6}
//	mergedUnsort := g.Merge(a, b)     // [1 3 5 2 4 6]
//	mergedSort := g.Merge(a, b, true) // [1 2 3 4 5 6]
//
// This function is generic and can work with any type T.
func Merge[T Verifiable](a []T, b []T, sort ...bool) []T {
	merged := make([]T, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	if All(sort...) {
		Sort(merged)
	}

	return merged
}

// In is a generic function that checks if a given value 'v' of type 'T' exists
// in a variadic parameter list of 'T'. The function uses the 'Verifiable' type
// constraint which allows it to operate on numeric types and strings.
//
// This function leverages goroutines for concurrent computation when the size
// of the list is large. It splits the list into chunks (based on the number
// of CPU cores) and checks each chunk in a separate goroutine. This allows the
// function to take advantage of multi-core processors and improves performance
// on large data sets.
//
// A 'sync.WaitGroup' is used to ensure all goroutines have completed, and a
// thread-safe structure 'Found' is used to safely access the shared 'found'
// variable across goroutines.
//
// Usage:
//
//	// Define a slice of integers.
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	// Check if '5' exists in the slice.
//	exists := g.In(5, numbers...)
//	fmt.Println(exists)  // Output: true
//
//	// Define a slice of strings.
//	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
//
//	// Check if 'date' exists in the slice
//	exists = g.In("date", words...)
//	fmt.Println(exists)  // Output: true
func In[T Verifiable](v T, list ...T) bool {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{}

	// If the length of the slice is less than or equal to
	// the minLoadPerGoroutine, then we do not need
	// to use goroutines.
	if l := len(list); l == 0 {
		return false
	} else if l/p < minLoadPerGoroutine {
		for _, b := range list {
			if b == v {
				return true
			}
		}

		return false
	}

	chunkSize := len(list) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(list)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range list[start:end] {
				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if b == v {
					found.SetValue(true)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// Range generates a slice of integers based on the provided parameters.
// Returns nil if the parameters are invalid.
//
//   - If a single parameter is passed (Range(n)), the function returns
//     a slice from 0 to n-1.
//   - If two parameters are passed (Range(n, m)), the function returns
//     a slice from n to m-1.
//   - If three parameters are passed (Range(n, m, s)), the function returns
//     a slice from n to m-1 with a step size of s.
//
// The function returns nil if:
//   - The step size is zero
//   - The range is decreasing but step is positive
//   - The range is increasing but step is negative
//   - The resulting sequence would exceed MaxRangeSize
//
// The maximum size of the generated slice is set by the MaxRangeSize constant.
//
// Example usage:
//
//	result := g.Range(5)
//	// Output: [0, 1, 2, 3, 4]
//
//	result := g.Range(3, 7)
//	// Output: [3, 4, 5, 6]
//
//	result := g.Range(1, 10, 2)
//	// Output: [1, 3, 5, 7, 9]
//
//	result := g.Range(10, 0, -1)
//	// Output: [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
//
//	result := g.Range(1, 10, 0)
//	// Output: nil - step size cannot be zero
//
//	result := g.Range(10, 1, 1)
//	// Output: nil - decreasing range with positive step
func Range(a int, opt ...int) []int {
	var n, m, s int = 0, a, 1

	// Sets range as n to m-1.
	if len(opt) > 0 {
		n = a
		m = opt[0]
	}

	// Sets step size.
	if len(opt) > 1 {
		s = opt[1]
	}

	// Ignore incorrect parameters.
	if s == 0 || s < 0 && n <= m || s > 0 && n >= m {
		return nil
	}

	// Calculate the number of steps and create the slice of this size.
	steps := Abs(int(math.Ceil(float64(m-n) / float64(s))))

	// Limit the size of the slice.
	if steps >= MaxRangeSize {
		return nil
	}

	result := make([]int, steps)
	for i := 0; i < steps; i++ {
		result[i] = n + i*s
	}

	return result
}

// Rangef generates a slice of values based on the provided parameters
// and a given function as func(int) T.
// Returns nil if the parameters are invalid.
//
//   - If a single parameter is passed (Range(fn, n)), the function returns
//     a slice from 0 to n-1 and pass it in fn.
//   - If two parameters are passed (Range(fn, n, m)), the function returns
//     a slice from n to m-1 and pass it in fn.
//   - If three parameters are passed (Range(fn,  n, m, s)), the function
//     returns a slice from n to m-1 with a step size of s, and pass it in fn.
//
// The function returns nil if:
//   - The step size is zero
//   - The range is decreasing but step is positive
//   - The range is increasing but step is negative
//   - The resulting sequence would exceed MaxRangeSize
//
// Example usage:
//
//	func appleFactory() func(int) string {
//	    var appleVarieties = []string{
//	        "Gala",
//	        "Fuji",
//	        "Honeycrisp",
//	        "Red Delicious",
//	        "Granny Smith",
//	        "Golden Delicious",
//	        "Pink Lady",
//	        "Braeburn",
//	        "McIntosh",
//	        "Jazz",
//	    }
//
//	    return func(i int) string {
//	        if i >= 0 && i < len(appleVarieties) {
//	            return appleVarieties[i]
//	        }
//	        return "-"
//	    }
//	}
//
//	result := g.Rangef(appleFactory(), 3)
//	// Output: [Gala Fuji Honeycrisp]
//
//	result := g.Rangef(appleFactory(), 4, 7)
//	// Output: [Granny Smith Golden Delicious Pink Lady]
//
//	result := g.Rangef(appleFactory(), 7, 12, 2)
//	// Output: [Braeburn Jazz -]
//
//	result := g.Rangef(appleFactory(), 1, 10, 0)
//	// Output: nil - step size cannot be zero
func Rangef[T any](fn func(int) T, a int, opt ...int) []T {
	var n, m, s int = 0, a, 1

	// Sets range as n to m-1.
	if len(opt) > 0 {
		n = a
		m = opt[0]
	}

	// Sets step size.
	if len(opt) > 1 {
		s = opt[1]
	}

	// Ignore incorrect parameters.
	if s == 0 || s < 0 && n <= m || s > 0 && n >= m {
		return nil
	}

	// Calculate the number of steps and create the slice of this size.
	steps := Abs(int(math.Ceil(float64(m-n) / float64(s))))
	steps = If(steps >= MaxRangeSize, 0, steps) // limit the size of the slice
	result := make([]T, steps)

	for i := 0; i < steps; i++ {
		result[i] = fn(n + i*s)
	}

	return result
}

// Reverse changes slice with elements in reverse order.
//
// It takes a slice 'v' of any type 'T' and swaps the elements
// in-place to reverse their order.
//
// Example usage:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	Reverse(numbers)
//	// numbers: [5, 4, 3, 2, 1]
//
//	names := []string{"Alice", "Bob", "Charlie", "Dave"}
//	Reverse(names)
//	// names: ["Dave", "Charlie", "Bob", "Alice"]
//
//	empty := []bool{}
//	Reverse(empty)
//	// empty: []
func Reverse[T any](v []T) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}
