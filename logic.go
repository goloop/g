package g

import (
	"context"
	"reflect"
	"sync"
)

// If is a substitute for the ternary operator (?:)
// which is not available in Go.
//
// In languages like C/C++ and Python, you can use a ternary
// operator for a concise conditional expression:
//
//	C/C++:  int max = (a > b) ? a : b;
//	Python: max = a if a > b else b
//
// It takes three parameters: a boolean expression e, and two
// values of any type T (t and f). If the expression e is true,
// it returns t, otherwise it returns f.
//
// Example:
//
//	// Condition is true.
//	max := g.If(3 > 2, 3, 2)  // Output: 3
//
//	// If condition is false.
//	max := g.If(2 > 3, 2, 3)  // Output: 3
//
//	// Using with strings.
//	greeting := g.If(user == "admin", "Hello, admin", "Hello, user")
//	fmt.Println(greeting) // Output: the appropriate greeting
func If[T any](e bool, t, f T) T {
	if e {
		return t
	}

	return f
}

// All returns true if all values in the provided slice
// are not zero values for their types.
//
// If at least one value is a zero value, it immediately returns false.
// If the slice is empty, it returns false.
//
// This function is generic and can work with any type T.
//
// Example usage:
//
//	allNonZero := g.All(1, 2, 3)
//	fmt.Println(allNonZero) // Output: true
//
//	someZero := g.All(1, 0, 3)
//	fmt.Println(someZero) // Output: false
//
//	allNonZeroMixed := g.All(1, "a", true)
//	fmt.Println(allNonZeroMixed) // Output: true
//
//	empty := g.All()
//	fmt.Println(empty) // Output: false
func All[T any](v ...T) bool {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: true}

	if len(v) == 0 {
		return false
	}

	chunkSize := len(v) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(v)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range v[start:end] {
				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if IsEmpty(b) {
					found.SetValue(false)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// Any returns true if at least one value in the provided slice
// is not a zero value for its type.
//
// As soon as it finds a value that is not a zero value, it returns true.
// If all values in the slice are zero values or the slice is empty,
// it returns false.
//
// Example usage:
//
//	// Check if any element in a slice of integers is non-zero.
//	ints := []int{0, 0, 0, 1, 0}
//	resultI := g.Any(ints...)
//	fmt.Println(resultI) // Output: true
//
//	// Check if any element in a slice of strings is non-empty.
//	strings := []string{"", "hello", "", ""}
//	resultS := g.Any(strings...)
//	fmt.Println(resultS) // Output: true
//
//	// Check if any element in a slice of booleans is true.
//	bools := []bool{false, false, true, false}
//	resultB := g.Any(bools...)
//	fmt.Println(resultB) // Output: true
func Any[T any](v ...T) bool {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: false}

	if len(v) == 0 {
		return false
	}

	chunkSize := len(v) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(v)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, b := range v[start:end] {
				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				if !IsEmpty(b) {
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

// IsEmpty checks if the  v of any type T is "zero value" for that type.
//
// Zero values in Go are values that the variables of respective types hold
// upon their declaration, if they do not have any explicit initialization.
//
// For example, zero value of type int is 0, for type float64 is 0.0,
// for a pointer is nil, for a string is "", for a boolean is false, etc.
//
// Example usage:
//
//	// Check if an integer variable is zero.
//	var num int
//	result := g.IsEmpty(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is zero.
//	var f float64
//	result := g.IsEmpty(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a pointer variable is nil.
//	var ptr *int
//	result := g.IsEmpty(ptr)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is empty.
//	var str string
//	result := g.IsEmpty(str)
//	fmt.Println(result) // Output: true
//
//	// Check if a boolean variable is false.
//	var flag bool
//	result := g.IsEmpty(flag)
//	fmt.Println(result) // Output: true
func IsEmpty[T any](v T) bool {
	t := reflect.TypeOf(v)
	if t == nil {
		return true
	}
	zero := reflect.Zero(t).Interface()
	return reflect.DeepEqual(v, zero)
}

// IsPointer checks if a value is a pointer.
//
// The function takes an interface{} value `v` and checks if it is
// a pointer type. It returns true if `v` is a pointer, and false
// otherwise.
//
// Example usage:
//
//	// Check if a variable holding a pointer to a string is a pointer.
//	str := "hello"
//	result := g.IsPointer(&str)
//	fmt.Println(result) // Output: true
//
//	// Check if a variable holding a string is a pointer.
//	result = g.IsPointer(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding an integer is a pointer.
//	result = g.IsPointer(10)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding nil is a pointer.
//	var ptr *int
//	result = g.IsPointer(ptr)
//	fmt.Println(result) // Output: false
func IsPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

// IsNumber checks if a value is a numeric type.
//
// The function takes an interface{} value `v` and checks if it is
// of a numeric type, including integer and floating-point types.
// It returns true if `v` is a numeric type, and false otherwise.
//
// Example usage:
//
//	// Check if an integer variable is a numeric type.
//	num := 10
//	result := g.IsNumber(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is a numeric type.
//	f := 3.14
//	result = g.IsNumber(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is a numeric type.
//	str := "hello"
//	result = g.IsNumber(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a boolean variable is a numeric type.
//	flag := true
//	result = g.IsNumber(flag)
//	fmt.Println(result) // Output: false
//
//	// Check if a slice of integers is a numeric type.
//	nums := []int{1, 2, 3}
//	result = g.IsNumber(nums)
//	fmt.Println(result) // Output: false
func IsNumber(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8,
		uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
