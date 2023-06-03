package do

import (
	"context"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// Create a new type Found that can be safely shared across multiple goroutines.
type Found struct {
	m     sync.Mutex
	value bool
}

// SetValue sets a new value for the Found. It locks the Mutex before
// changing the value and unlocks it after the change is complete.
func (f *Found) SetValue(value bool) {
	f.m.Lock()
	defer f.m.Unlock()
	f.value = value
}

// GetValue retrieves the current value of the Found. It locks the Mutex
// before reading the value and unlocks it after the read is complete.
func (f *Found) GetValue() bool {
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
// Example:
//
//	n := []int{2, 3, 4}
//	prod := do.Product(n...) // 24
//
// This function is generic and can work with any type T that satisfies
// the Numerable interface.
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
//	mergedUnsort := do.Merge(a, b)     // [1 3 5 2 4 6]
//	mergedSort := do.Merge(a, b, true) // [1 2 3 4 5 6]
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
//	// Define a slice of integers
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	// Check if '5' exists in the slice
//	exists := In(5, numbers...)
//	fmt.Println(exists)  // Output: true
//
//	// Define a slice of strings
//	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
//
//	// Check if 'date' exists in the slice
//	exists = In("date", words...)
//	fmt.Println(exists)  // Output: true
func In[T Verifiable](v T, list ...T) bool {
	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := runtime.NumCPU() * 2
	found := &Found{}

	if len(list) < p {
		return in(v, list...)
	}

	var wg sync.WaitGroup

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
					cancel()
					return
				}
			}
		}(start, end)
	}

	wg.Wait()

	return found.GetValue()
}

// The in performs a sequential entry search.
func in[T Verifiable](v T, list ...T) bool {
	for _, b := range list {
		if b == v {
			return true
		}
	}
	return false
}
