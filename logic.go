package g

import (
	"context"
	"reflect"
	"sync"

	"github.com/goloop/trit"
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
// Example usage:
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
//
// The function can work with Trit types, so it can get the status Unknown.
// In this case, if the function receives two possible solutions, it will
// return the second one (false):
//
//	max := g.If(trit.Unknown, 3, 2)  // Output: 2
//
// This function can accept an optional argument u, for cases of working
// with Trit types. In this case, if Trit is Unknow, the third value will
// be returned:
//
//	max := g.If(trit.Unknown, 3, 2, 1)  // Output: 1
//
// If more than 3 possible result values are passed, the others of
// the results will be ignored. The third and more likely results
// are ignored for boolean expression values.
//
//	max := g.If(trit.Unknown, 3, 2, 1, 5, 7)  // Output: 1 and 5, 7 are ignored
//	min := g.If(3 > 2, 3, 2, 1, 5, 7)  // Output: 3 and 1, 5, 7 are ignored
func If[L bool | trit.Trit, T any](e L, t, f T, u ...T) T {
	switch val := interface{}(e).(type) {
	case trit.Tritter:
		if val.IsTrue() {
			return t
		} else if len(u) != 0 && val.IsUnknown() {
			return u[0]
		}
	case bool:
		if val {
			return t
		}
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
//
// Warning: the function checks the list as the whole object, that is:
//
//	l := []bool{false, false, false}
//	g.All(l)    // Returns: true, because list is not an empty
//	            // and not an empty list is true
//	g.All(l...) // Returns: false, because not all elements
//	            // of the list are true
func All[T any](v ...T) bool {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: true}

	// If the length of the slice is less than or equal to
	// the minLoadPerGoroutine, then we do not need
	// to use goroutines.
	if l := len(v); l == 0 {
		return false
	} else if l/p < minLoadPerGoroutine {
		for _, b := range v {
			if IsFalse(b) {
				return false
			}
		}

		return true
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

				if IsFalse(b) {
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

// AllList is a synonym for the All function that accepts
// a set of elements as a slice.
//
// This function prevents accidentally passing a slice
// as a value (whole object).
func AllList[T any](v []T) bool {
	return All(v...)
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
//
// Warning: the function checks the list as the whole object, that is:
//
//	l := []bool{false, false, false}
//	g.Any(l)    // Returns: true, because list is not an empty
//	            // and not an empty list is true
//	g.Any(l...) // Returns: false, because not all elements
//	            // of the list are true
func Any[T any](v ...T) bool {
	var wg sync.WaitGroup

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &logicFoundValue{value: false}

	// If the length of the slice is less than or equal to
	// the minLoadPerGoroutine, then we do not need
	// to use goroutines.
	if l := len(v); l == 0 {
		return false
	} else if l/p < minLoadPerGoroutine {
		for _, b := range v {
			if !IsFalse(b) {
				return true
			}
		}

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

				if !IsFalse(b) {
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

// AnyList is a synonym for the Any function that accepts
// a set of elements as a slice.
//
// This function prevents accidentally passing a slice
// as a value (whole object).
func AnyList[T any](v []T) bool {
	return Any(v...)
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

	// For Trit empty state is Unknown only.
	switch val := interface{}(v).(type) {
	case trit.Tritter:
		return val.IsUnknown()
	}

	// For slices and arrays we check if the length is zero.
	kind := t.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return reflect.ValueOf(v).Len() == 0
	}

	zero := reflect.Zero(t).Interface()
	return reflect.DeepEqual(v, zero)
}

// IsFalse checks if the v of any type T is true value for that type.
//
// P.s. trit.False and trit.Unknown are false values.
func IsFalse[T any](v T) bool {
	t := reflect.TypeOf(v)
	if t == nil {
		return true
	}

	// For Trit not true states is: False and Unknown.
	switch val := interface{}(v).(type) {
	case trit.Tritter:
		// Anything that is not a true is a false, when converting
		// three-valued object to a boolean object the Unknown is False.
		return !val.IsTrue()
	}

	// For slices and arrays we check if the length is zero.
	kind := t.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return reflect.ValueOf(v).Len() == 0
	}

	zero := reflect.Zero(t).Interface()
	return reflect.DeepEqual(v, zero)
}

// IsTrue checks if the v of any type T is true value for that type.
//
// P.s. trit.False and trit.Unknown are false values.
func IsTrue[T any](v T) bool {
	// Anything that is not a true is a false, when converting
	// three-valued object to a boolean object the Unknown is False.
	return !IsFalse(v)
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
